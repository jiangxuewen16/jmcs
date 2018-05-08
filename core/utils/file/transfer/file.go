package transfer

import (
	"jmcs/core/utils/strings"
	"net"
	"math"
	"jmcs/core/utils/file"
	"os"
	"path/filepath"
)

/* socket 文件发送包*/
type SendPackage struct {
	//Type          int8   //类型 1-文件（文件信息），2-文件夹
	Size          int64  //文件总大小
	FileName      string //发送得到文件名
	Path          string //文件所在的文件夹
	RootPath      string //最终要存入到的文件路径 todo:这里有客户端控制？？？？
	MergeFileName string //待合并文件名称 
	Token         string //文件标记（用于标记每个文件，最后合并按此标记来）-> 用uuid
	Coroutine     int    //协程数量或拆分文件的数量
	BufSize       int    //单次发送数据的大小
	Position      int    //文件数据包在文件所在的位置
	Data          []byte //文件数据
}

/*socket 服务端在文件包发送成功后确认包*/
type ReceivePackage struct {
	FilePath  string //文件所在的路径 + 文件名  ，用于传输失败补传
	Token     string //文件标记（用于标记每个文件，最后合并按此标记来）-> 用uuid
	Position  int    //文件数据包在文件所在的位置
	Message   string //失败信息
	isSuccess bool   //是否发送成功
}

func (c SendPackage) Handle(conn net.Conn, rootPath string, head []byte, fileInfo file.FileInfo) {
	c.Size = fileInfo.Size
	c.FileName = fileInfo.Name
	c.Path = fileInfo.Path
	c.RootPath = rootPath
	c.MergeFileName = fileInfo.Name
	c.BufSize = 1024 * 1024
	clientTransfer := c.buildClientTransfer(conn, head)
	clientTransfer.SendFile()
}

func (c SendPackage) buildClientTransfer(conn net.Conn, head []byte) ClientTransfer {
	c.calculate() //计算些发送数据
	c.SetToken()  //file token
	fl, err := c.openfile()
	if err != nil {
		panic(err)
	}
	return ClientTransfer{c, conn, fl, head}

}

func (c SendPackage) openfile() (fl *os.File, err error) {
	fl, err = os.OpenFile(filepath.Join(c.Path, c.FileName), os.O_RDWR, 0666) //读写打开
	if err != nil {
		return
	}
	return
}

/*计算需要发送多少个包*/
func (c *SendPackage) calculate() {
	if c.BufSize > 0 {
		f := float64(c.Size) / float64(c.BufSize)
		c.Coroutine = int(math.Ceil(f))
	}
}

func (s *SendPackage) SetToken() {
	uuid := strings.Rand()
	s.Token = uuid.Hex()
}

//todo:接收数据的时候统计文件接收了多少个包
//todo:文件发送的时候，需要先在客户端储存好目录结构，上传完成一个从结构里面删除一个，还可以用来补上传

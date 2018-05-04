package transfer

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
)

//todo:接收数据的时候统计文件接收了多少个包

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type ServerTransfer struct {
	SendPackage
	receiveNum int      //接收包的数量
	conn       net.Conn //tcp连接通道
	filePath   string   //文件 + 路径
}

var receiveNum map[string]int //每个传输文件的信息

func (s *ServerTransfer) setFilePath(path string)  {
	s.filePath = path
}

/*
*   文件接收方法
*
*   con 连接成功的客户端连接
 */
func (s ServerTransfer) ReceiveFile() {
	var (
		//data    = make([]byte, 1024*1024) //用于保存接收的数据的切片
		//by      []byte
		//databuf = bytes.NewBuffer(by) //数据缓冲变量
		position = s.Position
		token    = s.Token
		//tempDir = library.Config  //todo:项目临时文件夹
		tempDir = "/temp/jmcs/"
	)

	_, ok := receiveNum[token]
	if !ok {
		tempFileName := token + "-" + strconv.Itoa(position) //	临时文件名
		filePath :=
		fmt.Println("创建临时文件：", tempFileName)
		fout, err := os.Create(tempFileName)
		if err != nil {
			fmt.Println("创建临时文件错误", tempFileName)
			return
		}
		fout.Close()

		receiveNum[s.Token] = 1
	} else {
		writeTempFileEnd(tempFileName, data[0:length])
		receiveNum[s.Token]++
	}

	receiveNum[s.Token] = 1

	fmt.Println("开始接受文件：【" + s.FileName + "】")
	j := 0 //标记接收数据的次数
	for {
		length, err := con.Read(data)
		if err != nil {
			da := databuf.Bytes()
			fmt.Printf("客户端 %v 已断开. %2d %d \n", con.RemoteAddr(), s.FileNum, len(da))
			return
		}

		if 0 == j {
			s.res = string(data[0:8])
			if "fileover" == res { //判断是否为发送结束指令，且结束指令会在第一次接收的数据中
				xienum := int(data[8])
				mergeFileName := string(data[9:length])
				go mainMergeFile(xienum, mergeFileName) //合并临时文件，生成有效文件
				res = "文件接收完成: " + mergeFileName
				con.Write([]byte(res))
				fmt.Println(mergeFileName, "文件接收完成")
				return

			} else { //创建临时文件
				fileNum = int(data[0])
				tempFileName = string(data[1:length]) + strconv.Itoa(fileNum)
				fmt.Println("创建临时文件：", tempFileName)
				fout, err := os.Create(tempFileName)
				if err != nil {
					fmt.Println("创建临时文件错误", tempFileName)
					return
				}
				fout.Close()
			}
		} else {
			// databuf.Write(data[0:length])
			writeTempFileEnd(tempFileName, data[0:length])
		}

		res = strconv.Itoa(fileNum) + " 接收完成"
		con.Write([]byte(res))
		j++
	}

}

func (s ServerTransfer) buildFileInfo()  {
	if _,ok := receiveNum[s.Token]; ok {
		return
	}
	receiveNum[s.Token] = s
}

func (s ServerTransfer) createTempFile() bool {
	tempFileName := s.Token + strconv.Itoa() //	临时文件名
	fmt.Println("创建临时文件：", tempFileName)
	fout, err := os.Create(tempFileName)
	if err != nil {
		fmt.Println("创建临时文件错误", tempFileName)
		return
	}
	fout.Close()

	receiveNum[s.Token] = 1
}

func (s ServerTransfer) createDir() bool {
	dir := s.RootPath + "/" + s.Path + "/" + s.FileName
	//todo:权限？？？？
	if err := os.MkdirAll(dir, 0777); err != nil {
		return false
	}
	s.setFilePath(dir)
	return true
}

/*
*   把数据写入指定的临时文件中
*
*   fileName    临时文件名
*   data        接收的数据
 */
func writeTempFileEnd(fileName string, data []byte) {
	// fmt.Println("追加：", name)
	tempFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// panic(err)
		fmt.Println("打开临时文件错误", err)
		return
	}
	defer tempFile.Close()
	tempFile.Write(data)
}

/*
*   根据临时文件数量及有效文件名称生成文件合并规则进行文件合并
*
*   connumber   临时文件数量
*   filename    有效文件名称
 */
func mainMergeFile(connumber int, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建有效文件错误", err)
		return
	}
	defer file.Close()

	//依次对临时文件进行合并
	for i := 0; i < connumber; i++ {
		mergeFile(filename+strconv.Itoa(i), file)
	}

	//删除生成的临时文件
	for i := 0; i < connumber; i++ {
		os.Remove(filename + strconv.Itoa(i))
	}

}

/*
*   将指定临时文件合并到有效文件中
*
*   rfilename   临时文件名称
*   wfile       有效文件
 */
func mergeFile(rfilename string, wfile *os.File) {

	// fmt.Println(rfilename, wfilename)
	rfile, err := os.OpenFile(rfilename, os.O_RDWR, 0666)
	defer rfile.Close()
	if err != nil {
		fmt.Println("合并时打开临时文件错误:", rfilename)
		return
	}

	stat, err := rfile.Stat()
	if err != nil {
		panic(err)
	}

	num := stat.Size()

	buf := make([]byte, 1024*1024)
	for i := 0; int64(i) < num; {
		length, err := rfile.Read(buf)
		if err != nil {
			fmt.Println("读取文件错误")
		}
		i += length

		wfile.Write(buf[:length])
	}

}

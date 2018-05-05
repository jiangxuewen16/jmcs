package transfer

import (
	"os"
	"runtime"
	"strconv"
	"jmcs/core/utils"
)

//todo:接收数据的时候统计文件接收了多少个包

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type ServerTransfer struct {
	SendPackage
	receiveNum int //接收包的数量
	//conn       net.Conn //tcp连接通道
	//filePath   string   //文件 + 路径
}

var fileReceiveInfos map[string]ServerTransfer //每个传输文件的信息

var tempPath = "" //todo:项目临时文件夹

/*func (s *ServerTransfer) setFilePath(path string) {
	s.filePath = path
}*/

func (s *ServerTransfer) getTempFilePath(position int) string {
	return tempPath + "/" + s.Token + "-" + strconv.Itoa(position)
}

/*创建最终文件存放目录*/
func (s *ServerTransfer) createFilePath() (string, error) {
	filePath := s.RootPath + "/" + s.Path
	ok, err := utils.PathExists(filePath)
	if !ok || err != nil {
		err := utils.MkDirAll(filePath)
		return "", err
	}
	return filePath, nil
}

/*
*   文件接收方法
*
*   con 连接成功的客户端连接
 */
func (s ServerTransfer) ReceiveFile() ReceivePackage {

	token := s.Token

	err := s.createTempFile() //创建临时文件
	if err != nil {
		return s.receiveFileFail(err)
	}

	fileReceiveInfo, ok := fileReceiveInfos[token]
	if !ok {
		s.receiveNum = 1
		fileReceiveInfos[token] = s
	} else {
		fileReceiveInfo.receiveNum ++

		//receiveNum[token].receiveNum = num + 1
		/*判断文件包是否放松完成，发送完成则合并文件*/
		if fileReceiveInfo.receiveNum >= s.Coroutine {
			dir, err := s.createFilePath()
			if err != nil { //创建最终文件存放的文件夹
				return s.receiveFileFail(err)
			}
			err = s.mainMergeFile(dir)
			if err != nil {
				return s.receiveFileFail(err)
			}
		}
	}

	return s.receiveFileSuccess()
}

/*发送失败*/
func (s ServerTransfer) receiveFileFail(err error) ReceivePackage {
	receivePackage := ReceivePackage{FilePath:s.Path + s.FileName, Token: s.Token, Position:s.Position, Message:err.Error(),isSuccess:false}
	return receivePackage
}

/*发送成功*/
func (s ServerTransfer) receiveFileSuccess() ReceivePackage {
	receivePackage := ReceivePackage{FilePath:s.Path + s.FileName, Token: s.Token, Position:s.Position, Message:"",isSuccess:true}
	return receivePackage
}

/*创建临时文件*/
func (s ServerTransfer) createTempFile() error {
	var (
		data     = s.Data
		position = s.Position
	)

	tempFileName := s.getTempFilePath(position) //	临时文件名
	fout, err := os.Create(tempFileName)
	if err != nil {
		return err
	}

	fout.Write(data) //写入数据
	fout.Close()

	return nil
}

/*
*   根据临时文件数量及有效文件名称生成文件合并规则进行文件合并
*
*   connumber   临时文件数量
*   filename    有效文件名称
 */
func (s ServerTransfer) mainMergeFile(path string) error {
	mergeFileName := s.MergeFileName
	coroutine := s.Coroutine
	filePath := path + "/" + mergeFileName
	file, err := os.Create(filePath)
	if err != nil {
		//fmt.Println("创建有效文件错误", err)
		return err
	}
	defer file.Close()

	//依次对临时文件进行合并
	for i := 0; i < coroutine; i++ {
		tempFilePath := s.getTempFilePath(i)
		err := mergeFile(tempFilePath, file)
		if err != nil {
			return err
		}
		os.Remove(tempFilePath)
	}

	//删除生成的临时文件
	/*for i := 0; i < coroutine; i++ {
		os.Remove(tempFilePath)
	}*/
	return nil
}

/*
*   将指定临时文件合并到有效文件中
*
*   rfilename   临时文件名称
*   wfile       有效文件
 */
func mergeFile(rfilename string, wfile *os.File) error {

	rfile, err := os.OpenFile(rfilename, os.O_RDWR, 0666)
	defer rfile.Close()
	if err != nil {
		//fmt.Println("合并时打开临时文件错误:", rfilename)
		return err
	}

	stat, err := rfile.Stat()
	if err != nil {
		//panic(err)
		return err
	}

	num := stat.Size()

	buf := make([]byte, 1024*1024)
	for i := 0; int64(i) < num; {
		length, err := rfile.Read(buf)
		if err != nil {
			//fmt.Println("读取文件错误")
			return err
		}
		i += length

		wfile.Write(buf[:length])
	}
	return nil
}

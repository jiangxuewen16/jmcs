package transfer

import (
	"os"
	"strconv"
	"jmcs/core/utils"
)

type ServerTransfer struct {
	SendPackage
	receiveNum int //接收包的数量
	Finished bool	//判断传输是否完成
}

const READ_BUFF = 1024 * 1024

var fileReceiveInfos = make(map[string]*ServerTransfer) //每个传输文件的信息

/*
 *   文件接收方法主程序
 */
func (s *ServerTransfer) ReceiveFile() ReceivePackage {
	s.Finished = false

	token := s.Token

	err := s.createTempFile() //创建临时文件
	if err != nil {
		return s.receiveData(err)
	}

	fileReceiveInfo, ok := fileReceiveInfos[token]
	if !ok {
		s.receiveNum = 1
		fileReceiveInfos[token] = s
	} else {
		fileReceiveInfo.receiveNum ++

		/*判断文件包是否发送完成，发送完成则合并文件*/
		if fileReceiveInfo.receiveNum >= s.Coroutine {
			dir, err := s.createFilePath()
			if err != nil { //创建最终文件存放的文件夹
				return s.receiveData(err)
			}
			err = s.mainMergeFile(dir)
			if err != nil {
				return s.receiveData(err)
			}

			s.Finished = true
		}
	}

	return s.receiveData(nil)
}

/*获取存放的临时路径*/
func (s *ServerTransfer) getTempFilePath(position int) string {
	tempPath := utils.AppConfig["temDir"]
	return tempPath.(string) + "/" + s.Token + "-" + strconv.Itoa(position)
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

/*构建文件接收信息：成功或者失败*/
func (s ServerTransfer) receiveData(err error) ReceivePackage {
	rBool := true
	message := ""

	if err != nil {
		rBool = false
		message = err.Error()
	}
	receivePackage := ReceivePackage{FilePath:s.Path + s.FileName, Token: s.Token, Position:s.Position, Message:message,isSuccess:rBool}
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
 */
func (s ServerTransfer) mainMergeFile(path string) error {
	mergeFileName := s.MergeFileName
	coroutine := s.Coroutine
	token := s.Token
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

	delete(fileReceiveInfos, token)

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

	buf := make([]byte, READ_BUFF)
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

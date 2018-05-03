package transfer

import (
	"net"
	"os"
	"bytes"
	"strconv"
	"fmt"
	"time"
)

type FileClientTransfer struct {
	conn *net.Conn `json:"conn"` //socket连接

	label string `json:"label"` //master-主服务端(中心服务器)  slave-从服务器

	FileName      string `json:"file_name"`       //待发送文件名称
	MergeFileName string `json:"merge_file_name"` //待合并文件名称
	Coroutine     int    `json:"coroutine"`       //协程数量或拆分文件的数量
	BufSize       int    `json:"buf_size"`        //单次发送数据的大小
}

func (f FileClientTransfer) TransferHandle() {

	fl, err := os.OpenFile(f.FileName, os.O_RDWR, 0666) //读写打开
	if err != nil {
		fmt.Println("userFile", err)
		return
	}

	stat, err := fl.Stat() //获取文件状态
	if err != nil {
		panic(err)
	}
	var size int64
	size = stat.Size()
	fl.Close()

	littleSize := size / int64(f.Coroutine)

	fmt.Printf("Size: %d  %d \n", size, littleSize)

	begintime := time.Now().Unix()
	//对待发送文件进行拆分计算并调用发送方法
	c := make(chan string)
	var begin int64 = 0
	for i := 0; i < f.Coroutine; i++ {
		if i == f.Coroutine-1 {
			go f.splitFile(c, i, begin, size)
			fmt.Println(begin, size, f.BufSize)
		} else {
			go f.splitFile(c, i, begin, begin+littleSize)
			fmt.Println(begin, begin+littleSize)
		}

		begin += littleSize
	}

	//同步等待发送文件的协程
	for j := 0; j < f.Coroutine; j++ {
		fmt.Println(<-c)
	}

	midtime := time.Now().Unix()
	sendtime := midtime - begintime
	fmt.Printf("发送耗时：%d 分 %d 秒 \n", sendtime/60, sendtime%60)

	f.sendMergeCommand() //发送文件合并指令及文件名
	endtime := time.Now().Unix()

	mergetime := endtime - midtime
	fmt.Printf("合并耗时：%d 分 %d 秒 \n", mergetime/60, mergetime%60)

	tot := endtime - begintime
	fmt.Printf("总计耗时：%d 分 %d 秒 \n", tot/60, tot%60)

}

/*
*   文件拆分发送方法
*
*   c               channel,用于同步协程
*   coroutineNum    协程顺序或拆分文件的顺序
*   size            发送数据的大小
*   fileName        待发送的文件名
*   mergeFileName   待合并的文件名
*   begin           当前协程拆分待发送文件中的开始位置
*   end             当前协程拆分待发送文件中的结束位置
 */
func (f FileClientTransfer) splitFile(c chan string, coroutineNum int, begin int64, end int64) {
	conn := *f.conn
	size := f.BufSize
	fileName := f.FileName

	var by [1]byte
	by[0] = byte(coroutineNum)
	var bys []byte
	databuf := bytes.NewBuffer(bys) //数据缓冲变量
	databuf.Write(by[:])
	databuf.WriteString(f.MergeFileName)
	bb := databuf.Bytes()

	in, err := conn.Write(bb) //向服务器发送当前协程的顺序，代表拆分文件的顺序
	if err != nil {
		fmt.Printf("向服务器发送数据错误: %d\n", in)
		os.Exit(0)
	}

	var msg = make([]byte, 1024)   //创建读取服务端信息的切片
	lengthh, err := conn.Read(msg) //确认服务器已收到顺序数据
	if err != nil {
		fmt.Printf("读取服务器数据错误.\n", lengthh)
		os.Exit(0)
	}
	// str := string(msg[0:lengthh])
	// fmt.Println("服务端信息：",str)

	//打开待发送文件，准备发送文件数据
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(fileName, "-文件打开错误.")
		os.Exit(0)
	}

	file.Seek(begin, 0) //设定读取文件的位置

	buf := make([]byte, size) //创建用于保存读取文件数据的切片

	var sendDtaTolNum int = 0 //记录发送成功的数据量（Byte）
	//读取并发送数据
	for i := begin; int64(i) < end; i += int64(size) {
		length, err := file.Read(buf) //读取数据到切片中
		if err != nil {
			fmt.Println("读文件错误", i, coroutineNum, end)
		}

		//判断读取的数据长度与切片的长度是否相等，如果不相等，表明文件读取已到末尾
		if length == size {
			//判断此次读取的数据是否在当前协程读取的数据范围内，如果超出，则去除多余数据，否则全部发送
			if int64(i)+int64(size) >= end {
				sendDataNum, err := conn.Write(buf[:size-int((int64(i) + int64(size) - end))])
				if err != nil {
					fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
					os.Exit(0)
				}
				sendDtaTolNum += sendDataNum
			} else {
				sendDataNum, err := conn.Write(buf)
				if err != nil {
					fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
					os.Exit(0)
				}
				sendDtaTolNum += sendDataNum
			}

		} else {
			sendDataNum, err := conn.Write(buf[:length])
			if err != nil {
				fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
				os.Exit(0)
			}
			sendDtaTolNum += sendDataNum
		}

		//读取服务器端信息，确认服务端已接收数据
		lengths, err := conn.Read(msg)
		if err != nil {
			fmt.Printf("读取服务器数据错误.\n", lengths)
			os.Exit(0)
		}
		// str := string(msg[0:lengths])
		// fmt.Println("服务端信息：",str)

	}

	fmt.Println(coroutineNum, "发送数据(Byte)：", sendDtaTolNum)

	c <- strconv.Itoa(coroutineNum) + " 协程退出"
}

/*
*   向服务端发送待合并文件的名称及合并指令
*
*   mergeFileName   待合并的文件名
*   coroutine       拆分文件的总个数
 */
func (f FileClientTransfer) sendMergeCommand() {
	conn := *f.conn
	mergeFileName := f.MergeFileName
	coroutine := f.Coroutine

	var by [1]byte
	by[0] = byte(coroutine)
	var bys []byte
	databuf := bytes.NewBuffer(bys) //数据缓冲变量
	databuf.WriteString("fileover")
	databuf.Write(by[:])
	databuf.WriteString(mergeFileName)
	cmm := databuf.Bytes()

	in, err := conn.Write(cmm)
	if err != nil {
		fmt.Printf("向服务器发送数据错误: %d\n", in)
	}

	var msg = make([]byte, 1024)
	lengthh, err := conn.Read(msg)
	if err != nil {
		fmt.Printf("读取服务器数据错误.\n", lengthh)
		os.Exit(0)
	}
	str := string(msg[0:lengthh])
	fmt.Println("传输完成（服务端信息）： ", str)
}

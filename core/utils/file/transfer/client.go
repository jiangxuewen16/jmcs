package transfer

import (
	"net"
	"os"
	"strconv"
	"fmt"
	"time"
	"bytes"
	"encoding/json"
)

//todo:文件发送的时候，需要先在客户端储存好目录结构，上传完成一个从结构里面删除一个，还可以用来补上传

type ClientTransfer struct {
	SendPackage SendPackage
	Conn net.Conn
	fileHandle *os.File
	head []byte
}

func (c ClientTransfer) SendFile() {
	var (
		coroutine = c.SendPackage.Coroutine
		bufSize   = c.SendPackage.BufSize
		size = c.SendPackage.Size
	)

	littleSize := size / int64(coroutine)

	fmt.Printf("Size: %d  %d \n", size, littleSize)

	begintime := time.Now().Unix()
	//对待发送文件进行拆分计算并调用发送方法
	ch := make(chan string)
	var begin int64 = 0
	for i := 0; i < coroutine; i++ {
		if i == coroutine-1 {
			go c.splitFile(ch, i, begin, size)
			fmt.Println(begin, size, bufSize)
		} else {
			go c.splitFile(ch, i, begin, begin+littleSize)
			fmt.Println(begin, begin+littleSize)
		}

		begin += littleSize
	}

	//同步等待发送文件的协程
	for j := 0; j < coroutine; j++ {
		fmt.Println(<-ch)
	}

	endTime := time.Now().Unix()
	sendtime := endTime - begintime
	fmt.Printf("发送耗时：%d 分 %d 秒 \n", sendtime/60, sendtime%60)

	tot := endTime - begintime
	fmt.Printf("总计耗时：%d 分 %d 秒 \n", tot/60, tot%60)

}

/*
*   文件拆分发送方法
*
*   ch               channel,用于同步协程
*   coroutineNum    协程顺序或拆分文件的顺序
*   begin           当前协程拆分待发送文件中的开始位置
*   end             当前协程拆分待发送文件中的结束位置
 */
func (c *ClientTransfer) splitFile(ch chan string, coroutineNum int, begin int64, end int64) {
	var (
		conn     = c.Conn
		size     = c.SendPackage.BufSize
		//fileName = c.FileName
		readMsg      = make([]byte, size) //创建读取服务端信息的切片
		head = c.head
		fileHandle = c.fileHandle
	)


	//打开待发送文件，准备发送文件数据

	fileHandle.Seek(begin, 0) //设定读取文件的位置

	buf := make([]byte, size) //创建用于保存读取文件数据的切片

	var sendDtaTolNum int = 0 //记录发送成功的数据量（Byte）

	var bytebuf bytes.Buffer
	bytebuf.Write(head)

	c.SendPackage.Position = coroutineNum

	//读取并发送数据
	for i := begin; int64(i) < end; i += int64(size) {
		length, err := fileHandle.Read(buf) //读取数据到切片中
		if err != nil {
			fmt.Println("读文件错误", i, coroutineNum, end)
		}

		//判断读取的数据长度与切片的长度是否相等，如果不相等，表明文件读取已到末尾
		if length == size {
			//判断此次读取的数据是否在当前协程读取的数据范围内，如果超出，则去除多余数据，否则全部发送
			if int64(i)+int64(size) >= end {
				c.SendPackage.Data = buf[:size-int((int64(i) + int64(size) - end))]

				b,_ := json.Marshal(c.SendPackage)

				bytebuf.Write(b)
				//wb := append(c.head, buf[:size-int((int64(i) + int64(size) - end))]...)
				sendDataNum, err := conn.Write(bytebuf.Bytes())
				if err != nil {
					fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
					os.Exit(0)
				}
				sendDtaTolNum += sendDataNum
			} else {
				c.SendPackage.Data = buf
				b,_ := json.Marshal(c.SendPackage)

				bytebuf.Write(b)
				sendDataNum, err := conn.Write(bytebuf.Bytes())
				if err != nil {
					fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
					os.Exit(0)
				}
				sendDtaTolNum += sendDataNum
			}

		} else {
			c.SendPackage.Data = buf[:length]
			b,_ := json.Marshal(c.SendPackage)

			bytebuf.Write(b)
			sendDataNum, err := conn.Write(bytebuf.Bytes())
			if err != nil {
				fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
				os.Exit(0)
			}
			sendDtaTolNum += sendDataNum
		}

		//读取服务器端信息，确认服务端已接收数据
		lengths, err := conn.Read(readMsg)
		if err != nil {
			fmt.Printf("读取服务器数据错误.\n", lengths)
			os.Exit(0)
		}
		// str := string(msg[0:lengths])
		// fmt.Println("服务端信息：",str)

	}

	fmt.Println(coroutineNum, "发送数据(Byte)：", sendDtaTolNum)

	ch <- strconv.Itoa(coroutineNum) + " 协程退出"
}

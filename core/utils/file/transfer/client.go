package transfer

import (
	"net"
	"os"
	"fmt"
	"time"
	"bytes"
	"encoding/json"
)

//todo:文件发送的时候，需要先在客户端储存好目录结构，上传完成一个从结构里面删除一个，还可以用来补上传

type ClientTransfer struct {
	SendPackage SendPackage
	Conn        net.Conn
	fileHandle  *os.File
	head        []byte
}

func (c ClientTransfer) SendFile() {
	var (
		sendPackage = c.SendPackage
		coroutine   = sendPackage.Coroutine
		//bufSize     = sendPackage.BufSize
		//size        = sendPackage.Size
	)

	//littleSize := size / int64(coroutine)

	//fmt.Printf("Size: %d  %d \n", size, littleSize)

	begintime := time.Now().Unix()
	//对待发送文件进行拆分计算并调用发送方法
	ch := make(chan string)
	//var begin int64 = 0
	for i := 0; i < coroutine; i++ {

		buf,err := c.getSendBytes(i)
		if err != nil {
			panic(err)
		}
		go c.send(ch,buf)
		//go c.receive()
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

func (c ClientTransfer) getSendBytes(i int) ([]byte,error) {
	var (
		sendPackage = c.SendPackage

		fileHandle = c.fileHandle
		bufSize = sendPackage.BufSize
		head = c.head

		buf = make([]byte, bufSize)
	)

	offset := i * bufSize		//偏移
	_, err := fileHandle.Seek(int64(offset), 0)
	if err != nil {
		return nil,err
	}
	_,err = fileHandle.Read(buf)
	if err != nil {
		return nil,err
	}

	bytes.TrimRight(buf, "\x00")

	sendPackage.Position = i
	sendPackage.Data = buf
	buf,err = json.Marshal(sendPackage)
	if err != nil {
		return nil,err
	}

	buf = append(head, buf...)
	return buf,nil
}

/*
*   文件拆分发送方法
*
*   ch          channel,用于同步协程
*   sendByte    需要发送的字节
 */
func (c ClientTransfer) send(ch chan string, sendByte []byte) (int, error){
	var (
		conn = c.Conn
	)

	n, err := conn.Write(sendByte)
	if err != nil {
		return 0,err
	}
	return n,err
}

/*func (c ClientTransfer) receive()  {
	var (
		conn = c.Conn
	)

	n, err := conn.Read
}*/

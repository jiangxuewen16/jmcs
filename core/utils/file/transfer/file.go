package transfer

/* socket 文件发送的第一个包约定*/
type FileFirst struct {
	Type          int8   //类型 1-文件，2-文件夹
	Size          int64  //文件总大小
	FileName      string //发送得到文件名
	MergeFileName string //待合并文件名称

	Coroutine     int    //协程数量或拆分文件的数量
	BufSize       int  //单次发送数据的大小
}

type FilePackage struct {
	BufSize int   //每个发送包的大小

	Body []byte	//文件数据字节
}

type FileEnd struct {

}
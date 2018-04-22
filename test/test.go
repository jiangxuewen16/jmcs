package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
	"path/filepath"
)

func main() {
	dirList, err := ioutil.ReadDir("C:/golang/src/jmcs/conf/")
	checkErr(err)

	for _,dir := range dirList {
		fmt.Println(dir);
	}

}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		checkErr(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

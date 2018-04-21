package main

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"strings"
	"path/filepath"
)

func main() {
	f,err := os.Open("C:/golang/src/jmcs/conf/app.yml")
	if err != nil {
		fmt.Println(err)
	}

	maps := make(map[interface{}]interface{})

	b,_ := ioutil.ReadAll(f)
	yaml.Unmarshal(b, &maps)
	fmt.Println(GetCurrentDirectory())
	fmt.Println(maps)
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

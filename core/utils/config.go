package utils

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"os"
)

type Config map[string]interface{}

var Conf = Config(make(map[string]interface{}))

var suffix = []string{"yml"}

func (c *Config) resolve(f io.Reader) {

	b, err := ioutil.ReadAll(f)
	CheckErr(err)

	yaml.Unmarshal(b, c)
}

func (c *Config) ResolveDir(dir string) {
	if ok, err := PathExists(dir); !ok {
		if err != nil {
			CheckErr(err)
		}
		return
	}

	fileNames := GetPathAllFileName(dir,suffix)

	for _,fileName := range fileNames {
		filePath := dir + fileName

		f,err := os.Open(filePath)
		CheckErr(err)

		c.resolve(f)
	}
}

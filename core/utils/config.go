package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"os"
	"strings"
	"path"
)

type Config map[string]interface{}

var Configs = make(map[string]Config) //todo:是否把总配置放这里，应不应该放到业务中去

var Suffix = []string{".yml"}		//todo:现在只支持yaml配置格式

func (c Config) Resolve(filePath string) {
	f, err := os.Open(filePath)
	CheckErr(err)

	fileName := strings.Trim(path.Base(filePath), path.Ext(filePath))
	b, err := ioutil.ReadAll(f)
	CheckErr(err)

	yaml.Unmarshal(b, &c)

	Configs[fileName] = c
}

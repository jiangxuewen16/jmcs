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

var Suffix = []string{".yml"}

func (c Config) Resolve(filePath string) {
	f, err := os.Open(filePath)
	CheckErr(err)

	fileName := strings.Trim(path.Base(filePath), path.Ext(filePath))
	b, err := ioutil.ReadAll(f)
	CheckErr(err)

	yaml.Unmarshal(b, &c)

	Configs[fileName] = c
}

/*func aa(conf *interface{}, confName string) {
	//socket配置
	confs := strings.Split(confName, ",")
	baseConfig, ok := Configs[confs[0]][confs[1]]
	if !ok {
		Conf = http{enable: false}
	}

	Conf = http{}
	err := mapstructure.Decode(baseConfig, &Conf) //解析socket配置
	utils.CheckErr(err)
}*/

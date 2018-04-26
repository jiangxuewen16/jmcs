package utils

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config map[string]interface{}

var Configs = map[string]Config{} //todo:是否把总配置放这里，应不应该放到业务中去

var Suffix = []string{".yml"}

func (c Config) Resolve(f io.Reader, configName string) {

	config := make(map[string]interface{})
	b, err := ioutil.ReadAll(f)
	CheckErr(err)

	yaml.Unmarshal(b, config)

	Configs[configName] = config
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

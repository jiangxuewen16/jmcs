package utilities

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config map[interface{}]interface{}

func (c *Config) resolve(f io.Reader) {

	b, err := ioutil.ReadAll(f)
	CheckErr(err)

	yaml.Unmarshal(b, c)
}

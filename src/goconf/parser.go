package goconf

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/CardInfoLink/log"
	"gopkg.in/yaml.v2"

	"github.com/wonsikin/dictionary/src/util"
)

// Config system config
var Config = &InstanceConfig{}

func init() {
	fileName := fmt.Sprintf("%s/config/config.yaml", util.WorkDir)
	fmt.Printf("config file:\t %s\n", fileName)

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("read config file %s error: %s\n", fileName, err)
		os.Exit(4)
	}

	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		log.Errorf("config file `%s` parser error: %s\n", fileName, err)
		os.Exit(5)
	}

}

// readConfigFile 读取配置文件的方法
func readConfigFile(filePath string) (*InstanceConfig, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("read config file error %s", err)
		return nil, err
	}

	log.Debugf("This is biz configuraton:\n\ncontent is %s", string(content))

	var config InstanceConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Errorf("unmarshal yaml file error %s", err)
		return nil, err
	}

	return &config, err
}

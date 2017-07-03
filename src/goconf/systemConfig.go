package goconf

import "github.com/CardInfoLink/log"

// InstanceConfig 配置文件的结果，存储在`config/config_***.yaml`文件里面
type InstanceConfig struct {
	Youdao struct {
		APPID     string `yaml:"appID"`
		APPSecret string `yaml:"appSecret"`
	}
	Iciba struct {
		Key string `yaml:"key"`
	}
}

// UnmarshalYAML 自定义的解析YAML的方法，用来读取配置文件
func (c *InstanceConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		Youdao struct {
			APPID     string `yaml:"appID"`
			APPSecret string `yaml:"appSecret"`
		} `yaml:"app"`

		Iciba struct {
			Key string `yaml:"key"`
		} `yaml:"iciba"`
	}

	if err := unmarshal(&aux); err != nil {
		log.Errorf("unmarshal error %s", err)
		return err
	}

	c.Youdao.APPID = aux.Youdao.APPID
	c.Youdao.APPSecret = aux.Youdao.APPSecret

	c.Iciba.Key = aux.Iciba.Key

	return nil
}

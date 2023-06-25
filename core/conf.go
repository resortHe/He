package core

import (
	"fmt"
	"go_vue/config"
	"go_vue/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// InitCore 读取yarm文件的配置
func InitCore() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml error : %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal :%v", err)
	}
	log.Println("config yamFile load Init success")
	global.Config = c
}

package core

import (
	"fmt"
	"go_vue/config"
	"go_vue/global"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitCore 读取yaml文件的配置
func InitCore() {
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
func SetYaml() error {
	bytes, err := yaml.Marshal(global.Config)

	if err != nil {

		return err
	}
	err = ioutil.WriteFile(ConfigFile, bytes, fs.ModePerm)
	if err != nil {

		return err
	}
	global.Log.Info("配置修改成功")
	return nil
}

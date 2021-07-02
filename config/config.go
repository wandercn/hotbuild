/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : config.go
#   Last Modified : 2021-07-02 13:52
#   Describe      :
#
# ====================================================*/
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

const (
	ConfFileName = ".hotbuild" // 不包括类型后缀
)

func New() *buildConfig {
	return &buildConfig{viper.New()}
}

type buildConfig struct {
	*viper.Viper
}

/* 初始化默认配置 */
func init() {
	configFile := ConfFileName + ".toml"
	_, err := os.Lstat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("create config file:", ConfFileName+".toml")
		} else {
			log.Printf("Lstat failed: %v", err)
			return

		}
	}

	base := viper.New()
	base.AddConfigPath(".")
	base.SetConfigName(ConfFileName)
	base.SetConfigType("toml")
	base.Set("buildCmd", "go build -o ./tmp_bin")
	base.Set("runCmd", "./tmp_bin")
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	base.Set("projectDir", pwd)
	base.Set("excludeDir", []string{".git", "tmp", "docs", "vendor"})
	base.SafeWriteConfig()

}

/* 读取配置 */
func (b *buildConfig) readConf() (*buildConfig, error) {
	var err error
	b.SetConfigName(ConfFileName)
	b.SetConfigType("toml")
	b.AddConfigPath(".")
	err = b.ReadInConfig()
	if err != nil {
		log.Fatalf("ReadConfig failed: %v", err)
		return nil, fmt.Errorf("ReadConfig failed: %v", err)
	}
	return b, nil
}

/* 获取编译命令 */
func (b *buildConfig) GetBuildCmd() (string, error) {
	conf, err := b.readConf()
	if err != nil {
		return "", fmt.Errorf("readConf failed: %v", err)
	}
	return conf.GetString("buildCmd"), nil
}

/* 获取运行命令 */
func (b *buildConfig) GetRunCmd() (string, error) {
	conf, err := b.readConf()
	if err != nil {
		return "", fmt.Errorf("readConf failed: %v", err)
	}
	return conf.GetString("runCmd"), nil
}

/* 获取代码项目路径 */
func (b *buildConfig) GetProjectDir() (string, error) {
	conf, err := b.readConf()
	if err != nil {
		return "", fmt.Errorf("readConf failed: %v", err)
	}
	return conf.GetString("projectDir"), nil
}

/* 获取排除目录 */
func (b *buildConfig) GetExcludeDirs() ([]string, error) {
	conf, err := b.readConf()
	if err != nil {
		return nil, fmt.Errorf("readConf failed: %v", err)
	}
	return conf.GetStringSlice("excludeDir"), nil
}

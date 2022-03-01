/*
   Copyright (c) 2021-present ffactory.org
   hotbuild is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/
package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

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
func InitConf() error {
	configFile, err := filepath.Abs(ConfFileName + ".toml")
	if err != nil {
		log.Printf("get configfile abs path failed:%v", err)
	}
	_, err = os.Lstat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("create config file:", ConfFileName+".toml")
		} else {
			log.Printf("Lstat failed: %v", err)
			return fmt.Errorf("Lstat failed: %v", err)
		}
	}

	base := viper.New()
	base.AddConfigPath(".")
	base.SetConfigName(ConfFileName)
	base.SetConfigType("toml")
	tmpPath := path.Dir("./")
	base.Set("buildCmd", "go build -o "+tmpPath+"tmp_bin")
	base.Set("runCmd", tmpPath+"mp_bin")
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Getwd faied: %v", err)
	}
	base.Set("projectDir", pwd)
	base.Set("excludeDir", []string{".git", "tmp", "docs", "vendor"})
	base.SafeWriteConfig()
	return nil

}

/* 读取配置 */
func (b *buildConfig) readConf() (*buildConfig, error) {
	var err error
	b.SetConfigName(ConfFileName)
	b.SetConfigType("toml")
	b.AddConfigPath(".")
	err = b.ReadInConfig()
	if err != nil {
		log.Fatalf("ReadConfig failed: %v\nyou need run command \"hotbuild initconf\" to initialize the configuration first!", err)
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

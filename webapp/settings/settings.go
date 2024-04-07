package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	//设置配置文件的名称和路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	//读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("read config filed err...", err)
		return err
	}

	//监听配置文件是否被修改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})

	return
}

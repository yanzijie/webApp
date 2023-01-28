package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 配置文件加载

func Init() (err error) {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")      // 在当前路径查找

	// 把配置读取到 viper对象中
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// ok 的话，说明配置文件没找到
			fmt.Println("config file not found")
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println("ReadInConfig error:", err)
		}
		return err
	}

	// 实时监控配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		// do some thing, 比如重新给配置文件变量赋值
		fmt.Println("config content has change")
	})

	return
}

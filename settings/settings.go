package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	_ "github.com/mitchellh/mapstructure"
)

// 配置文件加载

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigName("config") // 配置文件名称(无扩展名),如果有config.json和config.yaml,那就是先找到什么用什么
	//viper.SetConfigType("yaml")   // 专用于从远程加载配置文件时，指定文件类型
	viper.AddConfigPath(".") // 在当前路径查找

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
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("Unmarshal to Conf error:", err)
	}

	// 实时监控配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Println("Unmarshal to Conf when config change error:", err)
		}
	})

	return
}

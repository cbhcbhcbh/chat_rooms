package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type TomlConfig struct {
	AppName       string
	MySQL         MySQLConfig
	Log           LogConfig
	StaticPath    PathConfig
	MsgChanelType MsgChanelType
}

type LogConfig struct {
	Path  string
	Level string
}

type PathConfig struct {
	FilePath string
}

//gochannel为单机使用时默认的channel进行消息传递
//kafka是作为使用kafka作为消息队列，可以分布式扩展消息聊天程序
type MsgChanelType struct {
	ChannelType string

	KafkaHosts string
	KafkaTopic string
}

type MySQLConfig struct {
	Host        string
	Name        string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

var c TomlConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.Unmarshal(&c)
}

func GetConfig() TomlConfig {
	return c
}

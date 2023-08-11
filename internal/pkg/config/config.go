package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig *Config

type Config struct {
	Server    ServerConfig `mapstructure:"server"`
	JwtSecret string       `mapstructure:"jwt-secret"`
	MySQL     MySQLConfig  `mapstructure:"mysql"`
}

type ServerConfig struct {
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
	AppName string `mapstructure:"app-name"`
}

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Database    string `mapstructure:"database"`
	Charset     string `mapstructure:"charset"`
	ParseTime   bool   `mapstructure:"parse-time"`
	Loc         string `mapstructure:"loc"`
	MaxOpenSize int    `mapstructure:"max-open-size"`
	MaxIdleSize int    `mapstructure:"max-idle-size"`
}

// Load is a loader to load config file.
func NewViper(configFilePath string) *Config {
	resolveRealPath(configFilePath)
	// 初始化配置文件
	if err := initConfig(); err != nil {
		panic(err)
	}
	// 监控配置文件，并热加载
	watchConfig()

	return GlobalConfig
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APPLICATION")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析到struct
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		panic(err)
	}
	log.Println("The application configuration file is loaded successfully!")
	return nil
}

// 监控配置文件变动
// 注意：有些配置修改后，及时重新加载也要重新启动应用才行，比如端口
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Configuration file changed: %s, reload it", in.Name)
		// 忽略错误
		NewViper(in.Name)
	})
}

// 如果未传递配置文件路径将使用约定的环境配置文件
func resolveRealPath(path string) {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		// 设置默认的config
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
}

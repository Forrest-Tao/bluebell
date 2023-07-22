package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	*AppConfig       `mapstructure:"app"`
	*LogConfig       `mapstructure:"log"`
	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*SnowFlakeConfig `mapstructure:"snowflake"`
	*EncryptConfig   `mapstructure:"encrypt"`
	*GinConfig       `mapstructure:"gin"`
	*AuthConfig      `mapstructure:"auth"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mpastructure:"filename"`
	MaxSize    int    `mpastructure:"max_size"`
	MaxAge     int    `mpastructure:"max_age"`
	MaxBackups int    `mpastructure:"max_backups"`
}

type MySQLConfig struct {
	Host              string `mapstructure:"host"`
	User              string `mapstructure:"user"`
	Password          string `mapstructure:"password"`
	Port              int    `mapstructure:"port"`
	Db                string `mapstructure:"db"`
	MaxOpenConnection int    `mapstructure:"max_open_connection"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type AuthConfig struct {
	JWTExpire int `mapstructure:"jwt_expire"`
}

type GinConfig struct {
	Mode string `mapstructure:"mode"`
}

type EncryptConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

type SnowFlakeConfig struct {
	StartTime string `mapstructure:"start_time"`
	MachineId int64  `mapstructure:"machine_id"`
}

func Init() (err error) {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./conf/")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed: %v\n", err)
		return
	}
	//将读取到的文件反序列换到Config struct 中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal(Conf) failed: %v\n", err)
		return
	}

	//监视 config 文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("Config file has changed...\n")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal(Conf) failed: %v\n", err)
			return
		}
	})
	return
}

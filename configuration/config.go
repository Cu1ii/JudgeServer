package configuration

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type MySQLConfig struct {
	DSN                       string `mapstructure:"dsn" yaml:"dsn"`
	DefaultStringSize         int    `mapstructure:"default_string_size" yaml:"default_string_size"`
	DontSupportRenameIndex    bool   `mapstructure:"dont_support_rename_index" yaml:"dont_support_rename_index"`
	DisableDatetimePrecision  bool   `mapstructure:"disable_datetime_precision" yaml:"disable_datetime_precision"`
	DontSupportRenameColumn   bool   `mapstructure:"dont_support_rename_column" yaml:"dont_support_rename_column"`
	SkipInitializeWithVersion bool   `mapstructure:"skip_initialize_with_version" yaml:"skip_initialize_with_version"`
}

var MySQLConfiguration = (*MySQLConfig)(nil)

func GetMySQLConfig(isEnv bool) *MySQLConfig {
	if MySQLConfiguration == nil {
		if isEnv {
			MySQLConfiguration = getMySQLConfigByEnv()
		} else {
			MySQLConfiguration = getMySQLConfiguration()
		}
	}
	return MySQLConfiguration
}

func getMySQLConfigByEnv() *MySQLConfig {
	MySQLConfiguration := MySQLConfig{
		DisableDatetimePrecision:  true,
		DefaultStringSize:         256,
		DontSupportRenameColumn:   true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}
	usr := os.Getenv("mysql-user")
	pwd := os.Getenv("mysql-pwd")
	host := os.Getenv("mysql-host")
	port := os.Getenv("mysql-port")
	dbname := os.Getenv("mysql-dbname")
	MySQLConfiguration.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", usr, pwd, host, port, dbname)
	return &MySQLConfiguration
}

func getMySQLConfiguration() *MySQLConfig {
	MySQLConfiguration := MySQLConfig{}
	configPath := "resources/config/mysql.yaml"
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read mysql.yaml failed: ", err.Error())
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&MySQLConfiguration); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&MySQLConfiguration); err != nil {
		fmt.Println(err)
	}
	return &MySQLConfiguration
}

type JudgeEnvironment struct {
	SubmissionPath string `mapstructure:"submission_path" yaml:"submission_path"`
	ResolutionPath string `mapstructure:"resolution_path" yaml:"resolution_path"`
}

var JudgeEnvironmentConfiguration = (*JudgeEnvironment)(nil)

func GetJudgeEnvironmentConfig() *JudgeEnvironment {
	if JudgeEnvironmentConfiguration == nil {
		JudgeEnvironmentConfiguration = getJudgeEnvironmentConfiguration()
	}
	return JudgeEnvironmentConfiguration
}

func getJudgeEnvironmentConfiguration() *JudgeEnvironment {
	JudgeEnvironmentConfiguration := JudgeEnvironment{}
	configPath := "resources/config/judge-environment.yaml"
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read judge-environment.yaml failed: ", err.Error())
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&JudgeEnvironmentConfiguration); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&JudgeEnvironmentConfiguration); err != nil {
		fmt.Println(err)
	}
	return &JudgeEnvironmentConfiguration
}

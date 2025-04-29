package bootstrap

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
	App struct {
		Env          string `mapstructure:"env"`
		Port         int    `mapstructure:"port"`
		Version      string `mapstructure:"version"`
		FirebasePath string `mapstructure:"firebase_path"`
	} `mapstructure:"app"`

	Database struct {
		DBHost   string `mapstructure:"dbhost"`
		DBPort   string `mapstructure:"dbport"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`

	Files struct {
		Host   string `mapstructure:"host"`
		Port   string `mapstructure:"port"`
		Key    string `mapstructure:"key"`
		Bucket string `mapstructure:"bucket"`
		PathIp string `mapstructure:"path_ip"`
	} `mapstructure:"file"`

	JWT struct {
		AccessToken  string `mapstructure:"access_token"`
		RefreshToken string `mapstructure:"refresh_token"`
	} `mapstructure:"jwt"`
}

func NewEnv() *Env {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	var env Env
	if err := v.Unmarshal(&env); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	EnvRunning(env.App.Env, env.App.Port)
	return &env
}

func EnvRunning(env string, port int) {
	switch env {
	case "dev":
		log.Println("The App is running in development env on port:", port)
	case "uat":
		log.Println("The App is running in user acceptance test (UAT) env on port::", port)
	case "prd":
		log.Println("The App is running in production env on port:", port)
	}
}

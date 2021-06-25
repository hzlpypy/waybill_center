package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Model  string `yaml:"model"`
	Server struct {
		Ip     string `yaml:"ip"`
		Port   int    `yaml:"port"`
		Name   string `yaml:"name"`
		Ticker int    `yaml:"ticker"`
	} `yaml:"server"`
	Rabbitmq struct {
		Name string `yaml:"name"`
		User string `yaml:"user"`
		Pwd  string `yaml:"pwd"`
	} `yaml:"rabbitmq"`
	Mysql struct {
		Host                                     string `yaml:"host"`
		Port                                     int    `yaml:"port"`
		User                                     string `yaml:"user"`
		Pwd                                      string `yaml:"pwd"`
		DBName                                   string `yaml:"db_name"`
		Charset                                  string `yaml:"charset"`
		ConnMaxLifetime                          int    `yaml:"conn_max_lifetime"`
		MaxIdleConns                             int    `yaml:"max_idle_conns"`
		MaxOpenConns                             int    `yaml:"max_Open_conns"`
		DisableForeignKeyConstraintWhenMigrating bool   `yaml:"disable_foreign_key_constraint_when_migrating"`
	} `yaml:"mysql"`
	Log struct {
		AccessPath string `yaml:"access_path"`
		ErrorPath  string `yaml:"error_path"`
	} `yaml:"log"`
	Etcd struct {
		Ttl int `yaml:"ttl"`
	} `yaml:"etcd"`
	OrderCenter struct{
		Queue string `yaml:"queue"`
	} `yaml:"order_center"`
}

func NewConfig() *Config {
	fmt.Println(os.Getwd())
	bytes, err := ioutil.ReadFile("cmd/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

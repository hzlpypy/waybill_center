package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
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
		Name                                     string `yaml:"name"`
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
		Ttl  int    `yaml:"ttl"`
		Ip   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"etcd"`
	OrderCenter struct {
		Consumer struct {
			Queue        string `yaml:"queue"`
			Exchange     string `yaml:"exchange"`
			ExchangeType string `yaml:"exchange_type"`
			RoutingKey   string `yaml:"routingKey"`
			ContentType  string `yaml:"contentType"`
		} `yaml:"consumer"`
		Dead struct {
			Queue        string `yaml:"queue"`
			Exchange     string `yaml:"exchange"`
			ExchangeType string `yaml:"exchange_type"`
			RoutingKey   string `yaml:"routingKey"`
			ContentType  string `yaml:"contentType"`
		} `yaml:"dead"`
	} `yaml:"order_center"`
}

func currentDir() string {
	_, filename, _, _ := runtime.Caller(0) // get current filepath in runtime
	return path.Dir(filename)
}

func NewConfig() *Config {
	bytes, err := ioutil.ReadFile(path.Join(currentDir(), "config.yaml"))
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

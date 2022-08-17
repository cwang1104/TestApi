package config

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg                                           *ini.File
	Db_driver, Db_source                          string
	DbMaxOpenConns, DbMaxIdleConns, DbMaxLifeTime int

	EtcdSchema  string
	EtcdAddress string

	LoginPort, UserPort, GmPort string

	LoginName, UserName, GmName string

	ServerIp string
)

func init() {
	var err error
	Cfg, err = ini.Load("D:\\program\\exec\\golang\\TestApi\\config\\app.ini")
	if err != nil {
		log.Fatalln("config file load failed err", err)
		return
	}
	LoadDataBase()
	LoadEtcd()
	LoadRpcPort()
	LoadrpcRegisterName()
	LoadServer()

	logs.Info("--config load successful")

}

func LoadDataBase() {
	sec, err := Cfg.GetSection("DATABASE")
	if err != nil {
		log.Fatalln("getSection database err", err)
		return
	}
	Db_driver = sec.Key("db_driver").String()
	Db_source = sec.Key("db_source").String()
	DbMaxOpenConns, _ = sec.Key("dbMaxOpenConns").Int()
	DbMaxIdleConns, _ = sec.Key("dbMaxIdleConns").Int()
	DbMaxLifeTime, _ = sec.Key("dbMaxLifeTime").Int()
}

func LoadEtcd() {
	sec, err := Cfg.GetSection("ETCD")
	if err != nil {
		log.Fatalln("getSection etcd err", err)
		return
	}
	EtcdSchema = sec.Key("etcdSchema").String()
	EtcdAddress = sec.Key("etcdAddress").String()
}

func LoadRpcPort() {
	sec, err := Cfg.GetSection("rpcPort")
	if err != nil {
		log.Fatalln("getSection rpcPort err", err)
		return
	}
	LoginPort = sec.Key("loginPort").String()
	UserPort = sec.Key("userPort").String()
	GmPort = sec.Key("gmPort").String()
}

func LoadrpcRegisterName() {
	sec, err := Cfg.GetSection("rpcRegisterName")
	if err != nil {
		log.Fatalln("getSection rpcRegisterName err", err)
		return
	}
	LoginName = sec.Key("loginName").String()
	UserName = sec.Key("userName").String()
	GmName = sec.Key("gmName").String()
}

func LoadServer() {
	sec, err := Cfg.GetSection("SERVER")
	if err != nil {
		logs.Error("getSection SERVER err", err)
		return
	}

	ServerIp = sec.Key("server_ip").String()
}

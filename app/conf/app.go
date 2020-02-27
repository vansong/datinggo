package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf *Config
)

// Config struct.
type Config struct {
	AppMode	string
	AppName string
	AppPort string

	Database *Database
	Redis *Redis

	Weixin *Weixin
}

// Database conf.
type Database struct {
	Driver		string
	Host 		string
	Port 		int
	Username	string
	Password	string
	Database	string
}

func (d *Database) getDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Database)
}

// Redis conf.
type Redis struct {
	Host 	string
	Port 	int
	Auth	string
}

type Weixin struct {
	VerifyToken string
}

func init() {
	err := ConfInit();
	if err != nil {
		panic(err)
	}
}

// Init conf.
func ConfInit() (err error) {
	confPath = "conf.toml"
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

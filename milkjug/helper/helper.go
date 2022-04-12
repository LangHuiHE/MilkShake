package helper

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

func GetJwtKey() string {
	s, err := beego.AppConfig.String("jwtkey")
	if err != nil {
		panic(err)
	}
	return s
}

func GetJwtTimeOut() int {
	timeout, err := beego.AppConfig.Int("tokenexp")
	if err != nil {
		log.Println(err)
		timeout = 60
	}
	return timeout
}

func GetRedisHost() string {
	add, err := beego.AppConfig.String("redis_address")
	if err != nil {
		log.Println(err)
	}
	return add
}

func GetRedisDB() string {
	db, err := beego.AppConfig.String("redis_database")
	if err != nil {
		log.Println(err)
	}
	return db
}

func GetRedisPsw() string {
	psw, err := beego.AppConfig.String("redis_password")
	if err != nil {
		log.Println(err)
	}
	return psw
}

package redis

import (
	"fmt"
	"log"
	"milkshake/helper"
	"time"

	"github.com/beego/beego/v2/adapter/cache"
	_ "github.com/beego/beego/v2/adapter/cache/redis"
	"github.com/beego/beego/v2/core/berror"
)

var redisCache cache.Cache

type DelayJob struct {
	Key   string
	Start time.Time
	Delay time.Duration
}

var DelayQueue chan DelayJob

func DelayQueueWatcher() {
	for {
		for j := range DelayQueue {
			go func(job DelayJob) {
				now := time.Now()
				excuteTime := job.Start.Add(job.Delay)
				if now.Before(excuteTime) {
					time.Sleep(time.Until(excuteTime))
				}
				DelKey(job.Key)
			}(j)
		}
	}
}

func init() {
	redisHost := helper.GetRedisHost()
	dataBase := helper.GetRedisDB()
	password := helper.GetRedisPsw()
	config := fmt.Sprintf(`{"key":"demo","conn":"%s","dbNum":"%s","password":"%s"}`, redisHost, dataBase, password)
	var err error
	redisCache, err = cache.NewCache("redis", config)
	if err != nil || redisCache == nil {
		errMsg := "failed to init redis"
		errorCode, _ := berror.FromError(err)
		berror.Error(errorCode, errMsg)
		log.Panicln(err)
		panic(errMsg)
	}

	DelayQueue = make(chan DelayJob)
	go DelayQueueWatcher()
}

// if time_optional = 0 then use the same timeout as JWT
func SetStr(key, value string, time_optional time.Duration) (err error) {
	if time_optional == 0 {
		time_optional = time.Minute * time.Duration(helper.GetJwtTimeOut())
	}
	err = redisCache.Put(key, value, time_optional)
	if err != nil {
		errorCode, _ := berror.FromError(err)
		berror.Error(errorCode, fmt.Sprint("set key:", key, ",value:", value))
	}
	return
}

func GetStr(key string) string {
	v := redisCache.Get(key)
	if v != nil {
		return string(v.([]byte)) // The conversion here is very important ,Get The return is interface
	} else {
		log.Println("Can't get key: ", key)
		return ""
	}
}

func GetByte(key string) (value []byte) {
	if redisCache.Get(key) != nil {
		return redisCache.Get(key).([]byte)
	}
	return nil
}

func DelKey(key string) (err error) {
	return redisCache.Delete(key)
}

func CreateDelayJob(key string, t time.Duration) {
	j := DelayJob{Key: key, Start: time.Now(), Delay: t}
	DelayQueue <- j
}

func SwapKey(old, new string, t time.Duration) error {
	// log.Println("Swapkey.GetStr OLDkey:", GetStr(old))
	if err := SetStr(new, GetStr(old), 0); err == nil {
		// log.Println("new key is same as old key", old == new)
		// log.Println("Swapkey.GetStr Newkey:", GetStr(new))
		CreateDelayJob(old, t)
		return nil
	} else {
		log.Println(err)
		return err
	}
}

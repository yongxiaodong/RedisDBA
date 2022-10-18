package pkg

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

var c Config

func init() {
	file := fmt.Sprintf("%s/../conf/config.yml", GetExcPath())
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Read Config File Error: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.GetConnectAddr(),
		Password: c.Password,
		DB:       c.DB,
		PoolSize: 500,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Connect Redis %s Failed", c.GetConnectAddr())
		return err
	}
	log.Printf("Connect Redis %s Successed. DB: %v", c.GetConnectAddr(), c.DB)
	return
}

func (c Config) GetConnectAddr() (connectAddr string) {
	return fmt.Sprintf("%s:%d", c.Addr, c.Port)
}

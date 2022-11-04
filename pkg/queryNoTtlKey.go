package pkg

import (
	"bufio"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"sync"
	"time"
)

var keysCh chan string
var rdb *redis.Client
var wg sync.WaitGroup
var ctx context.Context
var keysCount int64
var keysProcessCount int
var lck sync.Mutex

// var pullKeysCount int64
// var pipeQueryCount int
var noTtlKey int

func ttlIsPermanment(t int64) bool {
	if t == -1 {
		return true
	}
	return false
}

// use pipe get TTL
func getKeysTtl(temp chan []string, f *os.File) {
	//wg.Add(1)
	defer wg.Done()
	pipe := rdb.Pipeline()
	write := bufio.NewWriter(f)
	defer write.Flush()
	for v := range temp {
		for _, keys := range v {
			lck.Lock()
			keysProcessCount += 1
			lck.Unlock()
			pipe.TTL(ctx, keys)
		}
		res, err := pipe.Exec(ctx)
		if err != nil {
			log.Println(err)
		}
		lck.Lock()
		for _, v := range res {
			t := v.(*redis.DurationCmd).Val().Nanoseconds()
			if IsPermanment := ttlIsPermanment(t); IsPermanment == true {
				noTtlKey = noTtlKey + 1
				c := fmt.Sprintf("%v", v.Args()[1])
				write.WriteString(c + "\n")
			}
		}
		lck.Unlock()
		write.Flush()
	}
}

func processStdout() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("Query Queue Remaining: %v, Got count: %v, Processed: %v, noTTL keys Total: %v\n", len(keysCh), keysCount, keysProcessCount, noTtlKey)
	}
}

func QueryNoTtlKey() {
	keysProcessCount = 0
	noTtlKey = 0
	ctx = context.Background()
	keysCh = make(chan string, 200000)
	go processStdout() // print info
	keysGB := make(chan []string, 20000)
	go keysGroupBy(keysCh, keysGB) // wg.Done

	var f *os.File
	f, err := openResultFile("noTTL.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < c.ConsumerNum; i++ {
		wg.Add(1)
		go getKeysTtl(keysGB, f)
	}
	getKeysToCh(keysCh)
	wg.Wait()
	f.Close()
	fmt.Printf("Result= Queue Remaining: %v, keys Total: %v, NoTTL keys Total: %v \n", len(keysCh), keysCount, noTtlKey)
	log.Printf("scan completed. result: %s/result/", GetExcPath())
}

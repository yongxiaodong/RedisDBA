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

func getKeysToCh() {
	var cursor uint64
	var keys []string
	var err error
	wg.Add(1)
	defer wg.Done()
	for {
		keys, cursor, err = rdb.Scan(ctx, cursor, "*", c.PullKeysCount).Result()
		if err != nil {
			panic(err)
		}
		for _, v := range keys {
			lck.Lock()
			keysCount += 1
			lck.Unlock()
			keysCh <- v
		}
		if cursor == 0 {
			fmt.Println("Product: Get Keys End...keys count: ", keysCount)
			return
		}
	}
}

func ttlIsPermanment(t int64) bool {
	if t == -1 {
		return true
	}
	return false
}

// use pipe get TTL
func getKeysTtl(temp chan []string) {
	wg.Add(1)
	defer wg.Done()
	var f *os.File
	pipe := rdb.Pipeline()
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
		f, err = openResultFile()
		write := bufio.NewWriter(f)
		if err != nil {
			panic(err)
		}
		for _, v := range res {
			t := v.(*redis.DurationCmd).Val().Nanoseconds()
			if IsPermanment := ttlIsPermanment(t); IsPermanment == true {
				lck.Lock()
				noTtlKey = noTtlKey + 1
				lck.Unlock()
				c := fmt.Sprintf("%v", v.Args()[1])
				write.WriteString(c + "\n")
			}
		}
		write.Flush()
		f.Close()
	}
}

func openResultFile() (f *os.File, err error) {
	workPath := GetExcPath()
	filePathRoot := fmt.Sprintf("%s/result", workPath)
	filePath := fmt.Sprintf("%s/noTTL.txt", filePathRoot)
	err = os.MkdirAll(filePathRoot, 0644)
	if err != nil {
		return nil, err
	}
	f, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	return f, err
	//	Warn: f.Close() outside
}

func keysGroupBy(keysch chan string, keysGB chan []string) {
	log.Println("start chan group by")
	wg.Add(1)
	defer wg.Done()
	var temp []string
	for {
		select {
		case key := <-keysch:
			temp = append(temp, key)
			if len(temp) >= c.PipeQueryCount {
				keysGB <- temp
				temp = nil
			}
		case <-time.After(time.Second * 3):
			keysGB <- temp
			close(keysGB)
			log.Printf("End...Close Goup By chin")
			return
		}
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
	for i := 0; i < c.ConsumerNum; i++ {
		go getKeysTtl(keysGB)
	}
	getKeysToCh()
	wg.Wait()
	fmt.Printf("Result= Queue Remaining: %v, keys Total: %v, NoTTL keys Total: %v \n", len(keysCh), keysCount, noTtlKey)
}

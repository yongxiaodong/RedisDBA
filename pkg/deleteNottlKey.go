package pkg

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var delCountProcess int64

//func checkMd5(s string) (matchRes bool) {
//	res, _ := regexp.MatchString("([a-f\\d]{32}|[A-F\\d]{32})", s)
//	return res
//}

func DelNoTTL(ch chan []string) {
	wg.Add(1)
	defer wg.Done()
	pipe := rdb.Pipeline()
	for n := range ch {
		for _, v := range n {
			//if v != "" && checkMd5(v) == true {
			if v != "" {
				pipe.Del(ctx, v)
			}
		}
		res, err := pipe.Exec(ctx)
		if err != nil {
			log.Println(err)
		}
		for _ = range res {
			lck.Lock()
			delCountProcess += 1
			lck.Unlock()
		}
	}
}
func processStdoutDel() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Printf("Deleted: %v + \n", delCountProcess)
	}
}

func DelNoTTLPre() {
	ctx = context.Background()
	filePath := fmt.Sprintf("%s/result/noTTL.txt", GetExcPath())
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("read nottl file error. Please execute ./redisDBA nottl first")
		panic(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	// use redis pipe
	filech := make(chan string, 2000)
	fileGBch := make(chan []string, 2000)

	go keysGroupBy(filech, fileGBch)
	go processStdoutDel()
	for i := 0; i <= c.ConsumerNum; i++ {
		go DelNoTTL(fileGBch)
	}
	func() {
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			filech <- string(a)
		}
	}()
	wg.Wait()
	fmt.Printf("Deleted: %v +\n", delCountProcess)
}

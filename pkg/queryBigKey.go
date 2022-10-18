package pkg

import (
	"bufio"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"sort"
	"time"
)

var bigKeyCount int

type KeyMem struct {
	Name string
	Mem  int64
}

//MEMORY usage java_InvitationCode_set_new [Byte]

func getTop50(tempList []KeyMem) (t []KeyMem) {
	if len(tempList) >= 49 {
		return tempList[:49]
	} else {
		return tempList
	}
}

func bigKeyProcessStdout() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Printf("Scan Big Key, Got count: %v, Has been completed: %v \n", keysCount, bigKeyCount)
	}
}

func getKeysUseMem(temp chan []string) {
	//var top50L []KeyMem
	go bigKeyProcessStdout()
	top50L := make([]KeyMem, 0)
	wg.Add(1)
	defer wg.Done()
	pipe := rdb.Pipeline()
	for v := range temp {
		var tempL []KeyMem
		for _, keys := range v {
			pipe.MemoryUsage(ctx, keys)
		}
		res, err := pipe.Exec(ctx)
		if err != nil {
			log.Println(err)
		}
		for _, v := range res {
			name := fmt.Sprintf("%s", v.(*redis.IntCmd).Args()[2])
			mem := v.(*redis.IntCmd).Val()
			tempL = append(tempL, KeyMem{name, mem})
			bigKeyCount += 1
		}
		// sort revert
		sort.Slice(tempL, func(i, j int) bool {
			return tempL[i].Mem >= tempL[j].Mem
		})
		r := getTop50(tempL)
		for _, v := range r {
			top50L = append(top50L, v)
		}
		sort.Slice(top50L, func(i, j int) bool {
			return top50L[i].Mem >= top50L[j].Mem
		})
		if len(top50L) > 49 {
			top50L = top50L[:50]
		}
	}
	var f *os.File
	var err error
	f, err = openResultFile("BigKey.txt")
	if err != nil {
		panic(err)
	}
	write := bufio.NewWriter(f)
	for _, v := range top50L {
		s := fmt.Sprintf("%s %d", v.Name, v.Mem)
		write.WriteString(s + "\n")
	}
	write.Flush()
	defer f.Close()

}

func BigKeyTOP() {
	var keysch chan string
	var keysGB chan []string
	ctx = context.Background()
	keysch = make(chan string)
	keysGB = make(chan []string)
	go keysGroupBy(keysch, keysGB)
	go getKeysUseMem(keysGB)
	getKeysToCh(keysch)
	wg.Wait()
}

package pkg

import (
	"RedisDBA/cmd"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueryNoTtlKey(t *testing.T) {
	wg := sync.WaitGroup{}
	cmd.InitConfig()
	err := InitClient(c)
	if err != nil {
		panic(err)
	}
	var keysch chan string
	var keysGB chan []string
	keysch = make(chan string)
	ctx = context.Background()
	keysGB = make(chan []string)
	go keysGroupBy(keysch, keysGB)
	go getKeysUseMem(keysGB)
	getKeysToCh(keysch)
	wg.Wait()
	fmt.Println("wait")
	time.Sleep(time.Second * 20)

	//ta := make(chan []string, 10)
	//s := []string{"a", "b"}
	//s2 := []string{"a", "c"}
	//ta <- s
	//ta <- s2
	//go func() {
	//	close(ta)
	//}()
	//getKeysUseMem(ta)
}

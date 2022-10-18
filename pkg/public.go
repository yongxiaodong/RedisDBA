package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

func getKeysToCh(chName chan string) {
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
			chName <- v
		}
		if cursor == 0 {
			fmt.Println("Product: Get Keys End...keys count: ", keysCount)
			return
		}
	}
}

func openResultFile(filename string) (f *os.File, err error) {
	workPath := GetExcPath()
	filePathRoot := fmt.Sprintf("%s/result", workPath)
	//ClearDir(filePathRoot)
	ClearFile(filePathRoot)
	filePath := fmt.Sprintf("%s/%s", filePathRoot, filename)
	err = os.MkdirAll(filePathRoot, 0644)
	if err != nil {
		return nil, err
	}
	f, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
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

//func ClearDir(path string) {
//	_, err := os.Stat(path)
//	if err == nil {
//		os.RemoveAll(path)
//	}
//}

func ClearFile(path string) {
	_, err := os.Stat(path)
	if err == nil {
		os.Remove(path)
	}
}

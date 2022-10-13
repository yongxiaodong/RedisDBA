package main

import (
	"RedisDBA/pkg"
	"flag"
	"log"
)

func main() {
	var action string
	flag.StringVar(&action, "a", "", "nottl/delnottl_md5")
	flag.Parse()
	err := pkg.InitClient()
	if err != nil {
		panic(err)
	}
	if action == "nottl" {
		pkg.QueryNoTtlKey()
	} else if action == "delnottl_md5" {
		pkg.DelNoTTLPre()
	} else {
		log.Println("exit....")
	}

}

package main

import (
	"flag"
	"log"

	"github.com/atotto/clipboard"
	"github.com/chirichan/rice"
)

var (
	length = flag.Int("length", 16, "生成的密码长度，【6, 2048】")
	level  = flag.Int("level", 4, "生成的密码强度等级，数字越大，强度越高，【1, 4】")
)

func main() {

	flag.Parse()

	s, err := rice.FullPassword(*level, *length)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	if err := clipboard.WriteAll(s); err != nil {
		log.Fatalf("err: %v\n", err)
	}
}

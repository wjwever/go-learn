package app

import (
	"log"
	"runtime"
)

func CheckErr(err error) {
	if err != nil {
		_, _, line, _ := runtime.Caller(1)
		log.Fatalf("error happens: %v \nline:%v", err, line)
	}
}

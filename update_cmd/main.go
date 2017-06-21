package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetPrefix("[i] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func usage() {
	fmt.Println("Usage: update_cmd: check update for given github released workflow.")
	fmt.Println("")
	fmt.Println("\tupdate_cmd command [args]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("\tcheck \t检查是否有版本更新")
	fmt.Println("\tupdate\t升级到最新版本")
	fmt.Println("")
	os.Exit(1)
}

func main() {
	if (len(os.Args)) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "check":
		check_updates_cmd(os.Args[2:])
	case "update":
		updates(os.Args[2:])
	default:
		usage()
	}
}

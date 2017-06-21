package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	fset  = flag.NewFlagSet("update_cmd", flag.PanicOnError)
	owner = fset.String("owner", "", "Github repo owner")
	repo  = fset.String("repo", "", "Github repo")
	vers  = fset.String("version", "", "Current version of the workflow")
)

func init() {
	log.SetPrefix("[i] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func usage() {
	fmt.Println("Usage: update_cmd: check update for given github released workflow.")
	fmt.Println("")
	fmt.Println("\tupdate_cmd command -repo {repo} -owner {owner} -version {version}")
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

	fset.Parse(os.Args[2:])
	if *owner == "" || *repo == "" || *vers == "" {
		usage()
	}

	switch os.Args[1] {
	case "check":
		check_updates_cmd()
	case "update":
		updates()
	default:
		usage()
	}
}

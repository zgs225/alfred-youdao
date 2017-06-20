package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zgs225/alfred-youdao/alfred"
)

type updateInfo struct {
	Updates bool   `json:"updates"`
	Version string `json:"version"`
	Time    int64  `json:"time"`
}

func check_updates_cmd(args []string) {
	f := flag.NewFlagSet("check_update", flag.PanicOnError)
	owner := f.String("owner", "", "Github repo owner")
	repo := f.String("repo", "", "Github repo")
	vers := f.String("version", "", "Current version of the workflow")
	f.Parse(args)

	if *owner == "" {
		f.PrintDefaults()
		os.Exit(1)
	}

	if *repo == "" {
		f.PrintDefaults()
		os.Exit(1)
	}

	if *vers == "" {
		f.PrintDefaults()
		os.Exit(1)
	}

	version, err := alfred.ParseVersion(*vers)
	if err != nil {
		log.Println("版本解析错误: ", err)
		os.Exit(1)
	}

	check_updates(*owner, *repo, version)
}

func check_updates(owner, repo string, version *alfred.Version) {
	g := &alfred.GithubProvider{owner, repo}
	u := &alfred.Updater{
		V: version,
		P: g,
	}
	finfo := updateInfo{
		Time: time.Now().Unix(),
	}
	b := new(bytes.Buffer)

	if u.CanUpdate() {
		text := fmt.Sprintf("可以升级到: %v", u.E.V)
		notify(text, repo, "", "default")
		finfo.Updates = true
		finfo.Version = u.E.V.String()
	}
	if err := json.NewEncoder(b).Encode(&finfo); err != nil {
		panic(err)
	}
	f, err := os.Create(".alfred_updates")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(b.Bytes())
	if err != nil {
		panic(err)
	}
}

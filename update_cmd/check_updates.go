package main

import (
	"bytes"
	"encoding/json"
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

func (info *updateInfo) Dumps() error {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(info); err != nil {
		return err
	}
	f, err := os.Create(".alfred_updates")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func check_updates_cmd() {
	version, err := alfred.ParseVersion(*vers)
	if err != nil {
		log.Println("版本解析错误: ", err)
		os.Exit(1)
	}

	check_updates(*owner, *repo, version)
}

func check_updates(owner, repo string, version *alfred.Version) {
	log.Printf("Checking updates for %s/%s...\n", owner, repo)
	g := &alfred.GithubProvider{owner, repo}
	u := &alfred.Updater{
		V: version,
		P: g,
	}
	finfo := updateInfo{
		Time: time.Now().Unix(),
	}
	if u.CanUpdate() {
		// Do not display notification
		// text := fmt.Sprintf("可以升级到: %v", u.E.V)
		// notify(text, repo, "", "default")
		finfo.Updates = true
		finfo.Version = u.E.V.String()
	}
	if err := finfo.Dumps(); err != nil {
		panic(err)
	}
}

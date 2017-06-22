package main

import (
	"encoding/json"
	"github.com/zgs225/alfred-youdao/alfred"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	VERSION = "1.3.1"
	OWNER   = "zgs225"
	REPO    = "alfred-youdao"
	F_INFO  = ".alfred_updates"
)

var (
	_info *updateInfo
)

func init() {
	_info, _ = loadUpdateInfo()
}

type updateInfo struct {
	Updates bool   `json:"updates"`
	Version string `json:"version"`
	Time    int64  `json:"time"`
}

func (i *updateInfo) CheckAvailable() bool {
	t := time.Unix(i.Time, 0)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return t.Before(today)
}

func canUpdates() bool {
	if _info == nil {
		return false
	}
	currentv, _ := alfred.ParseVersion(VERSION)
	version, err := alfred.ParseVersion(_info.Version)
	if err != nil {
		return false
	}
	return _info.Updates && version.After(currentv)
}

func checkUpdate() {
	if checkAvailable() {
		cmd := exec.Command("./update_cmd", "check", "--owner", OWNER, "--repo", REPO, "--version", VERSION)
		err := cmd.Start()
		if err != nil {
			log.Println("Check update error: ", err)
		}
	}
}

func doUpdate() {
	cmd := exec.Command("./update_cmd", "update", "--owner", OWNER, "--repo", REPO, "--version", VERSION)
	err := cmd.Start()
	if err != nil {
		log.Println("Check update error: ", err)
	}
}

// 每天检查一次更新
func checkAvailable() bool {
	if _info == nil {
		return true
	} else {
		return _info.CheckAvailable()
	}
}

// Load .alfred_updates
func loadUpdateInfo() (*updateInfo, error) {
	f, err := os.Open(F_INFO)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var info updateInfo

	if err := json.NewDecoder(f).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}

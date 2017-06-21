package main

import (
	"log"
	"os"

	"github.com/zgs225/alfred-youdao/alfred"
)

func updates() {
	version, err := alfred.ParseVersion(*vers)
	if err != nil {
		log.Println("版本解析错误: ", err)
		os.Exit(1)
	}
	g := &alfred.GithubProvider{*owner, *repo}
	u := &alfred.Updater{
		V: version,
		P: g,
	}

	if err := u.Update(); err != nil {
		log.Println("更新失败: ", err)
		os.Exit(1)
	} else {
		log.Println("更新成功")
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zgs225/alfred-youdao/alfred"
	"github.com/zgs225/youdao"
)

const (
	APPID     = "2f871f8481e49b4c"
	APPSECRET = "CQFItxl9hPXuQuVcQa5F2iPmZSbN0hYS"
)

func init() {
	log.SetPrefix("[I] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println(os.Args)

	c := &youdao.Client{
		AppID:     APPID,
		AppSecret: APPSECRET,
	}
	r, err := c.Query(strings.Join(os.Args[1:], " "))

	if err != nil {
		panic(err)
	}

	items := alfred.NewResult()

	if r.Basic != nil {
		item := alfred.ResultElement{
			Valid:    true,
			Title:    r.Basic.Explains[0],
			Subtitle: r.Basic.Phonetic,
		}
		items.Append(&item)
	}

	if r.Translation != nil {
		item := alfred.ResultElement{
			Valid:    true,
			Title:    (*r.Translation)[0],
			Subtitle: "翻译结果",
		}
		items.Append(&item)
	}

	if r.Web != nil {
		items.Append(&alfred.ResultElement{
			Valid:    true,
			Title:    "网络释义",
			Subtitle: "有道词典",
		})

		for _, elem := range *r.Web {
			items.Append(&alfred.ResultElement{
				Valid:    true,
				Title:    elem.Key,
				Subtitle: elem.Value[0],
			})
		}
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(items)
	fmt.Print(b.String())
}

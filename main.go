package main

import (
	"log"
	"os"
	"strings"

	"github.com/zgs225/alfred-youdao/alfred"
	"github.com/zgs225/youdao"
)

const (
	APPID     = "2f871f8481e49b4c"
	APPSECRET = "CQFItxl9hPXuQuVcQa5F2iPmZSbN0hYS"
	MAX_LEN   = 255

	UPDATECMD = "alfred-youdao:update"
)

func init() {
	log.SetPrefix("[i] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println(os.Args)
	checkUpdate()

	client := &youdao.Client{
		AppID:     APPID,
		AppSecret: APPSECRET,
	}
	agent := newAgent(client)
	q := strings.TrimSpace(strings.Join(os.Args[1:], " "))
	items := alfred.NewResult()

	if q == UPDATECMD {
		doUpdate()
		items.Append(&alfred.ResultElement{
			Valid:    true,
			Title:    "正在更新中...",
			Subtitle: "有道词典 for Alfred",
		})
		items.End()
	}

	if len(q) > 255 {
		items.Append(&alfred.ResultElement{
			Valid:    false,
			Title:    "错误: 最大查询字符数为255",
			Subtitle: q,
		})
		items.End()
	}

	r, err := agent.Query(q)
	if err != nil {
		panic(err)
	}

	if r.Basic != nil {
		item := alfred.ResultElement{
			Valid:    true,
			Title:    r.Basic.Explains[0],
			Subtitle: r.Basic.Phonetic,
			Arg:      r.Basic.Explains[0],
			Mods: map[string]*alfred.ModElement{
				alfred.Mods_Shift: &alfred.ModElement{
					Valid:    true,
					Arg:      toYoudaoDictUrl(q),
					Subtitle: "回车键打开词典网页",
				},
			},
		}
		items.Append(&item)
	}

	if r.Translation != nil {
		item := alfred.ResultElement{
			Valid:    true,
			Title:    (*r.Translation)[0],
			Subtitle: "翻译结果",
			Arg:      (*r.Translation)[0],
			Mods: map[string]*alfred.ModElement{
				alfred.Mods_Shift: &alfred.ModElement{
					Valid:    true,
					Arg:      toYoudaoDictUrl(q),
					Subtitle: "回车键打开词典网页",
				},
			},
		}
		items.Append(&item)
	}

	if r.Web != nil {
		items.Append(&alfred.ResultElement{
			Valid:    true,
			Title:    "网络释义",
			Subtitle: "有道词典 for Alfred",
		})

		for _, elem := range *r.Web {
			items.Append(&alfred.ResultElement{
				Valid:    true,
				Title:    elem.Key,
				Subtitle: elem.Value[0],
				Arg:      elem.Key,
				Mods: map[string]*alfred.ModElement{
					alfred.Mods_Shift: &alfred.ModElement{
						Valid:    true,
						Arg:      toYoudaoDictUrl(q),
						Subtitle: "回车键打开词典网页",
					},
				},
			})
		}
	}

	if canUpdates() {
		items.Append(&alfred.ResultElement{
			Valid:        true,
			Title:        "有新的更新可以更新",
			Subtitle:     "有道词典 for Alfred",
			Autocomplete: UPDATECMD,
		})
	}

	items.End()

	if agent.Dirty {
		agent.Cache.SaveFile(CACHE_FILE)
	}
}

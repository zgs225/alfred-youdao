package main

import (
	"encoding/gob"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zgs225/youdao"
)

const (
	CACHE_EXPIRES time.Duration = 30 * 24 * time.Hour
	CACHE_FILE    string        = "cache.dat"
)

type agentClient struct {
	Client *youdao.Client
	Cache  *cache.Cache
	Dirty  bool
}

func (a *agentClient) Query(q string) (*youdao.Result, error) {
	v, ok := a.Cache.Get(q)
	if ok {
		log.Println("Cache hit")
		return v.(*youdao.Result), nil
	}
	log.Println("Cache miss")
	r, err := a.Client.Query(q)
	if err != nil {
		return nil, err
	}
	a.Cache.Set(q, r, CACHE_EXPIRES)
	a.Dirty = true
	return r, nil
}

func newAgent(c *youdao.Client) *agentClient {
	gob.Register(&youdao.Result{})
	c2 := cache.New(CACHE_EXPIRES, CACHE_EXPIRES)
	err := c2.LoadFile(CACHE_FILE)
	if err != nil {
		log.Println(err)
	}
	log.Println("Cache count:", c2.ItemCount())
	return &agentClient{c, c2, false}
}

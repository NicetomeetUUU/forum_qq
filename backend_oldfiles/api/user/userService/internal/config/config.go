package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource string `json:"DataSource"`
	Cache      cache.CacheConf
}

type LogConf struct {
	ServiceName string `json:"ServiceName"`
	Mode        string `json:"Mode"`
	Path        string `json:"Path"`
	Level       string `json:"Level"`
	KeepDays    int    `json:"KeepDays"`
	Compress    bool   `json:"Compress"`
}

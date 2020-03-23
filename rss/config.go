package rss

import (
	"github.com/BurntSushi/toml"
)

type NextParser []string
type CName string

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func (n NextParser) Title() string {
	if len(n) != 0 {
		return n[0]
	}
	return ""
}

func (n NextParser) Body() string {
	if len(n) != 0 {
		return n[1]
	}
	return ""
}

type ConfigInterface interface {
	Load(string) bool
}

type Config struct {
	Owner   *Options
	Channel *COptions
}

type Options struct {
	Name    string
	Version string
	CName   string
}

type COptions struct {
	Title       string
	Link        string
	Description string
}

type CrawlerOptions struct {
	Node map[string]CrawlerNodeOptions
}

type CrawlerNodeOptions struct {
	Url        string
	ListParser string     `toml:"list_parser"`
	NextParser NextParser `toml:"next_parser"`
}

func (c *Config) Load(path string) bool {
	if _, err := toml.DecodeFile(path, c); err != nil {
		panic(err)
	}
	return true
}

func (c *CrawlerOptions) Load(path string) bool {
	if _, err := toml.DecodeFile(path, c); err != nil {
		panic(err)
	}
	return true
}

func NewRssOptions(path string) *Config {
	cfg := &Config{
		Owner:   &Options{},
		Channel: &COptions{},
	}
	if ok := cfg.Load(path); !ok {
		return nil
	}
	return cfg
}

func NewCrawlerOptions(path string) *CrawlerOptions {
	cfg := &CrawlerOptions{Node: make(map[string]CrawlerNodeOptions)}
	if ok := cfg.Load(path); ok {
		return cfg
	}
	return nil
}

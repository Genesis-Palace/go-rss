package test

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go-rss/crawler"
	"go-rss/rss"
	"io/ioutil"
	"testing"
)

func TestNewOptions(t *testing.T){
	opts := rss.NewRssOptions("../config/feeds.toml")
	cfg := rss.NewCrawlerOptions("../config/crawler.toml")
	for k, v := range cfg.Node{
		crawler.Crawler(k, v, opts)
	}
}

func TestNewFeeds(t *testing.T){
	content, err := ioutil.ReadFile("./feeds.xml");
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(content))
	app := iris.New()
	app.Get("/", func(ctx context.Context) {
		ctx.Text(string(content))
	})
	app.Run(iris.Addr(":8090"), iris.WithCharset("UTF-8"))
}
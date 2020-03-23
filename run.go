package main

import (
	"flag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go-rss/crawler"
	"go-rss/rss"
	"io/ioutil"
	"os"
)

var (
	server bool
)

func init(){
	flag.BoolVar(&server, "server", false, "")
	flag.Parse()
}

func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

func main(){
	if server{
		app := iris.New()
		app.Get("/feeds", func(context context.Context) {
			filename := context.URLParam("type")
			content, err := ioutil.ReadFile(filename+".feeds.xml");
			if err != nil{
				context.Text(err.Error())
			}
			context.Text(string(content))
		})
		app.Run(iris.Addr(":5530"))
	}else {
		opts := rss.NewRssOptions("config/feeds.toml")
		cfg := rss.NewCrawlerOptions("config/crawler.toml")
		for k, v := range cfg.Node{
			crawler.Crawler(k, v, opts)
		}
	}
}

package crawler

import (
	"encoding/xml"
	"fmt"
	"go-rss/rss"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)
import "github.com/PuerkitoBio/goquery"


func Crawler(typ string, conf rss.CrawlerNodeOptions, opts *rss.Config) {
	var feeds = rss.NewFeeds(opts)
	base, _:= url.Parse(conf.Url)
	doc, err := newDoc(conf.Url)
	if err != nil {
		fmt.Println(err, conf.Url)
		return
	}
	feeds.Channel.AddLastPubTime(time.Now().Format(rss.TimeFormat))
	doc.Find(conf.ListParser).Each(func(i int, selection *goquery.Selection) {
		if attr, ok := selection.Attr("href"); ok {
			u, err := url.Parse(attr)
			if err != nil {
				return
			}
			if u.Host == ""{
				attr = "http://" + base.Host + attr
			}
			fmt.Println(attr)
			item := nextCrawler(attr, conf)
			if item == nil{
				return
			}
			feeds.Channel.AddItem(item)
		}
	})
	html, err := xml.Marshal(feeds)
	if err != nil {
		return
	}
	if err := ioutil.WriteFile(typ + ".feeds.xml", html, 0666); err != nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		return
	}
}

func nextCrawler(url string, conf rss.CrawlerNodeOptions) *rss.Item {
	doc, err := newDoc(url)
	if err != nil {
		return nil
	}
	v := validator.New()
	d := &rss.Item{
		Title:       getElement(doc, conf.NextParser.Title()),
		Link:        url,
		PubData:     time.Now().Format(rss.TimeFormat),
		Description: getElement(doc, conf.NextParser.Body()),
	}
	err = v.Struct(d)
	if err != nil{
		return nil
	}
	return d
}

func newDoc(url string) (*goquery.Document, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func getElement(doc *goquery.Document, query string) string {
	doc.Find("span#imwp_tip").Each(func(i int, selection *goquery.Selection) {
		selection.Remove()
	})
	doc.Find(".big_ad").Each(func(i int, selection *goquery.Selection) {
		selection.Remove()
	})
	text, err := doc.Find(query).Html()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(text)
}

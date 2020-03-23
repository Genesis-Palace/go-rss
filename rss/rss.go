package rss

import (
	"encoding/xml"
	"io/ioutil"
)

type Feeds struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	LastBuildDate string `xml:"lastBuildDate"`
	Item          []*Item `xml:"item"`
}

func (c *Channel) AddItem(item *Item){
	c.Item = append(c.Item, item)
}

func (c *Channel) AddLastPubTime(pub string){
	c.LastBuildDate = pub
}

func (f *Feeds) Dumps(){
	feeds , err := xml.Marshal(f)
	if err != nil{
		panic(err)
	}
	if err := ioutil.WriteFile("feeds.xml", feeds, 0666); err != nil {
		panic(err)
	}

}

type Item struct {
	Title       string `xml:"title" validate:"required"`
	Link        string `xml:"link" validate:"required"`
	PubData     string `xml:"pubDate"`
	Description string `xml:"description" validate:"required"`
}

func NewFeeds(opts *Config) *Feeds {
	return &Feeds{
		XMLName: xml.Name{Local: "rss"},
		Version: opts.Owner.Version,
		Channel: &Channel{
			Title:         opts.Channel.Title,
			Link:          opts.Channel.Link,
			Description:   opts.Channel.Description,
			Item: []*Item{},
		},
	}
}

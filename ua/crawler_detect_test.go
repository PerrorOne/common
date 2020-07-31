package ua

import (
	"errors"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		ua string
	}
	tests := []struct {
		name string
		args args
		want func(obj CrawlerDetect) error
	}{
		{name: "baidu", args: args{ua: "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"}, want: func(obj CrawlerDetect) error {
			if !obj.iscrawler {
				return errors.New("spider not parse")
			}
			if obj.GetCrawler() != "Baiduspider" {
				return errors.New(fmt.Sprintf("spider name is not baidu, spider name: %s", obj.GetCrawler()))
			}
			return nil
		}},
		{name: "360", args: args{ua: "360spider (http://webscan.360.cn)"}, want: func(obj CrawlerDetect) error {
			if !obj.iscrawler {
				return errors.New("spider not parse")
			}
			if obj.GetCrawler() != "360spider" {
				return errors.New(fmt.Sprintf("spider name is not baidu, spider name: %s", obj.GetCrawler()))
			}
			return nil
		}},
		{name: "soso", args: args{ua: "Sosospider+(+http://help.soso.com/webspider.htm)"}, want: func(obj CrawlerDetect) error {
			if !obj.iscrawler {
				return errors.New("spider not parse")
			}
			if obj.GetCrawler() != "Sosospider" {
				return errors.New(fmt.Sprintf("spider name is not baidu, spider name: %s", obj.GetCrawler()))
			}
			return nil
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.args.ua)
			if err := tt.want(got); err != nil {
				t.Errorf("Parse() name = %s; error=%v", tt.name, err)
			}
		})
	}
}

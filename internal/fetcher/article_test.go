package fetcher

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/hi20160616/exhtml"
	"github.com/hi20160616/ms-voa/configs"
	"github.com/pkg/errors"
)

// pass test
func TestFetchArticle(t *testing.T) {
	tests := []struct {
		url string
		err error
	}{
		{
			"https://www.voachinese.com/a/us-house-hearing-blinken-china-alienating-world-20210607/5920478.html",
			ErrTimeOverDays,
		},
		{
			"https://www.voachinese.com/a/Trump-shuts-down-blog-nearly-erasing-online-presence/5914287.html",
			nil,
		},
	}
	for _, tc := range tests {
		a := NewArticle()
		a, err := a.fetchArticle(tc.url)
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			} else {
				fmt.Println("ignore old news pass test: ", tc.url)
			}
		} else {
			fmt.Println("pass test: ", a.Content)
		}
	}
}

func TestFetchTitle(t *testing.T) {
	tests := []struct {
		url   string
		title string
	}{
		{
			"https://www.voachinese.com/a/us-house-hearing-blinken-china-alienating-world-20210607/5920478.html",
			"塑造“可信、可爱、可敬”中国？ 布林肯：北京可能意识到战狼外交起反效果",
		},
		{
			"https://www.voachinese.com/a/Trump-shuts-down-blog-nearly-erasing-online-presence/5914287.html",
			"特朗普关闭博客 其网络平台几乎全部消失",
		},
	}
	for _, tc := range tests {
		a := NewArticle()
		u, err := url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		a.U = u
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		got, err := a.fetchTitle()
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			} else {
				fmt.Println("ignore pass test: ", tc.url)
			}
		} else {
			if tc.title != got {
				t.Errorf("\nwant: %s\n got: %s", tc.title, got)
			}
		}
	}

}

func TestFetchUpdateTime(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			"https://www.voachinese.com/a/us-house-hearing-blinken-china-alienating-world-20210607/5920478.html",
			"2021-06-08 10:55:33 +0800 UTC",
		},
		{
			"https://www.voachinese.com/a/Trump-shuts-down-blog-nearly-erasing-online-presence/5914287.html",
			"2021-06-03 08:50:11 +0800 UTC",
		},
	}
	var err error
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}

	for _, tc := range tests {
		a := NewArticle()
		a.U, err = url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		tt, err := a.fetchUpdateTime()
		if err != nil {
			t.Error(err)
		} else {
			ttt := tt.AsTime()
			got := shanghai(ttt)
			if got.String() != tc.want {
				t.Errorf("\nwant: %s\n got: %s", tc.want, got.String())
			}
		}
	}
}

func TestFetchContent(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			"https://www.voachinese.com/a/us-house-hearing-blinken-china-alienating-world-20210607/5920478.html",
			"2021-06-08 10:55:33 +0800 UTC",
		},
		{
			"https://www.voachinese.com/a/Trump-shuts-down-blog-nearly-erasing-online-presence/5914287.html",
			"2021-06-03 08:50:11 +0800 UTC",
		},
	}
	var err error
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}

	for _, tc := range tests {
		a := NewArticle()
		a.U, err = url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		c, err := a.fetchContent()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(c)
	}
}

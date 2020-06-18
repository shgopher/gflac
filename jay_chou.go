package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/golang/glog"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)
const MURL = "https://www.52flac.com/"
var (
	people       string
	getURLNumber int
	cookieValue  string
	finalURL     string
)

func init() {
	// if need
	flag.StringVar(&cookieValue, "coo", "", "cookie")

	flag.StringVar(&people, "peo", "周杰伦", "the artist of the music which you want to get")
	flag.IntVar(&getURLNumber, "num", 100, "the number of geting music")
	flag.Parse()
	finalURL = fmt.Sprintf("%ssearch.php?q=%s", MURL,people)
}

func main() {
	fmt.Println("开始任务🐬")
	putOut()
	fmt.Println("任务完成☕️🔥")
}

// get url address
func getDownloadAddress(url string, downLoadAddress chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	c := colly.NewCollector()
	// Find and visit all links
	var ma []string
	c.OnHTML(".bgb ul", func(e *colly.HTMLElement) {
		ma = e.ChildAttrs("li h2 a", "href")
	})
	if err := c.Visit(url); err != nil {
		glog.Error(err, url)
	}
	//

	for _, v := range ma {
		downLoadAddress <- v
	}
}

// get baidu pan address
func getBaiduPanAddress(downloadAddress string, outputurl chan string) {

	ce := strings.Split(downloadAddress, "/")
	a := ce[:len(ce)-2]
	ai := strings.Join(a, "/")
	b := ce[len(ce)-1]
	downloadAddress = fmt.Sprintf("%s/download/%s", ai, b)
	//
	c := colly.NewCollector()
	// Find and visit all links
	if err := c.SetCookies(downloadAddress, []*http.Cookie{
		{
			Name:  "Nobird_DownLoad",
			Value: cookieValue,
		},
	}); err != nil {
		glog.Error(err)
	}
	c.OnHTML(".con", func(e *colly.HTMLElement) {
		if e.Index == 1 {
			outputurl <- e.DOM.Text()
		}
	})
	if err := c.Visit(downloadAddress); err != nil {
		glog.Error(err)
	}
}

// put address
func putOut() {
	wg := new(sync.WaitGroup)
	downLoadAddress := make(chan string, 16)
	rootUrl := musicNumber(getURLNumber)
	for _, s := range rootUrl {
		wg.Add(1)
		s1 := s
		go getDownloadAddress(s1, downLoadAddress, wg)
	}
	go func() {
		wg.Wait()
		close(downLoadAddress)
	}()
	d := make(chan string)
	wg2 := new(sync.WaitGroup)
	wg2.Add(32)
	for i := 0; i < 32; i++ {
		go func() {
			defer wg2.Done()
			for address := range downLoadAddress {
				getBaiduPanAddress(address, d)
			}
		}()
	}
	go func() {
		wg2.Wait()
		close(d)
	}()
	wg3 := new(sync.WaitGroup)
	wg3.Add(32)
	fii,_ := os.Getwd()
	fii,_ = filepath.Split(fii)
	fmt.Println("输出路径：",fii)
	f, err := os.Create(fii+people+"_"+"music.csv")
	if err != nil {
		glog.Error(err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()
	w.Write([]string{"内容", "歌手"})
	for i := 0; i < 32; i++ {
		go func() {
			defer wg3.Done()
			for k := range d {
				if err := w.Write([]string{
					k,
					people,
				}); err != nil {
					glog.Error(err)
				}
			}
		}()
	}
	wg3.Wait()
}

// number

func musicNumber(number int) []string {
	number = int(math.Abs(float64(number)))
	//
	num := number >> 4
	result := make([]string, num+1)
	result[0] = finalURL
	//
	for i := 1; i <= num; i++ {
		ii := i
		result[ii] = fmt.Sprintf(finalURL+"&page=%d", ii+1)
	}
	return result
}

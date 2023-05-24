package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Telegrambot&cookie配置//--------------------------------------------------------|
const chatID = int64(114514) //chatid                                    |
const bottoken = "token"     //bottoken
const cookie = string("cookie")

// --------------------------------------------------------------------|
func option(host string, cookie string) {
	//参数定义
	url := "https://" + host + "/create-vps"
	var lastOptions []string
	noOptions := false
	client := &http.Client{}
	var wg sync.WaitGroup

	//检查配置正确性

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "PHPSESSID="+cookie+"; _ga=GA1.3.793762279.1679893292; __cf_bm=n1O7yJJxTW9hu.qhe6QRsCcE_kAZmAH1JHs6tEvmS.U-1680529193-0-Aemg00EIOnWChqBCRsyVnwuWHdxRP+evClmW8EKRQJQgiAb+BHppJXgtZc3FOuxR655TQ1GfpXRoYi62SQbA/5fTF3uyKoWeCOxr2tUojSIQVF8JcHQQAQcRo+HMqoOY8A; FCCDCF=%5Bnull%2Cnull%2Cnull%2C%5B%22CPpSJ0APpSJ0AEsABBENC9CoAP_AAG_AABAYINJB7D7FbSFCyP57aLsAMAhXRkCAQqQCAASBAmABQAKQIAQCkkAYFESgBAACAAAgIAJBIQIMCAgACUABQAAAAAEEAAAABAAIIAAAgAEAAAAIAAACAIAAEAAIAAAAEAAAmQhAAIIACAAAhAAAIAAAAAAAAAAAAgCAAAAAAAAAAAAAAAAAAQQaQD2F2K2kKEkfjWUWYAQBCujIEAhUAEAAECBIAAAAUgQAgFIIAwAIlACAAAAABAQAQCQgAQABAAAoACgAAAAAAAAAAAAAAQQAABAAIAAAAAAAAEAQAAIAAQAAAAAAABEhCAAQQAEAAAAAAAQAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAgAA%22%2C%221~2072.70.89.93.108.122.149.196.2253.2299.259.2357.311.317.323.2373.338.358.2415.415.2506.2526.482.486.494.495.2568.2571.2575.540.574.2624.624.2677.827.864.981.1048.1051.1095.1097.1171.1201.1205.1276.1301.1365.1415.1449.1570.1577.1651.1716.1753.1765.1870.1878.1889.1958.2012%22%2C%22ED667BE9-C3D9-464E-AE09-66952825F61F%22%5D%2Cnull%2Cnull%2C%5B%5D%5D; FCNEC=%5B%5B%22AKsRol8G5fkNHyoHbXzmXFTlKM7KH0iPAJbKDS4iyZ0XLicl7tvXtuH4sNjs4RPCJC30SL7NwBJC_E11HzjQmvk5cFN7VXN5qp94lNdxbHqZn86SmCcA1PlrgOgBm-1nKp4VBR9_8qwTZf8i3Q8BXnpdn-k3ZNSu0g%3D%3D%22%5D%2Cnull%2C%5B%5D%5D; PHPSESSID=c82418d070a8c8e767e0ce303a71159a")
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("clientrepuest error: %s", err))
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Sprintf("Body.Close error: %s", err))
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("NewDocumentFromReader error: %s", err))
	}

	if doc.Find("#datacenter option").Length() > 0 {
		fmt.Println("option find")
		sendmsg("检查配置成功", host)
	} else {
		fmt.Println("option not find")
		sendmsg("检查"+host+"配置不成功跳过", host)
		fmt.Println("option " + host + " not find")
		defer wg.Done()
	}

	///-----------------------------------------------------------------RUN-------------------------------
	for {
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Haxclient")
		req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("Host", host)
		req.Header.Set("Cookie", "PHPSESSID="+cookie+"; _ga=GA1.3.793762279.1679893292; __cf_bm=n1O7yJJxTW9hu.qhe6QRsCcE_kAZmAH1JHs6tEvmS.U-1680529193-0-Aemg00EIOnWChqBCRsyVnwuWHdxRP+evClmW8EKRQJQgiAb+BHppJXgtZc3FOuxR655TQ1GfpXRoYi62SQbA/5fTF3uyKoWeCOxr2tUojSIQVF8JcHQQAQcRo+HMqoOY8A; FCCDCF=%5Bnull%2Cnull%2Cnull%2C%5B%22CPpSJ0APpSJ0AEsABBENC9CoAP_AAG_AABAYINJB7D7FbSFCyP57aLsAMAhXRkCAQqQCAASBAmABQAKQIAQCkkAYFESgBAACAAAgIAJBIQIMCAgACUABQAAAAAEEAAAABAAIIAAAgAEAAAAIAAACAIAAEAAIAAAAEAAAmQhAAIIACAAAhAAAIAAAAAAAAAAAAgCAAAAAAAAAAAAAAAAAAQQaQD2F2K2kKEkfjWUWYAQBCujIEAhUAEAAECBIAAAAUgQAgFIIAwAIlACAAAAABAQAQCQgAQABAAAoACgAAAAAAAAAAAAAAQQAABAAIAAAAAAAAEAQAAIAAQAAAAAAABEhCAAQQAEAAAAAAAQAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAgAA%22%2C%221~2072.70.89.93.108.122.149.196.2253.2299.259.2357.311.317.323.2373.338.358.2415.415.2506.2526.482.486.494.495.2568.2571.2575.540.574.2624.624.2677.827.864.981.1048.1051.1095.1097.1171.1201.1205.1276.1301.1365.1415.1449.1570.1577.1651.1716.1753.1765.1870.1878.1889.1958.2012%22%2C%22ED667BE9-C3D9-464E-AE09-66952825F61F%22%5D%2Cnull%2Cnull%2C%5B%5D%5D; FCNEC=%5B%5B%22AKsRol8G5fkNHyoHbXzmXFTlKM7KH0iPAJbKDS4iyZ0XLicl7tvXtuH4sNjs4RPCJC30SL7NwBJC_E11HzjQmvk5cFN7VXN5qp94lNdxbHqZn86SmCcA1PlrgOgBm-1nKp4VBR9_8qwTZf8i3Q8BXnpdn-k3ZNSu0g%3D%3D%22%5D%2Cnull%2C%5B%5D%5D; PHPSESSID=c82418d070a8c8e767e0ce303a71159a")
		req.Header.Set("Connection", "keep-alive")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 1)
			continue
		}
		if err != nil {
			panic(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)

		//if resp.StatusCode != 200 {
		//	panic(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
		//}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}

		var options []string
		doc.Find("#datacenter option").Each(func(i int, s *goquery.Selection) {
			val, exists := s.Attr("value")
			if exists && val != "" && val != "-select-" {
				options = append(options, strings.TrimSpace(s.Text()))
			}
		})

		if len(options) > 0 {
			if len(lastOptions) == 0 {
				fmt.Println(strings.Join(options, "\n"))
				sendmsg(strings.Join(options, "\n"), host)
			} else {
				var newOptions []string
				for _, opt := range options {
					if !contains(lastOptions, opt) {
						newOptions = append(newOptions, opt)
					}
				}
				if len(newOptions) > 0 {
					fmt.Println(strings.Join(newOptions, "\n"))
					sendmsg(strings.Join(newOptions, "\n"), host)
				}
			}
			lastOptions = options
		} else {
			if !noOptions {
				fmt.Println("no options")
				sendmsg("no options", host)
				noOptions = true
			}
			lastOptions = nil
		}

		time.Sleep(1 * time.Microsecond)
	}
}
func contains(options []string, opt string) bool {
	for _, o := range options {
		if o == opt {
			return true
		}
	}
	fmt.Println("none")
	return false
}
func sendmsg(msg string, host string) { //chatid
	bot, err := telegram.NewBotAPI(bottoken) //bottoken
	if err != nil {
		panic(fmt.Sprintf("telegram bot error: %s", err))
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05.00000")
	msgText := host + " machine available:\n\n" + msg + "\n\n" + "time: " + currentTime //botsend
	msgconfig := telegram.NewMessage(chatID, msgText)
	_, err = bot.Send(msgconfig)
	if err != nil {
		return
	}
	fmt.Println(msgText)
}

func main() {
	// 网址列表
	urls := []string{
		"hax.co.id",
		"woiden.id",
		"free.vps.vc",
	}
	// 启动 goroutine
	for _, url := range urls {
		go option(url, cookie)
	}
	select {}
}

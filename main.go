package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"time"
)

func checkWebpage(bot *telegram.BotAPI, chatID int64, lastOptions map[string]bool) {
	url := "https://woiden.id/create-vps/" //网址

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	selectElement := doc.Find("select#datacenter")
	if selectElement.Length() == 0 {
		fmt.Println("No select element found")
		return
	}

	optionElements := selectElement.Find("option")
	if optionElements.Length() == 0 {
		fmt.Println("No option element found")
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05.00000")
	options := make(map[string]bool)
	optionElements.Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("value")
		if exists {
			options[value] = true
			if _, ok := lastOptions[value]; !ok {
				msgText := fmt.Sprintf("%s - %s", currentTime, s.Text()) //botsend
				msg := telegram.NewMessage(chatID, msgText)
				bot.Send(msg)
				fmt.Println(msgText)
			}
		}
	})

	for k := range lastOptions {
		if _, ok := options[k]; !ok {
			delete(lastOptions, k)
		}
	}

	for k := range options {
		if _, ok := lastOptions[k]; !ok {
			lastOptions[k] = true
		}
	}
}

func main() {
	chatID := int64(-1127432)                  //chatid
	bot, err := telegram.NewBotAPI("bottoken") //bottoken
	if err != nil {
		panic(err)
	}

	lastOptions := make(map[string]bool)

	for {
		res, err := http.Get("https://woiden.id/create-vps/") //网址
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		defer res.Body.Close()

		content, err := goquery.NewDocumentFromReader(res.Body) //获取网页内容并判断是否改变
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		selectElement := content.Find("select#datacenter")
		if selectElement.Length() == 0 {
			fmt.Println("No select element found")
			continue
		}

		optionElements := selectElement.Find("option")
		if optionElements.Length() == 0 {
			fmt.Println("No option element found")
			continue
		}

		checkWebpage(bot, chatID, lastOptions)

		time.Sleep(2 * time.Second)
	}
}

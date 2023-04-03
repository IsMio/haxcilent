package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"time"
)

func checkWebpage(bot *telegram.BotAPI, chatID int64) { //检查网页
	url := "https://hax.co.id/create-vps/		" //网址(此处为获取网页选择datacenter元素框中的服务器并推送到telegram)
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

	currentTime := time.Now().Format("2006-01-02 15:04:05:05:05") //TIMESTAMP与信息处理
	optionElements.Each(func(i int, s *goquery.Selection) {
		msgText := fmt.Sprintf("%s - %s", currentTime, s.Text())
		msg := telegram.NewMessage(chatID, msgText)
		bot.Send(msg)
		fmt.Println(msgText)
	})
}

func main() { //主函数
	lastContent := ""
	chatID := int64(-1145141919810)           //chat id
	bot, err := telegram.NewBotAPI("<token>") //bot token
	if err != nil {
		panic(err)
	}

	for {
		res, err := http.Get("https://hax.co.id/create-vps/") //网址
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

		if content.Text() != lastContent {
			fmt.Println("Webpage content has changed")
			checkWebpage(bot, chatID)
			lastContent = content.Text()
		} else {
			fmt.Println("Webpage content has not changed")
		}

		time.Sleep(1 * time.Second) //每秒检查一次
	}
}

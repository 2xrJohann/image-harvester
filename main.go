package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func imageDL(url string, name string) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	file, _ := os.Create("./images/" + name + ".png")
	defer file.Close()
	io.Copy(file, response.Body)
}

func getMsg(s *discordgo.Session, ch chan string, after string) (next string) {
	message, _ := s.ChannelMessages("", 1, after, "", "") //put ur channel id in the quotes (first param)
	if len(message) != 0 {
		if message[0].Content == "" {
			fmt.Printf("%+v\n", message[0].Attachments[0].URL)
			ch <- message[0].Attachments[0].URL
		}
		fmt.Println(message[0].ID)
		return message[0].ID
	} else {
		return "done"
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	ch := make(chan string)

	counter := 0

	go func(ch <-chan string) {
		for {
			url := <-ch
			fmt.Println(url)
			go imageDL(url, strconv.Itoa(counter))
			counter++
		}
	}(ch)

	next := getMsg(s, ch, "")

	go func(s *discordgo.Session, ch chan string, next string) {
		for ok := true; ok; ok = (next != "done") {
			next = getMsg(s, ch, next)
		}
		fmt.Println("done")
	}(s, ch, next)

	fmt.Printf("\n")

}

func main() {

	discord, err := discordgo.New("Bot " + "") //put ur bot token in the quotes

	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	discord.AddHandler(ready)

	err = discord.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	select {}
}

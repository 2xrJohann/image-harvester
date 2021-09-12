package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func readdir(s *discordgo.Session, ch chan string) {
	files, err := ioutil.ReadDir("../images/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		ch <- "../images/" + f.Name()
	}

}

func sned(s *discordgo.Session, ch chan string) {
	counter := 0

	for {
		select {
		case url := <-ch:
			var file, _ = os.OpenFile(url, os.O_RDWR, 0644)
			_, _ = s.ChannelFileSend("", strconv.Itoa(counter)+".png", file) //first param is channel id
			counter++
		}
	}

}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	ch := make(chan string)

	go readdir(s, ch)
	go sned(s, ch)
}

func main() {

	discord, err := discordgo.New("Bot " + "") //bot token here

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

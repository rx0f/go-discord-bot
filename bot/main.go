package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	CheckError(err)

	var TOKEN string = os.Getenv("TOKEN")

	fmt.Println("Trying to connect...")

	client, err := discordgo.New("Bot " + TOKEN)
	CheckError(err)

	client.AddHandler(MessageCreate)

	err = client.Open()
	CheckError(err)

	fmt.Println("Connected")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	client.Close()
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "pong!")
	}
}

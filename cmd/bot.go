package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RickChaves29/bot_pull_request/internal/data"
	"github.com/RickChaves29/bot_pull_request/utils"
	"github.com/bwmarrin/discordgo"
)

func main() {
	URL_REQUEST := fmt.Sprintf("https://api.github.com/repos/%v/%v/pulls?state=%v", os.Getenv("ORIGIN"), os.Getenv("REPOSITORY"), os.Getenv("STATE"))
	CHANNEL := os.Getenv("CHANNEL")

	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("LOG [error]: %v", err)
		return
	}
	pullrequests := data.ConnectionOnGithub(URL_REQUEST)
	err = s.Open()
	defer s.Close()
	if err != nil {
		log.Fatalf("LOG [error]: %v", err)
		return
	}
	if len(pullrequests) == 0 {
		embad := discordgo.MessageEmbed{
			Title: "Não a Pull Requests pendentes Hoje",
		}
		s.ChannelMessageSendEmbed(CHANNEL, &embad)
	} else {
		for _, pull := range pullrequests {
			date := utils.FormatDate(pull.DatePR)
			embad := discordgo.MessageEmbed{
				Title: pull.Title,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    pull.User.UserName,
					IconURL: pull.User.UserAvatar,
				},
				Description: pull.State,
				URL:         pull.Url,
				Footer: &discordgo.MessageEmbedFooter{
					Text: date,
				},
			}
			s.ChannelMessageSendEmbed(CHANNEL, &embad)
		}
	}
	log.Printf("LOG [bot]: is runnig !")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

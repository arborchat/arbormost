package main

import (
	"flag"
	"log"
	"net"

	arbor "github.com/arborchat/arbor-go"
	"github.com/mattermost/mattermost-server/model"
)

func postToMM(url, teamName, channelName, username, password string) (chan<- string, error) {

	client := model.NewAPIv4Client(url)
	_, response := client.Login(username, password)
	if response.Error != nil {
		log.Println(response)
		return nil, response.Error
	}
	log.Println("Login succeeded")
	team, response := client.GetTeamByName(teamName, "")
	if response.Error != nil {
		log.Println(response)
		return nil, response.Error
	}
	log.Println(team)
	channel, response := client.GetChannelByNameForTeamName(channelName, team.Name, "")
	if response.Error != nil {
		log.Println(response)
		return nil, response.Error
	}
	log.Println(channel)

	input := make(chan string)
	go func() {
		defer close(input)
		for value := range input {
			post, response := client.CreatePost(&model.Post{
				ChannelId: channel.Id,
				Message:   value})
			if response.Error != nil {
				log.Println(response)
				return
			}
			log.Println(post)
		}
	}()

	return input, nil
}

func main() {
	var username, password, url, team, channel, arborAddress string
	flag.StringVar(&username, "username", "", "mattermost server username")
	flag.StringVar(&password, "password", "", "mattermost server password")
	flag.StringVar(&team, "team", "", "mattermost server team")
	flag.StringVar(&channel, "channel", "", "mattermost server channel")
	flag.StringVar(&url, "url", "", "mattermost server url")
	flag.StringVar(&arborAddress, "arbor-address", "localhost:7777", "arbor server address")
	flag.Parse()
	sendChan, err := postToMM(url, team, channel, username, password)
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := net.Dial("tcp", arborAddress)
	if err != nil {
		log.Println(err)
		return
	}
	recvChan := arbor.MakeMessageReader(conn)
	for mesg := range recvChan {
		if mesg.Type == arbor.NewMessageType {
			sendChan <- arborAddress + " [id](" + mesg.UUID + ") [re](" + mesg.Parent + ") @" + mesg.Username + ": " + mesg.Content
		}
	}
}

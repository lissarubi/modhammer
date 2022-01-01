package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
	color "github.com/fatih/color"
	"github.com/gempir/go-twitch-irc"
	"github.com/joho/godotenv"
)

func main(){

	red := color.New(color.FgWhite).Add(color.BgRed).SprintFunc()

	if len(os.Args) < 2{
		fmt.Println(red("Modhammer needs a minium of one parameter"))	
		os.Exit(0)
	}

	if os.Args[1] == "--setup"{
		setup()
	}else{
		run()
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func setup() {

	// colors init
	green := color.New(color.FgBlack).Add(color.BgGreen).SprintFunc()

	// get data
	fmt.Println(green("Put your Twitch Username:"))
	username := prompt.Input("> ", completer)

	fmt.Println(green("Put your Twitch Token:"))
	token := prompt.Input("> ", completer)

	fmt.Println(green("Put the channels on which do you want to send the messages (split with comma)"))
	channels := prompt.Input("> ", completer)

	configFile := "USER=" + username + "\nTOKEN=" + token + "\nCHANNELS=" + channels

	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	WriteToFile( "/home/" + user.Username + "/.modhammer.env", configFile)
}

func run() {

	// colors init
	green := color.New(color.FgBlack).Add(color.BgGreen).SprintFunc()
	yellow := color.New(color.FgBlack).Add(color.BgYellow).SprintFunc()
	red := color.New(color.FgWhite).Add(color.BgRed).SprintFunc()
	white := color.New(color.FgBlack).Add(color.BgWhite).SprintFunc()

	// dotenv init
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	configPath := "/home/" + user.Username + "/.modhammer.env"

	err = godotenv.Load(configPath)
	if err != nil {
		fmt.Println(red("Config file doesn't exist. Create this file using modhammer --setup"))
		os.Exit(0)
	}

	message := ""
	for i := 1; i < len(os.Args); i++{
		if i == len(os.Args) - 1{
			message += os.Args[i]
		}else{
			message += os.Args[i] + " "
		}
	}

	modName := os.Getenv("USER")
	modToken := os.Getenv("TOKEN")
	modChannels := strings.Split(os.Getenv("CHANNELS"), ",")

	// irc connection
	client := twitch.NewClient(modName, modToken)

	fmt.Println(green("CONNECTING TO CHANNELS") + "\n")
	client.OnConnect(func() {

	fmt.Println(green("CONNECTED") + "\n")

		// do action at each channel
		for i := 0; i < len(modChannels); i++{
			fmt.Printf("%s %s %s %s\n", green("SENDING"), yellow(message), white("IN"), green(modChannels[i]))
			client.Say(modChannels[i], message)
		}
		time.Sleep(time.Second * 3)
		fmt.Println("\n" + green("DONE"))
		os.Exit(0)
	})

	// join channels
	client.Join(modChannels)

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

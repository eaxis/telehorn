package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"log"
	"os"
)

type HttpResponse struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
}

type SendingParams struct {
	Token   string `json:"token" binding:"required"`
	Chats   []int  `json:"chats" binding:"required,gte=1,dive,gte=1"`
	Message string `json:"message" binding:"required"`
}

func (s *SendingParams) validate() error {
	if s.Token == "" {
		return errors.New("Token can not be blank.")
	}

	if s.Message == "" {
		return errors.New("Message can not be blank.")
	}

	var validChats []int

	for _, v := range s.Chats {
		if v > 0 {
			validChats = append(validChats, v)
		}
	}

	if len(validChats) == 0 {
		return errors.New("Chats are not valid.")
	}

	return nil
}

func main() {
	app := cli.NewApp()

	app.Name = "TeleHorn"
	app.Usage = "simple & flexible tool to make newsletters in Telegram"

	app.UsageText = "To see help for GUI/API: 'telehorn web -h'\r\n   " +
		"To see help for CLI: 'telehorn cli -h'"

	app.Description += "You are free to use API, GUI, CLI or Go-bindings to send your messages via TeleHorn.\r\n   " +
		"TeleHorn sends about 30 messages per second and then sleep for 1 second to avoid limits of Telegram.\r\n   " +
		"While you using CLI version, you will see amount of sent messages (with failures).\r\n   " +
		"Read detailed docs at github.com/narrator69/telehorn."

	app.Authors = []cli.Author{{
		Name: "Narrator69",
		Email: "github.com/narrator69",
	}}

	app.Commands = []cli.Command{
		getWebCommand(),
		getCliCommand(),
	}

	app.HideVersion = true

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

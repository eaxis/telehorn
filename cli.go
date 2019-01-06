package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func getCliCommand() cli.Command {
	return cli.Command{
		Name:        "cli",
		Usage:       "Execute one time (CLI)",
		UsageText: 	 "cli [--token=abc --chat=1 --chat=2 --message=abc] | [--json=file.json]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "token",
				Usage: "Telegram Bot-API token",
			},
			cli.IntSliceFlag{
				Name:  "chat",
				Usage: "An ID of a chat",
			},
			cli.StringFlag{
				Name:  "chats",
				Usage: "A list of ID's of chats (comma-separated)",
			},
			cli.StringFlag{
				Name:  "message",
				Usage: "A message to send",
			},
			cli.StringFlag{
				Name:  "file",
				Usage: "A JSON-file to use (instead of flags)",
			},
		},
		Action: func(c *cli.Context) error {
			var chats []int

			if c.String("chats") != "" {
				split := strings.Split(c.String("chats"), ",")

				for _, v := range split {
					integer, err := strconv.Atoi(v)

					if integer > 0 && err == nil {
						chats = append(chats, integer)
					}
				}
			}

			chats = append(chats, c.IntSlice("chat")...)

			err := startCliHandler(
				c.String("token"),
				c.String("message"),
				c.String("file"),
				chats,
			)

			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}

			return nil
		},
	}
}

func startCliHandler(token string, message string, jsonFile string, chats []int) error {
	var params SendingParams

	if jsonFile != "" {
		path, err := filepath.Abs(jsonFile)

		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		err = json.Unmarshal(bytes, &params)

		if err != nil {
			return err
		}
	} else {
		params.Token = token
		params.Chats = chats
		params.Message = message
	}

	err := params.validate()

	if err != nil {
		return err
	}

	service, err := NewTeleHorn(params.Token)

	if err != nil {
		return errors.New("Incorrect token / Timeout.")
	}

	results := service.Send(params.Chats, params.Message)

	fmt.Println(
		fmt.Sprintf(
			"Done! Messages sent: %d, Failed: %d, Successful: %d",
			len(results.Failed)+len(results.Successful),
			results.Failed,
			results.Successful,
		),
	)

	return nil
}

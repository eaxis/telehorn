package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"net/http"
)

func getWebCommand() cli.Command {
	return cli.Command{
		Name:        "web",
		Usage:       "Start web server (GUI/API)",
		UsageText: 	 "web [--port=1337] [--user=admin] [--pass=supersonic]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "port",
				Usage: "Web server will run on a specific port",
				Value: "2019",
			},
			cli.StringFlag{
				Name:  "user",
				Usage: "Should be specified to protect web server with basic auth",
			},
			cli.StringFlag{
				Name:  "pass",
				Usage: "Should be specified to protect web server with basic auth",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(
				fmt.Sprintf(
					"Starting up HTTP-Server on localhost:%s...",
					c.String("port"),
				),
			)

			err := startWebHandler(
				c.String("port"),
				c.String("user"),
				c.String("pass"),
			)

			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}

			return nil
		},
	}
}

func startWebHandler(port string, user string, pass string) error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	if user != "" && pass != "" {
		r.Use(gin.BasicAuth(
			gin.Accounts{
				user: pass,
			},
		))
	}

	//r.Delims("{[{", "}]}")

	r.Static("static", "view/static/")
	r.LoadHTMLGlob("view/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", "")
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.tpl", "")
	})

	r.POST("/submit", func(c *gin.Context) {
		var params SendingParams

		c.ShouldBindJSON(&params)

		err := params.validate()

		if err != nil {
			c.JSON(500, HttpResponse{
				Success:     false,
				Description: err.Error(),
			})

			return
		}

		service, err := NewTeleHorn(params.Token)

		if err != nil {
			c.JSON(500, HttpResponse{
				Success:     false,
				Description: "Incorrect token / Timeout.",
			})

			return
		}

		c.JSON(200, HttpResponse{
			Success:     true,
			Description: "We are sending your messages.",
		})

		go service.Send(params.Chats, params.Message)
	})

	err := r.Run(":" + port)

	return err
}

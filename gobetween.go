package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/libgit2/git2go"
)

func main() {
	commands := []cli.Command{
		{
			Name:        "status",
			ShortName:   "s",
			Usage:       "./gobetween s",
			Description: "Get Status of current directories dependencies, based on glide.lock",
			Action: func(c *cli.Context) {
				workingDir, err := os.Getwd()
				if err != nil {
					log.Printf("Working Directory failed:%v\n", err)
					os.Exit(100)
				}
				repo, err := git.OpenRepository(workingDir)
				if err != nil {
					log.Printf("git.OpenRepository failed:%v\n", err)
					os.Exit(100)
				}
				state := repo.State()
				log.Printf("State:%v\n", state)
			},
		},
		{
			Name:        "Prepare",
			ShortName:   "p",
			Usage:       "./gobetween p",
			Description: "Get Status of current directories dependencies, based on glide.lock",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "orgId", Value: "", Usage: "A valid UUID"},
			},
			Action: func(c *cli.Context) {
				orgId := strings.TrimSpace(c.String("orgId"))

				if len(orgId) == 0 {
					log.Printf("Cannot pass a blank orgId")
					os.Exit(1)
				}
			},
		},
	}
	app := cli.NewApp()
	app.Commands = commands
	app.Name = "gobetween"
	app.Usage = "Glide Helper."
	app.Version = "0.0.1"

	app.Action = func(ctx *cli.Context) {
		if len(ctx.Args()) == 0 {
			cli.ShowAppHelp(ctx)
			os.Exit(1)
		}
		console := cli.NewApp()
		console.Commands = commands
		console.Action = func(c *cli.Context) {
			fmt.Println("Command not found. Type 'help' for a list of commands.")
		}
	}
	app.Run(os.Args)
	os.Exit(0)
}

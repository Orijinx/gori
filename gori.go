package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	// we create our commands
	app.Commands = []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a project",
			// the action, or code that will be executed when
			// we execute our `ns` command
			Action: func(c *cli.Context) error {

				fmt.Println(c.Args().Get(0))



				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	git "github.com/go-git/go-git" // with go modules disabled
	cli "github.com/urfave/cli"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()

	//Descriptions
	app.Authors = []*cli.Author{
		&cli.Author{Name: "Highjin", Email: "vladonnx@mail.ru"},
	}
	app.Copyright = "(C) Orijinx 2021"
	app.Name = "Gori"
	app.Version = "0.0.1"
	app.Description = "A CLI application for Gorijinx-framework"
	///////////////////////////////////////////////////////////////////////////

	// we create our commands
	app.Commands = []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a default project",
			// the action, or code that will be executed when
			// we execute our `ns` command
			Action: func(c *cli.Context) error {

				fmt.Println(c.Args().Get(0))

				Path := "./" + c.Args().Get(0)


				_, err := git.PlainClone(Path, false, &git.CloneOptions{
					SingleBranch: true,
					RemoteName: "framework",
					URL:      "https://github.com/Orijinx/Gorijinx",
					Progress: os.Stdout,
				})
				if err != nil {
					fmt.Println(err)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

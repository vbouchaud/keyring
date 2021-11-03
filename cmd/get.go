package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/zalando/go-keyring"
)

func Get() *cli.Command {
	return &cli.Command{
		Name:     "get",
		Aliases:  []string{"g"},
		Usage:    "get an entry from the keyring.",
		HideHelp: false,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "service",
				Required: true,
				Usage:    "The targeted service.",
			},
			&cli.StringFlag{
				Name:     "user",
				Required: true,
				Usage:    "The user to get for the targeted service.",
			},
		},
		Action: func(c *cli.Context) error {
			var (
				service = c.String("service")
				user    = c.String("user")
			)

			secret, err := keyring.Get(service, user)
			if err != nil {
				return err
			}

			fmt.Println(secret)

			return nil
		},
	}
}

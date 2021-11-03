package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/zalando/go-keyring"
)

func Del() *cli.Command {
	return &cli.Command{
		Name:     "delete",
		Aliases:  []string{"d", "del"},
		Usage:    "delete an entry from the keyring.",
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
				Usage:    "The user to remove for the targeted service.",
			},
		},
		Action: func(c *cli.Context) error {
			var (
				service = c.String("service")
				user    = c.String("user")
			)

			return keyring.Delete(service, user)
		},
	}
}

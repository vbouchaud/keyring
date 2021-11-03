package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"
	"github.com/zalando/go-keyring"
	"golang.org/x/term"
)

func readData(readLine func(screen io.ReadWriter) (string, error)) (string, error) {
	if !isatty.IsTerminal(os.Stdin.Fd()) && !isatty.IsCygwinTerminal(os.Stdin.Fd()) {
		return "", fmt.Errorf("stdin should be terminal")
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			// what should we do if we cannot restore terminal?
		}
	}(int(os.Stdin.Fd()), oldState)

	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}

	line, err := readLine(screen)
	if err != nil {
		return "", err
	}

	return line, nil

}

func password(screen io.ReadWriter) (string, error) {
	terminal := term.NewTerminal(screen, "")

	print("secret: ")

	line, err := terminal.ReadPassword("")

	if err == io.EOF || line == "" {
		return "", fmt.Errorf("secret cannot be empty")
	}

	return line, err
}

func Add() *cli.Command {
	return &cli.Command{
		Name:     "add",
		Aliases:  []string{"a"},
		Usage:    "add a secret to the keyring.",
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
				Usage:    "The user to use for the targeted service. (optional)",
			},
			&cli.StringFlag{
				Name:     "secret",
				Required: false,
				Usage:    "The secret to set.",
			}},
		Action: func(c *cli.Context) error {
			var (
				service = c.String("service")
				user    = c.String("user")
				secret  = c.String("secret")
			)

			if secret == "" {
				interactiveSecret, err := readData(password)
				print("\n")

				if err != nil {
					return err
				}

				secret = interactiveSecret
			}

			err := keyring.Set(service, user, secret)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

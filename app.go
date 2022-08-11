package cli

import (
	"errors"
)

type App struct {
	Name        string
	Usage       string
	Description string
	Commands    []*Command
}

func (app *App) Run(args []string) error {
	cmdCTX := ctx()

	cmd, err := app.checkCommand(cmdCTX)
	if err != nil {
		return err
	}

	cmd.Run(cmdCTX)

	return nil
}

func (app *App) checkCommand(ctx *Context) (*Command, error) {
	if ctx.args == nil || len(ctx.args) == 0 {
		return nil, errors.New("the app needs args")
	}

	cmdName := ctx.args[0]

	var cmd *Command
	for _, c := range app.Commands {
		if c.Name == cmdName {
			cmd = c
			break
		}
	}

	if cmd == nil {
		return nil, errors.New("not support command: " + cmdName)
	}

	return cmd, nil
}

// func (app *App)

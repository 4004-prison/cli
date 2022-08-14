package cli

type App struct {
	Name        string
	Usage       string
	Description string
	Commands    []*Command
}

func (app *App) Run(args []string) error {
	cmdCTX := initCTX(args)

	cmdName := cmdCTX.getCMD()

	var cmd *Command
	for _, c := range app.Commands {
		if c.Name == cmdName {
			cmd = c
		}
	}

	if err := cmdCTX.parse(cmd.Flags); err != nil {
		return err
	}

	cmd.Run(cmdCTX)

	return nil
}

// func (app *App)

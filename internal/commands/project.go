package commands

func GetProjectCommands() map[string]*Command {
	return map[string]*Command{
		"create": {
			Help:  "Create a new project",
			Usage: "mytodo project create [name] [--interactive | -i] --directory=[dir] --time=[hours]",
			Options: CommandOptions{
				"interactive": {
					DefaultValue: "false",
					Required: func(ctx *Context) bool {
						return false
					},
				},
				"name": {
					Required: func(ctx *Context) bool {
						return !isInteractive(ctx)
					},
				},
				"targetDuration": {
					Required: func(ctx *Context) bool {
						return !isInteractive(ctx)
					},
				},
			},
			Execute: func(ctx *Context) error {
				if err != nil {
					return err
				}
				return ctx.Service.CreateProject()
			},
		},
		"delete": {
			Help:  "delete a project",
			Usage: "mytodo project delete [name|id]",
			Options: CommandOptions{
				"directory": {
					Required: func(ctx *Context) bool {
						return true
					},
				},
			},
		},
	}
}

func isInteractive(ctx *Context) bool {
	return ctx.Options["interactive"] == "true" || ctx.Options["i"] == "true"
}

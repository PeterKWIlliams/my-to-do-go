package commands

import (
	"fmt"

	"github.com/PeterKWIlliams/my-to-do-go/internal/service"
)

type CommandOptions map[string]CommandOption

type CommandOption struct {
	DefaultValue string
	Required     func(*Context) bool
	Validation   func(string) error
}

type Command struct {
	Help        string
	Usage       string
	Options     CommandOptions
	SubCommands map[string]*Command
	Execute     func(ctx *Context) error
}

type Context struct {
	Service *service.Service
	Args    []string
	Options map[string]string
	Parent  *Context
}

func (c *Command) ValidateAndExectue(ctx *Context) error {
	for key, opt := range c.Options {

		value := ctx.Options[key]

		if opt.Required(ctx) && value == "" {
			return fmt.Errorf("required option %s not provided", key)
		}
		if opt.Validation != nil {
			if err := opt.Validation(value); err != nil {
				return fmt.Errorf("invalid value for option %s: %w", key, err)
			}
		}

	}
	return c.Execute(ctx)
}

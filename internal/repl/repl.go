package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PeterKWIlliams/my-to-do-go/internal/commands"
	"github.com/PeterKWIlliams/my-to-do-go/internal/todo"
)

func Start(cfg *todo.Config) {
	for {
		fmt.Print("This is a todo app")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		commandArgs := input[1:]
		command, ok := commands.GetCommands()[commandName]

		if ok {

			err := command.Callback(cfg, commandArgs...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("command not found type help for a list of commands")
		}

		continue

	}
}

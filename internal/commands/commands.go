package commands


type db struct{
  connecString string
}

type config struct {
  db *db
}

type command struct {
	help     string
	callback func(*config, ...string)
}



func getCommands() map[string]command {
	return map[string]command{
		"task": {
			help:     "Use this command to interact with tasks\nUsage: todo task [OPTIONS]\nOptions\n",
			callback: add,
		},
		"todos": {
			help:     "Show a list of all todos",
			callback: add,
		},
}

package main

import (
	"context"
	"os"
	"pip_toolkit_url_shortener/containers"
)

func main() {
	proc := containers.NewAliasProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(context.Background(), os.Args)
}

package main

import (
	"fmt"
	"os"

	"ccusage-rainbow/internal/frameworks/di"
)

func main() {
	container := di.NewContainer()
	rootCmd := container.GetCLIController().CreateRootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

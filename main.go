package main

import (
	"fmt"

	"github.com/ente-io/entectl/cmd"
	"github.com/ente-io/entectl/internal/config"
)

func main() {
	_, err := config.GetConfigDirPath()
	if err != nil {
		fmt.Println("Error getting config directory")
		return
	}

	cmd.Execute()
}

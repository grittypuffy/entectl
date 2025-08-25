package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	EnteSupportEndpoint = "https://support.ente.workers.dev/chat"
	TokenHeader         = "X-Auth-Token"
	TokenQuery          = "token"
	ClientPkgHeader     = "X-Client-Package"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command("docker", "compose", "stop").Run()
	},
}

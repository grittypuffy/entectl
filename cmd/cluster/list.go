package cluster

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all clusters",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command("docker", "compose", "ps").Run()
	},
}

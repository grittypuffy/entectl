package cluster

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command("docker", "compose", "down").Run()
	},
}

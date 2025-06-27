package cluster

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command("docker", "compose", "stop").Run()
	},
}

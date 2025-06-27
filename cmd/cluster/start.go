package cluster

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command("docker", "compose", "up", "-d").Run()
	},
}

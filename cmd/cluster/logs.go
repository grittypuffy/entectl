package cluster

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show cluster logs",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("docker", "compose", "logs", "-f")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

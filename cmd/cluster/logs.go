package cluster

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ente-io/entectl/internal/config"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show cluster logs",
	Run: func(cmd *cobra.Command, args []string) {
		clusterName, _ := cmd.Flags().GetString("name")

		if clusterName == "" {
			fmt.Println("Cluster name must be provided")
			return
		}

		clusterDir, err := config.GetClusterDir(clusterName)

		if err != nil {
			fmt.Println(fmt.Errorf("Error getting configuration directory: %w", err))
			return
		}

		err = os.Chdir(clusterDir)

		if err != nil {
			fmt.Println(fmt.Errorf("Error changing to cluster directory: %w", err))
			return
		}

		c := exec.Command("docker", "compose", "logs")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	logsCmd.Flags().String("name", "my-ente", "Name of the cluster")
}

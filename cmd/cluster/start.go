package cluster

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ente-io/entectl/internal/config"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a cluster",
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

		command := exec.Command("docker", "compose", "up", "-d")

		err = command.Run()

		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Printf("Docker Compose failed with exit code %d\n", exitErr.ExitCode())
			} else {
				fmt.Println("Error running Docker Compose:", err)
			}
		} else {
			fmt.Println("Started cluster", clusterName)
		}
	},
}

func init() {
	startCmd.Flags().String("name", "my-ente", "Name of the cluster")
}

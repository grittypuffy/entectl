package cluster

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/ente-io/entectl/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all services of a cluster",
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

		var stdout bytes.Buffer

		command := exec.Command("docker", "compose", "ps")

		command.Stdout = &stdout

		err = command.Run()
		if err != nil {
			fmt.Println("Error running command:", err)
			return
		}

		fmt.Println(stdout.String())
	},
}

func init() {
	listCmd.Flags().String("name", "my-ente", "Name of the cluster")
}

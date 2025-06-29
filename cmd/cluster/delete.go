package cluster

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ente-io/entectl/internal/config"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete cluster configuration",
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
		exec.Command("docker", "compose", "remove", "--volumes").Run()

		err = os.Chdir("..")
		if err != nil {
			fmt.Println("Error changing directory:", err)
			return
		}

		parentDir, _ := os.Getwd()
		err = os.RemoveAll(parentDir)
		if err != nil {
			fmt.Println("Error removing directory:", err)
		} else {
			fmt.Println("Directory removed successfully.")
		}
	},
}

func init() {
	deleteCmd.Flags().String("name", "my-ente", "Name of the cluster")
}

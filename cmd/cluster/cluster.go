package cluster

import (
	"github.com/spf13/cobra"
)

var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Run cluster commands",
}

func init() {
	ClusterCmd.AddCommand(initCmd)
	ClusterCmd.AddCommand(startCmd)
	ClusterCmd.AddCommand(listCmd)
	ClusterCmd.AddCommand(stopCmd)
	ClusterCmd.AddCommand(removeCmd)
	ClusterCmd.AddCommand(logsCmd)
	ClusterCmd.AddCommand(pullCmd)
}

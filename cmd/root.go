package cmd

import (
	"github.com/ente-io/entectl/cmd/cluster"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "entectl",
	Short: "CLI to manage self-hosted Ente clusters",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(cluster.ClusterCmd)
	rootCmd.AddCommand(chatCmd)
}

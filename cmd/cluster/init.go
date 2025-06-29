package cluster

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ente-io/entectl/internal/config"
	"github.com/spf13/cobra"
)

var configPath string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Docker Compose cluster",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		clusterName, _ := cmd.Flags().GetString("name")

		if configPath == "" {
			fmt.Println("Configuration file must be provided")
			return
		}

		if clusterName == "" {
			fmt.Println("Cluster name must be provided")
			return
		}

		configDir, err := config.GetConfigDir()
		if err != nil {
			fmt.Println(fmt.Errorf("Error getting configuration directory: %w", err))
			return
		}

		err = config.CreateClusterDir(clusterName)
		if err != nil {
			fmt.Println(fmt.Errorf("Error creating cluster: %w", err))
			return
		}

		clusterPath := filepath.Join(configDir, clusterName)

		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		err = os.Chdir(clusterPath)

		if err != nil {
			fmt.Println(fmt.Errorf("Error changing to cluster directory: %w", err))
			return
		}

		exePath, err := os.Executable()
		if err != nil {
			return
		}

		exeDir := filepath.Dir(exePath)

		composeTemplatePath := filepath.Join(exeDir, "internal/templates/compose.yaml.tmpl")

		err = config.RenderConfig(cfg, composeTemplatePath, "compose.yaml")
		if err != nil {
			fmt.Println("Error creating compose.yaml", err)
		} else {
			fmt.Println("compose.yaml generated.")
		}

		museumTemplatePath := filepath.Join(exeDir, "internal/templates/museum.yaml.tmpl")

		err = config.RenderConfig(cfg, museumTemplatePath, "museum.yaml")
		if err != nil {
			fmt.Println("Error creating museum.yaml:", err)
		} else {
			fmt.Println("museum.yaml generated.")
		}

		caddyTemplatePath := filepath.Join(exeDir, "internal/templates/reverse_proxy/Caddyfile.tmpl")

		err = config.RenderConfig(cfg, caddyTemplatePath, "Caddyfile")
		if err != nil {
			fmt.Println("Error creating Caddyfile:", err)
		} else {
			fmt.Println("Caddyfile generated.")
		}

	},
}

func init() {
	initCmd.Flags().String("config", "", "Path to config file")
	initCmd.Flags().String("name", "my-ente", "Name of the cluster")
}

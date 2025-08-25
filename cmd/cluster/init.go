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

		// Get template paths before changing directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting working directory:", err)
			return
		}

		composeTemplatePath := filepath.Join(wd, "internal/templates/compose.yaml.tmpl")
		museumTemplatePath := filepath.Join(wd, "internal/templates/museum.yaml.tmpl")
		caddyTemplatePath := filepath.Join(wd, "internal/templates/reverse_proxy/Caddyfile.tmpl")

		err = os.Chdir(clusterPath)

		if err != nil {
			fmt.Println(fmt.Errorf("Error changing to cluster directory: %w", err))
			return
		}

		err = config.RenderConfig(cfg, composeTemplatePath, "compose.yaml")
		if err != nil {
			fmt.Println("Error creating compose.yaml", err)
		} else {
			fmt.Println("compose.yaml generated.")
		}

		err = config.RenderConfig(cfg, museumTemplatePath, "museum.yaml")
		if err != nil {
			fmt.Println("Error creating museum.yaml:", err)
		} else {
			fmt.Println("museum.yaml generated.")
		}

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

package cluster

import (
	"fmt"

	"github.com/ente-io/entectl/internal/config"
	"github.com/ente-io/entectl/internal/template"
	"github.com/spf13/cobra"
)

var configPath string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Docker Compose cluster",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		if configPath == "" {
			fmt.Println("Configuration file must be provided")
			return
		}

		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		err = template.RenderDockerCompose(cfg, "internal/templates/compose.yaml.tmpl", "compose.yaml")
		if err != nil {
			fmt.Println("Error rendering compose file:", err)
		} else {
			fmt.Println("compose.yaml generated.")
		}
	},
}

func init() {
	initCmd.Flags().String("config", "", "Path to config file")
}

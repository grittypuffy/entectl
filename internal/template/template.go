package template

import (
	"os"
	"text/template"

	"github.com/ente-io/entectl/internal/config"
)

func RenderDockerCompose(cfg *config.Config, templatePath string, outputPath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return tmpl.Execute(outFile, cfg)
}

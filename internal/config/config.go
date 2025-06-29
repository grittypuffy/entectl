package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

type WebApp struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type Config struct {
	Domain     string         `yaml:"domain"`
	MuseumPort int            `yaml:"museum_port"`
	WebPorts   map[string]int `yaml:"web_ports"`

	DB struct {
		Password string `yaml:"password"`
	} `yaml:"db"`

	JWTSecret string `yaml:"jwt_secret"`
	EncKey    string `yaml:"enc_key"`
	HashKey   string `yaml:"hash_key"`

	S3 struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"s3"`
}

func LoadConfig(path string) (*Config, error) {
	absPath, err := filepath.Abs(path)

	if err != nil {
		return nil, fmt.Errorf("Error getting absolute path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	cfg.ApplyDefaults()

	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}

func CreateClusterDir(name string) error {
	configDir, err := GetConfigDir()
	if err != nil {
		return fmt.Errorf("Failed to get config directory: %w", err)
	}

	clusterPath := filepath.Join(configDir, name)

	info, err := os.Stat(clusterPath)

	if os.IsNotExist(err) {
		err := os.MkdirAll(clusterPath, 0755)
		if err != nil {
			return fmt.Errorf("Failed to create config directory: %w", err)
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("Error checking config directory: %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("Path exists but is not a directory: %s", clusterPath)
	}
	return nil
}

func GetClusterDir(name string) (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	clusterPath := filepath.Join(configDir, name)

	info, err := os.Stat(clusterPath)

	if os.IsNotExist(err) {
		return "", err
	}

	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return "", err
	}

	return clusterPath, nil
}

func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Failed to get home directory: %w", err)
	}

	entectlPath := filepath.Join(home, ".config", "entectl")

	info, err := os.Stat(entectlPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(entectlPath, 0755)
		if err != nil {
			return "", fmt.Errorf("Failed to create config directory: %w", err)
		}
		return entectlPath, nil
	}
	if err != nil {
		return "", fmt.Errorf("Error checking config directory: %w", err)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("Path exists but is not a directory: %s", entectlPath)
	}

	return entectlPath, nil
}

func (cfg *Config) ApplyDefaults() {
	if cfg.Domain == "" {
		cfg.Domain = "localhost"
	}
	if cfg.MuseumPort == 0 {
		cfg.MuseumPort = 8080
	}
	if len(cfg.WebPorts) == 0 {
		cfg.WebPorts = map[string]int{
			"photos":   3000,
			"accounts": 3001,
			"albums":   3002,
			"auth":     3003,
			"cast":     3004,
		}
	}
}

func GenerateFromTemplate(tmplFile string, cfg *Config, outputFile string) error {
	tmplData, err := os.ReadFile(tmplFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New(filepath.Base(tmplFile)).Parse(string(tmplData))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return err
	}

	return os.WriteFile(outputFile, buf.Bytes(), 0644)
}

func RenderConfig(cfg *Config, templatePath string, outputPath string) error {
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

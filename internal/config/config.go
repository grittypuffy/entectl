package config

import (
	"bytes"
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
	Email      string         `yaml:"email"`
	MuseumPort int            `yaml:"museum_port"`
	WebPorts   map[string]int `yaml:"web_ports"`

	DB struct {
		Host string `yaml:"host"`
	} `yaml:"db"`

	JWTSecret string `yaml:"jwt_secret"`
	EncKey    string `yaml:"enc_key"`
	HashKey   string `yaml:"hash_key"`

	Minio struct {
		Key      string `yaml:"key"`
		Secret   string `yaml:"secret"`
		Endpoint string `yaml:"endpoint"`
		Bucket   string `yaml:"bucket"`
		Region   string `yaml:"region"`
	} `yaml:"minio"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	cfg.ApplyDefaults()
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
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
			"photos": 3000,
			"public": 3002,
		}
	}

	if cfg.DB.Host == "" {
		cfg.DB.Host = "postgres"
	}

	if cfg.Minio.Endpoint == "" {
		cfg.Minio.Endpoint = "https://storage.ente.localhost"
	}
	if cfg.Minio.Region == "" {
		cfg.Minio.Region = "eu-central-2"
	}
	if cfg.Minio.Bucket == "" {
		cfg.Minio.Bucket = "b2-eu-cen"
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

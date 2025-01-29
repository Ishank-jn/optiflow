package config

import (
    "gopkg.in/yaml.v2"
    "os"
)

type Config struct {
    Server struct {
        Address string `yaml:"address"`
    } `yaml:"server"`
    Database struct {
        Host     string `yaml:"host"`
        Port     int    `yaml:"port"`
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        DBName   string `yaml:"dbname"`
    } `yaml:"database"`
    OAuth struct {
        ClientID     string `yaml:"client_id"`
        ClientSecret string `yaml:"client_secret"`
        RedirectURL  string `yaml:"redirect_url"`
        AuthURL      string `yaml:"auth_url"`
        TokenURL     string `yaml:"token_url"`
    } `yaml:"oauth"`
    Redis struct {
        Address  string `yaml:"address"`
        Password string `yaml:"password"`
        DB       int    `yaml:"db"`
    } `yaml:"redis"`
    JWT struct {
        SecretKey  string `yaml:"secret_key"`
        Expiration int    `yaml:"expiration"`
    } `yaml:"jwt"`
    Logging struct {
        Level  string `yaml:"level"`
        Format string `yaml:"format"`
    } `yaml:"logging"`
    Prometheus struct {
        Enabled bool   `yaml:"enabled"`
        Path    string `yaml:"path"`
    } `yaml:"prometheus"`
    Loki struct {
        Enabled bool   `yaml:"enabled"`
        URL     string `yaml:"url"`
    } `yaml:"loki"`
}

func LoadConfig(path string) (*Config, error) {
    config := &Config{}
    file, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(file, config)
    if err != nil {
        return nil, err
    }

    // Load secrets from environment variables (if any)
    if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
        config.Database.Password = dbPassword
    }
    if oauthClientSecret := os.Getenv("OAUTH_CLIENT_SECRET"); oauthClientSecret != "" {
        config.OAuth.ClientSecret = oauthClientSecret
    }
    if jwtSecretKey := os.Getenv("JWT_SECRET_KEY"); jwtSecretKey != "" {
        config.JWT.SecretKey = jwtSecretKey
    }

    return config, nil
}
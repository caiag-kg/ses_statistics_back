package models

// PostgreSQL configuration structure
type psqlConfig struct {
	DriverName string `yaml:"DriverName" default:"postgres"`
	Host       string `yaml:"Host" default:"localhost"`
	Port       string `yaml:"Port" default:"5432"`
	User       string `yaml:"User" default:"postgres"`
	Password   string `yaml:"Password" default:"postgres"`
	DbName     string `yaml:"DbName" default:"postgres"`
	TableName  string `yaml:"TableName" default:"users"`
}

// Server configuration structure
type serverConfig struct {
	Host string `yaml:"Host" default:"localhost"`
	Port string `yaml:"Port" default:"8080"`
}

// Application configuration structure
type AppConfig struct {
	Psql     psqlConfig
	Server   serverConfig
}
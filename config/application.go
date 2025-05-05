package config

import (
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
	"gopkg.in/yaml.v2"
)

type server struct {
	Port string `yaml:"port"`
	URL  string `yaml:"url"`
}

type db struct {
	Driver string `yaml:"driver"`
	URL    string `yaml:"url"`
}

type logFile struct {
	AccessLog string `yaml:"access-log"`
	ErrorLog  string `yaml:"error-log"`
}

type path struct {
	Public string `yaml:"public"`
	Theme  string `yaml:"theme"`
}

type MailConfig struct {
	Driver     string          `yaml:"driver"`
	Host       string          `yaml:"host"`
	Port       int             `yaml:"port"`
	Username   string          `yaml:"username"`
	Password   string          `yaml:"password"`
	From       string          `yaml:"from"`
	Encryption mail.Encryption `yaml:"encryption"`
}

type googleKey struct {
	ClientID     string `yaml:"client-id"`
	ClientSecret string `yaml:"client-secret"`
	CallbackURL  string `yaml:"callback-url"`
}

type applicationConfig struct {
	Server          *server     `yaml:"server"`
	DB              *db         `yaml:"db"`
	LogFile         *logFile    `yaml:"logfile"`
	Path            *path       `yaml:"path"`
	EMail           *MailConfig `yaml:"mail"`
	CipherKey       *string     `yaml:"cipherkey"`
	SessionLifetime *int        `yaml:"session-lifetime"`
	GoogleKey       *googleKey  `yaml:"google-key"`
}

var Server server
var DB db
var LogFile logFile
var AppConfig applicationConfig
var Path path
var EMail MailConfig
var Cipher string
var SessionLifetime int
var GoogleKey googleKey

func SetApplication(yamlFile []byte) {
	AppConfig = applicationConfig{
		Server:          &Server,
		DB:              &DB,
		LogFile:         &LogFile,
		Path:            &Path,
		EMail:           &EMail,
		CipherKey:       &Cipher,
		SessionLifetime: &SessionLifetime,
		GoogleKey:       &GoogleKey,
	}

	err := yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		log.Fatalln("Error Unmarshal application.yaml file:", err)
	}
}

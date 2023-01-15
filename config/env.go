package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
)

var ServerURL, Port, DBURL, AccessLogFile, ErrorLogFile, Cipher, PublicPath, ThemePath, GoogleClientID, GoogleClientSecret string
var SessionLifetime int64

var MailStruct struct {
	Host       string
	Port       int
	Username   string
	Password   string
	From       string
	Encryption mail.Encryption
}

func init() {
	godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Println(err)
	}

	ServerURL = os.Getenv("SERVER_URL")

	MailStruct.Host = os.Getenv("MAIL_HOST")
	MailStruct.Port = port
	MailStruct.Username = os.Getenv("MAIL_USERNAME")
	MailStruct.Password = os.Getenv("MAIL_PASSWORD")
	MailStruct.Encryption = mail.EncryptionTLS

	Port = os.Getenv("PORT")
	DBURL = os.Getenv("SQL_URL")
	AccessLogFile = os.Getenv("ACCESSLOGFILE")
	ErrorLogFile = os.Getenv("ERRORLOGFILE")

	/*
		Encryption Key

		key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	*/
	Cipher = os.Getenv("CIPHER")

	PublicPath = os.Getenv("PUBLICPATH")
	ThemePath = os.Getenv("THEMEPATH")

	session, err := strconv.ParseInt(os.Getenv("SESSIONLIFETIME"), 10, 64)
	if err != nil {
		log.Println(err)
	}

	/*
		Session Lifetime

		Here you may specify the number of minutes that you wish the session
		to be allowed to remain idle before it expires.
	*/
	SessionLifetime = session //minutes

	GoogleClientID = os.Getenv("GOOGLE_CLIENTID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENTSECRET")
}

package config

var DBUrl = "mongodb+srv://faizalom:Faiare123@cluster0.urcqt.mongodb.net/next_demo?retryWrites=true&w=majority"

var PublicPath = "public"

var ThemePath = "resources/templates/"
var ThemeView = "resources/views/"

var AccessLogFile = "logs/access.log"
var ErrorLogFile = "logs/error.log"

/*
Encryption Key

key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
*/
var Cipher = "AES-256-CBC"

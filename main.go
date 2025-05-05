// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Go-Web
// A Simple framework for web development using go language
// With auth, Middleware, session, Flash messages, CSRF, access Logs, error logs
// Database using MongoDB
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/middleware"
	"github.com/faizalom/go-web/routers"
)

func init() {
	yamlFile, err := os.ReadFile("application.yaml")
	if err != nil {
		log.Fatalln("Error loading application.yaml file: ", err)
	}

	// Create the log directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(config.LogFile.ErrorLog), os.ModePerm)
	if err != nil {
		log.Fatalln("Error creating log directory: ", err)
	}

	// Create the log directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(config.LogFile.ErrorLog), os.ModePerm)
	if err != nil {
		log.Fatalln("Error creating log directory: ", err)
	}

	config.SetApplication(yamlFile)
	lib.TemplateParseGlob(config.Path.Theme)
	lib.InitSession()
	lib.InitHash()
}

func main() {
	lib.LogErrors(config.LogFile.ErrorLog)

	fmt.Printf("Server is started on port %s\n", config.Server.Port)

	// Start a web server
	// Set your listening port here
	// Port :80 for http://
	log.Fatal(http.ListenAndServe(config.Server.Port, middleware.RequestLogger(routers.SetRoutes())))

	// To use ssl(https://) use this
	// Need certfile and keyfile to run this
	// log.Fatal(http.ListenAndServeTLS(":443", "certFile", "keyFile", middleware.RequestLogger(routers.SetRoutes())))
}

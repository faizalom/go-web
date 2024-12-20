// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Go-Web
// A Simple framework for web development using go language
// With auth, Middleware, session, Flash messages, CSRF, access Logs, error logs
// Database using MongoDB
package main

import (
	"log"
	"net/http"
	"os"

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
	config.SetApplication(yamlFile)
	lib.TemplateParseGlob(config.Path.Theme)
}

func main() {
	lib.LogErrors(config.LogFile.ErrorLog)

	// Start a web server
	// Set your listening port here
	// Port :80 for http://
	log.Fatal(http.ListenAndServe(config.Server.Port, middleware.RequestLogger(routers.SetRoutes())))

	// To use ssl(https://) use this
	// Need certfile and keyfile to run this
	// log.Fatal(http.ListenAndServeTLS(":443", "certFile", "keyFile", middleware.RequestLogger(routers.SetRoutes())))
}

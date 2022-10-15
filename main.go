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

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/lib"
	"github.com/faizalom/go-web/middleware"
	"github.com/faizalom/go-web/routers"
)

func main() {
	lib.LogErrors(config.ErrorLogFile)
	// Start a web server
	// Set your listening port here
	// Port :80 for http://
	log.Fatal(http.ListenAndServe(":8080", middleware.RequestLogger(routers.SetRoutes())))

	// To use ssl(https://) use this
	// Need certfile and keyfile to run this
	// log.Fatal(http.ListenAndServeTLS(":443", "certFile", "keyFile", middleware.RequestLogger(routers.SetRoutes())))
}

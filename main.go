// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	log.Fatal(http.ListenAndServe(":8080", middleware.RequestLogger(routers.SetRoutes())))
}

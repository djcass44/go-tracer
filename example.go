/*
 *    Copyright 2020 Django Cass
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package main

import (
	"flag"
	"fmt"
	"github.com/djcass44/go-tracer/tracer"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.DebugLevel)

	port := flag.Int("port", 8080, "port to run on")

	flag.Parse()

	// setup http handlers
	http.HandleFunc("/api/v1/my-resource", tracer.NewFunc(func(w http.ResponseWriter, r *http.Request) {
		// write the request id into the response
		_, _ = w.Write([]byte(tracer.GetRequestId(r)))
	}))
	// start http server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

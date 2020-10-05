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

package tracer

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type traceableRoundTripper struct {
	Proxied http.RoundTripper
}

func (rt *traceableRoundTripper) RoundTrip(r *http.Request) (res *http.Response, err error) {
	// inject the request id header
	if rawId := r.Context().Value("id"); rawId != nil {
		id := fmt.Sprintf("%v", rawId) // safe conversion to string
		if id != "" {
			r.Header.Add("X-Request-ID", id)
		}
	}
	res, err = rt.Proxied.RoundTrip(r)
	return
}

// Apply modifies an http client's transport to inject the X-Request-ID header
func Apply(c *http.Client) {
	log.Debug("injecting request tracing into http client")
	rt := traceableRoundTripper{http.DefaultTransport}
	c.Transport = &rt
}

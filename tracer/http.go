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
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// SetRequestId returns a shallow copy of the origin request, with a request ID created or extracted from the X-Request-ID header
func SetRequestId(r *http.Request) *http.Request {
	id := r.Header.Get("X-Request-ID")
	// if we didn't get an id, create one
	if id == "" {
		id = uuid.New().String()
		log.WithField("id", id).Debugf("failed to locate existing request ID, generating a new one...")
	}
	// update the request context
	return r.WithContext(context.WithValue(r.Context(), "id", id))
}

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
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Request-ID", r.Context().Value("id").(string))
}

func TestMiddleware(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://example.org/test", nil)
	req.Header.Add("X-Request-ID", "test-request-id")
	w := httptest.NewRecorder()

	h := NewHandler(&handler{})
	h.ServeHTTP(w, req)

	resp := w.Result()

	assert.EqualValues(t, "test-request-id", resp.Header.Get("X-Request-ID"))
}

func TestNewFunc(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Request-ID", r.Context().Value("id").(string))
	}
	req := httptest.NewRequest(http.MethodGet, "http://example.org/test", nil)
	w := httptest.NewRecorder()

	NewFunc(handler)(w, req)

	resp := w.Result()

	assert.NotEmpty(t, resp.Header.Get("X-Request-ID"))
}

func TestNewFuncWithHeader(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Request-ID", r.Context().Value("id").(string))
	}
	req := httptest.NewRequest(http.MethodGet, "http://example.org/test", nil)
	req.Header.Add("X-Request-ID", "test")
	w := httptest.NewRecorder()

	NewFunc(handler)(w, req)

	resp := w.Result()

	assert.Equal(t, "test", resp.Header.Get("X-Request-ID"))
}

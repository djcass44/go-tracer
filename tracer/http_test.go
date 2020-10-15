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
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSetRequestId(t *testing.T) {
	request := &http.Request{}
	newRequest := SetRequestId(request)

	assert.Empty(t, request.Context().Value("id"))
	assert.NotEmpty(t, newRequest.Context().Value("id"))
}

func TestSetRequestIdFromHeader(t *testing.T) {
	headers := http.Header{}
	headers.Add("X-Request-ID", "my-request-id")
	request := &http.Request{Header: headers}
	newRequest := SetRequestId(request)

	assert.Empty(t, request.Context().Value("id"))
	assert.Equal(t, "my-request-id", newRequest.Context().Value("id"))
}

func TestGetRequestIdNoContext(t *testing.T) {
	request := &http.Request{}

	assert.Empty(t, GetRequestId(request))
}

func TestGetRequestId(t *testing.T) {
	request := &http.Request{}
	r := request.WithContext(context.WithValue(context.TODO(), "id", "test"))

	assert.Equal(t, "test", GetRequestId(r))
}

func TestGetContextId(t *testing.T) {
	ctx := context.WithValue(context.Background(), "id", "test")

	assert.EqualValues(t, "test", GetContextId(ctx))
}

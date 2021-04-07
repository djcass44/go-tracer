# Go Tracer

This library provides a method for simplifying request tracing and correlation between GoLang services.

#### What it does

* Extracts the `X-Request-ID` value or creates a new ID from a UUIDv4
* Adds an entry to the `http.Request` context
* Injects the `X-Request-ID` header to outgoing HTTP requests

This allows you to 'follow' requests as they move between various services as they will all have the same ID.


#### Getting started

Have a look at [example.go](example.go) for a simple example.

```shell script
go get github.com/djcass44/go-tracer
```

Using handler functions
```go
http.HandleFunc("/something", tracer.NewFunc(myFunc))
```

Using handlers
```go
type myHandler struct {}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}
// later on
http.Handle("/something", tracer.NewHandler(myHandler{}))
```

**Accessing the RequestID**

The ID can be pulled from the context directly, or using a helper function

Helper:
```go
requestId := tracer.GetRequestId(request)
```

Direct:
```go
requestId := request.Context().Value(tracer.ContextKeyID).(string)
```

**Enabling outgoing RequestID injection**

This needs to be configured for all instances of `http.Client`

*Note: this will replace the `http.Transport` of the client*

```go
// enable outgoing-tracing for the default http client
tracer.Apply(http.DefaultClient)

var myClient = &http.Client{}
tracer.Apply(myClient)
```

#### Contribution

Contributions are more than welcome! Feel free to open a PR, or an issue if you have questions.
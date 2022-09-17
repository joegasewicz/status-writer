# Status Writer
Utility pkg that helps get the status code from the response writer

## Install
```
go get github.com/joegasewicz/status-writer
```

## Setup
Example shows setup for a middleware function. This could be a logging 
function or similar whereby you require access to the status code from
`http.ResponseWriter`
```go
func MyMiddleware(h http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        sw := status_writer.New(w) // Here we override http.ResponseWriter's `WriteHeader` function
		h(sw, r)
        statusCode := sw.Status // Now work with the status code
    }	
}
```

### Usage
Now you can set the status code as usual via the response writer in your 
handlers
```go
func (n *Notice) Get(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
```
# Status Writer
Utility pkg that helps get the status code from the response writer

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

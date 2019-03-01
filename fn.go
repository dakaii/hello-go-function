package hello
import (
   "fmt"
   "net/http"
   "go.opencensus.io/plugin/ochttp"
)
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    fn := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello 世界")
    }
    traced := &ochttp.Handler{
        Handler: http.HandlerFunc(fn),
    }
    traced.ServeHTTP(w, r)
}

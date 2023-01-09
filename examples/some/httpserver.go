package some

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// 请求拦截器
func handleIterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(">>> Receive: Method[%s] IP[%s]\n", r.Method, getIP(r))
		h(w, r)
	}
}

var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

func getIP(r *http.Request) string {
	var ip string

	if xrip := r.Header.Get(xRealIP); xrip != "" {
		ip = xrip
	} else if xff := r.Header.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	} else if clientIP, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		ip = clientIP
	}

	return ip
}

func handler1(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)
	fmt.Println(ip, xForwardedFor, xRealIP)
	fmt.Fprintf(w, "Hello handler1"+ip)
}

type handler2 struct {
	content string
}

func (ih *handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"message": "hello world"}`))
	fmt.Fprintf(w, ih.content)
}

// -----------Chain Interceptor-------------

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// -----------Chain Interceptor-------------

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, chain interceptor")
}

// TestHTTPServer ...
func TestHTTPServer() {
	fmt.Println("-----TestHTTPServer-----")
	// 1
	http.HandleFunc("/", handleIterceptor(handler1))
	// 2
	http.Handle("/test", &handler2{content: "Hello handler2"})

	fs := http.FileServer(http.Dir("/Users/tianwen/go"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/404", http.NotFoundHandler())

	//chain interceptor
	http.HandleFunc("/chain", Chain(hello, Method("GET"), Logging()))

	//Listen
	fmt.Println("ListenAndServe :8080")
	http.ListenAndServe(":8080", nil)
	fmt.Println("-----TestHTTPServer-----")
}

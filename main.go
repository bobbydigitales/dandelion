package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

// LoggingHandler is a custom http.Handler that logs requests and delegates them to the underlying handler.
type LoggingHandler struct {
	handler http.Handler
}

func (lh *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving file: %s", r.URL.Path)
	lh.handler.ServeHTTP(w, r)
}

func getAddressString(portNumber int) string {
	return (":" + strconv.Itoa(portNumber))
}

func portIsInUse(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return true
	}
	conn.Close()
	return false
}

func main() {
	dir := "./"
	startPort := 8000
	const maxAttempts = 10

	fmt.Println("Starting...")

	// Configure TLS with the self-signed certificate and private key
	// tlsConfig := &tls.Config{
	// 	MinVersion:               tls.VersionTLS12,
	// 	PreferServerCipherSuites: true,
	// 	InsecureSkipVerify:       true,
	// 	Certificates:             make([]tls.Certificate, 1),
	// }

	// Load the certificate and private key
	// cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	// if err != nil {
	// 	log.Fatalf("Failed to load certificate and key: %v", err)
	// }
	// tlsConfig.Certificates[0] = cert

	port := startPort
	for attempts := 0; attempts < maxAttempts; attempts++ {
		fmt.Printf("Trying port %d", port)

		if portIsInUse(port) {
			fmt.Println("...port in use!")
			port++
			continue
		}

		// If the port is not in use, bind your server here
		// For example:
		addr := ":" + strconv.Itoa(port)
		log.Printf("\nServing %s on %s using HTTP/2...", dir, addr)

		// Configure the HTTP/2 server
		server := &http.Server{
			Addr:    addr,
			Handler: &LoggingHandler{http.FileServer(http.Dir(dir))},
			// TLSConfig: tlsConfig,
		}
		fmt.Println("Configured...")

		fmt.Println("Serving...")
		// Start the server
		server.ListenAndServe()
		if err := server.ListenAndServeTLS("", ""); err != nil {
			log.Fatalf("Server failed: %v", err)
		}

		fmt.Println("Server started on port:", port)

	}
	log.Fatalf("Could not find an open port after %d attempts", maxAttempts)

}

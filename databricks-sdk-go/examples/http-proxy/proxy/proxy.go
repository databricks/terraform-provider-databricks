package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	createX509Certificate()
	defer cleanup()

	util := newCertUtil()
	runCommands(util.GetRegisterCertificateCommands("proxy.crt"))
	defer runCommands(util.GetDeregisterCertificateCommands("proxy.crt"))

	server := startServer()
	defer server.Close()

	waitForSigint()
}

func waitForSigint() {
	log.Printf("Proxy server started. Press Ctrl+C to stop.")
	log.Printf("Waiting for SIGINT...")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	<-sigChan
}

func createX509Certificate() {
	log.Printf("Creating X509 certificate and private key...")
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Your Organization"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,
		// KeyEncipherment is only used for RSA keys.
		KeyUsage:/* x509.KeyUsageKeyEncipherment | */ x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
		DNSNames:              []string{"localhost"},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		panic(err)
	}
	privateKey, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		panic(err)
	}

	certOut, err := os.Create("proxy.crt")
	if err != nil {
		panic(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, err := os.OpenFile("proxy.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKey})
	keyOut.Close()
}

func cleanup() {
	log.Printf("Cleaning up X509 certificate and private key...")
	os.Remove("proxy.crt")
	os.Remove("proxy.key")
}

func runCommands(cmds [][]string) {
	for _, c := range cmds {
		cmd := exec.Command(c[0], c[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

// The following code is a simple HTTPS proxy server.
func handleTunneling(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request: %s %s", r.Method, r.URL.String())
	if r.URL.Path == "/ping" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
		return
	}

	destConn, err := net.Dial("tcp", r.URL.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "HTTP Server does not support hijacking", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, "Error hijacking connection: "+err.Error(), http.StatusInternalServerError)
		return
	}

	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func startServer() *http.Server {
	log.Printf("Starting server...")
	server := &http.Server{
		Addr:    "localhost:8443",
		Handler: http.HandlerFunc(handleTunneling),
	}
	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		panic(err)
	}

	// Start the server
	log.Printf("Starting proxy server on https://%s", server.Addr)
	go func() {
		err := server.ListenAndServeTLS("proxy.crt", "proxy.key")
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	log.Printf("Waiting for server to start...")
	for {
		_, err := http.Get("https://localhost:8443/ping")
		if err == nil {
			break
		}
		log.Printf("Sleeping for 1 second...")
		time.Sleep(time.Second)
	}
	log.Printf("Server started.")
	return server
}

package server

import (
	"net/http"
	"time"
)

func (server *Server) Run(port string, handler http.Handler, readTimeout, writeTimeout int) error {
	server.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
	}

	return server.httpServer.ListenAndServe()
}

func (server *Server) RunTLS(port string, handler http.Handler, readTimeout, writeTimeout int, certFilePath, keyFilePath string) error {
	server.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
	}

	return server.httpServer.ListenAndServeTLS(certFilePath, keyFilePath)
}

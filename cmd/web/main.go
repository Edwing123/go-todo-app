package main

import (
	"crypto/tls"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Get command-line flags.
	config, err := getConfigFlags()
	if err != nil {
		errLogger.Fatalln(err)
	}

	// Get templates map.
	templates, err := createTemplatesMap("ui/html")
	if err != nil {
		errLogger.Fatalln(err)
	}

	// Get database connections pool.
	db, err := createDatabasePool(config.dsn)
	if err != nil {
		errLogger.Fatalln(err)
	}
	defer db.Close()

	// Setup application dependencies.
	app := &application{
		templates:      templates,
		infoLogger:     infoLogger,
		errLogger:      errLogger,
		sessionManager: createSessionManager(db),
	}

	// TSL config for the server.
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.CurveP256, tls.X25519},
	}

	// Setup server.
	server := &http.Server{
		Addr:      config.addr,
		Handler:   app.getRouter(),
		ErrorLog:  errLogger,
		TLSConfig: tlsConfig,

		// Server connection timeout settings.
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	// Start server using tls certificate and private key.
	infoLogger.Printf("Listening on %s\n", config.addr)
	errLogger.Fatalln(server.ListenAndServeTLS(config.tlsCertPath, config.privKeyPath))
}

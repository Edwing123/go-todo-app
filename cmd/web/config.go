package main

import (
	"errors"
	"flag"
	"strings"
)

// Define and return command-line flags for the application.
func getConfigFlags() (*configFlags, error) {
	addr := flag.String("addr", "", "Set server netowrk address")
	dsn := flag.String("dsn", "", "Set database data source name")
	tlsCertPath := flag.String("tlsCertPath", "", "Path to the TLS certificate file")
	privKeyPath := flag.String("privKeyPath", "", "Path of the private key associated with the TLS certificate")
	flag.Parse()

	// Return an error if the flags values are empty
	if strings.TrimSpace(*addr) == "" || strings.TrimSpace(*dsn) == "" {
		return nil, errors.New("all command-line flags are required")
	}

	return &configFlags{
		addr:        *addr,
		dsn:         *dsn,
		tlsCertPath: *tlsCertPath,
		privKeyPath: *privKeyPath,
	}, nil
}

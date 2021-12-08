package main

import (
	"log"
	"os"
)

var (
	infoLogger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	errLogger  = log.New(os.Stdout, "[ERROR] ", log.LstdFlags)
)

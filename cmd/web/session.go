package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
)

type contextKey string

var (
	contextKeyIsUserAuthenticated = contextKey("IsUserAuthenticated")
)

// Create and configure session manager.
func createSessionManager(db *sql.DB) *scs.SessionManager {
	sessionManager := scs.New()

	// Setting some field to change session manager
	// behavior, these settings are subjected to change
	// in the future.
	sessionManager.Lifetime = time.Hour
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true

	// Set session data store.
	sessionManager.Store = mysqlstore.New(db)

	return sessionManager
}

// Create session store for session manager.
// func createSessionStore(db *sql.DB) scs.Store {
// 	return mysqlstore.New(db, time.Hour)
// }

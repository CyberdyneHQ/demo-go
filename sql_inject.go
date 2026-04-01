package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

// SQL injection via string concatenation
func getUserByName(db *sql.DB, r *http.Request) (*sql.Rows, error) {
	username := r.URL.Query().Get("username")
	query := "SELECT * FROM users WHERE username = '" + username + "'"
	return db.Query(query)
}

// SQL injection via fmt.Sprintf
func deleteUser(db *sql.DB, userID string) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", userID)
	_, err := db.Exec(query)
	return err
}

// Unvalidated redirect — open redirect vulnerability
func handleLogin(w http.ResponseWriter, r *http.Request) {
	redirectURL := r.URL.Query().Get("redirect")
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// Path traversal via user input
func serveFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	http.ServeFile(w, r, "/uploads/"+filename)
}

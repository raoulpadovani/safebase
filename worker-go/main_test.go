package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Vérifie que l'endpoint /backup rejette les méthodes autres que POST.
func TestHandleBackupRejectsNonPost(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/backup", nil)
	rec := httptest.NewRecorder()

	handleBackup(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("attendu statut %d, obtenu %d", http.StatusMethodNotAllowed, rec.Code)
	}
}

// Vérifie que l'endpoint /backup rejette un corps JSON invalide.
func TestHandleBackupRejectsInvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/backup", nil)
	rec := httptest.NewRecorder()

	handleBackup(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("attendu statut %d, obtenu %d", http.StatusBadRequest, rec.Code)
	}
}

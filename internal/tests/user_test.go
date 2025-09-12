package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/DATA-DOG/go-sqlmock"
)

// setup mock db
func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	return db, mock
}

func TestGetUser_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	password := "secret"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery("SELECT id, name, email, password, role, verified FROM users WHERE email = ?").
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "verified"}).
			AddRow(1, "Test User", "test@example.com", string(hashedPassword), "user", true))

	h := &handlers.handlers{DB: db}

	body := `{"email":"test@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(body))
	w := httptest.NewRecorder()

	h.GetUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("expected %d, got %d", http.StatusAccepted, resp.StatusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestGetUser_WrongPassword(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	password := "secret"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery("SELECT id, name, email, password, role, verified FROM users WHERE email = ?").
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "verified"}).
			AddRow(1, "Test User", "test@example.com", string(hashedPassword), "user", true))

	h := &handlers.handlers{DB: db}

	body := `{"email":"test@example.com","password":"wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(body))
	w := httptest.NewRecorder()

	h.GetUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

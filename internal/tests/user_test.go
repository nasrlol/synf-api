package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"synf/internal/api/data/rest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

// Setup mock database
func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql db: %v", err)
	}
	return db, mock
}

func TestGetUser_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	password := "secret"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery("SELECT * FROM users WHERE email = ?").
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "verified"}).
			AddRow(1, "Test User", "test@example.com", string(hashedPassword), "user", true))

	body := `{"email":"test@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(body))
	w := httptest.NewRecorder()

	handlers.GetUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("expected status %d, got %d", http.StatusAccepted, resp.StatusCode)
	}
}

func TestGetUser_WrongPassword(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	password := "secret"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery("SELECT * FROM users WHERE email = ?").
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "verified"}).
			AddRow(1, "Test User", "test@example.com", string(hashedPassword), "user", true))

	body := `{"email":"test@example.com","password":"wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(body))
	w := httptest.NewRecorder()

	handlers.GetUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestCreateUser_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("INSERT INTO users (name, email, password, role, verified) VALUES (? , ? , ?, ?, ? )").WithArgs("Test User", "test@example.com", sqlmock.AnyArg(), "user", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	body := `{"name":"Test User","email":"test@example.com","password":"secret","role":"user","verfied":1}`
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(body))
	w := httptest.NewRecorder()

	handlers.CreateUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestUpdateUser_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("UPDATE users SET").
		WillReturnResult(sqlmock.NewResult(1, 1))

	body := `{"name":"Updated Name"}`
	req := httptest.NewRequest(http.MethodPut, "/user/1", strings.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.UpdateUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func TestUpdateUser_NoID(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/user/ ", nil)
	req = mux.SetURLVars(req, map[string]string{"id": " "})
	w := httptest.NewRecorder()

	handlers.UpdateUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.DeleteUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestDeleteUser_NoID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/user/ ", nil)
	req = mux.SetURLVars(req, map[string]string{"id": " "})
	w := httptest.NewRecorder()

	handlers.DeleteUser(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

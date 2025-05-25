package web

import (
	"canchitas-libres-transaction/internal/pgk/domain"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Transaction, error)
	GetByID(ctx context.Context, id string) (domain.Transaction, error)
	Add(ctx context.Context, user domain.Transaction) error
	Delete(ctx context.Context, id string) error
	Update(id int, user domain.Transaction) error
}
type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

var (
	getAllRegexp = regexp.MustCompile(`^\/transaction\/?$`)
	getOneRegexp = regexp.MustCompile(`^\/transaction\/([a-zA-Z0-9\-]+)$`)
)

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		handler.CreateTransaction(w, r)
		return
	case r.Method == http.MethodGet && getAllRegexp.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		handler.GetAllTransactions(w, r)
		return
	case r.Method == http.MethodGet && getOneRegexp.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		handler.GetTransactionByID(w, r)
		return
	case r.Method == http.MethodDelete:
		handler.DeleteTransaction(w, r)
		return
	case r.Method == http.MethodPut:
		handler.UpdateTransaction(w, r)
		return
	default:
		http.NotFound(w, r)
		return
	}
}

func (handler *Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to get transactions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (handler *Handler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	matches := getOneRegexp.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	id := matches[1] // UUID como string

	tx, err := handler.Service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}

func (handler *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tx domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := handler.Service.Add(r.Context(), tx)
	if err != nil {
		http.Error(w, "Failed to insert transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Transaction recorded"))
}

func (handler *Handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	matches := getOneRegexp.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	id := matches[1] // UUID como string

	err := handler.Service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction deleted successfully"))
}

func (handler *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")
	w.WriteHeader(http.StatusCreated)
}

package web

import (
	"canchitas-libres-transaction/internal/pgk/domain"
	"fmt"
	"net/http"
	"regexp"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	GetByID(id int) (domain.Transaction, error)
	Add(user domain.Transaction) error
	Delete(id int) error
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
	getOneRegexp = regexp.MustCompile(`^\/transaction\/(\d+)$`)
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
	fmt.Println("getall")
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getbyid")
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create")
	w.WriteHeader(http.StatusCreated)
}

func (handler *Handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")
	w.WriteHeader(http.StatusCreated)
}

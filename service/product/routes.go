package product

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dodo/ecom/types"
	"github.com/dodo/ecom/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleProductsList).Methods("GET")
	router.HandleFunc("/products/{id}", h.handleProductById).Methods("GET")
	router.HandleFunc("/products", h.handleProductCreate).Methods("POST")
}

func (h *Handler) handleProductsList(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetListProducts()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleProductById(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, fmt.Sprint("handle product by id"))
}

func (h *Handler) handleProductCreate(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, fmt.Sprint("handle product create"))
}

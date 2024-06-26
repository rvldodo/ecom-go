package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/dodo/ecom/lib/bcrypt"
	"github.com/dodo/ecom/lib/cookies"
	"github.com/dodo/ecom/lib/jwt"
	"github.com/dodo/ecom/types"
	"github.com/dodo/ecom/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParesJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err == nil && u.ID == 0 {
		utils.WriteError(
			w,
			http.StatusNotFound,
			fmt.Errorf("email %s not registered yet", user.Email),
		)
		return
	}

	if !bcrypt.ComparePassword(u.Password, []byte(user.Password)) {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("Invalid password"),
		)
		return
	}

	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	v := &cookies.CookieData{
		UserID: u.ID,
		Email:  u.Email,
	}
	if err = cookies.SetCookie(v, w, r); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParesJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err == nil && u.ID != 0 {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("user with email %s already exist", user.Email),
		)
		return
	}

	hashedPassword, err := bcrypt.HashPassword(user.Password)

	err = h.store.CreateUser(&types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

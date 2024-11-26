package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/motty93/clean-architecture/internal/usecase"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	user, err := uh.usecase.GetUserByID(ctx, id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

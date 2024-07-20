package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pedrogutierresbr/pos-goexpert/9-apis/internal/dto"
	"github.com/pedrogutierresbr/pos-goexpert/9-apis/internal/entity"
	"github.com/pedrogutierresbr/pos-goexpert/9-apis/internal/infa/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

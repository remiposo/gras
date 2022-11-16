package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/remiposo/gras/usecase"
)

type UserHandler struct {
	userUc *usecase.UserUc
}

func NewUserHandler(userUc *usecase.UserUc) *UserHandler {
	return &UserHandler{
		userUc: userUc,
	}
}

type UserRegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserRegisterResp struct {
	ID string `json:"id"`
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &UserRegisterReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		RespJSON(w, &ErrResp{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	data := usecase.UserRegisterDto{
		Name:     req.Name,
		Password: req.Password,
		Role:     req.Role,
	}
	id, err := uh.userUc.Register(ctx, data)
	if err != nil {
		RespJSON(w, &ErrResp{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	resp := &UserRegisterResp{
		ID: id,
	}
	RespJSON(w, resp, http.StatusOK)
}

package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/remiposo/gras/usecase"
)

type TaskHandler struct {
	taskUc *usecase.TaskUc
}

func NewTaskHandler(taskUc *usecase.TaskUc) *TaskHandler {
	return &TaskHandler{
		taskUc: taskUc,
	}
}

type TaskAddReq struct {
	Title string `json:"title"`
}

type TaskAddResp struct {
	ID string `json:"id"`
}

func (th *TaskHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &TaskAddReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		RespJSON(w, &ErrResp{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	data := usecase.TaskAddDto{
		Title: req.Title,
	}
	id, err := th.taskUc.Add(ctx, data)
	if err != nil {
		RespJSON(w, &ErrResp{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	resp := &TaskAddResp{
		ID: id,
	}
	RespJSON(w, resp, http.StatusOK)
}

type TaskShowResp struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskListResp []TaskShowResp

func (th *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := th.taskUc.List(ctx)
	if err != nil {
		RespJSON(w, &ErrResp{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	resp := make(TaskListResp, 0, len(data))
	for _, d := range data {
		showResp := TaskShowResp{
			ID:     d.ID,
			Title:  d.Title,
			Status: d.Status,
		}
		resp = append(resp, showResp)
	}
	RespJSON(w, &resp, http.StatusOK)
}

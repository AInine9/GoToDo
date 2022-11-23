package handler

import (
	"backend/cmd/api/usecase"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ItemHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
	Update(http.ResponseWriter, *http.Request, httprouter.Params)
}

type itemHandler struct {
	itemUseCase usecase.ItemUseCase
}

func NewItemHandler(iu usecase.ItemUseCase) ItemHandler {
	return &itemHandler{
		itemUseCase: iu,
	}
}

func (ih itemHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	items, err := ih.itemUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (ih itemHandler) Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	name := r.FormValue("name")
	status, _ := strconv.Atoi(r.FormValue("status"))

	err := ih.itemUseCase.Create(status, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (ih itemHandler) Update(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	status, _ := strconv.Atoi(r.FormValue("status"))

	err := ih.itemUseCase.Update(id, status, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

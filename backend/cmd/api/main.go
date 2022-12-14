package main

import (
	"backend/cmd/api/config"
	"backend/cmd/api/infrastructure/persistence"
	"backend/cmd/api/interface/handler"
	"backend/cmd/api/usecase"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	itemPersistence := persistence.NewItemPersistence(config.Connect())
	itemUseCase := usecase.NewItemUseCase(itemPersistence)
	itemHandler := handler.NewItemHandler(itemUseCase)

	router := httprouter.New()
	router.GET("/api/items", itemHandler.Index)
	router.POST("/api/items", itemHandler.Create)
	router.PUT("/api/items", itemHandler.Update)

	http.ListenAndServe(":8000", &Server{router})
	log.Fatal(http.ListenAndServe(":8000", router))
}

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET POST PUT")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}

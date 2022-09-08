package handler

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gomoku/pkg/logger"
	"net/http"
)

const (
	mainPageEndPoint = "/"
)

type Handler struct {
	upgrader *websocket.Upgrader
}

func NewHandler() *Handler {
	return &Handler{
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	logger.Infof("server listen endpoint for main page: %s", mainPageEndPoint)
	router.HandleFunc(mainPageEndPoint, h.mainPage)

	return router
}

func (h *Handler) mainPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

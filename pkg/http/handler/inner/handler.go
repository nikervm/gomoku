package inner

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gomoku/pkg/http/players"
	"gomoku/pkg/http/websocket_conn"
	"gomoku/pkg/logger"
	"net/http"
	"sync"
)

const (
	createSession = "/{nickname}"
)

type Handler struct {
	upgrader *websocket.Upgrader
	players  *sync.Map
}

func NewHandler() *Handler {
	return &Handler{
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		players: new(sync.Map),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	logger.Infof("server listen endpoint for main page: %s", createSession)
	router.HandleFunc(createSession, h.createSession)

	return router
}

func (h *Handler) createSession(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]

	ws, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Errorf("Server error ws [%v]", err)
		return
	}

	player := players.NewPlayer(websocket_conn.New(ws), h, nickname).Run()
	h.players.Store(nickname, player)

	return
}

func (h *Handler) GetPlayers() *sync.Map {
	return h.players
}

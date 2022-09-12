package players

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gomoku/pkg/http/handler"
	"gomoku/pkg/http/websocket_conn"
	"gomoku/pkg/logger"
	"sync"
)

type rPlayer struct {
	nickname string
	ws       websocket_conn.WebsocketConn
	ctx      context.Context
	cancel   context.CancelFunc
	logger   *logrus.Logger
	mx       *sync.RWMutex
	handler  handler.IHandler
}

func NewPlayer(ws websocket_conn.WebsocketConn, handler handler.IHandler, nickname string) Player {
	ctx, cancel := context.WithCancel(context.Background())
	return &rPlayer{
		ws:       ws,
		nickname: nickname,
		ctx:      ctx,
		cancel:   cancel,
		mx:       new(sync.RWMutex),
		logger:   logger.With("nickname", nickname).Logger,
		handler:  handler,
	}
}

func (p *rPlayer) Run() Player {
	logger.Infof("run player [%s]", p.nickname)

	go p.reader()

	return p
}

func (p *rPlayer) Close() {
	p.mx.Lock()
	defer p.mx.Unlock()

	p.logger.Infof("remove player [%s]", p.nickname)
	p.handler.GetPlayers().Delete(p.nickname)
	p.cancel()
	p.ws.Close()
}

func (p *rPlayer) reader() {
	channel := p.ws.Read()
	for {
		select {
		case <-p.ctx.Done():
			return
		case data, ok := <-channel:
			if !ok {
				p.Close()
				return
			}

			p.logger.Infof("get message [%s]", string(data))

			playerMsg, err := parseMsg(data)
			if err != nil {
				p.logger.Errorf("parse message error [%v]", err)
				continue
			}

			_ = playerMsg
		}
	}
}

func parseMsg(data []byte) (*PlayerMessage, error) {
	var AgMsg PlayerMessage

	if err := json.Unmarshal(data, &AgMsg); err != nil {
		return nil, err
	}

	return &AgMsg, nil
}

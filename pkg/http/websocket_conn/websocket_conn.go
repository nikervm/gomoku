package websocket_conn

import (
	"context"
	"github.com/gorilla/websocket"
	"gomoku/pkg/logger"
	"sync"
)

type WebsocketConn interface {
	Read() chan []byte
	Write() chan []byte
	Close()
	Closed() bool
}

type websocketConn struct {
	ws        *websocket.Conn
	writeChan chan []byte
	readChan  chan []byte
	ctx       context.Context
	cancel    context.CancelFunc
	closed    bool
	mx        *sync.RWMutex
}

func New(conn *websocket.Conn) WebsocketConn {
	ctx, cancel := context.WithCancel(context.Background())

	wsConn := &websocketConn{
		ws:        conn,
		writeChan: make(chan []byte, 1),
		readChan:  make(chan []byte, 1),
		ctx:       ctx,
		cancel:    cancel,
		mx:        new(sync.RWMutex),
	}

	go wsConn.reader()
	go wsConn.writer()

	return wsConn
}

func (w *websocketConn) reader() {
	defer w.Close()

	for {
		msgType, msg, err := w.ws.ReadMessage()
		if msgType < 0 {
			logger.Warnf("websocket message type incorrect [%d] with message [%v]", msgType, string(msg))
			return
		}

		if err != nil {
			logger.Errorf("websocket read err [%v]", err)
			return
		}

		w.readChan <- msg
	}
}

func (w *websocketConn) writer() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case data := <-w.writeChan:
			if err := w.ws.WriteMessage(websocket.TextMessage, data); err != nil {
				logger.Errorf("websocket write err [%v]", err)
			}
		}
	}
}

func (w *websocketConn) Read() chan []byte {
	return w.readChan
}

func (w *websocketConn) Write() chan []byte {
	return w.writeChan
}

func (w *websocketConn) Close() {
	w.mx.Lock()
	defer w.mx.Unlock()

	if !w.closed {
		logger.Info("close websocket connection")
		close(w.readChan)

		w.closed = true
		w.cancel()
		w.ws.Close()
		close(w.writeChan)
	}
}

func (w *websocketConn) Closed() bool {
	w.mx.RLock()
	defer w.mx.RUnlock()

	return w.closed
}

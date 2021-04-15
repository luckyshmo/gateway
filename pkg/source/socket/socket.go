package socket

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/luckyshmo/gateway/config"
	"github.com/luckyshmo/gateway/models"
	"github.com/sirupsen/logrus"
)

type SocketSource struct {
	*websocket.Conn
}

var authStr = `{
	"cmd": "auth_req",
	 "login": "root",
	 "password": "123"
  }`

func (ss *SocketSource) ReadData(ch chan<- models.RawData) error {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ss.WriteMessage(websocket.TextMessage, []byte(authStr))

	defer ss.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := ss.ReadMessage()
			if err != nil {
				logrus.Error(err)
				return
			}
			ch <- models.RawData{
				Id:   uuid.New(),
				Time: time.Now().UTC(),
				Data: message,
			}

		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			logrus.Info("DONE")
			return nil
		case <-interrupt:
			logrus.Info("Socket interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := ss.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logrus.Info("write close:", err)
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

func NewSocketSource(cfg *config.Config) *SocketSource {

	u := url.URL{Scheme: "ws", Host: cfg.SocketHost}
	logrus.Info(fmt.Sprintf("connecting to %s", u.String()))

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.Warn("dial:", err)
	}
	return &SocketSource{c}
}

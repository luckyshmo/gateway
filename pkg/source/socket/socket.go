package socket

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/luckyshmo/gateway/models"
	"github.com/sirupsen/logrus"
)

type SocketSource struct {
	*websocket.Conn
}

func (ss *SocketSource) ReadData(ch chan<- models.RawData) error { //data 41 + data 51 + same devEui
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	authStr := `{
		"cmd": "auth_req",
	 	"login": "root",
	 	"password": "123"
	  }`

	ss.WriteMessage(websocket.TextMessage, []byte(authStr))

	defer ss.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := ss.ReadMessage()
			if err != nil {
				// log.Println("read:", err)
				return
			}
			// log.Printf("recv: %s", message)
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
			return nil
		// case t := <-ticker.C:
		// 	err := ss.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := ss.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
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

func NewSocketSource() *SocketSource {

	// send(new Gson().toJson(AuthWS("auth_req", "root", "123")))
	// new URI("ws://89.109.190.198:8003")

	u := url.URL{Scheme: "ws", Host: "89.109.190.198:8003"}
	logrus.Info(fmt.Sprintf("connecting to %s", u.String()))

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.Warn("dial:", err)
	}
	return &SocketSource{c}
}

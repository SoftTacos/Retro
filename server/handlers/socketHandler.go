package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	dbModels "github.com/softtacos/retrobot/models/db"
)

func NewSocketManager() *SocketManager {
	return &SocketManager{
		connections: make(map[*SocketClient]bool),
	}
}

type SocketManager struct {
	connections map[*SocketClient]bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//if they hit this page, that means they are requesting the socket
func (sh *SocketManager) AddConnection(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Vis Socket Connected!")
	client := NewSocketClient(sh, ws)
	sh.connections[client] = true
	client.Listen()
}

func (sh *SocketManager) CloseConnection(client *SocketClient) error {
	delete(sh.connections, client)
	return nil
}

func (sh *SocketManager) Broadcase(message []byte) error {
	var err error
	for client := range sh.connections {
		err = client.Send(message)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func NewSocketClient(manager *SocketManager, conn *websocket.Conn) *SocketClient {
	client := &SocketClient{
		conn: conn,
	}
	conn.SetCloseHandler(SetupCloseHandler(manager, client))

	return client
}

type SocketClient struct {
	user    *dbModels.User
	manager *SocketManager
	conn    *websocket.Conn
	//chan to communicate with socket manager?
}

//this only sends, for now we don't need to listen to the socket
//messageType is an int and can be 1:Text([]uint8|[]byte), 2:binary(), 8:closemessage, 9:ping message, 10:pong message?
func (sc *SocketClient) Listen() {
	for open := true; open; {
		messageType, p, err := sc.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("MESSAGE RECIEVED: ", messageType, string(p))
		// if err = sc.conn.WriteMessage(1, []byte("DAVAI DAVAI")); err != nil {
		// 	log.Println(err)
		// 	return
		// }
		sc.Send(p)
	}
}

func (sc *SocketClient) Send(message []byte) error {
	if err := sc.conn.WriteMessage(1, message); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SetupCloseHandler(manager *SocketManager, client *SocketClient) func(int, string) error {
	return func(code int, text string) error {
		log.Println("CLOSING")
		return manager.CloseConnection(client)
	}
}

package controller

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	pwd, _        = os.Getwd()
	JSON          = websocket.JSON              // codec for JSON
	Message       = websocket.Message           // codec for string, []byte
	ActiveClients = make(map[ClientConn]string) // map containing clients
	User          = make(map[string]string)
)
// Initialize handlers and websocket handlers
func init() {
	User["aaa"] = "aaa"
	User["bbb"] = "bbb"
	User["test"] = "test"
	User["test2"] = "test2"
	User["test3"] = "test3"
}

// Client connection consists of the websocket and the client ip
type ClientConn struct {
	websocket *websocket.Conn
	clientIP  string
}
func Hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		var err error
		defer ws.Close()

		client := ws.Request().RemoteAddr
		log.Println("Client connected:", client)
		sockCli := ClientConn{ws, client}
		ActiveClients[sockCli] = ""
		log.Println("Number of clients connected:", len(ActiveClients))

		for {
			// Write
			timestr:= time.Now().Format("2006-01-02 15:04:05")
			if err = Message.Send(ws,timestr+"[]"+ client); err != nil {
				log.Println("Could not send message to ",client , err.Error())
			}
			//for cs, na := range ActiveClients {
			//	if na != "" {
			//		timestr:= time.Now().Format("2006-01-02 15:04:05")
			//		if err = Message.Send(cs.websocket,timestr+"[]"+ cs.clientIP); err != nil {
			//			log.Println("Could not send message to ", cs.clientIP, err.Error())
			//		}
			//	}
			//}
			//rmsg := "Hello," + sockCli.clientIP
			//err := websocket.Message.Send(ws, rmsg)

			//if err != nil {
			//	fmt.Println("senderror")
			//	log.Fatal(err)
			//}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				// If we cannot Read then the connection is closed
				log.Println("Websocket Disconnected waiting", err.Error())
				// remove the ws client conn from our active clients
				delete(ActiveClients, sockCli)
				log.Println("Number of clients still connected:", len(ActiveClients))
				return
			}
			ActiveClients[sockCli] ="a"
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

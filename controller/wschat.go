package controller

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"os"
	"time"
	"net/http"
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
func SelectPersion(c echo.Context) error{
	//接收参数，题目id
	//随机生成候选人
	//调用ws发送大屏幕候选人
	name :=c.QueryParam("name")
	fmt.Println(name)
	for cs, na := range ActiveClients {
		if na != "" {
			timestr:= time.Now().Format("2006-01-02 15:04:05")
			if err := Message.Send(cs.websocket,timestr+"[]"+ cs.clientIP+name); err != nil {
				log.Println("Could not send message to ", cs.clientIP, err.Error())
			}
		}
	}
	//返回手机遥控器
	return c.JSON(http.StatusOK,nil)
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
		ActiveClients[sockCli] = "a"
		log.Println("Number of clients connected:", len(ActiveClients))

		for {
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

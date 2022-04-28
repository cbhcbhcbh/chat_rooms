package server

import (
	"chat_rooms/pkg/common/constant"
	"chat_rooms/pkg/global/log"
	"chat_rooms/pkg/protocol"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Unregister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Logger.Error("client read message error", log.Any("client read message error", err.Error()))
			MyServer.Unregister <- c
			c.Conn.Close()
			break
		}

		//解析传回来的message
		msg := &protocol.Message{}
		proto.Unmarshal(message, msg)

		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongByte, err := proto.Marshal(pong)
			if err != nil {
				log.Logger.Error("client marshaling message error", log.Any("client marshaling message error", err.Error()))
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			//TODO: 将数据从kafka消息队列中读取出来
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}

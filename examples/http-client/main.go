//go:build examples
// +build examples

package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type MCPMessage struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/"}
	log.Printf("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			var msg MCPMessage
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println("read:", err)
				return
			}

			jsonBytes, _ := json.MarshalIndent(msg, "", "  ")
			log.Printf("Received: %s", string(jsonBytes))
		}
	}()

	// Test sequence
	testMessages := []MCPMessage{
		{
			JSONRPC: "2.0",
			ID:      1,
			Method:  "initialize",
			Params: map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"capabilities":    map[string]interface{}{},
				"clientInfo": map[string]interface{}{
					"name":    "test-client",
					"version": "1.0.0",
				},
			},
		},
		{
			JSONRPC: "2.0",
			ID:      2,
			Method:  "tools/list",
		},
		{
			JSONRPC: "2.0",
			ID:      3,
			Method:  "tools/call",
			Params: map[string]interface{}{
				"name": "web_search",
				"arguments": map[string]interface{}{
					"query":       "Go programming tutorial",
					"max_results": 3,
				},
			},
		},
		{
			JSONRPC: "2.0",
			ID:      4,
			Method:  "ping",
		},
	}

	// Send test messages with delays
	for i, msg := range testMessages {
		log.Printf("Sending message %d:", i+1)
		jsonBytes, _ := json.MarshalIndent(msg, "", "  ")
		log.Printf("%s", string(jsonBytes))

		err := c.WriteJSON(msg)
		if err != nil {
			log.Println("write:", err)
			return
		}

		// Wait a bit between messages to see responses
		if i < len(testMessages)-1 {
			select {
			case <-done:
				return
			case <-interrupt:
				log.Println("Interrupt received, closing connection...")
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}
				return
			default:
				// Continue to next message
			}
		}
	}

	// Wait for interrupt
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Interrupt received, closing connection...")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

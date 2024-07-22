package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r3labs/sse/v2"
)

func main() {
	e := echo.New()

	e.GET("/events", func(c echo.Context) error {
		// Set the headers for SSE
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")

		// Create an SSE client
		// From simple implementation. No channel parameter
		client := sse.NewClient("http://localhost:8080/sse")
		// From broadcast implementation. Required with channel parameter
		//client := sse.NewClient("http://localhost:8080/sse?stream=time")

		// Channel to receive messages
		messages := make(chan *sse.Event)

		// Subscribe to the upstream SSE server
		err := client.SubscribeChanRawWithContext(c.Request().Context(), messages)
		if err != nil {
			log.Printf("Failed to subscribe to messages: %v", err)
			return err
		}

		// Get the response writer
		flusher, ok := c.Response().Writer.(http.Flusher)
		if !ok {
			return c.String(http.StatusInternalServerError, "Streaming not supported")
		}

		// Forward messages from the upstream server to the client
		for msg := range messages {
			// Format the message according to SSE protocol (see the `README.md` for more details)
			message := fmt.Sprintf("data: %s\n\n", msg.Data)
			_, err := c.Response().Writer.Write([]byte(message))
			if err != nil {
				return err
			}
			flusher.Flush()
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":8181"))
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// HandlerFunc is the signature for our "pure" controllers.
// They return the data, the HTTP status code, and an error (if any).
type HandlerFunc func(c *gin.Context) (any, int, error)

// WSHandlerFunc is the signature for WebSocket controllers.
// They receive the context (for params) and the active connection.
type WSHandlerFunc func(c *gin.Context, conn *websocket.Conn) error

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins for development
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WrapHandler adapts our pure HandlerFunc to a gin.HandlerFunc.
// It handles the "dirty work" of setting the context for the middleware.
func WrapHandler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, status, err := h(c)

		if err != nil {
			// If there is an error, we return it directly
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}

		// If success, we set the result in the context.
		// The TimingMiddleware will pick this up and wrap it.
		c.Set("result", data)
		c.Status(status)
	}
}

// WrapWSHandler adapts a WSHandlerFunc to a gin.HandlerFunc.
// It handles the HTTP->WebSocket upgrade automatically.
func WrapWSHandler(h WSHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		println("WS Connection attempt from: " + c.Request.RemoteAddr)
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			println("WS Upgrade failed: " + err.Error())
			return
		}
		defer conn.Close()

		// Execute the WebSocket logic
		if err := h(c, conn); err != nil {
			// Log the error or handle disconnect
			// We can't send an HTTP response here because the connection is hijacked
		}
	}
}

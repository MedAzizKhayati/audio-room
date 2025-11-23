package handler

import (
	"app/internal/service"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type FibonacciHandler struct {
	service *service.FibonacciService
}

func NewFibonacciHandler(s *service.FibonacciService) *FibonacciHandler {
	return &FibonacciHandler{service: s}
}

// RegisterRoutes mimics the "Module" setup in NestJS.
// It maps the controller methods to specific routes.
func (h *FibonacciHandler) RegisterRoutes(router *gin.Engine) {
	router.BasePath()
	createPath := GetControllerPathCreator("/fibonacci")
	router.GET(createPath("/sequential"), WrapHandler(h.handleSequential))
	router.GET(createPath("/concurrent"), WrapHandler(h.handleConcurrent))
	router.GET(createPath("/ws"), WrapWSHandler(h.handleWebSocket))
}

// handleWebSocket is a WebSocket controller.
// It reads a number from the client and streams back the result.
func (h *FibonacciHandler) handleWebSocket(c *gin.Context, conn *websocket.Conn) error {
	// Use a mutex to serialize writes because gorilla/websocket
	// does not allow concurrent writes on the same connection.
	var writeMu sync.Mutex

	for {
		// Read message from client (expects a number string like "10")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err // Client disconnected or read error
		}

		// Parse input
		nStr := string(msg)
		n, err := strconv.Atoi(nStr)
		if err != nil {
			writeMu.Lock()
			_ = conn.WriteJSON(gin.H{"error": "Invalid number"})
			writeMu.Unlock()
			continue
		}

		// Basic validation and safety limit (avoid huge / slow calculations)
		if n < 0 || n > 92 {
			writeMu.Lock()
			_ = conn.WriteJSON(gin.H{"error": "n out of allowed range (0..92)"})
			writeMu.Unlock()
			continue
		}

		// Spawn a goroutine to process this message concurrently. It will
		// compute values and write them back, serializing writes via writeMu.
		go func(n int) {
			for i := 0; i <= n; i++ {
				// Calculate
				start := time.Now()
				val := h.service.Calculate(i)
				duration := time.Since(start)

				// Serialize writes
				writeMu.Lock()
				err := conn.WriteJSON(gin.H{
					"n":        i,
					"value":    val,
					"duration": duration.String(),
				})
				writeMu.Unlock()

				if err != nil {
					// If write fails (client disconnected), there's nothing
					// sensible to do here other than return and let other
					// goroutines eventually fail too.
					return
				}
			}
		}(n)
	}
}

// handleSequential is now a pure "Controller" method.
func (h *FibonacciHandler) handleSequential(c *gin.Context) (any, int, error) {
	nStr := c.Query("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var val uint64
	for i := range n {
		val = h.service.Calculate(i)
	}

	return gin.H{
		"n":     n,
		"value": val,
	}, http.StatusOK, nil
}

// handleConcurrent is now a pure "Controller" method.
func (h *FibonacciHandler) handleConcurrent(c *gin.Context) (any, int, error) {
	nStr := c.Query("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	results := h.service.CalculateConcurrent(n)
	var val uint64
	for _, result := range results {
		if result.Index == n-1 {
			val = result.Value
			break
		}
	}

	return gin.H{
		"n":     n,
		"value": val,
	}, http.StatusOK, nil
}

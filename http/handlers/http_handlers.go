package handlers

import (
	"go-challenge/trigram"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"log"

	"github.com/go-chi/render"
)

type HttpHandler struct {
	mutex   *sync.RWMutex
	trigram trigram.TrigramIntf
}

type HandlerOptions struct {
	Mutex   *sync.RWMutex
	State   map[string]interface{}
	Trigram trigram.TrigramIntf
}

func NewHTTPHandler(opt *HandlerOptions) *HttpHandler {
	return &HttpHandler{
		mutex:   &sync.RWMutex{},
		trigram: opt.Trigram,
	}
}

// HandleTeachRequest to teach trigram
func (h *HttpHandler) HandleTeachRequest(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Failed to decode request body."
		log.Print("decode request body", err)

		// Return error response
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": msg,
			"error":   err,
		})
		return
	}

	h.mutex.Lock()
	success := h.trigram.Add(string(buf))
	h.mutex.Unlock()
	if !success {
		// Return error response
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": "Invalid Body",
			"error":   success,
		})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "Trigram collected Successfully",
		"success": success,
	})
}

// HandleFetchRequest to get sentence from trigram
func (h *HttpHandler) HandleFetchRequest(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	h.mutex.RLock()

	query := q.Get("complexity")

	if q.Get("complexity") == "" {
		query = "5"
	}

	complexity, err := strconv.Atoi(query)
	if err != nil {
		msg := "Failed to covert request query."
		log.Print("Failed to covert request query", err)

		// Return error response
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": msg,
			"error":   err,
		})
		return
	}
	data := h.trigram.Query(complexity)
	h.mutex.RUnlock()
	if data == "" {
		// Return error response
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": "Empty data set",
			"success": false,
		})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": data,
		"success": true,
	})
}

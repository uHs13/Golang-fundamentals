package port

import "net/http"

type HandlerInterface interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

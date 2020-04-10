package handlers

import (
	"fmt"
	"net/http"
)

func LivenessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	fmt.Fprint(w)
}

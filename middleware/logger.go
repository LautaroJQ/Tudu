package middleware

import (
	"log"
	"net/http"
	"time"
)

// Las funciones handlefunc de los routers necesitan un handler por lo que logger envuelve el HANDLER de la ruta y devuelve una función del tipo handler que aplica cierta logica antes y después de que se haya ejecutado el handler de la ruta
func Logger(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.Printf("%s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}

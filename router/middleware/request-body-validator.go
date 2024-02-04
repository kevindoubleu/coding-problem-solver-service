package middleware

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestBodyValidatorMiddleware[T any] struct {
	logRequestBody bool
	parsedRequestBody T
}

func NewRequestBodyValidatorMiddleware[T any](
		logRequestBody bool,
	) RequestBodyValidatorMiddleware[T] {

	return RequestBodyValidatorMiddleware[T]{
		logRequestBody: logRequestBody,
	}
}

func (mw RequestBodyValidatorMiddleware[T]) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading request body, err: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if mw.logRequestBody {
			log.Printf("received request body: %s", body)
		}
		
		err = json.Unmarshal(body, &mw.parsedRequestBody)
		if err != nil {
			log.Printf("request body is not of expected type, err: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		enrichedContext := context.WithValue(r.Context(), "data", mw.parsedRequestBody)
		r = r.WithContext(enrichedContext)
		next.ServeHTTP(w, r)
	})
}

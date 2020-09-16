package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai4/komazz/omikuji"
)

var errJSONEncode = errors.New("Fail to Encode")

// Server サーバー型
type Server struct {
	Omikuji *omikuji.Omikuji
}

// Run サーバーを起動する
func Run() error {
	server := &Server{omikuji.New(time.Now())}
	if err := http.ListenAndServe(":8080", server); err != nil {
		return err
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/omikuji/", http.HandlerFunc(s.omikujiHandler))
	router.ServeHTTP(w, r)
}

func (s *Server) omikujiHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(s.Omikuji.Run()); err != nil {
		log.Println(errJSONEncode)
	}
}

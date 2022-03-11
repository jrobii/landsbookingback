package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/jrobii/landsbookingback/config"
)

func main() {

	mux := http.NewServeMux()

	configuration := config.NewConfiguration()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	err := http.ListenAndServe(configuration.Port, mux)
	if err != nil {
		log.Error().Err(err).Msg("starting http server")
	}
	log.Info().Str("port", configuration.Port).Msg("listening on")
}

package main

import (
	"fmt"
	"github.com/alliander/diva-go-backend/api"
	"github.com/alliander/diva-go-backend/config"
	"github.com/alliander/diva-go-backend/session"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/privacybydesign/irmago/server"
	"github.com/privacybydesign/irmago/server/irmaserver"
	"log"
	"net/http"
	"strconv"
)

func main() {
	c := config.GetConfig()
	fmt.Printf("Running with config: %v\n", c)

	session.Init(c)

	err := irmaserver.Initialize(&server.Configuration{
		// Replace with address that IRMA apps can reach
		URL:        c.IrmaURL,
		EnableSSE:  true,
		Production: c.IrmaProductionMode,
	})

	if err != nil {
		panic(fmt.Sprintf("Error Initializing Irma server: %v", err))
	}

	var corsOptions = cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "Cache-Control"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}

	router := chi.NewRouter()

	// TODO: Enable CORS only on /api/irma/* endpoints
	router.Use(cors.New(corsOptions).Handler)

	router.Handle("/api/irma/*", irmaserver.HandlerFunc())
	router.HandleFunc("/api/start-irma-session", api.StartIrmaSession)
	router.HandleFunc("/api/complete-irma-session", api.CompleteIrmaSession)
	router.HandleFunc("/api/get-session", api.GetSession)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Port), router))
}

package main

import (
	"fmt"
	"net/http"

	P_ServicesNasabah "api-nasabah/Services"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {
	goji.Get("/services/nasabah", P_ServicesNasabah.ServiceGetNasabahAll)
	goji.Get("/services/nasabah/:id", P_ServicesNasabah.ServiceGetNasabahById)
	goji.Post("/services/nasabah/:data", P_ServicesNasabah.ServiceAddNasabah)
	goji.Put("/services/nasabah/:data", P_ServicesNasabah.ServiceUpdateNasabahById)
	goji.Delete("/services/nasabah/:id", P_ServicesNasabah.ServiceDeleteNasabahById)
	goji.Post("/auth/nasabah/:data", P_ServicesNasabah.AuthNasabah)
	goji.Serve()
}

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
	goji.Get("/nasabah", P_ServicesNasabah.ServiceGetNasabahAll)
	goji.Get("/nasabah/:id", P_ServicesNasabah.ServiceGetNasabahById)
	goji.Post("/nasabah/:data", P_ServicesNasabah.ServiceAddNasabah)
	goji.Put("/nasabah/:data", P_ServicesNasabah.ServiceUpdateNasabahById)
	goji.Delete("/nasabah/:id", P_ServicesNasabah.ServiceDeleteNasabahById)
	goji.Serve()
}

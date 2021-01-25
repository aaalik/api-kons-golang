package routers

import (
	"net/http"
	"os"

	v1 "github.com/aaalik/ke-jepang/controller/v1"
	"github.com/aaalik/ke-jepang/helper"
)

func SetupRouter() {
	r := http.NewServeMux()

	//Redirect HTTP request
	go http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), http.HandlerFunc(redirect))

	r.HandleFunc("/v1/items/", v1.List)

	helper.Log.Info("Server started at " + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusPermanentRedirect)
}

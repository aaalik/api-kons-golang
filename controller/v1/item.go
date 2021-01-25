package v1

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/aaalik/ke-jepang/helper"
	"github.com/aaalik/ke-jepang/model"
)

// func saveItem(w http.ResponseWriter, r *http.Request) {
// 	r.FormValue("name")
// 	list = model.SaveItem()
// }

func List(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "GET":
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"message": "get called"}`))
	// case "POST":
	// 	saveItem(w, r)
	// case "PUT":
	// 	w.WriteHeader(http.StatusAccepted)
	// 	w.Write([]byte(`{"message": "put called"}`))
	// case "DELETE":
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"message": "delete called"}`))
	// default:
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte(`{"message": "not found"}`))
	// }

	tmp_id := strings.TrimPrefix(r.URL.Path, "/v1/items/")

	var list interface{}
	var err error

	if tmp_id != "" {
		id, _ := strconv.Atoi(tmp_id)
		list, err = model.GetSingleItem(id)
		if err != nil {
			helper.Response(w, err, http.StatusBadRequest, true)
			return
		}

	} else {
		list = model.GetItems()
	}

	helper.Response(w, list, http.StatusOK, false)
}

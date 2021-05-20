package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type webResponseAdd struct {
	Number1 int `json:"n1"` // input number1
	Number2 int `json:"n2"` // input number2
	Result  int `json:"r"`  // result of number1 + number2
}

func main() {
	http.HandleFunc("/add/", add)
	http.ListenAndServe(":8080", nil)
}

func add(w http.ResponseWriter, r *http.Request) {
	var n1_int int
	var n2_int int
	var err error
	r.ParseForm()
	
	if len(r.Form) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    for k := range r.Form {
        if len(r.Form[k]) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if n1, ok := r.Form["n1"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		n1_int, err = strconv.Atoi(n1[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
		
	if n2, ok := r.Form["n2"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		n2_int, err = strconv.Atoi(n2[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	responseJSON := webResponseAdd{Number1: n1_int, Number2: n2_int, 
		Result: n1_int + n2_int}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(responseJSON)
}
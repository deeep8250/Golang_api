package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {

	person := []map[string]interface{}{
		{"id": 1, "name": "Deep", "age": 18},
		{"id": 2, "name": "Asit", "age": 60},
		{"id": 3, "name": "Hena", "age": 80},
		{"id": 4, "name": "Kusom", "age": 18},
		{"id": 5, "name": "tommy", "age": 2},
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(person)

	})

	http.HandleFunc("/users/filter", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var secList []map[string]interface{}

		age := r.URL.Query().Get("age")

		age1, err := strconv.Atoi(age)
		if err != nil || age == "" {

			panic("error")
		}

		for _, i := range person {
			if i["age"] == age1 {

				secList = append(secList, i)
			}

		}

		json.NewEncoder(w).Encode(secList)

	})

	http.ListenAndServe(":8888", nil)

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	var personDetails []map[string]interface{}

	// get data ( get )
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Printf("%+v\n", personDetails)
		json.NewEncoder(w).Encode(personDetails)
	})

	// add data( post )
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var addPerson map[string]interface{}

		if r.Method != http.MethodPost {
			http.Error(w, "check if you select the post method or not", http.StatusBadRequest)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&addPerson)

		if err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		personDetails = append(personDetails, addPerson)
		json.NewEncoder(w).Encode(personDetails)

	})

	//filter result by age

	http.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var personDetails2 []map[string]interface{}

		if r.Method != http.MethodGet {
			http.Error(w, "double check if you selected post option or not ", http.StatusBadRequest)
			return
		}

		age := r.URL.Query().Get("age")

		for i := range personDetails {
			if personDetails[i]["age"] == age {
				personDetails2 = append(personDetails2, personDetails[i])

			}
		}

		json.NewEncoder(w).Encode(personDetails2)
		defer r.Body.Close()

	})

	fmt.Printf("%+v\n", personDetails)
	http.ListenAndServe(":8888", nil)

	//v : 0.0.1.1
}

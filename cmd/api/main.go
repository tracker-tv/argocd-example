package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := []Todo{
			{1, "Test argocd deployment", true},
			{2, "Create a new project", false},
			{3, "Create a new project", false},
			{4, "Create a new project", false},
		}
		js, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(js)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		return
	}
}

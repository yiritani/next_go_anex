package req_ping

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	fmt.Print(response) // Print the response to the console

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

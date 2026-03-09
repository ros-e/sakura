package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ros-e/sakura/packages/api/internal/config"
	"github.com/ros-e/sakura/packages/api/internal/json"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v", err)
		return
	}

	if err = os.MkdirAll(homeDir+"/.sakura", 0755); err != nil {
		fmt.Printf("Error creating config directory: %v", err)
		return
	}

	if _, err = os.Stat(homeDir + "/.sakura/config.yaml"); os.IsNotExist(err) {
		if err = config.Init(homeDir + "/.sakura/config.yaml"); err != nil {
			fmt.Printf("Error creating config: %v", err)
			return
		}
	}

	// not found
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.Write(w, http.StatusNotFound, []byte(`
			{"error": "true",
			"message": "not found"}`,
		))
	})

	http.HandleFunc("/setup", func(w http.ResponseWriter, r *http.Request) {
		// this will depend on if the server is setup...

	})

	fmt.Println("Listening on http://localhost:8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v", err)
	}
}

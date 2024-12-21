package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sysstats-cli/src/config"
	"sysstats-cli/src/utils"
	"time"
)

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := utils.GetStats()
	if err != nil {
		http.Error(w, "Error fetching stats", http.StatusInternalServerError)
		log.Println("Error fetching stats:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func Execute() {
	// Run HTTP server in background
	http.HandleFunc("/stats", statsHandler)

	port := config.GetPort()
	go func() {
		fmt.Println("Starting server on http://localhost" + port)
		log.Fatal(http.ListenAndServe(port, nil))
	}()

	// Simulate background running
	for {
		time.Sleep(10 * time.Second)
	}

}

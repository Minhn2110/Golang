package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		// logger.WithError(err).Error("failed to read request body")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	fmt.Printf("%s", body)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeaders(map[string]string{
			"K-Trigger-Id":    "2Edk0ZEWS8c",
			"K-Shared-Secret": "123456",
		}).
		SetBody(body).
		Post("https://dev-api.kincloud.io/v1/webhook/kin-cloud-module-common/v1-0/2Edk0ZESXVD")

	fmt.Printf("%s", resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start workflow"})
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Executed workflow success"})

}

func main() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	a := viper.AllSettings()
	b, _ := json.Marshal(a)

	fmt.Printf(string(b))

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

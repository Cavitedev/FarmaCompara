package web_scrap

import (
	"context"
	"encoding/json"

	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cavitedev/farma-compara/web_scrap/scraper"
)

var client *firestore.Client
var ctx context.Context

func init() {

	ctx = context.Background()
	conf := &firebase.Config{ProjectID: "farma-compara"}

	var err error

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Firestore: %v", err)
	}

	functions.HTTP("scrapWebsite", scrapWebsite)
}

// scrapWebsite is an HTTP Cloud Function with a request parameter.
func scrapWebsite(w http.ResponseWriter, r *http.Request) {

	var d struct {
		Website       string `json:"website"`
		ScrapItems    bool
		ScrapDelivery bool
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Printf("Could not decode json body with the website name")
		return
	}
	if d.Website == "" {
		http.Error(w, "Missing 'website' parameter", http.StatusBadRequest)
		return
	}
	if d.ScrapItems == false && d.ScrapDelivery == false {
		http.Error(w, "Missing 'scrapItems' or 'scrapDelivery' parameter", http.StatusBadRequest)
		return
	}

	result := scraper.Scrap(d.Website, client, d.ScrapItems, d.ScrapDelivery)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})

}

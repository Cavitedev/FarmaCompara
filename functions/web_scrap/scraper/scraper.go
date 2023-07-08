package scraper

import (
	"log"

	"cloud.google.com/go/firestore"
	"github.com/cavitedev/farma-compara/web_scrap/scraper/farmaciaencasa"

	"github.com/cavitedev/farma-compara/web_scrap/scraper/okfarma"
)

func Scrap(website string, client *firestore.Client, scrapItems bool, scrapDelivery bool) string {

	log.Println("Inicializando scraper")

	switch website {
	case okfarma.Domain:
		okfarma.Scrap(client, scrapItems, scrapDelivery)
	case farmaciaencasa.Domain:
		farmaciaencasa.Scrap(client, scrapItems, scrapDelivery)
	case "all":
		okfarma.Scrap(client, scrapItems, scrapDelivery)
		farmaciaencasa.Scrap(client, scrapItems, scrapDelivery)
	default:
		log.Fatalf("No se ha encontrado la página \"%v\" para scrappear los datos\n", website)
	}

	return "Scrapping of " + website + " complete"

}

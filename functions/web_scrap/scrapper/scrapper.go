package scrapper

import (
	"log"

	"cloud.google.com/go/firestore"
	"github.com/cavitedev/farma-compara/web_scrap/scrapper/dosfarma"
	"github.com/cavitedev/farma-compara/web_scrap/scrapper/farmaciaencasa"
	"github.com/cavitedev/farma-compara/web_scrap/scrapper/farmaciasdirect"
	"github.com/cavitedev/farma-compara/web_scrap/scrapper/okfarma"
)

func Scrap(website string, ref *firestore.CollectionRef) string {

	log.Println("Hola scrapper")

	switch website {
	case okfarma.Domain:
		okfarma.Scrap(ref)
	case farmaciasdirect.Domain:
		farmaciasdirect.Scrap(ref)
	case dosfarma.Domain:
		dosfarma.Scrap(ref)
	case farmaciaencasa.Domain:
		farmaciaencasa.Scrap(ref)
	}

	return "Scrapping of " + website + " complete"

}

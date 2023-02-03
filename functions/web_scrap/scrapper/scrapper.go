package scrapper

import (
	"log"

	"cloud.google.com/go/firestore"
	"github.com/cavitedev/farma-compara/web_scrap/scrapper/dosfarma"
	"github.com/cavitedev/farma-compara/web_scrap/scrapper/farmaciasdirect"
	"github.com/cavitedev/farma-compara/web_scrap/scrapper/okfarma"
)

func Scrap(website string, ref *firestore.CollectionRef) string {

	log.Println("Hola scrapper")

	if website == okfarma.Domain {
		okfarma.Scrap(ref)
	} else if website == farmaciasdirect.Domain {
		farmaciasdirect.Scrap(ref)
	} else if website == dosfarma.Domain {
		dosfarma.Scrap(ref)
	}

	return "Scrapping of " + website + " complete"

}

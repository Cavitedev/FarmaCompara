package scrapper

import (
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Cavitedev/terraform_tuto/web_scrap/scrapper/farmaciasdirect"
	"github.com/Cavitedev/terraform_tuto/web_scrap/scrapper/okfarma"
)

func Scrap(website string, ref *firestore.CollectionRef) string {

	log.Println("Hola scrapper")

	if website == okfarma.Domain {
		okfarma.Scrap(ref)
	} else if website == farmaciasdirect.Domain {
		farmaciasdirect.Scrap(ref)
	} else {
		return website + " not found"
	}

	return "Scrapping of " + website + " complete"

}

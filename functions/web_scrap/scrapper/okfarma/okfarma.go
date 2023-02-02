package okfarma

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/Cavitedev/terraform_tuto/web_scrap/firestore_utils"
	. "github.com/Cavitedev/terraform_tuto/web_scrap/scrapper/types"
	"github.com/Cavitedev/terraform_tuto/web_scrap/utils"
	"github.com/gocolly/colly/v2"
)

const Domain string = "okfarma.es"

func Scrap(ref *firestore.CollectionRef) {

	log.Println(Domain)

	items := []Item{}
	c := colly.NewCollector(
		// colly.Async(true),
		colly.AllowedDomains(Domain),
	)

	c.OnHTML("#product_list", func(h *colly.HTMLElement) {
		log.Println("Product List")

		h.ForEach(".product-container", func(_ int, e *colly.HTMLElement) {
			item := Item{}
			pageItem := WebsiteItem{}
			pageItem.Url = e.ChildAttr(".product-image-container a", "href")
			scrapDetailsPage(&item, &pageItem)
			if item.WebsiteItems == nil {
				item.WebsiteItems = make(map[string]WebsiteItem)
			}
			item.WebsiteItems[Domain] = pageItem
			items = append(items, item)
			firestore_utils.UpdateItem(item, ref)
			time.Sleep(50 * time.Millisecond)
		})
	})

	url := buildPageUrl()
	err := c.Visit(url)
	if err != nil {
		log.Printf("Error when visiting %v, err:%v", url, err)
	}

	bytes, _ := json.Marshal(items)
	log.Printf("%+v\n", string(bytes))

}

var productsVisited int = 0

func scrapDetailsPage(item *Item, pageItem *WebsiteItem) {
	c := colly.NewCollector(
		colly.AllowedDomains(Domain),
	)
	c.OnResponse(func(r *colly.Response) {
		productsVisited++
		log.Printf("Visit %d URL:%v\n", productsVisited, r.Request.URL)

	})

	c.OnHTML("div #center_column", func(h *colly.HTMLElement) {
		currentTime := time.Now()
		pageItem.LastUpdate = currentTime
		pageItem.Image = h.ChildAttr("#bigpic", "src")
		pageItem.Name = h.ChildText("h1.product-name")

		price := h.ChildText("#our_price_display")
		pageItem.Price = utils.ParseSpanishNumberStrToNumber(price)
		pageItem.Available = h.ChildText("#availability_value span") != "Este producto ya no está disponible"
		item.Ref = h.ChildAttr("#product_reference>span", "content")
	})

	c.Visit(pageItem.Url)
}

func buildPageUrl() string {

	url := fmt.Sprintf("https://%v/medicamentos?id_category=3&n=1192", Domain)
	return url
}

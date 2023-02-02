package firestore_utils

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Cavitedev/terraform_tuto/web_scrap/firestore_utils/transform"
	"github.com/Cavitedev/terraform_tuto/web_scrap/scrapper/types"
)

func UpdateItem(item types.Item, col *firestore.CollectionRef) {
	ctx := context.Background()
	id := item.Ref
	doc := col.Doc(id)

	m := transform.ToFirestoreMap(item)

	_, err := doc.Set(ctx, m, firestore.MergeAll)

	if err != nil {
		log.Printf("Could not insert %v\n", item)
	}
}

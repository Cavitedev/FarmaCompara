import {
  onDocumentWritten,
} from "firebase-functions/v2/firestore";
import {setGlobalOptions} from "firebase-functions/v2/options";

setGlobalOptions({region: "europe-west2"});

export const onItemUpdate = onDocumentWritten("items/{itemId}",
  async (event) => {
    const after: item = event.data?.after.data() as item;

    let name;
    let bestPrice;
    let lastUpdate;

    for (const value of Object.values(after.website_items)) {
      if (lastUpdate === undefined || value.last_update > lastUpdate) {
        lastUpdate = value.last_update;
      }

      if (value.available) {
        name = value.name;


        if (bestPrice === undefined || value.price < bestPrice) {
          bestPrice = value.price;
        }
      }
    }

    event.data?.after.ref.set(
      {name: name, best_price: bestPrice, last_update: lastUpdate},
      {merge: true}
    );
  }
);


interface websiteItem {
  available: true;
  image: string;
  last_update: Date;
  name: string;
  price: number;
  url: string;
}

interface item {
  ref: string;
  website_items: { [key: string]: websiteItem };
}

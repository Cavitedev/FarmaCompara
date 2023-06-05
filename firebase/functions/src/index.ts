import * as functions from "firebase-functions";

export const onItemUpdate = functions.region("europe-west2").firestore
  .document("items/{itemId}")
  .onWrite(async (event) => {
    const after: item = event.after.data() as item;

    let name;
    let bestPrice;
    let lastUpdate;

    for (const value of Object.values(after.website_items)) {
      if (value.available) {
        name = value.name;

        if (lastUpdate === undefined || value.last_update > lastUpdate) {
          lastUpdate = value.last_update;
        }
        if (bestPrice === undefined || value.price < bestPrice) {
          bestPrice = value.price;
        }
      }
    }

    event.after.ref.set(
      {name: name, best_price: bestPrice, last_update: lastUpdate},
      {merge: true}
    );
  });
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

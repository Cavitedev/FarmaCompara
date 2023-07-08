import {onDocumentWritten} from "firebase-functions/v2/firestore";
import {setGlobalOptions} from "firebase-functions/v2/options";
import {initializeApp} from "firebase-admin/app";
import {getFirestore} from "firebase-admin/firestore";

setGlobalOptions({region: "europe-west2"});

initializeApp();
const db = getFirestore();

db.settings({ignoreUndefinedProperties: true});

export const onItemUpdate = onDocumentWritten(
  "items/{itemId}",
  async (event) => {
    const after: item = event.data?.after.data() as item;

    let name;
    let bestPrice;
    let lastUpdate;
    const websiteNames = Object.keys(after.website_items);

    const websiteIterable = Object.values(after.website_items);

    const websitesCount = websiteIterable.length;


    for (const value of websiteIterable) {
      if (lastUpdate === undefined || value.last_update > lastUpdate) {
        lastUpdate = value.last_update;
      }
      name = value.name;

      if (value.available) {
        if (bestPrice === undefined || value.price < bestPrice) {
          bestPrice = value.price;
        }
      }
    }

    event.data?.after.ref.set(
      {
        name: name,
        best_price: bestPrice,
        last_update: lastUpdate,
        websites_count: websitesCount,
        website_names: websiteNames,
      },
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
  name: string;
  best_price: number;
  last_update: Date;
  website_items: Map<string, websiteItem>;
}

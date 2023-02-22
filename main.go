package main

import (
	fmt "fmt"

	itemv2 "github.com/donwellus/go-protobuf-oneof/gen/item/v2"
	"google.golang.org/protobuf/proto"
)

func main() {
	product := &itemv2.Item_ProductItem{
		ProductItem: &itemv2.Item_Product{
			Sku: string("HT-1234"),
		},
	};
	showcase := &itemv2.Item_ShowcaseItem{
		ShowcaseItem: &itemv2.Item_Showcase{
			Pos: "br",
			Category: "home",
			Slug: "home",
		},
	};
	collection := &itemv2.Collection{
		Items: []*itemv2.Item{
			{Item: product},
			{Item: showcase},
		},
	}

	data, err := proto.Marshal(collection)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	newCollection := &itemv2.Collection{}
	err = proto.Unmarshal(data, newCollection)
	if err != nil {
		panic(err)
	}

	for i, it := range newCollection.Items {
		switch item := it.GetItem().(type) {
		case *itemv2.Item_ProductItem:
			fmt.Println(i, item.ProductItem)
		case *itemv2.Item_ShowcaseItem:
			fmt.Println(i, item.ShowcaseItem)
		}
	}
}

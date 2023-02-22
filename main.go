package main

import (
	fmt "fmt"

	itemv2 "github.com/donwellus/go-protobuf-oneof/gen/item/v2"
	"google.golang.org/protobuf/proto"
)

func main() {
	product := &itemv2.Item_ProductItem{
		ProductItem: &itemv2.Item_Product{
			Sku: "ht-1234",
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

	fmt.Println(newCollection)
	fmt.Printf("%T\n", newCollection)
	fmt.Printf("%+v\n", newCollection)

	for i, it := range newCollection.Items {
		fmt.Printf("RAW\tIndex: %d\tType: %T\t\tValue: {%v}\n", i, it, it)
		switch item := it.GetItem().(type) {
		case *itemv2.Item_ProductItem:
			fmt.Printf("TYPED\tIndex: %d\tType: %T\tValue: {%v}\n", i, item.ProductItem, item.ProductItem)
		case *itemv2.Item_ShowcaseItem:
			fmt.Printf("TYPED\tIndex: %d\tType: %T\tValue: {%v}\n", i, item.ShowcaseItem, item.ShowcaseItem)
		}
	}
}

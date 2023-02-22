package main

import (
	fmt "fmt"

	itemv1 "github.com/donwellus/go-protobuf-oneof/gen/item/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	message := &itemv1.Message{
		Item: &itemv1.Message_ProductItem{
			ProductItem: &itemv1.ProductItem{
				Sku: string("HT-1234"),
			},
		},
	}

	// message := &itemv1.Message{
	// 	Item: &itemv1.Message_ShowcaseItem{
	// 		ShowcaseItem: &itemv1.ShowcaseItem{
	// 			Pos: "br",
	// 			Category: "home",
	// 			Slug: "hotels",
	// 		},
	// 	},
	// }

	data, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	newMessage := &itemv1.Message{}
	err = proto.Unmarshal(data, newMessage)
	if err != nil {
		panic(err)
	}

	switch item := message.Item.(type) {
	case *itemv1.Message_ProductItem:
		fmt.Println(item.ProductItem)
	case *itemv1.Message_ShowcaseItem:
		fmt.Println(item.ShowcaseItem)
	}
	// fmt.Println(newMessage.GetProductItem())
	// fmt.Println(newMessage.GetShowcaseItem())
}

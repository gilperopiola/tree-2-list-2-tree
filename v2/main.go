package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("tree-2-list-2-tree")

	tree := []TreeItem{
		NewTreeItem("Hamburger").With(NewTreeItem("Ketchup")),
		NewTreeItem("Soda"),
	}

	// tree := BuildTree([]TreeItem{})

	// tree := Tree{
	// 	{
	// 		Item: NewItem("Hamburger"),
	// 		Children: Tree{{
	// 			Item:     NewItem("Ketchup"),
	// 			Children: Tree{},
	// 		}},
	// 	},
	// 	{
	// 		Item:     NewItem("Soda"),
	// 		Children: Tree{},
	// 	},
	// }
}

func NewItem(name string) Item {
	return Item{
		ID:   uuid.New(),
		Name: name,
	}
}

func NewTreeItem(name string) TreeItem {
	return TreeItem{
		Item: NewItem(name),
	}
}

func NewListItem(name string) ListItem {
	return ListItem{
		Item: NewItem(name),
	}
}

func (tree Tree) ToList() List {
	var list List

	for _, rootItem := range tree {
		list = append(list, ListItem{
			Item: Item{
				ID:   rootItem.ID,
				Name: rootItem.Name,
			},
			ParentID: uuid.Nil,
		})

		for _, childItem := range rootItem.Children {
			list = append(list, ListItem{
				Item:     childItem.Item,
				ParentID: rootItem.ID,
			})
		}
	}

	return list
}

func (list List) ToTree() Tree {
	tree := Tree{}

	return tree
}

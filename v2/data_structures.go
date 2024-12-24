package main

import "github.com/google/uuid"

// Item 🍔 is anything that can be stored in a list or a tree.
// More a tree than a list, as our lists are hierarchical so they
// are basically trees.
type Item struct {
	ID   uuid.UUID
	Name string
}

// ListItem 📃 is an item that can be stored in a list,
// having to store the ID of its parent to be able to traverse
// the hierarchy.
type ListItem struct {
	Item
	ParentID uuid.UUID
}

// TreeItem 🛒 is an item that can be stored in a tree,
// having to store the IDs of its children to be able to traverse
// the hierarchy.
type TreeItem struct {
	Item
	Children []TreeItem
}

func (this TreeItem) With(children ...TreeItem) TreeItem {
	if this.Children == nil {
		this.Children = []TreeItem{}
	}
	this.Children = append(this.Children, children...)
	return this
}

// List 📃 is a list of items where each item can have a parent.
// The order first goes through all children of the first item
// before moving on to the next one.
//
// JSON example:
// [{
//   "id": "1",
//   "name": "Hamburger",
//   "parent_id": "0"
// }, {
//   "id": "2",
//   "name": "Ketchup",
//   "parent_id": "1" // child of the first item
// }, {
//   "id": "3",
//   "name": "Soda",
//   "parent_id": "0"
//}]
//
type List []ListItem

// Tree 🛒 is a tree of items.
// Empty 'Children' slices are marshalled as 'null' in JSON (nil in Go).
//
// JSON example:
// [{
//   "id": "1",
//   "name": "Hamburger",
//   "children": [
//     {
//       "id": "2",
//       "name": "Ketchup",
//       "children": []
//     }
//   ]
// }, {
//   "id": "3",
//   "name": "Soda",
//   "children": []
// }]
//
type Tree []TreeItem

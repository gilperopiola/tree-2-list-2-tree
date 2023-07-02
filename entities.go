package main

type List []ListItem
type Tree []TreeItem

type ListItem struct {
	ID           int
	Name         string
	ParentItemID int
}

func (i ListItem) ToTreeItemWithoutChildren() TreeItem {
	return TreeItem{
		ID:           i.ID,
		Name:         i.Name,
		ParentItemID: i.ParentItemID,
	}
}

type TreeItem struct {
	ID           int
	Name         string
	ParentItemID int
	Children     Tree
}

func (treeItem TreeItem) ToListItem() ListItem {
	return ListItem{
		ID:           treeItem.ID,
		Name:         treeItem.Name,
		ParentItemID: treeItem.ParentItemID,
	}
}

package main

func (i ListItem) ToTreeItemWithoutChildren() TreeItem {
	return TreeItem{
		ID:           i.ID,
		Name:         i.Name,
		ParentItemID: i.ParentItemID,
	}
}

func (i ListItem) IsRoot() bool {
	return i.ParentItemID == 0
}

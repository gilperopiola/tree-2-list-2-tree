package main

func (i ListItem) IsRoot() bool {
	return i.ParentItemID == 0
}

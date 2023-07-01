package main

type List []ListItem
type Tree []TreeItem

type ListItem struct {
	ID           int
	Name         string
	ParentItemID int
}

type TreeItem struct {
	ID           int
	Name         string
	ParentItemID int
	Children     Tree
}

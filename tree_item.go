package main

// AddChildrenToChildren is the starting point of the recursion from List to Tree
func (treeItem TreeItem) AddChildrenToChildren(modifiersList List) TreeItem {
	firstLevelModifiersIDs := treeItem.GetChildrenIDs()
	secondLevelSingleParentModifiers := modifiersList.Filter(firstLevelModifiersIDs)
	treeItem.Children = treeItem.Children.AddChildren(secondLevelSingleParentModifiers)

	for i, treeItemChild := range treeItem.Children {
		treeItem.Children[i] = treeItemChild.AddChildrenToChildren(modifiersList)
	}

	return treeItem
}

func (treeItem TreeItem) GetChildrenIDs() []int {
	childrenIDs := []int{}
	for _, treeItemChild := range treeItem.Children {
		childrenIDs = append(childrenIDs, treeItemChild.ID)
	}
	return childrenIDs
}

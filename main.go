package main

import "fmt"

func main() {
	fmt.Println("asd")
}

func (tree Tree) ToList() List {

}

func (list List) ToTree() Tree {

	// Append Root Items to the Tree
	// Also fill modifiersList and rootIDs
	outTree, rootIDs, modifiersList := list.prepareToConvert()

	// We add the first layer of modifiers
	firstLevelModifiers := modifiersList.Filter(rootIDs)
	outTree = outTree.AddChildren(firstLevelModifiers)

	// We traverse the root items and their children
	for i, treeItem := range outTree {
		outTree[i] = treeItem.AddChildrenToChildren(modifiersList)
	}

	return outTree
}

func (list List) prepareToConvert() (Tree, []int, List) {

	// outTree is the output
	outTree := Tree{}

	// modifiersList stores all the modifiers
	modifiersList := List{}

	// rootIDs holds the IDs of the root items
	rootIDs := []int{}

	for _, item := range list {
		if item.IsRoot() {
			outTree = append(outTree, item.ToTreeItemWithoutChildren())
			rootIDs = append(rootIDs, item.ID)
		} else {
			modifiersList = append(modifiersList, item)
		}
	}

	return outTree, rootIDs, modifiersList
}

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

func (tree Tree) AddChildren(children List) Tree {
	for i, item := range tree {
		for _, child := range children {
			if item.ID == child.ParentItemID {
				tree[i].Children = append(tree[i].Children, child.ToTreeItemWithoutChildren())
			}
		}
	}

	return tree
}

func (list List) Filter(parentIDs []int) List {
	out := List{}
	for _, item := range list {

		// Filter through parentIDs
		for _, parentID := range parentIDs {
			if item.ParentItemID == parentID {
				out = append(out, item)
			}
		}

	}
	return out
}

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

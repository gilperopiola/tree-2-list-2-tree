package main

import "fmt"

func main() {
	fmt.Println("tree-2-list-2-tree")
}

func (tree Tree) ToList() List {
	outList := List{}

	for _, treeItem := range tree {
		outList = append(outList, treeItem.ToListItem())
		outList = append(outList, treeItem.GetChildrenAsList()...)
	}

	return outList
}

func (list List) ToTree() Tree {

	// Get initial tree with the root items
	// Also get a list of all the modifiers, and the IDs of the root items
	outTree, modifiersList, rootIDs := list.prepareToConvertToTree()

	// We add the first layer of modifiers to the tree
	firstLevelModifiers := modifiersList.Filter(rootIDs)
	outTree = outTree.AddChildren(firstLevelModifiers)

	// We traverse the root items and their children
	for i, treeItem := range outTree {
		outTree[i] = treeItem.AddChildrenToChildren(modifiersList)
	}

	return outTree
}

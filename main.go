package main

import "fmt"

func main() {
	fmt.Println("asd")
}

func (tree Tree) ToList() List {
	return List{}
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

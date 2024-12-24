package main

// outTree contains the root items
// modifiersList stores all the modifiers
// rootIDs holds the IDs of the root items
func (list List) prepareToConvertToTree() (outTree Tree, modifiersList List, rootIDs []int) {
	for _, item := range list {
		if item.IsRoot() {
			outTree = append(outTree, item.ToTreeItemWithoutChildren())
			rootIDs = append(rootIDs, item.ID)
		} else {
			modifiersList = append(modifiersList, item)
		}
	}

	return outTree, modifiersList, rootIDs
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

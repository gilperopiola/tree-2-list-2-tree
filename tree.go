package main

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

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListToTree(t *testing.T) {

	// prepare

	list1 := List{
		{ID: 1},
		/* */ {ID: 3, ParentItemID: 1},
		/* */ /* */ {ID: 7, ParentItemID: 3},
		/* */ /* */ /* */ {ID: 10, ParentItemID: 7},
		/* */ /* */ /* */ /* */ {ID: 11, ParentItemID: 10},
		/* */ {ID: 4, ParentItemID: 1},
		/* */ /* */ {ID: 8, ParentItemID: 4},
		/* */ /* */ {ID: 9, ParentItemID: 4},
		{ID: 2},
		/* */ {ID: 5, ParentItemID: 2},
		/* */ {ID: 6, ParentItemID: 2},
	}

	tree1 := Tree{
		{
			ID: 1,
			Children: Tree{
				{ID: 3, ParentItemID: 1,
					Children: Tree{
						{ID: 7, ParentItemID: 3,
							Children: Tree{
								{ID: 10, ParentItemID: 7,
									Children: Tree{
										{ID: 11, ParentItemID: 10},
									},
								},
							}},
					}},
				{ID: 4, ParentItemID: 1,
					Children: Tree{
						{ID: 8, ParentItemID: 4},
						{ID: 9, ParentItemID: 4},
					}},
			},
		},
		{
			ID: 2,
			Children: []TreeItem{
				{ID: 5, ParentItemID: 2},
				{ID: 6, ParentItemID: 2},
			},
		},
	}

	// act

	out1 := list1.ToTree()

	// assert

	assert.Equal(t, tree1, out1)
}

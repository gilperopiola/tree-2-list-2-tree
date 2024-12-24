package main

//func TestTreeToListToTree(t *testing.T) {
//	// prepare
//	tree := Tree{
//		{
//			ID: 1,
//			Children: Tree{
//				{ID: 3, ParentItemID: 1,
//					Children: Tree{
//						{ID: 7, ParentItemID: 3,
//							Children: Tree{
//								{ID: 10, ParentItemID: 7,
//									Children: Tree{
//										{ID: 11, ParentItemID: 10},
//									},
//								},
//							}},
//					}},
//				{ID: 4, ParentItemID: 1,
//					Children: Tree{
//						{ID: 8, ParentItemID: 4},
//						{ID: 9, ParentItemID: 4},
//					}},
//			},
//		},
//		{
//			ID: 2,
//			Children: []TreeItem{
//				{ID: 5, ParentItemID: 2},
//				{ID: 6, ParentItemID: 2},
//			},
//		},
//	}
//
//	// act
//	outList := tree.ToList()
//	outTree := outList.ToTree()
//
//	// assert
//	assert.Equal(t, tree, outTree)
//}
//
//func TestTreeToList(t *testing.T) {
//	// prepare
//	tree := Tree{
//		{
//			ID: 1,
//			Children: Tree{
//				{ID: 3, ParentItemID: 1,
//					Children: Tree{
//						{ID: 7, ParentItemID: 3,
//							Children: Tree{
//								{ID: 10, ParentItemID: 7,
//									Children: Tree{
//										{ID: 11, ParentItemID: 10},
//									},
//								},
//							}},
//					}},
//				{ID: 4, ParentItemID: 1,
//					Children: Tree{
//						{ID: 8, ParentItemID: 4},
//						{ID: 9, ParentItemID: 4},
//					}},
//			},
//		},
//		{
//			ID: 2,
//			Children: Tree{
//				{ID: 5, ParentItemID: 2},
//				{ID: 6, ParentItemID: 2},
//			},
//		},
//	}
//
//	list := List{
//		{ID: 1},
//		/* */ {ID: 3, ParentItemID: 1},
//		/* */ /* */ {ID: 7, ParentItemID: 3},
//		/* */ /* */ /* */ {ID: 10, ParentItemID: 7},
//		/* */ /* */ /* */ /* */ {ID: 11, ParentItemID: 10},
//		/* */ {ID: 4, ParentItemID: 1},
//		/* */ /* */ {ID: 8, ParentItemID: 4},
//		/* */ /* */ {ID: 9, ParentItemID: 4},
//		{ID: 2},
//		/* */ {ID: 5, ParentItemID: 2},
//		/* */ {ID: 6, ParentItemID: 2},
//	}
//
//	// act
//	out := tree.ToList()
//
//	// assert
//	assert.Equal(t, list, out)
//}
//
//func TestListToTree(t *testing.T) {
//
//	// prepare
//	list := List{
//		{ID: 1},
//		/* */ {ID: 3, ParentItemID: 1},
//		/* */ /* */ {ID: 7, ParentItemID: 3},
//		/* */ /* */ /* */ {ID: 10, ParentItemID: 7},
//		/* */ /* */ /* */ /* */ {ID: 11, ParentItemID: 10},
//		/* */ {ID: 4, ParentItemID: 1},
//		/* */ /* */ {ID: 8, ParentItemID: 4},
//		/* */ /* */ {ID: 9, ParentItemID: 4},
//		{ID: 2},
//		/* */ {ID: 5, ParentItemID: 2},
//		/* */ {ID: 6, ParentItemID: 2},
//	}
//
//	tree := Tree{
//		{
//			ID: 1,
//			Children: Tree{
//				{ID: 3, ParentItemID: 1,
//					Children: Tree{
//						{ID: 7, ParentItemID: 3,
//							Children: Tree{
//								{ID: 10, ParentItemID: 7,
//									Children: Tree{
//										{ID: 11, ParentItemID: 10},
//									},
//								},
//							}},
//					}},
//				{ID: 4, ParentItemID: 1,
//					Children: Tree{
//						{ID: 8, ParentItemID: 4},
//						{ID: 9, ParentItemID: 4},
//					}},
//			},
//		},
//		{
//			ID: 2,
//			Children: []TreeItem{
//				{ID: 5, ParentItemID: 2},
//				{ID: 6, ParentItemID: 2},
//			},
//		},
//	}
//
//	// act
//	out := list.ToTree()
//
//	// assert
//	assert.Equal(t, tree, out)
//}

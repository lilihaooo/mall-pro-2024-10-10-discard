package product

import (
	"admin-gin/global"
	"admin-gin/model/product"
)

type CategoryService struct{}

//var CategoryServiceApp = new(CategoryService)

func (*CategoryService) GetCategoryTree() (tree []product.Category, err error) {
	err = global.XTK_DB.Order("level asc, sort asc").Find(&tree).Error
	if err != nil {
		return nil, err
	}
	tree = buildTree(tree, 0)
	return tree, nil
}

func buildTree(all []product.Category, parentID uint) (tree []product.Category) {
	for _, menu := range all {
		if menu.ParentID == parentID {
			subMenu := buildTree(all, menu.ID)
			menu.Children = subMenu
			tree = append(tree, menu)
		}
	}
	return
}

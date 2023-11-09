package db

func Assets_All() (assets_all []Assets, msg string) {
	db.Unscoped().Find(&assets_all)
	return assets_all, "获取所有资产组成功"
}

func Assets_Add(content string) (success bool, msg string) {
	result := db.Create(&Assets{Content: content})
	if result.RowsAffected == 1 {
		success = true
		msg = "新增资产组成功"
	} else {
		success = false
		msg = "新增资产组成功"
	}
	return success, msg
}

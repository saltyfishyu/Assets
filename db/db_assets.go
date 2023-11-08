package db

func Assets_All() (assets_all []Assets) {
	db.Unscoped().Find(&assets_all)
	return assets_all
}

func Assets_Add(content string) bool {
	result := db.Create(&Assets{Content: content})
	var success bool
	if result.RowsAffected == 1 {
		success = true
	} else {
		success = false
	}
	return success
}

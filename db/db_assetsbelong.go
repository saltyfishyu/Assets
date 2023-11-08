package db

func AssetsBelong_Add(aid int, domain string) bool {
	result := db.Create(&AssetsBelong{Aid: aid, Domain: domain})
	var success bool
	if result.RowsAffected == 1 {
		success = true
	} else {
		success = false
	}
	return success
}

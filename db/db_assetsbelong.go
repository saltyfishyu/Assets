package db

import (
	"Assets/util"
	"strings"
)

func AssetsBelong_Add(aid int, domain string) (success bool, msg string) {
	var content AssetsBelong
	db.Unscoped().Find(&content, "aid = ?", aid)
	if content == (AssetsBelong{}) {
		result := db.Create(&AssetsBelong{Domain: domain})
		if result.RowsAffected == 1 {
			success = true
			msg = "新增资产成功"
		} else {
			success = false
			msg = "新增资产失败"
		}
	} else {
		domains := strings.Split(content.Domain, ",")
		var finaldomain string
		if strings.Contains(domain, ",") {
			domainsplit := strings.Split(domain, ",")
			domainall := append(domains, domainsplit...)
			domainall = util.RemoveDuplicate(domainall)
			for _, str := range domainall {
				finaldomain = finaldomain + str + ","
			}
			finaldomain = finaldomain[:len(finaldomain)-1]
		} else {
			if util.In(domain, domains) {
				return false, "该资产已存在"
			} else {
				domains = append(domains, domain)
				for _, str := range domains {
					finaldomain = finaldomain + str + ","
				}
				finaldomain = finaldomain[:len(finaldomain)-1]
			}
		}
		result := db.Model(&content).Updates(&AssetsBelong{Domain: finaldomain})
		if result.RowsAffected == 1 {
			success = true
			msg = "更新资产成功"
		} else {
			success = false
			msg = "更新资产成功"
		}
	}

	return success, msg
}

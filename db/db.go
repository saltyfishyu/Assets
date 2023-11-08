package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// 资产组 比如 ID:1 Content:京东
type Assets struct {
	gorm.Model
	Content string `json:"content"` // 描述
}

// 资产组对应资产 比如Aid:1 即京东 domain:jd.com ...
type AssetsBelong struct {
	gorm.Model
	Aid    int    `json:"aid"`    // 资产id
	Domain string `json:"domain"` // 域名、IP
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("assets.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库出错：", err.Error())
		return
	}

	db.AutoMigrate(&Assets{})
	db.AutoMigrate(&AssetsBelong{})
}

func New() {

}

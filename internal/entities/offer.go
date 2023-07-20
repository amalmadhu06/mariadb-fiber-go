package entities

import (
	"time"
)

// OfferCompany represents offers
type OfferCompany struct {
	OfferID            uint      `gorm:"column:offer_id;primaryKey"`
	ClientID           uint      `gorm:"column:client_id"`
	Country            string    `gorm:"column:country"`
	Image              string    `gorm:"column:image"`
	ImageWidth         int       `gorm:"column:image_width"`
	ImageHeight        int       `gorm:"column:image_height"`
	TextLocale         string    `gorm:"column:text_locale"`
	ValidityTextLocale string    `gorm:"column:validity_text_locale"`
	Position           int       `gorm:"column:position"`
	ValidFrom          time.Time `gorm:"column:valid_from"`
	ShowFrom           time.Time `gorm:"column:show_from"`
	ValidTo            time.Time `gorm:"column:valid_to"`
	Flag               uint      `gorm:"column:flag"`
	PageCount          uint      `gorm:"column:page_count"`
	StoreURL           string    `gorm:"column:store_url"`
	StoreURLTitle      string    `gorm:"column:store_url_title"`
	OfferHome          int       `gorm:"column:offer_home"`
}

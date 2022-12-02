package db

import "gorm.io/gorm"

type CdCount struct {
	ID        int64  `gorm:"column:id" json:"id"`
	AllCount  int64  `gorm:"column:all_count" json:"all_count"`
	CountType string `gorm:"column:count_type" json:"count_type"` //  类型 全市/住宅
	Year      int64  `gorm:"column:year" json:"year"`
	Month     string `gorm:"column:month" json:"month"`
	Day       int64  `gorm:"column:day" json:"day"`
}

func (c *CdCount) Create() error {
	return db.Create(c).Error
}

func (c *CdCount) GetCount() ([]CdCount, error) {
	var ct []CdCount
	result := db.Model(&c).Where("count_type = 'dwelling'").Order("id DESC").Limit(10).Find(&ct)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return ct, nil
}

package db

import "gorm.io/gorm"

type Count struct {
	ID    int64  `gorm:"column:id" json:"id"`
	Area  string `gorm:"column:area" json:"area"`   //  区域
	Plate string `gorm:"column:plate" json:"plate"` //  板块
	Count int64  `gorm:"column:count" json:"count"` //  数量
	Year  int64  `gorm:"column:year" json:"year"`
	Month string `gorm:"column:month" json:"month"`
	Day   int64  `gorm:"column:day" json:"day"`
}

func (c *Count) Create() error {
	return db.Create(c).Error
}

func (c *Count) GetCount(area, plate string) ([]Count, error) {
	var ct []Count
	result := db.Model(&c).Where("area = ? AND plate = ?", area, plate).Order("id DESC").Limit(10).Find(&ct)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return ct, nil
}

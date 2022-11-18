package db

type HouseInfo struct {
	ID                  int64   `gorm:"column:id" json:"id"`                                       //  id
	HouseId             string  `gorm:"column:house_id" json:"house_id"`                           //  链家房屋id
	Title               string  `gorm:"column:title" json:"title"`                                 //  房源标题
	HouseType           string  `gorm:"column:house_type" json:"house_type"`                       //  几居室
	Face                string  `gorm:"column:face" json:"face"`                                   //  朝向
	Acreage             float64 `gorm:"column:acreage" json:"acreage"`                             //  面积
	TotalPrice          int64   `gorm:"column:total_price" json:"total_price"`                     //  总价
	UnitPrice           int64   `gorm:"column:unit_price" json:"unit_price"`                       //  单价
	Area                string  `gorm:"column:area" json:"area"`                                   //  区域
	Plate               string  `gorm:"column:plate" json:"plate"`                                 //  板块
	ResidentialAreaName string  `gorm:"column:residential_area_name" json:"residential_area_name"` //  小区名称
	HouseBirth          string  `gorm:"column:house_birth" json:"house_birth"`                     //  建成年限
	Floor               string  `gorm:"column:floor" json:"floor"`                                 //  楼层分布
	TotalHeight         int     `gorm:"column:total_height" json:"total_height"`                   //  楼层总高
	PublishTime         string  `gorm:"column:publish_time" json:"publish_time"`                   //  上架时间
	CreatedTimeMonth    string  `gorm:"column:created_time_month" json:"created_time_month"`       //  创建时间-月
	CreatedTimeDay      int     `gorm:"column:created_time_day" json:"created_time_day"`           //  创建时间-天
}

func CreatHouseInfo(houses []HouseInfo) {
	db.Table("house_info").Create(&houses)
}

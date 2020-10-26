/*
@Time : 2020/10/26 11:50
@Author : lai
@Description :
@File : vehicle_location_latest
*/
package model

import "time"

type VehicleLocationLatest struct {
	ID              int64      `gorm:"primary_key;column:id;type:bigint;default:id_generator()";sql:"index;unique" json:"id"` //主键ID
	Altitude        string     `gorm:"column:altitude;type:text" json:"altitude"`                                             //海拔
	Direction       string     `gorm:"column:direction;type:text" json:"direction"`                                           //方向
	DrivingSpeed    string     `gorm:"column:driving_speed;type:text" json:"driving_speed"`                                   //行驶记录仪速度
	Longitude       string     `gorm:"column:longitude;type:text" json:"longitude"`                                           //经度
	Latitude        string     `gorm:"column:latitude;type:text" json:"latitude"`                                             //纬度
	Mileage         string     `gorm:"column:mileage;type:text" json:"mileage"`                                               //里程
	Speed           string     `gorm:"column:speed;type:text" json:"speed"`                                                   //GPS速度
	VehicleID       string     `gorm:"column:vehicle_id;type:text" json:"vehicle_id"`                                         //车辆ID
	EncodeLongitude string     `gorm:"column:encode_longitude;type:text" json:"encode_longitude"`                             //纠偏后经度
	EncodeLatitude  string     `gorm:"column:encode_latitude;type:text" json:"encode_latitude"`                               //纠偏后纬度
	Status          string     `gorm:"column:status;type:text" json:"status"`                                                 //车辆状态
	LocateTime      time.Time  `gorm:"column:locate_time;type:timestamptz" json:"locate_time"`                                //定位时间
	RecordTime      time.Time  `gorm:"column:record_time;type:timestamptz" json:"record_time"`                                //记录时间
	CreatedAt       time.Time  `gorm:"column:create_at;type:timestamptz" json:"create_at"`                                    //创建时间
	UpdatedAt       time.Time  `gorm:"column:update_at;type:timestamptz" json:"update_at"`                                    //修改时间
	DeletedAt       *time.Time `gorm:"column:delete_at;type:timestamptz" json:"delete_at"`                                    //删除时间
}

func (VehicleLocationLatest) TableName() string {
	return "vehicle_location_latest"
}

/*
@Time : 2020/10/14 8:59
@Author : Administrator
@Description :
@File : create_test
@Software: GoLand
*/
package curd

import (
	"fmt"
	"github.com/rs/xid"
	"lai.com/GraphQL_Server/graph/model"
	"testing"
	"time"
)

var c Create

func init() {
	c.Init()
}

func TestCreateEnterpriseInfo(t *testing.T) {

	for i := 0; i < 10000000; i++ {
		enterpriseInfo := &model.EnterpriseInfo{
			EnterpriseID:    xid.New().String(),
			EnterpriseCode:  string(i),
			EnterpriseName:  "辉哥股份有限公司",
			EnterpriseLevel: i,
			Remark:          fmt.Sprintf("第%d个备注", i),
		}
		c.CreateEnterpriseInfo(enterpriseInfo)
	}
	//enterpriseInfo := &model.EnterpriseInfo{
	//	EnterpriseName: "小灰灰科技有限公司",
	//	EnterpriseID:   xid.New().String(),
	//}
	//err := c.CreateEnterpriseInfo(enterpriseInfo)
	//if err != nil {
	//	fmt.Printf("CreateEnterpriseInfo failed, err=%v\n", err)
	//}
}

const RULE_TIME_FORMAT = "2006-01-02 15:04:05"

var lt1, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:54:55", time.Local)
var lt2, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:28", time.Local)
var lt3, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:22", time.Local)
var lt4, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:13", time.Local)
var lt5, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:46:58", time.Local)
var lt6, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:46:54", time.Local)
var lt7, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:46:53", time.Local)
var lt8, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:46:51", time.Local)
var lt9, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:46:47", time.Local)
var lt10, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:46:46", time.Local)

var rt1, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:55:02", time.Local)
var rt2, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:58", time.Local)
var rt3, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:37", time.Local)
var rt4, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:28", time.Local)
var rt5, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:18", time.Local)
var rt6, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:18", time.Local)
var rt7, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:18", time.Local)
var rt8, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:18", time.Local)
var rt9, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:12", time.Local)
var rt10, _ = time.ParseInLocation(RULE_TIME_FORMAT, "2020-10-21 17:47:12", time.Local)

var VehicleLocationLatestList = []model.VehicleLocationLatest{
	{
		Altitude:        "124",
		Direction:       "302",
		DrivingSpeed:    "42",
		Latitude:        "25.178219",
		Longitude:       "118.269329",
		Mileage:         "38691.7",
		LocateTime:      lt1,
		RecordTime:      rt1,
		Speed:           "46",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1754370874615",
		EncodeLongitude: "118.274084217264",
		Status:          "397467",
	},
	{
		Altitude:        "100",
		Direction:       "160",
		DrivingSpeed:    "31",
		Latitude:        "25.194358",
		Longitude:       "118.285826",
		Mileage:         "38688",
		LocateTime:      lt2,
		RecordTime:      rt2,
		Speed:           "34",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1915933443648",
		EncodeLongitude: "118.290594017174",
		Status:          "397467",
	},
	{
		Altitude:        "99",
		Direction:       "160",
		DrivingSpeed:    "39",
		Latitude:        "25.194891",
		Longitude:       "118.285629",
		Mileage:         "38688",
		LocateTime:      lt3,
		RecordTime:      rt3,
		Speed:           "44",
		VehicleID:       "70876",
		EncodeLatitude:  "25.192126117598",
		EncodeLongitude: "118.290396860237",
		Status:          "430491",
	},
	{
		Altitude:        "100",
		Direction:       "160",
		DrivingSpeed:    "39",
		Latitude:        "25.195844",
		Longitude:       "118.285289",
		Mileage:         "38687.9",
		LocateTime:      lt4,
		RecordTime:      rt4,
		Speed:           "44",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1930787295609",
		EncodeLongitude: "118.290056594724",
		Status:          "397467",
	},
	{
		Altitude:        "99",
		Direction:       "163",
		DrivingSpeed:    "35",
		Latitude:        "25.197389",
		Longitude:       "118.284731",
		Mileage:         "38687.7",
		LocateTime:      lt5,
		RecordTime:      rt5,
		Speed:           "37",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1946230976798",
		EncodeLongitude: "118.289498165778",
		Status:          "397467",
	},
	{
		Altitude:        "100",
		Direction:       "166",
		DrivingSpeed:    "39",
		Latitude:        "25.197774",
		Longitude:       "118.284624",
		Mileage:         "38687.7",
		LocateTime:      lt6,
		RecordTime:      rt6,
		Speed:           "43",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1950079805417",
		EncodeLongitude: "118.289391091077",
		Status:          "430491",
	},
	{
		Altitude:        "100",
		Direction:       "167",
		DrivingSpeed:    "39",
		Latitude:        "25.197881",
		Longitude:       "118.284598",
		Mileage:         "38687.7",
		LocateTime:      lt7,
		RecordTime:      rt7,
		Speed:           "43",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1951149525961",
		EncodeLongitude: "118.28936507395",
		Status:          "397467",
	},
	{
		Altitude:        "101",
		Direction:       "171",
		DrivingSpeed:    "39",
		Latitude:        "25.198098",
		Longitude:       "118.284554",
		Mileage:         "38687.6",
		LocateTime:      lt8,
		RecordTime:      rt8,
		Speed:           "43",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1953319066185",
		EncodeLongitude: "118.289321047602",
		Status:          "430491",
	},
	{
		Altitude:        "101",
		Direction:       "174",
		DrivingSpeed:    "40",
		Latitude:        "25.198544",
		Longitude:       "118.284499",
		Mileage:         "38687.6",
		LocateTime:      lt9,
		RecordTime:      rt9,
		Speed:           "44",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1957778553988",
		EncodeLongitude: "118.289266027331",
		Status:          "397467",
	},
	{
		Altitude:        "101",
		Direction:       "175",
		DrivingSpeed:    "39",
		Latitude:        "25.198656",
		Longitude:       "118.284488",
		Mileage:         "38687.6",
		LocateTime:      lt10,
		RecordTime:      rt10,
		Speed:           "45",
		VehicleID:       "70876",
		EncodeLatitude:  "25.1958898459546",
		EncodeLongitude: "118.28925502493",
		Status:          "430491",
	},
}

func TestCreateLocationLatest(t *testing.T) {
	//for i := 0; i < 10000000; i++ {
	//	for _, locate := range VehicleLocationLatestList {
	//		err := c.CreateLocationLatest(locate)
	//		if err != nil {
	//			fmt.Printf("err:%v", err)
	//		}
	//	}
	//}
	start := time.Now()
	var vList []model.VehicleLocationLatest
	for i := 0; i < 100000; i++ {
		for _, v := range VehicleLocationLatestList {
			vList = append(vList, v)
		}
	}
	fmt.Printf("create models complete %.2fs\n", time.Since(start).Seconds())
	/*for i := 0; i < 40; i++ {
		BatchSave(c.db, vList)
	}*/
	fmt.Printf("add database complete %.2fs\n", time.Since(start).Seconds())

}

// BatchSave 批量插入数据
//func BatchSave(db *gorm.DB, vList []model.VehicleLocationLatest) error {
//	var buffer bytes.Buffer
//	sql := "insert into vehicle_location_latest(altitude,direction,driving_speed,longitude,latitude,mileage,speed,vehicle_id,encode_longitude,encode_latitude,status,locate_time,record_time) values"
//	if _, err := buffer.WriteString(sql); err != nil {
//		return err
//	}
//	for i, e := range vList {
//		if i == len(vList)-1 {
//			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');", e.Altitude, e.Direction, e.DrivingSpeed, e.Longitude, e.Latitude, e.Mileage, e.Speed, e.VehicleID, e.EncodeLongitude, e.EncodeLatitude, e.Status, e.LocateTime.Format(RULE_TIME_FORMAT), e.RecordTime.Format(RULE_TIME_FORMAT)))
//		} else {
//			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'),", e.Altitude, e.Direction, e.DrivingSpeed, e.Longitude, e.Latitude, e.Mileage, e.Speed, e.VehicleID, e.EncodeLongitude, e.EncodeLatitude, e.Status, e.LocateTime.Format(RULE_TIME_FORMAT), e.RecordTime.Format(RULE_TIME_FORMAT)))
//
//		}
//	}
//	//fmt.Println(buffer.String())
//	return db.Exec(buffer.String()).Error
//}

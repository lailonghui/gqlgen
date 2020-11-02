/*
@Time : 2020/11/2 9:54
@Author : lai
@Description :
@File : alarmInfo
*/
package graph

import "time"

// 经纬度
type Point struct {
	// 经度
	Longitude float64 `json:"longitude"`
	// 纬度
	Latitude float64 `json:"latitude"`
}

// 车辆报警信息
type VehicleAlarmData struct {
	// 主键id
	ID string `json:"id"`
	// 由golang程序生成的xid，暴露到外部使用
	AlarmDataID string `json:"alarm_data_id"`
	// 车辆ID
	VehicleID string `json:"vehicle_id"`
	// 所属区域ID
	Pid string `json:"pid"`
	// 车辆速度
	Speed float64 `json:"speed"`
	// 记录时间
	RecordTime time.Time `json:"record_time"`
	// 位置描述
	PosDesc string `json:"pos_desc"`
	// 空间数据类型point表示经纬度
	Coordinate *Point `json:"coordinate"`
	// 行车记录仪速度
	DrivingSpeed float64 `json:"driving_speed"`
	// 处理状态(1.处理中 2.处理完毕 3.不作处理 4.将来处理)
	DealingStatus int `json:"dealing_status"`
	// 处理人
	UserID string `json:"user_id"`
	// 处理描述
	DealDesc string `json:"deal_desc"`
	// 报警类型(001.超速报警 002.疲劳驾驶报警 066.凌晨2点到5点行驶报警)
	AlarmType int `json:"alarm_type"`
	// 报警开始时间
	AlarmBeginTime time.Time `json:"alarm_begin_time"`
	// 报警结束时间
	AlarmEndTime time.Time `json:"alarm_end_time"`
	// 车牌号
	LicensePlateNumber string `json:"license_plate_number"`
	// 车牌颜色
	LicensePlateColor int `json:"license_plate_color"`
}

// 车辆监控图片
type VehicleSupervisionPhoto struct {
	// 主键id
	ID string `json:"id"`
	// 由golang程序生成的xid，暴露到外部使用
	SupervisionPhotoID string `json:"supervision_photo_id"`
	// 车辆ID
	VehicleID string `json:"vehicle_id"`
	// 驾驶员id
	DriverID string `json:"driver_id"`
	// 所在企业id
	EnterpriseID string `json:"enterprise_id"`
	// 终端上报时间
	UpdateTime time.Time `json:"update_time"`
	// 监控图片地址
	MonitoringPic string `json:"monitoring_pic"`
	// 监控图片上传时间
	MonitoringPicUploadTime time.Time `json:"monitoring_pic_upload_time"`
	// 终端IMEI
	Imei string `json:"IMEI"`
	// SIM卡号
	SimNumber string `json:"SIM_number"`
}

// 车辆终端达标情况
type VehicleTerminalStatus struct {
	// 主键id
	ID string `json:"id"`
	// ACC状态
	Acc string `json:"acc"`
	// 制动信号(刹车)
	Brake string `json:"brake"`
	// 左转向灯信号
	LeftLamp string `json:"left_lamp"`
	// 右转向灯信号
	RightLamp string `json:"right_lamp"`
	// 近光灯信号
	NearLamp string `json:"near_lamp"`
	// 远光灯信号(大灯)
	FarLamp string `json:"far_lamp"`
	// 喇叭信号
	LoudSpeaker string `json:"loud_speaker"`
	// 定位状态
	Locate string `json:"locate"`
	// 前门
	FrontDoor string `json:"front_door"`
	// GPS信号
	GpsOpen string `json:"GPS_open"`
	// 北斗信号
	BdOpen string `json:"BD_open"`
	// 摄像头
	Camera string `json:"camera"`
	// 车速
	DrivingSpeed string `json:"driving_speed"`
	// 备注
	Remarks string `json:"remarks"`
	// 记录时间
	RecordTime time.Time `json:"record_time"`
}

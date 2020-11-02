/*
@Time : 2020/10/31 16:56
@Author : lai
@Description :
@File : driver
*/
package graph

import "time"

// 驾驶员身份验证信息(暂把数据表中所有信息列出，实际使用时根据具体业务和安全性选择要返回的信息)
type DriverIdentity struct {
	// 主键id
	ID string `json:"id"`
	// 由golang程序生成的xid，暴露到外部使用
	IdentityID string `json:"identity_id"`
	// 身份证号码
	IDCardNum string `json:"id_card_num"`
	// 身份证出生日期
	IDCardBirthday time.Time `json:"id_card_birthday"`
	// 身份证签发机关
	IDCardSignGovernment string `json:"id_card_sign_government"`
	// 身份证民族
	IDCardNation string `json:"id_card_nation"`
	// 身份证有效起始日期
	IDCardStartDate time.Time `json:"id_card_start_date"`
	// 身份证有效截止日期
	IDCardEndDate time.Time `json:"id_card_end_date"`
	// 身份证正面照，云存储地址
	IDCardFrontPic string `json:"id_card_front_pic"`
	// 身份证背面照，云存储地址
	IDCardBackPic string `json:"id_card_back_pic"`
	// 身份证住址
	IDCardAddress string `json:"id_card_address"`
	// 驾驶员手持身份证照片,云储存系统返回的路径
	DriverHoldingIDPhoto string `json:"driver_holding_id_photo"`
	// 驾驶员的正面照,云储存系统返回的路径
	DriverPhoto string `json:"driver_photo"`
	// 驾驶员签名,云储存系统返回的路径
	DriverSignature string `json:"driver_signature"`
	// 从业资格证号码
	OccupationalNumber string `json:"occupational_number"`
	// 从业资格证有效期至
	OccupationalExpireDate time.Time `json:"occupational_expire_date"`
	// 从业资格证发证机构
	OccupationalIssuingAuthority *string `json:"occupational_issuing_authority"`
	// 劳动合同,云储存系统返回的完整劳动合同的图片路径
	LaborContract []string `json:"labor_contract"`
	// 驾驶员驾驶证,云储存系统返回的路径
	DriverLicensePic string `json:"driver_license_pic"`
	// 驾驶证发证机关
	DriverLicenseIssuingAuthority string `json:"driver_license_issuing_authority"`
	// 年审日期（六合一）
	AnnualReviewDate time.Time `json:"annual_review_date"`
	// 换证日期（六合一）
	RenewalDate time.Time `json:"renewal_date"`
	// 累计积分（六合一）
	AccumulativedPoints float64 `json:"accumulatived_points"`
	// 清分日期（六合一）
	SortingDate time.Time `json:"sorting_date"`
	// 准驾车型（六合一）
	QuasiDrivingModels int `json:"quasi_driving_models"`
	// 驾驶证发证所在地的省份ID
	DriverLicenseProvinceID string `json:"driver_license_province_id"`
	// 驾驶证发证所在地的城市ID
	DriverLicenseCityID string `json:"driver_license_city_id"`
	// 驾驶证发证所在地的区域ID
	DriverLicenseDistrictID string `json:"driver_license_district_id"`
	// 驾驶证状态
	DriverLicenseStatus int `json:"driver_license_status"`
	// 驾驶证初次领证日期
	DriverLicenseIssueDate time.Time `json:"driver_license_issue_date"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy string `json:"delete_by"`
}

// 驾驶员信息(暂把数据表中所有信息列出，实际使用时根据具体业务和安全性选择要返回的信息)
type DriverInfo struct {
	// 主键id
	ID string `json:"id"`
	// 驾驶员外部编码,由golang程序生成的xid，暴露到外部使用
	DriverID string `json:"driver_id"`
	// 所在企业id
	EnterpriseID string `json:"enterprise_id"`
	// 所在部门id
	DepartmentID string `json:"department_id"`
	// 驾驶员身份验证信息ID
	DriverIdentityID string `json:"driver_identity_id"`
	// 驾驶员姓名
	DriverName string `json:"driver_name"`
	// 手机号码
	Telephone string `json:"telephone"`
	// 性别
	Sex int `json:"sex"`
	// 档案编号(后6位)
	FilesNumber string `json:"files_number"`
	// 联系地址
	ContactAddress string `json:"contact_address"`
	// 邮寄地址
	MailingAddress string `json:"mailing_address"`
	// 是否提交
	IsSubmit bool `json:"is_submit"`
	// 提交内容
	SubmitContent string `json:"submit_content"`
	// 提交时间
	SubmitAt time.Time `json:"submit_at"`
	// 提交人
	SubmitBy string `json:"submit_by"`
	// 是否手动录入
	IsManualInput bool `json:"is_manual_input"`
	// 是否录入
	IsInput bool `json:"is_input"`
	// 录入时间
	InputAt time.Time `json:"input_at"`
	// 录入人
	InputBy string `json:"input_by"`
	// 是否校验数据
	IsCheckData bool `json:"is_check_data"`
	// 检验时间
	CheckAt time.Time `json:"check_at"`
	// 校验人
	CheckBy string `json:"check_by"`
	// 驾驶员信息同步内网反馈信息
	RemarkIn string `json:"remark_in"`
	// 内网更新时间
	UpdateTimeIn time.Time `json:"update_time_in"`
	// 是否通过短信验证
	IsCheckSms bool `json:"is_check_sms"`
	// 是否黑名单
	IsBlack bool `json:"is_black"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy string `json:"delete_by"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
}

// 驾驶员信息的输入(具体输入信息待确认)
type NewDriverInfo struct {
	// 所在企业id
	EnterpriseID string `json:"enterprise_id"`
	// 所在部门id
	DepartmentID string `json:"department_id"`
	// 驾驶员身份验证信息ID
	DriverIdentityID string `json:"driver_identity_id"`
	// 驾驶员姓名
	DriverName string `json:"driver_name"`
	// 手机号码
	Telephone string `json:"telephone"`
	// 性别
	Sex int `json:"sex"`
	// 档案编号(后6位)
	FilesNumber string `json:"files_number"`
	// 联系地址
	ContactAddress string `json:"contact_address"`
	// 邮寄地址
	MailingAddress string `json:"mailing_address"`
}

// 车辆驾驶员绑定表(
type VehicleDriverBinding struct {
	// 主键id
	ID string `json:"id"`
	// 由golang程序生成的xid，暴露到外部使用
	VehicleDriverBindingID string `json:"vehicle_driver_binding_id"`
	// 驾驶员id
	DriverID string `json:"driver_id"`
	// 车辆id
	VehicleID string `json:"vehicle_id"`
	// 备注
	Remarks string `json:"remarks"`
	// 租赁合同,,云储存系统返回的完整租赁合同的图片路径
	LeaseContract []string `json:"lease_contract"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy string `json:"create_by"`
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 修改人
	UpdateBy string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy string `json:"delete_by"`
}

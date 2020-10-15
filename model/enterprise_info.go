package model

//type EnterpriseInfo2 struct {
//	ID             int64      `gorm:"primary_key;column:id;type:bigint;default:id_generator()";sql:"index;unique" json:"id"` //主键ID
//	EnterpriseID   string     `gorm:"column:enterprise_id;type:text;unique_index" json:"enterprise_id"`                      //企业id,联合主键
//	EnterpriseName string     `gorm:"column:enterprise_name;type:text" json:"enterprise_name"`                               //企业名称
//	EnterpriseCode string     `gorm:"column:enterprise_code;type:text" json:"enterprise_code"`                               //企业码
//	CreatedAt      time.Time  `gorm:"column:create_at;type:timestamptz" json:"create_at"`                                    //创建时间
//	UpdatedAt      time.Time  `gorm:"column:update_at;type:timestamptz" json:"update_at"`                                    //修改时间
//	DeletedAt      *time.Time `gorm:"column:delete_at;type:timestamptz" json:"delete_at"`                                    //删除时间
//}
//
//func (EnterpriseInfo2) TableName() string {
//	return "enterprise_info"
//}

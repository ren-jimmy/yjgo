// ==========================================================================
// 云捷GO自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成路径: app/model/module/online/online_entity.go
// 生成人：yunjie
// ==========================================================================

package online

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table sys_user_online.
type Entity struct {
	SessionId      string      `orm:"sessionId,primary" json:"sessionId"`       // 用户会话id
	LoginName      string      `orm:"login_name" json:"login_name"`             // 登录账号
	DeptName       string      `orm:"dept_name" json:"dept_name"`               // 部门名称
	Ipaddr         string      `orm:"ipaddr" json:"ipaddr"`                     // 登录IP地址
	LoginLocation  string      `orm:"login_location" json:"login_location"`     // 登录地点
	Browser        string      `orm:"browser" json:"browser"`                   // 浏览器类型
	Os             string      `orm:"os" json:"os"`                             // 操作系统
	Status         string      `orm:"status" json:"status"`                     // 在线状态on_line在线off_line离线
	StartTimestamp *gtime.Time `orm:"start_timestamp" json:"start_timestamp"`   // session创建时间
	LastAccessTime *gtime.Time `orm:"last_access_time" json:"last_access_time"` // session最后访问时间
	ExpireTime     int         `orm:"expire_time" json:"expire_time"`           // 超时时间，单位为分钟
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}

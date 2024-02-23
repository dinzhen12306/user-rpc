package mysql

import "time"

// 用户表
type User struct {
	Id       int64      `xorm:"pk autoincr"`
	Username string     `xorm:"varchar(50) comment('用户名')"`
	Password string     `xorm:"varchar(32) comment('密码32位md5加密')"`
	Phone    string     `xorm:"varchar(11) comment('电话号码')"`
	Email    string     `xorm:"varchar(30) comment('邮箱')"`
	Token    string     `xorm:"varchar(255) comment('用户鉴权')"`
	Created  time.Time  `xorm:"created"`
	Updated  time.Time  `xorm:"updated"`
	Deleted  *time.Time `xorm:"deleted"`
}

// 添加
func InsertUser(user User) (User, error) {
	_, err = XDB.Insert(&user)
	return user, err
}

// 修改
func UpdateUser(user User) error {
	_, err = XDB.Where("id = ?", user.Id).Update(&user)
	return err
}

// 删除
func DeletedUser(user User) error {
	_, err = XDB.Delete(&user)
	return err
}

// 查询单条
func AccordingToConditionGetUser(Condition map[string][]byte) (user User, err error) {
	_, err = XDB.Where(Condition).Get(&user)
	return
}

// 查询多条
func AccordingToConditionGetUsers(Condition map[string][]byte) (user []User, err error) {
	err = XDB.Where(Condition).Find(&user)
	return
}

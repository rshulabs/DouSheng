// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	User          int64   `gorm:"column:user;type:bigint;primaryKey" json:"user"`
	Username      *string `gorm:"column:username;type:varchar(16)" json:"username"`
	Phone         string  `gorm:"column:phone;type:char(11);not null" json:"phone"`
	Name          string  `gorm:"column:name;type:char(16);not null" json:"name"`
	FollowCount   int64   `gorm:"column:follow_count;type:int;not null" json:"follow_count"`
	FollowerCount int64   `gorm:"column:follower_count;type:int;not null" json:"follower_count"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
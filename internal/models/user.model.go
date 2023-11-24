package models

import "time"

type UserModel struct {
	Id        int        `db:"id" json:"id" valid:"-"`
	Username  string     `db:"username" form:"username" json:"username" valid:"required,minstringlength(6)"`
	Email     string     `db:"email" form:"email" json:"email" valid:"required,email"`
	Password  string     `db:"userpass" form:"password" json:"password" valid:"required,minstringlength(4)"`
	Role      string     `db:"user_role" json:"role" valid:"-"`
	CreatedAt *time.Time `db:"created_at" valid:"-"`
	UpdatedAt *time.Time `db:"updated_at" valid:"-"`
}

type GetUserInfoModel struct {
	Id       int    `db:"id" json:"id" valid:"-"`
	Username string `db:"username" form:"username" json:"username" valid:"-"`
	Email    string `db:"email" form:"email" json:"email" valid:"required,email"`
	Password string `db:"userpass" form:"password" json:"password" valid:"required,minstringlength(4)"`
	Role     string `db:"user_role" json:"role" valid:"-"`
}

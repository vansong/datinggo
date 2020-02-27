package dao

import (
	"database/sql"
	"goinspect.cn/app/account"
	"goinspect.cn/app/conf"
	"goinspect.cn/app/library/database"
)

type Dao struct {
	db *sql.DB
}

func New(c *conf.Config) *Dao {
	d := &Dao{
		db: database.NewMysql(c.Database),
	}
	return d
}

func (d *Dao) GetUser(id int) (u *account.User, err error) {
	sql := "SELECT id, nickname, avatar FROM users WHERE id = ?"
	row := d.db.QueryRow(sql, id)

	u = &account.User{}
	err = row.Scan(&u.Id, &u.Nickname, &u.Avatar)

	return
}
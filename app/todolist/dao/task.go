package dao

import (
	"database/sql"
	"fmt"
	"goinspect.cn/app/conf"
	"goinspect.cn/app/library/database"
	"goinspect.cn/app/todolist"
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

func (d *Dao) GetTasks(uid int) (res []*todolist.Task, err error) {
	var rows *sql.Rows
	sql := "SELECT id, uid, title, status, updated_at, created_at FROM todo_tasks WHERE uid = ?";
	if rows, err = d.db.Query(sql, uid); err != nil {
		return
	}
	defer rows.Close()

	res = []*todolist.Task{}
	for rows.Next() {
		t := &todolist.Task{}
		if err = rows.Scan(&t.Id, &t.Uid, &t.Title, &t.Status, &t.Updated_At, &t.Created_At); err != nil {
			return
		}
		res = append(res, t)
	}
	return
}

func (d *Dao) NewTask(uid int, title string) bool {
	sql := "INSERT INTO todo_tasks(uid,title,status,updated_at,created_at) VALUES (?,?,?,NOW(),NOW())";
	rs, err := d.db.Exec(sql, uid, title, todolist.STATUS_PENDING)
	if err != nil {
		fmt.Println(err)
		// TODO log
		return false
	}
	if id, _ := rs.LastInsertId(); id > 0 {
		return true
	}
	return false
}

func (d *Dao) RemoveTask(id int) bool {
	sql := "DELETE FROM todo_tasks WHERE id = ?"
	_, err := d.db.Exec(sql, id)
	if err != nil {
		return false
	}
	return true;
}

func (d *Dao) UpdateTaskStatus(id int, status int) bool {
	sql := "UPDATE todo_tasks SET status = ? WHERE id = ?"
	_, err := d.db.Exec(sql, status, id)
	if err != nil {
		return false
	}
	return true
}
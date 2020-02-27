package paper

import (
	"database/sql"
	"goinspect.cn/app/conf"
	"goinspect.cn/app/library/database"
	"goinspect.cn/app/questionair/model/paper"
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

func (d *Dao) GetPapers( offset int, limit int) (res []*paper.Paper, err error) {
	var rows *sql.Rows
	sql := "SELECT id, title, subtitle FROM papers ORDER BY id DESC LIMIT ?,?";
	if rows, err = d.db.Query(sql, offset, limit); err != nil {
		return
	}
	defer rows.Close()

	res = []*paper.Paper{}
	for rows.Next() {
		p := &paper.Paper{}
		if err = rows.Scan(&p.Id, &p.Title, &p.Subtitle); err != nil {
			return
		}
		res = append(res, p)
	}
	return
}

func (d *Dao) GetPaper(id int) (res *paper.Paper, err error) {
	sql := "SELECT id, title, subtitle FROM papers WHERE id = ?";
	row := d.db.QueryRow(sql, id)

	res = &paper.Paper{}
	_ = row.Scan(&res.Id, &res.Title, &res.Subtitle)
	return
}



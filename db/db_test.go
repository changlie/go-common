package db

import (
	"fmt"
	"testing"
)

func Test_sqlite(t *testing.T) {
	db := Sqlite("t.db")
	db.Exec("create table if not exists Users(id integer primary key , name text, addr text)")
	db.Insert("insert into Users(name, addr) values (?,?)", "深圳", "广东")
	list := db.GetRows("select * from Users")
	for _, item := range list {
		fmt.Println(item)
	}
}

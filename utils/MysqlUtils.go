package utils

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func openDB(db_user, db_password, db_name string) *sql.DB{
	db, err := sql.Open("mysql", db_user + ":" + db_password + "@/" + db_name + "?charset=utf8")
	if err != nil {
		log.Fatalf("utils OpenDB() db, err := sql.Open() error => %v\n", err)
		return nil
	}
	return db
}

//mysql插入操作（裸写），返回回填id
func Insert(db_user, db_password, db_name, db_sql string, params ...interface{}) (int64, error){
	db := openDB(db_user, db_password, db_name)
	defer db.Close()

	stmt, err := db.Prepare(db_sql)
	if err != nil {
		log.Fatalf("utils Insert() stmt, err := db.Prepare() error => %v\n", err)
		return -1, err
	}

	result, err := stmt.Exec(params...)
	if err != nil {
		log.Fatalf("utils Insert() result, err := stmt.Exec() error => %v\n", err)
		return -1, err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("utils Insert() affect, err := result.RowsAffected() error => %v\n", err)
		return -1, err
	}

	return affect, nil
}

//mysql查找操作（裸写），返回rows
func Select(db_user, db_password, db_name, db_sql string, params ...interface{})(*sql.Rows, error){
	db := openDB(db_user, db_password, db_name)
	defer db.Close()

	rows, err := db.Query(db_sql, params...)
	if err != nil {
		log.Fatalf("utils Select() rows, err := db.Query() error => %v\n", err)
		return nil, err
	}

	return rows, nil
}

package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"time"
)
var(
	user = "hqzx_sit_rw"
	base64dbpassword = "QPbFI57VQYUnUAn"
	addr = "10.210.110.86:3306"
	dbname = "strategy_db"

	qryBase = "select id from test"
	qryTest = qryBase+" where name=? AND gender=?"

	insertTable = "test"
	insertSql = "insert into "+ insertTable +"(id,name,gender) values(?,?,?)"
)
type TestManager interface {
	QueryTest(name string, gender string) error
	Insert(id int64,name,gender string) error
}

type SqlScan interface {
	Scan(dst ...interface{}) error
}
type SqlPrepare interface {
	ExecContext(ctx context.Context,args ...interface{})(sql.Result,error)
}


type TestDbManager struct {
	db *sql.DB
	ctx context.Context
	sm dbStmtManager
}

type dbStmtManager struct {
	querySt *sql.Stmt
	insertSt *sql.Stmt
}


type xormTest struct {
	db *xorm.Engine
}
func NewXormDb() (TestManager,error) {
	m :=&xormTest{}
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, base64dbpassword, addr, dbname)

	db,err := xorm.NewEngine("mysql",conn);
	if err != nil {
		return m,err
	}
	m.db = db
	return m,nil
}


func NewTestDbManager() (TestManager,error) {
	m := &TestDbManager{
		ctx:context.Background(),
	}
	err := m.reset()
	return m,err
}

func (t *TestDbManager)reset()error {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, base64dbpassword, addr, dbname)
	db,err := sql.Open("mysql",conn)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	err = db.Ping()
	if err != nil {
		return err
	}
	t.db = db
	return t.resetStmt()
}
func (t *TestDbManager)resetStmt() error{
	testSt,err := t.db.Prepare(qryBase)
	if err != nil {
		return err
	}
	if t.sm.querySt != nil {
		t.sm.querySt.Close()
	}
	t.sm.querySt = testSt
	insertSt,err := t.db.Prepare(insertSql)
	if err != nil {
		return err
	}
	if t.sm.insertSt != nil {
		t.sm.insertSt.Close()
	}
	t.sm.insertSt = insertSt
	return nil
}

type data struct {
	Id int64 `xorm:"id"`
}


func (t *TestDbManager) QueryTest(name string, gender string) error{
	ctx,_ := context.WithTimeout(t.ctx,time.Duration(10)*time.Second)

	rows,err := t.sm.querySt.QueryContext(ctx)
	if err != nil {
		return err
	}
	for rows.Next() {
		var m data
		t.parseTest(rows,&m)
		//fmt.Println(m.Id)
	}
	return nil
}
func (t *TestDbManager) parseTest(sc SqlScan,d *data) error {
	return sc.Scan(&d.Id)
}


func (x *xormTest) QueryTest(name string, gender string) error{
	lst := make([]data,0)
	err := x.db.SQL(fmt.Sprintf("select id from test")).Find(&lst)
	if err != nil {
		return err
	}
/*	for _,d :=range lst {
		fmt.Println(d.Id)
	}*/
	return nil
}

func (x *xormTest)Insert(id int64,name,gender string) error {
	r,err := x.db.Exec(fmt.Sprintf("insert into test(id,name,gender) value('%d','%s','%s')",id,name,gender))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(r)
	return nil
}

func (t *TestDbManager) Insert(id int64,name,gender string) error {
	r,err:= t.sm.insertSt.ExecContext(t.ctx,id,name,gender)
	if err != nil {
		return err
	}
	fmt.Println(r.LastInsertId())
	return nil
}

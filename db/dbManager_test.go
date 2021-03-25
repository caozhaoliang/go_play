package db

import (
	"fmt"
	"sync"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

var (
	onceTest sync.Once
	test TestManager
	xormInstance TestManager
)

func initTest(){
	onceTest.Do(func() {
		var err error
		test,err = NewTestDbManager()
		if err != nil {
			panic(err)
		}
		xormInstance,err = NewXormDb()
		if err != nil {
			panic(err)
		}
	})
}


func TestTestDbManager_InsertX(t *testing.T) {
	initTest()
	for i:=10000;i <100000; i++{
		xormInstance.Insert(int64(i),fmt.Sprintf("name_%d",i),fmt.Sprintf("%d",i))
	}
}

func TestTestDbManager_Insert2(t *testing.T) {
	initTest()
	for i:=10;i <10000; i++{
		test.Insert(int64(i),fmt.Sprintf("name_%d",i),fmt.Sprintf("%d",i))
	}
}


func TestTestDbManager_QueryTestX(t *testing.T) {
	initTest()

	for i:=0;i < 100000; i++ {
		xormInstance.QueryTest("name", "M")
	}
}


func TestTestDbManager_QueryTest2(t *testing.T) {
	initTest()
	for i:=0;i <100000; i++{
		test.QueryTest("name","M")
	}
}

func BenchmarkXormTest_QueryTest(b *testing.B) {
	initTest()
	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		xormInstance.QueryTest("name","M")
	}
	b.StopTimer()
}
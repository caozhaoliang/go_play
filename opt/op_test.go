package opt

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

type T1 struct {
	String func() string
}

func (T1) Error() string {
	return "T1.Error"
}

type T2 struct {
	Error func() string
}

func (T2) String() string {
	return "T2.String"
}

var t1 = T1{String: func() string { return "T1.String" }}
var t2 = T2{Error: func() string { return "T2.Error" }}

func TestErrString(t *testing.T) {
	fmt.Println(t1.Error())
	fmt.Println(t1.String())

	fmt.Println(t2.Error())
	fmt.Println(t2.String())

	fmt.Println(t1)
	fmt.Println(t2)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func checkFileIsExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}
func TestWrite(t *testing.T) {
	var wireteString = "测试n"
	/*var fileName = "./output1.txt"
	var f *os.File
	var err error
	if checkFileIsExist(fileName) {
		f, err = os.OpenFile(fileName, os.O_APPEND, 0660)
		fmt.Println("文件不存在")
	} else {
		f, err = os.Create(fileName)
	}
	check(err)
	n, err := io.WriteString(f, wireteString) // 1. io.WriteString
	check(err)
	fmt.Printf("写入%d个字节", n)*/
	// ====================
	/*var d1 = []byte(wireteString)
	err2 := ioutil.WriteFile("./output2.txt", d1, 0666) //写入文件(字节数组)	// 2. ioutil.WriteFile
	check(err2)*/
	// ====================== 3. file
	/*	var d1 = []byte(wireteString)
		f, err3 := os.Create("./output3.txt") //创建文件
		check(err3)
		defer f.Close()
		n2, err3 := f.Write(d1) //写入文件(字节数组)
		check(err3)
		fmt.Printf("写入 %d 个字节n", n2)
		n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
		fmt.Printf("写入 %d 个字节n", n3)
		f.Sync()*/
	// ==================== 4. bufio
	f, err3 := os.Create("./output4.txt") //创建文件
	check(err3)
	w := bufio.NewWriter(f)
	n4, err3 := w.WriteString(wireteString)
	check(err3)
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()

}

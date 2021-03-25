package demoArchive

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func writeBuf() bytes.Buffer {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)
	w.WriteString("hello,")
	w.WriteRune('W')
	w.WriteByte('o')
	w.Write([]byte("rld!"))
	w.Flush()
	return buf
}
func readBuf(path string) {
	var buf bytes.Buffer
	bf,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var r = bufio.NewReader(bytes.NewReader(bf))
	if s,err := r.ReadString('W');err!= nil {
		log.Fatal(err)
	}else {
		fmt.Println("--",s)
	}
	fmt.Println(r.Buffered())
	r.WriteTo(&buf)
	fmt.Printf("%s\n",buf.Bytes())
}

func TestScanner(t *testing.T) {
	input :="foo bar    baz"
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err();err != nil {
		t.Fatal(err)
	}
}

func TestScannerCustom (t *testing.T) {
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	spliter := func(data []byte,atEOF bool) (advance int,token []byte,err error) {
		advance,token,err = bufio.ScanWords(data,atEOF)
		if err == nil && token != nil {
			_,err = strconv.ParseInt(string(token),10,32)
		}
		return
	}
	scanner.Split(spliter)
	for scanner.Scan() {
		fmt.Printf("%s\n",scanner.Text())
	}
	if err := scanner.Err(); err!= nil {
		log.Fatal("Invalid input:",err)
	}
}

func TestScannerSplitWithComma(t *testing.T) {
	const intput  = "1,2,3,4"
	scanner := bufio.NewScanner(strings.NewReader(intput))
	onComma := func(data []byte,atEOF bool)(advance int,token []byte,err error) {
		for i:= 0;i < len(data);i++ {
			if data[i] == ',' {
				return i+1,data[:i],nil
			}
		}
		return 0,data,bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("%s\n",scanner.Text())
	}
	if err := scanner.Err();err!= nil {
		log.Fatal(err)
	}
}


const (
	filePathBuf = "testdata/bufio"
)

func TestBuf(t *testing.T) {
	buf:=writeBuf()
	fmt.Printf("%s\n",buf.Bytes())
	if err := ioutil.WriteFile(filePathBuf,buf.Bytes(),os.ModePerm);err != nil {
		log.Fatal(err)
	}
	readBuf(filePathBuf)
}

func TestBytes(t *testing.T) {
	var s = []byte("ba&cd,gh&a cc|dg&a")
	l := bytes.SplitN(s,[]byte("a "),-1)
	for _,d := range l {
		fmt.Printf("%s\n",string(d))
	}
}

func TestHexDump(t *testing.T) {
	content := []byte("Go is an open source programming language.")
	str := hex.Dump(content)
	fmt.Printf("%s\n",str)

	var buf bytes.Buffer
	dumper := hex.Dumper(&buf)
	defer dumper.Close()
	if _,err := dumper.Write(content);err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n",buf.Bytes())
}

func Test_regexp(t *testing.T) {
	//digit()
	//stringRe()
	//emailRe()
	domainRe()
}

func domainRe(){
	// ^[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(/.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+/.?
	data := "yxzq-g.ac"
	pattern:= `[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,10})?`
	ok,err := regexp.MatchString(pattern,data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}


func emailRe() {
	// ^\w+([-+.]\w+)\*@\w+([-.]\w+)\*\.\w+([-.]\w+)\*$
	data := "107cn@gmail.com.cn"
	pattern:= `^\w+([-+.]?\w+)@(\w+[-.]\w+)?\w+([-.]\w+)$`
	ok,err := regexp.MatchString(pattern,data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}

func stringRe() {
	//^[A-Za-z0-9]+$
	/*data := "bacdZZ901Z"
	pattern := `^[A-Za-z0-9]+$`*/
	// ^.{3,20}$
	/*data := "bacdZZ901Z"
	pattern := `^.{1,10}$`*/

	//^\w+$ 由数字、26个英文字母或者下划线组成的字符串
	data := "_bacdZZ901Z"
	pattern := `^\w+$`
	ok,err := regexp.MatchString(pattern,data)//Match(pattern,[]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}

func digit() {
	// ^[0-9]\*$
	//data := "01920*"
	//pattern := `^[0-9]+\*$`

	// ^\d{n}$
	// data := "11010"
	// pattern := `^\d{5}$`

	// ^\d{n,}$
	//data := "11010"
	//pattern := `^\d{4,8}$`

	// ^(0|[1-9][0-9]\*)$
	/*data := "2000*"
	pattern := `^([1-9][0-9]+\*)$`*/

	// ^([1-9][0-9]\*)+(\.[0-9]{1,2})?$
/*	data := "210101"
	pattern := `^([1-9][0-9]+)+(\.[0-9]{1,2})?$`*/

	// ^(\-)?\d+(\.\d{1,2})$   // 带1-2位小数的正数或负数
/*	data := "-7.87"
	pattern := `^(\-)?\d+(\.\d{1,2})$`
*/
	// ^-[1-9]\d\*$
	/*data := "-787"
	pattern := `^-[1-9](\d+)$`*/
	// ^\d+(\.\d+)?$
	data := "-787.98901"
	pattern := `^(-?)\d+(\.\d+)$`
	ok,err := regexp.MatchString(pattern,data)//Match(pattern,[]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}
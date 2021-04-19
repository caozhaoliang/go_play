package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fogleman/gg"
	goredislib "github.com/go-redis/redis"
	"github.com/go-redsync/redsync"
	"github.com/go-redsync/redsync/redis/goredis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kofoworola/godate"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

func getError() error {
	return nil
}

func testErrCause() {
	//e0 := getError()
	e1 := errors.New("error test")
	e2 := errors.Wrap(e1, "a")
	e3 := errors.Wrap(e2, "b")
	es := errors.Cause(e3)
	fmt.Printf("%#v %#v %#v %#v\n", es, e1, e2, e3)
	fmt.Println(es == e1, es == e2)
	fmt.Printf("%v", e3)
}

func timeTest() {
	f1()
}
func f1() {
	now, _ := godate.Parse("2006-01-02", "2021-01-30")
	fmt.Println(now)
	mf := now.Add(1, godate.MONTH)
	fmt.Println(mf)
}

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / GO"
	format                 = "%f / %d / %s"
)

func testScanln() {
	fmt.Println("plz input #")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("from the string we read:", f, i, s)
}

func bufioRead() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("plz input:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(input)
}

func args() {
	who := "czl "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("good morning", who)
}

var NewLineFlag = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	NewLine = "\n"
)

func flagTest() {

	//flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLineFlag {
				s += NewLine
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

func cat(r *bufio.Reader) {
	i := 0
	for {
		i++
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%d:%s", i, buf)
	}
	return
}

func catFile() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from:%s:%s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func playUrl() {
	a := 2
	b := 3
	fmt.Println(a & b)

	//urlTemp := "http://m1-sit.yxzq.com/webapp/stock-king/strategy.html#/manual/detail/10035"
	//fmt.Println(url.QueryEscape(urlTemp))
}
func stringsJoin() {
	/*sil := []string{"a", "b", "c"}
	fmt.Println(strings.Join(sil, ","))*/
	s := "中文ttaa" //"FuckFUCKFack"
	s2 := "中文TT"
	b := strings.Contains(strings.ToLower(s), strings.ToLower(s2))
	fmt.Println(b)
}

func stringsTrim() {
	s1 := ""
	s2 := strings.Split(s1, ",")

	//s2 := strings.TrimSpace(s1) // Trim(s1, " ")
	fmt.Println(len(s1), s2, len(s2))
}
func dateFormat() {
	date := time.Now()
	fmt.Println(date.Format("aaa:20060102"))
}

func mathRound() {
	f := 3.141926
	a1 := math.Round(f)
	a2 := math.Cbrt(f)
	a3 := math.Pow(1.4646436871034738, 3)
	fmt.Println(a1, a2, a3)
}
func rangeNil() {
	type s struct {
		A string
	}
	d := make([]*s, 0)
	fmt.Println(d)
	for _, i := range d {
		fmt.Println(i.A)
	}
}
func mapInsert() {
	mp := make(map[string]int64)
	mp["a"] = 1
	mp["a"] = 2
	fmt.Println(mp)
	mp = make(map[string]int64)
	fmt.Println(mp)
}
func switchDemo() {
	i = 100
	switch i {
	case 1:
		fmt.Println(i)
		//default:
		//fmt.Println("default")
	}
	fmt.Println("a")
}
func testfmt() {
	v := []string{"a"}
	s := fmt.Sprintf("%%%s%%", strings.Join(v, "%"))
	fmt.Println(s)
}

func main1() {
	//testErrCause()
	//timeTest()
	//testScanln()
	//bufioRead()
	//args()
	//flagTest()
	//catFile()
	//playUrl()
	//stringsJoin()
	// stringsTrim()
	//dateFormat()
	//mapInsert()
	//mathRound()
	//switchDemo()
	//play()
	/*	o1 := opt.WithCaching(false)
		o2 := opt.WithTimeout(time.Duration(1)*time.Second)
		opt.Connect("127.0.0.1",o1)
		opt.Connect("127.0.0.1",o1,o2)*/
	//timeTest()
	//testfmt()
	//deleteMap()
	//m1()
	//sort2()
	//playCron()
}

func playCron() {
	c := cron.New()
	err := c.AddFunc("0 */2 * * *", func() {
		fmt.Println("-1-", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		panic(err)
	}
	c.Start()
	select {}
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func m1() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		err := http.ListenAndServe("localhost:8000", nil)
		//log.Fatal()
		fmt.Println(err)
		return
	}
	lissajous(os.Stdout)
}
func sort2() {
	uIds := []string{"a", "b", "c"}
	sort.Sort(sort.Reverse(sort.StringSlice(uIds)))
	fmt.Println(uIds)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

type Hello struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Arr  []string `json:""`
}

func (h *Hello) Say() {
	fmt.Println("Hello World!")
}
func (h Hello) Says() {
	fmt.Println("Hello World!")
}

func playReflect() {
	hello := Hello{}
	t := reflect.TypeOf(hello) // main.Hello
	ty := reflect.TypeOf(&hello).Elem()
	fmt.Println(t, ty)
}

func unSafe() {
	var hello = Hello{}
	s := unsafe.Sizeof(hello.Name) // 字段所占字节数
	fmt.Println(s)
	f := unsafe.Offsetof(hello.Age) // 偏移量
	fmt.Println(f)
	a := unsafe.Alignof(hello) // 结构体对齐方式
	fmt.Println(a)
	/*d := unsafe.Pointer(&hello)
	fmt.Println(d)*/
}

/*给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。*/
func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	start, end := -1, len(nums)
	for i, d := range nums {
		if d < val {
			start = i
		}
		if d > val && end > i {
			end = i
		}
	}
	defer func() {
		fmt.Println(nums, start, end)
	}()
	if start < 0 && end == len(nums) {
		return 0
	} else if start < 0 {
		return len(nums[end:])
	} else if end == len(nums) {
		return len(nums[:start])
	}
	return len(nums[:start]) + len(nums[end:])
}

type SMap struct {
	A  string              `json:"a"`
	Mp map[int]interface{} `json:"mp"`
}

func deleteMap() {
	smp := SMap{Mp: make(map[int]interface{})}
	smp.Mp[1] = "嫦娥"
	smp.A = "a"
	content, err := json.Marshal(&smp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func play() {
	//removeElement([]int{3,2,2,3},3)
	//playReflect()
	// deleteMap()
	date := "2020-01-30 12:00:00"
	t, _ := time.Parse("2006-01-02 15:04:05", date)
	t1 := t.AddDate(0, 1, 0)
	fmt.Println(t1.Format("2006-01-02 15:04:05"))
	//fmt.Printf("%f", -1e10)
}

// intput "2019-12-30"
func GetDateInt64(t string) (int64, error) {
	tim, err := time.ParseInLocation("2006-01-02", t, time.Local)
	if err != nil {
		return 0, err
	}

	date, err := strconv.ParseInt(tim.Format("20060102150405"), 10, 64)
	if err != nil {
		return 0, err
	}

	return date, nil
}

func main2() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int) // 双向通道
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done(): // 避免goroutine 泄露
					return
				case ch <- i:
				}
			}
		}()
		return ch // 被转换为只读通道
	}(ctx)

	for v := range ch { // 只读通道
		fmt.Println(v)
		if v == 5 {
			cancel()
			break
		}
	}
}
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}
func main12() {
	t1 := time.NewTicker(time.Second)
	t5 := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t1.C:
			fmt.Println("t1-", time.Now().Format("2006-01-02 15:04:05"))
		case <-t5.C:
			fmt.Println("t5-", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

// ------
func deferCall() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	// 延时函数会依次 添加到先进后出队列，程序首先运行到panic 函数 并且结束进场，结束进程前延时函数将被执行，
	// 如果延时函数中包含recover函数,那么异常将被处理。
	panic("err")
}
func main3() {
	deferCall()
}

type student struct {
	Name string
	Age  int
}

func paseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu // for range slice 是遍历指针,而mp的value也是存放的指针，那么当stu指向最后一个元素的时候 mp中的value也是指向
	}
	fmt.Println(m)
}
func main4() {
	paseStudent()
}
func main5() {
	runtime.GOMAXPROCS(1) // 设置进一个逻辑进程 GPM模型中的P
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() { // 协程 轻量级线程 保存要执行指令及上下文 GPM模型中的G
			fmt.Println("i: ", i)
			wg.Done()
		}() // 由于只有一个逻辑进程 此时所有的G都放在队列当中 共用变量i 且i=10 并退出循环
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i) // 执行指令的是M-machine 也是真正执行机器指令的
			wg.Done()
		}(i) // 放在队列当中 且单独保存每个协程的变量i
	}

	wg.Wait() // 释放主协程的资源
}

// ------------
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main6() {
	t := Teacher{}
	t.ShowA() // 由于teacher 没有ShowA 方法 这里执行的是 t.people.ShowA()
}
func main7() {
	b := make([]int, 0)
	for i := 0; i < 22; i++ {
		b = append(b, i)
	}
	fmt.Println(b[len(b)-20:])
}

type Kline struct {
	Id      string `json:"id"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Open    string `json:"open"`
	Close   string `json:"close"`
	PClose  string `json:"p_close"`
	StockId string `json:"stock_id"`
	Date    string `json:"date"`
}

func main8() {
	file, err := ioutil.ReadFile("mysq.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lst := make([]*Kline, 0)
	var r = bufio.NewReader(bytes.NewReader(file))
	for {
		if s, err := r.ReadString('\n'); err != nil {
			break
		} else {
			d := strings.Split(s, "\t")
			fmt.Println(d)
			tmp := &Kline{
				Id:     d[0],
				Date:   d[1],
				Open:   d[2],
				High:   d[3],
				Low:    d[4],
				Close:  d[5],
				PClose: strings.Trim(d[6], "\r\n"),
			}
			lst = append(lst, tmp)
		}
	}
	fmt.Println(lst)
}
func main9() {
	s := "%7B\"content\"%3A\"hello\"%7D"
	s1, err := url.QueryUnescape(s)

	fmt.Println(s1, err)
}
func GetRandStr(n int) (randStr string) {
	chars := "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

// 渲染文字
func writeText(dc *gg.Context, text string, x, y float64) {
	xfload := 5 - rand.Float64()*10 + x
	yfload := 5 - rand.Float64()*10 + y

	radians := 40 - rand.Float64()*80
	dc.RotateAbout(gg.Radians(radians), x, y)
	dc.DrawStringAnchored(text, xfload, yfload, 0.2, 0.5)
	dc.RotateAbout(-1*gg.Radians(radians), x, y)
	dc.Stroke()
}

// 随机坐标
func getRandPos(width, height int) (x float64, y float64) {
	x = rand.Float64() * float64(width)
	y = rand.Float64() * float64(height)
	return x, y
}

// 随机颜色
func getRandColor(maxColor int) (r, g, b, a int) {
	r = int(uint8(rand.Intn(maxColor)))
	g = int(uint8(rand.Intn(maxColor)))
	b = int(uint8(rand.Intn(maxColor)))
	a = int(uint8(rand.Intn(255)))
	return r, g, b, a
}

// 随机颜色范围
func getRandColorRange(miniColor, maxColor int) (r, g, b, a int) {
	if miniColor > maxColor {
		miniColor = 0
		maxColor = 255
	}
	r = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	g = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	b = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	a = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	return r, g, b, a
}

type B struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}
type A struct {
	SB *B `json:"sb"`
}

func main10() {
	a := &A{
		SB: &B{Count: 10, List: make([]string, 0)},
	}
	a.SB.List = append(a.SB.List, []string{"1", "2", "1", "2", "1", "2"}...)
	a1 := &A{}
	*a1 = *a
	fmt.Println("%#v", a1)
}

func main11() {
	var array [10]int
	var slice = array[5:6]
	fmt.Println("length of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5], slice[0])
}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}
func main13() { // 13
	var slice []int
	slice = append(slice, 1, 2, 3)
	newSlice := AddElement(slice, 4)
	fmt.Println(&slice[0] == &newSlice[0]) // true
	newSlice = AddElement(newSlice, 5)
	fmt.Println(&slice[0] == &newSlice[0]) // false
}
func main14() { //14
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}

type Server struct {
	ServerName string `key1:"value1" key11:"value11"`
	ServerIP   string `key2:"value2"`
}

func main15() { //15
	s := Server{}
	st := reflect.TypeOf(s)
	field1 := st.Field(0)
	fmt.Printf("key1:%v\n", field1.Tag.Get("key1"))
	fmt.Printf("key11:%v\n", field1.Tag.Get("key11"))
	filed2 := st.Field(1)
	fmt.Printf("key2:%v\n", filed2.Tag.Get("key2"))
}

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 //const声明第0行，即iota==0
	bit1, mask1                          //const声明第1行，即iota==1, 表达式继承上面的语句
	_, _                                 //const声明第2行，即iota==2
	bit3, mask3                          //const声明第3行，即iota==3
)

func main16() {
	fmt.Println(bit0, mask0)
	fmt.Println(bit1, mask1)
	fmt.Println(bit3, mask3)
}

/*func main17() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Hello, World!")
	})
	log.Fatalln(zerodown.ListenAndServe(":8080", router))

}*/
func getRand(x int) int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := rd.Int63()
	return int(v % int64(x))
}

// 从 a b c d 中取两个不相同的值 且取出来的比例为 1：2：3：4
func get() []string {
	lst := []string{"a", "b", "b", "c", "c", "c", "d", "d", "d", "d"}
	var a, b *int
	var z int
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 2; i++ {
		x := rd.Intn(1000)
		x = x % (10 - z)
		fmt.Println(i, x)
		if a == nil {
			a = new(int)
			*a = x
			if x == 0 {
				z = 1
			} else if x < 3 {
				z = 2
			} else if x < 6 {
				z = 3
			} else {
				z = 4
			}
			continue
		}
		b = new(int)
		if z == 1 || (z == 2 && x > 0) || (z == 3 && x > 2) {
			*b = x + z
		} else {
			*b = x
		}

	}
	return []string{lst[*a], lst[*b]}
}

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func getSmallNode(r *LinkNode) (*LinkNode, bool) {
	Val := r.Val
	tmp := r
	swap := false
	for r != nil {
		if r.Val < Val {
			tmp = r
			Val = r.Val
			swap = true
		}
		r = r.Next
	}
	return tmp, swap
}

func swap(r, small *LinkNode) {
	r.Val, small.Val = small.Val, r.Val
}
func printNode(r *LinkNode) {
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
func Sort(r *LinkNode, n int) *LinkNode {
	//printNode(r)
	t := r
	for i := 0; i < n; i++ {
		small, sp := getSmallNode(t)
		if sp {
			swap(t, small)
		}
		t = t.Next
	}
	return r
}
func main17() {
	/*	fmt.Println(get())
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < 2; i++ {
				x := rd.Intn(10)
				fmt.Println(i, x)
			}*/
	l0 := &LinkNode{Val: 3}
	l1 := &LinkNode{Val: 1}
	l2 := &LinkNode{Val: 2}
	l3 := &LinkNode{Val: 0}
	l0.Next, l1.Next, l2.Next, l3.Next = l1, l2, l3, nil
	printNode(l0)
	r := Sort(l0, 4)
	printNode(r)
}

// leetcode
func TwoNum(lst []int, target int) []int {
	m := make(map[int]int)
	for i, v := range lst {
		if k, ok := m[target-v]; ok {
			return []int{i, k}
		}
		m[v] = i
	}
	return []int{}
}

func AddTwoNum(l1, l2 *LinkNode) *LinkNode {
	head := &LinkNode{Val: 0}
	n1, n2, carry, cur := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		cur.Next = &LinkNode{Val: (n1 + n2 + carry) % 10}
		cur = cur.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 3 在一个字符串重寻找没有重复字母的最长子串
func LongestSubString(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 4 找出两个数组的中位数
func findMedianSorted(n1, n2 []int) float64 {
	if len(n1) > len(n2) {
		return findMedianSorted(n2, n1)
	}
	low, high, k, n1Mid, n2Mid := 0, len(n1), (len(n1)+len(n2)+1)>>1, 0, 0
	for low <= high {
		n1Mid = low + (high-low)>>1
		n2Mid = k - n1Mid
		if n1Mid > 0 && n1[n1Mid-1] > n2[n2Mid] {
			high = n1Mid - 1
		} else if n1Mid != len(n1) && n1[n1Mid] < n2[n2Mid-1] {
			low = n1Mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if n1Mid == 0 {
		midLeft = n2[n2Mid-1]
	} else if n2Mid == 0 {
		midLeft = max(n1[n1Mid-1], n2[n2Mid-1])
	} else {
		midLeft = max(n1[n1Mid-1], n2[n2Mid-1])
	}
	if (len(n1)+len(n2))&1 == 1 {
		return float64(midLeft)
	}
	if n1Mid == len(n1) {
		midRight = n2[n2Mid]
	} else if n2Mid == len(n2) {
		midRight = n1[n1Mid]
	} else {
		midRight = min(n1[n1Mid], n2[n2Mid])
	}
	return float64(midLeft+midRight) / 2
}
func main18() {
	/*a := []int{2, 3, 4, 7}
	fmt.Println(TwoNum(a, 5))*/
	// 2
	/*	a, b, c := &LinkNode{Val: 3}, &LinkNode{Val: 1}, &LinkNode{Val: 2}
		a.Next, b.Next, c.Next = b, c, nil
		a1, b1, c1 := &LinkNode{Val: 8}, &LinkNode{Val: 3}, &LinkNode{Val: 8}
		a1.Next, b1.Next, c1.Next = b1, c1, nil
		r := AddTwoNum(a, a1)
		printNode(r)*/
	// 3 找出字符串最大不重复字符串
	/*	s := "abccdefgww"
		fmt.Println(LongestSubString(s))*/
	// 4 找出两个数组的中位数
	a := []int{1, 3, 4, 6}
	b := []int{2, 3, 4, 5, 9}
	r := findMedianSorted(a, b)
	fmt.Println(r)
}

type Student struct {
	Id   int
	Name string
}

func addr(s []int) []int {
	s = append(s, []int{4, 5, 6, 7, 8}...)
	return s
}
func (s *Student) SetName(name string) *Student {
	s.Name = name
	return s
}
func main19() {
	s := &Student{Id: 1, Name: "咖啡色的羊驼"}
	t := reflect.TypeOf(s)
	// v := reflect.ValueOf(s)
	f := make(map[int]int)
	f[1] = 1
	v := reflect.ValueOf(&f)
	v.Elem().SetMapIndex(reflect.ValueOf(1), reflect.ValueOf(10))
	fmt.Println(v)
	// 通过.Kind()来判断对比的值是否是struct类型
	if k := t.Kind(); k == reflect.Ptr {
		fmt.Println("bingo")
	}

	num := 1
	numType := reflect.TypeOf(num)
	if k := numType.Kind(); k == reflect.Int {
		fmt.Println("bingo")
	}
}

func main() {
	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	client := goredislib.NewClient(&goredislib.Options{
		Addr:"",	//
	})
	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutexName := "my-global-mutex"
	mutex := rs.NewMutex(mutexName)
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		panic(err)
	}
	// Do your work that requires the lock.

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
}

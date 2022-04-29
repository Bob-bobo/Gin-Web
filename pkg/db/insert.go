package db

import (
	"database/sql"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//mysql连接信息配置： 用户名、密码、ip、端口、库名、表名
const (
	strUserName  = "root"
	strPassword  = "19980320"
	strIP        = "127.0.0.1"
	strPort      = "3306"
	strDBName    = "myblog_db"
	strTableName = "user_tab"
)

//插入数据量配置
const (
	TOTAL_INSERT_NUM = 3000 //共插入多少行数据
	PER_INSERT_NUM   = 500  //单次向mysql插入多少行数据
	MAX_FAILNUM      = 10   //最大容许插入失败次数
)

//mysql整型范围 供预置数据范围使用
const (
	BIGINT_MIN = -9223372036854775808
	BIGINT_MAX = 9223372036854775807

	INT_MIN = -2147483648
	INT_MAX = 2147483647

	MEDIUMINT_MIN = -8388608
	MEDIUMINT_MAX = 8388607

	SMALLINT_MIN = -32768
	SMALLINT_MAX = 32767

	TINYINT_MIN = -128
	TINYINT_MAX = 127
)

//随机字符串种子
const STRCHAR = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&*+-./:;<=>?@[]^_{|}~"

var CITYS = []string{"ChengDu", "KunMing", "XiAn", "LaSa", "JiNan", "NanJing", "HangZhou", "FuZhou", "GuangZhou", "changsha",
	"HaiKou", "HaErBin", "ChangChun", "ShenYang", "ZhengZhou", "HeFei", "WuHan", "ChongQing", "BeiJing", "ShangHai"}

var FName = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "沈", "韩", "杨", "朱", "秦", "许", "何", "张", "杜", "曹", "谢", "苏", "马", "雷", "范", "唐", "葛"}

var NaNv = []string{"男", "女"}

//随机生成一个整型数据
func MakeRandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

//随机生成一个浮点型数据
func MakeRandFloat(base int64) float32 {
	return rand.Float32() * float32(base)
}

//随机生成一个双精度浮点型数据
func MakeRandDouble(base int64) float64 {
	return rand.Float64() * float64(base)
}

//随机生成一个汉字字符串  入参：字符串长度
func MakeChineseString(length int) string {
	a := make([]rune, length)
	for i := range a {
		a[i] = rune(MakeRandInt(19968, 40869))
	}
	return string(a)
}

//随机生成一个字符串（指定字符种子） 入参：（字符串长度，长度是否随机）
func MakeRandString(length int64, bRegular bool) string {
	var size int64
	if bRegular {
		size = length
	} else {
		size = MakeRandInt(1, length)
	}

	str := make([]byte, size)
	for i := 0; i < int(size); i++ {
		index := MakeRandInt(0, int64(len(STRCHAR)))
		str[i] = STRCHAR[index]
	}
	return string(str)
}

//随机生成一个字符串（任意字符） 入参：字符串长度
func MakeRandString2(length int) string {
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		index := MakeRandInt(0, 127)
		str[i] = byte(index)
	}
	return string(str)
}

//随机生成一个日期类型数据
func MakeRandDate() string {
	year := MakeRandInt(1970, 2021)
	month := MakeRandInt(1, 12)
	var day int64
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = MakeRandInt(1, 31)
	case 4, 6, 9, 11:
		day = MakeRandInt(1, 30)
	case 2:
		day = MakeRandInt(1, 28)
	}

	strDate := strconv.FormatInt(year, 10) + "-" + strconv.FormatInt(month, 10) + "-" + strconv.FormatInt(day, 10)
	return strDate
}

//生成一个递增数据
var iValue int64

//初始化递增数据的值 入参：初始值
func InitIncreaseInt(start int64) {
	iValue = start
}

//生成递增数据 入参：递增量
func MakeIncreaseInt(stage int64) int64 {
	iRet := iValue
	iValue += stage
	return iRet
}

//建立数据库连接
func InitDB() *sql.DB {

	source := strings.Join([]string{strUserName, ":", strPassword, "@tcp(", strIP, ":", strPort, ")/", strDBName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, err := sql.Open("mysql", source)
	if err != nil {
		panic(fmt.Sprintf("Open Kungate Connection:[%s] failed, error is [%v].", source, err))
	}

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)

	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return nil
	}

	fmt.Println("connnect success")
	return DB
}

//执行插入操作
func DoInsert(dbConn *sql.DB) {
	fmt.Printf("begin insert, total num:[%d]\n", TOTAL_INSERT_NUM)

	startTime := time.Now().Unix()
	failnum := 0

	//需要赋值的字段定义
	var id int64
	var phone string
	var username string
	var password string
	var gender string
	var trueName string
	var birthday string
	var email string
	var personalBrief string
	var avatarImgUrl string
	var recentlyLanded string

	//初始化递增数据
	InitIncreaseInt(7510)

	//拼接insert语句
	var strInsert string = "insert into " + strTableName + " values"
	var InsertBuf string = strInsert
	personalBrief = ""
	avatarImgUrl = ""
	recentlyLanded = ""
	password = util.HashAndSalt("test2022@")
	for i := 1; i <= TOTAL_INSERT_NUM; i++ {

		//为各个字段制造随机数据
		id = MakeIncreaseInt(1)
		phone = fmt.Sprintf("%s%d", "1", MakeRandInt(1000000000, 9999999999))
		username = MakeChineseString(4)
		gender = NaNv[int(MakeRandInt(0, 2))]
		trueName = fmt.Sprintf("%s%s", FName[int(MakeRandInt(0, 26))], MakeChineseString(2))
		birthday = MakeRandDate()
		email = fmt.Sprintf("%d@qq.com", MakeRandInt(70000000, 79999999))

		//拼接insert语句中的值
		InsertBuf += fmt.Sprintf(" (%d,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s' ) ",
			id, phone, username, password, gender, trueName, birthday, email, personalBrief, avatarImgUrl, recentlyLanded)

		//若达到单次插入行数 执行插入
		if i%PER_INSERT_NUM == 0 {
			InsertBuf += ";"

			_, err := dbConn.Exec(InsertBuf)
			if err != nil {
				fmt.Println(err)
				fmt.Println(InsertBuf)

				failnum++
				if failnum > MAX_FAILNUM {
					return
				}

				InsertBuf = strInsert
				continue
			}

			//重新初始化insert插入语句
			InsertBuf = strInsert
			CurTime := time.Now().Unix()
			fmt.Printf("complete:[%d]   failnum:[%d]  consume:[%ds]\n", i, failnum*PER_INSERT_NUM, CurTime-startTime)
		} else {
			InsertBuf += ","
		}
	}

	endTime := time.Now().Unix()
	fmt.Printf("finished: totalnum:[%d]  consum:[%d]s\n", TOTAL_INSERT_NUM, endTime-startTime)
}

//通道方法
func printHello(ch chan int) {
	fmt.Println("hello from printHello")
	//send a value on channel
	ch <- 2
}

func main() {
	//var db *gorm.DB
	//var actuser models.ActivityUserTab
	//err := db.Where("user_id = ? and activity_id = ?", 1, 4).First(&actuser).Error
	//if err != nil {
	//	fmt.Println("没有")
	//	return
	//}
	//fmt.Println("存在")
	//var in = "2010-2-2"
	//on := "2007-6-5"
	//if in > on {
	//	fmt.Println("in > on")
	//} else {
	//	fmt.Println("in < on")
	//}
	//in = fmt.Sprintf("%s love %s", "cyq", in)
	//fmt.Printf(in)
	//sql := "dzb"
	//sql += "cyqqqqqqq"
	//fmt.Println(sql)

	//向通道传递方法运行，也就是确保小于线程的routine
	//能够顺序在main结束前通知到，并执行routine的方法
	ch := make(chan int)

	go func() {
		fmt.Println("hello inline")
		//发送一个值给通道
		ch <- 1
	}()

	go printHello(ch)
	fmt.Println("hello from main")

	i := <-ch
	fmt.Println("Recieved", i)

	<-ch
}

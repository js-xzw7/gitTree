### 类型转换

#### int 转其他类型 
- int转string
```
s := strconv.Itoa(i)
等价于s := strconv.FormatInt(int64(i), 10)
```
- int 转 float

```
var money int
money = 1
fmt.Println(float64(money) / float64(100))
```

- (2)int64转string
```
i := int64(123)
s := strconv.FormatInt(i, 10)
```
> 第二个参数为基数，可选2~36 注：对于无符号整形，可以使用FormatUint(i uint64, base int)
- (3)string转int
```
i, err := strconv.Atoi(s)
```

- string转int64
```
i, err := strconv.ParseInt(s, 10, 64)
```

> 第二个参数为基数（2~36），第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，分别对应 int, int8, int16, int32和int64
  
- (5)float相关
> float转string：
  
```
v := 3.1415926535
s1 := strconv.FormatFloat(v, 'E', -1, 32)//float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
```
>string转float：
```
s := "3.1415926535"
v1, err := strconv.ParseFloat(v, 32)
v2, err := strconv.ParseFloat(v, 64)
``` 

- (5)decime 转 string
```
s := decimalObjr.String()
```

#### 汇总
```
//string到int
int,err:=strconv.Atoi(string)
//string到int64
int64, err := strconv.ParseInt(string, 10, 64)
//int到string
string:=strconv.Itoa(int)
//int64到string
string:=strconv.FormatInt(int64,10)
//string到float32(float64)
float,err := strconv.ParseFloat(string,32/64)
//float到string
string := strconv.FormatFloat(float32, 'E', -1, 32)
string := strconv.FormatFloat(float64, 'E', -1, 64)
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
```

### go的数学运算
- 安装命令
```
go get github.com/shopspring/decimal
```
####案例
```
初始化一个变量
d0 := decimal.NewFromFloat(0)
设置精度 为三位 四舍五入的精度
decimal.DivisionPrecision = 3 
```
- 加法 Add
```	
2.1 + 3 float和int相加
var num1 float64 = 2.1
var num2 int = 3
d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "5.1"

2.1 + 3.1 (float和float相加)
var num1 float64 = 2.1
var num2 float64 = 3.1
d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
fmt.Println(d1) // output: "5.2"

2 + 3 (int和int相加 可以直接相加) 
var num1 float64 = 2
var num2 float64 = 3
（d1 = num1+num2）
d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
fmt.Println(d1) // output: "5"

```
- 减法 Sub

```	
3.1 - 2 float和int相减
var num1 float64 = 3.1
var num2 int = 2
d1 := decimal.NewFromFloat(num1).Sub(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "1.1"

2.1 - 3.1 (float和float相减)
var num1 float64 = 2.1
var num2 float64 = 3.1
d1 := decimal.NewFromFloat(num1).Sub(decimal.NewFromFloat(num2))
fmt.Println(d1) // output: "-1"

2 - 3 (int和int相减)
var num1 int = 2
var num2 int = 3
(d1 = num1 - num2)
d1 := decimal.NewFromFloat(float64(num1)).Sub(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "-1"

```

- 乘法 Mul

```
	
3.1 * 2 float和int相乘
var num1 float64 = 3.1
var num2 int = 2
d1 := decimal.NewFromFloat(num1).Mul(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "6.2"

2.1 * 3.1 (float和float相乘)
var num1 float64 = 2.1
var num2 float64 = 3.1
d1 := decimal.NewFromFloat(num1).Mul(decimal.NewFromFloat(num2))
fmt.Println(d1) // output: "6.51"

2 * 3 (int和int相乘)
var num1 int = 2
var num2 int = 3
(d1 = num1 * num2)
d1 := decimal.NewFromFloat(float64(num1)).Mul(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "6"

```

除法 Div

```

int 和 int 相除
2 ➗ 3 = 0.6666666666666667 
var num1 int = 2
var num2 int = 3
d1 := decimal.NewFromFloat(float64(num1)).Div(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "0.6666666666666667"

float64 和 int 相除
var num1 float64 = 2.1
var num2 int = 3
d1 := decimal.NewFromFloat(num1).Div(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "0.7"

float64 和 float64 相除
2.1 ➗ 0.3 = 7
var num1 float64 = 2.1
var num2 float64 = 0.3
d2 := decimal.NewFromFloat(num1).Div(decimal.NewFromFloat(num2))
fmt.Println(d2) // output: "7"

int 和 int 相除 并保持3位精度
2 ➗ 3 = 0.667 
var num1 int = 2
var num2 int = 3
decimal.DivisionPrecision = 3
d1 := decimal.NewFromFloat(float64(num1)).Div(decimal.NewFromFloat(float64(num2)))
fmt.Println(d1) // output: "0.667"

```
- 项目例子
```

var betgolds, etstones, stoneuse2, users, rounds, tickets, betgold, betstone, redpacks int
d1 := decimal.NewFromFloat(0) //累加变量初始化值
for _, v := range info {
	betgold = v["robot_golds"].Intval + v["systax"].Intval
	betgolds = betgolds + betgold
	betstone = v["stonetax"].Intval + v["stoneuse"].Intval
	etstones = etstones + betstone
	stoneuse2 = stoneuse2 + v["stoneuse2"].Intval
	users = users + v["totalmembers"].Intval
	redpacks = redpacks + v["gift_redpack"].Intval
	rounds = rounds + v["rounds"].Intval
	tickets = tickets + v["gift_ticket"].Intval
	decimal.DivisionPrecision = 2
	d3 := decimal.NewFromFloat(float64(betgold)).Div(decimal.NewFromFloat(10000))
	d4 := decimal.NewFromFloat(float64(betstone)).Mul(decimal.NewFromFloat(0.22))
	d5 := decimal.NewFromFloat(float64(v["gift_ticket"].Intval)).Mul(decimal.NewFromFloat(0.001))
	d1 = d1.Add(d3).Add(d4).Sub(decimal.NewFromFloat(float64(v["gift_redpack"].Intval))).Sub(d5)
}
//for外输入d1
logger.info(d1)

```
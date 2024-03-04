package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sort"
	"time"
)

const (
	MAXBj  = 3.9
	MAXBs  = 7.8
	MAXCn  = 6.5
	MAXDgj = 5.8
	MAXGj  = 19
	MAXJt  = 23
	MAXDsm = 5.8
	MAXSm  = 299
	MAXDfy = 7.3
	MAXFy  = 23
	Bj     = 3.3
	Bs     = 6.6
	Cn     = 5.5
	Dgj    = 5.0
	Gj     = 17
	Jt     = 20
	Dsm    = 5.0
	Sm     = 254
	Dfy    = 6.2
	Fy     = 20
	MINBj  = 2.7
	MINBs  = 5.4
	MINCn  = 4.9
	MINDgj = 4.1
	MINGj  = 14
	MINJt  = 16
	MINMsm = 4.1
	MINSm  = 209
	MINDfy = 5.1
	MINFy  = 16
)

//var Date []EntryValue

//	func InItDate() {
//		Date = []EntryValue{
//			{BJ: Bj, BS: Bs, CN: Cn, DGj: Dgj},
//			{BJ: Bj, BS: Bs, CN: MAXCn, DGj: MAXDgj},
//		}
//	}
type res1 [][]int

func (m res1) Len() int {
	return len(m)
}
func (m res1) Less(i, j int) bool {
	// 在这里定义自定义的排序规则
	// 这里我们以每行的第一个元素进行升序排序
	if m[i][0] == m[j][0] {
		if m[i][1] == m[j][1] {
			if m[i][2] == m[j][2] {
				if m[i][3] == m[j][3] {

				} else {
					return m[i][3] < m[j][3]
				}
			} else {
				return m[i][2] < m[j][2]
			}
		} else {
			return m[i][1] < m[j][1]

		}
	}
	return m[i][0] < m[j][0]
}
func (m res1) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 210
// 840
// 3360
// 13440
// 53760
// 215040
// 没有小词条 35840
// 羽毛:21509
// 差羽毛:17908
// 好羽毛:5978
// 花:21509
// 差花:3580
// 好花:5978
// 沙漏:35838
// 差沙漏:1191
// 好沙漏:5978
// 杯子:35668
// 差杯子:1191
// 好杯子:5808
// 头冠:35838
// 差头冠:1191 三个小词条
// 好头冠:5978 没有小词条
func ain() {
	UntilRedis()
	initDB()
	var a1 []MY
	var cc MY
	for i := 0; i < 215040; i++ {
		fmt.Println(i)
		//aaa := FindUserPhotoByUserId(i)
		//bb := Add(aaa)
		//cc := MY{
		//	A: aaa,
		//	B: bb,
		//}
		key := fmt.Sprintf("v5:%d", i)
		//fmt.Println(cc)
		//j, _ := json.Marshal(cc)
		//SetRedis(key, j, 0)
		r := GetRedis(key)
		_ = json.Unmarshal([]byte(r), &cc)
		a1 = append(a1, cc)
	}
	fmt.Println(len(a1))
	ii := 0
	for _, my := range a1 {
		if my.A.FY == 0 && my.A.GJ == 0 && my.A.SM == 0 {

		} else if my.A.FY != 0 && my.A.GJ == 0 && my.A.SM == 0 {

		} else {
			key := fmt.Sprintf("普通杯子:%d", ii)
			j, _ := json.Marshal(my)
			SetRedis(key, j, 0)
			ii += 1
			fmt.Println(key)
		}
	}
}

func Add(s EntryValue) (b []int) {
	if s.BJ != 0 {
		b = append(b, 0)
	}
	if s.BS != 0 {
		b = append(b, 1)
	}
	if s.CN != 0 {
		b = append(b, 2)
	}
	if s.DGj != 0 {
		b = append(b, 3)
	}
	if s.GJ != 0 {
		b = append(b, 4)
	}
	if s.JT != 0 {
		b = append(b, 5)
	}
	if s.DSM != 0 {
		b = append(b, 6)
	}
	if s.SM != 0 {
		b = append(b, 7)
	}
	if s.DFY != 0 {
		b = append(b, 8)
	}
	if s.FY != 0 {
		b = append(b, 9)
	}
	sort.Ints(b)
	return b
}

var ConnectRedis *redis.Client

func UntilRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.20.21:6379", // Redis服务器地址和端口
		Password: "root",               // Redis密码，如果没有设置密码则为空
		DB:       0,                    // Redis数据库索引，默认为0
	})
	ConnectRedis = rdb
}
func SetRedis(key string, value any, mytime time.Duration) error {
	//设置键值对：使用Set方法设置键为"key"，值为"value"的数据，并设置过期时间为0
	err := ConnectRedis.Set(context.Background(), key, value, mytime).Err()
	return err
}
func GetRedis(key string) string {
	//获取键的值：使用Get方法获取键为"key"的值，并根据返回结果进行判断。
	value, err := ConnectRedis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		//fmt.Println("key does not exist")
		return ""
	} else {
		//fmt.Println("key:", value)
	}
	return value
}

// mongodb
func FindUserPhotoByUserId(id int) EntryValue {

	//指定操作的数据库名以及要操作的数据集合名
	c := client.Database("vote").Collection("v5")
	//获取指定字段的,指定数值的,所有符合要求的文档
	//c2, _ := c.Find(context.Background(), bson.D{})
	c2, _ := c.Find(context.Background(), bson.D{{"vys", id}})
	defer c2.Close(context.Background())
	var str EntryValue
	for c2.Next(context.Background()) {
		var userphoto EntryValue
		//数据绑定到结构体上
		err := c2.Decode(&userphoto)
		if err != nil {
			fmt.Println(err)
			return EntryValue{}
		}
		str = userphoto
		//获取整个数据
		//fmt.Printf("result: %v\n", result)
		////获取数据的数据部分
		//fmt.Printf("result.Map(): %v\n", result.Map())
		////获取数据的数据部分的指定字段
		//fmt.Printf("result.Map()[\"name\"]: %v\n", result.Map()["path"])
		//mystr := result.Map()["path"].(string)
		//str = strings.Fields(mystr)
		//fmt.Println(str)
	}
	//fmt.Printf("str:%s\n", str)
	if err := c2.Err(); err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return EntryValue{}
	}
	return str
}

// initDB()
//
//	var a1 []MY
//	var cc MY
//	for i := 0; i < 35838; i++ {
//		fmt.Println(i)
//		//aaa := FindUserPhotoByUserId(i)
//		//bb := Add(aaa)
//		//cc := MY{
//		//	A: aaa,
//		//	B: bb,
//		//}
//		//SetRedis(key, j, 0)
//		r := GetRedis(key)
//		_ = json.Unmarshal([]byte(r), &cc)
//		a1 = append(a1, cc)
//	}
//
// 获得圣遗物胚子
func InIt() (rrr []MY) {
	var res res1
	var ress res1
	//获取胚子属性种类
	for i := 0; i < 10; i++ {
		for i1 := 1; i1 < 10; i1++ {
			for i2 := 2; i2 < 10; i2++ {
				for i3 := 3; i3 < 10; i3++ {
					if i == i1 || i == i2 || i == i3 || i1 == i2 || i1 == i3 || i2 == i3 {
					} else {
						b := []int{i, i1, i2, i3}
						sort.Ints(b)
						res = append(res, b)
					}
				}
			}
		}
	}
	sort.Sort(res)
	for i := 1; i < len(res); i++ {
		if i == 1 {
			ress = append(ress, res[0])
		}
		if res[i][0] == res[i-1][0] && res[i][1] == res[i-1][1] && res[i][2] == res[i-1][2] && res[i][3] == res[i-1][3] {
		} else {
			ress = append(ress, res[i])
		}

	}
	//获得圣遗物胚子
	for _, ints := range ress {
		var vv MY
		var v EntryValue
		for _, i2 := range ints {
			switch i2 {
			case 0:
				v.BJ = Bj
			case 1:
				v.BS = Bs
			case 2:
				v.CN = Cn
			case 3:
				v.DGj = Dgj
			case 4:
				v.GJ = Gj
			case 5:
				v.JT = Jt
			case 6:
				v.DSM = Dsm
			case 7:
				v.SM = Sm
			case 8:
				v.DFY = Dfy
			case 9:
				v.FY = Fy
			}
		}
		vv = MY{
			A: v,
			B: ints,
		}
		rrr = append(rrr, vv)
	}
	return rrr
}

// 强化第一次
func Strong(rrr00, rrr01, rrr02, rrr03 []MY) (rrr1 []MY) {
	//对第一个属性强化
	for _, value := range rrr00 {
		for i, component := range value.B {
			if i == 0 {
				switch component {
				case 0:
					value.A.BJ += Bj
				case 1:
					value.A.BS += Bs

				case 2:
					value.A.CN += Cn

				case 3:
					value.A.DGj += Dgj

				case 4:
					value.A.GJ += Gj

				case 5:
					value.A.JT += Jt

				case 6:
					value.A.DSM += Dsm

				case 7:
					value.A.SM += Sm

				case 8:
					value.A.DFY += Dfy

				case 9:
					value.A.FY += Fy

				}
			}
		}
		rrr1 = append(rrr1, value)
	}
	//对第二个属性强化
	for _, value := range rrr01 {
		for i, component := range value.B {
			if i == 1 {
				switch component {
				case 0:
					value.A.BJ += Bj
				case 1:
					value.A.BS += Bs

				case 2:
					value.A.CN += Cn

				case 3:
					value.A.DGj += Dgj

				case 4:
					value.A.GJ += Gj

				case 5:
					value.A.JT += Jt

				case 6:
					value.A.DSM += Dsm

				case 7:
					value.A.SM += Sm

				case 8:
					value.A.DFY += Dfy

				case 9:
					value.A.FY += Fy

				}
			}
		}
		rrr1 = append(rrr1, value)
	}
	//对第三个属性强化
	for _, value := range rrr02 {
		for i, component := range value.B {
			if i == 2 {
				switch component {
				case 0:
					value.A.BJ += Bj
				case 1:
					value.A.BS += Bs

				case 2:
					value.A.CN += Cn

				case 3:
					value.A.DGj += Dgj

				case 4:
					value.A.GJ += Gj

				case 5:
					value.A.JT += Jt

				case 6:
					value.A.DSM += Dsm

				case 7:
					value.A.SM += Sm

				case 8:
					value.A.DFY += Dfy

				case 9:
					value.A.FY += Fy

				}
			}
		}
		rrr1 = append(rrr1, value)
	}
	//对第四个属性强化
	for _, value := range rrr03 {
		for i, component := range value.B {
			if i == 3 {
				switch component {
				case 0:
					value.A.BJ += Bj
				case 1:
					value.A.BS += Bs

				case 2:
					value.A.CN += Cn

				case 3:
					value.A.DGj += Dgj

				case 4:
					value.A.GJ += Gj

				case 5:
					value.A.JT += Jt

				case 6:
					value.A.DSM += Dsm

				case 7:
					value.A.SM += Sm

				case 8:
					value.A.DFY += Dfy

				case 9:
					value.A.FY += Fy

				}
			}
		}
		rrr1 = append(rrr1, value)
	}
	return rrr1
}
func FindDate() []MY {
	var a1 []MY
	var cc MY
	for i := 0; i < 35838; i++ {
		//fmt.Println(i)
		key := fmt.Sprintf("杯子:%d", i)
		r := GetRedis(key)
		_ = json.Unmarshal([]byte(r), &cc)
		a1 = append(a1, cc)
	}
	return a1
}
func Count(e EntryValue, a []MY) float64 {
	var res float64
	var nnn EntryValueComponents
	var aa []EntryValueComponents
	for _, my := range a {
		a1 := AddComponent(my.A)
		aa = append(aa, a1)
	}
	r := AddComponent(e)
	for i, components := range aa {
		re := AngleBetween(components, r)
		if i == 0 {
			res = re
			nnn = components
		}
		if res > re {
			res = re
			nnn = components
		}
	}
	fmt.Println(nnn)
	return res
}

var client *mongo.Client

// 连接mongodb数据库
func initDB() (err error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://root:root@192.168.20.21:27017")
	// 连接到MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}
func close() {
	client.Disconnect(context.TODO())
}

// 插入单条数据
func insertOne(s EntryValue, atr string) {
	//获取数据库连接驱动,client

	//指定连接的库以及连接的表(集合)
	c := client.Database("vote").Collection(atr)
	//插入数据
	ior, err := c.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
}

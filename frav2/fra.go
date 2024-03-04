package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
	"strconv"
)

//const (
//	Bj  = 3.9 * 5
//	Bs  = 7.8 * 5
//	Cn  = 6.5 * 5
//	Dgj = 5.8 * 5
//	Gj  = 19 * 5
//	Jt  = 23 * 5
//	Dsm = 5.8 * 5
//	Sm  = 299 * 5
//	Dfy = 7.3 * 5
//	Fy  = 23 * 5
//)

var (
	MaxBj  = EntryValue{BJ: Bj}
	MaxBS  = EntryValue{BS: Bs}
	MaxCN  = EntryValue{CN: Cn}
	MaxDGJ = EntryValue{DGj: Dgj}
	MaxGJ  = EntryValue{GJ: Gj}
	MaxJT  = EntryValue{JT: Jt}
	MaxDSM = EntryValue{DSM: Dsm}
	MaxSM  = EntryValue{SM: Sm}
	MaxDFY = EntryValue{DFY: Dfy}
	MaxFY  = EntryValue{FY: Fy}
)
var ()

type Entry struct {
	Key   EntryKey
	Value EntryValue
}
type EntryKey struct {
	//套装名称
	Name string
	//主词条
	Main string
	//部位名称
	Asse string
}
type EntryValue struct {
	//值
	BJ  float64
	BS  float64
	VYS int
	CN  float64
	DGj float64
	GJ  float64
	JT  float64
	DSM float64
	SM  float64
	DFY float64
	FY  float64
}

func AddComponent(e EntryValue) EntryValueComponents {
	var a []float64
	a = append(a, e.BJ)
	a = append(a, e.BS)
	a = append(a, e.CN)
	a = append(a, e.DGj)
	a = append(a, e.GJ)
	a = append(a, e.JT)
	a = append(a, e.DSM)
	a = append(a, e.SM)
	a = append(a, e.DFY)
	a = append(a, e.FY)
	ee := EntryValueComponents{Components: a}
	return ee
}

type EntryValueComponents struct {
	Components []float64
}
type MY struct {
	A EntryValue
	B []int
}

// 角色结构体
type Role struct {
	RoleName string
	//	沙漏
	RoleHourglass []string
	//  杯子
	RoleCup []string
	//  头冠
	RolePileum []string
}
type Res struct {
	Name   string
	Source string
}

var MyEntry Entry

// 计算向量的模（长度）
func (v EntryValueComponents) Magnitude() float64 {
	sumOfSquares := 0.0
	for _, component := range v.Components {
		sumOfSquares += component * component
	}
	return math.Sqrt(sumOfSquares)
}

// 计算向量之间的点积
func DotProduct(v1, v2 EntryValueComponents) float64 {
	if len(v1.Components) != len(v2.Components) {
		panic("向量维度不匹配")
	}

	dotProduct := 0.0
	for i := 0; i < len(v1.Components); i++ {
		dotProduct += v1.Components[i] * v2.Components[i]
	}
	return dotProduct
}

// 计算两个向量之间的夹角（以弧度为单位）
func AngleBetween(v1, v2 EntryValueComponents) float64 {
	dotProduct := DotProduct(v1, v2)
	magnitudeProduct := v1.Magnitude() * v2.Magnitude()
	return math.Acos(dotProduct / magnitudeProduct)
}

//func main() {
//	var (
//		a EntryValue
//		b EntryValue
//	)
//	a.Components=[]float64{}
//	b.Components=[]float64{}
//}

func main121() {
	UntilRedis()
	var MyRess []Res
	MyEntry.Key.Name = "翠绿之影"
	MyEntry.Key.Main = "风元素伤害加成"
	MyEntry.Key.Asse = "空之杯"
	MyEntry.Value.CN = 33
	MyEntry.Value.DGj = 5
	MyEntry.Value.GJ = 17
	MyEntry.Value.DSM = 5
	role := FindRole(MyEntry.Key.Name)
	if role[0].RoleName == "过渡使用" {
		fmt.Println("不推荐")
	}
	var resRoles []Role
	switch MyEntry.Key.Asse {
	case "时之沙":
		for _, r := range role {
			for _, hourglass := range r.RoleHourglass {
				if MyEntry.Key.Main == hourglass {
					resRoles = append(resRoles, r)
					break
				}
			}
		}
	case "空之杯":
		for _, r := range role {
			for _, cup := range r.RoleCup {
				if MyEntry.Key.Main == cup {
					resRoles = append(resRoles, r)
					break
				}
			}
		}
	case "理之冠":
		for _, r := range role {
			for _, pieum := range r.RolePileum {
				if MyEntry.Key.Main == pieum {
					resRoles = append(resRoles, r)
					break
				}
			}
		}
	default:
		for _, r := range role {
			resRoles = append(resRoles, r)
			return
		}
	}
	a := FindDate()
	for _, resRole := range resRoles {
		MyRess = append(MyRess, Res{
			Name:   resRole.RoleName,
			Source: strconv.FormatFloat(Count(MyEntry.Value, a), 'f', 1, 64),
		})
	}
	fmt.Println(MyRess)
}
func CtoE(ress []Res) []Res {
	var r []Res
	for _, res := range ress {
		switch res.Name {
		case "旅行者":
			res.Name = "Aether/Lumine"
		case "温迪":
			res.Name = "Venti"
		case "安柏":
			res.Name = "Amber"
		case "凯亚":
			res.Name = "Kaeya"
		case "丽莎":
			res.Name = "Lisa"
		case "琴":
			res.Name = "Jean"
		case "迪卢克":
			res.Name = "Diluc"
		case "罗莎莉亚":
			res.Name = "Rusaria"
		case "芭芭拉":
			res.Name = "Barbara"
		case "莫娜":
			res.Name = "Mona"
		case "班尼特":
			res.Name = "Bannett"
		case "优菈":
			res.Name = "Eula"
		case "阿贝多":
			res.Name = "Albedo"
		case "雷泽":
			res.Name = "Razor"
		case "迪奥娜":
			res.Name = "Diona"
		case "可莉":
			res.Name = "Klee"
		case "菲谢尔":
			res.Name = "Fischl"
		case "砂糖":
			res.Name = "Sucrose"
		case "诺艾尔":
			res.Name = "Nolle"
		case "米卡":
			res.Name = "Mika"
		case "钟离":
			res.Name = "Zhongli"
		case "香菱":
			res.Name = "Xiangling"
		case "魈":
			res.Name = "Xiao"
		case "刻晴":
			res.Name = "Keqing"
		case "甘雨":
			res.Name = "Ganyu"
		case "七七":
			res.Name = "Qiqi"
		case "行秋":
			res.Name = "Xingqiu"
		case "重云":
			res.Name = "Chongyun"
		case "凝光":
			res.Name = "Ningguang"
		case "北斗":
			res.Name = "Beidou"
		case "烟绯":
			res.Name = "Yanfei"
		case "申鹤":
			res.Name = "Shenhe"
		case "云堇":
			res.Name = "Yun Jin"
		case "胡桃":
			res.Name = "Hu Tao"
		case "辛焱":
			res.Name = "Xinyan"
		case "夜兰":
			res.Name = "Yelan"
		case "瑶瑶":
			res.Name = "Yaoyao"
		case "白术":
			res.Name = "Baizhu"
		case "雷电将军":
			res.Name = "Raiden Shogun"
		case "枫原万叶":
			res.Name = "Kaedehara Kazusha"
		case "托马":
			res.Name = "Thoma"
		case "神里绫华":
			res.Name = "Kaimsto Ayaka"
		case "八重神子":
			res.Name = "Yae Miko"
		case "宵宫":
			res.Name = "Yoimiya"
		case "九条娑罗":
			res.Name = "Kujo Sara"
		case "五郎":
			res.Name = "Gorou"
		case "珊瑚宫心海":
			res.Name = "Sagonomiya Kokomi"
		case "早柚":
			res.Name = "Sayu"
		case "神里绫人":
			res.Name = "Kaimsto Ayato"
		case "荒泷一斗":
			res.Name = "Arataki Itto"
		case "久岐忍":
			res.Name = "Kuki Shinobu"
		case "鹿野院平藏":
			res.Name = "Shikanoin Heizou"
		case "绮良良":
			res.Name = "Kirara"
		case "提纳里":
			res.Name = "Tighnari"
		case "柯莱":
			res.Name = "Collei"
		case "多莉":
			res.Name = "Dori"
		case "坎蒂丝":
			res.Name = "Candace"
		case "赛诺":
			res.Name = "Cyno"
		case "妮露":
			res.Name = "Nilou"
		case "纳西妲":
			res.Name = "Nahida"
		case "莱依拉":
			res.Name = "Layla"
		case "流浪者":
			res.Name = "Wanderer"
		case "珐露珊":
			res.Name = "Faruzan"
		case "艾尔海森":
			res.Name = "Alhaitham"
		case "迪希雅":
			res.Name = "Dehya"
		case "卡维":
			res.Name = "Kaveh"
		case "林尼":
			res.Name = "Lyney"
		case "琳妮特":
			res.Name = "Lynette"
		case "菲米尼":
			res.Name = "Freminet"
		case "那维莱特":
			res.Name = "Neuvillette"
		case "莱欧斯利":
			res.Name = "Wriothesley"
		case "芙宁娜":
			res.Name = "Furina"
		case "夏洛蒂":
			res.Name = "Charlotte"
		case "达达利亚":
			res.Name = "Tartaglia"
		case "埃洛伊":
			res.Name = "Aloy"
		}
		r = append(r, res)
	}
	return r
}
func maiaaan() {
	MyRess1 := []Res{
		{
			Name:   "温迪",
			Source: " 0.0",
		},
		{
			Name:   "安柏",
			Source: " 0.0",
		},
	}
	MyRess := CtoE(MyRess1)
	// 创建画布
	const width = 400
	const height = 200
	dc := gg.NewContext(width, height)

	// 设置画布背景色
	dc.SetRGB(1, 1, 1) // 白色
	dc.Clear()

	// 绘制结构体内容
	dc.SetRGB(0, 0, 0) // 黑色

	for i, ress := range MyRess {
		y1 := 25 * (i + 1)
		y2 := 33 * (i + 1)
		dc.DrawString("Name: "+ress.Name, 20, float64(y1))
		dc.DrawString("Source: "+ress.Source, 20, float64(y2))
	}
	////设置字体样式
	//err := dc.LoadFontFace("Wacko.ttf", 16)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 保存为图片文件
	dc.SavePNG("res.png")
}

// 找出主词条
func FindMain(str string) {
	switch str {
	case "BJ":
		MyEntry.Value.BJ = 0
	case "BS":
		MyEntry.Value.BS = 0
	case "YS":
		MyEntry.Value.VYS = 0
	case "CN":
		MyEntry.Value.CN = 0
	case "DGJ":
		MyEntry.Value.DGj = 0
	case "JT":
		MyEntry.Value.JT = 0
	case "DSM":
		MyEntry.Value.DSM = 0
	case "DFY":
		MyEntry.Value.DFY = 0
	case "SM":
		MyEntry.Value.SM = 0
	case "GJ":
		MyEntry.Value.GJ = 0
	}
}

// 套装适用角色
func FindRole(str string) (Roles []Role) {
	switch str {
	case "逐影猎人":
		Roles = []Role{
			{"林尼", []string{"攻击力", "元素精通"}, []string{"火元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"那维莱特", []string{"生命值"}, []string{"生命值", "水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"琳妮特", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击伤害"}},
			{"莱欧斯利", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "黄金剧团":
		Roles = []Role{
			{"菲谢尔", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"八重神子", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"阿贝多", []string{"防御力"}, []string{"岩元素伤害加成", "防御力"}, []string{"暴击率", "暴击伤害"}},
			{"纳西妲", []string{"元素精通"}, []string{"元素精通", "草元素伤害加成"}, []string{"暴击率", "暴击伤害", "元素精通"}},
		}
	case "花海甘露之光":
		Roles = []Role{
			{"迪希雅", []string{"元素充能效率", "攻击力"}, []string{"火元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"妮露", []string{"生命值"}, []string{"生命值"}, []string{"生命值"}},
		}
	case "水仙之梦":
		Roles = []Role{
			{"达达利亚", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "乐园遗落之花":
		Roles = []Role{
			{"久岐忍", []string{"元素充能效率", "元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
			{"托马", []string{"元素充能效率", "元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
		}
	case "沙上楼阁史话":
		Roles = []Role{
			{"流浪者", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"枫原万叶", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值", "元素精通"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "深林的记忆":
		Roles = []Role{
			{"纳西妲", []string{"元素精通"}, []string{"草元素伤害加成", "元素精通"}, []string{"元素精通", "暴击率", "暴击伤害"}},
			{"提纳里", []string{"元素精通", "攻击力"}, []string{"草元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"柯莱", []string{"元素精通", "攻击力", "元素充能效率"}, []string{"草元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "饰金之梦":
		Roles = []Role{
			{"纳西妲", []string{"元素精通"}, []string{"草元素伤害加成", "元素精通"}, []string{"元素精通", "暴击率", "暴击伤害"}},
			{"提纳里", []string{"元素精通", "攻击力"}, []string{"草元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值", "元素精通"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"八重神子", []string{"元素精通", "攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"甘雨", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"久岐忍", []string{"元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
			{"艾尔海森", []string{"元素精通", "元素充能效率"}, []string{"元素精通", "草元素伤害加成"}, []string{"暴击率", "暴击伤害", "元素精通"}},
		}
	case "辰砂往生录":
		Roles = []Role{
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}}}
	case "来歆余响":
		Roles = []Role{
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"宵宫", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "华馆梦醒形骸记":
		Roles = []Role{
			{"荒泷一斗", []string{"防御力"}, []string{"岩元素伤害加成", "防御力"}, []string{"暴击率", "暴击伤害", "防御力"}},
			{"阿贝多", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害", "防御力"}},
			{"诺艾尔", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害", "防御力"}},
			{"云堇", []string{"防御力", "元素充能效率"}, []string{"防御力"}, []string{"防御力"}},
		}
	case "海染砗磲":
		Roles = []Role{
			{"珊瑚宫心海", []string{"生命值"}, []string{"生命值", "水元素伤害加成"}, []string{"治疗加成"}},
			{"七七", []string{"攻击力"}, []string{"攻击力"}, []string{"治疗加成"}},
			{"芭芭拉", []string{"生命值"}, []string{"生命值"}, []string{"治疗加成"}},
		}
	case "绝缘之旗印":
		Roles = []Role{
			{"雷电将军", []string{"元素充能效率", "攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"莫娜", []string{"元素充能效率"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"夜兰", []string{"生命值"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"行秋", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"香菱", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"重云", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"凯亚", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"温迪", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"神里绫华", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"达达利亚", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"旅行者", []string{"元素充能效率"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "追忆之注连":
		Roles = []Role{
			{"宵宫", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"菲谢尔", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "千岩牢固":
		Roles = []Role{
			{"钟离", []string{"生命值"}, []string{"生命值", "岩元素伤害加成"}, []string{"暴击率", "暴击伤害", "生命值"}},
			{"妮露", []string{"生命值"}, []string{"生命值"}, []string{"生命值"}},
			{"瑶瑶", []string{"生命值", "元素充能效率"}, []string{"生命值"}, []string{"生命值", "治疗加成"}},
			{"夜兰", []string{"生命值"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"阿贝多", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"珊瑚宫心海", []string{"生命值"}, []string{"生命值"}, []string{"治疗加成"}},
			{"莫娜", []string{"元素充能效率"}, []string{"水元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"菲谢尔", []string{"元素充能效率"}, []string{"雷元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
		}
	case "苍白之火":
		Roles = []Role{
			{"雷泽", []string{"攻击力"}, []string{"攻击力", "物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"优菈", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "平息鸣雷的尊者":
		Roles = []Role{
			{"刻晴", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"北斗", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"菲谢尔", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "炽烈的炎之魔女":
		Roles = []Role{
			{"迪卢克", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"可莉", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值", "元素精通"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"香菱", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"烟绯", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "流浪大地的乐团":
		Roles = []Role{
			{"甘雨", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"可莉", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"安柏", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"烟绯", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "染血的骑士道":
		Roles = []Role{{"辛焱", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}}}
	case "被怜爱的少女":
		Roles = []Role{
			{"珊瑚宫心海", []string{"生命值"}, []string{"生命值"}, []string{"治疗加成"}},
			{"芭芭拉", []string{"生命值"}, []string{"生命值"}, []string{"治疗加成"}},
			{"迪奥娜", []string{"生命值", "元素充能效率"}, []string{"生命值"}, []string{"治疗加成"}},
			{"班尼特", []string{"生命值", "元素充能效率"}, []string{"生命值"}, []string{"治疗加成"}},
			{"七七", []string{"攻击力", "元素充能效率"}, []string{"攻击力"}, []string{"治疗加成"}},
			{"琴", []string{"攻击力", "元素充能效率"}, []string{"攻击力"}, []string{"攻击力", "治疗加成"}},
		}
	case "角斗士的终幕礼":
		Roles = []Role{
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "渡过烈火的贤人":
		Roles = []Role{
			{"迪卢克", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"可莉", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "悠古的磐岩":
		Roles = []Role{
			{"阿贝多", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"诺艾尔", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"凝光", []string{"攻击力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "如雷的盛怒":
		Roles = []Role{
			{"刻晴", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"北斗", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"菲谢尔", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"丽莎", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成", "元素精通"}, []string{"暴击率", "暴击伤害", "元素精通"}},
		}
	case "沉沦之心":
		Roles = []Role{
			{"达达利亚", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"珊瑚宫心海", []string{"生命值"}, []string{"水元素伤害加成"}, []string{"治疗加成"}},
			{"行秋", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"莫娜", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "逆飞的流星":
		Roles = []Role{
			{"钟离", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"诺艾尔", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "昔日宗室之仪":
		Roles = []Role{
			{"班尼特", []string{"生命值", "元素充能效率"}, []string{"生命值"}, []string{"生命值", "治疗加成"}},
			{"迪奥娜", []string{"生命值", "元素充能效率"}, []string{"生命值"}, []string{"生命值", "治疗加成"}},
			{"砂糖", []string{"元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
			{"莫娜", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"重云", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"达达利亚", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"行秋", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"申鹤", []string{"攻击力", "元素充能效率"}, []string{"攻击力"}, []string{}},
		}
	case "翠绿之影":
		Roles = []Role{
			{"温迪", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"旅行者", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"琴", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"枫原万叶", []string{"元素精通"}, []string{"元素精通", "风元素伤害加成"}, []string{"元素精通", "暴击率", "暴击伤害"}},
			{"砂糖", []string{"元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
			{"早柚", []string{"攻击力", "元素充能效率", "元素精通"}, []string{"元素精通"}, []string{"元素精通", "治疗加成", "攻击力"}},
		}
	case "冰风迷途的勇士":
		Roles = []Role{
			{"甘雨", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"神里绫华", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"重云", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"凯亚", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"莱欧斯利", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
		}
	default:
		Roles = []Role{
			{"过渡使用", []string{}, []string{}, []string{}},
		}
	}
	return Roles
}

// 计算通用评分
func TongYongSource() (res int) {
	num := ReturnResultForTY([]EntryValue{MaxBj, MaxBS, MaxDGJ, MaxCN}, MyEntry.Value)
	fmt.Printf("具体大小:%s\n", strconv.FormatFloat(num, 'f', 1, 64))
	res = 1
	if num > 168 {
		res = 2
	}
	if num > 175 {
		res = 3
	}
	if num > 200 {
		res = 4
	}
	if num > 230 {
		res = 5
	}
	return res
}

// 计算精通评分
func JinTongSource() (res int) {
	num := ReturnResult([]EntryValue{MaxBj, MaxBS, MaxJT, MaxCN}, MyEntry.Value)
	fmt.Printf("具体大小:%s\n", strconv.FormatFloat(num, 'f', 1, 64))
	res = 1
	if num > 280 {
		res = 2
	}
	if num > 330 {
		res = 3
	}
	if num > 380 {
		res = 4
	}
	if num > 450 {
		res = 5
	}
	return res
}

// 计算生命评分
func ShenMingSource() (res int) {
	num := ReturnResult([]EntryValue{MaxBj, MaxBS, MaxDSM, MaxCN}, MyEntry.Value)
	fmt.Printf("具体大小:%s\n", strconv.FormatFloat(num, 'f', 1, 64))
	res = 1
	if num > 168 {
		res = 2
	}
	if num > 175 {
		res = 3
	}
	if num > 200 {
		res = 4
	}
	if num > 230 {
		res = 5
	}
	return res
}

// 计算防御评分
func FangYuSource() (res int) {
	num := ReturnResult([]EntryValue{MaxBj, MaxBS, MaxDFY, MaxCN}, MyEntry.Value)
	fmt.Printf("具体大小:%s\n", strconv.FormatFloat(num, 'f', 1, 64))
	res = 1
	if num > 168 {
		res = 2
	}
	if num > 175 {
		res = 3
	}
	if num > 200 {
		res = 4
	}
	if num > 230 {
		res = 5
	}
	return res
}

// 计算契合程度
func ReturnResult(nums []EntryValue, n EntryValue) float64 {
	var source float64
	for _, num := range nums {
		source = source + (abs(n.BJ-num.BJ) + abs(n.BS-num.BS) + abs(n.CN-num.CN) + abs(n.DGj-num.DGj) + abs(n.GJ-num.GJ) +
			abs(n.JT-num.JT) + abs(n.DSM-num.DSM) + abs(n.SM-num.SM) + abs(n.DFY-num.DFY) + abs(n.FY-num.FY))
	}
	return source
}

// 计算契合程度
func ReturnResultForTY(nums []EntryValue, n EntryValue) float64 {
	var source float64
	var nnn float64
	for i, num := range nums {
		//fmt.Println(num)
		switch i {
		case 0:
			nnn = 5
		case 1:
			nnn = 2.5
		case 2:
			nnn = 1.5
		case 3:
			nnn = 1
		}
		source = source + nnn*0.1*(abs(n.BJ-num.BJ)+abs(n.BS-num.BS)+abs(n.CN-num.CN)+abs(n.DGj-num.DGj)+abs(n.GJ-num.GJ)+
			abs(n.JT-num.JT)+abs(n.DSM-num.DSM)+abs(n.SM-num.SM)+abs(n.DFY-num.DFY)+abs(n.FY-num.FY))
		fmt.Printf("res%d:%s\n", i, strconv.FormatFloat(source, 'f', 1, 64))
	}
	return source
}

// abs用于获取绝对值
func abs(a float64) float64 {
	if a < 0 {
		a = a * -1
	}
	return a
}

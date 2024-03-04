package tools

import (
	"Toupiao/frav2/model"
	"encoding/json"
	"fmt"
	"github.com/fogleman/gg"
)

// 结果转为英文
func CtoE(ress []model.Res) []model.Res {
	var r []model.Res
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

// 套装适用角色
func FindRole(str string) (Roles []model.Role) {
	switch str {
	case "逐影猎人":
		Roles = []model.Role{
			{"林尼", []string{"攻击力", "元素精通"}, []string{"火元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"那维莱特", []string{"生命值"}, []string{"生命值", "水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"琳妮特", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击伤害"}},
			{"莱欧斯利", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "黄金剧团":
		Roles = []model.Role{
			{"菲谢尔", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"八重神子", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"阿贝多", []string{"防御力"}, []string{"岩元素伤害加成", "防御力"}, []string{"暴击率", "暴击伤害"}},
			{"纳西妲", []string{"元素精通"}, []string{"元素精通", "草元素伤害加成"}, []string{"暴击率", "暴击伤害", "元素精通"}},
		}
	case "花海甘露之光":
		Roles = []model.Role{
			{"迪希雅", []string{"元素充能效率", "攻击力"}, []string{"火元素伤害加成", "攻击力"}, []string{"暴击率", "暴击伤害"}},
			{"妮露", []string{"生命值"}, []string{"生命值"}, []string{"生命值"}},
		}
	case "水仙之梦":
		Roles = []model.Role{
			{"达达利亚", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "乐园遗落之花":
		Roles = []model.Role{
			{"久岐忍", []string{"元素充能效率", "元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
			{"托马", []string{"元素充能效率", "元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
		}
	case "沙上楼阁史话":
		Roles = []model.Role{
			{"流浪者", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"枫原万叶", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值", "元素精通"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "深林的记忆":
		Roles = []model.Role{
			{"纳西妲", []string{"元素精通"}, []string{"草元素伤害加成", "元素精通"}, []string{"元素精通", "暴击率", "暴击伤害"}},
			{"提纳里", []string{"元素精通", "攻击力"}, []string{"草元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"柯莱", []string{"元素精通", "攻击力", "元素充能效率"}, []string{"草元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "饰金之梦":
		Roles = []model.Role{
			{"纳西妲", []string{"元素精通"}, []string{"草元素伤害加成", "元素精通"}, []string{"元素精通", "暴击率", "暴击伤害"}},
			{"提纳里", []string{"元素精通", "攻击力"}, []string{"草元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值", "元素精通"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"八重神子", []string{"元素精通", "攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"甘雨", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"久岐忍", []string{"元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
			{"艾尔海森", []string{"元素精通", "元素充能效率"}, []string{"元素精通", "草元素伤害加成"}, []string{"暴击率", "暴击伤害", "元素精通"}},
		}
	case "辰砂往生录":
		Roles = []model.Role{
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}}}
	case "来歆余响":
		Roles = []model.Role{
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"宵宫", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "华馆梦醒形骸记":
		Roles = []model.Role{
			{"荒泷一斗", []string{"防御力"}, []string{"岩元素伤害加成", "防御力"}, []string{"暴击率", "暴击伤害", "防御力"}},
			{"阿贝多", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害", "防御力"}},
			{"诺艾尔", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害", "防御力"}},
			{"云堇", []string{"防御力", "元素充能效率"}, []string{"防御力"}, []string{"防御力"}},
		}
	case "海染砗磲":
		Roles = []model.Role{
			{"珊瑚宫心海", []string{"生命值"}, []string{"生命值", "水元素伤害加成"}, []string{"治疗加成"}},
			{"七七", []string{"攻击力"}, []string{"攻击力"}, []string{"治疗加成"}},
			{"芭芭拉", []string{"生命值"}, []string{"生命值"}, []string{"治疗加成"}},
		}
	case "绝缘之旗印":
		Roles = []model.Role{
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
		Roles = []model.Role{
			{"宵宫", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"菲谢尔", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "千岩牢固":
		Roles = []model.Role{
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
		Roles = []model.Role{
			{"雷泽", []string{"攻击力"}, []string{"攻击力", "物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"优菈", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "平息鸣雷的尊者":
		Roles = []model.Role{
			{"刻晴", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"北斗", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"菲谢尔", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "炽烈的炎之魔女":
		Roles = []model.Role{
			{"迪卢克", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"可莉", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"胡桃", []string{"生命值", "元素精通"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"香菱", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"烟绯", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "流浪大地的乐团":
		Roles = []model.Role{
			{"甘雨", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"可莉", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"安柏", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"烟绯", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "染血的骑士道":
		Roles = []model.Role{{"辛焱", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}}}
	case "被怜爱的少女":
		Roles = []model.Role{
			{"珊瑚宫心海", []string{"生命值"}, []string{"生命值"}, []string{"治疗加成"}},
			{"芭芭拉", []string{"生命值"}, []string{"生命值"}, []string{"治疗加成"}},
			{"迪奥娜", []string{"生命值", "元素充能效率"}, []string{"生命值"}, []string{"治疗加成"}},
			{"班尼特", []string{"生命值", "元素充能效率"}, []string{"生命值"}, []string{"治疗加成"}},
			{"七七", []string{"攻击力", "元素充能效率"}, []string{"攻击力"}, []string{"治疗加成"}},
			{"琴", []string{"攻击力", "元素充能效率"}, []string{"攻击力"}, []string{"攻击力", "治疗加成"}},
		}
	case "角斗士的终幕礼":
		Roles = []model.Role{
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "渡过烈火的贤人":
		Roles = []model.Role{
			{"迪卢克", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"可莉", []string{"攻击力"}, []string{"火元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "悠古的磐岩":
		Roles = []model.Role{
			{"阿贝多", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"诺艾尔", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"凝光", []string{"攻击力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "如雷的盛怒":
		Roles = []model.Role{
			{"刻晴", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"北斗", []string{"攻击力"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"菲谢尔", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"丽莎", []string{"攻击力", "元素精通"}, []string{"雷元素伤害加成", "元素精通"}, []string{"暴击率", "暴击伤害", "元素精通"}},
		}
	case "沉沦之心":
		Roles = []model.Role{
			{"达达利亚", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"珊瑚宫心海", []string{"生命值"}, []string{"水元素伤害加成"}, []string{"治疗加成"}},
			{"行秋", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"莫娜", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"神里绫人", []string{"攻击力"}, []string{"水元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "逆飞的流星":
		Roles = []model.Role{
			{"钟离", []string{"攻击力"}, []string{"物理伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"诺艾尔", []string{"防御力"}, []string{"岩元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
		}
	case "昔日宗室之仪":
		Roles = []model.Role{
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
		Roles = []model.Role{
			{"温迪", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"魈", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"旅行者", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"琴", []string{"攻击力"}, []string{"风元素伤害加成"}, []string{"暴击率", "暴击伤害"}},
			{"枫原万叶", []string{"元素精通"}, []string{"元素精通", "风元素伤害加成"}, []string{"元素精通", "暴击率", "暴击伤害"}},
			{"砂糖", []string{"元素精通"}, []string{"元素精通"}, []string{"元素精通"}},
			{"早柚", []string{"攻击力", "元素充能效率", "元素精通"}, []string{"元素精通"}, []string{"元素精通", "治疗加成", "攻击力"}},
		}
	case "冰风迷途的勇士":
		Roles = []model.Role{
			{"甘雨", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"神里绫华", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"重云", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"凯亚", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
			{"莱欧斯利", []string{"攻击力"}, []string{"冰元素伤害加成"}, []string{"暴击伤害"}},
		}
	default:
		Roles = []model.Role{
			{"过渡使用", []string{}, []string{}, []string{}},
		}
	}
	return Roles
}

// 找到最想的向量
func Count(e model.EntryValue, a []model.MY) float64 {
	var res float64
	var nnn model.EntryValueComponents
	var aa []model.EntryValueComponents
	for _, my := range a {
		a1 := model.AddComponent(my.A)
		aa = append(aa, a1)
	}
	r := model.AddComponent(e)
	for i, components := range aa {
		re := model.AngleBetween(components, r)
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

// 获取向量库
func FindDate() []model.MY {
	var a1 []model.MY
	var cc model.MY
	for i := 0; i < 35838; i++ {
		//fmt.Println(i)
		key := fmt.Sprintf("杯子:%d", i)
		r := model.GetRedis(key)
		_ = json.Unmarshal([]byte(r), &cc)
		a1 = append(a1, cc)
	}
	return a1
}
func SavePhoto(MyRess1 []model.Res, path string) {
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
	// 保存为图片文件
	dc.SavePNG(path)
}

type Code struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

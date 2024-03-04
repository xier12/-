package model

import (
	"math"
	"sort"
)

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

// 根据结构体生成十维数组
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

// 存储的数据
type MY struct {
	//圣遗物结构体
	A EntryValue
	//当前圣遗物拥有的词条
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

// 结果
type Res struct {
	//角色
	Name string
	//推荐程度
	Source string
}
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

// 根据圣遗物生成词条数组
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

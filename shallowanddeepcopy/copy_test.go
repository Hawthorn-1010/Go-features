package shallowanddeepcopy

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 值类型的数据，默认都是深拷贝
//int、float、string、bool、array、struct
//引用类型的数据，默认都是浅拷贝
//slice、map、function

func TestPerCopy(t *testing.T) {
	// 在结构体当中没有指针，没有slice和map类型，是深拷贝
	p1 := Per{
		Name:    "lily",
		Age:     20,
		HouseId: [3]int{1, 2, 3},
	}

	p2 := p1

	fmt.Printf("%v,%p\n", p1, &p1)
	fmt.Printf("%v,%p\n", p2, &p2)

	// 指针浅拷贝
	p3 := Per{
		Name:    "lucas",
		Age:     20,
		HouseId: [3]int{1, 2, 3},
	}

	p4 := &p3

	fmt.Printf("%v,%p\n", *p4, p4)
	fmt.Printf("%v,%p\n", p3, &p3)

	p4.HouseId[0] = 100
	fmt.Println(*p4, p3)

	// 使用new也是浅拷贝
	p5 := new(Per)
	p5.Name = "lucas"
	p5.Age = 20
	p5.HouseId = [3]int{1, 2, 3}

	p6 := p5
	p6.Age = 22
	fmt.Printf("%v,%p\n", p5, &p5)
	fmt.Printf("%v,%p\n", p6, &p6)
}

func TestPer1Copy(t *testing.T) {
	// 结构体当中含有引用类型的字段，那对应的字段是浅拷贝。如果里面有字段会发生浅拷贝，为了避免浅拷贝可以将浅拷贝的字段拿出来单个挨个赋值。
	p1 := Per1{
		Name:    "lily",
		Age:     20,
		HouseId: [3]int{1, 2, 3},
		Tags:    map[string]string{"k1": "v1", "k2": "v2"},
	}

	p2 := p1

	// name是深拷贝
	p2.Name = "bob"
	// tags是浅拷贝
	p2.Tags["k1"] = "v10"
	fmt.Println(p1, p2)
}

/*
把结构体的引用类型变成深拷贝，方法一：引用类型全部重新赋值
*/
func TestPer1DeepCopy1(t *testing.T) {
	p1 := Per1{
		Name:    "lucas",
		Age:     20,
		HouseId: [3]int{1, 2, 3},
		Tags:    map[string]string{"k1": "v1", "k2": "v2"},
	}

	p2 := p1

	//tmpCarId := []int{}

	//for _, v := range p1.CarId {
	//	tmpCarId = append(tmpCarId, v)
	//}
	//fmt.Println(tmpCarId)

	tmpTags := map[string]string{}
	for k, v := range p1.Tags {
		tmpTags[k] = v
	}
	fmt.Println(tmpTags)

	p2.Tags = tmpTags
	//p2.CarId = tmpCarId
}

/*
把结构体的引用类型变成深拷贝，方法二：使用反射或者json
*/
func TestPer1DeepCopy2(t *testing.T) {
	m1 := Man{
		Name:    "jordon",
		HouseId: []int{1, 2, 3, 4},
		Tags:    map[string]string{"k1": "v1", "k2": "v2"},
	}

	var m3 Man
	data, _ := json.Marshal(&m1)
	fmt.Println(string(data))
	json.Unmarshal(data, &m3)
	m3.HouseId = []int{2, 4, 6}
	fmt.Printf("%v", m1)
	fmt.Printf("%v", m3)
}

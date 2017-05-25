package main

import "fmt"

type Enum int

const (
	Apple Enum = iota
	Pear
	Banana
)

type subS struct {
	Fruit Enum
}

type subsub struct {
	Sub subS
}

type S struct {
	Fruit1  Enum
	Fruit2  Enum
	Fruit3  Enum
	Fruit4  Enum
	Fruit5  Enum
	Sub     subS
	SubSub1 subsub
	SubSub2 subsub
}

func GetInt() int {
	return 10
}

func f1() {
	var f Enum

	f = Apple
	f = 10
	f = Enum(55.)

	fmt.Print(f)

	ch := make(chan int, 1)
	ch <- 123

	chEnum := make(chan Enum, 1)
	chEnum <- 321

	m := map[string]Enum{
		"aaa": 1234,
		"bbb": Banana,
	}

	s := S{
		Fruit1: 88,
		Fruit2: Pear,
		Fruit3: f,
		Fruit4: Enum(GetInt()),
		Fruit5: Enum(<-ch),
		Sub: subS{
			Fruit: 11,
		},
		SubSub1: subsub{
			Sub: subS{
				Fruit: 22,
			},
		},
		SubSub2: subsub{
			Sub: subS{
				Fruit: Enum(44),
			},
		},
	}

	if f == Enum(22) {
		fmt.Println("equal")
	}

	if s.SubSub2.Sub.Fruit == 1 {
		fmt.Println("equal")
	}

	fmt.Print(s, m)
}

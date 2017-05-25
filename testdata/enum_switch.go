package main

import "fmt"

type Enum int

const (
	Apple Enum = iota
	Pear
	Banana
)

type st struct {
	fruit Enum
}

type nested struct {
	nestedFruit st
}

type doubleNested struct {
	dbl nested
}

func f2() {
	a := Apple

	switch a {
	case Apple:
		fmt.Println("Apple")
	}

	s := st{
		fruit: Banana,
	}

	switch s.fruit {
	case Banana:
		fmt.Println("Banana")
	}

	nstd := nested{
		nestedFruit: st{
			fruit: Pear,
		},
	}

	switch nstd.nestedFruit.fruit {
	case Pear:
		fmt.Println("Pear")
	}

	dn := doubleNested{
		dbl: nested{
			nestedFruit: st{
				fruit: Pear,
			},
		},
	}

	switch dn.dbl.nestedFruit.fruit {
	case Pear:
		fmt.Println("Pear")
	}

	arr := []Enum{Apple, Pear, Banana}

	switch arr[0] {
	case Apple:
		fmt.Println("Apple")
	}

	ch := make(chan Enum, 1)
	ch <- Pear

	switch <-ch {
	case Pear:
		fmt.Println("Pear")
	}

	aa := dn.dbl.nestedFruit.fruit

	switch aa {
	case arr[2]:
		fmt.Println("Pear")
	case 1.0:
		fmt.Println("1.0")
	case 2:
		fmt.Println("2")
	case Apple:
		fmt.Println("Apple")
	}
}

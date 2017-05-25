Enum linter
===========

Enum linter checks:
 - incorrect type or basic literal is assigned to enumerated type value 
 - switch clause is not exhaustive (in progress)
    
Example
=======
```
type Enum int

const (
    Apple Enum = iota
    Pear
    Banana
)

...

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
```

Run
===
```
$ enumlinter -type Enum ./enum.go
```

```
./testdata/enum_assignment.go:40:6: Enum type constant must be used instead of a basic literal 10
./testdata/enum_assignment.go:41:6: ToEnum() like func must be used instead of a type casting Enum(55.)
./testdata/enum_assignment.go:49:12: Enum type constant must be used instead of a basic literal 321
./testdata/enum_assignment.go:52:10: Enum type constant must be used instead of a basic literal 1234
./testdata/enum_assignment.go:57:11: Enum type constant must be used instead of a basic literal 88
./testdata/enum_assignment.go:60:11: ToEnum() like func must be used instead of a type casting Enum(GetInt())
./testdata/enum_assignment.go:61:11: ToEnum() like func must be used instead of a type casting Enum(<-ch)
./testdata/enum_assignment.go:63:11: Enum type constant must be used instead of a basic literal 11
./testdata/enum_assignment.go:67:12: Enum type constant must be used instead of a basic literal 22
./testdata/enum_assignment.go:72:12: ToEnum() like func must be used instead of a type casting Enum(44)
./testdata/enum_assignment.go:77:10: ToEnum() like func must be used instead of a type casting Enum(22)
./testdata/enum_assignment.go:81:28: Enum type constant must be used instead of a basic literal 1
```

        
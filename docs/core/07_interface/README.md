# 接口和抽象

## type

Golang 具备完善的 type 体系，有 built-in types，也有自定义类型

Built-in types includes:

1. Built-in string type: string.
2. Built-in boolean type: bool.
3. Built-in numeric types:
    1. int8, uint8 (byte), int16, uint16, int32 (rune), uint32, int64, uint64, int, uint, uintptr.
    2. float32, float64.
    3. complex64, complex128. 
   
Note, byte is a built-in alias of uint8, and rune is a built-in alias of int32. We can also declare custom type aliases (see below).

Golang 支持的类型有：

- pointer types
- struct types
- function types
- container types
    - array types
    - slice types
    - map types
- channel types
- interface types
- [unsafe types](https://pkg.go.dev/unsafe)

```
*T         // a pointer type
[5]T       // an array type
[]T        // a slice type
map[Tkey]T // a map type

// a struct type
struct {
	name string
	age  int
}

// a function type
func(int) (bool, string)

// an interface type
interface {
	Method0(string) int
	Method1() (int, bool)
}

// some channel types
chan T
chan<- T
<-chan T
```

定义新的 type

```
// Define a solo new type.
type NewTypeName SourceType

// Define multiple new types together.
type (
	NewTypeName1 SourceType1
	NewTypeName2 SourceType2
)

// The following new defined and source types
// are all basic types.
type (
	MyInt int
	Age   int
	Text  string
)

// The following new defined and source types are
// all composite types.
type IntPtr *int
type Book struct{author, title string; pages int}
type Convert func(in0 int, in1 bool)(out0 int, out1 string)
type StringArray [5]string
type StringSlice []string

func f() {
	// The names of the three defined types
	// can be only used within the function.
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{Read([]byte) int}
}
```

type alias (类型重命名)

```
type (
	Name = string
	Age  = int
)

type table = map[string]int
type Table = map[Name]Age
```

Note: Type alias declarations are useful in refactoring some large Go projects, they are not intended for general purpose uses. 
We should use type definition declarations in general programming.

via:

- [Go Type System Overview](https://go101.org/article/type-system-overview.html)

# Go 编程范式

Go 可以被认为是一个OO的语言，OOP的基石是封装、继承、多态等特性，Go 语言中又是如何表达的呢？

## 用 Go 做 OOP

### Go 的设计哲学

Go 语言崇尚极致的简单性。

1. Go 语言没有类，没有对象，没有异常和模板
2. Go 语言支持垃圾回收和内建的并发
3. 有关面向对象方面的最显著的遗漏是Go语言中没有类型层次结构, 与大多数面向对象语言如C++，Java，C＃，Scala甚至动态语言（如Python和Ruby）形成对比

### Go OO 特性

Go语言没有类，但它支持类型。 特别是， 它支持structs。 Structs是用户定义的类型。 Struct类型(含方法)提供类似于其它语言中类的服务。

- struct

一个struct定义一个状态。Structs 只保存状态，不保存行为。

```golang
type User struct {
    Name string
    Age int
}
```

- 方法

方法是对特定类型进行操作的函数。他们有一个 receiver，命令它们对什么样的类型可进行操作

```golang
func (u User) Update() {
    // ...
}
```

- 嵌入

你可以将匿名的类型嵌入进彼此。 如果你嵌入一个无名的struct那么被嵌入的struct对接受嵌入的struct直接提供它自己的状态（和方法）。比如

```golang
type Creature struct {
  Name string
  Real bool
}

// FlyingCreature有一个无名子的被嵌入的Creature struct，这意味着一个FlyingCreature就是一个Creature
type FlyingCreature struct {
  Creature
  WingSpan int
}

// 如果你有一个FlyingCreature的实例，你可以直接访问它的Name和Real属性
dragon := &FlyingCreature{
    Creature{"Dragon", false, },
    15,
}

fmt.Println(dragon.Name)
fmt.Println(dragon.Real)
fmt.Println(dragon.WingSpan)
```

- interface

接口是Go语言对面向对象支持的标志。 接口是声明方法集的类型。 与其它语言中的接口类似，它们不包含方法的实现。
实现所有接口方法的对象自动地实现接口。 它没有继承或子类或“implements”关键字。这种特性叫 `duck-typing`.

### OOD: Go 的方式

OO 的核心：对象是语言的构造，它们有状态和行为，这些行为对状态进行操作并且选择性地把它们曝露给程序的其它部分。

- 封装

Go语言在包的级别进行封装。 以小写字母开头的名称只在该程序包中可见。 你可以隐藏私有包中的任何内容，只暴露特定的类型，接口和工厂函数。

```golang
// 在这里要隐藏Foo类型，只暴露接口，可以将其重命名为小写的foo，并提供一个NewFoo()函数，返回公共Fooer接口
type Fooer interface {
  Foo1()
  Foo2()
  Foo3()
}

type foo struct {
}

func (f foo) Foo1() {
    fmt.Println("Foo1() here")
}

func (f foo) Foo2() {
    fmt.Println("Foo2() here")
}

func (f foo) Foo3() {
    fmt.Println("Foo3() here")
}

func NewFoo() Fooer {
    return &Foo{}
}
// 然后来自另一个包的代码可以使用NewFoo()并访问由内部foo类型实现的Footer接口
f := NewFoo()
f.Foo1()
f.Foo2()
f.Foo3()
```

- 继承

继承或子类化始终是一个有争议的问题。 实现继承有许多问题（与接口继承相反）。
C++和Python以及其它语言实现的多重继承受着致命的死亡钻石问题，但即使是单一继承也会有面对脆弱的基类问题的处境。

现代语言和面向对象的思维现在倾向于组合而不是继承。 Go语言对此很严肃，它没有任何类型层次结构。
它允许你通过组合来共享实现的细节。 但Go语言，在一个非常奇怪的转变中（这可能源于实用的考量），允许嵌入匿名组合。

通过嵌入一个匿名类型的组合等同于实现继承，这是它所有意图和目的。 一个嵌入的struct与基类一样脆弱。
你还可以嵌入一个接口，这相当于在Java或C ++等语言中继承一个接口。 如果嵌入类型没有实现所有接口方法，
它甚至可能导致产生在编译时未被发现的运行错误。

```
// 这里SuperFoo嵌入Fooer接口，但是SuperFoo没有实现Foo的方法。
// Go编译器会愉快地让你创建一个新的SuperFood并调用Fooer的方法，但很显然这在运行时会失败。
type SuperFooer struct {
  Fooer
}

func main() {
  s := SuperFooer{}
  s.Foo2()
}

// error
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x28 pc=0x2a78]

goroutine 1 [running]:
panic(0xde180, 0xc82000a0d0)
  /usr/local/Cellar/go/1.6/libexec/src/runtime/panic.go:464 +0x3e6
main.main()
  /Users/gigi/Documents/dev/go/src/github.com/oop_test/main.go:104 +0x48
exit status 2

Process finished with exit code 1
```

- 多态性

多态性是面向对象编程的本质：只要对象坚持实现同样的接口，Go语言就能处理不同类型的那些对象。 Go接口以非常直接和直观的方式提供这种能力。

```golang
package main

import "fmt"

type Creature struct {
  Name string
  Real bool
}

func Dump(c*Creature) {
  fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

func (c Creature) Dump() {
  fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

type FlyingCreature struct {
  Creature
  WingSpan int
}

func (fc FlyingCreature) Dump() {
  fmt.Printf("Name: '%s', Real: %t, WingSpan: %d\n",
    fc.Name,
    fc.Real,
    fc.WingSpan)
}

type Unicorn struct {
  Creature
}

type Dragon struct {
  FlyingCreature
}

type Pterodactyl struct {
  FlyingCreature
}

func NewPterodactyl(wingSpan int) *Pterodactyl {
  pet := &Pterodactyl{
    FlyingCreature{
      Creature{
        "Pterodactyl",
        true,
      },
      wingSpan,
    },
  }
  return pet
}

type Dumper interface {
  Dump()
}

type Door struct {
  Thickness int
  Color     string
}

func (d Door) Dump() {
  fmt.Printf("Door => Thickness: %d, Color: %s", d.Thickness, d.Color)
}

func main() {
  creature := &Creature{
    "some creature",
    false,
  }

  uni := Unicorn{
    Creature{
      "Unicorn",
      false,
    },
  }

  pet1 := &Pterodactyl{
    FlyingCreature{
      Creature{
        "Pterodactyl",
        true,
      },
      5,
    },
  }

  pet2 := NewPterodactyl(8)

  door := &Door{3, "red"}

  Dump(creature)
  creature.Dump()
  uni.Dump()
  pet1.Dump()
  pet2.Dump()

  creatures := []Creature{
    *creature,
    uni.Creature,
    pet1.Creature,
    pet2.Creature}
  fmt.Println("Dump() through Creature embedded type")
  for _, creature := range creatures {
    creature.Dump()
  }

  dumpers := []Dumper{creature, uni, pet1, pet2, door}
  fmt.Println("Dump() through Dumper interface")
  for _, dumper := range dumpers {
    dumper.Dump()
  }
}
```
## 博客

- [与Go同行：Golang面向对象编程](https://code.tutsplus.com/zh-hans/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540)
- [37 | 编程范式游记（8）- Go 语言的委托模式](https://time.geekbang.org/column/article/2748)
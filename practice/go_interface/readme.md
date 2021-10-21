# Interface 接口

## v0 
这段代码中使用了一个 Printable 的接口，而 Country 和 City 都实现了接口方法 PrintStr() 把自己输出。
然而，这些代码都是一样的，能不能省掉呢？

## v1
其实，我们可以使用“结构体嵌入”的方式来完成这个事，如下所示：
引入一个叫 WithName的结构体，但是这会带来一个问题：在初始化的时候变得有点乱。那么，有没有更好的方法呢？再来看另外一个解。

## v2
在这段代码中，我们可以看到，我们使用了一个叫Stringable 的接口，
我们用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了。
于是，只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。

## summary
这种编程模式在 Go 的标准库有很多的示例，最著名的就是 io.Read 和 ioutil.ReadAll 的玩法，

其中 io.Read 是一个接口，你需要实现它的一个 Read(p []byte) (n int, err error) 接口方法，只要满足这个规则，就可以被 ioutil.ReadAll这个方法所使用。

这就是面向对象编程方法的黄金法则——“Program to an interface not an implementation”。

## 接口完整性检查

我们可以看到，Go 语言的编译器并没有严格检查一个对象是否实现了某接口所有的接口方法，如下面这个示例：

```go

type Shape interface {
    Sides() int
    Area() int
}
type Square struct {
    len int
}
func (s* Square) Sides() int {
    return 4
}
func main() {
    s := Square{len: 5}
    fmt.Printf("%d\n",s.Sides())
}
```

在 Go 语言编程圈里，有一个比较标准的做法：

```go
 var _ Shape = (*Square)(nil)
```

声明一个 _ 变量（没人用）会把一个 nil 的空指针从 Square 转成 Shape，这样，如果没有实现完相关的接口方法，编译器就会报错：





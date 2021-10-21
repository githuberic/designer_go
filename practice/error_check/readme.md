## Error Check Hell

说到 Go 语言的 if err !=nil 的代码了，这样的代码的确是能让人写到吐。

```go
func parse(r io.Reader) (*Point, error) {

    var p Point

    if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
        return nil, err
    }
}
```

要解决这个事，我们可以用<b>函数式编程</b>的方式，如下代码示例：

```go
func parse(r io.Reader) (*error_check.Point, error) {
	var p error_check.Point
	var err error
	
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}
	
	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)
	
	if err != nil {
		return &p, err
	}
	
	return &p, nil
}
```

从这段代码中，我们可以看到，我们通过使用 Closure 的方式把相同的代码给抽出来重新定义一个函数，

这样大量的 if err!=nil 处理得很干净了，但是会带来一个问题，那就是有一个 err 变量和一个内部的函数，感觉不是很干净。

那么，我们还能不能搞得更干净一点呢？我们从 Go 语言的 bufio.Scanner()中似乎可以学习到一些东西：

```go
scanner := bufio.NewScanner(input)

for scanner.Scan() {
    token := scanner.Text()
    // process token
}

if err := scanner.Err(); err != nil {
    // process the error
}
```

可以看到，scanner在操作底层的 I/O 的时候，那个 for-loop 中没有任何的 if err !=nil 的情况，退出循环后有一个 scanner.Err() 的检查，看来使用了结构体的方式。

模仿它，就可以对我们的代码进行重构了。

首先，定义一个结构体和一个成员函数：

```go
type Reader struct {
    r   io.Reader
    err error
}

func (r *Reader) read(data interface{}) {
    if r.err == nil {
        r.err = binary.Read(r.r, binary.BigEndian, data)
    }
}

func parse(input io.Reader) (*Point, error) {
var p Point
r := Reader{r: input}

r.read(&p.Longitude)
r.read(&p.Latitude)
r.read(&p.Distance)
r.read(&p.ElevationGain)
r.read(&p.ElevationLoss)

if r.err != nil {
    return nil, r.err
}

return &p, nil
}
```


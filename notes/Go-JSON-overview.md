
## Go 语言中的 JSON 简介

#### 动态内容解析

go 语言中的 json encode/decode 比较麻烦的是对动态内容的解析，需要转成 interface{} 的处理

```
data := []byte(`{"Name":"shniu","IsAdmin":true,"Followers":12,"address":"bj"}`)

// 更简便的方式
var u2 map[string]interface{}
err = json.Unmarshal(data, &u2)
assert.Nil(t, err)
fmt.Printf("%v\n", u2)

name := u2["Name"].(string)
fmt.Println(name)

// 解析过来的数字类型默认是 float64, 需要自己手动转成对应的类型
followers := int(u2["Followers"].(float64))
fmt.Println(followers)
```

#### 自定义解析方法

`encoding/json` 提供了两个接口：`Marshaler` 和 `Unmarshaler`, 如果希望自己控制怎么解析成 JSON，或者把 JSON 解析成自定义的类型，只需要实现对应的接口（interface）

比如 golang 标准库中的 time 的例子：

```go
package time

type Month struct {
    MonthNumber int
    YearNumber int
}

func (m Month) MarshalJSON() ([]byte, error){
    return []byte(fmt.Sprintf("%d/%d", m.MonthNumber, m.YearNumber)), nil
}

func (m *Month) UnmarshalJSON(value []byte) error {
    parts := strings.Split(string(value), "/")
    m.MonthNumber = strconv.ParseInt(parts[0], 10, 32)
    m.YearNumber = strconv.ParseInt(parts[1], 10, 32)

    return nil
}
```


#### 参考

- [Go 语言 JSON 简介](https://cizixs.com/2016/12/19/golang-json-guide/)
- [go and json](https://eager.io/blog/go-and-json/)
- [go dynamic json](https://eagain.net/articles/go-dynamic-json/)
- [go blog: json and go](https://blog.golang.org/json-and-go)
- [JSON 官网](https://json.org/)

![](https://www.json.org/img/object.png)
![](https://www.json.org/img/array.png)
# 编码规范

- 无错误的正常流程代码，将成为一条直线，而不是缩进代码

```
// Good
func goodDemo() {
    f, err := os.Open(path)
    if err != nil {
        // 处理错误
    }
    // 业务处理
    err = f.Chmod(f.Chmod(os.ModeAppend))
    if err != nil {
        // 处理错误
    }
    // 正常业务逻辑
}

// Bad
func badDemo() {
    f, err := os.Open(path)
    if err == nil {
        // 处理业务
    }
    // 处理异常
}
```

- 如果没有特殊的事情要做，减少不必要的判断

````
// Bad
func Auth(r *Request) error {
    err := authentiate(r.User)
    if err != nil {
        return nil
    }
    return nil
}
// Good
func Auth(r *Request) error {
    return authentiate(r.User)
}
````

- Eliminate error handling by eliminating errors (通过消除错误处理来消除错误)

```
// Demo 1
// Normal version
func CountLines(r io.Reader) (int, error) {
    br := bufio.NewReader(r)
    var lines int
    var err error
    
    for {
        _, err = br.ReadString('\n')
        lines++
        
        if err != nil {
            break
        }
    }
    
    if err != io.EOF {
        return 0, err
    }
    
    return lines, nil
}

// Better version
func CountLines(r io.Reader) (int, error) {
     s := bufio.NewScanner(r)
     lines := 0
     
     for sc.Scan() {
         lines++
     }
     return lines, sc.Err()
}
```

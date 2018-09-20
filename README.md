# gocontrol
[![Open Source Love png1](https://badges.frapsoft.com/os/v1/open-source.png?v=103)](https://github.com/ellerbrock/open-source-badges/) [![GitHub license](https://img.shields.io/github/license/ninedraft/gocontrol.svg)](https://github.com/ninedraft/gocontrol/blob/master/LICENSE) [![](https://godoc.org/github.com/ninedraft/gocontrol?status.svg)](https://godoc.org/github.com/ninedraft/gocontrol)

Micropackage for goroutine population control with dead simple API

Example:
```go
var guard = &gocontrol.Guard{}

for i := 0; i<42; i++ {
    i := i
    go func(){
        defer guard.Go()()

        fmt.Printf("START %d\n", i)
        time.Sleep(100*time.Duration(i)*time.Millisecond)
        fmt.Printf("END %d\n", i)
    }()
}

for guard.AliveN() > 0{
    time.Sleep(100*time.Millisecond)
    fmt.Println("%d goroutines are alive", guard.AliveN())
}
```

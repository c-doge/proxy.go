package base

import (
    "fmt"

    "github.com/c-doge/proxy.go/base/logger"
    "github.com/c-doge/proxy.go/base/setting"
)


func Start(path string) error {
    err := setting.Parse(path);
    if err != nil {
        return err
    }
    s := setting.Get()
    fmt.Printf("setting:\n\r%v\n", s)
    err = logger.New(s.Log.Level, s.Log.Path, "proxy.go");
    if err != nil{
        return err
    }
    return nil
}

func Stop() {

}
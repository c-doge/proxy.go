package main

import (
    "os"
    "fmt"
    "flag"

    "github.com/c-doge/proxy.go/base"
    "github.com/c-doge/proxy.go/base/logger"
)

var help bool
var configFile string

func usage() {
    fmt.Printf("proxy.go version: %s\r\n", _version)
    fmt.Printf("Usage: proxy.go [-ch]\r\n")
    fmt.Printf("           -h print this message\r\n")
    fmt.Printf("           -c config file path\r\n")
}

func init() {
    flag.BoolVar(&help,             "h", false,                "show this help")
    flag.StringVar(&configFile,     "c", "./proxy.yaml",       "config file path")
    flag.Usage = usage
}


func main() {
    fmt.Printf("proxy.go start")
    flag.Parse()
    if help {
        usage()
        os.Exit(0)
    }
    err := base.Start(configFile);
    if err != nil {
        panic(err);
    }
    logger.Infof("version:     %s\n", _version);
    logger.Infof("git branch:  %s\n", _gitBranch);
    logger.Infof("build time:  %s\n", _buildTime);
    return;
}

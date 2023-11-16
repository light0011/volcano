package main

import (
	"runtime"

	"github.com/light0011/volcano/cmd/scheduler/app"
	"github.com/light0011/volcano/cmd/scheduler/app/options"
	"github.com/spf13/pflag"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 解析命令行参数到 serverOption
	serverOption := options.NewServerOption()
	serverOption.AddFlags(pflag.CommandLine)
	
	app.Run(serverOption)

}

package app

import (
	"fmt"

	"github.com/light0011/volcano/cmd/scheduler/app/options"
	"github.com/light0011/volcano/pkg/kube"
)


func Run(opt *options.ServerOption) error {

	// 生成kube config
	config, err := kube.BuildConfig(opt.KubeClientOptions)

	fmt.Print(config)
	fmt.Print(err)
}
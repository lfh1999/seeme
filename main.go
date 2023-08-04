package main

import (
	"github.com/HyperMutekiXXX/stackman-public/serve"
	_ "github.com/lfh1999/seeme/controller"
)

func main() {
	// TODO 后续工具库增加读取配置文件，现在先写死配置
	serve.Run("127.0.0.1:8088")
}

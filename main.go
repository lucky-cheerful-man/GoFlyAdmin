package main

import (
	"fmt"
	"runtime"
	"strconv"

	"gofly/bootstrap"
	"gofly/global"
)

func main() {
	// 初始化配置
	global.App.Config.InitializeConfig()
	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("项目启动成功")
	//将对象，转换成json格式
	// data_config, err := json.Marshal(conf)

	// if err != nil {
	// 	fmt.Println("err:\t", err.Error())
	// 	return
	// }
	// fmt.Println("data_config:\t", string(data_config))
	// fmt.Println("config.Database.Driver=", conf.App.Port)
	//加载配置
	cpuNum, _ := strconv.Atoi(global.App.Config.App.CPUnum)
	mycpu := runtime.NumCPU()
	if cpuNum > mycpu { //如果配置cpu核数大于当前计算机核数，则等当前计算机核数
		cpuNum = mycpu
	}
	if cpuNum > 0 {
		runtime.GOMAXPROCS(cpuNum)
		global.App.Log.Info(fmt.Sprintf("当前计算机核数: %v个,调用：%v个", mycpu, cpuNum))
	} else {
		runtime.GOMAXPROCS(mycpu)
		global.App.Log.Info(fmt.Sprintf("当前计算机核数: %v个,调用：%v个", mycpu, mycpu))
	}

	// 启动服务器
	bootstrap.RunServer()
}

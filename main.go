package main

import (
	"fmt"
	"flag"
	// "errors"
	// "project/controllers"
	rout "project/router"
	"github.com/gin-gonic/gin"
)

var (
	configFile = flag.String("config", "conf/config.json", "config file for system")
	// config     common.IConfigInfo
)

func main()  {
	flag.Parse()
	if flag.Parsed() == false {
		flag.PrintDefaults()
		return
	}
  fmt.Printf("string=", configFile)
	// conf, err := common.LoadConfig(*configFile)
	// if err != nil {
	// 	panic("配置文件加载错误: " + err.Error())
	// }
  //
	// err = Validate(conf)
	// if err != nil {
	// 	panic("配置文件加载错误: " + err.Error())
	// }
  //
	// config = conf
  //
	// if conf.GetMode() == "release" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	// controllers.InitControllers(conf)

	router := gin.Default()
	rout.InitRouters(router)

	// fmt.Printf("%#v \r\n", conf)

	// router.Run(":" + conf.GetListeningPort())
	router.Run(":" + "7778")

}

// func Validate(conf common.IConfigInfo) error {
// 	if len(conf.GetStaticFileServer()) <= 0 {
// 		return errors.New("No StaticServerHost")
// 	}
// 	if len(conf.GetMainHost()) <= 0 {
// 		println("MainHost Not Specified")
// 		// return errors.New("No MainHost")
// 	}
// 	if len(conf.GetExamHost()) <= 0 {
// 		return errors.New("No ExamHost")
// 	}
// 	// if len(conf.GetUeConfigFile()) <= 0 {
// 	// 	return errors.New("No Ueconfig")
// 	// }
// 	if len(conf.GetListeningPort()) <= 0 {
// 		return errors.New("No listening port")
// 	}
//
// 	return nil
// }

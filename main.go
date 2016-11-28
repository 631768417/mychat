package main

import (
	"flag"
	"fmt"
	"os"

	"chat.DB"
	"chat.cluster"
	. "chat.common"
	"chat.logger"
	"chat.mysql/dao/basedao"
	daoService "chat.mysql/service"
	. "chat.protocol"
	"chat.service"
	"chat.ticker"
)

func init() {
	servername := fmt.Sprint(ProtocolversionName, " server")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("-------------------- " + servername + " ---------------------")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("----------------------------------------------------------")
}

func initbasedao() {
	if CF.Db_Exsit == 0 {
		return
	}
	logger.Debug("initbasedao")
	DB.Init()
	basedao.SetDB(DB.Master)
	basedao.SetAdapterType(basedao.MYSQL)
	gbs, err := basedao.ExecuteQuery("select 1")
	if err == nil {
		logger.Debug("test db ok", gbs[0].MapIndex(1).Value())
	}
}

func initLog(loglevel string) {
	logger.SetConsole(true)
	logger.SetRollingDaily(CF.GetLog())
	switch loglevel {
	case "debug":
		logger.SetLevel(logger.DEBUG)
	case "info":
		logger.SetLevel(logger.INFO)
	case "warn":
		logger.SetLevel(logger.WARN)
	case "error":
		logger.SetLevel(logger.ERROR)
	default:
		logger.SetLevel(logger.WARN)
	}
}

//chat f chat.xml c cluster.xml d debug
func main() {
	flag.Parse()
	wd, _ := os.Getwd()
	if flag.NArg() > 6 {
		fmt.Println("error:", "flag's length is", flag.NArg())
		os.Exit(1)
	}
	timconf := fmt.Sprint(wd, "/chat.xml")
	initconf := ""
	clusterconf := fmt.Sprint(fmt.Sprint(wd, "/cluster.xml"))
	for i := 0; i < flag.NArg(); i++ {
		if i%2 == 0 {
			switch flag.Arg(i) {
			case "f":
				timconf = flag.Arg(i + 1)
			case "c":
				clusterconf = flag.Arg(i + 1)
			case "d":
				initconf = flag.Arg(i + 1)
			default:
				fmt.Println("error:", "error arg:", flag.Arg(i))
				os.Exit(1)
			}
		}
	}
	CF.Init(timconf)
	initLog(initconf)
	cluster.InitCluster(clusterconf)
	initbasedao()
	daoService.InitDaoservice()
	ticker.TickerStart()
	service.ServerStart()
}

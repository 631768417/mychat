package common

import (
	"chat.conf"
)

/*版本*/
var VersionName = "1.0"
var VersionCode = 2
var Author = "lilinxuan"
var Email = "64180190@qq.com"

var CF = &conf.ConfBean{KV: make(map[string]string, 0), Db_Exsit: 1, MustAuth: 1}

var ClusterConf = &conf.ClusterBean{IsCluster: 1}

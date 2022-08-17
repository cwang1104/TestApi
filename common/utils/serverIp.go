package utils

import (
	"TestApi/common/config"
	"net"
)

var ServerIP = ""

func init() {
	//fixme 配置文件中ip优先，没有就获取本机有效网卡ip
	if config.ServerIp != "" {
		ServerIP = config.ServerIp
		return
	}
	//fixme 获取本机有效网卡ip
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(netInterfaces); i++ {
		//通过判断net.flagUp标志，排除无用网卡
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			address, _ := netInterfaces[i].Addrs()
			for _, addr := range address {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ServerIP = ipNet.IP.String()
						return
					}
				}
			}
		}
	}
}

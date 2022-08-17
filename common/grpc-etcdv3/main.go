package main

import (
	"TestApi/common/config"
	"TestApi/common/grpc-etcdv3/getcdv3"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func Get(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

func genRand(num int) int {
	return int(rand.Int31n(int32(num)))
}

func GetServer(client *getcdv3.EtcdClient) (string, error) {
	prefix := getcdv3.GetIpPrefix(config.LoginName)
	fmt.Println("gerserver prefix", prefix)
	servers, err := client.List(config.LoginName)
	if err != nil {
		return "", err
	}
	fmt.Println("----------", servers)

	return servers[genRand(len(servers))], nil
}

func main() {
	etcdclient, _ := getcdv3.RegistEtcdClintByip(config.EtcdAddress, config.ServerIp, config.LoginPort, config.LoginName, int64(10))
	
	//getcdv3.GetConn(config.EtcdAddress, config.LoginName)
	err := etcdclient.Connect()
	if err != nil {
		panic(err)
	}
	defer etcdclient.Close()

	for i := 0; i < 10; i++ {
		address, err := GetServer(etcdclient)
		fmt.Println("address : ", address)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data, err := Get(address)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
		time.Sleep(2 * time.Second)
	}
}

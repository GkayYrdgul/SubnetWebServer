package Models

import "fmt"

/*
Json TYPE
{
"IpAddr": "192.168.20.20",
"SubnetValue": 24,
"SubnetMaskAddr": "255.255.255.0",
"SubnetWildcard": "0.0.0.255",
"NetworkAddr": "192.168.20.0",
"BroadCastAdrr": "192.168.20.255",
"Hosts": 254
}
*/

type Jsonfile struct {
	IpAddress   string `json:"IpAddr"`
	SubValue    int    `json:"SubnetValue"`
	SubMaskAddr string `json:"SubnetMaskAddr"`
	SubWildcard string `json:"SubnetWildcard"`
	NetAddr     string `json:"NetworkAddr"`
	BrdAddr     string `json:"BroadCastAdrr"`
	HostsNet    int    `json:"Hosts"`
}

func (Jsn Jsonfile) Print() {
	fmt.Println("IP ADDRESS:", Jsn.IpAddress)
	fmt.Println("SUBNET VALUE:", Jsn.SubValue)
	fmt.Println("SUBNET MASK ADDRESS:", Jsn.SubMaskAddr)
	fmt.Println("WILDCARD:", Jsn.SubWildcard)
	fmt.Println("NETWORK ADDRESS:", Jsn.NetAddr)
	fmt.Println("BROADCAST ADDRESS:", Jsn.BrdAddr)
	fmt.Println("HOSTS:", Jsn.HostsNet)

}

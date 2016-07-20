package main

import (
	"fmt"
	"net"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func GetIpRangeByCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips[1 : len(ips)-1], nil // 去除掉网络号和广播地址
}

func GetIpStartAndEndPosition(ips []string, startIp, endIp string) (int, int) {
	var start, end int = -1, -1
	for index, ip := range ips {
		if ip == startIp {
			start = index
		}
		if ip == endIp {
			end = index
		}
		if start != -1 && end != -1 {
			break
		}
	}
	return start, end
}

func main() {
	if ips, err := GetIpRangeByCIDR("192.168.0.12/24"); err == nil {
		start, end := GetIpStartAndEndPosition(ips, "192.168.0.1", "192.168.0.254")
		fmt.Println(start, end)
		fmt.Println(ips[start : end+1])
	}
}

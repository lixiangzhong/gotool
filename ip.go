package gotool

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func IPv4toUint32(ip string) (uint32, error) {
	i := net.ParseIP(ip)
	if i == nil {
		return 0, errors.New("ParseIP error")
	}
	i = i.To4()
	return binary.BigEndian.Uint32(i), nil
}
func Uint32toIPv4(ipint uint32) string {
	a := ipint >> 24
	b := (ipint - (a << 24)) >> 16
	c := (ipint - (a << 24) - (b << 16)) >> 8
	d := ipint - (a << 24) - (b << 16) - (c << 8)
	return fmt.Sprintf("%v.%v.%v.%v", a, b, c, d)
}

//获取本机网卡IP
func GetLocalIP() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var ips []string
	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok {
			ip := ipnet.IP
			if ip.To4() != nil && !ip.IsLoopback() {
				ips = append(ips, ip.String())
			}
		}
	}
	return ips, nil
}

//将CIDR转成数字,如  1.0.0.0/24 转成 16777216 16777471
func CIDRToUint32(cidr string) (start uint32, end uint32, err error) {
	s := strings.Split(cidr, "/")
	if len(s) != 2 {
		err = errors.New(cidr + " is not CIDR")
		return
	}
	var i32 uint32 = 32
	start, err = IPv4toUint32(s[0])
	if err != nil {
		return
	}
	netbit, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return
	}
	end = 1<<(i32-uint32(netbit)) + start - 1
	return
}

/*
什么是外网IP和内网IP?

tcp/ip协议中，专门保留了三个IP地址区域作为私有地址，其地址范围如下：

 10.0.0.0/8：   10.0.0.0～10.255.255.255
172.16.0.0/12： 172.16.0.0～172.31.255.255
192.168.0.0/16：192.168.0.0～192.168.255.255

什么是内网IP

一些小型企业或者学校，通常都是申请一个固定的IP地址，然后通过IP共享（IP Sharing），使用整个公司
或学校的机器都能够访问互联网。而这些企业或学校的机器使用的IP地址就是内网IP，内网IP是在规划IPv4
协议时，考虑到IP地址资源可能不足，就专门为内部网设计私有IP地址（或称之为保留地址），一般常用内
网IP地址都是这种形式的：10.X.X.X、172.16.X.X-172.31.X.X、192.168.X.X等。需要注意的是，内网的计
算机可向Internet上的其他计算机发送连接请求，但Internet上其他的计算机无法向内网的计算机发送连接
请求。我们平时可能在内网机器上搭建过网站或者FTP服务器，而在外网是不能访问该网站和FTP服务器的，
原因就在于此。

什么是公网IP

公网IP就是除了保留IP地址以外的IP地址，可以与Internet上的其他计算机随意互相访问。我们通常所说的
IP地址，其实就是指的公网IP。互联网上的每台计算机都有一个独立的IP地址，该IP地址唯一确定互联网上
的一台计算机。这里的IP地址就是指的公网IP地址。

怎样理解互联网上的每台计算机都有一个唯一的IP地址

其实，互联网上的计算机是通过“公网IP＋内网IP”来唯一确定的，就像很多大楼都是201房间一样，房间号可
能一样，但是大楼肯定是唯一的。公网IP地址和内网IP地址也是同样，不同企业或学校的机器可能有相同的
内网IP地址，但是他们的公网IP地址肯定不同。那么这些企业或学校的计算机是怎样IP地址共享的呢？这就
需要使用NAT（Network Address Translation,网络地址转换）功能。当内部计算机要连接互联网时，首先需
要通过NAT技术，将内部计算机数据包中有关IP地址的设置都设成NAT主机的公共IP地址，然后再传送到
Internet，虽然内部计算机使用的是私有IP地址，但在连接Internet时，就可以通过NAT主机的NAT技术，
将内网我IP地址修改为公网IP地址，如此一来，内网计算机就可以向Internet请求数据了。

————————————————

What Is a Network Interface?

A network interface is the point of interconnection between a computer and a private or public
network. A network interface is generally a network interface card (NIC), but does not have to 
have a physical form. Instead, the network interface can be implemented in software. For example, 
the loopback interface (127.0.0.1 for IPv4 and ::1 for IPv6) is not a physical device but a piece 
of software simulating a network interface. The loopback interface is commonly used in test 
environments.

获取本地ip
*/
package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}

func externalIP() (string, error) {
    // get all interface.
    ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			fmt.Println(" ip.String()", ip.String())
			// return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

// 获取外网ip

func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//s := buf.String()
	return string(content)
}

// 10进制与ip相互转换

func inet_ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func inet_aton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

// 判断是否公网ip

func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// ip在区间内

func IpBetween(from net.IP, to net.IP, test net.IP) bool {
	if from == nil || to == nil || test == nil {
		fmt.Println("An ip input is nil") // or return an error!?
		return false
	}

	from16 := from.To16()
	to16 := to.To16()
	test16 := test.To16()
	if from16 == nil || to16 == nil || test16 == nil {
		fmt.Println("An ip did not convert to a 16 byte") // or return an error!?
		return false
	}

	if bytes.Compare(test16, from16) >= 0 && bytes.Compare(test16, to16) <= 0 {
		return true
	}
	return false
}

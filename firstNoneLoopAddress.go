package libnet

import (
    "fmt"
    "net"
)

func FirstNoneLoopAddress() (string, error) {
    ip, err := firstNoneLoopAddress()
    if err != nil {
        return "", err 
    }   
    return ip.String(), nil 
}

func firstNoneLoopAddress() (net.IP, error) {

    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err 
    }   
    addresses := []net.IP{}
    for _, iface := range ifaces {

        if iface.Flags&net.FlagUp == 0 { 
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 { 
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            continue
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
            addresses = append(addresses, ip) 
        }
    }   
    if len(addresses) == 0 { 
        return nil, fmt.Errorf("no address Found, net.InterfaceAddrs: %v", addresses)
    }   
    //  fmt.Println(addresses)
    //only need first
    return addresses[0], nil 
}

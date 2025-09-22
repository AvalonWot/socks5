package socks5

import (
	"net"
)

var Debug bool

func init() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var Resolve func(network string, addr string) (net.Addr, error) = func(network string, addr string) (net.Addr, error) {
	if network == "tcp" || network == "tcp4" || network == "tcp6" {
		return net.ResolveTCPAddr(network, addr)
	}
	return net.ResolveUDPAddr(network, addr)
}

var DialTCP func(network string, laddr, raddr string) (net.Conn, error) = func(network string, laddr, raddr string) (net.Conn, error) {
	var la, ra *net.TCPAddr
	if laddr != "" {
		var err error
		la, err = net.ResolveTCPAddr(network, laddr)
		if err != nil {
			return nil, err
		}
	}
	a, err := Resolve(network, raddr)
	if err != nil {
		return nil, err
	}
	ra = a.(*net.TCPAddr)
	return net.DialTCP(network, la, ra)
}

var DialUDP func(network string, laddr, raddr string) (net.Conn, error) = func(network string, laddr, raddr string) (net.Conn, error) {
	var la, ra *net.UDPAddr
	if laddr != "" {
		var err error
		la, err = net.ResolveUDPAddr(network, laddr)
		if err != nil {
			return nil, err
		}
	}
	a, err := Resolve(network, raddr)
	if err != nil {
		return nil, err
	}
	ra = a.(*net.UDPAddr)
	return net.DialUDP(network, la, ra)
}

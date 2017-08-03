package collectors

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"

	"github.com/iwondory/udpserver/event"
)

type Collector interface {
	Start(chan<- *event.Event) error
	Addr() net.Addr
}

func NewCollector(proto, iface, format string, tlsConfig *tls.Config) (Collector, error) {
	if strings.ToLower(proto) == "tcp" {
		//		return &TCPCollector{
		//			iface:     iface,
		//			format:    format,
		//			tlsConfig: tlsConfig,
		//		}, nil
	} else if strings.ToLower(proto) == "udp" {
		addr, err := net.ResolveUDPAddr("udp", iface)
		if err != nil {
			return nil, err
		}

		return &UDPCollector{addr: addr, format: format}, nil
	}
	return nil, fmt.Errorf("unsupport collector protocol")
}

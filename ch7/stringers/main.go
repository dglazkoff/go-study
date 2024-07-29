//package main
//
//import (
//	"fmt"
//)
//
//type IPAddr [4]byte
//
//func (addr IPAddr) String() string {
//	var s string
//
//	for i, b := range addr {
//		if i == 0 {
//			s = fmt.Sprint(b)
//			continue
//		}
//
//		s += "." + fmt.Sprint(b)
//	}
//	return s
//}
//
//// TODO: Add a "String() string" method to IPAddr.
//
//func main() {
//	hosts := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for name, ip := range hosts {
//		fmt.Printf("%v: %v\n", name, ip)
//	}
//}

package main

import (
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	r.r.Read(*p)

	for i, b := range p {
		p[i] = rot13(b)
	}

	return len(p), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

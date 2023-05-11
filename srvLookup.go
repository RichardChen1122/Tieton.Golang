package main

import (
	"fmt"
	"net"
)

func srvLookup() {
	cname, srvs, err := net.LookupSRV("alt", "tcp", "a-test-junbchen-srv.azconfig-test.io")
	if err != nil {
		panic(err)
	}
 
	fmt.Printf("\ncname: %s \n\n", cname)
 
	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func p(printObject ...interface{}) {
	_, _ = fmt.Fprintln(os.Stdout, printObject...)
}
func strToInt(stra string) (int, error) {
	var num int
	var err error
	num, err = strconv.Atoi(stra)
	if err != nil {
		p(err)
		return 0, err
	}
	return num, nil
}
func dirExist(_path string) bool {
	var err error
	var filestat os.FileInfo
	filestat, err = os.Stat(_path)
	if err != nil {
		return false
	}
	return filestat.IsDir()
}

func main() {
	var osargs = os.Args
	var argName string
	var argValue string
	var thisArg string
	var ipaddress = "127.0.0.1"
	var port = "80"
	//http.Dir(".")
	var workdir = "/var/www/html"
	for i := 1; i < len(osargs); i++ {
		thisArg = osargs[i]
		if strings.Index(thisArg, "--") != 0 {
			p("input arg '" + thisArg + "' prefix is not '--'!")
			os.Exit(1)
		}
		if strings.Index(thisArg, "=") == -1 {
			p("input arg '" + thisArg + "' is not like '--arg=value'!")
			os.Exit(1)
		}
		argName = thisArg[2:strings.Index(thisArg, "=")]
		argValue = thisArg[strings.Index(thisArg, "=")+1:]
		if argName == "ip" {
			if net.ParseIP(argValue) == nil {
				p("input arg ip is '" + argValue + "', is not like '1.1.1.1'!")
				os.Exit(1)
			}
			ipaddress = argValue
		} else if argName == "port" {
			portInt, err := strToInt(argValue)
			if err != nil || portInt < 0 || portInt > 65535 {
				p("input arg port is '" + thisArg + "', is not in 0~65535")
				os.Exit(1)
			}
			port = argValue
		} else if argName == "dir" {
			if !dirExist(argValue) {
				p("input arg dir is '" + thisArg + "', is not a directory!")
				os.Exit(1)
			}
			workdir = argValue
		}
	}
	http.Handle("/", http.FileServer(http.Dir(workdir)))

	err := http.ListenAndServe(ipaddress+":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}

}

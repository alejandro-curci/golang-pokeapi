package args

import "flag"

var Port string

func ParseArgs() {
	flag.StringVar(&Port, "port", "8080", "application listen port")
	flag.Parse()
}

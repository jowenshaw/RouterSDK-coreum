package main

import (
	"flag"
	"fmt"

	"github.com/anyswap/CrossChain-Router/v3/log"
	routersdk "github.com/anyswap/RouterSDK-coreum/sdk"
)

var (
	paramPublicKey string
	paramPrefix    string
)

func initFlags() {
	flag.StringVar(&paramPublicKey, "p", "", "publicKey")
	flag.StringVar(&paramPrefix, "prefix", "testcore", "prefix, eg. cosmos, sei, etc.")

	flag.Parse()
}

func main() {
	initFlags()

	if addr, err := routersdk.PublicKeyToAddress(paramPrefix, paramPublicKey); err != nil {
		log.Fatalf("err: %v\n", err)
	} else {
		fmt.Printf("addr: %v\n", addr)
	}
}

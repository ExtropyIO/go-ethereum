package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	log15 "github.com/ethereum/go-ethereum/log"
)

var ()

var logger = log15.New()

const extraPrefix = 32 + 1 // `1` is for vote

func main() {
	logger.SetHandler(
		log15.LvlFilterHandler(log15.LvlDebug,
			log15.StreamHandler(os.Stdout, log15.TerminalFormat(true))))

	flag.Parse()

	// the addresses can be optionally prefixed with "0x"
	var addrStrs = []string{
		"6bbc9092b4b21cf68d81cb4b5527965486a32434",
		"0xdcdc1a58c2666e230f8566ca350b8e6eded163e8",
	}
	fmt.Println("extraData:", createExtraData(addrStrs))
}

func createExtraData(addrs []string) string {
	addrs = normalize(addrs)
	// we multiple by 2 because 1 byte is encoded using 2hex
	var extra = []string{"0x", strings.Repeat("0", 2*extraPrefix)}
	extra = append(extra, addrs...)
	return strings.Join(extra, "")
}

// normalize asserts that all addresses are valid and removes 0x prefix if needed
func normalize(ads []string) []string {
	var out = make([]string, len(ads))
	var ok = true
	for i, s := range ads {
		if !common.IsHexAddress(s) {
			ok = false
			logger.Error("Found invalid address", "addr_index", i, "value", s)
		} else {
			out[i] = remove0x(s)
		}
	}
	if !ok {
		logger.Crit("to continue fix invalid addresses")
	}
	return out
}

func remove0x(s string) string {
	return strings.TrimPrefix(s, "0x")
}

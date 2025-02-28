package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

const offset = 0x26DB27 //where the default config is located
const maxSize = 0x44D   //safe size so i don't overwrite anything else

func usage() error {
	return errors.New("Usage: <config_to_inject.cfg> <zip0_xbox.xzp>")
}

func inject(in []byte, orig []byte) []byte {
	return bytes.Join(
		[][]byte{
			bytes.Clone(orig[:offset]),
			in,
			bytes.Clone(orig[offset+maxSize:]),
		},

		[]byte{},
	)
}

func open(path string) (*os.File, int64) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return file, stat.Size()
}

func main() {
	if len(os.Args) <= 2 {
		err := usage()
		log.Fatal(err)
	}

	cfg, cfgSize := open(os.Args[1])
	defer cfg.Close()

	orig, origSize := open(os.Args[2])
	defer orig.Close()

	if cfgSize > maxSize {
		decSize := fmt.Sprintf("%d", maxSize)
		err := errors.New("config file too large, max size " + decSize + " bytes")
		log.Fatal(err)
	}

	cfgData := make([]byte, maxSize)
	cfg.Read(cfgData)

	origData := make([]byte, origSize)
	orig.Read(origData)

	modified := inject(cfgData, origData)

	xzp, err := os.Create("modified.xzp")
	if err != nil {
		log.Fatal(err)
	}
	defer xzp.Close()

	xzp.Write(modified)

	modStat, err := xzp.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if modStat.Size() != origSize {
		err = errors.New("modified file not correct size")
		log.Fatal(err)
	}

	fmt.Println(string(modified[offset : offset+maxSize]))

	log.Println("now just rename modified.xzp to zip0_xbox.xzp :3")
	log.Println("and keep the original somewhere safe")
}

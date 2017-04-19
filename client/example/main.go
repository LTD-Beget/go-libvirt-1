package main

import (
	"image/png"
	"log"
	"net"
	"os"
	"time"

	"github.com/spakin/netpbm"
	client "github.com/vtolstov/go-libvirt/client"
)

func main() {
	c, err := net.DialTimeout("tcp", os.Args[1], 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}
	l := client.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	domain, err := l.DomainLookupByName("143177")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	stream, _, err := domain.Screenshot(0, 0)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	defer stream.Close()
	img, err := netpbm.Decode(stream, &netpbm.DecodeOptions{
		Target: netpbm.PPM,
		Exact:  true})
	if err != nil {
		log.Fatalf("failed: %#+v %v", img, err)
	}

	f, err := os.Create("img.png")
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	if err = l.Close(); err != nil {
		panic(err)
	}
}

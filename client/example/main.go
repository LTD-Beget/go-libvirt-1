package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	libvirt "github.com/vtolstov/go-libvirt"
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

	v, err := l.Version()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt version: %v", err)
	}
	fmt.Println("Version:", v)

	domains, err := l.ListAllDomains()
	if err != nil {
		log.Fatalf("failed to retrieve domains: %v", err)
	}

	fmt.Println("ID\tName\t\tUUID")
	fmt.Printf("--------------------------------------------------------\n")
	for _, d := range domains {
		fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
	}

	pool, err := l.StoragePoolLookupByName("sda")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	volsrc, err := pool.StorageVolumeLookupByName("132859")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	voldst, err := pool.StorageVolumeLookupByName("test")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	src, err := l.StreamNew()
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	defer src.Close()
	err = volsrc.Download(src, 0, 0, 0)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	dst, err := l.StreamNew()
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	err = voldst.Upload(dst, 0, 0, 0)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	defer dst.Close()

	var n int64
	buf := make([]byte, libvirt.NetMessageLegacyPayloadMax)
Loop:
	for {
		s, serr := src.Read(buf)
		if serr != nil && serr != io.EOF {
			fmt.Printf("src %s\n", serr.Error())
			break Loop
		}
		for {
			var d int
			c, err := dst.Write(buf[d:s])
			n += int64(c)
			d += c
			if err != nil {
				fmt.Printf("dst %s\n", err.Error())
				time.Sleep(2 * time.Second)
				continue
			} else {
				break
			}
		}
		if serr == io.EOF {
			break Loop
		}
	}
	fmt.Printf("close after %d\n", n)
	if err := l.Disconnect(); err != nil {
		log.Fatal("failed to disconnect: %v", err)
	}
	fmt.Printf("disconnected\n")
}

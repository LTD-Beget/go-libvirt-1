package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

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

	volsrc, err := pool.StorageVolumeLookupByName("143177")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	x := `<volume type='file'>
    <name>test</name>
    <allocation unit='bytes'>945627136</allocation>
    <target>
      <format type='qcow2'/>
    </target>
    </volume>`

	volnew, err := pool.StorageVolumeCreateXMLFrom(x, volsrc, 0)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	fmt.Printf("%#+v\n", volnew)

	if err := l.Disconnect(); err != nil {
		log.Fatal("failed to disconnect: %v", err)
	}
	fmt.Printf("disconnected\n")
}

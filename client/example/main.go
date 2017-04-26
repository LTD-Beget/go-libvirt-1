package main

import (
	"fmt"
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

	evtchan := make(chan *libvirt.Event)
	cids, err := l.DomainEventRegisterAny(nil, evtchan)
	if err != nil {
		log.Fatalf("failed : %v", err)
	}
	tick := time.Tick(10 * time.Second)
Loop:
	for {
		select {
		case <-tick:
			l.DomainEventDeregisterAny(cids)
			break Loop
		case evt := <-evtchan:
			fmt.Printf("%#+v\n", evt)
		}
	}
	/*
		domain, err := l.DomainLookupByName("143177")
		if err != nil {
			log.Fatalf("failed: %v", err)
		}

		stream, err := domain.Console("serial0", libvirt.DomainConsoleFlagForce)
		if err != nil {
			log.Fatalf("failed: %v", err)
		}

		conreader := bufio.NewReader(os.Stdin)
		r := bufio.NewReader(stream)
		w := bufio.NewWriter(stream)
		rw := bufio.NewReadWriter(r, w)

		term := terminal.NewTerminal(rw, "> ")
		go func() {
			for {
				text, _ := conreader.ReadString('\n')
				_, err := term.Write([]byte(text))
				w.Flush()
				if err != nil {
					fmt.Printf("err %s\n", err.Error())
				}
			}
		}()

		go func() {
			for {
				text, err := term.ReadLine()
				if err != nil {
					time.Sleep(1 * time.Second)
					//fmt.Printf("err %s\n", err.Error())
				}
				if text != "" {
					fmt.Printf("%s", text)
				}
			}
		}()

		select {}
		/*
			stream, _, err := domain.Screenshot(0, 0)
			if err != nil {
				log.Fatalf("failed: %v", err)
			}
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

	*/
	/*
		stream.Close()
	*/
	if err = l.Close(); err != nil {
		panic(err)
	}
}

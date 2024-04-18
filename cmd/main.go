package main

import (
	com "github.com/slinky55/KioskOrderProcessor"
	"go.bug.st/serial"
	"log"
	"time"
)

var ports = make(map[string]*com.PortInterface, 12)

var DefaultSerialMode = &serial.Mode{
	BaudRate: 115200,
}

func main() {
	// Open serial ports

	portList, err := serial.GetPortsList()
	if err != nil {
		log.Fatalln("Failed to get COM port list: ", err.Error())
	}

	for _, name := range portList {
		port, err := serial.Open(name, DefaultSerialMode)
		if err != nil {
			log.Println("Failed to open port ", name)
			continue
		}

		_, err = port.Write([]byte("IDENT"))
		if err != nil {
			log.Println("Failed to ident on port ", name)
			continue
		}

		buff := make([]byte, 8)
		for {
			err := port.SetReadTimeout(time.Second * 5)
			if err != nil {
				log.Println("Failed to set read timeout on port ", name)
				log.Println(err)
				break
			}
			n, err := port.Read(buff)
			if err != nil {
				log.Println("Failed to read ident response on port ", name)
				log.Println(err)
				break
			}

			if n == 0 {
				log.Println("No response received on port ", name)
				break
			}

			slot := string(buff[:n])

			inf := com.NewPortInterface(&port)
			ports[slot] = inf

			log.Printf("Found slot %s on port %s", slot, name)
		}
	}

}

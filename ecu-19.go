package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/distributed/sers"
)

var (
	ecu19SpecificInitCommand = []byte {0x7C}

	ecu19WokeResponse = []byte {0x55, 0x76, 0x83}
  ecu19SpecificInitResponse = []byte {ecu19SpecificInitCommand[0], 0xE9} // includes our echo
)

func readFirstBytesFromPortEcu19(fn string) ([]byte, error) {

	fmt.Println("Connecting to MEMS 1.9 ECU - RW")
	globalConnected = false

	sp, err := sers.Open(fn)
	if err != nil {
		return nil, err
	}
	defer sp.Close()

	err = sp.SetMode(9600, 8, sers.N, 1, sers.NO_HANDSHAKE)
	if err != nil {
		return nil, err
	}

	// setting:
	// minread = 0: minimal buffering on read, return characters as early as possible
	// timeout = 1.0: time out if after 1.0 seconds nothing is received
	err = sp.SetReadParams(0, 0.001)
	if err != nil {
		return nil, err
	}

	mode, err := sp.GetMode()
	fmt.Println("Serial cable set to:")
	fmt.Println(mode)

	fmt.Println( "try the normal method first")
	ecu1xLoop(sp, true)

  // clear the line
	sp.SetBreak(false)
	time.Sleep(2000 * time.Millisecond)

  start := time.Now()

  fmt.Printf("Time A: %s\n", time.Now());
  // start bit
	sp.SetBreak(true)
	//fmt.Println("break 1")
  sleepUntil(start, 200)

  // send the byte
  ecuAddress := 0x16
  for i:=0; i<8; i++ {

    bit := (ecuAddress >> i) & 1;
    if (bit > 0) {
        sp.SetBreak(false)
		//fmt.Println("break 1")
    } else {
        sp.SetBreak(true)
		//fmt.Println("break 0")
    }

    sleepUntil(start, 200 + ((i+1)*200))

  }
  // stop bit
	sp.SetBreak(false)
	//fmt.Println("break 0")
	//fmt.Printf("Time B - before stop bit: %s\n", time.Now());
  sleepUntil(start, 200 + (8*200) + 200)
  fmt.Printf("Time C - after stop bit: %s\n", time.Now());
	buffer := make([]byte, 0)

	readLoops := 0
	readLoopsLimit := 200
	for readLoops < readLoopsLimit {
		//fmt.Printf("Time D - after stop bit: %s, %d", time.Now(), readLoops);
		readLoops++
		if readLoops > 1 {
			time.Sleep(10 * time.Millisecond)
		}

		rb := make([]byte, 128)
		n, _ := sp.Read(rb[:])
		rb = rb[0:n] // chop down to actual data size
		buffer = append(buffer, rb...)
		if n > 0 {
			readLoops = 0 // reset timeout
		}

		// clear leading zeros (from our wake up)

		
		for len(buffer) > 0 && buffer[0] == 0x00 {
			fmt.Printf("Time D - Dropping 0x00\n", time.Now());
			fmt.Printf("\n")
			buffer = buffer[1:]
		}

		if len(buffer) == 0 { continue }
		
		if slicesEqual(buffer, ecu19WokeResponse) {
      fmt.Println("1.9 ECU woke up - init stage 1!")
	  fmt.Printf("Time F\n", time.Now());

	  fmt.Printf("RW buffer a data: got %d bytes \n%s\n", len(buffer), hex.Dump(buffer))
	  fmt.Printf("RW ecu19WokeResponse data: got %d bytes \n%s\n", len(ecu19WokeResponse), hex.Dump(ecu19WokeResponse))

	  
			buffer = nil
			time.Sleep(50 * time.Millisecond) // TODO: is this the right sleep?
      // todo: invert (xor) byte 2 (x83) and send back to ecu
      // 0x83, 1000 0011 -> 0x7C 0111 1100
      // doing manually for now (doesn't hurt)
      sp.Write(ecu19SpecificInitCommand)
			continue
		}

		
    if slicesEqual(buffer, ecu19SpecificInitResponse) {
		fmt.Printf("Time H\n", time.Now());
		fmt.Printf("RW buffer b data: got %d bytes \n%s", len(buffer), hex.Dump(buffer))
		fmt.Printf("RW ecu19SpecificInitResponse data: got %d bytes \n%s\n", len(ecu19SpecificInitResponse), hex.Dump(ecu19SpecificInitResponse))		


      fmt.Println("1.9 ECU init stage 2!")
	  fmt.Printf("RW buffer data: got %d bytes \n%s\n", len(buffer), hex.Dump(buffer))
      buffer = nil
      ecu1xLoop(sp, true)
      continue
    }

	}
	if readLoops >= readLoopsLimit {
		fmt.Printf("1.9 had buffer data: got %d bytes \n%s\n", len(buffer), hex.Dump(buffer))
		return nil, errors.New("MEMS 1.9 timed out")
	}
	fmt.Println("fell out of 1.9 readloop")

	return nil, err
}

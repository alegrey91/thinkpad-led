package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

/**
----------------------------
thinkpad-led -- by alegrey91
This little tool allow you to
manage the red back led of
your Thinkpad.
The allowed commands are:
on → to enable the led
off → to disable the led
blink → to make it blink
----------------------------
*/

const programName string = "thinkpad-led"
const ledSystemFile string = "/sys/kernel/debug/ec/ec0/io"
const offset int64 = 12
const version string = "1.0.1"

/**
Write binary instruction into the dedicated file
to manage red back led of the thinkpad.
*/
func writeToLed(file *os.File, data *[]byte) (bool, error) {
	err := binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		return false, fmt.Errorf("Error with binary.Write: %s", err)
	}
	return true, nil
}

/**
Prepare the right instruction to
turn ON the red back led.
*/
func doOn(file *os.File) (bool, error) {
	data := []byte{0x8a}
	return writeToLed(file, &data)
}

/**
Prepare the right instruction to
turn OFF the red back led.
*/
func doOff(file *os.File) (bool, error) {
	data := []byte{0x0a}
	return writeToLed(file, &data)
}

/**
Prepare the right instruction to
turn BLINK the red back led.
*/
func doBlink(file *os.File) (bool, error) {
	data := []byte{0xca}
	return writeToLed(file, &data)
}

func main() {
	// Input management.
	allowedArguments := []string{"on", "off", "blink", "help", "version"}
	if len(os.Args) < 2 {
		log.Fatalf("Not enough arguments.\n")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		log.Fatalf("Too much arguments.\n")
		os.Exit(1)
	}

	// Check an allowed argument has been passed.
	var argumentFound bool = false
	for _, argument := range allowedArguments {
		if os.Args[1] == argument {
			argumentFound = true
		}
	}
	if !argumentFound {
		log.Fatalf("Wrong argument passed. Type: './%s help' to get help.\n", programName)
		os.Exit(1)
	}
	allowedArgument := os.Args[1]

	// Check for led file existance.
	if _, err := os.Stat(ledSystemFile); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist. Try to load kernel module: sudo modprobe ec_sys write_support=1\n", ledSystemFile)
		os.Exit(1)
	}

	// Open file in WRITE ONLY mode, maintaining the right permissions.
	file, err := os.OpenFile(ledSystemFile, os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Error with os.Open: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Position the writing to the right offset.
	if _, err := file.Seek(offset, 0); err != nil {
		log.Fatalf("Error with file.Seek: %s\n", err)
		os.Exit(1)
	}

	// Parse the input flag to do the right action.
	switch action := allowedArgument; action {
	case "on":
		if _, err := doOn(file); err != nil {
			log.Fatalf("Error with binary.Write in doOn function: %s\n", err)
		}
		os.Exit(0)
	case "off":
		if _, err := doOff(file); err != nil {
			log.Fatalf("Error with binary.Write in doOff function: %s\n", err)
		}
		os.Exit(0)
	case "blink":
		if _, err := doBlink(file); err != nil {
			log.Fatalf("Error with binary.Write in doBlink function: %s\n", err)
		}
		os.Exit(0)
	case "help":
		fmt.Printf("Usage: # %s c̲o̲m̲m̲a̲n̲d̲\nAvailable commands:\non → to enable the led\noff → to disable the led\nblink → to make it blink\nhelp → show this page\nversion → show the version\n", programName)
		os.Exit(0)
	case "version":
		fmt.Printf("%s: v%s\n", programName, version)
		os.Exit(0)
	default:
		log.Fatal("Wrong argument.\n")
		os.Exit(1)
	}
}

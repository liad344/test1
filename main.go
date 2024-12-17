package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cancel := make(chan os.Signal)
	r, w := io.Pipe()
	config := readConfigNoViper("./config.toml")
	go slowWrite(w, config)

	read := Read(r)
	go func() {
		for r := range read {
			fmt.Println(string(r))
		}
	}()
	signal.Notify(cancel, os.Interrupt, syscall.SIGTERM)
	fmt.Printf("waiting for ctrl+C\n")
	<-cancel
	fmt.Printf("got for ctrl+C\n")

}

func readConfigViper() {
	viper.

}

func readConfigNoViper(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("could not read file %v\n", err)
		return ""
	}
	fmt.Printf("loaded file %v\n", string(file))
	type config struct {
		Name string `toml:"name"`
	}
	var c config
	_, err = toml.Decode(string(file), &c)
	if err != nil {
		fmt.Printf("could not decode toml %v\n", err)
		return ""
	}
	fmt.Printf("loaded with config %v\n", c)
	return c.Name
}
func slowWrite(w io.Writer, config string) {
	for _ = range time.NewTicker(time.Second).C {
		_, err := w.Write([]byte(config))
		if err != nil {
			fmt.Printf("could not write %v", err)
		}
	}
}

func Read(r io.Reader) chan []byte {
	buffer := make([]byte, 4)
	read := make(chan []byte)

	go func() {
		for {
			_, err := r.Read(buffer)
			if err != nil {
				fmt.Printf("could not read %v", err)
			}
			read <- buffer

		}

	}()

	return read
}

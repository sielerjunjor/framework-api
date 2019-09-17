package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
	Collection string
}

// Read and parse the configuration file
func (c *Config) Read() {
	fmt.Printf("%s\n", c.Server)
	file, err := os.Open("config.json")
	if err != nil {  log.Fatal(err) }

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {  log.Fatal(err) }

	fmt.Printf("asd %s\n", c.Server)
}
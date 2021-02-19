package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	"gopkg.in/ini.v1"
)

var (
	file    string
	section string
	key     string
	pretty  bool
)

func init() {
	flag.StringVarP(&file, "file", "f", "", "Path to INI file")
	flag.StringVarP(&section, "section", "s", "", "INI section to get")
	flag.StringVarP(&key, "key", "k", "", "INI key to get")
	flag.BoolVarP(&pretty, "pretty", "p", false, "Pretty output")
	flag.Parse()
	if !pretty {
		ini.PrettyFormat = false
		ini.PrettySection = false
	}
}

func main() {
	if file == "" {
		log.Fatal("--file is required")
	}
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatal(err)
	}
	if section == "" {
		cfg.WriteTo(os.Stdout)
		os.Exit(0)
	}
	s := cfg.Section(section)
	if key != "" {
		fmt.Printf(s.Key(key).String())
		os.Exit(0)
	}
	tmpCfg := ini.Empty()
	*tmpCfg.Section(section) = *s
	tmpCfg.WriteTo(os.Stdout)
}

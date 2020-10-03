package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"os/user"
)

type listEntry struct {
	ServerURL  string
	SecretType string
}

type fullEntry struct {
	ServerURL  string
	SecretType string
	Secret     string
}

func version() {
	fmt.Printf("version 0.1\n")
}

func list(cfg *ini.File) {
	host := cfg.Section("client").Key("host").String()
	if len(host) < 1 {
		host = "localhost"
	}
	entry := listEntry{
		cfg.Section("client").Key("user").String() + "@" + host,
		"password",
	}
	j, err := json.Marshal([]listEntry{entry})
	if err != nil {
		fmt.Printf("Fail to encode json %v", err)
		os.Exit(1)
	}
	println(string(j))
}

func get(cfg *ini.File) {
	host := cfg.Section("client").Key("host").String()
	if len(host) < 1 {
		host = "localhost"
	}
	entry := fullEntry{
		cfg.Section("client").Key("user").String() + "@" + host,
		"password",
		cfg.Section("client").Key("password").String(),
	}
	j, err := json.Marshal(entry)
	if err != nil {
		fmt.Printf("Fail to encode json %v", err)
		os.Exit(1)
	}
	println(string(j))
}

func main() {
	usr, _ := user.Current()

	cfg, err := ini.LoadSources(ini.LoadOptions{AllowBooleanKeys: true},
		usr.HomeDir+"/.my.cnf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	switch command {
	case "version":
		version()
		break
	case "list":
		list(cfg)
		break
	case "get":
		get(cfg)
		break
	default:
		fmt.Printf("Unknown or unsupported command: '%s', try 'version', 'list' or 'get'\n", command)
		os.Exit(1)
	}
}

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mdh67899/openfalcon-flume-monitor/model"
)

func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}

func ParseConfig(filename string) *model.Cfg {
	if filename == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	_, err := os.Stat(filename)
	if err != nil {
		log.Fatalln("config file", filename, ", failed:", err, ", try to mv cfg.example.yaml to cfg.yaml")
	}

	var cfg = model.Cfg{}

	configContent, err := ToTrimString(filename)
	if err != nil {
		log.Fatalln("read config file:", filename, "fail:", err)
	}

	err = json.Unmarshal([]byte(configContent), &cfg)
	if err != nil {
		log.Fatalln("parse config file:", filename, "fail:", err)
	}

	log.Println("read config file:", filename, "successfully")

	return &cfg
}

package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	configMap map[string]string
}

func (config *Config) Get(key string) string {
	return config.configMap[key]
}
func loadConfig() map[string]string {
	configMap := make(map[string]string)
	f, err := os.Open("./config.ini")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		f.Close()
		log.Println("============================loadConfig========================================================")
		for s := range configMap {
			log.Printf("%s=%s", s, configMap[s])
		}
		log.Println("==============================================================================================")
	}()

	reader := bufio.NewReader(f)
	for {
		readString, err := reader.ReadString('\n')
		index := strings.Index(readString, "=")
		key := readString[:index]
		val := readString[index+1:]
		configMap[strings.TrimSpace(key)] = strings.TrimSpace(val)
		if err != nil {
			if err == io.EOF {
				return configMap
			}
		}
	}
}

func LoadConfig() *Config {
	return &Config{configMap: loadConfig()}
}

var CONFIG = LoadConfig()

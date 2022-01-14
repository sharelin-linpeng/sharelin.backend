package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func loadConfig() map[string]string {
	configMap := make(map[string]string)
	f, err := os.Open("./src/config/config.ini")
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
		split := strings.Split(readString, "=")
		configMap[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
		if err != nil {
			if err == io.EOF {
				return configMap
			}
		}
	}
}

var config = loadConfig()

func Config(key string) string {
	return config[key]
}

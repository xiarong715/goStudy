package conf

import (
	"log"
	"os"
)

type Config struct {
	DataPath   string
	PathPrefix string
}

var Conf Config = Config{}

func fromEnv() error {
	dataPathEnv := os.Getenv("DATA_PATH")
	if dataPathEnv != "" {
		Conf.DataPath = dataPathEnv
	}

	pathPrefixEnv := os.Getenv("PATH_PREFIX")
	if pathPrefixEnv != "" {
		Conf.PathPrefix = pathPrefixEnv
	}
	return nil
}

var runnerAfterConfInit []func() error

func RunAfterConfInit(f func() error) {
	runnerAfterConfInit = append(runnerAfterConfInit, f)
}

func InitConf() error {
	log.Println("init conf")
	if err := fromEnv(); err != nil {
		return err
	}

	for _, runner := range runnerAfterConfInit {
		if err := runner(); err != nil {
			return err
		}
	}
	return nil
}

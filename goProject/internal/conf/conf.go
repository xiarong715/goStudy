package conf

import "log"

type security struct {
}

type server struct {
}

type Config struct {
	Security security
	Server   server
}

var C = Config{
	// default
	Security: security{},
	Server:   server{},
}

var runnerAfterConf []func() error

func fromEnv() error { return nil }

func RunAfterConfInit(f func() error) {
	runnerAfterConf = append(runnerAfterConf, f)
}

func InitConfig() error {
	if err := fromEnv(); err != nil {
		log.Fatal("form env:", err)
	}

	for _, f := range runnerAfterConf {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

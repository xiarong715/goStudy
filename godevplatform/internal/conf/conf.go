package conf

import (
	"log"
	"os"
)

type cookie struct {
	CookieName  string
	CookieValue string
	CookiePath  string
}

type Config struct {
	DataPath   string
	PathPrefix string
	Cookie     cookie
	Domain     string
}

var Conf Config = Config{
	Cookie: cookie{CookieName: "lsy", CookieValue: "hello_my_soom_to_2", CookiePath: "/"},
	Domain: "localhost",
}

func fromEnv() error {
	dataPathEnv := os.Getenv("DATA_PATH")
	if dataPathEnv != "" {
		Conf.DataPath = dataPathEnv
	}

	pathPrefixEnv := os.Getenv("PATH_PREFIX")
	if pathPrefixEnv != "" {
		Conf.PathPrefix = pathPrefixEnv
	}

	cookieNameEnv := os.Getenv("COOKIE_NAME")
	if cookieNameEnv != "" {
		Conf.Cookie.CookieName = cookieNameEnv
	}

	cookieValueEnv := os.Getenv("COOKIE_VALUE")
	if cookieValueEnv != "" {
		Conf.Cookie.CookieValue = cookieValueEnv
	}

	cookiePathEnv := os.Getenv("COOKIE_PATH")
	if cookiePathEnv != "" {
		Conf.Cookie.CookiePath = cookiePathEnv
	}

	domainEnv := os.Getenv("DOMAIN")
	if domainEnv != "" {
		Conf.Domain = domainEnv
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

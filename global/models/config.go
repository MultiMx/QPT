package models

import "github.com/Mmx233/config"

type Config struct {
	config.Config
	HttpProxy  string
	HttpsProxy string
}

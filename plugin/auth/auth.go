package auth

import oauth "github.com/breezymind/comqtt/server/listeners/auth"

type Auth interface {
	Open() error
	Close()
	oauth.Controller
}

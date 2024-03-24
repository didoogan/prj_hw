package route

import (
	"maps/transport"
)

type EuroRoute struct {
	transports []transport.Transport
}

func (er *EuroRoute) AddTransport(t transport.Transport) {
	er.transports = append(er.transports, t)
}

func (er *EuroRoute) GetTransports() []transport.Transport {
	return er.transports
}

func (er *EuroRoute) ShowTransports() {
	basicShowTransport(er)
}

func (er *EuroRoute) Run() {
	basicRun(er)
}

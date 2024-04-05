package route

type EuroRoute struct {
	transports []Transport
}

func (er *EuroRoute) AddTransport(t Transport) {
	er.transports = append(er.transports, t)
}

func (er *EuroRoute) GetTransports() []Transport {
	return er.transports
}

func (er *EuroRoute) ShowTransports() {
	basicShowTransport(er)
}

func (er *EuroRoute) Run() {
	basicRun(er)
}

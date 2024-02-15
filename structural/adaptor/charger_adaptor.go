package adaptor

type DatabaseDriver interface {
	establishConnection(s string)
}

type CurrentImplementation struct{}

func (CurrentImplementation) getConnection() string {
	return "connection"
}

type DatabaseAdaptor struct {
	currentImplementation CurrentImplementation
}

func (d *DatabaseAdaptor) establishConnection(s string) {
	d.currentImplementation.getConnection()
}

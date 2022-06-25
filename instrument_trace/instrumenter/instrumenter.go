package instrumenter

type Instrument interface {
	Instrument(string) ([]byte, error)
}

package dhcp

type Option byte

const (
	// cba keeping a list of all the options, see the RFC:
	// https://datatracker.ietf.org/doc/html/rfc2132

	OptionEnd Option = 255
	OptionPad Option = 0
)

type Options map[Option][]byte

func (m Message) parseOptions(opts []byte) Options {
	// TODO
	return nil
}

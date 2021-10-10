package dhcp

import (
	"fmt"
)

// https://datatracker.ietf.org/doc/html/rfc2132
type Option byte

func (m Message) parseOptions(opts []byte) error {
	for len(opts) > 0 {
		// first octet of an option contains the option code
		opt := Option(opts[0])

		switch opt {
		case 0:
			// pad
			opts = opts[1:]
		case 255:
			// end
			return nil
		default:
			if _, exists := m.options[opt]; exists {
				// no duplicates
				return fmt.Errorf("duplicate option %d", opt)
			}

			if len(opts) < 2 {
				return fmt.Errorf("option %d has no length", opt)
			}

			// second octet of an option contains the option payload size
			size := int(opts[1])
			if len(opts[2:]) < size {
				return fmt.Errorf("option %d claims to have more bytes than remaining payload", opt)
			}

			// the remaining octets of defined size contains the value of the option
			m.options[opt] = opts[2 : 2+size]

			// shift the slice
			opts = opts[2+size:]
		}
	}

	return fmt.Errorf("options parse was terminated early")
}

func (m Message) GetOption(opt Option) []byte {
	return m.options[opt]
}

func (m Message) SetOption(opt Option, value []byte) {
	m.options[opt] = value
}

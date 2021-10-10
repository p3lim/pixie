package dhcp

// https://datatracker.ietf.org/doc/html/rfc2131#section-2
// see Figure 2 and the associated description

// flags only really contain the broadcast flag, everything else is reserved for future use.

type Flags []byte

func (f Flags) Broadcast() bool {
	return f[0] > 127
}

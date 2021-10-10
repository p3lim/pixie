package dhcp

// https://datatracker.ietf.org/doc/html/rfc2131
/*
   0                   1                   2                   3
   0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |     op (1)    |   htype (1)   |   hlen (1)    |   hops (1)    |
   +---------------+---------------+---------------+---------------+
   |                            xid (4)                            |
   +-------------------------------+-------------------------------+
   |           secs (2)            |           flags (2)           |
   +-------------------------------+-------------------------------+
   |                          ciaddr  (4)                          |
   +---------------------------------------------------------------+
   |                          yiaddr  (4)                          |
   +---------------------------------------------------------------+
   |                          siaddr  (4)                          |
   +---------------------------------------------------------------+
   |                          giaddr  (4)                          |
   +---------------------------------------------------------------+
   |                                                               |
   |                          chaddr  (16)                         |
   |                                                               |
   |                                                               |
   +---------------------------------------------------------------+
   |                                                               |
   |                          sname   (64)                         |
   +---------------------------------------------------------------+
   |                                                               |
   |                          file    (128)                        |
   +---------------------------------------------------------------+
   |                                                               |
   |                          options (variable)                   |
   +---------------------------------------------------------------+
*/

type Message []byte

func (m Message) GetOP() byte {
	return m[0]
}

func (m Message) GetHTYPE() byte {
	return m[1]
}

func (m Message) GetHLEN() byte {
	return m[2]
}

func (m Message) GetHOPS() byte {
	return m[3]
}

func (m Message) GetXID() []byte {
	return m[4:8]
}

func (m Message) GetSECS() []byte {
	return m[8:10]
}

func (m Message) GetFLAGS() []byte {
	return m[10:12]
}

func (m Message) GetCIADDR() []byte {
	return m[12:16]
}

func (m Message) GetYIADDR() []byte {
	return m[16:20]
}

func (m Message) GetSIADDR() []byte {
	return m[20:24]
}

func (m Message) GetGIADDR() []byte {
	return m[24:28]
}

func (m Message) GetCHADDR() []byte {
	return m[28:44]
}

func (m Message) GetSNAME() []byte {
	return m[44:108]
}

func (m Message) GetFILE() []byte {
	return m[108:236]
}

func (m Message) GetOPTIONS() []byte {
	return m[236:]
}

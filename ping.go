package jipmi

import (
	"encoding/binary"
	"errors"
	"net"
	"time"
)

type rmcp_hdr struct {
	ver      byte
	reserved byte
	seq      byte
	class    byte
}

type asf_hdr struct {
	iana     uint32
	msg_type byte
	msg_tag  byte
	reserved byte
	length   byte
}

// 13.2.3 RMCP/ASF Presence Ping Message, table 13-6
type rmcp_ping struct {
	rmcp_hdr
	asf_hdr
}

// 13.2.4 RMCP/ASF Pong Message (Ping Response), table 13-7
type rmcp_pong struct {
	rmcp         rmcp_hdr
	asf          asf_hdr
	iana         uint32
	oem          uint32
	sup_entities byte
	sup_interact byte
	reserved     [6]byte
}

// send one RMCP/ASF Presence Ping Message to host
func Ping(host net.IP) (ok bool, err error) {
	req := &rmcp_ping{
		ver:      0x06,
		seq:      0xFF,
		class:    0x06,
		iana:     4542,
		msg_type: PRESENCE_PING,
		msg_tag:  'J',
	}

	conn, err := net.DialUDP("udp", net.JoinHostPort(host.String(), PRI_RMCP_PORT))
	if err != nil {
		return
	}
	defer conn.Close()
	conn.SetDeadline(PING_TIMEOUT)

	_, err = binary.Write(con, binary.BigEndian, req)
	if err != nil {
		return
	}

	res := &rmcp_pong{}

	n, err := binary.Read(c, binary.BigEndian, res)
	if err != nil {
		return
	}
	if n < binary.Size(res) {
		err = errors.New("short read")
		return
	}

	if res.asf.msg_type == PRESENCE_PONG &&
		res.asf.msg_tag == 'J' &&
		res.sup_entities == 0x81 {
		ok = true
	}
	return
}

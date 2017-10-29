package jipmi

import (
	"time"
)

const (
	PRI_RMCP_PORT = "623" // 13.1.2 RMCP Port Numbers
	PING_TIMEOUT  = time.Second * 3
	PRESENCE_PING = 0x80
	PRESENCE_PONG = 0x40
)

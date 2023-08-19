package telnet

import (
	"strings"
	"time"

	"github.com/reiver/go-telnet"
)

type Telnet struct {
	conn *telnet.Conn
}

func (t Telnet) ReaderTelnet(expect string) (out string) {
	var buffer [1]byte
	recvData := buffer[:]
	var n int
	var err error

	for {
		n, err = t.conn.Read(recvData)
		if n <= 0 || err != nil || strings.Contains(out, expect) {
			break
		} else {
			out += string(recvData)
		}
	}
	return out
}

func (t Telnet) SenderTelnet(command string) {
	time.Sleep(50 * time.Millisecond)
	var commandBuffer []byte
	for _, char := range command {
		commandBuffer = append(commandBuffer, byte(char))
	}
	t.conn.Write(commandBuffer)
}

func (t *Telnet) Connect(host string, port string) error {
	var err error
	t.conn, err = telnet.DialTo(host + ":" + port)
	if err != nil {
		return err
	}
	return nil
}

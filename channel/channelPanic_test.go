package channel

import "testing"

func TestCloseNilChannel(t *testing.T) {
	closeNilChannel()
}

func TestReadNilChannel(t *testing.T) {
	readNilChannel()
}

func TestWriteNilChannel(t *testing.T) {
	writeNilChannel()
}

func TestCloseClosedChannel(t *testing.T) {
	channel := closeChannel()
	close(channel)
}

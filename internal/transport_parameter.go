package internal

import (
	"bytes"
	"fmt"
	"github.com/lucas-clemente/quic-go/quicvarint"
)

type transportParameterID uint64

const (
	originalDestinationConnectionIDParameterID transportParameterID = 0x0
	maxIdleTimeoutParameterID                  transportParameterID = 0x1
	statelessResetTokenParameterID             transportParameterID = 0x2
	maxUDPPayloadSizeParameterID               transportParameterID = 0x3
	initialMaxDataParameterID                  transportParameterID = 0x4
	initialMaxStreamDataBidiLocalParameterID   transportParameterID = 0x5
	initialMaxStreamDataBidiRemoteParameterID  transportParameterID = 0x6
	initialMaxStreamDataUniParameterID         transportParameterID = 0x7
	initialMaxStreamsBidiParameterID           transportParameterID = 0x8
	initialMaxStreamsUniParameterID            transportParameterID = 0x9
	ackDelayExponentParameterID                transportParameterID = 0xa
	maxAckDelayParameterID                     transportParameterID = 0xb
	disableActiveMigrationParameterID          transportParameterID = 0xc
	preferredAddressParameterID                transportParameterID = 0xd
	activeConnectionIDLimitParameterID         transportParameterID = 0xe
	initialSourceConnectionIDParameterID       transportParameterID = 0xf
	retrySourceConnectionIDParameterID         transportParameterID = 0x10
	// RFC 9221
	maxDatagramFrameSizeParameterID transportParameterID = 0x20
	// TODO Not IANA Registered! 0x40 is used temporarily (see RFC9000 section 22.3)
	extraStreamEncryptionParameterID transportParameterID = 0x40
)

var parameterID2Name = map[transportParameterID]string{
	originalDestinationConnectionIDParameterID: "original_destination_connection_id",
	maxIdleTimeoutParameterID:                  "max_idle_timeout",
	statelessResetTokenParameterID:             "stateless_reset_token",
	maxUDPPayloadSizeParameterID:               "max_udp_payload_size",
	initialMaxDataParameterID:                  "initial_max_data",
	initialMaxStreamDataBidiLocalParameterID:   "initial_max_stream_data_bidi_local",
	initialMaxStreamDataBidiRemoteParameterID:  "initial_max_stream_data_bidi_remote",
	initialMaxStreamDataUniParameterID:         "initial_max_stream_data_uni",
	initialMaxStreamsBidiParameterID:           "initial_max_streams_bidi",
	initialMaxStreamsUniParameterID:            "initial_max_streams_uni",
	ackDelayExponentParameterID:                "ack-delay_exponent",
	maxAckDelayParameterID:                     "max_ack_delay",
	disableActiveMigrationParameterID:          "disable_active_migration",
	preferredAddressParameterID:                "preferred_address",
	activeConnectionIDLimitParameterID:         "active_connection_id_limit",
	initialSourceConnectionIDParameterID:       "initial_source_connection_id",
	retrySourceConnectionIDParameterID:         "retry_source_connection_id",
	// RFC 9221
	maxDatagramFrameSizeParameterID: "max_datagram_frame_size",
	// TODO Not IANA Registered! 0x40 is used temporarily (see RFC9000 section 22.3)
	extraStreamEncryptionParameterID: "extra_stream_encryption",
}
var name2ParameterID = InvertMap(parameterID2Name)

type TransportParameter struct {
	id        transportParameterID
	byteValue []byte
}

// return name;
// if name is unknown return id
func (e *TransportParameter) name() string {
	name, ok := parameterID2Name[e.id]
	if !ok {
		return fmt.Sprintf("0x%02x", e.id)
	}
	return name
}

// return value in a human-readable format;
// if parameter is unknown return byte array;
func (e *TransportParameter) stringValue() (string, error) {
	switch e.id {
	case // varint
		maxIdleTimeoutParameterID,
		maxUDPPayloadSizeParameterID,
		initialMaxDataParameterID,
		initialMaxStreamDataBidiLocalParameterID,
		initialMaxStreamDataBidiRemoteParameterID,
		initialMaxStreamDataUniParameterID,
		initialMaxStreamsBidiParameterID,
		initialMaxStreamsUniParameterID,
		maxAckDelayParameterID,
		activeConnectionIDLimitParameterID,
		maxDatagramFrameSizeParameterID,
		ackDelayExponentParameterID:
		val, err := quicvarint.Read(bytes.NewReader(e.byteValue))
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", val), nil
	default: // unknown
		return fmt.Sprintf("%v", e.byteValue), nil
	}

}

func (e *TransportParameter) String() (string, error) {
	stringValue, err := e.stringValue()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s, %s", e.name(), stringValue), nil
}

func ParseNextTransportParameter(r *bytes.Reader) (TransportParameter, error) {
	tpe := TransportParameter{}
	paramIDInt, err := quicvarint.Read(r)
	if err != nil {
		return tpe, fmt.Errorf("failed to read param id: %s", err)
	}
	tpe.id = transportParameterID(paramIDInt)
	paramLen, err := quicvarint.Read(r)
	if err != nil {
		return tpe, fmt.Errorf("failed to read param length of %s: %s", tpe.name(), err)
	}
	tpe.byteValue = make([]byte, paramLen)
	if paramLen == 0 {
		return tpe, nil
	}
	n, err := r.Read(tpe.byteValue)
	if err != nil {
		return tpe, fmt.Errorf("failed to read %d byte value of %s: %s", paramLen, tpe.name(), err)
	}
	if len(tpe.byteValue) != n {
		return tpe, fmt.Errorf("invalid length of %s", tpe.name())
	}
	return tpe, nil
}

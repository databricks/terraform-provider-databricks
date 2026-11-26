package traceparent

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

type traceparentFlag byte

const (
	traceparentFlagSampled traceparentFlag = 1 << iota
)

const TRACEPARENT_HEADER = "traceparent"

type Traceparent struct {
	version  byte
	traceId  [16]byte
	parentId [8]byte
	flags    traceparentFlag
}

func (t *Traceparent) String() string {
	b := strings.Builder{}
	b.WriteString(hex.EncodeToString([]byte{t.version}))
	b.WriteRune('-')
	b.WriteString(hex.EncodeToString(t.traceId[:]))
	b.WriteRune('-')
	b.WriteString(hex.EncodeToString(t.parentId[:]))
	b.WriteRune('-')
	b.WriteString(hex.EncodeToString([]byte{byte(t.flags)}))
	return b.String()
}

func (t *Traceparent) Equals(other *Traceparent) bool {
	return t.version == other.version &&
		t.flags == other.flags &&
		string(t.traceId[:]) == string(other.traceId[:]) &&
		string(t.parentId[:]) == string(other.parentId[:])
}

func NewTraceparent() *Traceparent {
	traceId := [16]byte{}
	parentId := [8]byte{}
	rand.Read(traceId[:])
	rand.Read(parentId[:])
	return &Traceparent{
		version:  0,
		traceId:  traceId,
		parentId: parentId,
		flags:    traceparentFlagSampled,
	}
}

func FromString(s string) (*Traceparent, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid traceparent string: %s", s)
	}
	t := &Traceparent{
		traceId:  [16]byte{},
		parentId: [8]byte{},
	}
	version, err := hex.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	if len(version) != 1 {
		return nil, fmt.Errorf("invalid version: %s, expected 1 byte", parts[0])
	}
	t.version = version[0]
	n, err := hex.Decode(t.traceId[:], []byte(parts[1]))
	if err != nil {
		return nil, err
	}
	if n != 16 {
		return nil, fmt.Errorf("invalid traceId: %s, expected 16 bytes", parts[1])
	}
	n, err = hex.Decode(t.parentId[:], []byte(parts[2]))
	if err != nil {
		return nil, err
	}
	if n != 8 {
		return nil, fmt.Errorf("invalid parentId: %s, expected 8 bytes", parts[2])
	}
	flags, err := hex.DecodeString(parts[3])
	if err != nil {
		return nil, err
	}
	if len(flags) != 1 {
		return nil, fmt.Errorf("invalid flags: %s, expected 1 byte", parts[3])
	}
	t.flags = traceparentFlag(flags[0])
	return t, nil
}

func AddTraceparent(r *http.Request) {
	if r.Header.Get(TRACEPARENT_HEADER) == "" {
		r.Header.Set(TRACEPARENT_HEADER, NewTraceparent().String())
	}
}

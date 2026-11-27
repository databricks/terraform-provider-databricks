package traceparent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tp := NewTraceparent()
	assert.Equal(t, byte(0), tp.version)
	assert.Equal(t, byte(1), byte(tp.flags))
}

func TestEqual(t *testing.T) {
	tp1 := NewTraceparent()
	tp2 := &Traceparent{
		version:  tp1.version,
		traceId:  tp1.traceId,
		parentId: tp1.parentId,
		flags:    tp1.flags,
	}
	assert.True(t, tp1.Equals(tp2))
}

func TestTwoNewTraceparentsAreNotEqual(t *testing.T) {
	tp1 := NewTraceparent()
	tp2 := NewTraceparent()
	assert.False(t, tp1.Equals(tp2))
}

var testTraceId = [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
var testParentId = [8]byte{0, 1, 2, 3, 4, 5, 6, 7}

func TestString(t *testing.T) {
	tp := &Traceparent{
		version:  0,
		traceId:  testTraceId,
		parentId: testParentId,
		flags:    1,
	}
	res := tp.String()
	assert.Equal(t, "00-000102030405060708090a0b0c0d0e0f-0001020304050607-01", res)
}

func TestFromString(t *testing.T) {
	tp, err := FromString("00-000102030405060708090a0b0c0d0e0f-0001020304050607-01")
	assert.NoError(t, err)
	assert.Equal(t, byte(0), tp.version)
	assert.Equal(t, testTraceId, tp.traceId)
	assert.Equal(t, testParentId, tp.parentId)
	assert.Equal(t, byte(1), byte(tp.flags))
}

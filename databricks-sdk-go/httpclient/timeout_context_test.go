package httpclient

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeoutContextTimeout(t *testing.T) {
	ctx := context.Background()
	ctx, _ = newTimeoutContext(ctx, time.Millisecond*50)
	time.Sleep(time.Millisecond * 100)

	// The context should have timed out.
	assert.Equal(t, context.Canceled, ctx.Err())
}

func TestTimeoutContextTick(t *testing.T) {
	ctx := context.Background()
	ctx, ticker := newTimeoutContext(ctx, time.Millisecond*50)

	// Extend the deadline a couple of times.
	for i := 0; i < 5; i++ {
		ticker.Tick()
		time.Sleep(time.Millisecond * 25)
	}

	// The context should not have timed out.
	assert.Nil(t, ctx.Err())
}

func TestTimeoutContextCancel(t *testing.T) {
	ctx := context.Background()
	ctx, ticker := newTimeoutContext(ctx, time.Millisecond*50)

	// Cancel the context.
	ticker.Cancel()

	// The context should have timed out.
	assert.Equal(t, context.Canceled, ctx.Err())
}

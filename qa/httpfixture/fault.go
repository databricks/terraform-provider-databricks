package httpfixture

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// Fault names what the server does to the connection instead of completing a
// response. The resulting client error depends on the client's HTTP stack.
type Fault string

const (
	// FaultReset closes the connection with linger 0, so the peer gets a TCP reset.
	FaultReset Fault = "reset"
	// FaultClose closes the connection without writing a response.
	FaultClose Fault = "close"
	// FaultTruncateBody declares the full Content-Length, writes the first half
	// of the body and closes the connection.
	FaultTruncateBody Fault = "truncate_body"
	// FaultStallBody writes the first half of the body, then waits until the
	// client closes the connection.
	FaultStallBody Fault = "stall_body"
)

var statusLinePattern = regexp.MustCompile(`^HTTP/1\.[01] ([0-9]{3})(?: [^\r\n]*)?$`)

// writesStatus reports whether the response starts with a status line, which
// needs a valid status. Unknown faults are left to compileResponse.
func (r Response) writesStatus() bool {
	return r.Fault == "" || r.Fault == FaultTruncateBody || r.Fault == FaultStallBody
}

// compileResponse validates the response and encodes its body once.
func compileResponse(e *exchange, r Response) error {
	if r.Delay != "" {
		delay, err := time.ParseDuration(r.Delay)
		if err != nil || delay <= 0 {
			return fmt.Errorf("delay %q must be a positive duration, such as 2s", r.Delay)
		}
		e.delay = delay
	}
	switch r.Fault {
	case FaultReset, FaultClose:
		if r.Status != 0 || r.Headers != nil || r.Body != nil || r.StatusLine != "" {
			return fmt.Errorf("fault %s writes no response; omit status, headers, body and status_line", r.Fault)
		}
		return nil
	case "", FaultTruncateBody, FaultStallBody:
	default:
		return fmt.Errorf("unknown fault %q; expected reset, close, truncate_body or stall_body", r.Fault)
	}
	if r.Body != nil {
		var err error
		if e.responseBody, err = r.Body.literalBytes(); err != nil {
			return err
		}
	}
	if r.Fault != "" && len(e.responseBody) == 0 {
		return fmt.Errorf("fault %s requires a nonempty response body", r.Fault)
	}
	if r.Fault != "" && !bodyAllowed(e.call.Request.Method, r.Status) {
		return fmt.Errorf("fault %s needs a response body, which net/http ignores for HEAD requests and 1xx, 204 and 304 responses", r.Fault)
	}
	// Status stays the source of truth for the code; the line only changes its spelling.
	if match := statusLinePattern.FindStringSubmatch(r.StatusLine); r.StatusLine != "" &&
		(match == nil || match[1] != strconv.Itoa(r.Status)) {
		return fmt.Errorf("status_line %q must be HTTP/1.1 %d, optionally followed by a reason phrase", r.StatusLine, r.Status)
	}
	return nil
}

// bodyAllowed mirrors net/http, whose client never reads these responses' bodies,
// so a body fault on them would pass without reaching the client.
func bodyAllowed(method string, status int) bool {
	return method != http.MethodHead && status >= 200 && status != http.StatusNoContent && status != http.StatusNotModified
}

// respond waits for the delay, then writes the response. It runs outside the
// handler lock, so a slow response does not hold up concurrent requests.
func respond(w http.ResponseWriter, req *http.Request, e *exchange) error {
	if e.delay > 0 {
		timer := time.NewTimer(e.delay)
		defer timer.Stop()
		select {
		case <-timer.C:
		case <-req.Context().Done():
			// The client closed the connection, so nobody reads the response.
			return nil
		}
	}
	if e.call.Response.Fault == "" && e.call.Response.StatusLine == "" {
		writeResponse(w, e)
		return nil
	}
	return writeRaw(w, e)
}

// writeRaw hijacks the connection to write what net/http cannot: a literal
// status line, a broken connection, or a body shorter than its Content-Length.
// Write errors are ignored, as in writeResponse: they mean the client left.
func writeRaw(w http.ResponseWriter, e *exchange) error {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		return errors.New("fault and status_line require a connection that supports hijacking")
	}
	conn, client, err := hijacker.Hijack()
	if err != nil {
		return err
	}
	defer conn.Close()
	switch e.call.Response.Fault {
	case FaultReset:
		return setLingerZero(conn)
	case FaultClose:
		return nil
	}
	body := e.responseBody
	if e.call.Response.Fault != "" {
		body = body[:len(body)/2]
	}
	_, err = conn.Write(append(rawHead(e), body...))
	if err == nil && e.call.Response.Fault == FaultStallBody {
		// The client sends nothing more, so the read ends when it closes the connection.
		_, _ = io.Copy(io.Discard, client)
	}
	return nil
}

// setLingerZero makes the deferred Close send a reset instead of a normal shutdown.
func setLingerZero(conn net.Conn) error {
	tcp, ok := conn.(interface{ SetLinger(int) error })
	if !ok {
		return fmt.Errorf("fault reset requires a TCP connection, got %T", conn)
	}
	return tcp.SetLinger(0)
}

// rawHead frames the body by its full length, even when a fault writes less of
// it, and announces the close, replacing any such headers in the fixture.
func rawHead(e *exchange) []byte {
	r := e.call.Response
	line := r.StatusLine
	if line == "" {
		line = fmt.Sprintf("HTTP/1.1 %d %s", r.Status, http.StatusText(r.Status))
	}
	header := http.Header{}
	setResponseHeaders(header, e)
	header.Del("Transfer-Encoding")
	header.Set("Content-Length", strconv.Itoa(len(e.responseBody)))
	header.Set("Connection", "close")
	var head bytes.Buffer
	head.WriteString(line + "\r\n")
	_ = header.Write(&head)
	head.WriteString("\r\n")
	return head.Bytes()
}

package httpfixture

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Times is how many times in a row a call is expected. The zero value means
// once. Min below Max allows a count that depends on timing, such as retries
// that end with a timeout.
type Times struct {
	Min int
	Max int
	// err defers a YAML shape error to compileCalls, which names the step and call.
	err error
}

// Exactly expects a call n times in a row. A count below 1 fails when the
// fixtures compile, rather than becoming the zero value, which means once.
func Exactly(n int) Times {
	if n < 1 {
		return Times{err: fmt.Errorf("times %d must be a positive count", n)}
	}
	return Times{Min: n, Max: n}
}

// IsZero reports an omitted times field, so YAML encoding omits it too.
func (t Times) IsZero() bool { return t.Min == 0 && t.Max == 0 && t.err == nil }

// UnmarshalYAML accepts a count, such as 3, or a range, such as {min: 2, max: 3}.
func (t *Times) UnmarshalYAML(node *yaml.Node) error {
	*t = Times{}
	switch node.Kind {
	case yaml.ScalarNode:
		n, err := strconv.Atoi(node.Value)
		if err != nil || n < 1 {
			t.err = fmt.Errorf("times %q must be a positive count or {min, max}", node.Value)
			return nil
		}
		*t = Exactly(n)
	case yaml.MappingNode:
		t.err = t.decodeRange(node)
	default:
		t.err = errors.New("times must be a positive count or {min, max}")
	}
	return nil
}

func (t *Times) decodeRange(node *yaml.Node) error {
	seen := map[string]bool{}
	for i := 0; i+1 < len(node.Content); i += 2 {
		key, value := node.Content[i].Value, node.Content[i+1]
		n, err := strconv.Atoi(value.Value)
		if value.Kind != yaml.ScalarNode || err != nil {
			return fmt.Errorf("times.%s %q must be an integer", key, value.Value)
		}
		switch {
		case seen[key]:
			return fmt.Errorf("times.%s is set twice", key)
		case key == "min":
			t.Min = n
		case key == "max":
			t.Max = n
		default:
			return fmt.Errorf("unknown times field %q; expected min and max", key)
		}
		seen[key] = true
	}
	if !seen["min"] || !seen["max"] {
		return errors.New("times range requires both min and max")
	}
	// Checked here as well, because {min: 0, max: 0} decodes to the zero value.
	return t.validate()
}

func (t Times) validate() error {
	if t.Min < 0 || t.Max < 1 || t.Min > t.Max {
		return fmt.Errorf("times {min: %d, max: %d} needs 0 <= min <= max and max >= 1",
			t.Min, t.Max)
	}
	return nil
}

// MarshalYAML writes the count form when the range is exact.
func (t Times) MarshalYAML() (any, error) {
	if t.Min == t.Max {
		return t.Max, nil
	}
	return struct {
		Min int `yaml:"min"`
		Max int `yaml:"max"`
	}{t.Min, t.Max}, nil
}

// repeat is a call's compiled times, gap bounds and arrivals.
type repeat struct {
	min, max       int
	repeated       bool
	minGap, maxGap time.Duration
	count          int
	// closed stops further repetitions once ordered replay has moved past the call.
	closed bool
}

func compileRepeat(c Call) (repeat, error) {
	r := repeat{min: 1, max: 1}
	if !c.Times.IsZero() {
		if c.Times.err != nil {
			return r, c.Times.err
		}
		if err := c.Times.validate(); err != nil {
			return r, err
		}
		r.min, r.max, r.repeated = c.Times.Min, c.Times.Max, true
	}
	var err error
	if r.minGap, err = parseGap("min_gap", c.MinGap); err != nil {
		return r, err
	}
	if r.maxGap, err = parseGap("max_gap", c.MaxGap); err != nil {
		return r, err
	}
	if r.minGap > 0 && r.maxGap > 0 && r.minGap > r.maxGap {
		return r, fmt.Errorf("min_gap %s exceeds max_gap %s", c.MinGap, c.MaxGap)
	}
	if r.hasGap() && r.repeated && r.max < 2 {
		return r, errors.New("with times, min_gap and max_gap bound the gaps between " +
			"repetitions, so times needs a max of at least 2")
	}
	return r, nil
}

func parseGap(field, value string) (time.Duration, error) {
	if value == "" {
		return 0, nil
	}
	gap, err := time.ParseDuration(value)
	if err != nil || gap <= 0 {
		return 0, fmt.Errorf("%s %q must be a positive duration, such as 1s", field, value)
	}
	return gap, nil
}

func (r *repeat) hasGap() bool { return r.minGap > 0 || r.maxGap > 0 }

// open reports whether the call can match another request.
func (r *repeat) open() bool { return !r.closed && r.count < r.max }

func (r *repeat) satisfied() bool { return r.count >= r.min }

// arrival is a matched call's index, arrival time and request line.
type arrival struct {
	call    int
	at      time.Time
	request string
}

// gapFailures checks gap bounds in arrival order, because concurrent requests
// can take the lock in another order after reading their bodies. A repeated
// call measures from its previous repetition and leaves its first unbounded; a
// call without times measures from the previous matched call in the step.
func (h *Handler) gapFailures() []error {
	ordered := slices.Clone(h.arrivals)
	slices.SortStableFunc(ordered, func(a, b arrival) int { return a.at.Compare(b.at) })
	var failures []error
	repetitions, previous := map[int]int{}, map[int]time.Time{}
	for i, a := range ordered {
		e := &h.calls[a.call]
		repetitions[a.call]++
		n := repetitions[a.call]
		var err error
		switch {
		case !e.hasGap() || e.repeated && n == 1:
		case e.repeated:
			err = e.checkGap(a.at.Sub(previous[a.call]), fmt.Sprintf("repetition %d arrived", n), fmt.Sprintf("repetition %d", n-1))
		case i == 0:
			err = errors.New("arrived first in its step, so min_gap and max_gap have no " +
				"earlier call to measure from")
		default:
			p := ordered[i-1]
			err = e.checkGap(a.at.Sub(p.at), "arrived", fmt.Sprintf("call %d (%s)", p.call+1, p.request))
		}
		previous[a.call] = a.at
		if err != nil {
			failures = append(failures, fmt.Errorf("%s: call %d: %s %w", h.label, a.call+1, a.request, err))
		}
	}
	return failures
}

func (r *repeat) checkGap(gap time.Duration, arrived, from string) error {
	observed := fmt.Sprintf("%s %s after %s", arrived, gap.Round(time.Millisecond), from)
	if r.minGap > 0 && gap < r.minGap {
		return fmt.Errorf("%s, below min_gap %s", observed, r.minGap)
	}
	if r.maxGap > 0 && gap > r.maxGap {
		return fmt.Errorf("%s, above max_gap %s", observed, r.maxGap)
	}
	return nil
}

// countDetail explains an unsatisfied repeated call's count.
func (r *repeat) countDetail() string {
	switch {
	case r.max == 1:
		return ""
	case r.min == r.max:
		return fmt.Sprintf(": expected %d times, matched %d", r.max, r.count)
	default:
		return fmt.Sprintf(": expected %d to %d times, matched %d", r.min, r.max, r.count)
	}
}

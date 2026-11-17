// Package iptypes provides a Terraform string type whose values compare
// equal when they spell the same IPv4 or IPv6 address, so that a service that
// normalizes addresses does not cause drift.
package iptypes

import (
	"context"
	"fmt"
	"net/netip"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// ParseIPAddress parses an IPv4 or IPv6 address. Two strings spell the same
// address if and only if their parsed values are ==.
//
// It accepts what [netip.ParseAddr] accepts, plus decimal leading zeros in
// IPv4 octets (also in the dotted tail of an IPv6 address), and it returns
// IPv4-mapped IPv6 addresses as IPv4. Other IPv6 forms that embed an IPv4
// address, such as IPv4-compatible or NAT64 addresses, stay distinct from it.
// It rejects zone identifiers, prefixes, ports, brackets, surrounding
// whitespace, and integer or hexadecimal IPv4.
func ParseIPAddress(s string) (netip.Addr, error) {
	// netip accepts scoped IPv6 such as "fe80::1%eth0", but a zone is local to
	// one host and is not part of the address.
	if strings.Contains(s, "%") {
		return netip.Addr{}, fmt.Errorf("invalid IP address %q: zone identifiers are not supported", s)
	}
	// netip rejects leading zeros in dotted octets as possibly octal, but they
	// are decimal padding here. Empty octets stay empty so that malformed input
	// still fails to parse.
	normalized := s
	tailStart := strings.LastIndex(s, ":") + 1
	if octets := strings.Split(s[tailStart:], "."); len(octets) == 4 {
		for i, octet := range octets {
			for len(octet) > 1 && octet[0] == '0' {
				octet = octet[1:]
			}
			octets[i] = octet
		}
		normalized = s[:tailStart] + strings.Join(octets, ".")
	}
	addr, err := netip.ParseAddr(normalized)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("invalid IP address %q: %w", s, err)
	}
	return addr.Unmap(), nil
}

var _ basetypes.StringTypable = IPAddressType{}

// IPAddressType is the Terraform type of [IPAddress] values.
type IPAddressType struct {
	basetypes.StringType
}

// String returns a human-readable name of the type.
func (t IPAddressType) String() string {
	return "iptypes.IPAddressType"
}

// ValueType returns the value type of this type.
func (t IPAddressType) ValueType(context.Context) attr.Value {
	return IPAddress{}
}

// Equal reports whether o is an IPAddressType.
func (t IPAddressType) Equal(o attr.Type) bool {
	_, ok := o.(IPAddressType)
	return ok
}

// ValueFromString wraps in without changing its spelling or state.
func (t IPAddressType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPAddress{StringValue: in}, nil
}

// ValueFromTerraform converts a Terraform string without changing its
// spelling or state.
func (t IPAddressType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	value, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	stringValue, ok := value.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", value)
	}
	return IPAddress{StringValue: stringValue}, nil
}

var _ basetypes.StringValuableWithSemanticEquals = IPAddress{}

// IPAddress is a Terraform string that holds an IPv4 or IPv6 address. It
// keeps the exact spelling it was given, because Terraform rejects an applied
// value that differs from the configured one. Terraform uses semantic equality
// to keep the prior spelling when the service returns an equivalent address.
//
// The type does not validate values: the service owns the accepted input
// contract, and a malformed value is never semantically equal to another
// value. The zero value is null.
type IPAddress struct {
	basetypes.StringValue
}

// NewIPAddressValue returns a known IPAddress with the given spelling.
func NewIPAddressValue(value string) IPAddress {
	return IPAddress{StringValue: basetypes.NewStringValue(value)}
}

// NewIPAddressNull returns a null IPAddress.
func NewIPAddressNull() IPAddress {
	return IPAddress{StringValue: basetypes.NewStringNull()}
}

// NewIPAddressUnknown returns an unknown IPAddress.
func NewIPAddressUnknown() IPAddress {
	return IPAddress{StringValue: basetypes.NewStringUnknown()}
}

// Type returns an IPAddressType.
func (v IPAddress) Type(context.Context) attr.Type {
	return IPAddressType{}
}

// Equal reports whether o is an IPAddress with the same state and the exact
// same spelling. Terraform relies on exact equality; use
// [IPAddress.StringSemanticEquals] to compare addresses.
func (v IPAddress) Equal(o attr.Value) bool {
	other, ok := o.(IPAddress)
	return ok && v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals reports whether both values are known and spell the
// same address. Null, unknown, and malformed values are never semantically
// equal, even to an identical string.
func (v IPAddress) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newValue, ok := newValuable.(IPAddress)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			fmt.Sprintf("An unexpected value type was received while performing semantic equality checks. Please report this to the provider developers.\n\nExpected Value Type: %T\nGot Value Type: %T", v, newValuable),
		)
		return false, diags
	}
	address, err := v.ValueIPAddress()
	newAddress, newErr := newValue.ValueIPAddress()
	return err == nil && newErr == nil && address == newAddress, diags
}

// ValueIPAddress parses the value with [ParseIPAddress]. It returns an error
// if the value is null, unknown, or malformed: null and unknown values hold
// the empty string, which does not parse.
func (v IPAddress) ValueIPAddress() (netip.Addr, error) {
	return ParseIPAddress(v.ValueString())
}

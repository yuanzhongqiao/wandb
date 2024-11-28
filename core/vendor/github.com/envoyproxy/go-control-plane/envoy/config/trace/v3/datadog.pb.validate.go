//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v3/datadog.proto

package tracev3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DatadogRemoteConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DatadogRemoteConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DatadogRemoteConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DatadogRemoteConfigMultiError, or nil if none found.
func (m *DatadogRemoteConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DatadogRemoteConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPollingInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DatadogRemoteConfigValidationError{
					field:  "PollingInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DatadogRemoteConfigValidationError{
					field:  "PollingInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPollingInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DatadogRemoteConfigValidationError{
				field:  "PollingInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DatadogRemoteConfigMultiError(errors)
	}

	return nil
}

// DatadogRemoteConfigMultiError is an error wrapping multiple validation
// errors returned by DatadogRemoteConfig.ValidateAll() if the designated
// constraints aren't met.
type DatadogRemoteConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DatadogRemoteConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DatadogRemoteConfigMultiError) AllErrors() []error { return m }

// DatadogRemoteConfigValidationError is the validation error returned by
// DatadogRemoteConfig.Validate if the designated constraints aren't met.
type DatadogRemoteConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DatadogRemoteConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DatadogRemoteConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DatadogRemoteConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DatadogRemoteConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DatadogRemoteConfigValidationError) ErrorName() string {
	return "DatadogRemoteConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DatadogRemoteConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDatadogRemoteConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DatadogRemoteConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DatadogRemoteConfigValidationError{}

// Validate checks the field values on DatadogConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DatadogConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DatadogConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DatadogConfigMultiError, or
// nil if none found.
func (m *DatadogConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DatadogConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCollectorCluster()) < 1 {
		err := DatadogConfigValidationError{
			field:  "CollectorCluster",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetServiceName()) < 1 {
		err := DatadogConfigValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CollectorHostname

	if all {
		switch v := interface{}(m.GetRemoteConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DatadogConfigValidationError{
					field:  "RemoteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DatadogConfigValidationError{
					field:  "RemoteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRemoteConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DatadogConfigValidationError{
				field:  "RemoteConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DatadogConfigMultiError(errors)
	}

	return nil
}

// DatadogConfigMultiError is an error wrapping multiple validation errors
// returned by DatadogConfig.ValidateAll() if the designated constraints
// aren't met.
type DatadogConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DatadogConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DatadogConfigMultiError) AllErrors() []error { return m }

// DatadogConfigValidationError is the validation error returned by
// DatadogConfig.Validate if the designated constraints aren't met.
type DatadogConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DatadogConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DatadogConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DatadogConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DatadogConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DatadogConfigValidationError) ErrorName() string { return "DatadogConfigValidationError" }

// Error satisfies the builtin error interface
func (e DatadogConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDatadogConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DatadogConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DatadogConfigValidationError{}
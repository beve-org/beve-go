package beve

import (
	"encoding/binary"
	"fmt"
	"time"
)

// EncodeTimestamp encodes a Timestamp with optional timezone (Extension 4)
// Size: 14 bytes (UTC) or 16 bytes (with timezone)
func EncodeTimestamp(ts Timestamp) ([]byte, error) {
	hasTimezone := ts.TimezoneOffset != nil

	// Calculate total size
	size := 1 + 1 + 8 + 4 // header + precision + seconds + nanoseconds
	if hasTimezone {
		size += 2 // timezone offset
	}

	buf := make([]byte, size)
	offset := 0

	// Header
	buf[offset] = ExtTimestamp
	offset++

	// Precision + timezone flag
	precision := PrecisionNanoseconds
	if hasTimezone {
		precision |= FlagHasTimezone
	}
	buf[offset] = precision
	offset++

	// Epoch seconds (little-endian)
	binary.LittleEndian.PutUint64(buf[offset:], uint64(ts.Seconds))
	offset += 8

	// Nanoseconds (little-endian)
	binary.LittleEndian.PutUint32(buf[offset:], ts.Nanoseconds)
	offset += 4

	// Optional timezone offset
	if hasTimezone {
		binary.LittleEndian.PutUint16(buf[offset:], uint16(*ts.TimezoneOffset))
	}

	return buf, nil
}

// DecodeTimestamp decodes Extension 4 timestamp
func DecodeTimestamp(data []byte) (Timestamp, error) {
	if len(data) < 14 || data[0] != ExtTimestamp {
		return Timestamp{}, fmt.Errorf("invalid timestamp header")
	}

	offset := 1

	// Read precision + flags
	precision := data[offset]
	hasTimezone := (precision & FlagHasTimezone) != 0
	offset++

	// Read epoch seconds
	seconds := int64(binary.LittleEndian.Uint64(data[offset:]))
	offset += 8

	// Read nanoseconds
	nanoseconds := binary.LittleEndian.Uint32(data[offset:])
	offset += 4

	// Read optional timezone
	var timezoneOffset *int16
	if hasTimezone {
		if len(data) < 16 {
			return Timestamp{}, fmt.Errorf("invalid timestamp: missing timezone")
		}
		tz := int16(binary.LittleEndian.Uint16(data[offset:]))
		timezoneOffset = &tz
	}

	return Timestamp{
		Seconds:        seconds,
		Nanoseconds:    nanoseconds,
		TimezoneOffset: timezoneOffset,
	}, nil
}

// TimestampFromTime converts time.Time to Timestamp
func TimestampFromTime(t time.Time) Timestamp {
	seconds := t.Unix()
	nanoseconds := uint32(t.Nanosecond())

	// Check if timezone is not UTC
	_, offset := t.Zone()
	if offset != 0 {
		offsetMinutes := int16(offset / 60)
		return NewTimestampWithTZ(seconds, nanoseconds, offsetMinutes)
	}

	return NewTimestampUTC(seconds, nanoseconds)
}

// ToTime converts Timestamp to time.Time
func (ts Timestamp) ToTime() time.Time {
	if ts.TimezoneOffset == nil {
		// UTC
		return time.Unix(ts.Seconds, int64(ts.Nanoseconds)).UTC()
	}

	// With timezone
	offsetSeconds := int(*ts.TimezoneOffset) * 60
	loc := time.FixedZone("", offsetSeconds)
	return time.Unix(ts.Seconds, int64(ts.Nanoseconds)).In(loc)
}

// MarshalTimestamp marshals a time.Time to BEVE timestamp bytes
func MarshalTimestamp(t time.Time) ([]byte, error) {
	ts := TimestampFromTime(t)
	return EncodeTimestamp(ts)
}

// UnmarshalTimestamp unmarshals BEVE timestamp bytes to time.Time
func UnmarshalTimestamp(data []byte) (time.Time, error) {
	ts, err := DecodeTimestamp(data)
	if err != nil {
		return time.Time{}, err
	}
	return ts.ToTime(), nil
}

// EncodeDuration encodes a time.Duration (Extension 5)
func EncodeDuration(d time.Duration) ([]byte, error) {
	// Duration layout: header + sign_precision + seconds + nanoseconds
	buf := make([]byte, 14)
	offset := 0

	// Header
	buf[offset] = ExtDuration
	offset++

	// Sign and precision
	sign := byte(0)
	if d < 0 {
		sign = 1
		d = -d
	}
	precision := PrecisionNanoseconds
	buf[offset] = (precision) | sign
	offset++

	// Seconds
	seconds := uint64(d / time.Second)
	binary.LittleEndian.PutUint64(buf[offset:], seconds)
	offset += 8

	// Nanoseconds
	nanos := uint32(d % time.Second)
	binary.LittleEndian.PutUint32(buf[offset:], nanos)

	return buf, nil
}

// DecodeDuration decodes Extension 5 duration
func DecodeDuration(data []byte) (time.Duration, error) {
	if len(data) < 14 || data[0] != ExtDuration {
		return 0, fmt.Errorf("invalid duration header")
	}

	offset := 1

	// Read sign and precision
	signPrecision := data[offset]
	isNegative := (signPrecision & 0x01) != 0
	offset++

	// Read seconds
	seconds := binary.LittleEndian.Uint64(data[offset:])
	offset += 8

	// Read nanoseconds
	nanos := binary.LittleEndian.Uint32(data[offset:])

	// Calculate duration
	d := time.Duration(seconds)*time.Second + time.Duration(nanos)*time.Nanosecond

	if isNegative {
		d = -d
	}

	return d, nil
}

// EncodeInterval encodes a time interval (start-end)
func EncodeInterval(start, end time.Time) ([]byte, error) {
	startTS := TimestampFromTime(start)
	endTS := TimestampFromTime(end)

	startBytes, err := EncodeTimestamp(startTS)
	if err != nil {
		return nil, err
	}

	endBytes, err := EncodeTimestamp(endTS)
	if err != nil {
		return nil, err
	}

	// Layout: header + start_timestamp + end_timestamp
	buf := make([]byte, 1+len(startBytes)+len(endBytes))
	buf[0] = ExtInterval
	copy(buf[1:], startBytes)
	copy(buf[1+len(startBytes):], endBytes)

	return buf, nil
}

// DecodeInterval decodes Extension 6 interval
func DecodeInterval(data []byte) (start, end time.Time, err error) {
	if len(data) < 2 || data[0] != ExtInterval {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid interval header")
	}

	// Decode start timestamp
	// Note: Timestamp size is variable (14 or 16 bytes)
	// We need to parse the first timestamp to know where the second starts
	startTS, err := DecodeTimestamp(data[1:])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	// Calculate start timestamp size
	startSize := 14
	if startTS.TimezoneOffset != nil {
		startSize = 16
	}

	// Decode end timestamp
	endTS, err := DecodeTimestamp(data[1+startSize:])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startTS.ToTime(), endTS.ToTime(), nil
}

package beve

import (
	"fmt"
	"regexp"
)

// RegExpFlags represents regex flags (Extension 9)
const (
	FlagCaseInsensitive byte = 0x01 // i - Case insensitive
	FlagMultiline       byte = 0x02 // m - Multi-line mode
	FlagDotAll          byte = 0x04 // s - Dot matches newline
	FlagUnicode         byte = 0x08 // u - Unicode mode
	FlagGlobal          byte = 0x10 // g - Global search
)

// RegExpData represents a compiled regular expression
type RegExpData struct {
	Pattern string
	Flags   byte
}

// EncodeRegExp encodes a regular expression (Extension 9)
func EncodeRegExp(pattern string, flags byte) ([]byte, error) {
	if len(pattern) == 0 {
		return nil, fmt.Errorf("empty regex pattern")
	}

	// Validate pattern compiles
	_, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	patternBytes := stringToBytes(pattern)

	// Layout: header + flags + pattern_size + pattern
	headerSize := 1 + 1 + sizeOfCompressedInt(len(patternBytes))
	buf := make([]byte, headerSize+len(patternBytes))

	offset := 0

	// Header
	buf[offset] = ExtRegex
	offset++

	// Flags
	buf[offset] = flags
	offset++

	// Pattern size
	offset += writeCompressedSize(buf[offset:], len(patternBytes))

	// Pattern
	copy(buf[offset:], patternBytes)

	return buf, nil
}

// DecodeRegExp decodes Extension 9 regular expression
func DecodeRegExp(data []byte) (RegExpData, error) {
	if len(data) < 3 || data[0] != ExtRegex {
		return RegExpData{}, fmt.Errorf("invalid regex header")
	}

	offset := 1

	// Read flags
	flags := data[offset]
	offset++

	// Read pattern size
	patternSize, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return RegExpData{}, fmt.Errorf("failed to read pattern size: %w", err)
	}
	offset += bytesRead

	// Read pattern
	pattern := bytesToString(data[offset : offset+patternSize])

	return RegExpData{
		Pattern: pattern,
		Flags:   flags,
	}, nil
}

// MarshalRegExp marshals a regexp.Regexp to BEVE bytes
func MarshalRegExp(r *regexp.Regexp) ([]byte, error) {
	if r == nil {
		return nil, fmt.Errorf("nil regexp")
	}

	// Extract flags from Go regexp (limited support)
	pattern := r.String()
	flags := byte(0)

	// Go regexp doesn't expose flags directly, so we can't reliably extract them
	// Default to no flags
	return EncodeRegExp(pattern, flags)
}

// UnmarshalRegExp unmarshals BEVE regex bytes to regexp.Regexp
func UnmarshalRegExp(data []byte) (*regexp.Regexp, error) {
	regexData, err := DecodeRegExp(data)
	if err != nil {
		return nil, err
	}

	// Build Go regex pattern with flags
	pattern := regexData.Pattern

	// Go's regexp package has limited flag support
	// We can use (?flags) syntax for some flags
	if regexData.Flags&FlagCaseInsensitive != 0 {
		pattern = "(?i)" + pattern
	}
	if regexData.Flags&FlagMultiline != 0 {
		pattern = "(?m)" + pattern
	}
	if regexData.Flags&FlagDotAll != 0 {
		pattern = "(?s)" + pattern
	}

	return regexp.Compile(pattern)
}

// EncodeRegExpString is a convenience function for string patterns
func EncodeRegExpString(pattern string, caseInsensitive, multiline, dotAll bool) ([]byte, error) {
	flags := byte(0)
	if caseInsensitive {
		flags |= FlagCaseInsensitive
	}
	if multiline {
		flags |= FlagMultiline
	}
	if dotAll {
		flags |= FlagDotAll
	}

	return EncodeRegExp(pattern, flags)
}

// DecodeRegExpString decodes and returns pattern and flags as booleans
func DecodeRegExpString(data []byte) (pattern string, caseInsensitive, multiline, dotAll bool, err error) {
	regexData, err := DecodeRegExp(data)
	if err != nil {
		return "", false, false, false, err
	}

	return regexData.Pattern,
		(regexData.Flags & FlagCaseInsensitive) != 0,
		(regexData.Flags & FlagMultiline) != 0,
		(regexData.Flags & FlagDotAll) != 0,
		nil
}

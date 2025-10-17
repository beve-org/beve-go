package beve

// Extension headers following BEVE v1.0 spec section 6
// Format: 0x86 (extension base) | (extension_id << 3)
const (
	// Category 1: Performance & Optimization Extensions (0-3)
	ExtFieldIndex       byte = 0x86 // 0b00000'110 - Fast partial field access
	ExtTypedArray       byte = 0x8E // 0b00001'110 - Deduplicate field names in arrays
	ExtTypedNestedArray byte = 0x96 // 0b00010'110 - Hierarchical schema for nested objects
	ExtCompressionHint  byte = 0x9E // 0b00011'110 - Reserved for compression metadata

	// Category 2: Temporal Extensions (4-7)
	ExtTimestamp      byte = 0xA6 // 0b00100'110 - High-precision timestamp
	ExtDuration       byte = 0xAE // 0b00101'110 - Time duration
	ExtInterval       byte = 0xB6 // 0b00110'110 - Time interval (start-end)
	ExtRecurringEvent byte = 0xBE // 0b00111'110 - Cron-like recurring events

	// Category 3: Identifier & Pattern Extensions (8-11)
	ExtUUID  byte = 0xC6 // 0b01000'110 - Binary UUID/ULID
	ExtRegex byte = 0xCE // 0b01001'110 - Regular expression with flags
)

// Extension type codes for nested schema
const (
	TypeAny    byte = 0
	TypeInt    byte = 1
	TypeString byte = 2
	TypeObject byte = 3
	TypeArray  byte = 4
	TypeFloat  byte = 5
	TypeBool   byte = 6
)

// Timestamp precision flags
const (
	PrecisionSeconds      byte = 0 << 1
	PrecisionMilliseconds byte = 1 << 1
	PrecisionMicroseconds byte = 2 << 1
	PrecisionNanoseconds  byte = 3 << 1

	FlagHasTimezone byte = 0x01 // Bit 0: timezone present
)

// Nested schema constraints
const (
	MaxNestingDepth = 16 // Maximum nesting levels for Extension 2
)

// Field flags
const (
	FlagOmitEmpty byte = 0x01 // Bit 0: field can be omitted if empty
	FlagNested    byte = 0x02 // Bit 1: field is a nested object
)

// MarshalOptions controls encoding behavior
type MarshalOptions struct {
	UseTypedSchema  bool // Enable Extension 1/2 (typed arrays)
	UseFieldIndex   bool // Enable Extension 0 (field indexing)
	IncludeFallback bool // Include generic encoding for old parsers (hybrid mode)
	AutoDetect      bool // Automatically choose best encoding
	MinArraySize    int  // Minimum array size for typed encoding (default: 5)
}

// DefaultMarshalOptions provides backward-compatible defaults
var DefaultMarshalOptions = MarshalOptions{
	UseTypedSchema:  false, // Opt-in for typed schema
	UseFieldIndex:   false, // Opt-in for field index
	IncludeFallback: false, // No hybrid encoding by default
	AutoDetect:      false, // Explicit opt-in only
	MinArraySize:    5,     // Use typed arrays for N >= 5
}

// FieldSchema represents a field definition in typed array schema
type FieldSchema struct {
	Name           string
	TypeCode       byte // TypeAny, TypeInt, TypeString, etc.
	NestedSchemaID int  // Only used if TypeCode == TypeObject
}

// SchemaNode represents a schema at a specific nesting depth
type SchemaNode struct {
	ID       int
	ParentID int // Parent schema ID (-1 for root)
	Fields   []FieldSchema
}

// FieldIndexEntry represents field metadata for Extension 0
type FieldIndexEntry struct {
	Offset uint32 // Relative offset to field data start
	Size   uint16 // Field size (0 = variable length)
	Flags  byte   // FlagOmitEmpty, FlagNested, etc.
}

// Timestamp represents a high-precision timestamp with optional timezone
type Timestamp struct {
	Seconds        int64  // Unix epoch seconds
	Nanoseconds    uint32 // Sub-second precision (0-999,999,999)
	TimezoneOffset *int16 // Minutes from UTC (nil = UTC)
}

// NewTimestampUTC creates a UTC timestamp (no timezone offset)
func NewTimestampUTC(seconds int64, nanoseconds uint32) Timestamp {
	return Timestamp{
		Seconds:        seconds,
		Nanoseconds:    nanoseconds,
		TimezoneOffset: nil,
	}
}

// NewTimestampWithTZ creates a timestamp with timezone offset
func NewTimestampWithTZ(seconds int64, nanoseconds uint32, tzOffsetMinutes int16) Timestamp {
	return Timestamp{
		Seconds:        seconds,
		Nanoseconds:    nanoseconds,
		TimezoneOffset: &tzOffsetMinutes,
	}
}

// Capabilities represents parser/serializer capabilities for negotiation
type Capabilities struct {
	SupportsExtensions     bool
	SupportsFieldIndex     bool // Extension 0
	SupportsTypedArray     bool // Extension 1
	SupportsTypedNested    bool // Extension 2
	SupportsTimestamp      bool // Extension 4
	SupportsDuration       bool // Extension 5
	SupportsInterval       bool // Extension 6
	SupportsUUID           bool // Extension 8
	SupportsRegex          bool // Extension 9
	MaxNestingDepth        int  // Maximum supported nesting (0 = no limit)
	SupportsHybridEncoding bool // Supports fallback mode
}

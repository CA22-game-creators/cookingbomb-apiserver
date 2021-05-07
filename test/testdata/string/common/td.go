package testdata

var ULID = struct {
	Valid, Invalid string
}{
	Valid:   "00000000000000000000000001",
	Invalid: "invalid_ulid",
}

var UUID = struct {
	Valid, Invalid, Encrypted string
}{
	Valid:     "00000000-0000-0000-0000-000000000001",
	Invalid:   "invalid_uuid",
	Encrypted: "$2a$10$ILQMw9IZ8XLn9bBz6hboDOdZSQl21lCZd1eNjYA0trInaAbL6h21C",
}

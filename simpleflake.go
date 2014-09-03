package simpleflake

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"
)

const (
	nano = 1000 * 1000
)

var (
	// default epoch: 2000-01-01 00:00:00 +0000 UTC = 946684800000
	epoch         int64  = int64(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()) / nano
	timestampBits uint32 = 41
	randomBits    uint32 = 64 - timestampBits
)

// Generate a new 64-bit, roughly-ordered, unique ID
func New() (id uint64, err error) {
	seq, err := randomSequence()
	if err != nil {
		return
	}
	id = buildId(customTimestamp(time.Now()), seq)
	return
}

// Parse a previously generated ID
func Parse(id uint64) [2]uint64 {
	return [2]uint64{
		extractBits(id, randomBits, timestampBits) + uint64(epoch), // timestamp
		extractBits(id, 0, randomBits),                             // sequence
	}
}

// Set the epoch to a custom time
func SetCustomEpoch(t time.Time) {
	epoch = t.UTC().UnixNano() / nano
}

// Set the precision level of the timestamp
func SetCustomPrecision(bits uint32) {
	timestampBits = bits
	// reset random bit length
	randomBits = 64 - timestampBits
}

// Build a new 64-bit ID from the timestamp and random sequence
func buildId(ts int64, seq uint64) uint64 {
	return (uint64(ts) << randomBits) | seq
}

// Get a custom timestamp to be used to generate a new ID
func customTimestamp(t time.Time) int64 {
	return t.UnixNano()/nano - epoch
}

// Extract bits from a uint64
func extractBits(data uint64, shift uint32, length uint32) uint64 {
	bitmask := uint64(((1 << length) - 1) << shift)
	return ((data & bitmask) >> shift)
}

// Get a random sequence to be used to generate a new ID
func randomSequence() (seq uint64, err error) {
	// the maximum random sequence we can generate is 2^randomBits-1
	max := big.NewInt(int64((math.Pow(2, float64(randomBits))) - 1))
	random, err := rand.Int(rand.Reader, max)
	if err == nil {
		seq = random.Uint64()
	}
	return
}

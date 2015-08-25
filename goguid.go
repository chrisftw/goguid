package goguid

import (
	"fmt"
	"math/rand"
	"time"
)

// Version 4 UUIDs use a scheme relying only on random numbers. This algorithm sets the version number (4 bits) as well as two reserved bits.
// All other bits (the remaining 122 bits) are set using a random or pseudorandom data source.
// Version 4 UUIDs have the form xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx where x is any hexadecimal digit and y is one of 8, 9, A, or B (e.g., f47ac10b-58cc-4372-a567-0e02b2c3d479).
// From Wikipedia (https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29)
func GenV4() string {
	a := rand.Int63()
	b := rand.Int63()
	return fmt.Sprintf("%08x-%04x-4%03x-%x%03x-%012x", a&0xffffffff, (a>>32)&0xffff, (a>>48)&0xfff, ((a>>60)&0x3)|0x8, b&0xfff, (b>>12)&0xffffffffffff)
}

func init() {
	fmt.Print("") // Unsure why, but printing an empty string speeds up the benchmarks... -cr-
	rand.Seed(time.Now().UTC().UnixNano())
}

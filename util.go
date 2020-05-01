/*
#######
##           __         __
##       ___/ /__ ____ / /___ ____ _
##      / _  / _ `(_-</ __/ // /  ' \
##      \_,_/\_,_/___/\__/\_,_/_/_/_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package dastum

import (
	"crypto/rand"
	"fmt"
	mathrand "math/rand"
	"os"
)

// FileExists vérifie l'existence d'un fichier.
func FileExists(file string) (bool, error) {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// NewUUID génère un UUID V4.
func NewUUID() string {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		mathrand.Read(b) // Version dégradée
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//
	// https://tools.ietf.org/html/rfc4122
	//
	// The version 4 UUID is meant for generating UUIDs from truly-random or
	// pseudo-random numbers.
	//
	// typedef struct {
	//    unsigned32  time_low;
	//    unsigned16  time_mid;
	//    unsigned16  time_hi_and_version;
	//    unsigned8   clock_seq_hi_and_reserved;
	//    unsigned8   clock_seq_low;
	//    byte        node[6];
	// } uuid_t;
	//
	// The algorithm is as follows:
	//
	// - Set the two most significant bits (bits 6 and 7) of the clock_seq_hi_and_reserved
	//   to zero and one, respectively.
	//

	b[8] = (b[8] & 0x7F) | 0x40

	//
	// - Set the four most significant bits (bits 12 through 15) of the time_hi_and_version
	//   field to the 4-bit version number from Section 4.1.3.
	//

	b[6] = (b[6] & 0x0F) | 0x40

	//
	// Set all the other bits to randomly (or pseudo-randomly) chosen
	// values.
	//
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

/*
######################################################################################################## @(°_°)@ #######
*/

package ext4

import (
	"fmt"
)

// blockGroup is a structure holding the data about a single block group
//
//nolint:unused // will be used in the future, not yet
type blockGroup struct {
	inodeBitmap    *bitmap
	blockBitmap    *bitmap
	blockSize      int
	number         int
	inodeTableSize int
	firstDataBlock int
}

// blockGroupFromBytes create a blockGroup struct from bytes
// it does not load the inode table or data blocks into memory, rather holding pointers to where they are
//
//nolint:unused // will be used in the future, not yet
func blockGroupFromBytes(b []byte, blockSize, groupNumber int) (*blockGroup, error) {
	expectedSize := 2 * blockSize
	actualSize := len(b)
	if actualSize != expectedSize {
		return nil, fmt.Errorf("expected to be passed %d bytes for 2 blocks of size %d, instead received %d", expectedSize, blockSize, actualSize)
	}
	inodeBitmap := bitmapFromBytes(b[0:blockSize])
	blockBitmap := bitmapFromBytes(b[blockSize : 2*blockSize])

	bg := blockGroup{
		inodeBitmap: inodeBitmap,
		blockBitmap: blockBitmap,
		number:      groupNumber,
		blockSize:   blockSize,
	}
	return &bg, nil
}

// toBytes returns bitmaps ready to be written to disk
//
//nolint:unused // will be used in the future, not yet
func (bg *blockGroup) toBytes() ([]byte, error) {
	b := make([]byte, 2*bg.blockSize)
	inodeBitmapBytes := bg.inodeBitmap.toBytes()
	blockBitmapBytes := bg.blockBitmap.toBytes()

	b = append(b, inodeBitmapBytes...)
	b = append(b, blockBitmapBytes...)

	return b, nil
}

package nox

/*
#include <stdio.h>
*/
import "C"
import (
	"encoding/binary"
	"io"
	"os"
	"unsafe"

	"nox/v1/common/memmap"
)

var (
	cryptFile           *Binfile
	cryptFileMode       int
	cryptFileRollingXOR int32
	cryptFileNoKey      bool
	cryptFileXOREnabled bool
	cryptFileOffsSaved  []cryptFileBookmark
)

type cryptFileBookmark struct {
	Before uint32
	After  uint32
}

//export nox_xxx_cryptGetXxx
func nox_xxx_cryptGetXxx() C.int {
	return C.int(cryptFileMode)
}

//export nox_xxx_cryptSetTypeMB_426A50
func nox_xxx_cryptSetTypeMB_426A50(a1 C.int) {
	cryptFileXOREnabled = a1 != 0
}

//export nox_xxx_cryptOpen_426910
func nox_xxx_cryptOpen_426910(a1 *C.char, cmode, key C.int) int32 {
	if err := cryptFileOpen(GoString(a1), int(cmode), int(key)); err != nil {
		if !os.IsNotExist(err) {
			binFileLog.Println(err)
		}
		return 0
	}
	return 1
}

func cryptFileOpen(path string, cmode int, key int) error {
	if cryptFile != nil {
		nox_xxx_cryptClose_4269F0()
	}
	nox_xxx_cryptSetTypeMB_426A50(0)
	cryptFileNoKey = key == -1
	cryptFileRollingXOR = -1
	var fmode BinFileMode
	cryptFileMode = cmode
	if cmode == 1 {
		fmode = BinFileRO
	} else if cmode == 2 {
		cryptFileMode = 0
		fmode = BinFileRW
	} else {
		fmode = BinFileWO
	}
	f, err := BinfileOpen(path, fmode)
	if err != nil {
		return err
	}
	cryptFile = f
	if err = f.SetKey(key); err != nil {
		return err
	}
	if cmode == 2 {
		if err := cryptFile.FileSeek(0, io.SeekEnd); err != nil {
			return err
		}
	}
	cryptFileOffsSaved = cryptFileOffsSaved[:0]
	return nil
}

//export nox_xxx_cryptFlush_4268E0
func nox_xxx_cryptFlush_4268E0() C.int {
	res, _ := cryptFile.FileFlush()
	return C.int(res)
}

//export nox_xxx_cryptClose_4269F0
func nox_xxx_cryptClose_4269F0() {
	if cryptFile != nil {
		cryptFile.Close()
		nox_xxx_cryptSetTypeMB_426A50(0)
		cryptFile = nil
	}
}

//export nox_xxx_mapgenGetSomeFile_426A60
func nox_xxx_mapgenGetSomeFile_426A60() *C.FILE {
	return newFileHandle(cryptFile.file)
}

//export sub_426AA0
func sub_426AA0(a1 C.int) C.int {
	cryptFile.FileSeek(int64(a1), io.SeekCurrent)
	return 0
}

func sub_426BD0(a1 []byte) {
	if cryptFile == nil {
		return
	}
	v3 := cryptFileRollingXOR
	for i := range a1 {
		v6 := int32(a1[i])
		result := int32(memmap.Uint32(0x581450, 7288+4*uintptr(v6^int32(byte(v3)))))
		v3 = result ^ (v3 >> 8)
		cryptFileRollingXOR = int32(v3)
	}
	cryptFileRollingXOR = int32(^v3)
}

//export nox_xxx_fileReadWrite_426AC0_file3_fread
func nox_xxx_fileReadWrite_426AC0_file3_fread(a1 *C.uchar, a2 C.size_t) C.size_t {
	buf := unsafe.Slice((*byte)(unsafe.Pointer(a1)), int(a2))
	_, err := cryptFileReadWrite(buf)
	if err != nil {
		return 0
	}
	return 1
}

func cryptFileReadWrite(a1 []byte) (int, error) {
	if cryptFileMode != 0 {
		var (
			n   int
			err error
		)
		if cryptFileXOREnabled {
			n, err = cryptFile.file.Read(a1)
			netCryptXor(126, a1)
		} else {
			n, err = cryptFile.Read(a1)
		}
		if n != 0 {
			sub_426BD0(a1)
		}
		return n, err
	}
	sub_426BD0(a1)
	if cryptFileXOREnabled {
		v2 := make([]byte, len(a1))
		copy(v2, a1)
		netCryptXor(126, v2)
		return cryptFile.file.Write(v2)
	}
	return cryptFile.Write(a1)
}

//export nox_xxx_fileCryptReadCrcMB_426C20
func nox_xxx_fileCryptReadCrcMB_426C20(a1 *C.uchar, a2 C.size_t) {
	buf := unsafe.Slice((*byte)(unsafe.Pointer(a1)), int(a2))
	cryptFileReadMaybeAlign(buf)
}

func cryptFileReadMaybeAlign(a1 []byte) {
	if cryptFileMode == 1 {
		if cryptFileXOREnabled {
			cryptFile.file.Read(a1)
			netCryptXor(126, a1)
		} else if cryptFileNoKey {
			cryptFile.Read(a1)
		} else {
			cryptFile.ReadAligned(a1)
		}
	}
}

//export nox_xxx_crypt_426C90
func nox_xxx_crypt_426C90() {
	if cryptFileMode == 0 {
		if cryptFileXOREnabled {
			v2, _ := cryptFile.file.Seek(0, io.SeekCurrent)
			cryptFileOffsSaved = append(cryptFileOffsSaved, cryptFileBookmark{
				After: uint32(v2),
			})
			var b [4]byte
			binary.LittleEndian.PutUint32(b[:], uint32(v2))
			netCryptXor(126, b[:])
			cryptFile.file.Write(b[:])
		} else {
			off1, _ := cryptFile.FileFlush()
			off2 := cryptFile.Written()
			cryptFileOffsSaved = append(cryptFileOffsSaved, cryptFileBookmark{
				Before: uint32(off1),
				After:  uint32(off2),
			})
		}
	}
}

//export sub_4268F0
func sub_4268F0(off C.int) {
	cryptFile.WriteUint32At(cryptFileRollingXOR, int64(off))
}

//export nox_xxx_crypt_426D40
func nox_xxx_crypt_426D40() {
	if cryptFileMode != 0 {
		return
	}
	cnt := len(cryptFileOffsSaved)
	offs := cryptFileOffsSaved[cnt-1]
	cryptFileOffsSaved = cryptFileOffsSaved[:cnt-1]

	if cryptFileXOREnabled {
		v1, _ := cryptFile.file.Seek(0, io.SeekCurrent)
		v3 := uint32(v1 - int64(offs.After) - 4)
		cryptFile.FileSeek(int64(offs.After), io.SeekStart)
		var b [4]byte
		binary.LittleEndian.PutUint32(b[:], v3)
		netCryptXor(126, b[:])
		cryptFile.file.Write(b[:])
		cryptFile.FileSeek(v1, io.SeekStart)
	} else {
		v5 := cryptFile.Written()
		cryptFile.WriteUint32At(int32(v5-int64(offs.After)), int64(offs.Before))
	}
}

//export sub_41C200
func sub_41C200() int32 {
	var b [2]byte
	binary.LittleEndian.PutUint16(b[:], 1)
	cryptFileReadWrite(b[:2])
	if int16(binary.LittleEndian.Uint16(b[:])) > 1 {
		return 0
	}
	if v1 := 8 - cryptFile.Written()%8; v1 > 0 {
		var empty [8]byte
		cryptFileReadWrite(empty[:v1])
	}
	return 1
}

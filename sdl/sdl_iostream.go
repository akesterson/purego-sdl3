package sdl

import "unsafe"

type IOStatus uint32

const (
	IOStatusReady IOStatus = iota
	IOStatusError
	IOStatusEof
	IOStatusNotReady
	IOStatusReadOnly
	IOStatusWriteOnly
)

type IOWhence uint32

const (
	IOSeekSet IOWhence = iota
	IOSeekCur
	IOSeekEnd
)

type IOStream struct{}

// IOFromConstMem returns a read-only memory buffer for use with [IOStream] or nil on failure.
func IOFromConstMem(mem []byte) *IOStream {
	return sdlIOFromConstMem(mem, len(mem))
}

// CloseIO closes and frees an allocated [IOStream] structure.
func CloseIO(context *IOStream) bool {
	return sdlCloseIO(context)
}

// func FlushIO(context *IOStream) bool {
//	return sdlFlushIO(context)
// }

// func GetIOProperties(context *IOStream) PropertiesID {
//	return sdlGetIOProperties(context)
// }

// func GetIOSize(context *IOStream) int64 {
//	return sdlGetIOSize(context)
// }

// func GetIOStatus(context *IOStream) IOStatus {
//	return sdlGetIOStatus(context)
// }

// func IOFromDynamicMem() *IOStream {
//	return sdlIOFromDynamicMem()
// }

// IOFromFile returns an [IOStream] for the named file. The mode can be "r" for read-only.
func IOFromFile(file string, mode string) *IOStream {
	return sdlIOFromFile(file, mode)
}

func IOFromMem(mem []byte) *IOStream {
	return sdlIOFromMem(mem, len(mem))
}

// func IOprintf(context *IOStream, fmt string) uint64 {
//	return sdlIOprintf(context, fmt)
// }

// func IOvprintf(context *IOStream, fmt string, ap va_list) uint64 {
//	return sdlIOvprintf(context, fmt, ap)
// }

func LoadFile(file string, dataSize *uint64) unsafe.Pointer {
	return sdlLoadFile(file, dataSize)
}

// func LoadFile_IO(src *IOStream, datasize *uint64, closeio bool) unsafe.Pointer {
//	return sdlLoadFile_IO(src, datasize, closeio)
// }

// func OpenIO(iface *IOStreamInterface, userdata unsafe.Pointer) *IOStream {
//	return sdlOpenIO(iface, userdata)
// }

func ReadIO(context *IOStream, ptr unsafe.Pointer, size uint64) uint64 {
	return sdlReadIO(context, ptr, size)
}

// func ReadS16BE(src *IOStream, value *int16) bool {
//	return sdlReadS16BE(src, value)
// }

// func ReadS16LE(src *IOStream, value *int16) bool {
//	return sdlReadS16LE(src, value)
// }

// func ReadS32BE(src *IOStream, value *int32) bool {
//	return sdlReadS32BE(src, value)
// }

// func ReadS32LE(src *IOStream, value *int32) bool {
//	return sdlReadS32LE(src, value)
// }

// func ReadS64BE(src *IOStream, value *int64) bool {
//	return sdlReadS64BE(src, value)
// }

// func ReadS64LE(src *IOStream, value *int64) bool {
//	return sdlReadS64LE(src, value)
// }

// func ReadS8(src *IOStream, value *int8) bool {
//	return sdlReadS8(src, value)
// }

// func ReadU16BE(src *IOStream, value *uint16) bool {
//	return sdlReadU16BE(src, value)
// }

// func ReadU16LE(src *IOStream, value *uint16) bool {
//	return sdlReadU16LE(src, value)
// }

// func ReadU32BE(src *IOStream, value *uint32) bool {
//	return sdlReadU32BE(src, value)
// }

// func ReadU32LE(src *IOStream, value *uint32) bool {
//	return sdlReadU32LE(src, value)
// }

// func ReadU64BE(src *IOStream, value *uint64) bool {
//	return sdlReadU64BE(src, value)
// }

// func ReadU64LE(src *IOStream, value *uint64) bool {
//	return sdlReadU64LE(src, value)
// }

// func ReadU8(src *IOStream, value *uint8) bool {
//	return sdlReadU8(src, value)
// }

// func SaveFile(file string, data unsafe.Pointer, datasize uint64) bool {
//	return sdlSaveFile(file, data, datasize)
// }

// func SaveFile_IO(src *IOStream, data unsafe.Pointer, datasize uint64, closeio bool) bool {
//	return sdlSaveFile_IO(src, data, datasize, closeio)
// }

// func SeekIO(context *IOStream, offset int64, whence IOWhence) int64 {
//	return sdlSeekIO(context, offset, whence)
// }

// func TellIO(context *IOStream) int64 {
//	return sdlTellIO(context)
// }

// func WriteIO(context *IOStream, ptr unsafe.Pointer, size uint64) uint64 {
//	return sdlWriteIO(context, ptr, size)
// }

// func WriteS16BE(dst *IOStream, value int16) bool {
//	return sdlWriteS16BE(dst, value)
// }

// func WriteS16LE(dst *IOStream, value int16) bool {
//	return sdlWriteS16LE(dst, value)
// }

// func WriteS32BE(dst *IOStream, value int32) bool {
//	return sdlWriteS32BE(dst, value)
// }

// func WriteS32LE(dst *IOStream, value int32) bool {
//	return sdlWriteS32LE(dst, value)
// }

// func WriteS64BE(dst *IOStream, value int64) bool {
//	return sdlWriteS64BE(dst, value)
// }

// func WriteS64LE(dst *IOStream, value int64) bool {
//	return sdlWriteS64LE(dst, value)
// }

// func WriteS8(dst *IOStream, value int8) bool {
//	return sdlWriteS8(dst, value)
// }

// func WriteU16BE(dst *IOStream, value uint16) bool {
//	return sdlWriteU16BE(dst, value)
// }

// func WriteU16LE(dst *IOStream, value uint16) bool {
//	return sdlWriteU16LE(dst, value)
// }

// func WriteU32BE(dst *IOStream, value uint32) bool {
//	return sdlWriteU32BE(dst, value)
// }

// func WriteU32LE(dst *IOStream, value uint32) bool {
//	return sdlWriteU32LE(dst, value)
// }

// func WriteU64BE(dst *IOStream, value uint64) bool {
//	return sdlWriteU64BE(dst, value)
// }

// func WriteU64LE(dst *IOStream, value uint64) bool {
//	return sdlWriteU64LE(dst, value)
// }

// func WriteU8(dst *IOStream, value uint8) bool {
//	return sdlWriteU8(dst, value)
// }

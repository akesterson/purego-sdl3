package sdl

import "github.com/ebitengine/purego"

// func AddTimer(interval uint32, callback TimerCallback, userdata unsafe.Pointer) TimerID {
//	return sdlAddTimer(interval, callback, userdata)
// }

// func AddTimerNS(interval uint64, callback NSTimerCallback, userdata unsafe.Pointer) TimerID {
//	return sdlAddTimerNS(interval, callback, userdata)
// }

// func Delay(ms uint32)  {
//	sdlDelay(ms)
// }

// func DelayNS(ns uint64)  {
//	sdlDelayNS(ns)
// }

// func DelayPrecise(ns uint64)  {
//	sdlDelayPrecise(ns)
// }

func GetPerformanceCounter() uint64 {
	return sdlGetPerformanceCounter()
}

func GetPerformanceFrequency() uint64 {
	ret, _, _ := purego.SyscallN(sdlGetPerformanceFrequency)
	return uint64(ret)
}

// func GetTicks() uint64 {
//	return sdlGetTicks()
// }

// GetTicksNS returns the number of nanoseconds since SDL library initialization.
func GetTicksNS() uint64 {
	ret, _, _ := purego.SyscallN(sdlGetTicksNS)
	return uint64(ret)
}

// func RemoveTimer(id TimerID) bool {
//	return sdlRemoveTimer(id)
// }

package main

import "syncdemo/lib"

func main() {
	// lib.OnceCall()
	// lib.MutexCall()
	// lib.RWMutexCall()
	// lib.AtomicCall()
	// lib.SyncCondCall()
	lib.SyncPollCall()
}

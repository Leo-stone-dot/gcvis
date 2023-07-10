package main

type scvgtrace struct {
	ElapsedTime float64 // in seconds
	inuse       int64
	idle        int64
	sys         int64
	released    int64
	consumed    int64
}

type gctrace struct {
	ElapsedTime float64 // in seconds
	NumGC       int64
	Nproc       int64
	t1          int64
	t2          int64
	t3          int64
	t4          int64
	Heap0       int64 // heap size before, in megabytes
	Heap1       int64 // heap size after, in megabytes
	Obj         int64
	NMalloc     int64
	NFree       int64
	NSpan       int64
	NGoRoutines int64
	NBGSweep    int64
	NPauseSweep int64
	NHandoff    int64
	NHandoffCnt int64
	NSteal      int64
	NStealCnt   int64
	NProcYield  int64
	NOsYield    int64
	NSleep      int64
	// wall-clock
	STWSclock float64 //STW        : Write-Barrier - Wait for all Ps to reach a GC safe-point.
	MASclock  float64 //Concurrent : Marking
	STWMclock float64 //STW        : Mark Term     - Write Barrier off and clean up.
	// CPU time
	STWScpu      float64 //STW        : Write-Barrier
	MASAssistcpu float64 //Concurrent : Mark - Assist Time (GC performed in line with allocation)
	MASBGcpu     float64 //Concurrent : Mark - Background GC time
	MASIdlecpu   float64 //Concurrent : Mark - Idle GC time
	STWMcpu      float64 //STW        : Mark Term
}

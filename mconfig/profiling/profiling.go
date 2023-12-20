package profiling

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"
)

// Start starts the profiling server
func Start(port string, log *logger.Logger) {
	spawn := panics.GoroutineWrapperFunc(log)
	spawn("profiling.Start", func() {
		listenAddr := net.JoinHostPort("", port)
		log.Infof("Profile server listening on %s", listenAddr)
		profileRedirect := http.RedirectHandler("/debug/pprof", http.StatusSeeOther)
		http.Handle("/", profileRedirect)
		log.Error(http.ListenAndServe(listenAddr, nil))
	})
}

// TrackHeap tracks the size of the heap and dumps a profile if it passes a limit
func TrackHeap(appDir string, log *logger.Logger) {
	spawn := panics.GoroutineWrapperFunc(log)
	spawn("profiling.TrackHeap", func() {
		dumpFolder := filepath.Join(appDir, "dumps")
		err := os.MkdirAll(dumpFolder, 0700)
		if err != nil {
			log.Errorf("Could not create heap dumps folder at %s", dumpFolder)
			return
		}
		const limitInGigabytes = 7 // We want to support 8 GB RAM, so we profile at 7
		trackHeapSize(limitInGigabytes*1024*1024*1024, dumpFolder, log)
	})
}

func trackHeapSize(heapLimit uint64, dumpFolder string, log *logger.Logger) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		memStats := &runtime.MemStats{}
		runtime.ReadMemStats(memStats)
		// If we passed the expected heap limit, dump the heap profile to a file
		if memStats.HeapAlloc > heapLimit {
			dumpHeapProfile(heapLimit, dumpFolder, memStats, log)
		}
	}
}

func dumpHeapProfile(heapLimit uint64, dumpFolder string, memStats *runtime.MemStats, log *logger.Logger) {
	heapFile := filepath.Join(dumpFolder, heapDumpFileName)
	log.Infof("Saving heap statistics into %s (HeapAlloc=%d > %d=heapLimit)", heapFile, memStats.HeapAlloc, heapLimit)
	f, err := os.Create(heapFile)
	defer f.Close()
	if err != nil {
		log.Infof("Could not create heap profile: %s", err)
		return
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Infof("Could not write heap profile: %s", err)
	}
}

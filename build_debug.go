//go:build debug

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

var (
	flagCpuprofile = flag.String("cpuprofile", "", "Write CPU Profile.")
	flagTrace      = flag.String("trace", "", "Write trace data.")
	traceFile      *os.File
)

func debuginit() {
	if *flagCpuprofile != "" {
		f, err := os.Create(*flagCpuprofile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		_ = pprof.StartCPUProfile(f)
	}
	if *flagTrace != "" {
		traceFile, err := os.Create(*flagTrace)
		if err != nil {
			panic(err)
		}
		err = trace.Start(traceFile)
		if err != nil {
			panic(err)
		}
	}
}

func debugfinish() {
	if *flagCpuprofile != "" {
		pprof.StopCPUProfile()
	}
	if traceFile != nil {
		trace.Stop()
		traceFile.Close()
	}
}

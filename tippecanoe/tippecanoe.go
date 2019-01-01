package tippecanoe

/*
#include "bridge.h"
#cgo CPPFLAGS: -I./ -I/usr/local/include
#cgo LDFLAGS: -lm -lz -lsqlite3 -lpthread
#cgo CXXFLAGS: -std=c++11
*/
import "C"
import "unsafe"

func ReadFilter(fname string) unsafe.Pointer {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))
	return C.bridge_read_filter(cfname)
}

func ParseFilter(s string) unsafe.Pointer {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return C.bridge_parse_filter(cs)
}

func JSONFree(o unsafe.Pointer) {
	C.bridge_json_free(o)
}

func InitCPUs() {
	C.init_cpus()
}

func SetProjectionOrExit(optarg string) {
	coptarg := C.CString(optarg)
	defer C.free(unsafe.Pointer(coptarg))
	C.bridge_set_projection_or_exit(coptarg)
}

func SetAttributeType() {

}

func SetAttributeAccum() {

}

func MbtilesOpen(dbname string, args []string, forcetable int) unsafe.Pointer {
	cdbname := C.CString(dbname)
	defer C.free(unsafe.Pointer(cdbname))

	argv := make([]*C.char, len(args))
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}

	return C.bridge_mbtiles_open(cdbname, (**C.char)(&argv[0]), C.int(forcetable))
}

func MbtilesClose(sqlite3 unsafe.Pointer, pgm string) {
	cpgm := C.CString(pgm)
	defer C.free(unsafe.Pointer(cpgm))

	C.bridge_mbtiles_close(sqlite3, cpgm)
}

func ReadInput(
	sources []unsafe.Pointer,
	fname string,
	maxZoom int,
	minZoom int,
	baseZoom int,
	baseZoomMarkerWidth float64,
	outdb *C.void,
	outdir string,
	exclude []string,
	include []string,
	excludeAll int,
	filter *C.void,
	dropRate float64,
	buffer int,
	tmpdir string,
	gamme float64,
	readParallel int,
	forcetable int,
	attribution string,
	usesGamme bool,
	fileBBox *uint32,
	preFilter string,
	postFilter string,
	description string,
	guessMaxZoom bool,
	attributeTypes map[string]int,
	pgm string,
	attributeAccum map[string]int,
	attributeDescriptions map[string]string,
) int {
	return 0
}

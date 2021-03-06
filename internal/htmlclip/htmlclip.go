package htmlclip

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa
#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>

int clipboard_write_html(const void *bytes, NSInteger n);
NSInteger clipboard_change_count();
*/
import "C"
import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

var (
	errInvalidOperation = errors.New("invalid operation")
)


// Due to the limitation on operating systems (such as darwin),
// concurrent read can even cause panic, use a global lock to
// guarantee one read at a time.
var lock = sync.Mutex{}

// darwinHTMLWrite writes the given data to clipboard and
// returns true if success or false if failed.
func darwinHTMLWrite(buf []byte) (<-chan struct{}, error) {
	var ok C.int
	if len(buf) == 0 {
		ok = C.clipboard_write_html(unsafe.Pointer(nil), 0)
	} else {
		ok = C.clipboard_write_html(unsafe.Pointer(&buf[0]),
			C.NSInteger(len(buf)))
	}
	if ok != 0 {
		return nil, errInvalidOperation
	}

	// use unbuffered data to prevent goroutine leak
	changed := make(chan struct{}, 1)
	cnt := C.long(C.clipboard_change_count())
	go func() {
		for {
			time.Sleep(time.Second)
			cur := C.long(C.clipboard_change_count())
			if cnt != cur {
				changed <- struct{}{}
				close(changed)
				return
			}
		}
	}()
	return changed, nil
}

// lockedHTMLWrite writes bytes to the clipboard (both as text and HTML) on macOS
func lockedHTMLWrite(buf []byte) <-chan struct{} {
	lock.Lock()
	defer lock.Unlock()

	changed, err := darwinHTMLWrite(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write to clipboard err: %v\n", err)
		return nil
	}
	return changed
}

// ClipAsHTML takes copies a byte string to the clipboard as HTML
func ClipAsHTML(output []byte) {
	arch := runtime.GOOS

	// MacOS
	if arch == "darwin" {
		_ = lockedHTMLWrite(output)
		return
	}
	log.Fatalf("The following architecture is unsupported: %s", arch);
}

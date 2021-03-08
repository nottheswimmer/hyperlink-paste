#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>

int clipboard_write_html(const void *bytes, NSInteger n) {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    NSData *data = [NSData dataWithBytes: bytes length: n];
    [pasteboard clearContents];
    BOOL ok = [pasteboard setData: data forType:NSPasteboardTypeHTML];
    if (!ok) {
        return -1;
    }
    ok = [pasteboard setData: data forType:NSPasteboardTypeString];
    if (!ok) {
        return -1;
    }
    return 0;
}

NSInteger clipboard_change_count() {
    return [[NSPasteboard generalPasteboard] changeCount];
}
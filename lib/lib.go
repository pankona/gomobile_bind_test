package lib

/*
extern void goCallback();
static inline void callCallback() {
	goCallback();
}
*/
import "C"

func SetCallback() {
	C.callCallback()
}

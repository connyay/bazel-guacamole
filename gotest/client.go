package gotest

/*
#include "guacamole/client.h"
*/
import "C"

func NewClient() {
	client := C.guac_client_alloc()
	if client == nil {
		panic("no client")
	}
}

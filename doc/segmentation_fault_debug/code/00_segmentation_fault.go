package main

func main() {
	var ptr *int = nil
	*ptr = 100
}

/*
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x0 pc=0x400c02]


(gdb) run

Program received signal SIGSEGV, Segmentation fault.
main.main ()
    at ~00_segmentation_fault.go:5
5		*ptr = 100
*/

package main

func main() {
	myfunc := retfunc()

}

func retfunc() func() bool {
	return func() bool {
		return true
	}
}

package main

func main() {

	var sid int
	b := getNextChar()

	println(b)

	isCritical := checkWithDB(sid)

	println(isCritical)
}

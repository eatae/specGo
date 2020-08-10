package main

func main() {
	/* typed const */
	const constA int = 22

	/* use typed const */
	var tempA8 int8 = constA //ERROR

	/* untyped const */
	const constB = 34

	/* use untyped const */
	var tempB8 int8 = constB   // приведется к типу
	var tempB16 int16 = constB // приведется к типу
	var tempB32 int32 = constB // приведется к типу

}

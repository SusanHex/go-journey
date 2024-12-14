package main
const 
func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

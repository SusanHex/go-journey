package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

func zeroS() Entry {
	return Entry{}
}

// initialized by the user
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: Y}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

// initialized by Go
func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

// initialized by user, returns pointer
func initPtoS(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "Unknown", Year: Y}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)
	s2 := initS("Susan", "Hex", 2024)
	p2 := initPtoS("Susan", "Hex", 2024)
	fmt.Println("s2:", s2, "p2", *p2)
	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)
	pS := new(Entry)
	fmt.Println("pS:", pS)
}

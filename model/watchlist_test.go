package model

import (
	"fmt"
	"testing"
)

func TestWatchlistSave001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	setup()

	w, err := NewWatchList("Volitility")
	if err != nil {
		fmt.Println(err)
	}
	w.UpdateSymbols([]string{"IMNP", "AG", "GSV", "PTX", "GEVO"})
	fmt.Println(w)
}

func TestWatchlistLoad001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	setup()

	w, err := LoadList("Volitility")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w)
}

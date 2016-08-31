package model

import (
	"fmt"
	"testing"

	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/mariadb"
)

func setup() {
	conf := config.ReadConfigPath("./test.conf")

	pool, err := mariadb.NewPool(conf.Database)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	SetRepository(pool)
}

func TestSave001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	setup()

	s := NewSymbol("UVXY")
	s.Save()
	fmt.Println(s)
}

func TestLoadSymbol001(t *testing.T) {
	setup()

	s := NewSymbol("UVXY")
	s.Load()
	fmt.Printf("\n\nLoaded: %s\n\n", s)
}

func TestDelete001(t *testing.T) {
	setup()

	s := NewSymbol("UVXY")
	s.Load()
	fmt.Printf("\n\nDeleting: %s\n\n", s)
	s.Delete()
}

/*func TestSaveAll001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	setup()

	fmt.Println("Saving a bunch of symbols")
	s := NewSymbolCollection()
	sym1 := NewSymbol()
	sym1.Symbol = "UVXY"
	sym2 := NewSymbol()
	sym2.Symbol = "AG"
	s.Symbols = append(s.Symbols, sym1)
	s.Symbols = append(s.Symbols, sym2)
	fmt.Println(s)
	s.SaveAll()
	fmt.Println(s)
}

func TestSymbol001(t *testing.T) { //works well
	setup()
	searchSymbol := "x"
	fmt.Println("Finding Symbols like " + searchSymbol)
	s := NewSymbolCollection()
	s.FindLike(searchSymbol)
	fmt.Println(s)
}*/

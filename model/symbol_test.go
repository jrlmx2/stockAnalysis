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

	s := NewSymbol()
	s.Symbol = "UVXY"
	fmt.Println(s)
	fmt.Println(s.Save())
	fmt.Println(s)
}

func TestSaveAll001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	setup()

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

func TestSymbol001(t *testing.T) {
	setup()
	s := NewSymbolCollection()
	fmt.Println(s)
	s.FindLike("X")
	fmt.Println(s)
}

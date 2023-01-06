package store

import (
	"fmt"
	"strings"
	"testing"
)

/*
Возвращает тестовый стор и функцию, при выборе которой будут отчищатся
все таблицы, заполненные в хоте теста .
*/
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		//Проверка, переданны ли таблицы.
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}

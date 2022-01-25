package parser_test

import (
	p "parser"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	var tcs = []struct {
		s    string
		stmt *p.SelectStatement
		err  string
	}{
		{
			s: `SELECT name FROM tbl`,
			stmt: &p.SelectStatement{
				Fields:    []string{"name"},
				TableName: "tbl",
			},
		},
		{
			s: `SELECT first_name, last_name, age FROM my_table`,
			stmt: &p.SelectStatement{
				Fields:    []string{"first_name", "last_name", "age"},
				TableName: "my_table",
			},
		},
		{
			s: `SELECT * FROM my_table`,
			stmt: &p.SelectStatement{
				Fields:    []string{"*"},
				TableName: "my_table",
			},
		},
		{
			s: `select UpPer, lOWEr from MY_tabLe`,
			stmt: &p.SelectStatement{
				Fields:    []string{"UpPer", "lOWEr"},
				TableName: "MY_tabLe",
			},
		},
		// Errors
		{s: `foo`, err: `found "foo", expected SELECT`},
		{s: `SELECT !`, err: `found "!", expected field`},
		{s: `SELECT field xxx`, err: `found "xxx", expected FROM`},
		{s: `SELECT field FROM *`, err: `found "*", expected table name`},
	}
	for i, tt := range tcs {
		stmt, err := p.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  want=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nwant=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

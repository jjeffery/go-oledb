package oledb_test

import (
	"github.com/jjeffery/go-oledb/oledb"
	"testing"
)

func TestCreateDBInitialize(t *testing.T) {
	oledb.Initialize()

	dbi, err := oledb.CreateDBInitialize("SQLOLEDB")
	if err != nil {
		t.Fatal("Cannot create SQLOLEDB instance", err)
	}

	if dbi == nil {
		t.Fatal("Expected dbi to be non-nil")
	}

	defer dbi.Release()

	dbp, err := dbi.QueryDBProperties()
	if err != nil {
		t.Fatal("Cannot create DBProperties: ", err)
	}

	defer dbp.Release()

	/*	if err = dbi.Initialize(); err != nil {
			t.Fatal(err.Error())
		}
	*/
}

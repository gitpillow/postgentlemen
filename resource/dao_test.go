package resource

import (
	"github.com/gitpillow/postgentlemen/db"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

func after() {
	db.RemoveDB()
}

func before() {
	db.DBName = "test.sqlite"
	db.RemoveDB()
	db.CreateDB()
	db.Migrate()
}

func Test(t *testing.T) {
	r := Resource{
		Url:       "http://test.com/test",
		Method:    "GET",
		UrlParams: `{"name":"Tom"}`,
	}
	err := Add(&r)
	if err != nil {
		t.Error(err)
	}

	all, _ := All()
	t.Log(len(all))
	for _, r := range all {
		t.Log(r)
	}
}

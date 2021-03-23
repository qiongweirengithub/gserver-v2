package admin_test

import (
	"testing"
	"gserver.v2/admin/dbs"
)

func TestMachine(t *testing.T) {
	dbs.RawQueryField()
	dbs.StructInsert()
	dbs.RawQueryField()
}


func TestGService(t *testing.T) {
	dbs.CreateGservice("test_servcie", 1, "localhost", "8090", "test_image", "container_id", "g_web");
	dbs.GetGServices();
}




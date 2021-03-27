package admin_test

import (
	"fmt"
	"testing"

	"gserver.v2/admin/dbs"
)

func TestMachine(t *testing.T) {
	dbs.RawQueryField()
	dbs.StructInsert()
	dbs.RawQueryField()
}


func TestGService(t *testing.T) {
	// dbs.CreateGservice("test_servcie", 1, "localhost", "8090", "test_image", "container_id", "g_web");
	dbs.GetGService("10");
	fmt.Println("loading all service =============================")
	dbs.GetGServices();

}




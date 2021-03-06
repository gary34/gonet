package inspect

import (
	"agent/gsdb"
	"os"
	"testing"
	. "types"
)

func TestInspect(t *testing.T) {
	var sess Session
	sess.User = &User{Id: 1}
	gsdb.RegisterOnline(&sess, sess.User.Id)
	Inspect(1, os.Stdout)
	println("inspect field:")
	InspectField(1, ".User", os.Stdout)
	InspectField(1, ".Events", os.Stdout)
}

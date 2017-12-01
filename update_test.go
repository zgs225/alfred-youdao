package main

//
// import (
// 	"encoding/json"
// 	"os"
// 	"testing"
// 	"time"
// )
//
// func TestLoadUpdateInfo(t *testing.T) {
// 	prepare(false, "1.2.4", 1498039310)
// 	info, err := loadUpdateInfo()
// 	if err != nil {
// 		t.Error("loadUpdateInfo error: ", err)
// 	}
// 	if info.Updates {
// 		t.Error("loadUpdateInfo error: Updates")
// 	}
// 	if info.Version != "1.2.4" {
// 		t.Error("loadUpdateInfo error: Version")
// 	}
// 	if info.Time != 1498039310 {
// 		t.Error("loadUpdateInfo error: Time")
// 	}
// }
//
// func TestCheckAvailable(t *testing.T) {
// 	os.Remove(F_INFO)
// 	if b := checkAvailable(); !b {
// 		t.Errorf("checkAvaible error in no %s file\n", F_INFO)
// 	}
//
// 	prepare(true, "1.2.0", time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local).Unix())
// 	if b := checkAvailable(); !b {
// 		t.Error("checkAvaible error in before today time")
// 	}
//
// 	prepare(true, "1.2.0", time.Now().Unix())
// 	if b := checkAvailable(); b {
// 		t.Error("checkAvaible error duplicated check")
// 	}
// }
//
// func prepare(b bool, v string, t int64) {
// 	info := &updateInfo{b, v, t}
// 	f, _ := os.Create(F_INFO)
// 	n, _ := json.Marshal(info)
// 	f.Write(n)
// 	f.Close()
// }

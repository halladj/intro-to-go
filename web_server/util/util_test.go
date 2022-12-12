package util

import "testing"

var u = NewUtil()

func Test_uppercaseAndRemoveSpaces(t *testing.T) {

	resualt := u.uppercaseAndRemoveSpaces("Hamza Halladj")
	if resualt != "HAMZAHALLADJ" {
		t.Error("incorrrect result: expected HAMZAHALLADJ, got", resualt)
	}
}

func Test_cleanInput(t *testing.T) {

	resualt := u.cleanInput("!!!Hamza ?Halladj?")
	if resualt != "HAMZAHALLADJ" {
		t.Error("incorrrect result: expected HAMZAHALLADJ, got", resualt)
	}
}

func Test_VigenereEncryption(t *testing.T) {

	resualt := u.VigenereEncryption("!!!Hamza ?Halladj?", "r@ody !  !22")
	if resualt != "YOPXRVDJCOGH" {
		t.Error("incorrrect result: expected HAMZAHALLADJ, got", resualt)
	}
}

func Test_VigenereDecryption(t *testing.T) {

	resualt := u.VigenereDecryption("!!!YOPXRVDJ #@COGH??", "r@ody !  !22")
	if resualt != "HAMZAHALLADJ" {
		t.Error("incorrrect result: expected HAMZAHALLADJ, got", resualt)
	}
}

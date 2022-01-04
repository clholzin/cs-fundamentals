package storage

import "testing"

func TestAdd(t *testing.T) {
	storage := new(Storage)
	storage.dictionary = make(map[string]string)
	key := "basket"
	value := "a container used to hold or carry things, typically made from interwoven strips of cane or wire: a laundry basket."
	err := storage.Add(key, value)
	if err != nil {
		t.Error("failed to add")
	}
	err = storage.Add(key, value)
	if err == nil {
		t.Error("failed to reject")
	}
}

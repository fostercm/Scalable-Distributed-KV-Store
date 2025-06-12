package server_test

import (
	kvstore "kvstore/pkg/server"
	"testing"
)

func TestNewKVStore(t *testing.T) {
	store := kvstore.NewKVServer(4)
	if store == nil {
		t.Fatalf("Expected non-nil KVStore instance")
	}
}

func TestSetAndGet(t *testing.T) {
	store := kvstore.NewKVServer(4)

	setArgs := &kvstore.SetArgs{Key: "foo", Value: "bar"}
	setReply := &kvstore.SetReply{}
	if err := store.Set(setArgs, setReply); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	getArgs := &kvstore.GetArgs{Key: "foo"}
	getReply := &kvstore.GetReply{}
	if err := store.Get(getArgs, getReply); err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if !getReply.Exists {
		t.Errorf("Expected key to exist")
	}
	if getReply.Value != "bar" {
		t.Errorf("Expected value 'bar', got '%s'", getReply.Value)
	}
}

func TestDelete(t *testing.T) {
	store := kvstore.NewKVServer(4)

	_ = store.Set(&kvstore.SetArgs{Key: "temp", Value: "123"}, &kvstore.SetReply{})

	delArgs := &kvstore.DeleteArgs{Key: "temp"}
	delReply := &kvstore.DeleteReply{}
	if err := store.Delete(delArgs, delReply); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	getReply := &kvstore.GetReply{}
	_ = store.Get(&kvstore.GetArgs{Key: "temp"}, getReply)

	if getReply.Exists {
		t.Errorf("Expected key to be deleted")
	}
}

func TestExists(t *testing.T) {
	store := kvstore.NewKVServer(4)

	_ = store.Set(&kvstore.SetArgs{Key: "present", Value: "yes"}, &kvstore.SetReply{})

	existsArgs := &kvstore.ExistsArgs{Key: "present"}
	existsReply := &kvstore.ExistsReply{}
	if err := store.Exists(existsArgs, existsReply); err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if !existsReply.Exists {
		t.Errorf("Expected key to exist")
	}

	existsArgsMissing := &kvstore.ExistsArgs{Key: "missing"}
	existsReplyMissing := &kvstore.ExistsReply{}
	if err := store.Exists(existsArgsMissing, existsReplyMissing); err != nil {
		t.Fatalf("Exists for missing key failed: %v", err)
	}
	if existsReplyMissing.Exists {
		t.Errorf("Expected key not to exist")
	}
}

func TestLength(t *testing.T) {
	store := kvstore.NewKVServer(4)

	store.Set(&kvstore.SetArgs{Key: "a", Value: "1"}, &kvstore.SetReply{})
	store.Set(&kvstore.SetArgs{Key: "b", Value: "2"}, &kvstore.SetReply{})

	lengthReply := &kvstore.LengthReply{}
	if err := store.Length(&kvstore.LengthArgs{}, lengthReply); err != nil {
		t.Fatalf("Length failed: %v", err)
	}
	if lengthReply.Length != 2 {
		t.Errorf("Expected length 2, got %d", lengthReply.Length)
	}
}

package DB_Worker

import (
	"errors"
	"testing"

	"github.com/tidwall/buntdb"
)

func setupTestDB(t *testing.T) (*DB, func()) {
	db, err := New(":memory:")
	if err != nil {
		t.Fatalf("Failed to create in-memory db: %v", err)
	}
	cleanup := func() {
		if err := db.Close(); err != nil {
			t.Errorf("Failed to close db: %v", err)
		}
	}

	return db, cleanup
}

func TestDB_StoreAndGet(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	alias := "my-server"
	mac := "AA:BB:CC:DD:EE:FF"
	if err := db.StoreMachine(alias, mac); err != nil {
		t.Fatalf("StoreMachine() failed: %v", err)
	}
	retrievedMac, err := db.GetStoredMac(alias)
	if err != nil {
		t.Fatalf("GetStoredMac() failed: %v", err)
	}

	if retrievedMac != mac {
		t.Errorf("expected MAC %q, but got %q", mac, retrievedMac)
	}
	_, err = db.GetStoredMac("non-existent-alias")
	if err == nil {
		t.Error("expected an error when getting a non-existent alias, but got nil")
	}
	if !errors.Is(err, buntdb.ErrNotFound) {
		t.Errorf("expected error to be buntdb.ErrNotFound, but got %v", err)
	}
}

func TestDB_EditAndDelete(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	alias := "my-desktop"
	initialMac := "11:22:33:44:55:66"
	updatedMac := "FF:EE:DD:CC:BB:AA"

	// Store the initial MAC
	err := db.StoreMachine(alias, initialMac)
	if err != nil {
		return
	}

	// Edit the MAC address
	if err := db.EditMachineDetails(alias, updatedMac); err != nil {
		t.Fatalf("EditMachineDetails() failed: %v", err)
	}

	// Verify the edit
	retrievedMac, _ := db.GetStoredMac(alias)
	if retrievedMac != updatedMac {
		t.Errorf("after edit, expected MAC %q, but got %q", updatedMac, retrievedMac)
	}

	// Test deleting the entry
	if err := db.DeleteEntry(alias); err != nil {
		t.Fatalf("DeleteEntry() failed: %v", err)
	}

	// Verify the deletion
	_, err = db.GetStoredMac(alias)
	if !errors.Is(err, buntdb.ErrNotFound) {
		t.Errorf("expected ErrNotFound after deletion, but got %v", err)
	}
}

package client

import (
	"fmt"
	"testing"

	libvirt "github.com/vtolstov/go-libvirt"
	"github.com/vtolstov/go-libvirt/libvirttest"
)

func TestConnect(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)

	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
}

func TestDisconnect(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)

	err := l.Disconnect()
	if err != nil {
		t.Error(err)
	}
}

/*
func TestMigrate(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)

	var flags libvirt.DomainMigrateFlags
	flags = libvirt.DomainMigrateFlagLive |
		libvirt.DomainMigrateFlagPeerToPeer |
		libvirt.DomainMigrateFlagPersistDestination |
		libvirt.DomainMigrateFlagChangeProtection |
		libvirt.DomainMigrateFlagAbortOnError |
		libvirt.DomainMigrateFlagAutoConverge |
		libvirt.DomainMigrateFlagNonSharedDisk

	d, err := l.DomainLookupByName("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := d.Migrate("qemu+tcp://foo/system", "", flags); err != nil {
		t.Fatalf("unexpected live migration error: %v", err)
	}
}

func TestMigrateInvalidDest(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)

	var flags libvirt.DomainMigrateFlags
	flags = libvirt.DomainMigrateFlagLive |
		libvirt.DomainMigrateFlagPeerToPeer |
		libvirt.DomainMigrateFlagPersistDestination |
		libvirt.DomainMigrateFlagChangeProtection |
		libvirt.DomainMigrateFlagAbortOnError |
		libvirt.DomainMigrateFlagAutoConverge |
		libvirt.DomainMigrateFlagNonSharedDisk

	d, err := l.DomainLookupByName("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dest := ":$'"
	if err := d.Migrate(dest, "", flags); err == nil {
		t.Fatalf("expected invalid dest uri %q to fail", dest)
	}
}
*/

func TestMigrateSetMaxSpeed(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)

	d, err := l.DomainLookupByName("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := d.MigrateSetMaxSpeed(100, 0); err != nil {
		t.Fatalf("unexpected error setting max speed for migrate: %v", err)
	}
}

func TestDomains(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	domains, err := l.ListAllDomains()
	if err != nil {
		t.Error(err)
	}

	wantLen := 2
	gotLen := len(domains)
	if gotLen != wantLen {
		t.Errorf("expected %d domains to be returned, got %d", wantLen, gotLen)
	}

	for i, d := range domains {
		wantID := i + 1
		if d.ID != wantID {
			t.Errorf("expected domain ID %q, got %q", wantID, d.ID)
		}

		wantName := fmt.Sprintf("aaaaaaa-%d", i+1)
		if d.Name != wantName {
			t.Errorf("expected domain name %q, got %q", wantName, d.Name)
		}
	}
}

func TestDomainState(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	d, err := l.DomainLookupByName("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantState := libvirt.DomainState(libvirt.DomainStateRunning)
	gotState, err := d.State()
	if err != nil {
		t.Error(err)
	}

	if gotState != wantState {
		t.Errorf("expected domain state %d, got %d", wantState, gotState)
	}
}

func TestSecrets(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	secrets, err := l.ListAllSecrets()
	if err != nil {
		t.Fatal(err)
	}

	wantLen := 1
	gotLen := len(secrets)
	if gotLen != wantLen {
		t.Fatalf("expected %d secrets, got %d", wantLen, gotLen)
	}

	s := secrets[0]
	wantType := 1
	if s.UsageType != wantType {
		t.Errorf("expected usage type %d, got %d", wantType, s.UsageType)
	}

	wantID := "/tmp"
	if s.UsageID != wantID {
		t.Errorf("expected usage id %q, got %q", wantID, s.UsageID)
	}

	// 19fdc2f2-fa64-46f3-bacf-42a8aafca6dd
	wantUUID := libvirt.UUID{
		0x19, 0xfd, 0xc2, 0xf2, 0xfa, 0x64, 0x46, 0xf3,
		0xba, 0xcf, 0x42, 0xa8, 0xaa, 0xfc, 0xa6, 0xdd,
	}
	if s.UUID != wantUUID {
		t.Errorf("expected UUID %q, got %q", wantUUID, s.UUID)
	}
}

func TestStoragePool(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	wantName := "default"
	pool, err := l.StoragePoolLookupByName(wantName)
	if err != nil {
		t.Error(err)
	}

	gotName := pool.Name
	if gotName != wantName {
		t.Errorf("expected name %q, got %q", wantName, gotName)
	}

	// bb30a11c-0846-4827-8bba-3e6b5cf1b65f
	wantUUID := libvirt.UUID{
		0xbb, 0x30, 0xa1, 0x1c, 0x08, 0x46, 0x48, 0x27,
		0x8b, 0xba, 0x3e, 0x6b, 0x5c, 0xf1, 0xb6, 0x5f,
	}
	gotUUID := pool.UUID
	if gotUUID != wantUUID {
		t.Errorf("expected UUID %q, got %q", wantUUID, gotUUID)
	}
}

func TestStoragePools(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	pools, err := l.ListAllStoragePools(libvirt.StoragePoolsFlagActive)
	if err != nil {
		t.Error(err)
	}

	wantLen := 1
	gotLen := len(pools)
	if gotLen != wantLen {
		t.Errorf("expected %d storage pool, got %d", wantLen, gotLen)
	}

	wantName := "default"
	gotName := pools[0].Name
	if gotName != wantName {
		t.Errorf("expected name %q, got %q", wantName, gotName)
	}

	// bb30a11c-0846-4827-8bba-3e6b5cf1b65f
	wantUUID := libvirt.UUID{
		0xbb, 0x30, 0xa1, 0x1c, 0x08, 0x46, 0x48, 0x27,
		0x8b, 0xba, 0x3e, 0x6b, 0x5c, 0xf1, 0xb6, 0x5f,
	}
	gotUUID := pools[0].UUID
	if gotUUID != wantUUID {
		t.Errorf("expected UUID %q, got %q", wantUUID, gotUUID)
	}
}

func TestStoragePoolRefresh(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	pool, err := l.StoragePoolLookupByName("default")
	if err != nil {
		t.Error(err)
	}

	err = pool.Refresh(0)
	if err != nil {
		t.Error(err)
	}
}

func TestUndefine(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	d, err := l.DomainLookupByName("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var flags libvirt.DomainUndefineFlags
	if err := d.Undefine(flags); err != nil {
		t.Fatalf("unexpected undefine error: %v", err)
	}
}

func TestDestroy(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	var flags libvirt.DomainDestroyFlags
	d, err := l.DomainLookupByName("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := d.Destroy(flags); err != nil {
		t.Fatalf("unexpected destroy error: %v", err)
	}
}

func TestVersion(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	version, err := l.Version()
	if err != nil {
		t.Error(err)
	}

	expected := "1.3.4"
	if version != expected {
		t.Errorf("expected version %q, got %q", expected, version)
	}
}

func TestDefineXML(t *testing.T) {
	conn := libvirttest.New()
	l := New(conn)
	err := l.Connect()
	if err != nil {
		t.Error(err)
	}
	defer l.Disconnect()

	var flags libvirt.DomainDefineXMLFlags
	var buf string
	if err := l.DomainDefineXML(buf, flags); err != nil {
		t.Fatalf("unexpected define error: %v", err)
	}
}

package uimetrics

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func writeJSON(t *testing.T, dir string, name string, v any) string {
	t.Helper()
	p := filepath.Join(dir, name)
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if err := os.WriteFile(p, b, 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	return p
}

func TestDefault_NotZero(t *testing.T) {
	d := Default()
	if d.List.Margin == 0 || d.List.ItemH == 0 || d.List.ItemGap == 0 {
		t.Fatalf("unexpected zero in defaults: %+v", d.List)
	}
	if len(d.List.HeaderColumnsWeapons) < 3 || len(d.List.RowColumnsWeapons) < 3 {
		t.Fatalf("expected columns arrays to be seeded")
	}
}

func TestLoadOrDefault_MasterFallback(t *testing.T) {
	td := t.TempDir()
	// master only
	master := Default()
	master.List.Margin = 42
	mp := writeJSON(t, td, "mst.json", master)

	got := LoadOrDefault("/no/such/user.json", mp)
	if got.List.Margin != 42 {
		t.Fatalf("want Margin=42, got=%d", got.List.Margin)
	}
}

func TestLoadOrDefault_UserOverrides(t *testing.T) {
	td := t.TempDir()
	// master
	master := Default()
	master.List.Margin = 24
	mp := writeJSON(t, td, "mst.json", master)
	// user overrides margin only
	user := Default()
	user.List.Margin = 33
	up := writeJSON(t, td, "usr.json", user)

	got := LoadOrDefault(up, mp)
	if got.List.Margin != 33 {
		t.Fatalf("want Margin=33, got=%d", got.List.Margin)
	}
}

package service

import (
    "testing"
    "github.com/hajimehoshi/ebiten/v2"
)

// helper: stateful key provider
type keyState struct{ m map[ebiten.Key]bool }
func (s *keyState) set(k ebiten.Key, v bool) { s.m[k] = v }
func (s *keyState) isDown(k ebiten.Key) bool { return s.m[k] }

func TestPressAndDown(t *testing.T) {
    in := NewInput()
    ks := &keyState{m: map[ebiten.Key]bool{}}

    // frame 1: none
    in.SnapshotWith(ks.isDown)
    if in.Press(Confirm) || in.Down(Confirm) { t.Fatalf("unexpected press/down on idle") }

    // frame 2: Z down -> Confirm press
    ks.set(ebiten.KeyZ, true)
    in.SnapshotWith(ks.isDown)
    if !in.Press(Confirm) { t.Fatalf("expected Confirm press on rising edge") }
    if !in.Down(Confirm)  { t.Fatalf("expected Confirm down true when pressed") }

    // frame 3: keep Z down -> no press, still down
    in.SnapshotWith(ks.isDown)
    if in.Press(Confirm) { t.Fatalf("unexpected Confirm press while held") }
    if !in.Down(Confirm) { t.Fatalf("expected Confirm down while held") }

    // frame 4: release -> no press, not down
    ks.set(ebiten.KeyZ, false)
    in.SnapshotWith(ks.isDown)
    if in.Press(Confirm) { t.Fatalf("unexpected Confirm press on release") }
    if in.Down(Confirm)  { t.Fatalf("unexpected Confirm down on release") }
}

func TestMenuLongPressPressOnce(t *testing.T) {
    in := NewInput()
    ks := &keyState{m: map[ebiten.Key]bool{}}
    // hold Tab for 3 frames
    ks.set(ebiten.KeyTab, true)
    in.SnapshotWith(ks.isDown)
    if !in.Press(Menu) { t.Fatalf("expected Menu press on first frame") }
    for i := 0; i < 2; i++ {
        in.SnapshotWith(ks.isDown)
        if in.Press(Menu) { t.Fatalf("unexpected Menu press on hold frame %d", i) }
        if !in.Down(Menu) { t.Fatalf("expected Menu down on hold frame %d", i) }
    }
}

func TestTerrainShiftMapping(t *testing.T) {
    in := NewInput()
    ks := &keyState{m: map[ebiten.Key]bool{}}

    // frame 1: Shift+2 -> Def2
    ks.set(ebiten.KeyShift, true)
    ks.set(ebiten.Key2, true)
    in.SnapshotWith(ks.isDown)
    if !in.Press(TerrainDef2) { t.Fatalf("expected TerrainDef2 press on Shift+2") }
    if in.Press(TerrainAtt2) { t.Fatalf("did not expect TerrainAtt2 press on Shift+2") }

    // frame 2: release then 2 only -> Att2
    ks.set(ebiten.KeyShift, false)
    ks.set(ebiten.Key2, false)
    in.SnapshotWith(ks.isDown)
    // next press
    ks.set(ebiten.Key2, true)
    in.SnapshotWith(ks.isDown)
    if !in.Press(TerrainAtt2) { t.Fatalf("expected TerrainAtt2 press on 2") }
    if in.Press(TerrainDef2) { t.Fatalf("did not expect TerrainDef2 press on 2") }
}


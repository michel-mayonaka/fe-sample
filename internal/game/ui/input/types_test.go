package input

import (
    "testing"
    "github.com/hajimehoshi/ebiten/v2"
    gamesvc "ui_sample/internal/game/service"
)

// helper: stateful key provider reused via service.Input#SnapshotWith
type keyState struct{ m map[ebiten.Key]bool }
func (s *keyState) set(k ebiten.Key, v bool) { s.m[k] = v }
func (s *keyState) isDown(k ebiten.Key) bool { return s.m[k] }

func TestServiceAdapter_PressDown(t *testing.T) {
    s := gamesvc.NewInput()
    ks := &keyState{m: map[ebiten.Key]bool{}}
    r := WrapService(s)

    // frame1: idle
    s.SnapshotWith(ks.isDown)
    if r.Press(Confirm) || r.Down(Confirm) { t.Fatalf("unexpected press/down on idle") }

    // rising edge
    ks.set(ebiten.KeyZ, true)
    s.SnapshotWith(ks.isDown)
    if !r.Press(Confirm) { t.Fatalf("expected Confirm press on rising edge") }
    if !r.Down(Confirm)  { t.Fatalf("expected Confirm down when pressed") }

    // hold
    s.SnapshotWith(ks.isDown)
    if r.Press(Confirm) { t.Fatalf("no press while held") }
    if !r.Down(Confirm) { t.Fatalf("down should remain true while held") }

    // release
    ks.set(ebiten.KeyZ, false)
    s.SnapshotWith(ks.isDown)
    if r.Press(Confirm) { t.Fatalf("no press on release") }
    if r.Down(Confirm)  { t.Fatalf("not down on release") }
}

func TestServiceAdapter_TerrainShiftMapping(t *testing.T) {
    s := gamesvc.NewInput()
    ks := &keyState{m: map[ebiten.Key]bool{}}
    r := WrapService(s)

    // Shift+2 => Def2
    ks.set(ebiten.KeyShift, true)
    ks.set(ebiten.Key2, true)
    s.SnapshotWith(ks.isDown)
    if !r.Press(TerrainDef2) { t.Fatalf("expected TerrainDef2 on Shift+2") }
    if r.Press(TerrainAtt2)  { t.Fatalf("did not expect TerrainAtt2 on Shift+2") }

    // release then 2 only => Att2
    ks.set(ebiten.KeyShift, false)
    ks.set(ebiten.Key2, false)
    s.SnapshotWith(ks.isDown)
    ks.set(ebiten.Key2, true)
    s.SnapshotWith(ks.isDown)
    if !r.Press(TerrainAtt2) { t.Fatalf("expected TerrainAtt2 on 2") }
}


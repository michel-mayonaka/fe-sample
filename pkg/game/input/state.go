package input

// ControlState は 1 フレーム分の論理入力のスナップショットです。
// UI はこの“意味”を参照し、物理入力の違いを意識しません。
type ControlState struct {
    Up       bool
    Down     bool
    Left     bool
    Right    bool
    Confirm  bool
    Cancel   bool
    Menu     bool
    Next     bool
    Prev     bool
    OpenWeapons  bool
    OpenItems    bool
    EquipToggle  bool
    Slot1    bool
    Slot2    bool
    Slot3    bool
    Slot4    bool
    Slot5    bool
    Unassign bool
    // Terrain
    TerrainAtt1 bool
    TerrainAtt2 bool
    TerrainAtt3 bool
    TerrainDef1 bool
    TerrainDef2 bool
    TerrainDef3 bool
}

// Equal は 2 つの状態が同一かどうかを返します。
func (c ControlState) Equal(o ControlState) bool {
    return c.Up == o.Up && c.Down == o.Down && c.Left == o.Left && c.Right == o.Right &&
        c.Confirm == o.Confirm && c.Cancel == o.Cancel && c.Menu == o.Menu &&
        c.Next == o.Next && c.Prev == o.Prev && c.OpenWeapons == o.OpenWeapons &&
        c.OpenItems == o.OpenItems && c.EquipToggle == o.EquipToggle &&
        c.Slot1 == o.Slot1 && c.Slot2 == o.Slot2 && c.Slot3 == o.Slot3 &&
        c.Slot4 == o.Slot4 && c.Slot5 == o.Slot5 && c.Unassign == o.Unassign &&
        c.TerrainAtt1 == o.TerrainAtt1 && c.TerrainAtt2 == o.TerrainAtt2 && c.TerrainAtt3 == o.TerrainAtt3 &&
        c.TerrainDef1 == o.TerrainDef1 && c.TerrainDef2 == o.TerrainDef2 && c.TerrainDef3 == o.TerrainDef3
}

// Set は指定 Action の値を設定します。
func (c *ControlState) Set(a Action, v bool) {
    switch a {
    case ActionUp:
        c.Up = v
    case ActionDown:
        c.Down = v
    case ActionLeft:
        c.Left = v
    case ActionRight:
        c.Right = v
    case ActionConfirm:
        c.Confirm = v
    case ActionCancel:
        c.Cancel = v
    case ActionMenu:
        c.Menu = v
    case ActionNext:
        c.Next = v
    case ActionPrev:
        c.Prev = v
    case ActionOpenWeapons:
        c.OpenWeapons = v
    case ActionOpenItems:
        c.OpenItems = v
    case ActionEquipToggle:
        c.EquipToggle = v
    case ActionSlot1:
        c.Slot1 = v
    case ActionSlot2:
        c.Slot2 = v
    case ActionSlot3:
        c.Slot3 = v
    case ActionSlot4:
        c.Slot4 = v
    case ActionSlot5:
        c.Slot5 = v
    case ActionUnassign:
        c.Unassign = v
    case ActionTerrainAtt1:
        c.TerrainAtt1 = v
    case ActionTerrainAtt2:
        c.TerrainAtt2 = v
    case ActionTerrainAtt3:
        c.TerrainAtt3 = v
    case ActionTerrainDef1:
        c.TerrainDef1 = v
    case ActionTerrainDef2:
        c.TerrainDef2 = v
    case ActionTerrainDef3:
        c.TerrainDef3 = v
    }
}

// Get は指定 Action の現在値を返します。
func (c ControlState) Get(a Action) bool {
    switch a {
    case ActionUp:
        return c.Up
    case ActionDown:
        return c.Down
    case ActionLeft:
        return c.Left
    case ActionRight:
        return c.Right
    case ActionConfirm:
        return c.Confirm
    case ActionCancel:
        return c.Cancel
    case ActionMenu:
        return c.Menu
    case ActionNext:
        return c.Next
    case ActionPrev:
        return c.Prev
    case ActionOpenWeapons:
        return c.OpenWeapons
    case ActionOpenItems:
        return c.OpenItems
    case ActionEquipToggle:
        return c.EquipToggle
    case ActionSlot1:
        return c.Slot1
    case ActionSlot2:
        return c.Slot2
    case ActionSlot3:
        return c.Slot3
    case ActionSlot4:
        return c.Slot4
    case ActionSlot5:
        return c.Slot5
    case ActionUnassign:
        return c.Unassign
    case ActionTerrainAtt1:
        return c.TerrainAtt1
    case ActionTerrainAtt2:
        return c.TerrainAtt2
    case ActionTerrainAtt3:
        return c.TerrainAtt3
    case ActionTerrainDef1:
        return c.TerrainDef1
    case ActionTerrainDef2:
        return c.TerrainDef2
    case ActionTerrainDef3:
        return c.TerrainDef3
    default:
        return false
    }
}


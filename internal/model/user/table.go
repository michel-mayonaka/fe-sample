package user

// Table はユーザのキャラクター状態を ID で引ける索引です。
type Table struct {
	byID map[string]Character
	rows []Character
}

// NewTable は行スライスから索引付きのテーブルを構築します。
func NewTable(rows []Character) *Table {
	t := &Table{byID: make(map[string]Character, len(rows)), rows: append([]Character(nil), rows...)}
	for _, c := range rows {
		t.byID[c.ID] = c
	}
	return t
}

// Find はIDでユーザ状態を返します。
func (t *Table) Find(id string) (Character, bool) {
	if t == nil {
		return Character{}, false
	}
	c, ok := t.byID[id]
	return c, ok
}

// ByID は内部マップを返します（読み取り専用の想定）。
func (t *Table) ByID() map[string]Character { return t.byID }

// Slice は定義順のスライスを返します。
func (t *Table) Slice() []Character { return append([]Character(nil), t.rows...) }

// UpdateCharacter は ID が一致するレコードを更新します。
func (t *Table) UpdateCharacter(c Character) {
	if t == nil {
		return
	}
	if _, ok := t.byID[c.ID]; !ok {
		return
	}
	t.byID[c.ID] = c
	for i := range t.rows {
		if t.rows[i].ID == c.ID {
			t.rows[i] = c
			break
		}
	}
}

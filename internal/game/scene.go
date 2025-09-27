package game

import "github.com/hajimehoshi/ebiten/v2"

// Scene は最小のシーンインタフェースです。
// Update の戻り値 next が非 nil の場合、Push/Replace 等は呼び出し側の戦略に従います。
type Scene interface {
    Update(ctx *Ctx) (next Scene, err error)
    Draw(screen *ebiten.Image)
}

// SceneStack は単純な LIFO スタックです。
type SceneStack struct{ stack []Scene }

// Current は先頭（最上位）シーンを返します。空なら nil。
func (s *SceneStack) Current() Scene {
    if len(s.stack) == 0 {
        return nil
    }
    return s.stack[len(s.stack)-1]
}

// Push は新しいシーンを先頭に積みます（nil は無視）。
func (s *SceneStack) Push(sc Scene) {
    if sc != nil {
        s.stack = append(s.stack, sc)
    }
}

// Pop は先頭シーンを外して返します。空なら nil。
func (s *SceneStack) Pop() Scene {
    n := len(s.stack)
    if n == 0 {
        return nil
    }
    top := s.stack[n-1]
    s.stack = s.stack[:n-1]
    return top
}

// Replace は先頭シーンを置換します。空の場合は Push と同等です。
func (s *SceneStack) Replace(sc Scene) {
    if sc == nil {
        return
    }
    n := len(s.stack)
    if n == 0 {
        s.stack = append(s.stack, sc)
        return
    }
    s.stack[n-1] = sc
}

// Size は積まれているシーン数を返します。
func (s *SceneStack) Size() int { return len(s.stack) }


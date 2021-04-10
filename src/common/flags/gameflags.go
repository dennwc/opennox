package noxflags

import (
	"encoding/json"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

var (
	game       GameFlag
	gameChange []func(f GameFlag)
	gameSet    []func(f GameFlag)
	gameUnset  []func(f GameFlag)
	gameByName = make(map[string]GameFlag)
)

func init() {
	cur := GameFlag(1)
	for i := 0; i < 32; i++ {
		gameByName[cur.String()] = cur
		cur <<= 1
	}
}

const (
	GameHost        = GameFlag(0x1)
	GameFlag2       = GameFlag(0x2)
	GameFlag3       = GameFlag(0x4)
	GameFlag4       = GameFlag(0x8)
	GameFlag5       = GameFlag(0x10)
	GameFlag6       = GameFlag(0x20)
	GameFlag7       = GameFlag(0x40)
	GameFlag8       = GameFlag(0x80)
	GameFlag9       = GameFlag(0x100)
	GameFlag10      = GameFlag(0x200)
	GameFlag11      = GameFlag(0x400)
	GameSolo        = GameFlag(0x800)
	GameFlag13      = GameFlag(0x1000)
	GameServerXxx   = GameFlag(0x2000)
	GameFlag15      = GameFlag(0x4000)
	GameFlag16      = GameFlag(0x8000)
	GameNotQuest    = GameFlag(0x10000)
	GameFlag18      = GameFlag(0x20000)
	GamePause       = GameFlag(0x40000)
	GameFlag20      = GameFlag(0x80000)
	GameFlag21      = GameFlag(0x100000)
	GameFlag22      = GameFlag(0x200000)
	GameFlag23      = GameFlag(0x400000)
	GameFlag24      = GameFlag(0x800000)
	GameFlag25      = GameFlag(0x1000000)
	GameFlag26      = GameFlag(0x2000000)
	GameSuddenDeath = GameFlag(0x4000000)
	GameFlag28      = GameFlag(0x8000000)
	GameFlag29      = GameFlag(0x10000000)
	GameFlag30      = GameFlag(0x20000000)
	GameFlag31      = GameFlag(0x40000000)
	GameFlag32      = GameFlag(0x80000000)
)

const (
	GameServerSettings = GameFlag(0x7FFF0)
)

var (
	_ json.Marshaler   = GameFlag(0)
	_ json.Unmarshaler = (*GameFlag)(nil)
)

func ParseGameFlag(s string) (GameFlag, error) {
	if strings.Contains(s, ",") {
		var out GameFlag
		for _, s := range strings.Split(s, ",") {
			v, err := ParseGameFlag(s)
			if err != nil {
				return out, err
			}
			out |= v
		}
		return out, nil
	} else if strings.Contains(s, "|") {
		var out GameFlag
		for _, s := range strings.Split(s, "|") {
			v, err := ParseGameFlag(s)
			if err != nil {
				return out, err
			}
			out |= v
		}
		return out, nil
	}
	if v, ok := gameByName[s]; ok {
		return v, nil
	}
	s = strings.TrimPrefix(s, "Game")
	if v, ok := gameByName[s]; ok {
		return v, nil
	}
	v, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return 0, fmt.Errorf("cannot parse game flag: %q", s)
	}
	return GameFlag(v), nil
}

// GameFlag is a type for Nox game flags.
type GameFlag uint

func (f *GameFlag) UnmarshalJSON(p []byte) error {
	var arr []string
	err := json.Unmarshal(p, &arr)
	if err == nil {
		*f = 0
		for _, s := range arr {
			v, err := ParseGameFlag(s)
			if err != nil {
				return err
			}
			*f |= v
		}
		return nil
	}
	var v uint32
	if err2 := json.Unmarshal(p, &v); err2 != nil {
		return err
	}
	*f = GameFlag(v)
	return nil
}

func (f GameFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.SplitString())
}

func (f GameFlag) String() string {
	switch f {
	case GameHost:
		return "Host"
	case GameFlag2:
		return "Flag2"
	case GameFlag3:
		return "Flag3"
	case GameFlag4:
		return "Flag4"
	case GameFlag5:
		return "Flag5"
	case GameFlag6:
		return "Flag6"
	case GameFlag7:
		return "Flag7"
	case GameFlag8:
		return "Flag8"
	case GameFlag9:
		return "Flag9"
	case GameFlag10:
		return "Flag10"
	case GameFlag11:
		return "Flag11"
	case GameSolo:
		return "Solo"
	case GameFlag13:
		return "Flag13"
	case GameServerXxx:
		return "ServerXxx"
	case GameFlag15:
		return "Flag15"
	case GameFlag16:
		return "Flag16"
	case GameNotQuest:
		return "NotQuest"
	case GameFlag18:
		return "Flag18"
	case GamePause:
		return "Pause"
	case GameFlag20:
		return "Flag20"
	case GameFlag21:
		return "Flag21"
	case GameFlag22:
		return "Flag22"
	case GameFlag23:
		return "Flag23"
	case GameFlag24:
		return "Flag24"
	case GameFlag25:
		return "Flag25"
	case GameFlag26:
		return "Flag26"
	case GameSuddenDeath:
		return "SuddenDeath"
	case GameFlag28:
		return "Flag28"
	case GameFlag29:
		return "Flag29"
	case GameFlag30:
		return "Flag30"
	case GameFlag31:
		return "Flag31"
	case GameFlag32:
		return "Flag32"
	}
	if bits.OnesCount(uint(f)) == 1 {
		return fmt.Sprintf("GameFlag(0x%x)", uint(f))
	}
	return strings.Join(f.SplitString(), " | ")
}

func (f GameFlag) Split() []GameFlag {
	var (
		list []GameFlag
		cur  = GameFlag(1)
	)
	for f != 0 {
		if f&0x1 != 0 {
			list = append(list, cur)
		}
		cur <<= 1
		f >>= 1
	}
	return list
}

func (f GameFlag) SplitString() []string {
	var arr []string
	for _, v := range f.Split() {
		arr = append(arr, v.String())
	}
	return arr
}

// Has checks if one of the flags is set.
func (f GameFlag) Has(f2 GameFlag) bool {
	return f&f2 != 0
}

// HasAll checks if all of the flags are set.
func (f GameFlag) HasAll(f2 GameFlag) bool {
	return f&f2 == f2
}

// GetGame returns all game flags.
func GetGame() GameFlag {
	return game
}

// ResetGame resets all game flags.
func ResetGame() {
	game = 0
}

// HasGame check if one of the game flags is set.
func HasGame(f GameFlag) bool {
	return game.Has(f)
}

// SetGame sets one or more game flags.
func SetGame(f GameFlag) {
	if game&f == f {
		return // already set
	}
	game |= f
	for _, fnc := range gameChange {
		fnc(f)
	}
	for _, fnc := range gameSet {
		fnc(f)
	}
}

// UnsetGame resets one or more game flags.
func UnsetGame(f GameFlag) {
	if game&f == 0 {
		return // already unset
	}
	game &= ^f
	for _, fnc := range gameChange {
		fnc(f)
	}
	for _, fnc := range gameUnset {
		fnc(f)
	}
}

// OnGameChange registers a hook that triggers when game flags change.
func OnGameChange(fnc func(f GameFlag)) {
	gameChange = append(gameChange, fnc)
}

// OnGameSet registers a hook that triggers when a game flag is set.
func OnGameSet(fnc func(f GameFlag)) {
	gameSet = append(gameSet, fnc)
}

// OnGameUnset registers a hook that triggers when a game flag is unset.
func OnGameUnset(fnc func(f GameFlag)) {
	gameUnset = append(gameUnset, fnc)
}

package mapv0

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"

	"nox/v1/common/types"
	"nox/v1/server/script"
)

var (
	rtPointf = reflect.TypeOf(types.Pointf{})
)

func New(vm *lua.LState, tm *script.Timers, g script.Game) *lua.LTable {
	a := &api{
		s: vm, g: g, tm: tm,
		root:    vm.NewTable(),
		methods: make(map[string]*method),
	}
	a.initMeta()
	a.init()
	return a.root
}

type api struct {
	g    script.Game
	tm   *script.Timers
	s    *lua.LState
	root *lua.LTable
	meta struct {
		tostring *lua.LFunction
		metaWaypoint
		metaWall
		metaObject
		metaUnit
		metaPlayer
		metaTimer
	}
	methods map[string]*method
}

func (vm *api) newMeta(key string) *lua.LTable {
	t := vm.s.NewTable()
	t.Metatable = t
	t.RawSetString("__tostring", vm.meta.tostring)
	if key != "" {
		vm.root.RawSetString(key, t)
	}
	return t
}

func (vm *api) initMeta() {
	vm.meta.tostring = vm.newObjMethod("__tostring", func(obj interface{}) string {
		if s, ok := obj.(fmt.Stringer); ok {
			return s.String()
		}
		return fmt.Sprintf("%T", obj)
	}).lua
	vm.initMetaWaypoint()
	vm.initMetaWall()
	vm.initMetaObject()
	vm.initMetaUnit()
	vm.initMetaPlayer()
	vm.initMetaTimer()
}

func (vm *api) init() {
	vm.initWaypoint()
	vm.initWall()
	vm.initWallGroup()
	vm.initObject()
	vm.initObjectGroup()
	vm.initObjectType()
	vm.initUnit()
	vm.initUnitGroup()
	vm.initPlayer()
	vm.initTimer()
}

// receiverValue is a pseudo-value that helps LUA chaining by returning the receiver itself, without affecting its type.
type receiverValue struct{}

type method struct {
	name string
	this reflect.Type
	fnc  reflect.Value
	lua  *lua.LFunction
}

func (m *method) Call(vm *api, s *lua.LState) int {
	// check the receiver first, it must implement an interface
	lthis := s.CheckUserData(1)
	if lthis == nil {
		s.Error(lua.LString(fmt.Sprintf("object is nil for method %s.%s", m.this, m.name)), 0)
		return 0
	}
	this := lthis.Value
	if this == nil {
		s.Error(lua.LString(fmt.Sprintf("object is nil for method %s.%s", m.this, m.name)), 0)
		return 0
	}
	tv := reflect.ValueOf(this)
	if !tv.Type().Implements(m.this) {
		s.Error(lua.LString(fmt.Sprintf("object %T doesn't implement for method %s.%s", tv.Type().Name(), m.this, m.name)), 0)
		return 0
	}

	ft := m.fnc.Type()
	// construct the argument list
	in := make([]reflect.Value, ft.NumIn())
	li := 1 // LUA arguments may be longer, e.g. for positions
	for i := range in {
		if i == 0 { // receiver
			in[i] = tv.Convert(m.this)
			li++
			continue
		}
		argt := ft.In(i)
		var argv reflect.Value
		// try to match the argument type exactly
		// mostly used for custom handling of some well-known types
		switch argt {
		case rtPointf:
			// two cases: a pair of coordinates or a Positioner
			switch s.Get(li).(type) {
			case lua.LNumber:
				x := s.CheckNumber(li + 0)
				y := s.CheckNumber(li + 1)
				li++
				argv = reflect.ValueOf(types.Pointf{X: float32(x), Y: float32(y)})
			default:
				pos, ok := s.CheckUserData(li).Value.(script.Positioner)
				if !ok {
					s.ArgError(li, "argument doesn't doesn't have a position")
					return 0
				}
				argv = reflect.ValueOf(pos.Pos())
			}
		default:
			// match by argument kind
			switch argt.Kind() {
			case reflect.Interface:
				if s.Get(li) == lua.LNil {
					argv = reflect.Zero(argt)
				} else {
					lv := s.CheckUserData(li)
					if lv == nil {
						s.ArgError(li, fmt.Sprintf("argument is nil"))
						return 0
					}
					arg := lv.Value
					if arg == nil {
						s.ArgError(li, fmt.Sprintf("argument is nil"))
						return 0
					}
					argv = reflect.ValueOf(arg)
					if !argv.Type().Implements(argt) {
						s.ArgError(li, fmt.Sprintf("argument doesn't implement interface %s", argt.Name()))
						return 0
					}
					argv = argv.Convert(argt)
				}
			case reflect.String:
				lv := s.CheckString(li)
				argv = reflect.ValueOf(lv).Convert(argt)
			case reflect.Int, reflect.Uint, reflect.Float32, reflect.Float64:
				lv := s.CheckNumber(li)
				argv = reflect.ValueOf(lv).Convert(argt)
			case reflect.Bool:
				lv := s.CheckBool(li)
				argv = reflect.ValueOf(lv).Convert(argt)
			default:
				// that's a programming error, not script error
				panic(fmt.Errorf("unsupported argument type %v", argt))
			}
		}
		in[i] = argv
		li++
	}

	// call the method
	out := m.fnc.Call(in)
	if len(out) == 0 {
		return 0
	}

	// convert returns
	// TODO: support "v, ok := fnc()" pattern
	// TODO: support "v, err := fnc()" pattern
	var ln int
	for _, v := range out {
		var lv lua.LValue
		switch v := v.Interface().(type) {
		case nil:
			lv = lua.LNil
		case receiverValue:
			lv = lthis
		case int:
			lv = lua.LNumber(v)
		case uint:
			lv = lua.LNumber(v)
		case float32:
			lv = lua.LNumber(v)
		case float64:
			lv = lua.LNumber(v)
		case string:
			lv = lua.LString(v)
		case bool:
			lv = lua.LBool(v)
		// order is important here! larger interfaces should go first
		case script.Unit:
			lv = vm.newUnit(v)
		case *script.UnitGroup:
			lv = vm.newUnitGroup(v)
		case script.Object:
			lv = vm.newObject(v)
		case *script.ObjectGroup:
			lv = vm.newObjectGroup(v)
		case script.ObjectType:
			lv = vm.newObjectType(v)
		case script.Wall:
			lv = vm.newWall(v)
		case *script.WallGroup:
			lv = vm.newWallGroup(v)
		case script.Waypoint:
			lv = vm.newWaypoint(v)
		case *script.Timer:
			lv = vm.newTimer(v)
		case script.Duration:
			lv = vm.newDuration(v)
		default:
			// that's a programming error, not script error
			panic(fmt.Errorf("unsupported argument type %T", v))
		}
		s.Push(lv)
		ln++
	}
	return ln
}

func (vm *api) newObjMethod(name string, fnc interface{}) *method {
	rv := reflect.ValueOf(fnc)
	rt := rv.Type()
	if rt.Kind() != reflect.Func {
		panic("must be a function")
	} else if rt.NumIn() < 1 {
		panic("must have at least a receiver")
	} else if rt.IsVariadic() {
		panic("variadic functions are not supported yet")
	}
	this := rt.In(0)
	if this.Kind() != reflect.Interface {
		panic("receiver must be an interface")
	}
	m := &method{
		name: name,
		this: this,
		fnc:  rv,
	}
	m.lua = vm.s.NewFunction(func(s *lua.LState) int {
		return m.Call(vm, s)
	})
	return m
}

func (vm *api) registerObjMethod(name string, fnc interface{}) {
	if m, ok := vm.methods[name]; ok {
		panic(fmt.Errorf("method already registered: %s.%s: %s", m.this, name, m.fnc))
	}
	m := vm.newObjMethod(name, fnc)
	vm.methods[name] = m
}

func (vm *api) setIndexFunction(meta *lua.LTable, fnc func(val interface{}, key string) (lua.LValue, bool)) {
	// Node[key]
	meta.RawSetString("__index", vm.s.NewFunction(func(s *lua.LState) int {
		val := s.CheckUserData(1).Value
		key := s.CheckString(2)
		// automatically handle accesses for well-known interfaces
		if v, ok := vm.indexInterfaceV0(val, key); ok {
			s.Push(v)
			return 1
		}
		if fnc != nil {
			// custom fields for specific types
			if v, ok := fnc(val, key); ok {
				if v == nil {
					s.Push(lua.LNil)
				} else {
					s.Push(v)
				}
				return 1
			}
		}
		// check interface methods for that key
		if m, ok := vm.methods[key]; ok {
			s.Push(m.lua)
			return 1
		}
		// fallback to the metatable (e.g. custom class methods)
		s.Push(s.RawGet(meta, lua.LString(key)))
		return 1
	}))
}

func (vm *api) setSetIndexFunction(meta *lua.LTable, fnc func(obj interface{}, key string, val lua.LValue) bool) {
	meta.RawSetString("__newindex", vm.s.NewFunction(func(s *lua.LState) int {
		obj := s.CheckUserData(1).Value
		key := s.CheckString(2)
		if vm.setindexInterfaceV0(s, obj, key) {
			return 0
		}
		if fnc != nil {
			if fnc(obj, key, s.Get(3)) {
				return 0
			}
		}
		s.ArgError(2, fmt.Sprintf("no such key: %q", key))
		return 0
	}))
}

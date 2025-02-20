package server

import (
	"fmt"
	"math"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/types"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/unit/ai"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

const (
	defaultFirstObjectScriptID = ObjectScriptID(1000000000)
)

var (
	_ Obj                 = &Object{}
	_ script.Identifiable = &Object{}
)

type Obj interface {
	SObj() *Object
}

func toObject(p Obj) *Object {
	if p == nil {
		return nil
	}
	return p.SObj()
}

func toObjectC(p Obj) unsafe.Pointer {
	if p == nil {
		return nil
	}
	return p.SObj().CObj()
}

func asObjectP(p unsafe.Pointer) *Object {
	return (*Object)(p)
}

func (s *Server) CanSee(obj, targ *Object, flags int) bool {
	if obj.HasEnchant(ENCHANT_BLINDED) {
		return false
	}
	if flags&0x1 == 0 && targ.HasEnchant(ENCHANT_INVISIBLE) {
		if noxflags.HasGame(noxflags.GameModeQuest) {
			switch int(obj.TypeInd) {
			case s.Types.HecubahID(), s.Types.NecromancerID():
				goto check
			}
		} else if obj.HasEnchant(ENCHANT_INFRAVISION) {
			goto check
		}
		if int(obj.TypeInd) != s.Types.PixieID() || obj.ObjOwner == nil || !obj.ObjOwner.HasEnchant(ENCHANT_INFRAVISION) {
			if math.Abs(float64(targ.VelVec.X)) <= 6 && math.Abs(float64(targ.VelVec.Y)) <= 6 {
				return false
			}
		}
	}
check:
	ok := true
	if obj.Class().Has(object.ClassMonster) && obj.SubClass().AsMonster().Has(object.MonsterBomber) {
		if s.Abils.IsActive(targ, AbilityTreadLightly) {
			ok = false
		}
	}
	return ok
}

func (s *Server) CanInteract(obj, targ *Object, flags int) bool {
	ok := s.CanSee(obj, targ, flags)
	if ok {
		ok = s.MapTraceVision(obj, targ)
	}
	return ok
}

func (s *Server) CanSeeDir(obj, targ *Object) bool {
	if !s.CanInteract(obj, targ, 0) {
		return false
	}
	dp := targ.Pos().Sub(obj.Pos())
	dx, dy := float64(dp.X), float64(dp.Y)
	dir := obj.Direction1.Vec()
	dist := math.Sqrt(dx*dx+dy*dy) + 0.001
	return dy/dist*float64(dir.Y)+dx/dist*float64(dir.X) > 0.5
}

type serverObjects struct {
	handle          uintptr
	alloc           alloc.ClassT[Object]
	Alive           int
	MaxAlive        int
	Created         int
	CreatedSimple   int
	CreatedImmobile int
	List            *Object
	Pending         *Object
	UpdatableList   *Object
	UpdatableList2  *Object
	DeletedList     *Object
	firstScriptID   ObjectScriptID
	lastScriptID    ObjectScriptID
	PendingActions  []func()

	XFerInvLight ObjectXferFunc
}

func (s *serverObjects) init(h uintptr) {
	s.handle = h
}

func (s *serverObjects) LastObjectScriptID() ObjectScriptID {
	return s.lastScriptID
}

func (s *serverObjects) NextObjectScriptID() ObjectScriptID {
	id := s.lastScriptID
	s.lastScriptID++
	return id
}

func (s *serverObjects) SetLastObjectScriptID(id ObjectScriptID) {
	s.lastScriptID = id
}

func (s *serverObjects) SetFirstObjectScriptID(id ObjectScriptID) {
	s.firstScriptID = id
}

func (s *serverObjects) ResetObjectScriptIDs() {
	s.SetLastObjectScriptID(defaultFirstObjectScriptID)
	if s.firstScriptID != 0 {
		s.SetLastObjectScriptID(s.firstScriptID)
	}
}

func (s *serverObjects) GetAndZeroObjects() *Object { // nox_get_and_zero_server_objects_4DA3C0
	l := s.List
	s.List = nil
	return l
}

func (s *serverObjects) SetObjects(list *Object) { // nox_set_server_objects_4DA3E0
	s.List = list
}

func (s *serverObjects) Init(cnt int) bool {
	s.alloc = alloc.NewClassT("objectMemClass", Object{}, cnt)
	s.alloc.Keep(36) // NetCode
	if cnt == 0 {
		s.alloc.ResetStats()
		return true
	}
	var last *Object
	for i := 1; i <= cnt; i++ {
		obj := s.alloc.NewObject()
		if obj == nil {
			var next *Object
			for it := last; it != nil; it = next {
				next = it.ObjNext
				s.alloc.FreeObjectFirst(it)
			}
			return false
		}
		obj.NetCode = uint32(i)
		obj.ObjNext = last
		last = obj
	}
	var next *Object
	for it := last; it != nil; it = next {
		next = it.ObjNext
		code := it.NetCode
		s.alloc.FreeObjectFirst(it)
		it.NetCode = code
	}
	s.alloc.ResetStats()
	return true
}

func (s *serverObjects) FreeObjects() {
	s.alloc.Free()
}

func (s *serverObjects) FreeObject(obj *Object) int {
	if obj.Class().Has(object.ClassMonsterGenerator) {
		ud := obj.UpdateData
		arr := unsafe.Slice((**Object)(ud), 12)
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				it := arr[i*4+j]
				if it != nil {
					s.FreeObject(it)
				}
			}
		}
	}
	if !noxflags.HasGame(noxflags.GameFlag22) {
		var next *Object
		for it := obj.InvFirstItem; it != nil; it = next {
			next = it.InvNextItem
			s.FreeObject(it)
		}
	}
	if obj.IDPtr != nil {
		obj.IDPtr = nil
	}
	if obj.HealthData != nil {
		obj.HealthData = nil
	}
	if obj.ScriptVars != nil {
		obj.ScriptVars = nil
	}
	if obj.Field189 != nil {
		obj.Field189 = nil
	}
	if obj.InitData != nil {
		obj.InitData = nil
	}
	if obj.CollideData != nil {
		obj.CollideData = nil
	}
	if obj.UseData != nil {
		obj.UseData = nil
	}
	if obj.UpdateData != nil {
		obj.UpdateData = nil
	}
	code := obj.NetCode
	s.alloc.FreeObjectLast(obj.SObj())
	obj.NetCode = code
	s.Alive--
	return s.Alive
}

func (s *serverObjects) NewObject(t *ObjectType) *Object {
	obj := s.alloc.NewObject()
	*obj = Object{
		NetCode:      obj.NetCode,      // it is persisted by the allocator; so we basically reuse ID of the older object
		TypeInd:      uint16(t.Ind2()), // TODO: why is it setting it and then overwriting again?
		ObjClass:     uint32(t.Class()),
		ObjSubClass:  uint32(t.SubClass()),
		ObjFlags:     uint32(t.Flags()),
		Field5:       t.Field9,
		Material:     uint16(t.Material()),
		Experience:   t.Experience,
		Worth:        uint32(t.Worth),
		Float28:      t.Field13,
		Mass:         t.Mass,
		ZSize1:       t.ZSize1,
		ZSize2:       t.ZSize2,
		serverHandle: s.handle,
	}
	obj.Shape = t.Shape
	if !obj.Flags().Has(object.FlagNoCollide) {
		obj.SObj().UpdateCollider(obj.PosVec)
	}
	obj.Weight = t.Weight
	obj.CarryCapacity = uint16(t.CarryCap)
	obj.SpeedCur = t.Speed
	obj.SpeedBase = t.SpeedBase
	obj.SpeedBonus = t.SpeedBonus
	obj.HealthData = nil
	obj.Field38 = math.MaxUint32
	obj.TypeInd = uint16(t.Ind())
	if t.Health() != nil {
		data, _ := alloc.New(HealthData{})
		obj.HealthData = data
		*data = *t.Health()
	}
	obj.Init = t.Init
	if t.InitDataSize != 0 {
		data, _ := alloc.Make([]byte{}, t.InitDataSize)
		obj.InitData = unsafe.Pointer(&data[0])
		copy(data, unsafe.Slice((*byte)(t.InitData), t.InitDataSize))
	}
	obj.Collide = t.Collide
	if t.CollideDataSize != 0 {
		data, _ := alloc.Make([]byte{}, t.CollideDataSize)
		obj.CollideData = unsafe.Pointer(&data[0])
		copy(data, unsafe.Slice((*byte)(t.CollideData), t.CollideDataSize))
	}
	obj.Xfer = t.Xfer
	obj.Use = t.Use
	if t.UseDataSize != 0 {
		data, _ := alloc.Make([]byte{}, t.UseDataSize)
		obj.UseData = unsafe.Pointer(&data[0])
		copy(data, unsafe.Slice((*byte)(t.UseData), t.UseDataSize))
	}
	obj.Update = t.Update
	if t.UpdateDataSize != 0 {
		data, _ := alloc.Make([]byte{}, t.UpdateDataSize)
		obj.UpdateData = unsafe.Pointer(&data[0])
		copy(data, unsafe.Slice((*byte)(t.UpdateData), t.UpdateDataSize))
	}
	obj.Pickup = t.Pickup
	obj.Drop = t.Drop
	obj.Damage = t.Damage
	obj.DamageSound = t.DamageSound
	obj.Death = t.Death
	obj.ScriptVars = nil
	obj.DeathData = t.DeathData
	obj.Field192 = -1
	if noxflags.HasGame(noxflags.GameFlag22|noxflags.GameFlag23) && (obj.Class().HasAny(0x20A02) || obj.Xfer.Equals(s.XFerInvLight) || obj.Weight != 0xff) {
		obj.Field189, _ = alloc.Malloc(2572)
	}
	if fnc := t.Create.Get(); fnc != nil {
		fnc(obj)
	}
	if !noxflags.HasGame(noxflags.GameFlag22) {
		obj.ScriptIDVal = int(s.NextObjectScriptID())
	}
	if obj.Class().Has(object.ClassSimple) {
		s.CreatedSimple++
	} else if obj.Class().Has(object.ClassImmobile) {
		s.CreatedImmobile++
	}
	v8 := s.Alive + 1
	s.Created++
	s.Alive++
	if s.Alive > s.MaxAlive {
		s.MaxAlive = v8
	}
	return obj
}

func (s *Server) NewObjectByTypeInd(ind int) *Object { // nox_xxx_newObjectWithTypeInd_4E3450
	typ := s.Types.ByInd(ind)
	if typ == nil {
		return nil
	}
	return s.Objs.NewObject(typ)
}

func (s *Server) NewObjectByTypeID(id string) *Object { // nox_xxx_newObjectByTypeID_4E3810
	typ := s.Types.ByID(id)
	if typ == nil {
		return nil
	}
	return s.Objs.NewObject(typ)
}

func (s *Server) FirstServerObject() *Object { // nox_server_getFirstObject_4DA790
	return s.Objs.List
}

func (s *Server) GetObjects() []*Object {
	var out []*Object
	for p := s.FirstServerObject(); p != nil; p = p.ObjNext {
		out = append(out, p)
	}
	return out
}

func (s *Server) GetObjectsUpdatable2() []*Object {
	var out []*Object
	for p := s.Objs.UpdatableList2; p != nil; p = p.ObjNext {
		out = append(out, p)
	}
	return out
}

func (s *serverObjects) AddToUpdatable(obj *Object) {
	if obj.IsUpdatable == 0 && !obj.Class().Has(object.ClassMissile) {
		obj.UpdatablePrev = nil
		obj.UpdatableNext = s.UpdatableList
		if s.UpdatableList != nil {
			s.UpdatableList.UpdatablePrev = obj
		}
		s.UpdatableList = obj
		obj.IsUpdatable = 1
		obj.Obj130 = nil
	}
}

func (s *serverObjects) RemoveFromUpdatable(obj *Object) {
	if obj.IsUpdatable == 0 {
		return
	}
	prev := obj.UpdatablePrev
	if prev != nil {
		prev.UpdatableNext = obj.UpdatableNext
	} else {
		s.UpdatableList = obj.UpdatableNext
	}
	if next := obj.UpdatableNext; next != nil {
		next.UpdatablePrev = prev
	}
	obj.IsUpdatable = 0
	obj.Obj130 = nil
}

func (s *Server) getPending() []*Object {
	var out []*Object
	for p := s.Objs.Pending; p != nil; p = p.ObjNext {
		out = append(out, p)
	}
	return out
}

func (s *Server) GetObjectByInd(ind int) *Object { // aka nox_xxx_netGetUnitByExtent_4ED020
	for p := s.FirstServerObject(); p != nil; p = p.ObjNext {
		if !p.Flags().Has(object.FlagDestroyed) && p.Ind() == ind {
			return p
		}
	}
	return nil
}

func (s *Server) GetObjectByID(id string) *Object {
	for obj := s.Objs.List; obj != nil; obj = obj.ObjNext {
		if p := obj.FindByID(id); p != nil {
			return p
		}
	}
	for obj := s.Objs.Pending; obj != nil; obj = obj.ObjNext {
		if p := obj.FindByID(id); p != nil {
			return p
		}
	}
	return nil
}

func (s *serverObjects) QueueAction(fnc func()) {
	s.PendingActions = append(s.PendingActions, fnc)
}

func (s *Server) ObjectsClearPending() {
	var next *Object
	for it := s.Objs.Pending; it != nil; it = next {
		next = it.ObjNext
		it.ObjFlags &^= uint32(object.FlagPending)
		if s.Objs.List != nil {
			s.Objs.List.ObjPrev = it
		}
		it.ObjNext = s.Objs.List
		it.ObjPrev = nil
		s.Objs.List = it
	}
	s.Objs.Pending = nil
}

type Dir16 uint16

func (v Dir16) Vec() types.Pointf {
	x, y := SinCosDir(byte(v))
	return types.Pointf{X: x, Y: y}
}

type ObjectIndex struct {
	X     uint16
	Y     uint16
	Next4 *ObjectIndex
	Prev8 *ObjectIndex
	Obj12 *Object
}

type ScriptCallback struct {
	Flags uint32
	Func  int32
}

type Object struct {
	IDPtr         *byte                             // 0, 0
	TypeInd       uint16                            // 1, 4
	Field1_2      uint16                            // 1, 6
	ObjClass      uint32                            // 2, 8
	ObjSubClass   uint32                            // 3, 12
	ObjFlags      uint32                            // 4, 16
	Field5        uint32                            // 5, 20
	Material      uint16                            // 6, 24
	Field6_2      uint16                            // 6, 26
	Experience    float32                           // 7, 28
	Worth         uint32                            // 8, 32
	NetCode       uint32                            // 9, 36
	Extent        uint32                            // 10, 40
	ScriptIDVal   int                               // 11, 44
	TeamVal       ObjectTeam                        // 12, 48
	PosVec        types.Pointf                      // 14, 56
	NewPos        types.Pointf                      // 16, 64
	PrevPos       types.Pointf                      // 18, 72
	VelVec        types.Pointf                      // 20, 80
	ForceVec      types.Pointf                      // 22, 88
	Pos24         types.Pointf                      // 24, 96, // TODO: something related to acceleration/direction
	ZVal          float32                           // 26, 104
	Field27       float32                           // 27, 108
	Float28       float32                           // 28, 112, // TODO: damping/drag?
	Field29       uint32                            // 29, 116
	Mass          float32                           // 30, 120
	Direction1    Dir16                             // 31, 124
	Direction2    Dir16                             // 31, 126
	Field32       uint32                            // 32, 128, TODO: some frame/timestamp
	Field33       uint32                            // 33, 132
	Field34       uint32                            // 34, 136, TODO: some frame/timestamp
	Field35       uint32                            // 35, 140
	Field36       uint32                            // 36, 144
	Field37       uint32                            // 37, 148
	Field38       uint32                            // 38, 152
	Pos39         types.Pointf                      // 39, 156
	Field41       uint32                            // 41, 164
	Field42       uint32                            // 42, 168
	Shape         Shape                             // 43, 172
	ZSize1        float32                           // 56, 224
	ZSize2        float32                           // 57, 228
	CollideP1     types.Pointf                      // 58, 232
	CollideP2     types.Pointf                      // 60, 240
	Field62       [getInRectStackSize]uint32        // 62, 248
	ObjIndexBase  ObjectIndex                       // 64, 256
	ObjIndex      [4]ObjectIndex                    // 68, 272
	ObjIndexCur   uint32                            // 84, 336
	Buffs         uint32                            // 85, 340
	BuffsDur      [32]uint16                        // 86, 344
	BuffsPower    [32]uint8                         // 102, 408
	Field110      uint32                            // 110, 440
	ObjNext       *Object                           // 111, 444
	ObjPrev       *Object                           // 112, 448
	DeletedNext   *Object                           // 113, 452
	DeletedAt     uint32                            // 114, 456
	Field115      uint32                            // 115, 460
	Field116      uint32                            // 116, 464
	Field117      *Object                           // 117, 468
	Field118      *Object                           // 118, 472
	UpdatableNext *Object                           // 119, 476
	UpdatablePrev *Object                           // 120, 480
	IsUpdatable   uint32                            // 121, 484
	Weight        uint8                             // 122, 488
	Field122_1    uint8                             // 122, 489
	CarryCapacity uint16                            // 122, 490
	InvHolder     *Object                           // 123, 492 // Also health data, possibly same as 556, see 4E4560
	InvNextItem   *Object                           // 124, 496
	Field125      *Object                           // 125, 500, TODO: an Object*? see 4ED0C0
	InvFirstItem  *Object                           // 126, 504
	ObjOwner      *Object                           // 127, 508
	Field128      *Object                           // 128, 512
	Field129      *Object                           // 129, 516
	Obj130        *Object                           // 130, 520
	Field131      uint32                            // 131, 524
	Pos132        types.Pointf                      // 132, 528
	Frame134      uint32                            // 134, 536, TODO: some timestamp
	Field540      byte                              // 135, 540
	Field541      byte                              // 135, 541
	Field542      uint16                            // 135, 542
	SpeedCur      float32                           // 136, 544
	SpeedBase     float32                           // 137, 548
	SpeedBonus    float32                           // 138, 552
	HealthData    *HealthData                       // 139, 556
	Field140      [32]uint32                        // 140, 560
	Init          ccall.Func[ObjectInitFunc]        // 172, 688
	InitData      unsafe.Pointer                    // 173, 692
	Collide       ccall.Func[ObjectCollideFunc]     // 174, 696
	CollideData   unsafe.Pointer                    // 175, 700
	Xfer          ccall.Func[ObjectXferFunc]        // 176, 704
	Pickup        ccall.Func[ObjectPickupFunc]      // 177, 708
	Drop          ccall.Func[ObjectDropFunc]        // 178, 712
	Damage        ccall.Func[ObjectDamageFunc]      // 179, 716
	DamageSound   ccall.Func[ObjectDamageSoundFunc] // 180, 720
	Death         ccall.Func[ObjectDeathFunc]       // 181, 724
	DeathData     unsafe.Pointer                    // 182, 728
	Use           ccall.Func[ObjectUseFunc]         // 183, 732
	UseData       unsafe.Pointer                    // 184, 736
	Field185      uint32                            // 185, 740
	Update        ccall.Func[ObjectUpdateFunc]      // 186, 744
	UpdateData    unsafe.Pointer                    // 187, 748
	Field188      uint32                            // 188, 752
	Field189      unsafe.Pointer                    // 189, 756
	ScriptVars    unsafe.Pointer                    // 190, 760; []uint32
	Field191      uint32                            // 191, 764
	Field192      int                               // 192, 768
	serverHandle  uintptr                           // EXT
}

func (obj *Object) CObj() unsafe.Pointer {
	return unsafe.Pointer(obj)
}

func (obj *Object) SObj() *Object {
	if obj == nil {
		return nil
	}
	return obj
}

func (obj *Object) Server() *Server {
	if obj == nil {
		return nil
	}
	return getServer(obj.serverHandle)
}

func (obj *Object) ScriptID() int {
	if obj == nil {
		return 0
	}
	return obj.ScriptIDVal
}

func (obj *Object) ObjScriptID() int {
	if obj == nil {
		return 0
	}
	return obj.ScriptIDVal
}

func (obj *Object) ScriptVarPtr(i int) *uint32 {
	if obj == nil || obj.ScriptVars == nil {
		return nil
	}
	return (*uint32)(unsafe.Add(obj.ScriptVars, 4*i))
}

func (obj *Object) Ind() int { // aka "extent"
	return int(obj.Extent)
}

func (obj *Object) ID() string {
	if obj == nil {
		return ""
	}
	return alloc.GoString(obj.IDPtr)
}

func (obj *Object) EqualID(id2 string) bool {
	id := obj.ID()
	if id == "" {
		return false
	}
	return id == id2 || strings.HasSuffix(id, ":"+id2)
}

func (obj *Object) FindByID(id string) *Object {
	if obj.EqualID(id) {
		return obj
	}
	for p := obj.InvFirstItem; p != nil; p = p.InvNextItem {
		if p.EqualID(id) {
			return p
		}
	}
	return nil
}

func (obj *Object) Next() *Object {
	if obj == nil {
		return nil
	}
	return obj.ObjNext
}

func (obj *Object) Prev() *Object {
	if obj == nil {
		return nil
	}
	return obj.ObjPrev
}

func (obj *Object) FirstItem() *Object {
	if obj == nil {
		return nil
	}
	return obj.InvFirstItem
}

func (obj *Object) NextItem() *Object {
	if obj == nil {
		return nil
	}
	return obj.InvNextItem
}

func (obj *Object) HasItem(item *Object) bool {
	if item == nil || item.InvHolder != obj {
		return false
	}
	for p := obj.InvFirstItem; p != nil; p = p.InvNextItem {
		if p == item {
			return true
		}
	}
	return false
}

func (obj *Object) NextOwned512() *Object {
	return obj.Field128
}

func (obj *Object) FirstOwned516() *Object {
	return obj.Field129
}

func (obj *Object) stringAs(typ string) string {
	if obj == nil {
		return fmt.Sprintf("%s(<nil>)", typ)
	}
	var oid string
	if id := obj.ID(); id != "" {
		oid = fmt.Sprintf("%q", id)
	} else if ind := obj.Ind(); ind != 0 {
		oid = fmt.Sprintf("#%d", ind)
	}
	if t := obj.TypeInd; t != 0 {
		return fmt.Sprintf("%s(%s,T:%q)", typ, oid, t)
	}
	return fmt.Sprintf("%s(%s)", typ, oid)
}

func (obj *Object) String() string {
	return obj.stringAs("Object")
}

func (obj *Object) Pos() types.Pointf {
	if obj == nil {
		return types.Pointf{}
	}
	return obj.PosVec
}

func (obj *Object) Vel() types.Pointf {
	if obj == nil {
		return types.Pointf{}
	}
	return obj.VelVec
}

func (obj *Object) Force() types.Pointf {
	if obj == nil {
		return types.Pointf{}
	}
	return obj.ForceVec
}

func (obj *Object) Class() object.Class {
	if obj == nil {
		return 0
	}
	return object.Class(obj.ObjClass)
}

func (obj *Object) SubClass() object.SubClass {
	if obj == nil {
		return 0
	}
	return object.SubClass(obj.ObjSubClass)
}

func (obj *Object) ArmorClass() object.ArmorClass {
	if !obj.Class().Has(object.ClassArmor) {
		return 0
	}
	return object.ArmorClass(obj.SubClass())
}

func (obj *Object) Flags() object.Flags {
	if obj == nil {
		return 0
	}
	return object.Flags(obj.ObjFlags)
}

func (obj *Object) SetFlags(v object.Flags) {
	// TODO: do we need to update any cheksums?
	obj.ObjFlags = uint32(v)
}

func (obj *Object) FlagsEnable(v object.Flags) {
	obj.SetFlags(obj.Flags() | v)
}

func (obj *Object) FlagsDisable(v object.Flags) {
	obj.SetFlags(obj.Flags() &^ v)
}

func (obj *Object) IsEnabled() bool {
	if obj == nil {
		return false
	}
	return obj.Flags().Has(object.FlagEnabled)
}

func (obj *Object) IsMovable() bool {
	if obj.Flags().HasAny(object.FlagNotMovableMask) {
		return false
	}
	return !obj.Class().Has(object.ClassImmobile)
}

func (obj *Object) Dir1() Dir16 {
	if obj == nil {
		return 0
	}
	return obj.Direction1
}

func (obj *Object) Dir2() Dir16 {
	if obj == nil {
		return 0
	}
	return obj.Direction2
}

func (obj *Object) InitDataGlyph() *GlyphInitData {
	if alloc.IsDead(obj.UpdateData) {
		panic("object already deleted")
	}
	// TODO: verify this conversion by checking ObjectType
	return (*GlyphInitData)(obj.InitData)
}

func (obj *Object) InitDataShopkeeper() *ShopkeeperInitData {
	if alloc.IsDead(obj.UpdateData) {
		panic("object already deleted")
	}
	// TODO: verify this conversion by checking ObjectType
	return (*ShopkeeperInitData)(obj.InitData)
}

func (obj *Object) InitDataModifier() *ModifierInitData {
	if alloc.IsDead(obj.UpdateData) {
		panic("object already deleted")
	}
	// TODO: verify this conversion by checking ObjectType
	return (*ModifierInitData)(obj.InitData)
}

func (obj *Object) UpdateDataPlayer() *PlayerUpdateData {
	if !obj.Class().Has(object.ClassPlayer) {
		panic(obj.Class().String())
	}
	if alloc.IsDead(obj.UpdateData) {
		panic("object already deleted")
	}
	// TODO: verify this conversion by checking ObjectType
	return (*PlayerUpdateData)(obj.UpdateData)
}

func (obj *Object) ControllingPlayer() *Player {
	if obj == nil {
		return nil
	}
	ud := obj.UpdateDataPlayer()
	return ud.Player
}

func (obj *Object) UpdateDataMonster() *MonsterUpdateData {
	if !obj.Class().Has(object.ClassMonster) {
		panic(obj.Class().String())
	}
	if alloc.IsDead(obj.UpdateData) {
		panic("object already deleted")
	}
	// TODO: verify this conversion by checking ObjectType
	return (*MonsterUpdateData)(obj.UpdateData)
}

func (obj *Object) TeamPtr() *ObjectTeam {
	if obj == nil {
		return nil
	}
	return &obj.TeamVal
}

func (obj *Object) HasTeam() bool {
	if obj == nil {
		return false
	}
	return obj.TeamVal.Has()
}

func (obj *Object) Team() *Team {
	if obj == nil {
		return nil
	}
	if !obj.HasTeam() {
		return nil
	}
	return obj.Server().Teams.ByID(obj.TeamVal.ID)
}

func (obj *Object) Push(p types.Pointf, force float32) {
	vec := obj.Pos().Sub(p).Normalize()
	obj.ApplyForce(vec.Mul(force))
}

// ApplyForce adds a new force vector to the object. If another force in effect, it will adds up.
func (obj *Object) ApplyForce(p types.Pointf) {
	obj.ForceVec = obj.ForceVec.Add(p)
}

func (obj *Object) PushTo(p types.Pointf) {
	obj.ApplyForce(p)
}

func (obj *Object) SetDir(v Dir16) {
	obj.Direction1 = v
	obj.Direction2 = v
}

func (obj *Object) LookAt(p types.Pointf) {
	p = p.Sub(obj.Pos())
	v := DirFromVec(p)
	obj.SetDir(v)
}

func (obj *Object) Sub548600(dp types.Pointf) {
	obj.Pos24 = obj.Pos24.Add(dp.Div(obj.Mass))
}

func (obj *Object) Health() (cur, max int) {
	if obj == nil {
		return
	}
	h := obj.HealthData
	if h == nil || alloc.IsDead(unsafe.Pointer(h)) {
		return
	}
	cur = int(h.Cur)
	max = int(h.Max)
	return
}

func (obj *Object) Mana() (cur, max int) {
	if obj == nil {
		return
	}
	p := obj.UpdateDataPlayer()
	if p == nil {
		return
	}
	cur = int(p.ManaCur)
	max = int(p.ManaMax)
	return
}

func (obj *Object) CurrentSpeed() float32 {
	if obj == nil {
		return 0
	}
	return obj.SpeedCur
}

func (obj *Object) BaseSpeed() float32 {
	if obj == nil {
		return 0
	}
	return obj.SpeedBase
}

func (obj *Object) SetBaseSpeed(v float32) {
	// TODO: do we need to update any cheksums?
	obj.SpeedBase = v
}

func (obj *Object) Strength() int {
	if obj == nil {
		return 0
	}
	cl := obj.Class()
	switch {
	case cl.Has(object.ClassPlayer):
		ud := obj.UpdateDataPlayer()
		pl := ud.Player
		return int(*(*uint32)(unsafe.Add(pl.C(), 2239)))
	case cl.Has(object.ClassMonster):
		ud := obj.UpdateDataMonster()
		sub := obj.SubClass().AsMonster()
		switch {
		case sub.Has(object.MonsterNPC):
			return int(uint8(ud.Field331))
		default:
			return 30
		}
	default:
		return 0
	}
}

func (obj *Object) SetStrength(v int) {
	cl := obj.Class()
	switch {
	case cl.Has(object.ClassPlayer):
		ud := obj.UpdateDataPlayer()
		pl := ud.Player
		// TODO: do we need to update any cheksums?
		*(*uint32)(unsafe.Add(pl.C(), 2239)) = uint32(v)
	case cl.Has(object.ClassMonster):
		ud := obj.UpdateDataMonster()
		sub := obj.SubClass().AsMonster()
		switch {
		case sub.Has(object.MonsterNPC):
			// TODO: do we need to update any cheksums?
			ud.Field331 = uint32(uint8(v))
		}
	}
}

func (obj *Object) HasEnchant(v EnchantID) bool { // nox_xxx_testUnitBuffs_4FF350
	if obj == nil || v >= 32 {
		return false
	}
	return obj.Buffs&(uint32(1)<<v) != 0
}

func (obj *Object) EnchantDur(v EnchantID) int { // nox_xxx_unitGetBuffTimer_4FF550
	if obj == nil || v >= 32 {
		return 0
	}
	return int(obj.BuffsDur[v])
}

func (obj *Object) EnchantPower(v EnchantID) int { // nox_xxx_buffGetPower_4FF570
	if obj == nil || v >= 32 {
		return 0
	}
	return int(obj.BuffsDur[v])
}

func (obj *Object) Aggression() float32 {
	if obj.Class().Has(object.ClassMonster) {
		return 0
	}
	ud := obj.UpdateDataMonster()
	return ud.Aggression // TODO: which one?
}

func (obj *Object) SetAggression(v float32) {
	if obj.Class().Has(object.ClassMonster) {
		ud := obj.UpdateDataMonster()
		ud.SetAggression(v)
	}
}

func (obj *Object) RetreatLevel() float32 {
	if obj.Class().Has(object.ClassMonster) {
		return 0
	}
	ud := obj.UpdateDataMonster()
	return ud.RetreatLevel
}

func (obj *Object) SetRetreatLevel(v float32) {
	if obj.Class().Has(object.ClassMonster) {
		ud := obj.UpdateDataMonster()
		ud.RetreatLevel = v
	}
}

func (obj *Object) RegroupLevel() float32 {
	if obj.Class().Has(object.ClassMonster) {
		return 0
	}
	ud := obj.UpdateDataMonster()
	return ud.ResumeLevel
}

func (obj *Object) SetRegroupLevel(v float32) {
	if obj.Class().Has(object.ClassMonster) {
		ud := obj.UpdateDataMonster()
		ud.ResumeLevel = v
	}
}

func (obj *Object) HasOwner(owner *Object) bool {
	if obj == nil {
		return false
	}
	for it := obj; it != nil; it = it.ObjOwner {
		if it == owner {
			return true
		}
	}
	return false
}

func (obj *Object) FindOwnerChainPlayer() *Object {
	if obj == nil {
		return nil
	}
	res := obj
	for it := obj.ObjOwner; it != nil; it = it.ObjOwner {
		if it.Class().Has(object.ClassPlayer) {
			res = it
			break
		}
		res = it
	}
	return res
}

func (obj *Object) UpdateCollider(pos types.Pointf) {
	sh := &obj.Shape
	switch sh.Kind {
	case ShapeKindCenter:
		obj.CollideP1 = pos
		obj.CollideP2 = pos
	case ShapeKindCircle:
		r := types.Ptf(sh.Circle.R, sh.Circle.R)
		obj.CollideP1 = pos.Sub(r)
		obj.CollideP2 = pos.Add(r)
	case ShapeKindBox:
		obj.CollideP1 = pos.Add(types.Ptf(sh.Box.LeftBottom2, sh.Box.LeftBottom))
		obj.CollideP2 = pos.Add(types.Ptf(sh.Box.RightTop, sh.Box.RightTop2))
	}
}

func (obj *Object) Nox_xxx_objectUnkUpdateCoords_4E7290() {
	obj.UpdateCollider(obj.PosVec)
}

func (obj *Object) NeedSync() { // nox_xxx_unitNeedSync_4E44F0
	obj.Field38 = math.MaxUint32
}

func (obj *Object) SetDialogPortrait(name string) {
	if !obj.Class().Has(object.ClassMonster) {
		return
	}
	ud := obj.UpdateDataMonster()
	alloc.StrCopy(ud.DialogPortraitBuf[:], name)
}

func (obj *Object) CallUpdate() {
	if fnc := obj.Update.Get(); fnc != nil {
		fnc(obj)
	}
}

func (obj *Object) CallCollide(obj2 *Object, pos *types.Pointf) {
	if fnc := obj.Collide.Get(); fnc != nil {
		fnc(obj, obj2, pos)
	}
}

func (obj *Object) CallDamage(who Obj, a3 Obj, dmg int, typ object.DamageType) int {
	if fnc := obj.Damage.Get(); fnc != nil {
		return fnc(obj, toObject(who), toObject(a3), dmg, typ)
	}
	return 0
}

func (obj *Object) CallDrop(it Obj, pos types.Pointf) int {
	if fnc := obj.Drop.Get(); fnc != nil {
		cpos, free := alloc.New(types.Pointf{})
		defer free()
		*cpos = pos

		return fnc(obj, toObject(it), cpos)
	}
	return 0
}

func (obj *Object) CallXfer(a2 unsafe.Pointer) error {
	if obj.Xfer.Get()(obj, a2) == 0 {
		return fmt.Errorf("xfer for %s failed", obj.String())
	}
	return nil
}

func (s *Server) IsFish(obj *Object) bool {
	if obj == nil {
		return false
	}
	return int(obj.TypeInd) == s.Types.FishSmallID() || int(obj.TypeInd) == s.Types.FishBigID()
}

func (s *Server) IsRat(obj *Object) bool {
	if obj == nil {
		return false
	}
	return int(obj.TypeInd) == s.Types.RatID()
}

func (s *Server) IsFrog(obj *Object) bool {
	if obj == nil {
		return false
	}
	return int(obj.TypeInd) == s.Types.GreenFrogID()
}

func (s *Server) IsPlant(obj *Object) bool {
	if obj == nil {
		return false
	}
	return int(obj.TypeInd) == s.Types.CarnivorousPlantID()
}

func (s *Server) IsMimic(obj *Object) bool {
	if obj == nil {
		return false
	}
	return int(obj.TypeInd) == s.Types.MimicID()
}

func (obj *Object) IsEnemyTo(obj2 *Object) bool { // nox_xxx_unitIsEnemyTo_5330C0
	if obj == nil || obj2 == nil {
		return false
	}
	if obj.SObj() == obj2.SObj() {
		return false
	}
	return obj.Server().IsEnemyTo(obj, obj2)
}

func (s *Server) IsEnemyTo(obj, obj2 *Object) bool {
	if obj == nil || obj2 == nil {
		return false
	}
	if obj.SObj() == obj2.SObj() {
		return false
	}
	if obj2.Class().HasAny(object.ClassMonster) {
		if ud := obj2.UpdateDataMonster(); ud.StatusFlags.Has(object.MonStatusMorphed) {
			return false
		}
	}
	if obj.Class().HasAny(object.ClassPlayer) && int(obj2.TypeInd) == s.Types.PolypID() {
		return true
	}
	if obj.Class().HasAny(object.ClassPlayer) && obj2.Class().HasAny(object.ClassMonsterGenerator) {
		return true
	}
	if obj.Class().HasAny(object.ClassMonster) && obj2.Class().HasAny(object.ClassMonsterGenerator) {
		return false
	}
	if obj.Class().HasAny(object.ClassPlayer) && obj2.Class().HasAny(object.ClassMonster) && (obj2.SubClass()&0x20 != 0) {
		return false
	}
	if obj.Class().HasAny(object.ClassMonster) && obj2.Class().HasAny(object.ClassPlayer) && (obj.SubClass()&0x20 != 0) {
		return false
	}
	if s.IsFish(obj.SObj()) || s.IsFrog(obj.SObj()) {
		return !s.IsFish(obj2.SObj()) && !s.IsFrog(obj2.SObj())
	}
	if s.IsRat(obj.SObj()) {
		return !s.IsRat(obj2.SObj())
	}
	if s.IsFish(obj2.SObj()) {
		return false
	}
	if obj2.Class().HasAny(object.ClassMonster) && (obj2.SubClass()&0x8 != 0) {
		return false
	}
	if obj.Class().HasAny(object.ClassMonster) && (obj.SubClass()&0x8 != 0) {
		return false
	}
	if obj.HasEnchant(ENCHANT_CHARMING) || obj2.HasEnchant(ENCHANT_CHARMING) {
		return false
	}
	if obj2.Class().HasAny(object.ClassPlayer) {
		if pl := obj2.ControllingPlayer(); pl.Field3680&0x1 != 0 {
			return false
		}
	}
	own1 := obj.FindOwnerChainPlayer()
	own2 := obj2.FindOwnerChainPlayer()
	if own1 == own2 {
		return false
	}
	if own1.TeamPtr().SameAs(own2.TeamPtr()) {
		return false
	}
	if own1.Class().HasAny(object.ClassPlayer) && own2.Class().HasAny(object.ClassMonsterGenerator) {
		return true
	}
	if own1.Class().HasAny(object.ClassMonster) && own2.Class().HasAny(object.ClassMonsterGenerator) {
		return false
	}
	if !noxflags.HasGame(noxflags.GameModeQuest) && obj.Class().HasAny(object.ClassMonster) && int(obj2.TypeInd) == s.Types.WillOWispID() {
		return obj2.UpdateDataMonster().HasAction(ai.ACTION_FIGHT)
	}
	if own1.HasTeam() || own2.HasTeam() {
		return true
	}
	if own1.Class().HasAny(object.ClassMonster) && own2.Class().HasAny(object.ClassMonster) {
		return false
	}
	return true
}

func (s *Server) ItemsApplyUpdateEffect(obj *Object) {
	for it := obj.InvFirstItem; it != nil; it = it.InvNextItem {
		const maskItems = object.ClassFlag | object.ClassWeapon | object.ClassArmor | object.ClassWand
		if it.Flags().Has(object.FlagEquipped) && it.Class().HasAny(maskItems) {
			idata := unsafe.Slice((*unsafe.Pointer)(it.InitData), 4)
			for _, mod := range idata {
				if mod != nil {
					if fnc := *(*unsafe.Pointer)(unsafe.Add(mod, 100)); fnc != nil {
						ccall.AsFunc[func(unsafe.Pointer, *Object, unsafe.Pointer)](fnc)(mod, it, nil)
					}
				}
			}
		}
	}
}

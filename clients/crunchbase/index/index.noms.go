// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_index_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeEnumType("QuarterEnum", "Q1", "Q2", "Q3", "Q4"),
		types.MakeStructType("Quarter",
			[]types.Field{
				types.Field{"Year", types.MakePrimitiveType(types.Int32Kind), false},
				types.Field{"Quarter", types.MakeType(ref.Ref{}, 0), false},
			},
			types.Choices{},
		),
		types.MakeStructType("Key",
			[]types.Field{},
			types.Choices{
				types.Field{"Category", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Quarter", types.MakeType(ref.Ref{}, 1), false},
				types.Field{"Region", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"RoundType", types.MakeType(ref.Ref{}, 3), false},
				types.Field{"Year", types.MakePrimitiveType(types.Int32Kind), false},
			},
		),
		types.MakeEnumType("RoundTypeEnum", "Seed", "SeriesA", "SeriesB", "SeriesC", "SeriesD", "SeriesE", "SeriesF", "SeriesG", "SeriesH", "UnknownRoundType"),
	}, []ref.Ref{
		ref.Parse("sha1-d0c7c9389c0500f635bb58d1580ed988cc6324e9"),
		ref.Parse("sha1-ce23a43307c0d14735de5e3a4bfbe19e69c12282"),
	})
	__mainPackageInFile_index_CachedRef = types.RegisterPackage(&p)
}

// QuarterEnum

type QuarterEnum uint32

const (
	Q1 QuarterEnum = iota
	Q2
	Q3
	Q4
)

func NewQuarterEnum() QuarterEnum {
	return QuarterEnum(0)
}

var __typeForQuarterEnum types.Type

func (e QuarterEnum) Type() types.Type {
	return __typeForQuarterEnum
}

func init() {
	__typeForQuarterEnum = types.MakeType(__mainPackageInFile_index_CachedRef, 0)
	types.RegisterEnum(__typeForQuarterEnum, builderForQuarterEnum, readerForQuarterEnum)
}

func builderForQuarterEnum(v uint32) types.Value {
	return QuarterEnum(v)
}

func readerForQuarterEnum(v types.Value) uint32 {
	return uint32(v.(QuarterEnum))
}

func (e QuarterEnum) Equals(other types.Value) bool {
	return e == other
}

func (e QuarterEnum) Ref() ref.Ref {
	throwaway := ref.Ref{}
	return types.EnsureRef(&throwaway, e)
}

func (e QuarterEnum) Chunks() []ref.Ref {
	return nil
}

func (e QuarterEnum) ChildValues() []types.Value {
	return nil
}

// Quarter

type Quarter struct {
	_Year    int32
	_Quarter QuarterEnum

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewQuarter(cs chunks.ChunkStore) Quarter {
	return Quarter{
		_Year:    int32(0),
		_Quarter: NewQuarterEnum(),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type QuarterDef struct {
	Year    int32
	Quarter QuarterEnum
}

func (def QuarterDef) New(cs chunks.ChunkStore) Quarter {
	return Quarter{
		_Year:    def.Year,
		_Quarter: def.Quarter,
		cs:       cs,
		ref:      &ref.Ref{},
	}
}

func (s Quarter) Def() (d QuarterDef) {
	d.Year = s._Year
	d.Quarter = s._Quarter
	return
}

var __typeForQuarter types.Type

func (m Quarter) Type() types.Type {
	return __typeForQuarter
}

func init() {
	__typeForQuarter = types.MakeType(__mainPackageInFile_index_CachedRef, 1)
	types.RegisterStruct(__typeForQuarter, builderForQuarter, readerForQuarter)
}

func builderForQuarter(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := Quarter{ref: &ref.Ref{}, cs: cs}
	s._Year = int32(values[i].(types.Int32))
	i++
	s._Quarter = values[i].(QuarterEnum)
	i++
	return s
}

func readerForQuarter(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Quarter)
	values = append(values, types.Int32(s._Year))
	values = append(values, s._Quarter)
	return values
}

func (s Quarter) Equals(other types.Value) bool {
	return other != nil && __typeForQuarter.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Quarter) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Quarter) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForQuarter.Chunks()...)
	return
}

func (s Quarter) ChildValues() (ret []types.Value) {
	ret = append(ret, types.Int32(s._Year))
	ret = append(ret, s._Quarter)
	return
}

func (s Quarter) Year() int32 {
	return s._Year
}

func (s Quarter) SetYear(val int32) Quarter {
	s._Year = val
	s.ref = &ref.Ref{}
	return s
}

func (s Quarter) Quarter() QuarterEnum {
	return s._Quarter
}

func (s Quarter) SetQuarter(val QuarterEnum) Quarter {
	s._Quarter = val
	s.ref = &ref.Ref{}
	return s
}

// Key

type Key struct {
	__unionIndex uint32
	__unionValue types.Value
	cs           chunks.ChunkStore
	ref          *ref.Ref
}

func NewKey(cs chunks.ChunkStore) Key {
	return Key{
		__unionIndex: 0,
		__unionValue: types.NewString(""),
		cs:           cs,
		ref:          &ref.Ref{},
	}
}

type KeyDef struct {
	__unionIndex uint32
	__unionValue types.Value
}

func (def KeyDef) New(cs chunks.ChunkStore) Key {
	return Key{
		__unionIndex: def.__unionIndex,
		__unionValue: def.__unionValue,
		cs:           cs,
		ref:          &ref.Ref{},
	}
}

func (s Key) Def() (d KeyDef) {
	d.__unionIndex = s.__unionIndex
	d.__unionValue = s.__unionValue
	return
}

var __typeForKey types.Type

func (m Key) Type() types.Type {
	return __typeForKey
}

func init() {
	__typeForKey = types.MakeType(__mainPackageInFile_index_CachedRef, 2)
	types.RegisterStruct(__typeForKey, builderForKey, readerForKey)
}

func builderForKey(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := Key{ref: &ref.Ref{}, cs: cs}
	s.__unionIndex = uint32(values[i].(types.Uint32))
	i++
	s.__unionValue = values[i]
	i++
	return s
}

func readerForKey(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Key)
	values = append(values, types.Uint32(s.__unionIndex))
	values = append(values, s.__unionValue)
	return values
}

func (s Key) Equals(other types.Value) bool {
	return other != nil && __typeForKey.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Key) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Key) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForKey.Chunks()...)
	chunks = append(chunks, s.__unionValue.Chunks()...)
	return
}

func (s Key) ChildValues() (ret []types.Value) {
	ret = append(ret, s.__unionValue)
	return
}

func (s Key) Category() (val string, ok bool) {
	if s.__unionIndex != 0 {
		return
	}
	return s.__unionValue.(types.String).String(), true
}

func (s Key) SetCategory(val string) Key {
	s.__unionIndex = 0
	s.__unionValue = types.NewString(val)
	s.ref = &ref.Ref{}
	return s
}

func (def KeyDef) Category() (val string, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return def.__unionValue.(types.String).String(), true
}

func (def KeyDef) SetCategory(cs chunks.ChunkStore, val string) KeyDef {
	def.__unionIndex = 0
	def.__unionValue = types.NewString(val)
	return def
}

func (s Key) Quarter() (val Quarter, ok bool) {
	if s.__unionIndex != 1 {
		return
	}
	return s.__unionValue.(Quarter), true
}

func (s Key) SetQuarter(val Quarter) Key {
	s.__unionIndex = 1
	s.__unionValue = val
	s.ref = &ref.Ref{}
	return s
}

func (def KeyDef) Quarter() (val QuarterDef, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(Quarter).Def(), true
}

func (def KeyDef) SetQuarter(cs chunks.ChunkStore, val QuarterDef) KeyDef {
	def.__unionIndex = 1
	def.__unionValue = val.New(cs)
	return def
}

func (s Key) Region() (val string, ok bool) {
	if s.__unionIndex != 2 {
		return
	}
	return s.__unionValue.(types.String).String(), true
}

func (s Key) SetRegion(val string) Key {
	s.__unionIndex = 2
	s.__unionValue = types.NewString(val)
	s.ref = &ref.Ref{}
	return s
}

func (def KeyDef) Region() (val string, ok bool) {
	if def.__unionIndex != 2 {
		return
	}
	return def.__unionValue.(types.String).String(), true
}

func (def KeyDef) SetRegion(cs chunks.ChunkStore, val string) KeyDef {
	def.__unionIndex = 2
	def.__unionValue = types.NewString(val)
	return def
}

func (s Key) RoundType() (val RoundTypeEnum, ok bool) {
	if s.__unionIndex != 3 {
		return
	}
	return s.__unionValue.(RoundTypeEnum), true
}

func (s Key) SetRoundType(val RoundTypeEnum) Key {
	s.__unionIndex = 3
	s.__unionValue = val
	s.ref = &ref.Ref{}
	return s
}

func (def KeyDef) RoundType() (val RoundTypeEnum, ok bool) {
	if def.__unionIndex != 3 {
		return
	}
	return def.__unionValue.(RoundTypeEnum), true
}

func (def KeyDef) SetRoundType(cs chunks.ChunkStore, val RoundTypeEnum) KeyDef {
	def.__unionIndex = 3
	def.__unionValue = val
	return def
}

func (s Key) Year() (val int32, ok bool) {
	if s.__unionIndex != 4 {
		return
	}
	return int32(s.__unionValue.(types.Int32)), true
}

func (s Key) SetYear(val int32) Key {
	s.__unionIndex = 4
	s.__unionValue = types.Int32(val)
	s.ref = &ref.Ref{}
	return s
}

func (def KeyDef) Year() (val int32, ok bool) {
	if def.__unionIndex != 4 {
		return
	}
	return int32(def.__unionValue.(types.Int32)), true
}

func (def KeyDef) SetYear(cs chunks.ChunkStore, val int32) KeyDef {
	def.__unionIndex = 4
	def.__unionValue = types.Int32(val)
	return def
}

// RoundTypeEnum

type RoundTypeEnum uint32

const (
	Seed RoundTypeEnum = iota
	SeriesA
	SeriesB
	SeriesC
	SeriesD
	SeriesE
	SeriesF
	SeriesG
	SeriesH
	UnknownRoundType
)

func NewRoundTypeEnum() RoundTypeEnum {
	return RoundTypeEnum(0)
}

var __typeForRoundTypeEnum types.Type

func (e RoundTypeEnum) Type() types.Type {
	return __typeForRoundTypeEnum
}

func init() {
	__typeForRoundTypeEnum = types.MakeType(__mainPackageInFile_index_CachedRef, 3)
	types.RegisterEnum(__typeForRoundTypeEnum, builderForRoundTypeEnum, readerForRoundTypeEnum)
}

func builderForRoundTypeEnum(v uint32) types.Value {
	return RoundTypeEnum(v)
}

func readerForRoundTypeEnum(v types.Value) uint32 {
	return uint32(v.(RoundTypeEnum))
}

func (e RoundTypeEnum) Equals(other types.Value) bool {
	return e == other
}

func (e RoundTypeEnum) Ref() ref.Ref {
	throwaway := ref.Ref{}
	return types.EnsureRef(&throwaway, e)
}

func (e RoundTypeEnum) Chunks() []ref.Ref {
	return nil
}

func (e RoundTypeEnum) ChildValues() []types.Value {
	return nil
}

// MapOfRefOfKeyToSetOfRefOfRound

type MapOfRefOfKeyToSetOfRefOfRound struct {
	m   types.Map
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewMapOfRefOfKeyToSetOfRefOfRound(cs chunks.ChunkStore) MapOfRefOfKeyToSetOfRefOfRound {
	return MapOfRefOfKeyToSetOfRefOfRound{types.NewTypedMap(cs, __typeForMapOfRefOfKeyToSetOfRefOfRound), cs, &ref.Ref{}}
}

type MapOfRefOfKeyToSetOfRefOfRoundDef map[ref.Ref]SetOfRefOfRoundDef

func (def MapOfRefOfKeyToSetOfRefOfRoundDef) New(cs chunks.ChunkStore) MapOfRefOfKeyToSetOfRefOfRound {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, NewRefOfKey(k), v.New(cs))
	}
	return MapOfRefOfKeyToSetOfRefOfRound{types.NewTypedMap(cs, __typeForMapOfRefOfKeyToSetOfRefOfRound, kv...), cs, &ref.Ref{}}
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Def() MapOfRefOfKeyToSetOfRefOfRoundDef {
	def := make(map[ref.Ref]SetOfRefOfRoundDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(RefOfKey).TargetRef()] = v.(SetOfRefOfRound).Def()
		return false
	})
	return def
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Equals(other types.Value) bool {
	return other != nil && __typeForMapOfRefOfKeyToSetOfRefOfRound.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfRefOfKeyToSetOfRefOfRound) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfRefOfKeyToSetOfRefOfRound.
var __typeForMapOfRefOfKeyToSetOfRefOfRound types.Type

func (m MapOfRefOfKeyToSetOfRefOfRound) Type() types.Type {
	return __typeForMapOfRefOfKeyToSetOfRefOfRound
}

func init() {
	__typeForMapOfRefOfKeyToSetOfRefOfRound = types.MakeCompoundType(types.MapKind, types.MakeCompoundType(types.RefKind, types.MakeType(__mainPackageInFile_index_CachedRef, 2)), types.MakeCompoundType(types.SetKind, types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-ce23a43307c0d14735de5e3a4bfbe19e69c12282"), 1))))
	types.RegisterValue(__typeForMapOfRefOfKeyToSetOfRefOfRound, builderForMapOfRefOfKeyToSetOfRefOfRound, readerForMapOfRefOfKeyToSetOfRefOfRound)
}

func builderForMapOfRefOfKeyToSetOfRefOfRound(cs chunks.ChunkStore, v types.Value) types.Value {
	return MapOfRefOfKeyToSetOfRefOfRound{v.(types.Map), cs, &ref.Ref{}}
}

func readerForMapOfRefOfKeyToSetOfRefOfRound(v types.Value) types.Value {
	return v.(MapOfRefOfKeyToSetOfRefOfRound).m
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Empty() bool {
	return m.m.Empty()
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Len() uint64 {
	return m.m.Len()
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Has(p RefOfKey) bool {
	return m.m.Has(p)
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Get(p RefOfKey) SetOfRefOfRound {
	return m.m.Get(p).(SetOfRefOfRound)
}

func (m MapOfRefOfKeyToSetOfRefOfRound) MaybeGet(p RefOfKey) (SetOfRefOfRound, bool) {
	v, ok := m.m.MaybeGet(p)
	if !ok {
		return NewSetOfRefOfRound(m.cs), false
	}
	return v.(SetOfRefOfRound), ok
}

func (m MapOfRefOfKeyToSetOfRefOfRound) Set(k RefOfKey, v SetOfRefOfRound) MapOfRefOfKeyToSetOfRefOfRound {
	return MapOfRefOfKeyToSetOfRefOfRound{m.m.Set(k, v), m.cs, &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfRefOfKeyToSetOfRefOfRound) Remove(p RefOfKey) MapOfRefOfKeyToSetOfRefOfRound {
	return MapOfRefOfKeyToSetOfRefOfRound{m.m.Remove(p), m.cs, &ref.Ref{}}
}

type MapOfRefOfKeyToSetOfRefOfRoundIterCallback func(k RefOfKey, v SetOfRefOfRound) (stop bool)

func (m MapOfRefOfKeyToSetOfRefOfRound) Iter(cb MapOfRefOfKeyToSetOfRefOfRoundIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(RefOfKey), v.(SetOfRefOfRound))
	})
}

type MapOfRefOfKeyToSetOfRefOfRoundIterAllCallback func(k RefOfKey, v SetOfRefOfRound)

func (m MapOfRefOfKeyToSetOfRefOfRound) IterAll(cb MapOfRefOfKeyToSetOfRefOfRoundIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(RefOfKey), v.(SetOfRefOfRound))
	})
}

func (m MapOfRefOfKeyToSetOfRefOfRound) IterAllP(concurrency int, cb MapOfRefOfKeyToSetOfRefOfRoundIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(RefOfKey), v.(SetOfRefOfRound))
	})
}

type MapOfRefOfKeyToSetOfRefOfRoundFilterCallback func(k RefOfKey, v SetOfRefOfRound) (keep bool)

func (m MapOfRefOfKeyToSetOfRefOfRound) Filter(cb MapOfRefOfKeyToSetOfRefOfRoundFilterCallback) MapOfRefOfKeyToSetOfRefOfRound {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(RefOfKey), v.(SetOfRefOfRound))
	})
	return MapOfRefOfKeyToSetOfRefOfRound{out, m.cs, &ref.Ref{}}
}

// RefOfKey

type RefOfKey struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfKey(target ref.Ref) RefOfKey {
	return RefOfKey{target, &ref.Ref{}}
}

func (r RefOfKey) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfKey) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfKey) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfKey.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfKey) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfKey) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfKey.
var __typeForRefOfKey types.Type

func (r RefOfKey) Type() types.Type {
	return __typeForRefOfKey
}

func (r RefOfKey) Less(other types.OrderedValue) bool {
	return r.TargetRef().Less(other.(types.RefBase).TargetRef())
}

func init() {
	__typeForRefOfKey = types.MakeCompoundType(types.RefKind, types.MakeType(__mainPackageInFile_index_CachedRef, 2))
	types.RegisterRef(__typeForRefOfKey, builderForRefOfKey)
}

func builderForRefOfKey(r ref.Ref) types.Value {
	return NewRefOfKey(r)
}

func (r RefOfKey) TargetValue(cs chunks.ChunkStore) Key {
	return types.ReadValue(r.target, cs).(Key)
}

func (r RefOfKey) SetTargetValue(val Key, cs chunks.ChunkSink) RefOfKey {
	return NewRefOfKey(types.WriteValue(val, cs))
}

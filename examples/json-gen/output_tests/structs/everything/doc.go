// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type ByteAlias byte

//+k8s:json-gen=false
type BoolAlias bool

//+k8s:json-gen=false
type Int32Alias int32

//+k8s:json-gen=false
type Float32Alias float32

//+k8s:json-gen=false
type StringAlias string

//+k8s:json-gen=false
type SliceStringAlias []string

//+k8s:json-gen=false
type SlicePtrStringAlias []*string

//+k8s:json-gen=false
type MapStringStringAlias map[string]string

//+k8s:json-gen=false
type Inner struct {
	Byte            byte
	BytePtr         *byte
	ByteAlias       ByteAlias
	ByteAliasPtr    *ByteAlias
	Bool            bool
	BoolPtr         *bool
	BoolAlias       BoolAlias
	BoolAliasPtr    *BoolAlias
	Int8            int8
	Int8Ptr         *int8
	Int16           int16
	Int16Ptr        *int16
	Int32           int32
	Int32Ptr        *int32
	Int32Alias      Int32Alias
	Int32AliasPtr   *Int32Alias
	Uint8           uint8
	Uint8Ptr        *uint8
	Uint16          uint16
	Uint16Ptr       *uint16
	Uint32          uint32
	Uint32Ptr       *uint32
	Float32         float32
	Float32Ptr      *float32
	Float32Alias    Float32Alias
	Float32AliasPtr *Float32Alias
	Float64         float64
	Float64Ptr      *float64
	String          string
	StringPtr       *string
	StringAlias     StringAlias
	StringAliasPtr  *StringAlias
	Struct          struct {
		Byte                byte
		BytePtr             *byte
		ByteAlias           ByteAlias
		ByteAliasPtr        *ByteAlias
		Bool                bool
		BoolPtr             *bool
		BoolAlias           BoolAlias
		BoolAliasPtr        *BoolAlias
		Int8                int8
		Int8Ptr             *int8
		Int16               int16
		Int16Ptr            *int16
		Int32               int32
		Int32Ptr            *int32
		Int32Alias          Int32Alias
		Int32AliasPtr       *Int32Alias
		Uint8               uint8
		Uint8Ptr            *uint8
		Uint16              uint16
		Uint16Ptr           *uint16
		Uint32              uint32
		Uint32Ptr           *uint32
		Float32             float32
		Float32Ptr          *float32
		Float32Alias        Float32Alias
		Float32AliasPtr     *Float32Alias
		Float64             float64
		Float64Ptr          *float64
		String              string
		StringPtr           *string
		StringAlias         StringAlias
		StringAliasPtr      *StringAlias
		Struct              struct{}
		StructPtr           *Inner
		SliceString         []string
		SliceStringAlias    SliceStringAlias
		SlicePtrString      []*string
		SliceStringPtrAlias SlicePtrStringAlias
		SliceStringPtr      *[]string
		SliceByte           []byte
	}
	StructPtr *struct {
		Byte                    byte
		BytePtr                 *byte
		ByteAlias               ByteAlias
		ByteAliasPtr            *ByteAlias
		Bool                    bool
		BoolPtr                 *bool
		BoolAlias               BoolAlias
		BoolAliasPtr            *BoolAlias
		Int8                    int8
		Int8Ptr                 *int8
		Int16                   int16
		Int16Ptr                *int16
		Int32                   int32
		Int32Ptr                *int32
		Int32Alias              Int32Alias
		Int32AliasPtr           *Int32Alias
		Uint8                   uint8
		Uint8Ptr                *uint8
		Uint16                  uint16
		Uint16Ptr               *uint16
		Uint32                  uint32
		Uint32Ptr               *uint32
		Float32                 float32
		Float32Ptr              *float32
		Float32Alias            Float32Alias
		Float32AliasPtr         *Float32Alias
		Float64                 float64
		Float64Ptr              *float64
		String                  string
		StringPtr               *string
		StringAlias             StringAlias
		StringAliasPtr          *StringAlias
		Struct                  struct{}
		StructPtr               *Inner
		SliceString             []string
		SliceStringAlias        SliceStringAlias
		SlicePtrString          []*string
		SliceStringPtrAlias     SlicePtrStringAlias
		SliceStringPtr          *[]string
		SliceByte               []byte
		MapStringString         map[string]string
		MapStringStringPtr      *map[string]string
		MapStringStringAlias    MapStringStringAlias
		MapStringStringAliasPtr *MapStringStringAlias
	}
	SliceString             []string
	SliceStringAlias        SliceStringAlias
	SlicePtrString          []*string
	SliceStringPtrAlias     SlicePtrStringAlias
	SliceStringPtr          *[]string
	SliceByte               []byte
	MapStringString         map[string]string
	MapStringStringPtr      *map[string]string
	MapStringStringAlias    MapStringStringAlias
	MapStringStringAliasPtr *MapStringStringAlias
}

//+k8s:json-gen=false
type Tstd struct {
	Byte            byte
	BytePtr         *byte
	ByteAlias       ByteAlias
	ByteAliasPtr    *ByteAlias
	Bool            bool
	BoolPtr         *bool
	BoolAlias       BoolAlias
	BoolAliasPtr    *BoolAlias
	Int8            int8
	Int8Ptr         *int8
	Int16           int16
	Int16Ptr        *int16
	Int32           int32
	Int32Ptr        *int32
	Int32Alias      Int32Alias
	Int32AliasPtr   *Int32Alias
	Uint8           uint8
	Uint8Ptr        *uint8
	Uint16          uint16
	Uint16Ptr       *uint16
	Uint32          uint32
	Uint32Ptr       *uint32
	Float32         float32
	Float32Ptr      *float32
	Float32Alias    Float32Alias
	Float32AliasPtr *Float32Alias
	Float64         float64
	Float64Ptr      *float64
	String          string
	StringPtr       *string
	StringAlias     StringAlias
	StringAliasPtr  *StringAlias
	StructPtr       *Inner
	Struct          struct {
		Struct struct {
			Struct struct {
				Struct struct {
					String string
				}
			}
		}
	}
	SliceString             []string
	SliceStringAlias        SliceStringAlias
	SlicePtrString          []*string
	SliceStringPtrAlias     SlicePtrStringAlias
	SliceStringPtr          *[]string
	MapStringString         map[string]string
	MapStringStringPtr      *map[string]string
	MapStringStringAlias    MapStringStringAlias
	MapStringStringAliasPtr *MapStringStringAlias
}

type Ttest Tstd

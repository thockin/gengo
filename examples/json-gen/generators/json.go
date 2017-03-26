/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generators

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"

	"github.com/golang/glog"
)

// CustomArgs is used tby the go2idl framework to pass args specific to this
// generator.
type CustomArgs struct {
	BoundingDirs []string // Only deal with types rooted under these dirs.
}

// This is the comment tag that carries parameters for deep-copy generation.
const tagName = "k8s:json-gen"

// Known values for the comment tag.
const tagValuePackage = "package"

// tagValue holds parameters from a tagName tag.
type tagValue struct {
	value string
}

func extractTag(comments []string) *tagValue {
	tagVals := types.ExtractCommentTags("+", comments)[tagName]
	if tagVals == nil {
		// No match for the tag.
		return nil
	}
	// If there are multiple values, abort.
	if len(tagVals) > 1 {
		glog.Fatalf("Found %d %s tags: %q", len(tagVals), tagName, tagVals)
	}

	// If we got here we are returning something.
	tag := &tagValue{}

	// Get the primary value.
	parts := strings.Split(tagVals[0], ",")
	if len(parts) >= 1 {
		tag.value = parts[0]
	}

	// Parse extra arguments.
	parts = parts[1:]
	for i := range parts {
		kv := strings.SplitN(parts[i], "=", 2)
		k := kv[0]
		v := ""
		if len(kv) == 2 {
			v = kv[1]
		}
		switch k {
		default:
			_ = k
			_ = v
			glog.Fatalf("Unsupported %s param: %q", tagName, parts[i])
		}
	}
	return tag
}

func jsonNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Join: func(pre string, in []string, post string) string {
			return strings.Join(in, "_")
		},
		PrependPackageNames: 1,
	}
}

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public": jsonNamer(),
		"raw":    namer.NewRawNamer("", nil),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		glog.Fatalf("Failed loading boilerplate: %v", err)
	}

	header := append([]byte(fmt.Sprintf("// +build !%s\n\n", arguments.GeneratedBuildTag)), boilerplate...)
	header = append(header, []byte(
		`
		// This file was autogenerated by json-gen. Do not edit it manually!

		`)...)

	boundingDirs := []string{}
	if customArgs, ok := arguments.CustomArgs.(*CustomArgs); ok {
		for i := range customArgs.BoundingDirs {
			// Strip any trailing slashes - they are not exactly "correct" but
			// this is friendlier.
			boundingDirs = append(boundingDirs, strings.TrimRight(customArgs.BoundingDirs[i], "/"))
		}
	}

	// Results accumulator.
	packages := generator.Packages{}

	// De-dup inputs.
	inputs := map[string]bool{}
	for _, i := range context.Inputs {
		inputs[i] = true
	}

	for i := range inputs {
		glog.V(5).Infof("considering pkg %q", i)
		pkg := context.Universe[i]
		if pkg == nil {
			glog.V(5).Infof("package %q is nil in context.Universe: skipping", i)
			// If the input had no Go files, for example.
			continue
		}

		// Find the tags, if present.
		ptag := extractTag(pkg.DocComments)
		ptagValue := ""
		if ptag != nil {
			ptagValue = ptag.value
			if ptagValue != tagValuePackage {
				glog.Fatalf("Package %v: unsupported %s value: %q", i, tagName, ptagValue)
			}
			glog.V(5).Infof("  tag.value: %q", ptagValue)
		} else {
			glog.V(5).Infof("  no tag")
		}

		// If the pkg-scoped tag says to generate, we can skip scanning types.
		pkgNeedsGeneration := (ptagValue == tagValuePackage)
		if !pkgNeedsGeneration {
			// If the pkg-scoped tag did not exist, scan all types for one that
			// explicitly wants generation.
			for _, t := range pkg.Types {
				glog.V(5).Infof("  considering type %q", t.Name.String())
				ttag := extractTag(t.CommentLines)
				if ttag != nil && ttag.value == "true" {
					glog.V(5).Infof("    tag=true")
					pkgNeedsGeneration = true
					break
				}
			}
		}

		if pkgNeedsGeneration {
			packages = append(packages,
				&generator.DefaultPackage{
					PackageName: pkg.Name,
					PackagePath: pkg.Path,
					HeaderText:  header,
					GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
						return []generator.Generator{
							newGenerator(
								arguments.OutputFileBaseName, pkg.Path, //FIXME: use outputbase
								boundingDirs, (ptagValue == tagValuePackage)),
						}
					},
					FilterFunc: func(c *generator.Context, t *types.Type) bool {
						return t.Name.Package == pkg.Path
					},
				})
		}
	}
	return packages
}

// jsonGenerator produces a file with autogenerated deep-copy functions.
type jsonGenerator struct {
	generator.DefaultGen
	targetPackage string
	boundingDirs  []string
	emitAllTypes  bool
	imports       namer.ImportTracker
	astEmitted    map[string]bool
	astNeeded     []*types.Type
}

func newGenerator(sanitizedName string, targetPackage string, boundingDirs []string, emitAllTypes bool) generator.Generator {
	return &jsonGenerator{
		DefaultGen: generator.DefaultGen{
			OptionalName: sanitizedName,
		},
		targetPackage: targetPackage,
		boundingDirs:  boundingDirs,
		emitAllTypes:  emitAllTypes,
		imports:       generator.NewImportTracker(),
		astEmitted:    map[string]bool{},
	}
}

func (g *jsonGenerator) Namers(c *generator.Context) namer.NameSystems {
	// Have the raw namer for this file track what it imports.
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.targetPackage, g.imports),
	}
}

func (g *jsonGenerator) Filter(c *generator.Context, t *types.Type) bool {
	// Filter out types that are not being processed.
	if g.emitAllTypes {
		return true
	}
	if ttag := extractTag(t.CommentLines); ttag != nil && ttag.value == "true" {
		return true
	}
	return false
}

func (g *jsonGenerator) inBounds(t *types.Type) bool {
	// Only packages within the restricted range can be processed.
	if !isRootedUnder(t.Name.Package, g.boundingDirs) {
		return false
	}
	return true
}

func isRootedUnder(pkg string, roots []string) bool {
	// Add trailing / to avoid false matches, e.g. foo/bar vs foo/barn.  This
	// assumes that bounding dirs do not have trailing slashes.
	pkg = pkg + "/"
	for _, root := range roots {
		if strings.HasPrefix(string(pkg), string(root)+"/") {
			return true
		}
	}
	return false
}

var nameOfByteSlice = types.Name{Name: "[]byte"}
var nameOfError = types.Name{Name: "error"}

// hasMarshalMethod returns true if an appropriate MarshalJSON() method is
// defined for the given type.
func hasMarshalMethod(t *types.Type) bool {
	for mn, mt := range t.Methods {
		if mn != "MarshalJSON" {
			continue
		}
		if len(mt.Signature.Parameters) != 0 {
			return false
		}
		if len(mt.Signature.Results) != 2 ||
			mt.Signature.Results[0].Name != nameOfByteSlice ||
			mt.Signature.Results[1].Name != nameOfError {
			return false
		}
		return true
	}
	return false
}

// hasUnmarshalMethod returns true if an appropriate UnmarshalJSON() method is
// defined for the given type.
func hasUnmarshalMethod(t *types.Type) bool {
	for mn, mt := range t.Methods {
		if mn != "UnmarshalJSON" {
			continue
		}
		if len(mt.Signature.Parameters) != 1 || mt.Signature.Parameters[0].Name != nameOfByteSlice {
			return false
		}
		if len(mt.Signature.Results) != 1 || mt.Signature.Results[0].Name != nameOfError {
			return false
		}
		return true
	}
	return false
}

// hasTextMarshalMethod returns true if an appropriate MarshalText() method is
// defined for the given type.
func hasTextMarshalMethod(t *types.Type) bool {
	for mn, mt := range t.Methods {
		if mn != "MarshalText" {
			continue
		}
		if len(mt.Signature.Parameters) != 0 {
			return false
		}
		if len(mt.Signature.Results) != 2 ||
			mt.Signature.Results[0].Name != nameOfByteSlice ||
			mt.Signature.Results[1].Name != nameOfError {
			return false
		}
		return true
	}
	return false
}

func (g *jsonGenerator) isOtherPackage(importLine string) bool {
	if strings.HasSuffix(string(importLine), "\""+string(g.targetPackage)+"\"") {
		return false
	}
	return true
}

func (g *jsonGenerator) Imports(c *generator.Context) []string {
	importLines := []string{}
	for _, singleImport := range g.imports.ImportLines() {
		if g.isOtherPackage(singleImport) {
			importLines = append(importLines, singleImport)
		}
	}
	return importLines
}

func argsFromType(t *types.Type) generator.Args {
	return generator.Args{
		"type": t,
	}
}

func (g *jsonGenerator) Init(c *generator.Context, w io.Writer) error {
	g.imports.Add("k8s.io/gengo/examples/json-gen/libjson")
	return nil
}

func (g *jsonGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	glog.V(5).Infof("generating for type %v", t)
	if !g.needsGeneration(t) {
		glog.V(5).Infof("type %v does not need generation", t)
		return nil
	}

	sw := generator.NewSnippetWriter(w, c, "$", "$")
	g.emitAST(t, c, sw)
	g.emitMethods(t, sw)
	return sw.Error()
}

func (g *jsonGenerator) needsGeneration(t *types.Type) bool {
	if t.Kind == types.DeclarationOf {
		return false
	}

	tag := extractTag(t.CommentLines)
	tv := ""
	if tag != nil {
		tv = tag.value
		if tv != "true" && tv != "false" {
			glog.Fatalf("Type %v: unsupported %s value: %q", t, tagName, tag.value)
		}
	}
	if g.emitAllTypes && tv == "false" {
		// The whole package is being generated, but this type has opted out.
		glog.V(5).Infof("not generating for type %v because type opted out", t)
		return false
	}
	if !g.emitAllTypes && tv != "true" {
		// The whole package is NOT being generated, and this type has NOT opted in.
		glog.V(5).Infof("not generating for type %v because type did not opt in", t)
		return false
	}
	return true
}

func (g *jsonGenerator) emitAST(t *types.Type, c *generator.Context, sw *generator.SnippetWriter) {
	glog.V(3).Infof("emitting AST code for type %v", t)
	sw.Do(`
		func ast_$.type|public$(obj *$.type|raw$) (libjson.Value, error) {
			`+g.emitBodyFor(t, c)+`
		}
		`, argsFromType(t))
	g.astEmitted[t.String()] = true

}

func (g *jsonGenerator) emitMethods(t *types.Type, sw *generator.SnippetWriter) {
	if rootType(t).Kind == types.Pointer {
		//FIXME: should be fatal?
		glog.Errorf("not emitting methods for pointer type %v", t)
		return
	}
	if !hasMarshalMethod(t) {
		g.imports.Add("bytes")
		sw.Do(`
		func (obj $.type|raw$) MarshalJSON() ([]byte, error) {
			jv, err := ast_$.type|public$(&obj)
			if err != nil {
				return nil, err
			}
			var buf bytes.Buffer
			if err := jv.Render(&buf); err != nil {
				return nil, err
			}
			return buf.Bytes(), nil
		}
		`, argsFromType(t))
	}
	if !hasUnmarshalMethod(t) {
		g.imports.Add("bytes")
		sw.Do(`
		func (obj *$.type|raw$) UnmarshalJSON(data []byte) error {
			jv, err := ast_$.type|public$(obj)
			if err != nil {
				return err
			}
			return jv.Parse(data)
		}
		`, argsFromType(t))
	}
}

// Just a shortcut helper function.
func formatName(c *generator.Context, namer string, t *types.Type) string {
	return c.Namers[namer].Name(t)
}

// emitBodyFor emits a block of code which returns a libjson.Value
// representing an instance of t, or an error.
func (g *jsonGenerator) emitBodyFor(t *types.Type, c *generator.Context) string {
	/*
		// If the type implements Marshaler on its own, use that.
		if hasMarshalMethod(t) {
			glog.V(3).Infof("type %v has a MarshalJSON() method", t)
			g.imports.Add("fmt")
			return `
				if b, err := obj.MarshalJSON(); err != nil {
					return nil, fmt.Errorf("failed %T.MarshalJSON: %v", obj, err)
				} else {
					return libjson.Raw(string(b)), nil
				}
				`
		}
		if hasTextMarshalMethod(t) {
			glog.V(3).Infof("type %v has a MarshalText() method", t)
			return `
				if b, err := obj.MarshalText(); err != nil {
					return nil, fmt.Errorf("failed %T.MarshalText: %v", obj, err)
				} else {
					return libjson.String(string(b)), nil
				}
				`
		}
	*/

	// Just call another function for simple cases.
	if t.Kind == types.Alias {
		glog.V(4).Infof("type %v is alias to %v", t, t.Underlying)
		return `return ` + g.emitCallFor(t.Underlying, c)
	}

	var f func(*types.Type, *generator.Context) string

	// Handle more complex cases.  Each of these functions will emit a block
	// of code to marshal the respective types.  The emitted code should return
	// if it hits an error, but NOT if it is successful.  This allows these
	// blocks to be used in multiple contexts, and the contexts will handle
	// successful returns.
	glog.V(4).Infof("type %v is kind %s", t, t.Kind)
	switch t.Kind {
	case types.Builtin:
		f = g.emitBodyForBuiltin
	case types.Pointer:
		f = g.emitBodyForPointer
	case types.Struct:
		f = g.emitBodyForStruct
	case types.Slice:
		f = g.emitBodyForSlice
		/*
			case types.Map:
				f = g.emitBodyForMap
			default:
				// A reasonable argument could be made to just ignore it, but I'd
				// rather know if this happens.  The likely case is interfaces which,
				// obviously, we can't codegen for.
				panic("unsupported kind: " + string(t.Kind) + " for type " + string(t.String()))
		*/
	}
	if f != nil {
		return f(t, c)
	}
	return "//FIXME\nreturn libjson.Raw(`\"fixme\"`), nil"
}

func (g *jsonGenerator) emitCallFor(t *types.Type, c *generator.Context) string {
	g.astNeeded = append(g.astNeeded, t)
	//FIXME: name for anon structs is horrible - hash?
	return `ast_` + formatName(c, "public", t) + `((*` + formatName(c, "raw", t) + `)(obj))`
}

func (g *jsonGenerator) emitBodyForPointer(t *types.Type, c *generator.Context) string {
	//FIXME: check that obj in the func isn't reusing the same ptr
	return `
		{
			p := obj
			obj := *obj
			if obj == nil {
				obj = new(` + formatName(c, "raw", t.Elem) + `)
			}
			jv, err := ` + g.emitCallFor(t.Elem, c) + `
			if err != nil {
				return nil, err
			}
			setNull := func(b bool) {
				if b {
					*p = nil
				} else {
					*p = obj
				}
			}
			return libjson.NewNullable(jv, *p == nil, setNull), nil
		}
		`
}

func (g *jsonGenerator) emitBodyForStruct(t *types.Type, c *generator.Context) string {
	result := `
		result := libjson.Object{}
		`

	if len(t.Members) == 0 {
		// at least do something with args to avoid "not used" errors
		result += "_ = obj\n"
	}

	structMeta := collectFields(t, 0, "")
	for _, name := range structMeta.FieldNames {
		field := structMeta.Fields[name]
		//FIXME: doesn't handle recursive types.
		glog.V(4).Infof("descending into field %v.%s", t, field.FieldName)
		result += `
			// ` + field.FieldName + ` ` + field.Type.String() + `
			{
			    obj := &obj.` + field.FieldName + `
				_ = obj //FIXME: remove when other Kinds are done
			`

		if field.OmitEmpty {
			glog.V(4).Infof("field %v.%s has omitempty tag", t, field.FieldName)
			result += `
				empty := func(jv libjson.Value) bool { return jv.Empty() }
				`
		} else {
			result += `
				empty := func(libjson.Value) bool { return false }
				`
		}

		if field.String {
			glog.V(4).Infof("field %v.%s has string tag", t, field.FieldName)
			result += `
				finalize := func(jv libjson.Value) (libjson.Value, error) {
					buf := bytes.Buffer{}
					if err := jv.Render(&buf); err != nil {
						return nil, err
					}
					p := new(string)
					*p = buf.String()
					return libjson.NewString(func() string { return *p }, func(s string) { *p = s }), nil
				}
				`
		} else {
			result += `
				finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }
				`
		}

		//FIXME: the empty test fails for decode?
		result += `
				jv, err := ` + g.emitCallFor(field.Type, c) + `
				if err != nil {
					return nil, err
				}
				if !empty(jv) {
					fv, err := finalize(jv)
					if err != nil {
						return nil, err
					}
					p := new(string)
					*p = "` + name + `"
					nv := libjson.NamedValue{
						Name: libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
						Value: fv,
					}
					result = append(result, nv)
				} else { panic("TIM: ` + name + ` was empty") } //FIXME:
			}
			`
	}
	result += `
	    return result, nil
		`
	return result
}

// Keep track of fields as we process embedded structs.  This is required to be
// compatible with Go's `json` package.
type fieldMeta struct {
	FieldName string
	Type      *types.Type
	Nesting   int  // how many nesting levels
	Tagged    bool // had a tag-defined name
	OmitEmpty bool // the `omitempty` tag opion
	String    bool // the `string` tag option
}

type structMeta struct {
	Type       *types.Type
	FieldNames []string
	Fields     map[string]*fieldMeta
}

func newStructMeta(t *types.Type) structMeta {
	return structMeta{
		Type:       t,
		FieldNames: []string{},
		Fields:     map[string]*fieldMeta{},
	}
}

func mergeField(sm *structMeta, name string, fm *fieldMeta) {
	existing := sm.Fields[name]
	if existing == nil {
		sm.FieldNames = append(sm.FieldNames, name)
		sm.Fields[name] = fm
	} else if existing.Nesting != fm.Nesting {
		keep := existing
		drop := fm
		if fm.Nesting < existing.Nesting {
			sm.Fields[name] = fm
			keep = fm
			drop = existing
		}
		glog.Errorf("WARNING: JSON field %s.%s conflict: keeping %s, dropping %s (less nested)", sm.Type, name, keep.FieldName, drop.FieldName)
	} else if existing.Tagged != fm.Tagged {
		keep := existing
		drop := fm
		if fm.Tagged {
			sm.Fields[name] = fm
			keep = fm
			drop = existing
		}
		glog.Errorf("WARNING: JSON field %s.%s conflict: keeping %s, dropping %s (tagged)", sm.Type, name, keep.FieldName, drop.FieldName)
	} else {
		glog.Errorf("WARNING: JSON field %s.%s conflict: dropping both %s and %s", sm.Type, name, existing.FieldName, fm.FieldName)
	}
}

// collectFields returns a flattened map of field information for the given
// type. This processes embedded fields according to Go's `json` package docs.
func collectFields(t *types.Type, nesting int, fieldpath string) structMeta {
	structMeta := newStructMeta(t)

	for _, m := range t.Members {
		fm := &fieldMeta{
			FieldName: prefixFieldName(m.Name, fieldpath),
			Type:      m.Type,
			Nesting:   nesting,
		}

		name := ""
		if m.Tags != "" {
			glog.V(3).Infof("found struct tags for %v.%s", t, m.Name)
			tag := parseTag(m.Tags)
			name = tag.name
			if name == "-" {
				// Skip this field
				continue
			}
			fm.Tagged = true
			fm.OmitEmpty = tag.omitempty
			fm.String = tag.asString
		}
		if name == "" {
			name = m.Name
		}

		// If the field is either not embedded or is not a struct (e.g. an
		// embedded string), save it.
		if !m.Embedded || m.Type.Kind != types.Struct {
			mergeField(&structMeta, name, fm)
			continue
		}

		embeddedStructMeta := collectFields(m.Type, nesting+1, prefixFieldName(m.Name, fieldpath))
		for _, name := range embeddedStructMeta.FieldNames {
			field := embeddedStructMeta.Fields[name]
			mergeField(&structMeta, name, field)
		}
	}
	return structMeta
}

func prefixFieldName(name string, path string) string {
	if path == "" {
		return name
	}
	return path + "." + name
}

type jsonTag struct {
	name      string
	omitempty bool
	asString  bool
}

func parseTag(str string) jsonTag {
	tag := reflect.StructTag(str)

	result := jsonTag{}

	jt := tag.Get("json")
	parts := strings.Split(jt, ",")
	if len(parts) >= 1 {
		result.name = parts[0]
	}
	for i := 1; i < len(parts); i++ {
		if parts[i] == "omitempty" {
			result.omitempty = true
		} else if parts[i] == "string" {
			result.asString = true
		} else {
			glog.Errorf("WARNING: unknown json tag option %q in json tag", parts[i])
		}
	}
	return result
}

func (g *jsonGenerator) emitBodyForSlice(t *types.Type, c *generator.Context) string {
	//FIXME: need rootType here and elsewhere?
	result := ""
	if rootType(t.Elem) == types.Byte {
		// Go's json package special-cases []byte
		g.imports.Add("encoding/base64")
		result += `
			buf := bytes.Buffer{}
			b64 := base64.NewEncoder(base64.StdEncoding, &buf)
			`
		if t.Elem == types.Byte {
			result += `
				b64.Write(*obj)
				`
		} else {
			// Can't just cast []ByteAlias to []byte. :(
			result += `
				for _, b := range *obj {
					b64.Write([]byte{byte(b)})
				}
				`
		}
		result += `
			b64.Close()
			p := new(string)
			*p = buf.String()
			return libjson.NewString(func() string { return *p }, func(s string) { *p = s }), nil
			`
	} else {
		result += `
			get := func() ([]libjson.Value, error) {
				if *obj == nil {
					return nil, nil
				}
				result := []libjson.Value{}
				for i := range *obj {
					obj := &(*obj)[i]
					//FIXME: do any of these ACTUALLY return an error?
					jv, err := ` + g.emitCallFor(t.Elem, c) + `
					if err != nil {
						return nil, err
					}
					result = append(result, jv)
				}
				return result, nil
			}
			add := func() libjson.Value {
				var x ` + formatName(c, "raw", t.Elem) + `
				*obj = append(*obj, x)
				obj := &(*obj)[len(*obj)-1]
				jv, _ := ` + g.emitCallFor(t.Elem, c) + `
				//FIXME: handle error?
				return jv
			}
			setNull := func(b bool) {
				if b {
					*obj = nil
				} else {
					*obj = []` + formatName(c, "raw", t.Elem) + `{}
				}
			}
			return libjson.NewArray(*obj == nil, get, add, setNull), nil
		`
	}
	return result
}

func rootType(t *types.Type) *types.Type {
	// Peel away layers of alias.
	for t.Kind == types.Alias {
		t = t.Underlying
	}
	return t
}

/* FIXME: not done yet
func (g *jsonGenerator) emitBodyForMap(t *types.Type, c *generator.Context) string {
	//FIXME: need root Type here and elsewhere?
	result := ""

	// Map keys must be derived from string, encoding.TextMarshaler, or integral types.
	if rootType(t.Key) == types.String {
		result += `
			stringify := func(s ` + formatName(c, "raw", t.Key) + `) (string, error) {
				return string(s), nil
			}
			`
	} else if hasTextMarshalMethod(t.Key) {
		glog.V(3).Infof("type %v has a MarshalText() method", t)
		result += `
			stringify := func(tm ` + formatName(c, "raw", t.Key) + `) (string, error) {
				if b, err := tm.MarshalText(); err != nil {
					return "", fmt.Errorf("failed %T.MarshalText: %v", obj, err)
				} else {
					return string(b), nil
				}
			}
			`
	} else {
		switch rootType(t.Key) {
		case types.Int, types.Int64, types.Int32, types.Int16, types.Int8:
			g.imports.Add("strconv")
			result += `
			stringify := func(i ` + formatName(c, "raw", t.Key) + `) (string, error) {
				return strconv.FormatInt(int64(i), 10), nil
			}
			`
		case types.Uint, types.Uint64, types.Uint32, types.Uint16, types.Uint8, types.Uintptr:
			g.imports.Add("strconv")
			result += `
			stringify := func(i ` + formatName(c, "raw", t.Key) + `) (string, error) {
				return strconv.FormatUint(uint64(i), 10), nil
			}
			`
		default:
			glog.Error("WARNING: map key %v is not string, int, or TextMarshaler: ignoring", t.Key)
		}
	}

	//TODO(thockin): It would be nice to sort ints by numeric value.  The
	//standard `sort` package doesn't have functions for a int64 and uint64,
	//and I am too lazy to emit those myself right now.  Idea is to move the
	//defn of `keys` into the above sections and provide a `keyify` func and a
	//`sort` func.  Main loop extracts into `keys` via `keyify`, and then calls
	//`sort`, then emits via stringify.
	g.imports.Add("sort")
	return result + `
		m := make(map[string]libjson.Value, len(obj))
		keys := make([]string, 0, len(obj))
		for k, v := range obj {
			ks, err := stringify(k)
			if err != nil {
				return nil, err
			}
			keys = append(keys, ks)
			obj := v
			jv, err := ` + g.emitCallFor(t.Elem, c) + `
			if err != nil {
				return nil, err
			}
			m[ks] = jv
		}
		result := libjson.Object{}
		sort.Strings(keys)
		for _, ks := range keys {
			nv := libjson.NamedValue{
				Name: libjson.String(ks),
				Value: m[ks],
			}
			result = append(result, nv)
		}
	    return result, nil
		`
}
*/

func (g *jsonGenerator) Finalize(c *generator.Context, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	// loop because the list can be mutated while emitting
	for len(g.astNeeded) > 0 {
		todo := g.astNeeded
		g.astNeeded = nil
		glog.V(5).Infof("%d more types to emit", len(todo))
		for _, t := range todo {
			if g.astEmitted[t.String()] == true {
				continue
			}
			g.emitAST(t, c, sw)
		}
	}
	return sw.Error()
}

func (g *jsonGenerator) emitBodyForBuiltin(t *types.Type, c *generator.Context) string {
	switch t {
	case types.String:
		return `return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil`
	case types.Bool:
		return `return libjson.NewBool(func() bool { return *obj }, func(b bool) { *obj = b }), nil`
	case types.Int, types.Int64, types.Int32, types.Int16, types.Int8:
		fallthrough
	case types.Uint, types.Uint64, types.Uint32, types.Uint16, types.Uint8, types.Uintptr, types.Byte:
		return `
			get := func() float64 {
				return float64(*obj)
			}
			set := func(f float64) {
				*obj = ` + formatName(c, "raw", t) + `(f)
			}
			return libjson.NewInt(get, set), nil
			`
	case types.Float64:
		return `
			get := func() float64 {
				return float64(*obj)
			}
			set := func(f float64) {
				*obj = ` + formatName(c, "raw", t) + `(f)
			}
			return libjson.NewFloat(64, get, set), nil
			`
	case types.Float32:
		return `
			get := func() float64 {
				return float64(*obj)
			}
			set := func(f float64) {
				*obj = ` + formatName(c, "raw", t) + `(f)
			}
			return libjson.NewFloat(32, get, set), nil
			`
	default:
		// This is a legit bug in the tool.
		panic("unknown builtin \"" + t.String() + "\"")
	}
}

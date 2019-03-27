package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const (
	analyzedFile = "hero_stats.go"

	mainStructName = "HeroStats"

	tagName        = "stats"
	tagStartName   = "start"
	tagEndName     = "end"
	tagDefaultName = "default"
)

// TODO: Use value from "default" tag if set.
// Generate setters and getters for the selecting fields
func main() {
	var (
		pkgScan string
		getOut  string
		setOut  string
		addOut  string
	)
	flag.StringVar(&pkgScan, "scan", "./backend/model", "Path to directory which should be scanned for the *.go files.")
	flag.StringVar(&getOut, "getters", "./backend/ranking/getters.go", "Output file for getters")
	flag.StringVar(&setOut, "setters", "./backend/ranking/setters.go", "Output file for setters")
	flag.StringVar(&addOut, "adders", "./backend/ranking/adders.go", "Output file for adders")
	flag.Parse()

	generate(pkgScan, getOut, setOut, addOut)
}

func generate(pkgScan, getOut, setOut, addOut string) {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, pkgScan, nil, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	astf := []*ast.File{}
	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			if strings.HasSuffix(fileName, analyzedFile) {
				astf = append(astf, file)
			}
		}
	}

	config := &types.Config{
		Error: func(err error) {
			log.Fatal(err)
		},
		Importer: importer.Default(),
	}
	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	_, err = config.Check(pkgScan, fset, astf, &info)
	if err != nil {
		log.Fatal(err)
	}

	w := &Walker{
		info:          &info,
		pkgScan:       pkgScan,
		gettersOutput: getOut,
		settersOutput: setOut,
		addersOutput:  addOut,
	}

	for _, f := range astf {
		ast.Walk(w, f)
	}

	w.RenderFile()
}

type Walker struct {
	info          *types.Info
	pkgScan       string
	gettersOutput string
	settersOutput string
	addersOutput  string

	Getters []Getter
	Setters []Setter
	Adders  []Adder

	structName string
}

// Visit is an implementation of ast.Visitor
func (w *Walker) Visit(node ast.Node) ast.Visitor {
	if node != nil {
		switch n := node.(type) {
		case *ast.TypeSpec:
			w.structName = n.Name.String()
		case *ast.StructType:
			// Simple validation of anonymous/embedded structs
			if w.structName == "" {
				log.Fatal("Oopsie, anonymous struct - can't be")
			}

			w.analyzeStruct(n)

			// Reset to test if trying to save anonymous struct
			// (in case of validation)
			w.structName = ""
		}
	}

	return w
}

type Field struct {
	StatsKey       string
	FieldName      string
	EmbeddedStruct string
	FullFieldName  string
	DefaultValue   string
}

type Getter struct {
	Field     Field
	GetterTpl string
}

type Setter struct {
	Field     Field
	SetterTpl string
}

type Adder struct {
	Field    Field
	AdderTpl string
}

func (w *Walker) analyzeStruct(str *ast.StructType) {
	for _, field := range str.Fields.List {
		typ := w.info.TypeOf(field.Type)

		if w.IsEmbededField(typ) {
			log.Printf("tag: field with type %v is embedded, ommiting...", typ.String())
			continue
		}

		fieldName := field.Names[0].String()

		// Parse tag
		tags := ParseFieldTags(field.Tag)

		// Extract identifier from tags
		statsKey, err := tags.StringVal(tagName)
		if err != nil {
			log.Printf("tag: field %s without tag '%s', ignoring...", field.Names[0], tagName)
			continue
		}

		start, err := tags.IntVal(tagStartName)
		if err != nil {
			start = -1
		}

		end, err := tags.IntVal(tagEndName)
		if err != nil && start != -1 {
			log.Fatalf("tag: field with name %s has set 'start' tag but 'end' is missing", field.Names[0])
		}

		// Customize getters and setters for specified statsKey
		var (
			customGetter = "basicGetter"
			customSetter = "basicSetter"
			customAdder  = "basicAdder"

			defaultValue string
		)

		switch statsKey {
		case "c_apr", "c_emo", "c_eqp", "c_items":
			customGetter = "inventoryGetter"
			customSetter = "inventorySetter"

			customAdder = ""
		}

		switch typ.String() {
		case "string", "map[int]string":
			def, err := tags.StringVal(tagDefaultName)
			if err == nil {
				defaultValue = `"` + def + `"`
				customGetter = "basicGetterWithDefault"
			}
		case "[]string":
			def, err := tags.StringVal(tagDefaultName)
			if err == nil {
				defaultValue = `"` + def + `"`
				customGetter = "inventoryGetterWithDefault"
			}
		}

		switch true {
		case start == -1:
			// FIXME: currently ignoring multi-named variables
			w.AddGetterAndSetter(
				statsKey,
				fieldName,
				customGetter,
				customSetter,
				customAdder,
				defaultValue,
			)

		case end > 0:
			// Add as many as ends
			for index := start; index <= end; index++ {
				w.AddGetterAndSetter(
					fmt.Sprintf("%s%d", statsKey, index),
					fmt.Sprintf("%s[%d]", fieldName, index),
					customGetter,
					customSetter,
					customAdder,
					defaultValue,
				)
			}
		}
	}
}

func (w *Walker) IsEmbededField(typ types.Type) bool {
	return strings.HasPrefix(typ.String(), w.pkgScan)
}

func (w *Walker) AddGetterAndSetter(statsKey, fieldName, getterTpl, setterTpl, adderTpl, defaultValue string) {
	var (
		embeddedStruct string
		fullFieldName  = fieldName
	)

	if w.structName != mainStructName {
		embeddedStruct = w.structName
		fullFieldName = w.structName + "." + fieldName
	}

	field := Field{
		StatsKey:       statsKey,
		FieldName:      fieldName,
		EmbeddedStruct: embeddedStruct,
		FullFieldName:  fullFieldName,
		DefaultValue:   defaultValue,
	}

	// Add getter
	w.AddGetter(Getter{
		Field:     field,
		GetterTpl: getterTpl,
	})

	// Add setter
	w.AddSetter(Setter{
		Field:     field,
		SetterTpl: setterTpl,
	})

	if adderTpl != "" {
		w.AddAdder(Adder{
			Field:    field,
			AdderTpl: adderTpl,
		})
	}
}

func (w *Walker) AddGetter(getter ...Getter) {
	w.Getters = append(w.Getters, getter...)
}

func (w *Walker) AddSetter(setter ...Setter) {
	w.Setters = append(w.Setters, setter...)
}

func (w *Walker) AddAdder(adder ...Adder) {
	w.Adders = append(w.Adders, adder...)
}

func (w *Walker) RenderFile() {
	if err := w.renderGettersGo(); err != nil {
		log.Fatal(err)
	}

	if err := w.renderSettersGo(); err != nil {
		log.Fatal(err)
	}

	if err := w.renderAddersGo(); err != nil {
		log.Fatal(err)
	}
}

func (w *Walker) renderGettersGo() error {
	sort.Slice(w.Getters, func(i, j int) bool {
		return w.Getters[i].Field.StatsKey < w.Getters[i].Field.StatsKey

	})

	t := template.New("getters.go")

	w.registerCustomFuncs(t)
	w.registerGettersGo(t)
	w.registerHeader(t)
	w.registerBasicGetter(t)
	w.registerBasicGetterWithDefault(t)
	w.registerInventoryGetter(t)
	w.registerInventoryGetterWithDefault(t)

	file, err := os.Create(w.gettersOutput)
	if err != nil {
		return err
	}

	return t.Execute(file, w)
}

func (w *Walker) renderSettersGo() error {
	sort.Slice(w.Setters, func(i, j int) bool {
		return w.Setters[i].Field.StatsKey < w.Setters[i].Field.StatsKey
	})

	t := template.New("setters.go")

	w.registerCustomFuncs(t)
	w.registerSettersGo(t)
	w.registerHeader(t)
	w.registerBasicSetter(t)
	w.registerInventorySetter(t)

	file, err := os.Create(w.settersOutput)
	if err != nil {
		return err
	}

	return t.Execute(file, w)
}

func (w *Walker) renderAddersGo() error {
	sort.Slice(w.Setters, func(i, j int) bool {
		return w.Setters[i].Field.StatsKey < w.Setters[i].Field.StatsKey
	})

	t := template.New("adders.go")

	w.registerCustomFuncs(t)
	w.registerAddersGo(t)
	w.registerHeader(t)
	w.registerBasicAdder(t)

	file, err := os.Create(w.addersOutput)
	if err != nil {
		return err
	}

	return t.Execute(file, w)
}

func (w *Walker) registerCustomFuncs(t *template.Template) *template.Template {
	return t.Funcs(map[string]interface{}{
		"nestedTemplate": func(name string, data interface{}) (string, error) {
			buf := new(bytes.Buffer)
			err := t.ExecuteTemplate(buf, name, data)
			return buf.String(), err
		},
	})
}

func (w *Walker) registerHeader(t *template.Template) *template.Template {
	return template.Must(t.New("header").Parse(
		`package ranking

// This file was generated by go:generate
// Do not do any modifications here
`,
	))
}

func (w *Walker) registerAddersGo(t *template.Template) *template.Template {
	return template.Must(t.Parse(
		`{{ template "header" }}
import (
	"strconv"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
)

var (
	adders = map[string]Adder{
{{range $a := .Adders }}		"{{ $a.Field.StatsKey }}": {{ nestedTemplate $a.AdderTpl $a }},
{{end}}	}
)
`,
	))
}

func (w *Walker) registerBasicAdder(t *template.Template) *template.Template {
	return template.Must(t.New("basicAdder").Parse(
		`func(p *model.` + mainStructName + `, value string, pt int) error {
			addition, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			current, err := strconv.ParseFloat(p.{{ .Field.FullFieldName }}, 64)
			if err != nil {
				return err
			}
			sum := current + addition
			p.{{ .Field.FullFieldName }} = strconv.FormatFloat(sum, 'f', 4, 64)
			return nil
		}`,
	))
}

func (w *Walker) registerGettersGo(t *template.Template) *template.Template {
	return template.Must(t.Parse(
		`{{ template "header" }}
import (
	"strings"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
)

var (
	getters = map[string]Getter{
{{range $g := .Getters }}		"{{ $g.Field.StatsKey }}": {{ nestedTemplate $g.GetterTpl $g }},
{{end}}	}
)
`,
	))
}

func (w *Walker) registerBasicGetter(t *template.Template) *template.Template {
	return template.Must(t.New("basicGetter").Parse(
		`func(p *model.` + mainStructName + `) (string, error) {
			return p.{{ .Field.FullFieldName }}, nil
		}`,
	))
}

func (w *Walker) registerBasicGetterWithDefault(t *template.Template) *template.Template {
	return template.Must(t.New("basicGetterWithDefault").Parse(
		`func(p *model.` + mainStructName + `) (string, error) {
			if p.{{ .Field.FullFieldName }} == "" {
				return {{ .Field.DefaultValue }}, nil
			}
			return p.{{ .Field.FullFieldName }}, nil
		}`,
	))
}

func (w *Walker) registerInventoryGetter(t *template.Template) *template.Template {
	return template.Must(t.New("inventoryGetter").Parse(
		`func(p *model.` + mainStructName + `) (string, error) {
			return strings.Join(p.{{ .Field.FullFieldName }}, ";"), nil
		}`,
	))
}

func (w *Walker) registerInventoryGetterWithDefault(t *template.Template) *template.Template {
	return template.Must(t.New("inventoryGetterWithDefault").Parse(
		`func(p *model.` + mainStructName + `) (string, error) {
			if len(p.{{ .Field.FullFieldName }}) == 0 {
				return {{ .Field.DefaultValue }}, nil
			}
			return strings.Join(p.{{ .Field.FullFieldName }}, ";"), nil
		}`,
	))
}

func (w *Walker) registerSettersGo(t *template.Template) *template.Template {
	return template.Must(t.Parse(
		`{{ template "header" . }}
import (
	"strings"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
)

var (
	setters = map[string]Setter{
{{range $s := .Setters }}		"{{ $s.Field.StatsKey }}": {{ nestedTemplate $s.SetterTpl $s }},
{{end}}	}
)
`,
	))
}

func (w *Walker) registerBasicSetter(t *template.Template) *template.Template {
	return template.Must(t.New("basicSetter").Parse(
		`func(p *model.` + mainStructName + `, value string, pt int) error {
			p.{{ .Field.FullFieldName }} = value
			return nil
		}`,
	))
}

func (w *Walker) registerInventorySetter(t *template.Template) *template.Template {
	return template.Must(t.New("inventorySetter").Parse(
		`func(p *model.` + mainStructName + `, value string, pt int) error {
			p.{{ .Field.FullFieldName }} = strings.Split(value, ";")
			return nil
		}`,
	))
}

func ParseFieldTags(field *ast.BasicLit) FieldTag {
	ft := FieldTag{
		tags: map[string]string{},
	}
	if field == nil {
		return ft
	}

	rawTags := strings.Trim(field.Value, "`")

	for _, t := range strings.Split(rawTags, " ") {
		key, value := separateTags(t)

		ft.tags[key] = value
	}

	return ft
}

// separate tags to be recognized as key-value
func separateTags(singleTag string) (string, string) {
	tuple := strings.SplitN(singleTag, ":", 2)
	return tuple[0], tuple[1]
}

type FieldTag struct {
	tags map[string]string
}

func (ft *FieldTag) StringVal(tag string) (string, error) {
	value, ok := ft.tags[tag]
	if !ok {
		return "", fmt.Errorf("tag: %s not found", tag)
	}

	removedQuotes := strings.Trim(value, `"`)

	return removedQuotes, nil
}

func (ft *FieldTag) StringArr(tag string) ([]string, error) {
	s, err := ft.StringVal(tag)
	if err != nil {
		return nil, err
	}

	return strings.Split(s, ","), nil
}

func (ft *FieldTag) IntVal(tag string) (int, error) {
	s, err := ft.StringVal(tag)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}

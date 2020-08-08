package simpleyaml

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"

	"pkg.re/yaml.v2"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Yaml is YAML struct
type Yaml struct {
	data interface{}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Errors
var (
	ErrYAMLIsNil                = errors.New("Yaml struct or data is nil")
	ErrMapTypeAssertion         = errors.New("Type assertion to map[string]interface{} failed")
	ErrArrayTypeAssertion       = errors.New("Type assertion to []interface{} failed")
	ErrStringArrayTypeAssertion = errors.New("Type assertion to []string failed")
	ErrBoolTypeAssertion        = errors.New("Type assertion to bool failed")
	ErrStringTypeAssertion      = errors.New("Type assertion to string failed")
	ErrByteTypeAssertion        = errors.New("Type assertion to []byte failed")
	ErrIntTypeAssertion         = errors.New("Type assertion to int failed")
	ErrFloatTypeAssertion       = errors.New("Type assertion to float64 failed")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// New returns a pointer to a new, empty `Yaml` object
func New() *Yaml {
	return &Yaml{
		data: make(map[string]interface{}),
	}
}

// NewYaml returns a pointer to a new `Yaml` object after unmarshaling `body` bytes
func NewYaml(body []byte) (*Yaml, error) {
	var data interface{}

	err := yaml.Unmarshal(body, &data)

	if err != nil {
		return nil, fmt.Errorf("Can't unmarshal data: %v", err)
	}

	return &Yaml{data}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Interface returns the underlying data
func (y *Yaml) Interface() interface{} {
	if y == nil || y.data == nil {
		return nil
	}

	return y.data
}

// Encode returns it's marshaled data as `[]byte`
func (y *Yaml) Encode() ([]byte, error) {
	return y.MarshalYAML()
}

// MarshalYAML implements the yaml.Marshaler interface
func (y *Yaml) MarshalYAML() ([]byte, error) {
	if y == nil || y.data == nil {
		return nil, ErrYAMLIsNil
	}

	return yaml.Marshal(&y.data)
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
func (y *Yaml) UnmarshalYAML(p []byte) error {
	if y == nil {
		return ErrYAMLIsNil
	}

	return yaml.Unmarshal(p, &y.data)
}

// Int type asserts to an `int`
func (y *Yaml) Int() (int, error) {
	if y == nil || y.data == nil {
		return 0, ErrYAMLIsNil
	}

	i, ok := (y.data).(int)

	if ok {
		return i, nil
	}

	return 0, ErrIntTypeAssertion
}

// Float type asserts to a `float64`
func (y *Yaml) Float() (float64, error) {
	if y == nil || y.data == nil {
		return 0.0, ErrYAMLIsNil
	}

	f, ok := (y.data).(float64)

	if ok {
		return f, nil
	}

	return 0.0, ErrFloatTypeAssertion
}

// Bool type asserts to a `bool`
func (y *Yaml) Bool() (bool, error) {
	if y == nil || y.data == nil {
		return false, ErrYAMLIsNil
	}

	b, ok := (y.data).(bool)

	if ok {
		return b, nil
	}

	return false, ErrBoolTypeAssertion
}

// String type asserts to a `string`
func (y *Yaml) String() (string, error) {
	if y == nil || y.data == nil {
		return "", ErrYAMLIsNil
	}

	s, ok := (y.data).(string)

	if ok {
		return s, nil
	}

	return "", ErrStringTypeAssertion
}

// Bytes type asserts to a `[]byte`
func (y *Yaml) Bytes() ([]byte, error) {
	if y == nil || y.data == nil {
		return nil, ErrYAMLIsNil
	}

	s, ok := (y.data).(string)

	if ok {
		return []byte(s), nil
	}

	return nil, ErrByteTypeAssertion
}

// Map type asserts to an `map`
func (y *Yaml) Map() (map[interface{}]interface{}, error) {
	if y == nil || y.data == nil {
		return nil, ErrYAMLIsNil
	}

	m, ok := (y.data).(map[interface{}]interface{})

	if ok {
		return m, nil
	}

	return nil, ErrMapTypeAssertion
}

// Array type asserts to an `array`
func (y *Yaml) Array() ([]interface{}, error) {
	if y == nil || y.data == nil {
		return nil, ErrYAMLIsNil
	}

	a, ok := (y.data).([]interface{})

	if ok {
		return a, nil
	}

	return nil, ErrArrayTypeAssertion
}

// Dump returns string representation of any kind of data
func (y *Yaml) Dump() string {
	if y == nil || y.data == nil {
		return ""
	}

	return fmt.Sprintf("%v", y.data)
}

// StringArray type asserts to an `array` of `string`
func (y *Yaml) StringArray() ([]string, error) {
	if y == nil || y.data == nil {
		return nil, ErrYAMLIsNil
	}

	a, err := y.Array()

	if err != nil {
		return nil, err
	}

	return convertSlice(a)
}

// MustArray guarantees the return of a `[]interface{}` (with optional default)
//
// useful when you want to interate over array values in a succinct manner:
//		for i, v := range yaml.Get("results").MustArray() {
//			fmt.Println(i, v)
//
func (y *Yaml) MustArray(args ...[]interface{}) []interface{} {
	var def []interface{}

	if len(args) > 0 {
		def = args[0]
	}

	a, err := y.Array()

	if err != nil {
		return def
	}

	return a
}

// MustMap guarantees the return of a `map[string]interface{}` (with optional default)
//
// useful when you want to interate over map values in a succinct manner:
//		for k, v := range yaml.Get("dictionary").MustMap() {
//			fmt.Println(k, v)
//		}
func (y *Yaml) MustMap(args ...map[interface{}]interface{}) map[interface{}]interface{} {
	var def map[interface{}]interface{}

	if len(args) > 0 {
		def = args[0]
	}

	m, err := y.Map()

	if err != nil {
		return def
	}

	return m
}

// MustString guarantees the return of a `string` (with optional default)
//
// useful when you explicitly want a `string` in a single value return context:
//     myFunc(yaml.Get("param1").MustString(), yaml.Get("optional_param").MustString("my_default"))
func (y *Yaml) MustString(args ...string) string {
	var def string

	if len(args) > 0 {
		def = args[0]
	}

	s, err := y.String()

	if err != nil {
		return def
	}

	return s
}

// MustStringArray guarantees the return of a `[]string` (with optional default)
//
// useful when you want to interate over array values in a succinct manner:
//		for i, s := range yaml.Get("results").MustStringArray() {
//			fmt.Println(i, s)
//		}
func (y *Yaml) MustStringArray(args ...[]string) []string {
	var def []string

	if len(args) > 0 {
		def = args[0]
	}

	a, err := y.StringArray()

	if err != nil {
		return def
	}

	return a
}

// MustInt guarantees the return of an `int` (with optional default)
//
// useful when you explicitly want an `int` in a single value return context:
//     myFunc(yaml.Get("param1").MustInt(), yaml.Get("optional_param").MustInt(5150))
func (y *Yaml) MustInt(args ...int) int {
	var def int

	if len(args) > 0 {
		def = args[0]
	}

	i, err := y.Int()

	if err != nil {
		return def
	}

	return i
}

// MustFloat guarantees the return of a `float64` (with optional default)
//
// useful when you explicitly want a `float64` in a single value return context:
//     myFunc(yaml.Get("param1").MustFloat64(), yaml.Get("optional_param").MustFloat64(5.150))
func (y *Yaml) MustFloat(args ...float64) float64 {
	var def float64

	if len(args) > 0 {
		def = args[0]
	}

	f, err := y.Float()

	if err != nil {
		return def
	}

	return f
}

// MustBool guarantees the return of a `bool` (with optional default)
//
// useful when you explicitly want a `bool` in a single value return context:
//     myFunc(yaml.Get("param1").MustBool(), yaml.Get("optional_param").MustBool(true))
func (y *Yaml) MustBool(args ...bool) bool {
	var def bool

	if len(args) > 0 {
		def = args[0]
	}

	b, err := y.Bool()

	if err != nil {
		return def
	}

	return b
}

// Get returns a pointer to a new `Yaml` object
// for `key` in it's `map` representation
func (y *Yaml) Get(key string) *Yaml {
	m, err := y.Map()

	if err != nil {
		return &Yaml{nil}
	}

	data, present := m[key]

	if present {
		return &Yaml{data}
	}

	return &Yaml{nil}
}

// GetPath searches for the item as specified by the branch
// without the need to deep dive using Get()'s
func (y *Yaml) GetPath(branch ...string) *Yaml {
	yin := y

	for _, key := range branch {
		yin = yin.Get(key)
	}

	return yin
}

// GetByIndex returns a pointer to a new `Yaml` object
// for `index` in it's `array` representation
func (y *Yaml) GetByIndex(index int) *Yaml {
	a, err := y.Array()

	if err != nil {
		return &Yaml{nil}
	}

	if len(a) > index {
		return &Yaml{a[index]}
	}

	return &Yaml{nil}
}

// CheckGet returns a pointer to a new `Yaml` object and
// a `bool` identifying success or failure
//
// useful for chained operations when success is important:
//    if data, ok := yaml.Get("top_level").CheckGet("inner"); ok {
//        log.Println(data)
//    }
func (y *Yaml) CheckGet(key string) (*Yaml, bool) {
	m, err := y.Map()

	if err != nil {
		return &Yaml{nil}, false
	}

	data, present := m[key]

	if present {
		return &Yaml{data}, true
	}

	return &Yaml{nil}, false
}

// IsArray return true if yaml part is array
func (y *Yaml) IsArray() bool {
	_, err := y.Array()
	return err == nil
}

// IsMap return true if yaml part is map
func (y *Yaml) IsMap() bool {
	_, err := y.Map()
	return err == nil
}

// IsExist return false if object with given key
// is not present in `Yaml` object
func (y *Yaml) IsExist(key string) bool {
	return y.Get(key).data != nil
}

// IsPathExist return false if object with given path
// is not present in `Yaml` object
func (y *Yaml) IsPathExist(path ...string) bool {
	return y.GetPath(path...).data != nil
}

// IsIndexExist return false if object with given index
// is not present in `Yaml` array
func (y *Yaml) IsIndexExist(index int) bool {
	return y.GetByIndex(index).data != nil
}

// GetMapKeys return slice with all map keys
func (y *Yaml) GetMapKeys() ([]string, error) {
	m, err := y.Map()

	if err != nil {
		return nil, err
	}

	var result []string

	for key := range m {
		skey, ok := key.(string)

		if ok {
			result = append(result, skey)
		}
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// convertSlice convert []interface{} slice to string slice
func convertSlice(a []interface{}) ([]string, error) {
	result := make([]string, 0, len(a))

	for _, item := range a {
		if item == nil {
			result = append(result, "")
			continue
		}

		str, ok := item.(string)

		if !ok {
			return nil, ErrStringArrayTypeAssertion
		}

		result = append(result, str)
	}

	return result, nil
}

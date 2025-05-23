package be_reflected_test

import (
	"fmt"
	"reflect"

	"github.com/welltodohamm/be/be_reflected"
	"github.com/welltodohamm/be/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

// TODO: unify tests. Let's make all tests like in `be_math`

var _ = Describe("MatchersReflected", func() {
	Context("AsKind", func() {
		DescribeTable("should match kind", func(actual interface{}, expected reflect.Kind) {
			Expect(actual).To(be_reflected.AsKind(expected), reflect.TypeOf(actual).Kind().String())
		},
			Entry("int", 1, reflect.Int),
			Entry("int8", int8(1), reflect.Int8),
			Entry("int16", int16(1), reflect.Int16),
			Entry("int32", int32(1), reflect.Int32),
			Entry("int64", int64(1), reflect.Int64),
			Entry("uint", uint(1), reflect.Uint),
			Entry("uint8", uint8(1), reflect.Uint8),
			Entry("uint16", uint16(1), reflect.Uint16),
			Entry("uint32", uint32(1), reflect.Uint32),
			Entry("uint64", uint64(1), reflect.Uint64),
			Entry("float32", float32(1), reflect.Float32),
			Entry("float64", float64(1), reflect.Float64),
			Entry("complex64", complex64(1), reflect.Complex64),
			Entry("complex128", complex128(1), reflect.Complex128),
			Entry("string", "1", reflect.String),
			Entry("bool", true, reflect.Bool),
			Entry("chan", make(chan int), reflect.Chan),
			Entry("func", func() {}, reflect.Func),
			Entry("map", map[string]int{}, reflect.Map),
			Entry("ptr", new(int), reflect.Ptr),
			Entry("slice", []int{}, reflect.Slice),
			Entry("struct", struct{}{}, reflect.Struct),

			// todo add test for reflect.Interface
		)

		DescribeTable("should properly fail on matching", func(actual interface{}, expected reflect.Kind) {
			matcher := be_reflected.AsKind(expected)

			success, err := matcher.Match(actual)
			Expect(err).To(Succeed())
			Expect(success).To(BeFalse())

			message := matcher.FailureMessage(actual)
			Expect(message).To(Equal(fmt.Sprintf("Expected\n%s\nto be kind of %s", format.Object(actual, 1), expected.String())))
		},
			Entry("int", 1, reflect.Uint),
			Entry("int8", int8(1), reflect.Uint8),
			Entry("int16", int16(1), reflect.Uint16),

			Entry("int32", int32(1), reflect.Uint32),
			Entry("int64", int64(1), reflect.Uint64),
			Entry("uint", uint(1), reflect.Int),
			Entry("uint8", uint8(1), reflect.Int8),
			Entry("uint16", uint16(1), reflect.Int16),
			Entry("uint32", uint32(1), reflect.Int32),
			Entry("uint64", uint64(1), reflect.Int64),
			Entry("float32", float32(1), reflect.Float64),

			Entry("float64", float64(1), reflect.Float32),
			Entry("complex64", complex64(1), reflect.Complex128),
			Entry("complex128", complex128(1), reflect.Complex64),
			Entry("string as not bool", "1", reflect.Bool),
			Entry("bool as not string", true, reflect.String),
			Entry("func", func() {}, reflect.Chan),
			Entry("chan as not func", make(chan int), reflect.Func),
			Entry("map", map[string]int{}, reflect.Ptr),
			Entry("ptr", new(int), reflect.Map), // <*int | 0xc...>: de-referencad value
		)
	})

	Context("AssignableTo", func() {

		// Named booleans so table is more readable
		const WillPass = true
		const WontPass = false

		type CustomString string

		DescribeTable("should be assignable to a simple type", func(m types.BeMatcher, actual any, result bool) {
			matched, err := m.Match(actual)
			Expect(err).To(Succeed())
			Expect(matched).To(Equal(result))
		},
			Entry("int", be_reflected.AssignableTo[int](), 5, WillPass),
			Entry("uint", be_reflected.AssignableTo[uint8](), uint8(2), WillPass),
			Entry("int vs uint", be_reflected.AssignableTo[uint8](), 2, WontPass),
			Entry("float64", be_reflected.AssignableTo[float64](), 100.0, WillPass),
			Entry("string", be_reflected.AssignableTo[string](), "hello world", WillPass),

			Entry("CustomString", be_reflected.AssignableTo[CustomString](), CustomString("hello world"), WillPass),
			Entry("CustomString as string: no", be_reflected.AssignableTo[string](), CustomString("hello world"), WontPass),
			Entry("string as CustomString: no", be_reflected.AssignableTo[CustomString](), "hello world", WontPass),
		)
	})
})

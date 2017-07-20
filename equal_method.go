package assertions

import (
	"reflect"

	"github.com/smartystreets/logging"
)

type equalityMethodSpecification struct {
	a interface{}
	b interface{}

	aType reflect.Type
	bType reflect.Type

	equalMethod reflect.Value

	log *logging.Logger
}

func newEqualityMethodSpecification(a, b interface{}) *equalityMethodSpecification {
	return &equalityMethodSpecification{
		a:   a,
		b:   b,
		log: logging.Capture(),
	}
}

func (this *equalityMethodSpecification) IsSatisfied() bool {
	if !this.sameType() {
		return false
	}
	if !this.hasEqualMethod() {
		return false
	}
	if !this.equalMethodReceivesSameTypeForComparison() {
		return false
	}
	if !this.equalMethodReturnsBool() {
		return false
	}
	return true
}

func (this *equalityMethodSpecification) sameType() bool {
	this.aType = reflect.TypeOf(this.a)
	if this.aType.Kind() == reflect.Ptr {
		this.aType = this.aType.Elem()
	}
	this.bType = reflect.TypeOf(this.b)
	return this.aType == this.bType
}
func (this *equalityMethodSpecification) hasEqualMethod() bool {
	aInstance := reflect.ValueOf(this.a)
	this.equalMethod = aInstance.MethodByName("Equal")
	return this.equalMethod != reflect.Value{}
}

func (this *equalityMethodSpecification) equalMethodReceivesSameTypeForComparison() bool {
	signature := this.equalMethod.Type()
	return signature.NumIn() == 1 && signature.In(0) == this.aType
}

func (this *equalityMethodSpecification) equalMethodReturnsBool() bool {
	signature := this.equalMethod.Type()
	return signature.NumOut() == 1 && signature.Out(0) == reflect.TypeOf(true)
}

func (this *equalityMethodSpecification) AreEqual() bool {
	argument := reflect.ValueOf(this.b)
	argumentList := []reflect.Value{argument}
	result := this.equalMethod.Call(argumentList)
	return result[0].Bool()
}

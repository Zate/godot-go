package visualscript

import (
	"log"
	"reflect"

	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
Undocumented
*/
type VisualScriptOperator struct {
	VisualScriptNode
}

func (o *VisualScriptOperator) BaseClass() string {
	return "VisualScriptOperator"
}

/*
   Undocumented
*/
func (o *VisualScriptOperator) GetOperator() gdnative.Int {
	log.Println("Calling VisualScriptOperator.GetOperator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_operator", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptOperator) GetTyped() gdnative.Int {
	log.Println("Calling VisualScriptOperator.GetTyped()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_typed", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptOperator) SetOperator(op gdnative.Int) {
	log.Println("Calling VisualScriptOperator.SetOperator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(op)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_operator", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VisualScriptOperator) SetTyped(aType gdnative.Int) {
	log.Println("Calling VisualScriptOperator.SetTyped()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(aType)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_typed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptOperatorImplementer is an interface for VisualScriptOperator objects.
*/
type VisualScriptOperatorImplementer interface {
	Class
}
package class

import (
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

//func NewVisualScriptIteratorFromPointer(ptr gdnative.Pointer) VisualScriptIterator {
func NewVisualScriptIteratorFromPointer(ptr gdnative.Pointer) VisualScriptIterator {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptIterator{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptIterator struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptIterator) BaseClass() string {
	return "VisualScriptIterator"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptIterator) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *VisualScriptIterator) GetBaseObject() gdnative.Object {
	return o.owner
}
package capsuleshape

import (
	"log"
	"reflect"
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
Capsule shape for collisions.
*/
type CapsuleShape struct {
	Shape
}

func (o *CapsuleShape) BaseClass() string {
	return "CapsuleShape"
}

/*
   Undocumented
*/
func (o *CapsuleShape) GetHeight() gdnative.Float {
	log.Println("Calling CapsuleShape.GetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_height", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *CapsuleShape) GetRadius() gdnative.Float {
	log.Println("Calling CapsuleShape.GetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_radius", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *CapsuleShape) SetHeight(height gdnative.Float) {
	log.Println("Calling CapsuleShape.SetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(height)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_height", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *CapsuleShape) SetRadius(radius gdnative.Float) {
	log.Println("Calling CapsuleShape.SetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   CapsuleShapeImplementer is an interface for CapsuleShape objects.
*/
type CapsuleShapeImplementer interface {
	Class
}
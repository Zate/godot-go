package godot

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

//func NewRangeFromPointer(ptr gdnative.Pointer) Range {
func newRangeFromPointer(ptr gdnative.Pointer) Range {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Range{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Range is a base class for [Control] nodes that change a floating point [i]value[/i] between a [i]minimum[/i] and a [i]maximum[/i], using [i]step[/i] and [i]page[/i], for example a [ScrollBar].
*/
type Range struct {
	Control
	owner gdnative.Object
}

func (o *Range) BaseClass() string {
	return "Range"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Range) GetAsRatio() gdnative.Real {
	//log.Println("Calling Range.GetAsRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "get_as_ratio")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Range) GetMax() gdnative.Real {
	//log.Println("Calling Range.GetMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "get_max")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Range) GetMin() gdnative.Real {
	//log.Println("Calling Range.GetMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "get_min")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Range) GetPage() gdnative.Real {
	//log.Println("Calling Range.GetPage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "get_page")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Range) GetStep() gdnative.Real {
	//log.Println("Calling Range.GetStep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "get_step")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Range) GetValue() gdnative.Real {
	//log.Println("Calling Range.GetValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "get_value")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Range) IsRatioExp() gdnative.Bool {
	//log.Println("Calling Range.IsRatioExp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "is_ratio_exp")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Range) IsUsingRoundedValues() gdnative.Bool {
	//log.Println("Calling Range.IsUsingRoundedValues()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "is_using_rounded_values")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false value float}], Returns: void
*/
func (o *Range) SetAsRatio(value gdnative.Real) {
	//log.Println("Calling Range.SetAsRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_as_ratio")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Range) SetExpRatio(enabled gdnative.Bool) {
	//log.Println("Calling Range.SetExpRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_exp_ratio")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false maximum float}], Returns: void
*/
func (o *Range) SetMax(maximum gdnative.Real) {
	//log.Println("Calling Range.SetMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(maximum)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_max")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false minimum float}], Returns: void
*/
func (o *Range) SetMin(minimum gdnative.Real) {
	//log.Println("Calling Range.SetMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(minimum)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_min")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pagesize float}], Returns: void
*/
func (o *Range) SetPage(pagesize gdnative.Real) {
	//log.Println("Calling Range.SetPage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(pagesize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_page")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false step float}], Returns: void
*/
func (o *Range) SetStep(step gdnative.Real) {
	//log.Println("Calling Range.SetStep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(step)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_step")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Range) SetUseRoundedValues(enabled gdnative.Bool) {
	//log.Println("Calling Range.SetUseRoundedValues()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_use_rounded_values")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false value float}], Returns: void
*/
func (o *Range) SetValue(value gdnative.Real) {
	//log.Println("Calling Range.SetValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "set_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Binds two Ranges together along with any Ranges previously grouped with either of them. When any of Range's member variables change, it will share the new value with all other Ranges in its group.
	Args: [{ false with Object}], Returns: void
*/
func (o *Range) Share(with ObjectImplementer) {
	//log.Println("Calling Range.Share()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(with.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "share")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stop Range from sharing its member variables with any other Range.
	Args: [], Returns: void
*/
func (o *Range) Unshare() {
	//log.Println("Calling Range.Unshare()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Range", "unshare")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RangeImplementer is an interface that implements the methods
// of the Range class.
type RangeImplementer interface {
	ControlImplementer
	GetAsRatio() gdnative.Real
	GetMax() gdnative.Real
	GetMin() gdnative.Real
	GetPage() gdnative.Real
	GetStep() gdnative.Real
	GetValue() gdnative.Real
	IsRatioExp() gdnative.Bool
	IsUsingRoundedValues() gdnative.Bool
	SetAsRatio(value gdnative.Real)
	SetExpRatio(enabled gdnative.Bool)
	SetMax(maximum gdnative.Real)
	SetMin(minimum gdnative.Real)
	SetPage(pagesize gdnative.Real)
	SetStep(step gdnative.Real)
	SetUseRoundedValues(enabled gdnative.Bool)
	SetValue(value gdnative.Real)
	Share(with ObjectImplementer)
	Unshare()
}

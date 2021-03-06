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

//func NewNativeScriptFromPointer(ptr gdnative.Pointer) NativeScript {
func newNativeScriptFromPointer(ptr gdnative.Pointer) NativeScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := NativeScript{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type NativeScript struct {
	Script
	owner gdnative.Object
}

func (o *NativeScript) BaseClass() string {
	return "NativeScript"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *NativeScript) GetClassName() gdnative.String {
	//log.Println("Calling NativeScript.GetClassName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NativeScript", "get_class_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: GDNativeLibrary
*/
func (o *NativeScript) GetLibrary() GDNativeLibraryImplementer {
	//log.Println("Calling NativeScript.GetLibrary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NativeScript", "get_library")

	// Call the parent method.
	// GDNativeLibrary
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newGDNativeLibraryFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(GDNativeLibraryImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "GDNativeLibrary" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(GDNativeLibraryImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Object
*/
func (o *NativeScript) New() ObjectImplementer {
	//log.Println("Calling NativeScript.New()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NativeScript", "new")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false class_name String}], Returns: void
*/
func (o *NativeScript) SetClassName(className gdnative.String) {
	//log.Println("Calling NativeScript.SetClassName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(className)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NativeScript", "set_class_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false library GDNativeLibrary}], Returns: void
*/
func (o *NativeScript) SetLibrary(library GDNativeLibraryImplementer) {
	//log.Println("Calling NativeScript.SetLibrary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(library.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NativeScript", "set_library")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// NativeScriptImplementer is an interface that implements the methods
// of the NativeScript class.
type NativeScriptImplementer interface {
	ScriptImplementer
	GetClassName() gdnative.String
	GetLibrary() GDNativeLibraryImplementer
	New() ObjectImplementer
	SetClassName(className gdnative.String)
	SetLibrary(library GDNativeLibraryImplementer)
}

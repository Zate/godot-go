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

//func NewWindowDialogFromPointer(ptr gdnative.Pointer) WindowDialog {
func newWindowDialogFromPointer(ptr gdnative.Pointer) WindowDialog {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WindowDialog{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Windowdialog is the base class for all window-based dialogs. It's a by-default toplevel [Control] that draws a window decoration and allows motion and resizing.
*/
type WindowDialog struct {
	Popup
	owner gdnative.Object
}

func (o *WindowDialog) BaseClass() string {
	return "WindowDialog"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *WindowDialog) X_Closed() {
	//log.Println("Calling WindowDialog.X_Closed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WindowDialog", "_closed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *WindowDialog) X_GuiInput(arg0 InputEvent) {
	//log.Println("Calling WindowDialog.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WindowDialog", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Return the close [TextureButton].
	Args: [], Returns: TextureButton
*/
func (o *WindowDialog) GetCloseButton() TextureButtonImplementer {
	//log.Println("Calling WindowDialog.GetCloseButton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WindowDialog", "get_close_button")

	// Call the parent method.
	// TextureButton
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureButtonFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureButtonImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "TextureButton" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureButtonImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *WindowDialog) GetResizable() gdnative.Bool {
	//log.Println("Calling WindowDialog.GetResizable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WindowDialog", "get_resizable")

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
	Args: [], Returns: String
*/
func (o *WindowDialog) GetTitle() gdnative.String {
	//log.Println("Calling WindowDialog.GetTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WindowDialog", "get_title")

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
	Args: [{ false resizable bool}], Returns: void
*/
func (o *WindowDialog) SetResizable(resizable gdnative.Bool) {
	//log.Println("Calling WindowDialog.SetResizable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(resizable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WindowDialog", "set_resizable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false title String}], Returns: void
*/
func (o *WindowDialog) SetTitle(title gdnative.String) {
	//log.Println("Calling WindowDialog.SetTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(title)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WindowDialog", "set_title")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// WindowDialogImplementer is an interface that implements the methods
// of the WindowDialog class.
type WindowDialogImplementer interface {
	PopupImplementer
	X_Closed()
	GetCloseButton() TextureButtonImplementer
	GetResizable() gdnative.Bool
	GetTitle() gdnative.String
	SetResizable(resizable gdnative.Bool)
	SetTitle(title gdnative.String)
}

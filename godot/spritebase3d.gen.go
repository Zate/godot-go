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

// SpriteBase3DAlphaCutMode is an enum for AlphaCutMode values.
type SpriteBase3DAlphaCutMode int

const (
	SpriteBase3DAlphaCutDisabled      SpriteBase3DAlphaCutMode = 0
	SpriteBase3DAlphaCutDiscard       SpriteBase3DAlphaCutMode = 1
	SpriteBase3DAlphaCutOpaquePrepass SpriteBase3DAlphaCutMode = 2
)

// SpriteBase3DDrawFlags is an enum for DrawFlags values.
type SpriteBase3DDrawFlags int

const (
	SpriteBase3DFlagDoubleSided SpriteBase3DDrawFlags = 2
	SpriteBase3DFlagMax         SpriteBase3DDrawFlags = 3
	SpriteBase3DFlagShaded      SpriteBase3DDrawFlags = 1
	SpriteBase3DFlagTransparent SpriteBase3DDrawFlags = 0
)

//func NewSpriteBase3DFromPointer(ptr gdnative.Pointer) SpriteBase3D {
func newSpriteBase3DFromPointer(ptr gdnative.Pointer) SpriteBase3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SpriteBase3D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A node that displays 2D texture information in a 3D environment.
*/
type SpriteBase3D struct {
	GeometryInstance
	owner gdnative.Object
}

func (o *SpriteBase3D) BaseClass() string {
	return "SpriteBase3D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *SpriteBase3D) X_ImUpdate() {
	//log.Println("Calling SpriteBase3D.X_ImUpdate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "_im_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *SpriteBase3D) X_QueueUpdate() {
	//log.Println("Calling SpriteBase3D.X_QueueUpdate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "_queue_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.SpriteBase3D::AlphaCutMode
*/
func (o *SpriteBase3D) GetAlphaCutMode() SpriteBase3DAlphaCutMode {
	//log.Println("Calling SpriteBase3D.GetAlphaCutMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_alpha_cut_mode")

	// Call the parent method.
	// enum.SpriteBase3D::AlphaCutMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return SpriteBase3DAlphaCutMode(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.Vector3::Axis
*/
func (o *SpriteBase3D) GetAxis() gdnative.Vector3Axis {
	//log.Println("Calling SpriteBase3D.GetAxis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_axis")

	// Call the parent method.
	// enum.Vector3::Axis
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Vector3Axis(ret)
}

/*
        Undocumented
	Args: [{ false flag int}], Returns: bool
*/
func (o *SpriteBase3D) GetDrawFlag(flag gdnative.Int) gdnative.Bool {
	//log.Println("Calling SpriteBase3D.GetDrawFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_draw_flag")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: Rect2
*/
func (o *SpriteBase3D) GetItemRect() gdnative.Rect2 {
	//log.Println("Calling SpriteBase3D.GetItemRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_item_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *SpriteBase3D) GetModulate() gdnative.Color {
	//log.Println("Calling SpriteBase3D.GetModulate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_modulate")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *SpriteBase3D) GetOffset() gdnative.Vector2 {
	//log.Println("Calling SpriteBase3D.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *SpriteBase3D) GetOpacity() gdnative.Real {
	//log.Println("Calling SpriteBase3D.GetOpacity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_opacity")

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
func (o *SpriteBase3D) GetPixelSize() gdnative.Real {
	//log.Println("Calling SpriteBase3D.GetPixelSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "get_pixel_size")

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
func (o *SpriteBase3D) IsCentered() gdnative.Bool {
	//log.Println("Calling SpriteBase3D.IsCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "is_centered")

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
func (o *SpriteBase3D) IsFlippedH() gdnative.Bool {
	//log.Println("Calling SpriteBase3D.IsFlippedH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "is_flipped_h")

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
func (o *SpriteBase3D) IsFlippedV() gdnative.Bool {
	//log.Println("Calling SpriteBase3D.IsFlippedV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "is_flipped_v")

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
	Args: [{ false mode int}], Returns: void
*/
func (o *SpriteBase3D) SetAlphaCutMode(mode gdnative.Int) {
	//log.Println("Calling SpriteBase3D.SetAlphaCutMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_alpha_cut_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false axis int}], Returns: void
*/
func (o *SpriteBase3D) SetAxis(axis gdnative.Int) {
	//log.Println("Calling SpriteBase3D.SetAxis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(axis)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_axis")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false centered bool}], Returns: void
*/
func (o *SpriteBase3D) SetCentered(centered gdnative.Bool) {
	//log.Println("Calling SpriteBase3D.SetCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(centered)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_centered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flag int} { false enabled bool}], Returns: void
*/
func (o *SpriteBase3D) SetDrawFlag(flag gdnative.Int, enabled gdnative.Bool) {
	//log.Println("Calling SpriteBase3D.SetDrawFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_draw_flag")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flip_h bool}], Returns: void
*/
func (o *SpriteBase3D) SetFlipH(flipH gdnative.Bool) {
	//log.Println("Calling SpriteBase3D.SetFlipH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(flipH)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_flip_h")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flip_v bool}], Returns: void
*/
func (o *SpriteBase3D) SetFlipV(flipV gdnative.Bool) {
	//log.Println("Calling SpriteBase3D.SetFlipV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(flipV)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_flip_v")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false modulate Color}], Returns: void
*/
func (o *SpriteBase3D) SetModulate(modulate gdnative.Color) {
	//log.Println("Calling SpriteBase3D.SetModulate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(modulate)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_modulate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *SpriteBase3D) SetOffset(offset gdnative.Vector2) {
	//log.Println("Calling SpriteBase3D.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false opacity float}], Returns: void
*/
func (o *SpriteBase3D) SetOpacity(opacity gdnative.Real) {
	//log.Println("Calling SpriteBase3D.SetOpacity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(opacity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_opacity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pixel_size float}], Returns: void
*/
func (o *SpriteBase3D) SetPixelSize(pixelSize gdnative.Real) {
	//log.Println("Calling SpriteBase3D.SetPixelSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(pixelSize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpriteBase3D", "set_pixel_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SpriteBase3DImplementer is an interface that implements the methods
// of the SpriteBase3D class.
type SpriteBase3DImplementer interface {
	GeometryInstanceImplementer
	X_ImUpdate()
	X_QueueUpdate()
	GetDrawFlag(flag gdnative.Int) gdnative.Bool
	GetItemRect() gdnative.Rect2
	GetModulate() gdnative.Color
	GetOffset() gdnative.Vector2
	GetOpacity() gdnative.Real
	GetPixelSize() gdnative.Real
	IsCentered() gdnative.Bool
	IsFlippedH() gdnative.Bool
	IsFlippedV() gdnative.Bool
	SetAlphaCutMode(mode gdnative.Int)
	SetAxis(axis gdnative.Int)
	SetCentered(centered gdnative.Bool)
	SetDrawFlag(flag gdnative.Int, enabled gdnative.Bool)
	SetFlipH(flipH gdnative.Bool)
	SetFlipV(flipV gdnative.Bool)
	SetModulate(modulate gdnative.Color)
	SetOffset(offset gdnative.Vector2)
	SetOpacity(opacity gdnative.Real)
	SetPixelSize(pixelSize gdnative.Real)
}

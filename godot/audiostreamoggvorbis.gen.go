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

//func NewAudioStreamOGGVorbisFromPointer(ptr gdnative.Pointer) AudioStreamOGGVorbis {
func newAudioStreamOGGVorbisFromPointer(ptr gdnative.Pointer) AudioStreamOGGVorbis {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamOGGVorbis{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type AudioStreamOGGVorbis struct {
	AudioStream
	owner gdnative.Object
}

func (o *AudioStreamOGGVorbis) BaseClass() string {
	return "AudioStreamOGGVorbis"
}

/*
        Undocumented
	Args: [], Returns: PoolByteArray
*/
func (o *AudioStreamOGGVorbis) X_GetData() gdnative.PoolByteArray {
	//log.Println("Calling AudioStreamOGGVorbis.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamOGGVorbis", "_get_data")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false data PoolByteArray}], Returns: void
*/
func (o *AudioStreamOGGVorbis) X_SetData(data gdnative.PoolByteArray) {
	//log.Println("Calling AudioStreamOGGVorbis.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamOGGVorbis", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamOGGVorbis) GetLoopOffset() gdnative.Real {
	//log.Println("Calling AudioStreamOGGVorbis.GetLoopOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamOGGVorbis", "get_loop_offset")

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
func (o *AudioStreamOGGVorbis) HasLoop() gdnative.Bool {
	//log.Println("Calling AudioStreamOGGVorbis.HasLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamOGGVorbis", "has_loop")

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
	Args: [{ false enable bool}], Returns: void
*/
func (o *AudioStreamOGGVorbis) SetLoop(enable gdnative.Bool) {
	//log.Println("Calling AudioStreamOGGVorbis.SetLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamOGGVorbis", "set_loop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false seconds float}], Returns: void
*/
func (o *AudioStreamOGGVorbis) SetLoopOffset(seconds gdnative.Real) {
	//log.Println("Calling AudioStreamOGGVorbis.SetLoopOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(seconds)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamOGGVorbis", "set_loop_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioStreamOGGVorbisImplementer is an interface that implements the methods
// of the AudioStreamOGGVorbis class.
type AudioStreamOGGVorbisImplementer interface {
	AudioStreamImplementer
	X_GetData() gdnative.PoolByteArray
	X_SetData(data gdnative.PoolByteArray)
	GetLoopOffset() gdnative.Real
	HasLoop() gdnative.Bool
	SetLoop(enable gdnative.Bool)
	SetLoopOffset(seconds gdnative.Real)
}

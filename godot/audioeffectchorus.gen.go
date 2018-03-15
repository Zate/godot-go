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

//func NewAudioEffectChorusFromPointer(ptr gdnative.Pointer) AudioEffectChorus {
func newAudioEffectChorusFromPointer(ptr gdnative.Pointer) AudioEffectChorus {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectChorus{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Adds a chorus audio effect. The effect applies a filter with voices to duplicate the audio source and manipulate it through the filter.
*/
type AudioEffectChorus struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectChorus) BaseClass() string {
	return "AudioEffectChorus"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectChorus) GetDry() gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetDry()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_dry")

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
	Args: [], Returns: int
*/
func (o *AudioEffectChorus) GetVoiceCount() gdnative.Int {
	//log.Println("Calling AudioEffectChorus.GetVoiceCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_voice_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false voice_idx int}], Returns: float
*/
func (o *AudioEffectChorus) GetVoiceCutoffHz(voiceIdx gdnative.Int) gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetVoiceCutoffHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_voice_cutoff_hz")

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
	Args: [{ false voice_idx int}], Returns: float
*/
func (o *AudioEffectChorus) GetVoiceDelayMs(voiceIdx gdnative.Int) gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetVoiceDelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_voice_delay_ms")

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
	Args: [{ false voice_idx int}], Returns: float
*/
func (o *AudioEffectChorus) GetVoiceDepthMs(voiceIdx gdnative.Int) gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetVoiceDepthMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_voice_depth_ms")

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
	Args: [{ false voice_idx int}], Returns: float
*/
func (o *AudioEffectChorus) GetVoiceLevelDb(voiceIdx gdnative.Int) gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetVoiceLevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_voice_level_db")

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
	Args: [{ false voice_idx int}], Returns: float
*/
func (o *AudioEffectChorus) GetVoicePan(voiceIdx gdnative.Int) gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetVoicePan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_voice_pan")

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
	Args: [{ false voice_idx int}], Returns: float
*/
func (o *AudioEffectChorus) GetVoiceRateHz(voiceIdx gdnative.Int) gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetVoiceRateHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_voice_rate_hz")

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
func (o *AudioEffectChorus) GetWet() gdnative.Real {
	//log.Println("Calling AudioEffectChorus.GetWet()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "get_wet")

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
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectChorus) SetDry(amount gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetDry()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_dry")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false voices int}], Returns: void
*/
func (o *AudioEffectChorus) SetVoiceCount(voices gdnative.Int) {
	//log.Println("Calling AudioEffectChorus.SetVoiceCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(voices)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_voice_count")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false voice_idx int} { false cutoff_hz float}], Returns: void
*/
func (o *AudioEffectChorus) SetVoiceCutoffHz(voiceIdx gdnative.Int, cutoffHz gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetVoiceCutoffHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)
	ptrArguments[1] = gdnative.NewPointerFromReal(cutoffHz)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_voice_cutoff_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false voice_idx int} { false delay_ms float}], Returns: void
*/
func (o *AudioEffectChorus) SetVoiceDelayMs(voiceIdx gdnative.Int, delayMs gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetVoiceDelayMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)
	ptrArguments[1] = gdnative.NewPointerFromReal(delayMs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_voice_delay_ms")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false voice_idx int} { false depth_ms float}], Returns: void
*/
func (o *AudioEffectChorus) SetVoiceDepthMs(voiceIdx gdnative.Int, depthMs gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetVoiceDepthMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)
	ptrArguments[1] = gdnative.NewPointerFromReal(depthMs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_voice_depth_ms")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false voice_idx int} { false level_db float}], Returns: void
*/
func (o *AudioEffectChorus) SetVoiceLevelDb(voiceIdx gdnative.Int, levelDb gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetVoiceLevelDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)
	ptrArguments[1] = gdnative.NewPointerFromReal(levelDb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_voice_level_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false voice_idx int} { false pan float}], Returns: void
*/
func (o *AudioEffectChorus) SetVoicePan(voiceIdx gdnative.Int, pan gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetVoicePan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)
	ptrArguments[1] = gdnative.NewPointerFromReal(pan)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_voice_pan")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false voice_idx int} { false rate_hz float}], Returns: void
*/
func (o *AudioEffectChorus) SetVoiceRateHz(voiceIdx gdnative.Int, rateHz gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetVoiceRateHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(voiceIdx)
	ptrArguments[1] = gdnative.NewPointerFromReal(rateHz)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_voice_rate_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectChorus) SetWet(amount gdnative.Real) {
	//log.Println("Calling AudioEffectChorus.SetWet()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectChorus", "set_wet")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectChorusImplementer is an interface that implements the methods
// of the AudioEffectChorus class.
type AudioEffectChorusImplementer interface {
	AudioEffectImplementer
	GetDry() gdnative.Real
	GetVoiceCount() gdnative.Int
	GetVoiceCutoffHz(voiceIdx gdnative.Int) gdnative.Real
	GetVoiceDelayMs(voiceIdx gdnative.Int) gdnative.Real
	GetVoiceDepthMs(voiceIdx gdnative.Int) gdnative.Real
	GetVoiceLevelDb(voiceIdx gdnative.Int) gdnative.Real
	GetVoicePan(voiceIdx gdnative.Int) gdnative.Real
	GetVoiceRateHz(voiceIdx gdnative.Int) gdnative.Real
	GetWet() gdnative.Real
	SetDry(amount gdnative.Real)
	SetVoiceCount(voices gdnative.Int)
	SetVoiceCutoffHz(voiceIdx gdnative.Int, cutoffHz gdnative.Real)
	SetVoiceDelayMs(voiceIdx gdnative.Int, delayMs gdnative.Real)
	SetVoiceDepthMs(voiceIdx gdnative.Int, depthMs gdnative.Real)
	SetVoiceLevelDb(voiceIdx gdnative.Int, levelDb gdnative.Real)
	SetVoicePan(voiceIdx gdnative.Int, pan gdnative.Real)
	SetVoiceRateHz(voiceIdx gdnative.Int, rateHz gdnative.Real)
	SetWet(amount gdnative.Real)
}

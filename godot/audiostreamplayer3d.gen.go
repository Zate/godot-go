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

//func NewAudioStreamPlayer3DFromPointer(ptr gdnative.Pointer) AudioStreamPlayer3D {
func newAudioStreamPlayer3DFromPointer(ptr gdnative.Pointer) AudioStreamPlayer3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamPlayer3D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Plays a sound effect with directed sound effects, dampens with distance if needed, generates effect of hearable position in space.
*/
type AudioStreamPlayer3D struct {
	Spatial
	owner gdnative.Object
}

func (o *AudioStreamPlayer3D) BaseClass() string {
	return "AudioStreamPlayer3D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AudioStreamPlayer3D) X_BusLayoutChanged() {
	//log.Println("Calling AudioStreamPlayer3D.X_BusLayoutChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "_bus_layout_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer3D) X_IsActive() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer3D.X_IsActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "_is_active")

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
func (o *AudioStreamPlayer3D) X_SetPlaying(enable gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer3D.X_SetPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "_set_playing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *AudioStreamPlayer3D) GetAreaMask() gdnative.Int {
	//log.Println("Calling AudioStreamPlayer3D.GetAreaMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_area_mask")

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
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetAttenuationFilterCutoffHz() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetAttenuationFilterCutoffHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_attenuation_filter_cutoff_hz")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetAttenuationFilterDb() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetAttenuationFilterDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_attenuation_filter_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.AudioStreamPlayer3D::AttenuationModel
*/

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AudioStreamPlayer3D) GetBus() gdnative.String {
	//log.Println("Calling AudioStreamPlayer3D.GetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_bus")

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
	Args: [], Returns: enum.AudioStreamPlayer3D::DopplerTracking
*/

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetEmissionAngle() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetEmissionAngle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_emission_angle")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetEmissionAngleFilterAttenuationDb() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetEmissionAngleFilterAttenuationDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_emission_angle_filter_attenuation_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetMaxDb() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetMaxDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_max_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetMaxDistance() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_max_distance")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.AudioStreamPlayer3D::OutOfRangeMode
*/

/*
        Returns the position in the [AudioStream].
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetPlaybackPosition() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetPlaybackPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_playback_position")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: AudioStream
*/
func (o *AudioStreamPlayer3D) GetStream() AudioStreamImplementer {
	//log.Println("Calling AudioStreamPlayer3D.GetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_stream")

	// Call the parent method.
	// AudioStream
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAudioStreamFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AudioStreamImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AudioStream" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AudioStreamImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetUnitDb() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetUnitDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_unit_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer3D) GetUnitSize() gdnative.Float {
	//log.Println("Calling AudioStreamPlayer3D.GetUnitSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "get_unit_size")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer3D) IsAutoplayEnabled() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer3D.IsAutoplayEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "is_autoplay_enabled")

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
func (o *AudioStreamPlayer3D) IsEmissionAngleEnabled() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer3D.IsEmissionAngleEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "is_emission_angle_enabled")

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
func (o *AudioStreamPlayer3D) IsPlaying() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer3D.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "is_playing")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Plays the audio from the given position 'from_position', in seconds.
	Args: [{0 true from_position float}], Returns: void
*/
func (o *AudioStreamPlayer3D) Play(fromPosition gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(fromPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position from which audio will be played, in seconds.
	Args: [{ false to_position float}], Returns: void
*/
func (o *AudioStreamPlayer3D) Seek(toPosition gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.Seek()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(toPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "seek")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetAreaMask(mask gdnative.Int) {
	//log.Println("Calling AudioStreamPlayer3D.SetAreaMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_area_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetAttenuationFilterCutoffHz(degrees gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetAttenuationFilterCutoffHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_attenuation_filter_cutoff_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false db float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetAttenuationFilterDb(db gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetAttenuationFilterDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(db)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_attenuation_filter_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false model int}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetAttenuationModel(model gdnative.Int) {
	//log.Println("Calling AudioStreamPlayer3D.SetAttenuationModel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(model)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_attenuation_model")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetAutoplay(enable gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer3D.SetAutoplay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_autoplay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bus String}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetBus(bus gdnative.String) {
	//log.Println("Calling AudioStreamPlayer3D.SetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(bus)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_bus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetDopplerTracking(mode gdnative.Int) {
	//log.Println("Calling AudioStreamPlayer3D.SetDopplerTracking()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_doppler_tracking")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetEmissionAngle(degrees gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetEmissionAngle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_emission_angle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetEmissionAngleEnabled(enabled gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer3D.SetEmissionAngleEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_emission_angle_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false db float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetEmissionAngleFilterAttenuationDb(db gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetEmissionAngleFilterAttenuationDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(db)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_emission_angle_filter_attenuation_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false max_db float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetMaxDb(maxDb gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetMaxDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(maxDb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_max_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false metres float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetMaxDistance(metres gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(metres)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_max_distance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetOutOfRangeMode(mode gdnative.Int) {
	//log.Println("Calling AudioStreamPlayer3D.SetOutOfRangeMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_out_of_range_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stream AudioStream}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetStream(stream AudioStream) {
	//log.Println("Calling AudioStreamPlayer3D.SetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(stream.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_stream")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false unit_db float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetUnitDb(unitDb gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetUnitDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(unitDb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_unit_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false unit_size float}], Returns: void
*/
func (o *AudioStreamPlayer3D) SetUnitSize(unitSize gdnative.Float) {
	//log.Println("Calling AudioStreamPlayer3D.SetUnitSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(unitSize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "set_unit_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stops the audio.
	Args: [], Returns: void
*/
func (o *AudioStreamPlayer3D) Stop() {
	//log.Println("Calling AudioStreamPlayer3D.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer3D", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioStreamPlayer3DImplementer is an interface that implements the methods
// of the AudioStreamPlayer3D class.
type AudioStreamPlayer3DImplementer interface {
	SpatialImplementer
	X_BusLayoutChanged()
	X_IsActive() gdnative.Bool
	X_SetPlaying(enable gdnative.Bool)
	GetAreaMask() gdnative.Int
	GetAttenuationFilterCutoffHz() gdnative.Float
	GetAttenuationFilterDb() gdnative.Float
	GetBus() gdnative.String
	GetEmissionAngle() gdnative.Float
	GetEmissionAngleFilterAttenuationDb() gdnative.Float
	GetMaxDb() gdnative.Float
	GetMaxDistance() gdnative.Float
	GetPlaybackPosition() gdnative.Float
	GetStream() AudioStreamImplementer
	GetUnitDb() gdnative.Float
	GetUnitSize() gdnative.Float
	IsAutoplayEnabled() gdnative.Bool
	IsEmissionAngleEnabled() gdnative.Bool
	IsPlaying() gdnative.Bool
	Play(fromPosition gdnative.Float)
	Seek(toPosition gdnative.Float)
	SetAreaMask(mask gdnative.Int)
	SetAttenuationFilterCutoffHz(degrees gdnative.Float)
	SetAttenuationFilterDb(db gdnative.Float)
	SetAttenuationModel(model gdnative.Int)
	SetAutoplay(enable gdnative.Bool)
	SetBus(bus gdnative.String)
	SetDopplerTracking(mode gdnative.Int)
	SetEmissionAngle(degrees gdnative.Float)
	SetEmissionAngleEnabled(enabled gdnative.Bool)
	SetEmissionAngleFilterAttenuationDb(db gdnative.Float)
	SetMaxDb(maxDb gdnative.Float)
	SetMaxDistance(metres gdnative.Float)
	SetOutOfRangeMode(mode gdnative.Int)
	SetStream(stream AudioStream)
	SetUnitDb(unitDb gdnative.Float)
	SetUnitSize(unitSize gdnative.Float)
	Stop()
}

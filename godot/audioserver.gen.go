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

//func NewaudioServerFromPointer(ptr gdnative.Pointer) audioServer {
func newAudioServerFromPointer(ptr gdnative.Pointer) audioServer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := audioServer{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonAudioServer() *audioServer {
	return &audioServer{}
}

/*
   AudioServer is a low level server interface for audio access. It is in charge of creating sample data (playable audio) as well as its playback via a voice interface.
*/
var AudioServer = newSingletonAudioServer()

/*
AudioServer is a low level server interface for audio access. It is in charge of creating sample data (playable audio) as well as its playback via a voice interface.
*/
type audioServer struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *audioServer) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("AudioServer")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *audioServer) BaseClass() string {
	return "AudioServer"
}

/*
        Adds a bus at [code]at_position[/code].
	Args: [{-1 true at_position int}], Returns: void
*/
func (o *audioServer) AddBus(atPosition gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.AddBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(atPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "add_bus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds an [AudioEffect] effect to the bus [code]bus_idx[/code] at [code]at_position[/code].
	Args: [{ false bus_idx int} { false effect AudioEffect} {-1 true at_position int}], Returns: void
*/
func (o *audioServer) AddBusEffect(busIdx gdnative.Int, effect AudioEffect, atPosition gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.AddBusEffect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromObject(effect.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(atPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "add_bus_effect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Generates an [AudioBusLayout] using the available busses and effects.
	Args: [], Returns: AudioBusLayout
*/
func (o *audioServer) GenerateBusLayout() AudioBusLayoutImplementer {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GenerateBusLayout()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "generate_bus_layout")

	// Call the parent method.
	// AudioBusLayout
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAudioBusLayoutFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AudioBusLayoutImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AudioBusLayout" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AudioBusLayoutImplementer)
	}

	return &ret
}

/*
        Returns the number of available busses.
	Args: [], Returns: int
*/
func (o *audioServer) GetBusCount() gdnative.Int {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the [AudioEffect] at position [code]effect_idx[/code] in bus [code]bus_idx[/code].
	Args: [{ false bus_idx int} { false effect_idx int}], Returns: AudioEffect
*/
func (o *audioServer) GetBusEffect(busIdx gdnative.Int, effectIdx gdnative.Int) AudioEffectImplementer {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusEffect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(effectIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_effect")

	// Call the parent method.
	// AudioEffect
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAudioEffectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AudioEffectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AudioEffect" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AudioEffectImplementer)
	}

	return &ret
}

/*
        Returns the number of effects on the bus at [code]bus_idx[/code].
	Args: [{ false bus_idx int}], Returns: int
*/
func (o *audioServer) GetBusEffectCount(busIdx gdnative.Int) gdnative.Int {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusEffectCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_effect_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the index of the bus with the name [code]bus_name[/code].
	Args: [{ false bus_name String}], Returns: int
*/
func (o *audioServer) GetBusIndex(busName gdnative.String) gdnative.Int {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(busName)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the name of the bus with the index [code]bus_idx[/code].
	Args: [{ false bus_idx int}], Returns: String
*/
func (o *audioServer) GetBusName(busIdx gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the peak volume of the left speaker at bus index [code]bus_idx[/code] and channel index [code]channel[/code].
	Args: [{ false bus_idx int} { false channel int}], Returns: float
*/
func (o *audioServer) GetBusPeakVolumeLeftDb(busIdx gdnative.Int, channel gdnative.Int) gdnative.Float {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusPeakVolumeLeftDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(channel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_peak_volume_left_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns the peak volume of the right speaker at bus index [code]bus_idx[/code] and channel index [code]channel[/code].
	Args: [{ false bus_idx int} { false channel int}], Returns: float
*/
func (o *audioServer) GetBusPeakVolumeRightDb(busIdx gdnative.Int, channel gdnative.Int) gdnative.Float {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusPeakVolumeRightDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(channel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_peak_volume_right_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns the name of the bus that the bus at index [code]bus_idx[/code] sends to.
	Args: [{ false bus_idx int}], Returns: String
*/
func (o *audioServer) GetBusSend(busIdx gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusSend()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_send")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the volume of the bus at index [code]bus_idx[/code] in dB.
	Args: [{ false bus_idx int}], Returns: float
*/
func (o *audioServer) GetBusVolumeDb(busIdx gdnative.Int) gdnative.Float {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetBusVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_bus_volume_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns the sample rate at the output of the audioserver.
	Args: [], Returns: float
*/
func (o *audioServer) GetMixRate() gdnative.Float {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.GetMixRate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "get_mix_rate")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns the speaker configuration.
	Args: [], Returns: enum.AudioServer::SpeakerMode
*/

/*
        If [code]true[/code] the bus at index [code]bus_idx[/code] is bypassing effects.
	Args: [{ false bus_idx int}], Returns: bool
*/
func (o *audioServer) IsBusBypassingEffects(busIdx gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.IsBusBypassingEffects()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "is_bus_bypassing_effects")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code] the effect at index [code]effect_idx[/code] on the bus at index [code]bus_idx[/code] is enabled.
	Args: [{ false bus_idx int} { false effect_idx int}], Returns: bool
*/
func (o *audioServer) IsBusEffectEnabled(busIdx gdnative.Int, effectIdx gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.IsBusEffectEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(effectIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "is_bus_effect_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code] the bus at index [code]bus_idx[/code] is muted.
	Args: [{ false bus_idx int}], Returns: bool
*/
func (o *audioServer) IsBusMute(busIdx gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.IsBusMute()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "is_bus_mute")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code] the bus at index [code]bus_idx[/code] is in solo mode.
	Args: [{ false bus_idx int}], Returns: bool
*/
func (o *audioServer) IsBusSolo(busIdx gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.IsBusSolo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "is_bus_solo")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Locks the audio drivers mainloop. Remember to unlock it afterwards.
	Args: [], Returns: void
*/
func (o *audioServer) Lock() {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.Lock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "lock")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Moves the bus from index [code]index[/code] to index [code]to_index[/code].
	Args: [{ false index int} { false to_index int}], Returns: void
*/
func (o *audioServer) MoveBus(index gdnative.Int, toIndex gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.MoveBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromInt(toIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "move_bus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes the bus at index [code]index[/code].
	Args: [{ false index int}], Returns: void
*/
func (o *audioServer) RemoveBus(index gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.RemoveBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "remove_bus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes the effect at index [code]effect_idx[/code] from the bus at index [code]bus_idx[/code].
	Args: [{ false bus_idx int} { false effect_idx int}], Returns: void
*/
func (o *audioServer) RemoveBusEffect(busIdx gdnative.Int, effectIdx gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.RemoveBusEffect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(effectIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "remove_bus_effect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code] the bus at index [code]bus_idx[/code] is bypassing effects.
	Args: [{ false bus_idx int} { false enable bool}], Returns: void
*/
func (o *audioServer) SetBusBypassEffects(busIdx gdnative.Int, enable gdnative.Bool) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusBypassEffects()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_bypass_effects")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds and removes busses to make the number of busses match [code]amount[/code].
	Args: [{ false amount int}], Returns: void
*/
func (o *audioServer) SetBusCount(amount gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_count")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code] the effect at index [code]effect_idx[/code] on the bus at index [code]bus_idx[/code] is enabled.
	Args: [{ false bus_idx int} { false effect_idx int} { false enabled bool}], Returns: void
*/
func (o *audioServer) SetBusEffectEnabled(busIdx gdnative.Int, effectIdx gdnative.Int, enabled gdnative.Bool) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusEffectEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(effectIdx)
	ptrArguments[2] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_effect_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Overwrites the currently used [AudioBusLayout].
	Args: [{ false bus_layout AudioBusLayout}], Returns: void
*/
func (o *audioServer) SetBusLayout(busLayout AudioBusLayout) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusLayout()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(busLayout.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_layout")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code] the bus at index [code]bus_idx[/code] is muted.
	Args: [{ false bus_idx int} { false enable bool}], Returns: void
*/
func (o *audioServer) SetBusMute(busIdx gdnative.Int, enable gdnative.Bool) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusMute()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_mute")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the name of the bus at index [code]bus_idx[/code] to [code]name[/code].
	Args: [{ false bus_idx int} { false name String}], Returns: void
*/
func (o *audioServer) SetBusName(busIdx gdnative.Int, name gdnative.String) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Connects the output of the bus at [code]bus_idx[/code] to the bus named [code]send[/code].
	Args: [{ false bus_idx int} { false send String}], Returns: void
*/
func (o *audioServer) SetBusSend(busIdx gdnative.Int, send gdnative.String) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusSend()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromString(send)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_send")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code] the bus at index [code]bus_idx[/code] is in solo mode.
	Args: [{ false bus_idx int} { false enable bool}], Returns: void
*/
func (o *audioServer) SetBusSolo(busIdx gdnative.Int, enable gdnative.Bool) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusSolo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_solo")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the volume of the bus at index [code]bus_idx[/code] to [code]volume_db[/code].
	Args: [{ false bus_idx int} { false volume_db float}], Returns: void
*/
func (o *audioServer) SetBusVolumeDb(busIdx gdnative.Int, volumeDb gdnative.Float) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SetBusVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromFloat(volumeDb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "set_bus_volume_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Swaps the position of two effects in bus [code]bus_idx[/code].
	Args: [{ false bus_idx int} { false effect_idx int} { false by_effect_idx int}], Returns: void
*/
func (o *audioServer) SwapBusEffects(busIdx gdnative.Int, effectIdx gdnative.Int, byEffectIdx gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.SwapBusEffects()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(busIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(effectIdx)
	ptrArguments[2] = gdnative.NewPointerFromInt(byEffectIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "swap_bus_effects")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Unlocks the audiodriver's main loop. After locking it always unlock it.
	Args: [], Returns: void
*/
func (o *audioServer) Unlock() {
	o.ensureSingleton()
	//log.Println("Calling AudioServer.Unlock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioServer", "unlock")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioServerImplementer is an interface that implements the methods
// of the AudioServer class.
type AudioServerImplementer interface {
	ObjectImplementer
	AddBus(atPosition gdnative.Int)
	AddBusEffect(busIdx gdnative.Int, effect AudioEffect, atPosition gdnative.Int)
	GenerateBusLayout() AudioBusLayoutImplementer
	GetBusCount() gdnative.Int
	GetBusEffect(busIdx gdnative.Int, effectIdx gdnative.Int) AudioEffectImplementer
	GetBusEffectCount(busIdx gdnative.Int) gdnative.Int
	GetBusIndex(busName gdnative.String) gdnative.Int
	GetBusName(busIdx gdnative.Int) gdnative.String
	GetBusPeakVolumeLeftDb(busIdx gdnative.Int, channel gdnative.Int) gdnative.Float
	GetBusPeakVolumeRightDb(busIdx gdnative.Int, channel gdnative.Int) gdnative.Float
	GetBusSend(busIdx gdnative.Int) gdnative.String
	GetBusVolumeDb(busIdx gdnative.Int) gdnative.Float
	GetMixRate() gdnative.Float
	IsBusBypassingEffects(busIdx gdnative.Int) gdnative.Bool
	IsBusEffectEnabled(busIdx gdnative.Int, effectIdx gdnative.Int) gdnative.Bool
	IsBusMute(busIdx gdnative.Int) gdnative.Bool
	IsBusSolo(busIdx gdnative.Int) gdnative.Bool
	Lock()
	MoveBus(index gdnative.Int, toIndex gdnative.Int)
	RemoveBus(index gdnative.Int)
	RemoveBusEffect(busIdx gdnative.Int, effectIdx gdnative.Int)
	SetBusBypassEffects(busIdx gdnative.Int, enable gdnative.Bool)
	SetBusCount(amount gdnative.Int)
	SetBusEffectEnabled(busIdx gdnative.Int, effectIdx gdnative.Int, enabled gdnative.Bool)
	SetBusLayout(busLayout AudioBusLayout)
	SetBusMute(busIdx gdnative.Int, enable gdnative.Bool)
	SetBusName(busIdx gdnative.Int, name gdnative.String)
	SetBusSend(busIdx gdnative.Int, send gdnative.String)
	SetBusSolo(busIdx gdnative.Int, enable gdnative.Bool)
	SetBusVolumeDb(busIdx gdnative.Int, volumeDb gdnative.Float)
	SwapBusEffects(busIdx gdnative.Int, effectIdx gdnative.Int, byEffectIdx gdnative.Int)
	Unlock()
}

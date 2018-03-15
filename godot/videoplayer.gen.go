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

//func NewVideoPlayerFromPointer(ptr gdnative.Pointer) VideoPlayer {
func newVideoPlayerFromPointer(ptr gdnative.Pointer) VideoPlayer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VideoPlayer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This control has the ability to play video streams. The only format accepted is the OGV Theora, so any other format must be converted before using in a project.
*/
type VideoPlayer struct {
	Control
	owner gdnative.Object
}

func (o *VideoPlayer) BaseClass() string {
	return "VideoPlayer"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *VideoPlayer) GetAudioTrack() gdnative.Int {
	//log.Println("Calling VideoPlayer.GetAudioTrack()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_audio_track")

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
	Args: [], Returns: int
*/
func (o *VideoPlayer) GetBufferingMsec() gdnative.Int {
	//log.Println("Calling VideoPlayer.GetBufferingMsec()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_buffering_msec")

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
	Args: [], Returns: String
*/
func (o *VideoPlayer) GetBus() gdnative.String {
	//log.Println("Calling VideoPlayer.GetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_bus")

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
	Args: [], Returns: VideoStream
*/
func (o *VideoPlayer) GetStream() VideoStreamImplementer {
	//log.Println("Calling VideoPlayer.GetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_stream")

	// Call the parent method.
	// VideoStream
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newVideoStreamFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(VideoStreamImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "VideoStream" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(VideoStreamImplementer)
	}

	return &ret
}

/*
        Get the name of the video stream.
	Args: [], Returns: String
*/
func (o *VideoPlayer) GetStreamName() gdnative.String {
	//log.Println("Calling VideoPlayer.GetStreamName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_stream_name")

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
	Args: [], Returns: float
*/
func (o *VideoPlayer) GetStreamPosition() gdnative.Real {
	//log.Println("Calling VideoPlayer.GetStreamPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_stream_position")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Get the current frame of the video as a [Texture].
	Args: [], Returns: Texture
*/
func (o *VideoPlayer) GetVideoTexture() TextureImplementer {
	//log.Println("Calling VideoPlayer.GetVideoTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_video_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *VideoPlayer) GetVolume() gdnative.Real {
	//log.Println("Calling VideoPlayer.GetVolume()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_volume")

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
func (o *VideoPlayer) GetVolumeDb() gdnative.Real {
	//log.Println("Calling VideoPlayer.GetVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "get_volume_db")

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
func (o *VideoPlayer) HasAutoplay() gdnative.Bool {
	//log.Println("Calling VideoPlayer.HasAutoplay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "has_autoplay")

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
func (o *VideoPlayer) HasExpand() gdnative.Bool {
	//log.Println("Calling VideoPlayer.HasExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "has_expand")

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
func (o *VideoPlayer) IsPaused() gdnative.Bool {
	//log.Println("Calling VideoPlayer.IsPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "is_paused")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Get whether or not the video is playing.
	Args: [], Returns: bool
*/
func (o *VideoPlayer) IsPlaying() gdnative.Bool {
	//log.Println("Calling VideoPlayer.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "is_playing")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Start the video playback.
	Args: [], Returns: void
*/
func (o *VideoPlayer) Play() {
	//log.Println("Calling VideoPlayer.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false track int}], Returns: void
*/
func (o *VideoPlayer) SetAudioTrack(track gdnative.Int) {
	//log.Println("Calling VideoPlayer.SetAudioTrack()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(track)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_audio_track")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *VideoPlayer) SetAutoplay(enabled gdnative.Bool) {
	//log.Println("Calling VideoPlayer.SetAutoplay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_autoplay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false msec int}], Returns: void
*/
func (o *VideoPlayer) SetBufferingMsec(msec gdnative.Int) {
	//log.Println("Calling VideoPlayer.SetBufferingMsec()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(msec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_buffering_msec")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bus String}], Returns: void
*/
func (o *VideoPlayer) SetBus(bus gdnative.String) {
	//log.Println("Calling VideoPlayer.SetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(bus)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_bus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *VideoPlayer) SetExpand(enable gdnative.Bool) {
	//log.Println("Calling VideoPlayer.SetExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_expand")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false paused bool}], Returns: void
*/
func (o *VideoPlayer) SetPaused(paused gdnative.Bool) {
	//log.Println("Calling VideoPlayer.SetPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(paused)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_paused")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stream VideoStream}], Returns: void
*/
func (o *VideoPlayer) SetStream(stream VideoStreamImplementer) {
	//log.Println("Calling VideoPlayer.SetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(stream.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_stream")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false position float}], Returns: void
*/
func (o *VideoPlayer) SetStreamPosition(position gdnative.Real) {
	//log.Println("Calling VideoPlayer.SetStreamPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_stream_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false volume float}], Returns: void
*/
func (o *VideoPlayer) SetVolume(volume gdnative.Real) {
	//log.Println("Calling VideoPlayer.SetVolume()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(volume)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_volume")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false db float}], Returns: void
*/
func (o *VideoPlayer) SetVolumeDb(db gdnative.Real) {
	//log.Println("Calling VideoPlayer.SetVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(db)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "set_volume_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stop the video playback.
	Args: [], Returns: void
*/
func (o *VideoPlayer) Stop() {
	//log.Println("Calling VideoPlayer.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoPlayer", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VideoPlayerImplementer is an interface that implements the methods
// of the VideoPlayer class.
type VideoPlayerImplementer interface {
	ControlImplementer
	GetAudioTrack() gdnative.Int
	GetBufferingMsec() gdnative.Int
	GetBus() gdnative.String
	GetStream() VideoStreamImplementer
	GetStreamName() gdnative.String
	GetStreamPosition() gdnative.Real
	GetVideoTexture() TextureImplementer
	GetVolume() gdnative.Real
	GetVolumeDb() gdnative.Real
	HasAutoplay() gdnative.Bool
	HasExpand() gdnative.Bool
	IsPaused() gdnative.Bool
	IsPlaying() gdnative.Bool
	Play()
	SetAudioTrack(track gdnative.Int)
	SetAutoplay(enabled gdnative.Bool)
	SetBufferingMsec(msec gdnative.Int)
	SetBus(bus gdnative.String)
	SetExpand(enable gdnative.Bool)
	SetPaused(paused gdnative.Bool)
	SetStream(stream VideoStreamImplementer)
	SetStreamPosition(position gdnative.Real)
	SetVolume(volume gdnative.Real)
	SetVolumeDb(db gdnative.Real)
	Stop()
}

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

// TimerTimerProcessMode is an enum for TimerProcessMode values.
type TimerTimerProcessMode int

const (
	TimerTimerProcessIdle    TimerTimerProcessMode = 1
	TimerTimerProcessPhysics TimerTimerProcessMode = 0
)

//func NewTimerFromPointer(ptr gdnative.Pointer) Timer {
func newTimerFromPointer(ptr gdnative.Pointer) Timer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Timer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Counts down a specified interval and emits a signal on reaching 0. Can be set to repeat or "one shot" mode.
*/
type Timer struct {
	Node
	owner gdnative.Object
}

func (o *Timer) BaseClass() string {
	return "Timer"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Timer) GetTimeLeft() gdnative.Real {
	//log.Println("Calling Timer.GetTimeLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "get_time_left")

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
	Args: [], Returns: enum.Timer::TimerProcessMode
*/
func (o *Timer) GetTimerProcessMode() TimerTimerProcessMode {
	//log.Println("Calling Timer.GetTimerProcessMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "get_timer_process_mode")

	// Call the parent method.
	// enum.Timer::TimerProcessMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return TimerTimerProcessMode(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Timer) GetWaitTime() gdnative.Real {
	//log.Println("Calling Timer.GetWaitTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "get_wait_time")

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
func (o *Timer) HasAutostart() gdnative.Bool {
	//log.Println("Calling Timer.HasAutostart()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "has_autostart")

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
func (o *Timer) IsOneShot() gdnative.Bool {
	//log.Println("Calling Timer.IsOneShot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "is_one_shot")

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
func (o *Timer) IsPaused() gdnative.Bool {
	//log.Println("Calling Timer.IsPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "is_paused")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the timer is stopped.
	Args: [], Returns: bool
*/
func (o *Timer) IsStopped() gdnative.Bool {
	//log.Println("Calling Timer.IsStopped()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "is_stopped")

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
func (o *Timer) SetAutostart(enable gdnative.Bool) {
	//log.Println("Calling Timer.SetAutostart()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "set_autostart")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Timer) SetOneShot(enable gdnative.Bool) {
	//log.Println("Calling Timer.SetOneShot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "set_one_shot")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false paused bool}], Returns: void
*/
func (o *Timer) SetPaused(paused gdnative.Bool) {
	//log.Println("Calling Timer.SetPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(paused)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "set_paused")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Timer) SetTimerProcessMode(mode gdnative.Int) {
	//log.Println("Calling Timer.SetTimerProcessMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "set_timer_process_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false time_sec float}], Returns: void
*/
func (o *Timer) SetWaitTime(timeSec gdnative.Real) {
	//log.Println("Calling Timer.SetWaitTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(timeSec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "set_wait_time")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Starts the timer. This also resets the remaining time to [code]wait_time[/code]. Note: this method will not resume a paused timer. See [method set_paused].
	Args: [], Returns: void
*/
func (o *Timer) Start() {
	//log.Println("Calling Timer.Start()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "start")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stop (cancel) the Timer.
	Args: [], Returns: void
*/
func (o *Timer) Stop() {
	//log.Println("Calling Timer.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Timer", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TimerImplementer is an interface that implements the methods
// of the Timer class.
type TimerImplementer interface {
	NodeImplementer
	GetTimeLeft() gdnative.Real
	GetWaitTime() gdnative.Real
	HasAutostart() gdnative.Bool
	IsOneShot() gdnative.Bool
	IsPaused() gdnative.Bool
	IsStopped() gdnative.Bool
	SetAutostart(enable gdnative.Bool)
	SetOneShot(enable gdnative.Bool)
	SetPaused(paused gdnative.Bool)
	SetTimerProcessMode(mode gdnative.Int)
	SetWaitTime(timeSec gdnative.Real)
	Start()
	Stop()
}

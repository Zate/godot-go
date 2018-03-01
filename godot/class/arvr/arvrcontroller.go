package arvr

import (
	"log"
	"reflect"

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

/*
This is a helper spatial node that is linked to the tracking of controllers. It also offers several handy pass throughs to the state of buttons and such on the controllers. Controllers are linked by their id. You can create controller nodes before the controllers are available. Say your game always uses two controllers (one for each hand) you can predefine the controllers with id 1 and 2 and they will become active as soon as the controllers are identified. If you expect additional controllers to be used you should react to the signals and add ARVRController nodes to your scene. The position of the controller node is automatically updated by the ARVR Server. This makes this node ideal to add child nodes to visualise the controller.
*/
type ARVRController struct {
	Spatial
}

func (o *ARVRController) BaseClass() string {
	return "ARVRController"
}

/*
   Undocumented
*/
func (o *ARVRController) GetControllerId() gdnative.Int {
	log.Println("Calling ARVRController.GetControllerId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_controller_id", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If active, returns the name of the associated controller if provided by the AR/VR SDK used.
*/
func (o *ARVRController) GetControllerName() gdnative.String {
	log.Println("Calling ARVRController.GetControllerName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_controller_name", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the hand holding this controller, if known. See TRACKER_* constants in [ARVRPositionalTracker].
*/
func (o *ARVRController) GetHand() gdnative.Int {
	log.Println("Calling ARVRController.GetHand()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_hand", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the bound controller is active. ARVR systems attempt to track active controllers.
*/
func (o *ARVRController) GetIsActive() gdnative.Bool {
	log.Println("Calling ARVRController.GetIsActive()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_is_active", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the value of the given axis for things like triggers, touchpads, etc. that are embedded into the controller.
*/
func (o *ARVRController) GetJoystickAxis(axis gdnative.Int) gdnative.Float {
	log.Println("Calling ARVRController.GetJoystickAxis()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(axis)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joystick_axis", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the ID of the joystick object bound to this. Every controller tracked by the ARVR Server that has buttons and axis will also be registered as a joystick within Godot. This means that all the normal joystick tracking and input mapping will work for buttons and axis found on the AR/VR controllers. This ID is purely offered as information so you can link up the controller with its joystick entry.
*/
func (o *ARVRController) GetJoystickId() gdnative.Int {
	log.Println("Calling ARVRController.GetJoystickId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_joystick_id", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRController) GetRumble() gdnative.Float {
	log.Println("Calling ARVRController.GetRumble()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_rumble", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the button at index [code]button[/code] is pressed.
*/
func (o *ARVRController) IsButtonPressed(button gdnative.Int) gdnative.Int {
	log.Println("Calling ARVRController.IsButtonPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(button)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_button_pressed", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRController) SetControllerId(controllerId gdnative.Int) {
	log.Println("Calling ARVRController.SetControllerId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(controllerId)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_controller_id", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ARVRController) SetRumble(rumble gdnative.Float) {
	log.Println("Calling ARVRController.SetRumble()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rumble)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_rumble", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ARVRControllerImplementer is an interface for ARVRController objects.
*/
type ARVRControllerImplementer interface {
	Class
}
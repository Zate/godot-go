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

//func NewSpatialFromPointer(ptr gdnative.Pointer) Spatial {
func newSpatialFromPointer(ptr gdnative.Pointer) Spatial {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Spatial{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Most basic 3D game object, with a 3D [Transform] and visibility settings. All other 3D game objects inherit from Spatial. Use Spatial as a parent node to move, scale, rotate and show/hide children in a 3D project. Affine operations (rotate, scale, translate) happen in parent's local coordinate system, unless the Spatial object is set as top level. Affine operations in this coordinate system correspond to direct affine operations on the Spatial's transform. The word local below refers to this coordinate system. The coordinate system that is attached to the Spatial object itself is referred to as object-local coordinate system.
*/
type Spatial struct {
	Node
	owner gdnative.Object
}

func (o *Spatial) BaseClass() string {
	return "Spatial"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Spatial) X_UpdateGizmo() {
	//log.Println("Calling Spatial.X_UpdateGizmo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "_update_gizmo")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: SpatialGizmo
*/
func (o *Spatial) GetGizmo() SpatialGizmoImplementer {
	//log.Println("Calling Spatial.GetGizmo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_gizmo")

	// Call the parent method.
	// SpatialGizmo
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSpatialGizmoFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SpatialGizmoImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "SpatialGizmo" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SpatialGizmoImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Transform
*/
func (o *Spatial) GetGlobalTransform() gdnative.Transform {
	//log.Println("Calling Spatial.GetGlobalTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_global_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	return ret
}

/*
        Returns the parent [code]Spatial[/code], or an empty [Object] if no parent exists or parent is not of type [code]Spatial[/code].
	Args: [], Returns: Spatial
*/
func (o *Spatial) GetParentSpatial() SpatialImplementer {
	//log.Println("Calling Spatial.GetParentSpatial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_parent_spatial")

	// Call the parent method.
	// Spatial
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSpatialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SpatialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Spatial" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SpatialImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *Spatial) GetRotation() gdnative.Vector3 {
	//log.Println("Calling Spatial.GetRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_rotation")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *Spatial) GetRotationDegrees() gdnative.Vector3 {
	//log.Println("Calling Spatial.GetRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_rotation_degrees")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *Spatial) GetScale() gdnative.Vector3 {
	//log.Println("Calling Spatial.GetScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_scale")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Transform
*/
func (o *Spatial) GetTransform() gdnative.Transform {
	//log.Println("Calling Spatial.GetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *Spatial) GetTranslation() gdnative.Vector3 {
	//log.Println("Calling Spatial.GetTranslation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_translation")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the current [World] resource this Spatial node is registered to.
	Args: [], Returns: World
*/
func (o *Spatial) GetWorld() WorldImplementer {
	//log.Println("Calling Spatial.GetWorld()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "get_world")

	// Call the parent method.
	// World
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newWorldFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(WorldImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "World" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(WorldImplementer)
	}

	return &ret
}

/*
        Rotates the global (world) transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in global coordinate system.
	Args: [{ false axis Vector3} { false angle float}], Returns: void
*/
func (o *Spatial) GlobalRotate(axis gdnative.Vector3, angle gdnative.Float) {
	//log.Println("Calling Spatial.GlobalRotate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(axis)
	ptrArguments[1] = gdnative.NewPointerFromFloat(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "global_rotate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false scale Vector3}], Returns: void
*/
func (o *Spatial) GlobalScale(scale gdnative.Vector3) {
	//log.Println("Calling Spatial.GlobalScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "global_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Moves the global (world) transformation by [Vector3] offset. The offset is in global coordinate system.
	Args: [{ false offset Vector3}], Returns: void
*/
func (o *Spatial) GlobalTranslate(offset gdnative.Vector3) {
	//log.Println("Calling Spatial.GlobalTranslate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "global_translate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Disables rendering of this node. Change Spatial Visible property to false.
	Args: [], Returns: void
*/
func (o *Spatial) Hide() {
	//log.Println("Calling Spatial.Hide()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "hide")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns whether node notifies about its local transformation changes. Spatial will not propagate this by default.
	Args: [], Returns: bool
*/
func (o *Spatial) IsLocalTransformNotificationEnabled() gdnative.Bool {
	//log.Println("Calling Spatial.IsLocalTransformNotificationEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "is_local_transform_notification_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns whether this node is set as Toplevel, that is whether it ignores its parent nodes transformations.
	Args: [], Returns: bool
*/
func (o *Spatial) IsSetAsToplevel() gdnative.Bool {
	//log.Println("Calling Spatial.IsSetAsToplevel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "is_set_as_toplevel")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns whether the node notifies about its global and local transformation changes. Spatial will not propagate this by default.
	Args: [], Returns: bool
*/
func (o *Spatial) IsTransformNotificationEnabled() gdnative.Bool {
	//log.Println("Calling Spatial.IsTransformNotificationEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "is_transform_notification_enabled")

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
func (o *Spatial) IsVisible() gdnative.Bool {
	//log.Println("Calling Spatial.IsVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "is_visible")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns whether the node is visible, taking into consideration that its parents visibility.
	Args: [], Returns: bool
*/
func (o *Spatial) IsVisibleInTree() gdnative.Bool {
	//log.Println("Calling Spatial.IsVisibleInTree()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "is_visible_in_tree")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Rotates itself to point into direction of target position. Operations take place in global space.
	Args: [{ false target Vector3} { false up Vector3}], Returns: void
*/
func (o *Spatial) LookAt(target gdnative.Vector3, up gdnative.Vector3) {
	//log.Println("Calling Spatial.LookAt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(target)
	ptrArguments[1] = gdnative.NewPointerFromVector3(up)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "look_at")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Moves the node to specified position and then rotates itself to point into direction of target position. Operations take place in global space.
	Args: [{ false position Vector3} { false target Vector3} { false up Vector3}], Returns: void
*/
func (o *Spatial) LookAtFromPosition(position gdnative.Vector3, target gdnative.Vector3, up gdnative.Vector3) {
	//log.Println("Calling Spatial.LookAtFromPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromVector3(position)
	ptrArguments[1] = gdnative.NewPointerFromVector3(target)
	ptrArguments[2] = gdnative.NewPointerFromVector3(up)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "look_at_from_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Resets this node's transformations (like scale, skew and taper) preserving its rotation and translation by performing Gram-Schmidt orthonormalization on this node's [Transform3D].
	Args: [], Returns: void
*/
func (o *Spatial) Orthonormalize() {
	//log.Println("Calling Spatial.Orthonormalize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "orthonormalize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians.
	Args: [{ false axis Vector3} { false angle float}], Returns: void
*/
func (o *Spatial) Rotate(axis gdnative.Vector3, angle gdnative.Float) {
	//log.Println("Calling Spatial.Rotate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(axis)
	ptrArguments[1] = gdnative.NewPointerFromFloat(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "rotate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in object-local coordinate system.
	Args: [{ false axis Vector3} { false angle float}], Returns: void
*/
func (o *Spatial) RotateObjectLocal(axis gdnative.Vector3, angle gdnative.Float) {
	//log.Println("Calling Spatial.RotateObjectLocal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(axis)
	ptrArguments[1] = gdnative.NewPointerFromFloat(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "rotate_object_local")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Rotates the local transformation around the X axis by angle in radians
	Args: [{ false angle float}], Returns: void
*/
func (o *Spatial) RotateX(angle gdnative.Float) {
	//log.Println("Calling Spatial.RotateX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "rotate_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Rotates the local transformation around the Y axis by angle in radians.
	Args: [{ false angle float}], Returns: void
*/
func (o *Spatial) RotateY(angle gdnative.Float) {
	//log.Println("Calling Spatial.RotateY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "rotate_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Rotates the local transformation around the Z axis by angle in radians.
	Args: [{ false angle float}], Returns: void
*/
func (o *Spatial) RotateZ(angle gdnative.Float) {
	//log.Println("Calling Spatial.RotateZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "rotate_z")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Scales the local transformation by given 3D scale factors in object-local coordinate system.
	Args: [{ false scale Vector3}], Returns: void
*/
func (o *Spatial) ScaleObjectLocal(scale gdnative.Vector3) {
	//log.Println("Calling Spatial.ScaleObjectLocal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "scale_object_local")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Makes the node ignore its parents transformations. Node transformations are only in global space.
	Args: [{ false enable bool}], Returns: void
*/
func (o *Spatial) SetAsToplevel(enable gdnative.Bool) {
	//log.Println("Calling Spatial.SetAsToplevel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_as_toplevel")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false gizmo SpatialGizmo}], Returns: void
*/
func (o *Spatial) SetGizmo(gizmo SpatialGizmo) {
	//log.Println("Calling Spatial.SetGizmo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_gizmo")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false global Transform}], Returns: void
*/
func (o *Spatial) SetGlobalTransform(global gdnative.Transform) {
	//log.Println("Calling Spatial.SetGlobalTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform(global)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_global_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Reset all transformations for this node. Set its [Transform3D] to identity matrix.
	Args: [], Returns: void
*/
func (o *Spatial) SetIdentity() {
	//log.Println("Calling Spatial.SetIdentity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_identity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set whether the node ignores notification that its transformation (global or local) changed.
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Spatial) SetIgnoreTransformNotification(enabled gdnative.Bool) {
	//log.Println("Calling Spatial.SetIgnoreTransformNotification()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_ignore_transform_notification")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set whether the node notifies about its local transformation changes. Spatial will not propagate this by default.
	Args: [{ false enable bool}], Returns: void
*/
func (o *Spatial) SetNotifyLocalTransform(enable gdnative.Bool) {
	//log.Println("Calling Spatial.SetNotifyLocalTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_notify_local_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set whether the node notifies about its global and local transformation changes. Spatial will not propagate this by default.
	Args: [{ false enable bool}], Returns: void
*/
func (o *Spatial) SetNotifyTransform(enable gdnative.Bool) {
	//log.Println("Calling Spatial.SetNotifyTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_notify_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false euler Vector3}], Returns: void
*/
func (o *Spatial) SetRotation(euler gdnative.Vector3) {
	//log.Println("Calling Spatial.SetRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(euler)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_rotation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false euler_degrees Vector3}], Returns: void
*/
func (o *Spatial) SetRotationDegrees(eulerDegrees gdnative.Vector3) {
	//log.Println("Calling Spatial.SetRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(eulerDegrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_rotation_degrees")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scale Vector3}], Returns: void
*/
func (o *Spatial) SetScale(scale gdnative.Vector3) {
	//log.Println("Calling Spatial.SetScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false local Transform}], Returns: void
*/
func (o *Spatial) SetTransform(local gdnative.Transform) {
	//log.Println("Calling Spatial.SetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform(local)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false translation Vector3}], Returns: void
*/
func (o *Spatial) SetTranslation(translation gdnative.Vector3) {
	//log.Println("Calling Spatial.SetTranslation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(translation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_translation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false visible bool}], Returns: void
*/
func (o *Spatial) SetVisible(visible gdnative.Bool) {
	//log.Println("Calling Spatial.SetVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(visible)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "set_visible")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Enables rendering of this node. Change Spatial Visible property to "True".
	Args: [], Returns: void
*/
func (o *Spatial) Show() {
	//log.Println("Calling Spatial.Show()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "show")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Transforms [Vector3] "local_point" from this node's local space to world space.
	Args: [{ false local_point Vector3}], Returns: Vector3
*/
func (o *Spatial) ToGlobal(localPoint gdnative.Vector3) gdnative.Vector3 {
	//log.Println("Calling Spatial.ToGlobal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(localPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "to_global")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Transforms [Vector3] "global_point" from world space to this node's local space.
	Args: [{ false global_point Vector3}], Returns: Vector3
*/
func (o *Spatial) ToLocal(globalPoint gdnative.Vector3) gdnative.Vector3 {
	//log.Println("Calling Spatial.ToLocal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(globalPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "to_local")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Changes the node's position by given offset [Vector3].
	Args: [{ false offset Vector3}], Returns: void
*/
func (o *Spatial) Translate(offset gdnative.Vector3) {
	//log.Println("Calling Spatial.Translate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "translate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false offset Vector3}], Returns: void
*/
func (o *Spatial) TranslateObjectLocal(offset gdnative.Vector3) {
	//log.Println("Calling Spatial.TranslateObjectLocal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "translate_object_local")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Updates the [SpatialGizmo] of this node.
	Args: [], Returns: void
*/
func (o *Spatial) UpdateGizmo() {
	//log.Println("Calling Spatial.UpdateGizmo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Spatial", "update_gizmo")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SpatialImplementer is an interface that implements the methods
// of the Spatial class.
type SpatialImplementer interface {
	NodeImplementer
	X_UpdateGizmo()
	GetGizmo() SpatialGizmoImplementer
	GetGlobalTransform() gdnative.Transform
	GetParentSpatial() SpatialImplementer
	GetRotation() gdnative.Vector3
	GetRotationDegrees() gdnative.Vector3
	GetScale() gdnative.Vector3
	GetTransform() gdnative.Transform
	GetTranslation() gdnative.Vector3
	GetWorld() WorldImplementer
	GlobalRotate(axis gdnative.Vector3, angle gdnative.Float)
	GlobalScale(scale gdnative.Vector3)
	GlobalTranslate(offset gdnative.Vector3)
	Hide()
	IsLocalTransformNotificationEnabled() gdnative.Bool
	IsSetAsToplevel() gdnative.Bool
	IsTransformNotificationEnabled() gdnative.Bool
	IsVisible() gdnative.Bool
	IsVisibleInTree() gdnative.Bool
	LookAt(target gdnative.Vector3, up gdnative.Vector3)
	LookAtFromPosition(position gdnative.Vector3, target gdnative.Vector3, up gdnative.Vector3)
	Orthonormalize()
	Rotate(axis gdnative.Vector3, angle gdnative.Float)
	RotateObjectLocal(axis gdnative.Vector3, angle gdnative.Float)
	RotateX(angle gdnative.Float)
	RotateY(angle gdnative.Float)
	RotateZ(angle gdnative.Float)
	ScaleObjectLocal(scale gdnative.Vector3)
	SetAsToplevel(enable gdnative.Bool)
	SetGizmo(gizmo SpatialGizmo)
	SetGlobalTransform(global gdnative.Transform)
	SetIdentity()
	SetIgnoreTransformNotification(enabled gdnative.Bool)
	SetNotifyLocalTransform(enable gdnative.Bool)
	SetNotifyTransform(enable gdnative.Bool)
	SetRotation(euler gdnative.Vector3)
	SetRotationDegrees(eulerDegrees gdnative.Vector3)
	SetScale(scale gdnative.Vector3)
	SetTransform(local gdnative.Transform)
	SetTranslation(translation gdnative.Vector3)
	SetVisible(visible gdnative.Bool)
	Show()
	ToGlobal(localPoint gdnative.Vector3) gdnative.Vector3
	ToLocal(globalPoint gdnative.Vector3) gdnative.Vector3
	Translate(offset gdnative.Vector3)
	TranslateObjectLocal(offset gdnative.Vector3)
	UpdateGizmo()
}

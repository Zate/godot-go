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

//func NewPhysicsDirectBodyStateFromPointer(ptr gdnative.Pointer) PhysicsDirectBodyState {
func newPhysicsDirectBodyStateFromPointer(ptr gdnative.Pointer) PhysicsDirectBodyState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PhysicsDirectBodyState{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type PhysicsDirectBodyState struct {
	Object
	owner gdnative.Object
}

func (o *PhysicsDirectBodyState) BaseClass() string {
	return "PhysicsDirectBodyState"
}

/*

	Args: [{ false force Vector3} { false position Vector3}], Returns: void
*/
func (o *PhysicsDirectBodyState) AddForce(force gdnative.Vector3, position gdnative.Vector3) {
	//log.Println("Calling PhysicsDirectBodyState.AddForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(force)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "add_force")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false position Vector3} { false j Vector3}], Returns: void
*/
func (o *PhysicsDirectBodyState) ApplyImpulse(position gdnative.Vector3, j gdnative.Vector3) {
	//log.Println("Calling PhysicsDirectBodyState.ApplyImpulse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(position)
	ptrArguments[1] = gdnative.NewPointerFromVector3(j)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "apply_impulse")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false j Vector3}], Returns: void
*/
func (o *PhysicsDirectBodyState) ApplyTorqeImpulse(j gdnative.Vector3) {
	//log.Println("Calling PhysicsDirectBodyState.ApplyTorqeImpulse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(j)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "apply_torqe_impulse")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetAngularVelocity() gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_angular_velocity")

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
func (o *PhysicsDirectBodyState) GetCenterOfMass() gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetCenterOfMass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_center_of_mass")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: RID
*/
func (o *PhysicsDirectBodyState) GetContactCollider(contactIdx gdnative.Int) gdnative.Rid {
	//log.Println("Calling PhysicsDirectBodyState.GetContactCollider()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_collider")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: int
*/
func (o *PhysicsDirectBodyState) GetContactColliderId(contactIdx gdnative.Int) gdnative.Int {
	//log.Println("Calling PhysicsDirectBodyState.GetContactColliderId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_collider_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: Object
*/
func (o *PhysicsDirectBodyState) GetContactColliderObject(contactIdx gdnative.Int) ObjectImplementer {
	//log.Println("Calling PhysicsDirectBodyState.GetContactColliderObject()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_collider_object")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*

	Args: [{ false contact_idx int}], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetContactColliderPosition(contactIdx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetContactColliderPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_collider_position")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: int
*/
func (o *PhysicsDirectBodyState) GetContactColliderShape(contactIdx gdnative.Int) gdnative.Int {
	//log.Println("Calling PhysicsDirectBodyState.GetContactColliderShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_collider_shape")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetContactColliderVelocityAtPosition(contactIdx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetContactColliderVelocityAtPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_collider_velocity_at_position")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: int
*/
func (o *PhysicsDirectBodyState) GetContactCount() gdnative.Int {
	//log.Println("Calling PhysicsDirectBodyState.GetContactCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetContactLocalNormal(contactIdx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetContactLocalNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_local_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetContactLocalPosition(contactIdx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetContactLocalPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_local_position")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false contact_idx int}], Returns: int
*/
func (o *PhysicsDirectBodyState) GetContactLocalShape(contactIdx gdnative.Int) gdnative.Int {
	//log.Println("Calling PhysicsDirectBodyState.GetContactLocalShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_contact_local_shape")

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
	Args: [], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetInverseInertia() gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetInverseInertia()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_inverse_inertia")

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
	Args: [], Returns: float
*/
func (o *PhysicsDirectBodyState) GetInverseMass() gdnative.Float {
	//log.Println("Calling PhysicsDirectBodyState.GetInverseMass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_inverse_mass")

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
	Args: [], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetLinearVelocity() gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_linear_velocity")

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
	Args: [], Returns: Basis
*/
func (o *PhysicsDirectBodyState) GetPrincipalInertiaAxes() gdnative.Basis {
	//log.Println("Calling PhysicsDirectBodyState.GetPrincipalInertiaAxes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_principal_inertia_axes")

	// Call the parent method.
	// Basis
	retPtr := gdnative.NewEmptyBasis()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBasisFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: PhysicsDirectSpaceState
*/
func (o *PhysicsDirectBodyState) GetSpaceState() PhysicsDirectSpaceStateImplementer {
	//log.Println("Calling PhysicsDirectBodyState.GetSpaceState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_space_state")

	// Call the parent method.
	// PhysicsDirectSpaceState
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newPhysicsDirectSpaceStateFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(PhysicsDirectSpaceStateImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "PhysicsDirectSpaceState" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(PhysicsDirectSpaceStateImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PhysicsDirectBodyState) GetStep() gdnative.Float {
	//log.Println("Calling PhysicsDirectBodyState.GetStep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_step")

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
func (o *PhysicsDirectBodyState) GetTotalAngularDamp() gdnative.Float {
	//log.Println("Calling PhysicsDirectBodyState.GetTotalAngularDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_total_angular_damp")

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
	Args: [], Returns: Vector3
*/
func (o *PhysicsDirectBodyState) GetTotalGravity() gdnative.Vector3 {
	//log.Println("Calling PhysicsDirectBodyState.GetTotalGravity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_total_gravity")

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
	Args: [], Returns: float
*/
func (o *PhysicsDirectBodyState) GetTotalLinearDamp() gdnative.Float {
	//log.Println("Calling PhysicsDirectBodyState.GetTotalLinearDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_total_linear_damp")

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
	Args: [], Returns: Transform
*/
func (o *PhysicsDirectBodyState) GetTransform() gdnative.Transform {
	//log.Println("Calling PhysicsDirectBodyState.GetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "get_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: void
*/
func (o *PhysicsDirectBodyState) IntegrateForces() {
	//log.Println("Calling PhysicsDirectBodyState.IntegrateForces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "integrate_forces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *PhysicsDirectBodyState) IsSleeping() gdnative.Bool {
	//log.Println("Calling PhysicsDirectBodyState.IsSleeping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "is_sleeping")

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
	Args: [{ false velocity Vector3}], Returns: void
*/
func (o *PhysicsDirectBodyState) SetAngularVelocity(velocity gdnative.Vector3) {
	//log.Println("Calling PhysicsDirectBodyState.SetAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(velocity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "set_angular_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false velocity Vector3}], Returns: void
*/
func (o *PhysicsDirectBodyState) SetLinearVelocity(velocity gdnative.Vector3) {
	//log.Println("Calling PhysicsDirectBodyState.SetLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(velocity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "set_linear_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *PhysicsDirectBodyState) SetSleepState(enabled gdnative.Bool) {
	//log.Println("Calling PhysicsDirectBodyState.SetSleepState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "set_sleep_state")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false transform Transform}], Returns: void
*/
func (o *PhysicsDirectBodyState) SetTransform(transform gdnative.Transform) {
	//log.Println("Calling PhysicsDirectBodyState.SetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform(transform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectBodyState", "set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PhysicsDirectBodyStateImplementer is an interface that implements the methods
// of the PhysicsDirectBodyState class.
type PhysicsDirectBodyStateImplementer interface {
	ObjectImplementer
	AddForce(force gdnative.Vector3, position gdnative.Vector3)
	ApplyImpulse(position gdnative.Vector3, j gdnative.Vector3)
	ApplyTorqeImpulse(j gdnative.Vector3)
	GetAngularVelocity() gdnative.Vector3
	GetCenterOfMass() gdnative.Vector3
	GetContactCollider(contactIdx gdnative.Int) gdnative.Rid
	GetContactColliderId(contactIdx gdnative.Int) gdnative.Int
	GetContactColliderObject(contactIdx gdnative.Int) ObjectImplementer
	GetContactColliderPosition(contactIdx gdnative.Int) gdnative.Vector3
	GetContactColliderShape(contactIdx gdnative.Int) gdnative.Int
	GetContactColliderVelocityAtPosition(contactIdx gdnative.Int) gdnative.Vector3
	GetContactCount() gdnative.Int
	GetContactLocalNormal(contactIdx gdnative.Int) gdnative.Vector3
	GetContactLocalPosition(contactIdx gdnative.Int) gdnative.Vector3
	GetContactLocalShape(contactIdx gdnative.Int) gdnative.Int
	GetInverseInertia() gdnative.Vector3
	GetInverseMass() gdnative.Float
	GetLinearVelocity() gdnative.Vector3
	GetPrincipalInertiaAxes() gdnative.Basis
	GetSpaceState() PhysicsDirectSpaceStateImplementer
	GetStep() gdnative.Float
	GetTotalAngularDamp() gdnative.Float
	GetTotalGravity() gdnative.Vector3
	GetTotalLinearDamp() gdnative.Float
	GetTransform() gdnative.Transform
	IntegrateForces()
	IsSleeping() gdnative.Bool
	SetAngularVelocity(velocity gdnative.Vector3)
	SetLinearVelocity(velocity gdnative.Vector3)
	SetSleepState(enabled gdnative.Bool)
	SetTransform(transform gdnative.Transform)
}

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

//func NewKinematicBodyFromPointer(ptr gdnative.Pointer) KinematicBody {
func newKinematicBodyFromPointer(ptr gdnative.Pointer) KinematicBody {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := KinematicBody{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Kinematic bodies are special types of bodies that are meant to be user-controlled. They are not affected by physics at all (to other types of bodies, such a character or a rigid body, these are the same as a static body). They have however, two main uses: Simulated Motion: When these bodies are moved manually, either from code or from an AnimationPlayer (with process mode set to fixed), the physics will automatically compute an estimate of their linear and angular velocity. This makes them very useful for moving platforms or other AnimationPlayer-controlled objects (like a door, a bridge that opens, etc). Kinematic Characters: KinematicBody also has an API for moving objects (the [method move_and_collide] and [method move_and_slide] methods) while performing collision tests. This makes them really useful to implement characters that collide against a world, but that don't require advanced physics.
*/
type KinematicBody struct {
	PhysicsBody
	owner gdnative.Object
}

func (o *KinematicBody) BaseClass() string {
	return "KinematicBody"
}

/*
        Undocumented
	Args: [{ false axis int}], Returns: bool
*/
func (o *KinematicBody) GetAxisLock(axis gdnative.Int) gdnative.Bool {
	//log.Println("Calling KinematicBody.GetAxisLock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(axis)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "get_axis_lock")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the velocity of the floor. Only updates when calling [method move_and_slide].
	Args: [], Returns: Vector3
*/
func (o *KinematicBody) GetFloorVelocity() gdnative.Vector3 {
	//log.Println("Calling KinematicBody.GetFloorVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "get_floor_velocity")

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
func (o *KinematicBody) GetSafeMargin() gdnative.Real {
	//log.Println("Calling KinematicBody.GetSafeMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "get_safe_margin")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns a [KinematicCollision], which contains information about a collision that occurred during the last [method move_and_slide] call. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_count] - 1).
	Args: [{ false slide_idx int}], Returns: KinematicCollision
*/
func (o *KinematicBody) GetSlideCollision(slideIdx gdnative.Int) KinematicCollisionImplementer {
	//log.Println("Calling KinematicBody.GetSlideCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(slideIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "get_slide_collision")

	// Call the parent method.
	// KinematicCollision
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newKinematicCollisionFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(KinematicCollisionImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "KinematicCollision" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(KinematicCollisionImplementer)
	}

	return &ret
}

/*
        Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
	Args: [], Returns: int
*/
func (o *KinematicBody) GetSlideCount() gdnative.Int {
	//log.Println("Calling KinematicBody.GetSlideCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "get_slide_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the body is on the ceiling. Only updates when calling [method move_and_slide].
	Args: [], Returns: bool
*/
func (o *KinematicBody) IsOnCeiling() gdnative.Bool {
	//log.Println("Calling KinematicBody.IsOnCeiling()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "is_on_ceiling")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the body is on the floor. Only updates when calling [method move_and_slide].
	Args: [], Returns: bool
*/
func (o *KinematicBody) IsOnFloor() gdnative.Bool {
	//log.Println("Calling KinematicBody.IsOnFloor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "is_on_floor")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the body is on a wall. Only updates when calling [method move_and_slide].
	Args: [], Returns: bool
*/
func (o *KinematicBody) IsOnWall() gdnative.Bool {
	//log.Println("Calling KinematicBody.IsOnWall()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "is_on_wall")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Moves the body along the vector [code]rel_vec[/code]. The body will stop if it collides. Returns a [KinematicCollision], which contains information about the collision.
	Args: [{ false rel_vec Vector3}], Returns: KinematicCollision
*/
func (o *KinematicBody) MoveAndCollide(relVec gdnative.Vector3) KinematicCollisionImplementer {
	//log.Println("Calling KinematicBody.MoveAndCollide()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(relVec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "move_and_collide")

	// Call the parent method.
	// KinematicCollision
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newKinematicCollisionFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(KinematicCollisionImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "KinematicCollision" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(KinematicCollisionImplementer)
	}

	return &ret
}

/*
        Moves the body along a vector. If the body collides with another, it will slide along the other body rather than stop immediately. If the other body is a [code]KinematicBody[/code] or [RigidBody], it will also be affected by the motion of the other body. You can use this to make moving or rotating platforms, or to make nodes push other nodes. [code]linear_velocity[/code] is a value in pixels per second. Unlike in for example [method move_and_collide], you should [i]not[/i] multiply it with [code]delta[/code] — this is done by the method. [code]floor_normal[/code] is the up direction, used to determine what is a wall and what is a floor or a ceiling. If set to the default value of [code]Vector3(0, 0, 0)[/code], everything is considered a wall. This is useful for topdown games. If the body is standing on a slope and the horizontal speed (relative to the floor's speed) goes below [code]slope_stop_min_velocity[/code], the body will stop completely. This prevents the body from sliding down slopes when you include gravity in [code]linear_velocity[/code]. When set to lower values, the body will not be able to stand still on steep slopes. If the body collides, it will change direction a maximum of [code]max_slides[/code] times before it stops. [code]floor_max_angle[/code] is the maximum angle (in radians) where a slope is still considered a floor (or a ceiling), rather than a wall. The default value equals 45 degrees. Returns the movement that remained when the body stopped. To get more detailed information about collisions that occurred, use [method get_slide_collision].
	Args: [{ false linear_velocity Vector3} {(0, 0, 0) true floor_normal Vector3} {0.05 true slope_stop_min_velocity float} {4 true max_slides int} {0.785398 true floor_max_angle float}], Returns: Vector3
*/
func (o *KinematicBody) MoveAndSlide(linearVelocity gdnative.Vector3, floorNormal gdnative.Vector3, slopeStopMinVelocity gdnative.Real, maxSlides gdnative.Int, floorMaxAngle gdnative.Real) gdnative.Vector3 {
	//log.Println("Calling KinematicBody.MoveAndSlide()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromVector3(linearVelocity)
	ptrArguments[1] = gdnative.NewPointerFromVector3(floorNormal)
	ptrArguments[2] = gdnative.NewPointerFromReal(slopeStopMinVelocity)
	ptrArguments[3] = gdnative.NewPointerFromInt(maxSlides)
	ptrArguments[4] = gdnative.NewPointerFromReal(floorMaxAngle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "move_and_slide")

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
	Args: [{ false axis int} { false lock bool}], Returns: void
*/
func (o *KinematicBody) SetAxisLock(axis gdnative.Int, lock gdnative.Bool) {
	//log.Println("Calling KinematicBody.SetAxisLock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(axis)
	ptrArguments[1] = gdnative.NewPointerFromBool(lock)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "set_axis_lock")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pixels float}], Returns: void
*/
func (o *KinematicBody) SetSafeMargin(pixels gdnative.Real) {
	//log.Println("Calling KinematicBody.SetSafeMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(pixels)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "set_safe_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Checks for collisions without moving the body. Virtually sets the node's position, scale and rotation to that of the given [Transform], then tries to move the body along the vector [code]rel_vec[/code]. Returns [code]true[/code] if a collision would occur.
	Args: [{ false from Transform} { false rel_vec Vector3}], Returns: bool
*/
func (o *KinematicBody) TestMove(from gdnative.Transform, relVec gdnative.Vector3) gdnative.Bool {
	//log.Println("Calling KinematicBody.TestMove()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromTransform(from)
	ptrArguments[1] = gdnative.NewPointerFromVector3(relVec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("KinematicBody", "test_move")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// KinematicBodyImplementer is an interface that implements the methods
// of the KinematicBody class.
type KinematicBodyImplementer interface {
	PhysicsBodyImplementer
	GetAxisLock(axis gdnative.Int) gdnative.Bool
	GetFloorVelocity() gdnative.Vector3
	GetSafeMargin() gdnative.Real
	GetSlideCollision(slideIdx gdnative.Int) KinematicCollisionImplementer
	GetSlideCount() gdnative.Int
	IsOnCeiling() gdnative.Bool
	IsOnFloor() gdnative.Bool
	IsOnWall() gdnative.Bool
	MoveAndCollide(relVec gdnative.Vector3) KinematicCollisionImplementer
	MoveAndSlide(linearVelocity gdnative.Vector3, floorNormal gdnative.Vector3, slopeStopMinVelocity gdnative.Real, maxSlides gdnative.Int, floorMaxAngle gdnative.Real) gdnative.Vector3
	SetAxisLock(axis gdnative.Int, lock gdnative.Bool)
	SetSafeMargin(pixels gdnative.Real)
	TestMove(from gdnative.Transform, relVec gdnative.Vector3) gdnative.Bool
}

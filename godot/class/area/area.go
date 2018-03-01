package area

import (
	"log"
	"reflect"
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
3D area that detects [CollisionObject] nodes overlapping, entering, or exiting. Can also alter or override local physics parameters (gravity, damping).
*/
type Area struct {
	CollisionObject
}

func (o *Area) BaseClass() string {
	return "Area"
}

/*
   Undocumented
*/
func (o *Area) X_AreaEnterTree(id gdnative.Int) {
	log.Println("Calling Area.X_AreaEnterTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_area_enter_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) X_AreaExitTree(id gdnative.Int) {
	log.Println("Calling Area.X_AreaExitTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_area_exit_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) X_AreaInout(arg0 gdnative.Int, arg1 *RID, arg2 gdnative.Int, arg3 gdnative.Int, arg4 gdnative.Int) {
	log.Println("Calling Area.X_AreaInout()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)
	goArguments[2] = reflect.ValueOf(arg2)
	goArguments[3] = reflect.ValueOf(arg3)
	goArguments[4] = reflect.ValueOf(arg4)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_area_inout", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) X_BodyEnterTree(id gdnative.Int) {
	log.Println("Calling Area.X_BodyEnterTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_body_enter_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) X_BodyExitTree(id gdnative.Int) {
	log.Println("Calling Area.X_BodyExitTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_body_exit_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) X_BodyInout(arg0 gdnative.Int, arg1 *RID, arg2 gdnative.Int, arg3 gdnative.Int, arg4 gdnative.Int) {
	log.Println("Calling Area.X_BodyInout()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)
	goArguments[2] = reflect.ValueOf(arg2)
	goArguments[3] = reflect.ValueOf(arg3)
	goArguments[4] = reflect.ValueOf(arg4)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_body_inout", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) GetAngularDamp() gdnative.Float {
	log.Println("Calling Area.GetAngularDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_angular_damp", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetAudioBus() gdnative.String {
	log.Println("Calling Area.GetAudioBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_audio_bus", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetCollisionLayer() gdnative.Int {
	log.Println("Calling Area.GetCollisionLayer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_layer", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns an individual bit on the layer mask.
*/
func (o *Area) GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool {
	log.Println("Calling Area.GetCollisionLayerBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_layer_bit", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetCollisionMask() gdnative.Int {
	log.Println("Calling Area.GetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns an individual bit on the collision mask.
*/
func (o *Area) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	log.Println("Calling Area.GetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask_bit", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetGravity() gdnative.Float {
	log.Println("Calling Area.GetGravity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gravity", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetGravityDistanceScale() gdnative.Float {
	log.Println("Calling Area.GetGravityDistanceScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gravity_distance_scale", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetGravityVector() *Vector3 {
	log.Println("Calling Area.GetGravityVector()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gravity_vector", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetLinearDamp() gdnative.Float {
	log.Println("Calling Area.GetLinearDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_linear_damp", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a list of intersecting [code]Area[/code]s. For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (o *Area) GetOverlappingAreas() *Array {
	log.Println("Calling Area.GetOverlappingAreas()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_overlapping_areas", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a list of intersecting [PhysicsBody]s. For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (o *Area) GetOverlappingBodies() *Array {
	log.Println("Calling Area.GetOverlappingBodies()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_overlapping_bodies", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetPriority() gdnative.Float {
	log.Println("Calling Area.GetPriority()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_priority", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetReverbAmount() gdnative.Float {
	log.Println("Calling Area.GetReverbAmount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_reverb_amount", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetReverbBus() gdnative.String {
	log.Println("Calling Area.GetReverbBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_reverb_bus", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetReverbUniformity() gdnative.Float {
	log.Println("Calling Area.GetReverbUniformity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_reverb_uniformity", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) GetSpaceOverrideMode() gdnative.Int {
	log.Println("Calling Area.GetSpaceOverrideMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_space_override_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) IsGravityAPoint() gdnative.Bool {
	log.Println("Calling Area.IsGravityAPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_gravity_a_point", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) IsMonitorable() gdnative.Bool {
	log.Println("Calling Area.IsMonitorable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_monitorable", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) IsMonitoring() gdnative.Bool {
	log.Println("Calling Area.IsMonitoring()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_monitoring", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) IsOverridingAudioBus() gdnative.Bool {
	log.Println("Calling Area.IsOverridingAudioBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_overriding_audio_bus", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) IsUsingReverbBus() gdnative.Bool {
	log.Println("Calling Area.IsUsingReverbBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_using_reverb_bus", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If [code]true[/code] the given area overlaps the Area. Note that the result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
func (o *Area) OverlapsArea(area *Object) gdnative.Bool {
	log.Println("Calling Area.OverlapsArea()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(area)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "overlaps_area", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If [code]true[/code] the given body overlaps the Area. Note that the result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
func (o *Area) OverlapsBody(body *Object) gdnative.Bool {
	log.Println("Calling Area.OverlapsBody()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(body)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "overlaps_body", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Area) SetAngularDamp(angularDamp gdnative.Float) {
	log.Println("Calling Area.SetAngularDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angularDamp)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_angular_damp", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetAudioBus(name gdnative.String) {
	log.Println("Calling Area.SetAudioBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_audio_bus", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetAudioBusOverride(enable gdnative.Bool) {
	log.Println("Calling Area.SetAudioBusOverride()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_audio_bus_override", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetCollisionLayer(collisionLayer gdnative.Int) {
	log.Println("Calling Area.SetCollisionLayer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(collisionLayer)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_layer", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set/clear individual bits on the layer mask. This simplifies editing this [code]Area[code]'s layers.
*/
func (o *Area) SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool) {
	log.Println("Calling Area.SetCollisionLayerBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_layer_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetCollisionMask(collisionMask gdnative.Int) {
	log.Println("Calling Area.SetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(collisionMask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set/clear individual bits on the collision mask. This simplifies editing which [code]Area[/code] layers this [code]Area[/code] scans.
*/
func (o *Area) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	log.Println("Calling Area.SetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetGravity(gravity gdnative.Float) {
	log.Println("Calling Area.SetGravity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(gravity)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gravity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetGravityDistanceScale(distanceScale gdnative.Float) {
	log.Println("Calling Area.SetGravityDistanceScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(distanceScale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gravity_distance_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetGravityIsPoint(enable gdnative.Bool) {
	log.Println("Calling Area.SetGravityIsPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gravity_is_point", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetGravityVector(vector *Vector3) {
	log.Println("Calling Area.SetGravityVector()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(vector)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gravity_vector", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetLinearDamp(linearDamp gdnative.Float) {
	log.Println("Calling Area.SetLinearDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(linearDamp)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_linear_damp", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetMonitorable(enable gdnative.Bool) {
	log.Println("Calling Area.SetMonitorable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_monitorable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetMonitoring(enable gdnative.Bool) {
	log.Println("Calling Area.SetMonitoring()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_monitoring", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetPriority(priority gdnative.Float) {
	log.Println("Calling Area.SetPriority()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(priority)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_priority", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetReverbAmount(amount gdnative.Float) {
	log.Println("Calling Area.SetReverbAmount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_reverb_amount", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetReverbBus(name gdnative.String) {
	log.Println("Calling Area.SetReverbBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_reverb_bus", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetReverbUniformity(amount gdnative.Float) {
	log.Println("Calling Area.SetReverbUniformity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_reverb_uniformity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetSpaceOverrideMode(enable gdnative.Int) {
	log.Println("Calling Area.SetSpaceOverrideMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_space_override_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Area) SetUseReverbBus(enable gdnative.Bool) {
	log.Println("Calling Area.SetUseReverbBus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_reverb_bus", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AreaImplementer is an interface for Area objects.
*/
type AreaImplementer interface {
	Class
}
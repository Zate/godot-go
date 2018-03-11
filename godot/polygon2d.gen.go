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

//func NewPolygon2DFromPointer(ptr gdnative.Pointer) Polygon2D {
func newPolygon2DFromPointer(ptr gdnative.Pointer) Polygon2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Polygon2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A Polygon2D is defined by a set of points. Each point is connected to the next, with the final point being connected to the first, resulting in a closed polygon. Polygon2Ds can be filled with color (solid or gradient) or filled with a given texture.
*/
type Polygon2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Polygon2D) BaseClass() string {
	return "Polygon2D"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Polygon2D) GetAntialiased() gdnative.Bool {
	//log.Println("Calling Polygon2D.GetAntialiased()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_antialiased")

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
	Args: [], Returns: Color
*/
func (o *Polygon2D) GetColor() gdnative.Color {
	//log.Println("Calling Polygon2D.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Polygon2D) GetInvert() gdnative.Bool {
	//log.Println("Calling Polygon2D.GetInvert()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_invert")

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
	Args: [], Returns: float
*/
func (o *Polygon2D) GetInvertBorder() gdnative.Float {
	//log.Println("Calling Polygon2D.GetInvertBorder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_invert_border")

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
	Args: [], Returns: Vector2
*/
func (o *Polygon2D) GetOffset() gdnative.Vector2 {
	//log.Println("Calling Polygon2D.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: PoolVector2Array
*/
func (o *Polygon2D) GetPolygon() gdnative.PoolVector2Array {
	//log.Println("Calling Polygon2D.GetPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_polygon")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *Polygon2D) GetTexture() TextureImplementer {
	//log.Println("Calling Polygon2D.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_texture")

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
	Args: [], Returns: Vector2
*/
func (o *Polygon2D) GetTextureOffset() gdnative.Vector2 {
	//log.Println("Calling Polygon2D.GetTextureOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_texture_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Polygon2D) GetTextureRotation() gdnative.Float {
	//log.Println("Calling Polygon2D.GetTextureRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_texture_rotation")

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
func (o *Polygon2D) GetTextureRotationDegrees() gdnative.Float {
	//log.Println("Calling Polygon2D.GetTextureRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_texture_rotation_degrees")

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
	Args: [], Returns: Vector2
*/
func (o *Polygon2D) GetTextureScale() gdnative.Vector2 {
	//log.Println("Calling Polygon2D.GetTextureScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_texture_scale")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: PoolVector2Array
*/
func (o *Polygon2D) GetUv() gdnative.PoolVector2Array {
	//log.Println("Calling Polygon2D.GetUv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_uv")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: PoolColorArray
*/
func (o *Polygon2D) GetVertexColors() gdnative.PoolColorArray {
	//log.Println("Calling Polygon2D.GetVertexColors()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "get_vertex_colors")

	// Call the parent method.
	// PoolColorArray
	retPtr := gdnative.NewEmptyPoolColorArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolColorArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false antialiased bool}], Returns: void
*/
func (o *Polygon2D) SetAntialiased(antialiased gdnative.Bool) {
	//log.Println("Calling Polygon2D.SetAntialiased()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(antialiased)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_antialiased")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *Polygon2D) SetColor(color gdnative.Color) {
	//log.Println("Calling Polygon2D.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false invert bool}], Returns: void
*/
func (o *Polygon2D) SetInvert(invert gdnative.Bool) {
	//log.Println("Calling Polygon2D.SetInvert()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(invert)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_invert")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false invert_border float}], Returns: void
*/
func (o *Polygon2D) SetInvertBorder(invertBorder gdnative.Float) {
	//log.Println("Calling Polygon2D.SetInvertBorder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(invertBorder)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_invert_border")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *Polygon2D) SetOffset(offset gdnative.Vector2) {
	//log.Println("Calling Polygon2D.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false polygon PoolVector2Array}], Returns: void
*/
func (o *Polygon2D) SetPolygon(polygon gdnative.PoolVector2Array) {
	//log.Println("Calling Polygon2D.SetPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector2Array(polygon)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_polygon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *Polygon2D) SetTexture(texture Texture) {
	//log.Println("Calling Polygon2D.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture_offset Vector2}], Returns: void
*/
func (o *Polygon2D) SetTextureOffset(textureOffset gdnative.Vector2) {
	//log.Println("Calling Polygon2D.SetTextureOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(textureOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_texture_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture_rotation float}], Returns: void
*/
func (o *Polygon2D) SetTextureRotation(textureRotation gdnative.Float) {
	//log.Println("Calling Polygon2D.SetTextureRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(textureRotation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_texture_rotation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture_rotation float}], Returns: void
*/
func (o *Polygon2D) SetTextureRotationDegrees(textureRotation gdnative.Float) {
	//log.Println("Calling Polygon2D.SetTextureRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(textureRotation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_texture_rotation_degrees")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture_scale Vector2}], Returns: void
*/
func (o *Polygon2D) SetTextureScale(textureScale gdnative.Vector2) {
	//log.Println("Calling Polygon2D.SetTextureScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(textureScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_texture_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false uv PoolVector2Array}], Returns: void
*/
func (o *Polygon2D) SetUv(uv gdnative.PoolVector2Array) {
	//log.Println("Calling Polygon2D.SetUv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector2Array(uv)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_uv")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false vertex_colors PoolColorArray}], Returns: void
*/
func (o *Polygon2D) SetVertexColors(vertexColors gdnative.PoolColorArray) {
	//log.Println("Calling Polygon2D.SetVertexColors()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolColorArray(vertexColors)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Polygon2D", "set_vertex_colors")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Polygon2DImplementer is an interface that implements the methods
// of the Polygon2D class.
type Polygon2DImplementer interface {
	Node2DImplementer
	GetAntialiased() gdnative.Bool
	GetColor() gdnative.Color
	GetInvert() gdnative.Bool
	GetInvertBorder() gdnative.Float
	GetOffset() gdnative.Vector2
	GetPolygon() gdnative.PoolVector2Array
	GetTexture() TextureImplementer
	GetTextureOffset() gdnative.Vector2
	GetTextureRotation() gdnative.Float
	GetTextureRotationDegrees() gdnative.Float
	GetTextureScale() gdnative.Vector2
	GetUv() gdnative.PoolVector2Array
	GetVertexColors() gdnative.PoolColorArray
	SetAntialiased(antialiased gdnative.Bool)
	SetColor(color gdnative.Color)
	SetInvert(invert gdnative.Bool)
	SetInvertBorder(invertBorder gdnative.Float)
	SetOffset(offset gdnative.Vector2)
	SetPolygon(polygon gdnative.PoolVector2Array)
	SetTexture(texture Texture)
	SetTextureOffset(textureOffset gdnative.Vector2)
	SetTextureRotation(textureRotation gdnative.Float)
	SetTextureRotationDegrees(textureRotation gdnative.Float)
	SetTextureScale(textureScale gdnative.Vector2)
	SetUv(uv gdnative.PoolVector2Array)
	SetVertexColors(vertexColors gdnative.PoolColorArray)
}

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

//func NewLargeTextureFromPointer(ptr gdnative.Pointer) LargeTexture {
func newLargeTextureFromPointer(ptr gdnative.Pointer) LargeTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := LargeTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A Texture capable of storing many smaller Textures with offsets. You can dynamically add pieces([Texture]) to this [code]LargeTexture[/code] using different offsets.
*/
type LargeTexture struct {
	Texture
	owner gdnative.Object
}

func (o *LargeTexture) BaseClass() string {
	return "LargeTexture"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *LargeTexture) X_GetData() gdnative.Array {
	//log.Println("Calling LargeTexture.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "_get_data")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false data Array}], Returns: void
*/
func (o *LargeTexture) X_SetData(data gdnative.Array) {
	//log.Println("Calling LargeTexture.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add another [Texture] to this [code]LargeTexture[/code], starting on offset "ofs".
	Args: [{ false ofs Vector2} { false texture Texture}], Returns: int
*/
func (o *LargeTexture) AddPiece(ofs gdnative.Vector2, texture Texture) gdnative.Int {
	//log.Println("Calling LargeTexture.AddPiece()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector2(ofs)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "add_piece")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Clears the [code]LargeTexture[/code].
	Args: [], Returns: void
*/
func (o *LargeTexture) Clear() {
	//log.Println("Calling LargeTexture.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the number of pieces currently in this [code]LargeTexture[/code].
	Args: [], Returns: int
*/
func (o *LargeTexture) GetPieceCount() gdnative.Int {
	//log.Println("Calling LargeTexture.GetPieceCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "get_piece_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the offset of the piece with index "idx".
	Args: [{ false idx int}], Returns: Vector2
*/
func (o *LargeTexture) GetPieceOffset(idx gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling LargeTexture.GetPieceOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "get_piece_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the [Texture] of the piece with index "idx".
	Args: [{ false idx int}], Returns: Texture
*/
func (o *LargeTexture) GetPieceTexture(idx gdnative.Int) TextureImplementer {
	//log.Println("Calling LargeTexture.GetPieceTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "get_piece_texture")

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
        Sets the offset of the piece with index "idx" to "ofs".
	Args: [{ false idx int} { false ofs Vector2}], Returns: void
*/
func (o *LargeTexture) SetPieceOffset(idx gdnative.Int, ofs gdnative.Vector2) {
	//log.Println("Calling LargeTexture.SetPieceOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector2(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "set_piece_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the [Texture] of the piece with index "idx" to "ofs".
	Args: [{ false idx int} { false texture Texture}], Returns: void
*/
func (o *LargeTexture) SetPieceTexture(idx gdnative.Int, texture Texture) {
	//log.Println("Calling LargeTexture.SetPieceTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "set_piece_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the size of this [code]LargeTexture[/code].
	Args: [{ false size Vector2}], Returns: void
*/
func (o *LargeTexture) SetSize(size gdnative.Vector2) {
	//log.Println("Calling LargeTexture.SetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LargeTexture", "set_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// LargeTextureImplementer is an interface that implements the methods
// of the LargeTexture class.
type LargeTextureImplementer interface {
	TextureImplementer
	X_GetData() gdnative.Array
	X_SetData(data gdnative.Array)
	AddPiece(ofs gdnative.Vector2, texture Texture) gdnative.Int
	Clear()
	GetPieceCount() gdnative.Int
	GetPieceOffset(idx gdnative.Int) gdnative.Vector2
	GetPieceTexture(idx gdnative.Int) TextureImplementer
	SetPieceOffset(idx gdnative.Int, ofs gdnative.Vector2)
	SetPieceTexture(idx gdnative.Int, texture Texture)
	SetSize(size gdnative.Vector2)
}

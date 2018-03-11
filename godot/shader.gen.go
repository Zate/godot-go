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

//func NewShaderFromPointer(ptr gdnative.Pointer) Shader {
func newShaderFromPointer(ptr gdnative.Pointer) Shader {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Shader{}
	obj.SetBaseObject(owner)

	return obj
}

/*
To be changed, ignore.
*/
type Shader struct {
	Resource
	owner gdnative.Object
}

func (o *Shader) BaseClass() string {
	return "Shader"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *Shader) GetCode() gdnative.String {
	//log.Println("Calling Shader.GetCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shader", "get_code")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false param String}], Returns: Texture
*/
func (o *Shader) GetDefaultTextureParam(param gdnative.String) TextureImplementer {
	//log.Println("Calling Shader.GetDefaultTextureParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shader", "get_default_texture_param")

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

	Args: [], Returns: enum.Shader::Mode
*/

/*

	Args: [{ false name String}], Returns: bool
*/
func (o *Shader) HasParam(name gdnative.String) gdnative.Bool {
	//log.Println("Calling Shader.HasParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shader", "has_param")

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
	Args: [{ false code String}], Returns: void
*/
func (o *Shader) SetCode(code gdnative.String) {
	//log.Println("Calling Shader.SetCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(code)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shader", "set_code")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false param String} { false texture Texture}], Returns: void
*/
func (o *Shader) SetDefaultTextureParam(param gdnative.String, texture Texture) {
	//log.Println("Calling Shader.SetDefaultTextureParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(param)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shader", "set_default_texture_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ShaderImplementer is an interface that implements the methods
// of the Shader class.
type ShaderImplementer interface {
	ResourceImplementer
	GetCode() gdnative.String
	GetDefaultTextureParam(param gdnative.String) TextureImplementer
	HasParam(name gdnative.String) gdnative.Bool
	SetCode(code gdnative.String)
	SetDefaultTextureParam(param gdnative.String, texture Texture)
}

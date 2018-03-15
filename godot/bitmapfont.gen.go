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

//func NewBitmapFontFromPointer(ptr gdnative.Pointer) BitmapFont {
func newBitmapFontFromPointer(ptr gdnative.Pointer) BitmapFont {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BitmapFont{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Renders text using [code]*.fnt[/code] fonts containing texture atlases. Supports distance fields. For using vector font files like TTF directly, see [DynamicFont].
*/
type BitmapFont struct {
	Font
	owner gdnative.Object
}

func (o *BitmapFont) BaseClass() string {
	return "BitmapFont"
}

/*
        Undocumented
	Args: [], Returns: PoolIntArray
*/
func (o *BitmapFont) X_GetChars() gdnative.PoolIntArray {
	//log.Println("Calling BitmapFont.X_GetChars()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "_get_chars")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: PoolIntArray
*/
func (o *BitmapFont) X_GetKernings() gdnative.PoolIntArray {
	//log.Println("Calling BitmapFont.X_GetKernings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "_get_kernings")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *BitmapFont) X_GetTextures() gdnative.Array {
	//log.Println("Calling BitmapFont.X_GetTextures()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "_get_textures")

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
	Args: [{ false arg0 PoolIntArray}], Returns: void
*/
func (o *BitmapFont) X_SetChars(arg0 gdnative.PoolIntArray) {
	//log.Println("Calling BitmapFont.X_SetChars()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolIntArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "_set_chars")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 PoolIntArray}], Returns: void
*/
func (o *BitmapFont) X_SetKernings(arg0 gdnative.PoolIntArray) {
	//log.Println("Calling BitmapFont.X_SetKernings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolIntArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "_set_kernings")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Array}], Returns: void
*/
func (o *BitmapFont) X_SetTextures(arg0 gdnative.Array) {
	//log.Println("Calling BitmapFont.X_SetTextures()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "_set_textures")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a character to the font, where [code]character[/code] is the unicode value, [code]texture[/code] is the texture index, [code]rect[/code] is the region in the texture (in pixels!), [code]align[/code] is the (optional) alignment for the character and [code]advance[/code] is the (optional) advance.
	Args: [{ false character int} { false texture int} { false rect Rect2} {(0, 0) true align Vector2} {-1 true advance float}], Returns: void
*/
func (o *BitmapFont) AddChar(character gdnative.Int, texture gdnative.Int, rect gdnative.Rect2, align gdnative.Vector2, advance gdnative.Real) {
	//log.Println("Calling BitmapFont.AddChar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(character)
	ptrArguments[1] = gdnative.NewPointerFromInt(texture)
	ptrArguments[2] = gdnative.NewPointerFromRect2(rect)
	ptrArguments[3] = gdnative.NewPointerFromVector2(align)
	ptrArguments[4] = gdnative.NewPointerFromReal(advance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "add_char")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a kerning pair to the [code]BitmapFont[/code] as a difference. Kerning pairs are special cases where a typeface advance is determined by the next character.
	Args: [{ false char_a int} { false char_b int} { false kerning int}], Returns: void
*/
func (o *BitmapFont) AddKerningPair(charA gdnative.Int, charB gdnative.Int, kerning gdnative.Int) {
	//log.Println("Calling BitmapFont.AddKerningPair()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(charA)
	ptrArguments[1] = gdnative.NewPointerFromInt(charB)
	ptrArguments[2] = gdnative.NewPointerFromInt(kerning)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "add_kerning_pair")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a texture to the [code]BitmapFont[/code].
	Args: [{ false texture Texture}], Returns: void
*/
func (o *BitmapFont) AddTexture(texture TextureImplementer) {
	//log.Println("Calling BitmapFont.AddTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "add_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears all the font data and settings.
	Args: [], Returns: void
*/
func (o *BitmapFont) Clear() {
	//log.Println("Calling BitmapFont.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a BitmapFont from the [code]*.fnt[/code] file at [code]path[/code].
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *BitmapFont) CreateFromFnt(path gdnative.String) gdnative.Error {
	//log.Println("Calling BitmapFont.CreateFromFnt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "create_from_fnt")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Returns the size of a character, optionally taking kerning into account if the next character is provided.
	Args: [{ false char int} {0 true next int}], Returns: Vector2
*/
func (o *BitmapFont) GetCharSize(char gdnative.Int, next gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling BitmapFont.GetCharSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(char)
	ptrArguments[1] = gdnative.NewPointerFromInt(next)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "get_char_size")

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
	Args: [], Returns: BitmapFont
*/
func (o *BitmapFont) GetFallback() BitmapFontImplementer {
	//log.Println("Calling BitmapFont.GetFallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "get_fallback")

	// Call the parent method.
	// BitmapFont
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newBitmapFontFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(BitmapFontImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "BitmapFont" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(BitmapFontImplementer)
	}

	return &ret
}

/*
        Returns a kerning pair as a difference.
	Args: [{ false char_a int} { false char_b int}], Returns: int
*/
func (o *BitmapFont) GetKerningPair(charA gdnative.Int, charB gdnative.Int) gdnative.Int {
	//log.Println("Calling BitmapFont.GetKerningPair()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(charA)
	ptrArguments[1] = gdnative.NewPointerFromInt(charB)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "get_kerning_pair")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the font atlas texture at index [code]idx[/code].
	Args: [{ false idx int}], Returns: Texture
*/
func (o *BitmapFont) GetTexture(idx gdnative.Int) TextureImplementer {
	//log.Println("Calling BitmapFont.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "get_texture")

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
        Returns the number of textures in the BitmapFont atlas.
	Args: [], Returns: int
*/
func (o *BitmapFont) GetTextureCount() gdnative.Int {
	//log.Println("Calling BitmapFont.GetTextureCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "get_texture_count")

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
	Args: [{ false px float}], Returns: void
*/
func (o *BitmapFont) SetAscent(px gdnative.Real) {
	//log.Println("Calling BitmapFont.SetAscent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(px)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "set_ascent")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *BitmapFont) SetDistanceFieldHint(enable gdnative.Bool) {
	//log.Println("Calling BitmapFont.SetDistanceFieldHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "set_distance_field_hint")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false fallback BitmapFont}], Returns: void
*/
func (o *BitmapFont) SetFallback(fallback BitmapFontImplementer) {
	//log.Println("Calling BitmapFont.SetFallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(fallback.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "set_fallback")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false px float}], Returns: void
*/
func (o *BitmapFont) SetHeight(px gdnative.Real) {
	//log.Println("Calling BitmapFont.SetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(px)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitmapFont", "set_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// BitmapFontImplementer is an interface that implements the methods
// of the BitmapFont class.
type BitmapFontImplementer interface {
	FontImplementer
	X_GetChars() gdnative.PoolIntArray
	X_GetKernings() gdnative.PoolIntArray
	X_GetTextures() gdnative.Array
	X_SetChars(arg0 gdnative.PoolIntArray)
	X_SetKernings(arg0 gdnative.PoolIntArray)
	X_SetTextures(arg0 gdnative.Array)
	AddChar(character gdnative.Int, texture gdnative.Int, rect gdnative.Rect2, align gdnative.Vector2, advance gdnative.Real)
	AddKerningPair(charA gdnative.Int, charB gdnative.Int, kerning gdnative.Int)
	AddTexture(texture TextureImplementer)
	Clear()
	GetCharSize(char gdnative.Int, next gdnative.Int) gdnative.Vector2
	GetFallback() BitmapFontImplementer
	GetKerningPair(charA gdnative.Int, charB gdnative.Int) gdnative.Int
	GetTexture(idx gdnative.Int) TextureImplementer
	GetTextureCount() gdnative.Int
	SetAscent(px gdnative.Real)
	SetDistanceFieldHint(enable gdnative.Bool)
	SetFallback(fallback BitmapFontImplementer)
	SetHeight(px gdnative.Real)
}

package linkbutton

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
This kind of buttons are primarily used when the interaction with the button causes a context change (like linking to a web page).
*/
type LinkButton struct {
	BaseButton
}

func (o *LinkButton) BaseClass() string {
	return "LinkButton"
}

/*
   Undocumented
*/
func (o *LinkButton) GetText() gdnative.String {
	log.Println("Calling LinkButton.GetText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_text", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *LinkButton) GetUnderlineMode() gdnative.Int {
	log.Println("Calling LinkButton.GetUnderlineMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_underline_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *LinkButton) SetText(text gdnative.String) {
	log.Println("Calling LinkButton.SetText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(text)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_text", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *LinkButton) SetUnderlineMode(underlineMode gdnative.Int) {
	log.Println("Calling LinkButton.SetUnderlineMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(underlineMode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_underline_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   LinkButtonImplementer is an interface for LinkButton objects.
*/
type LinkButtonImplementer interface {
	Class
}
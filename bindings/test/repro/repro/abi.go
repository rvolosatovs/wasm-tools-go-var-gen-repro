// Code generated by wit-bindgen-go. DO NOT EDIT.

package repro

func lower_Var(v Var) (f0 uint32, f1 uint32) {
	f0 = (uint32)(v.Tag())
	switch f0 {
	case 0: // var
		v1 := (uint32)(*v.Var())
		f1 = (uint32)(v1)
	}
	return
}

//! This file was generated automatically by the Snowball to Go compiler
//! http://snowballstem.org/

package porter

import (
	snowballRuntime "github.com/blevesearch/snowballstem"
)

var A_0 = []*snowballRuntime.Among{
	{Str: "s", A: -1, B: 3, F: nil},
	{Str: "ies", A: 0, B: 2, F: nil},
	{Str: "sses", A: 0, B: 1, F: nil},
	{Str: "ss", A: 0, B: -1, F: nil},
}

var A_1 = []*snowballRuntime.Among{
	{Str: "", A: -1, B: 3, F: nil},
	{Str: "bb", A: 0, B: 2, F: nil},
	{Str: "dd", A: 0, B: 2, F: nil},
	{Str: "ff", A: 0, B: 2, F: nil},
	{Str: "gg", A: 0, B: 2, F: nil},
	{Str: "bl", A: 0, B: 1, F: nil},
	{Str: "mm", A: 0, B: 2, F: nil},
	{Str: "nn", A: 0, B: 2, F: nil},
	{Str: "pp", A: 0, B: 2, F: nil},
	{Str: "rr", A: 0, B: 2, F: nil},
	{Str: "at", A: 0, B: 1, F: nil},
	{Str: "tt", A: 0, B: 2, F: nil},
	{Str: "iz", A: 0, B: 1, F: nil},
}

var A_2 = []*snowballRuntime.Among{
	{Str: "ed", A: -1, B: 2, F: nil},
	{Str: "eed", A: 0, B: 1, F: nil},
	{Str: "ing", A: -1, B: 2, F: nil},
}

var A_3 = []*snowballRuntime.Among{
	{Str: "anci", A: -1, B: 3, F: nil},
	{Str: "enci", A: -1, B: 2, F: nil},
	{Str: "abli", A: -1, B: 4, F: nil},
	{Str: "eli", A: -1, B: 6, F: nil},
	{Str: "alli", A: -1, B: 9, F: nil},
	{Str: "ousli", A: -1, B: 12, F: nil},
	{Str: "entli", A: -1, B: 5, F: nil},
	{Str: "aliti", A: -1, B: 10, F: nil},
	{Str: "biliti", A: -1, B: 14, F: nil},
	{Str: "iviti", A: -1, B: 13, F: nil},
	{Str: "tional", A: -1, B: 1, F: nil},
	{Str: "ational", A: 10, B: 8, F: nil},
	{Str: "alism", A: -1, B: 10, F: nil},
	{Str: "ation", A: -1, B: 8, F: nil},
	{Str: "ization", A: 13, B: 7, F: nil},
	{Str: "izer", A: -1, B: 7, F: nil},
	{Str: "ator", A: -1, B: 8, F: nil},
	{Str: "iveness", A: -1, B: 13, F: nil},
	{Str: "fulness", A: -1, B: 11, F: nil},
	{Str: "ousness", A: -1, B: 12, F: nil},
}

var A_4 = []*snowballRuntime.Among{
	{Str: "icate", A: -1, B: 2, F: nil},
	{Str: "ative", A: -1, B: 3, F: nil},
	{Str: "alize", A: -1, B: 1, F: nil},
	{Str: "iciti", A: -1, B: 2, F: nil},
	{Str: "ical", A: -1, B: 2, F: nil},
	{Str: "ful", A: -1, B: 3, F: nil},
	{Str: "ness", A: -1, B: 3, F: nil},
}

var A_5 = []*snowballRuntime.Among{
	{Str: "ic", A: -1, B: 1, F: nil},
	{Str: "ance", A: -1, B: 1, F: nil},
	{Str: "ence", A: -1, B: 1, F: nil},
	{Str: "able", A: -1, B: 1, F: nil},
	{Str: "ible", A: -1, B: 1, F: nil},
	{Str: "ate", A: -1, B: 1, F: nil},
	{Str: "ive", A: -1, B: 1, F: nil},
	{Str: "ize", A: -1, B: 1, F: nil},
	{Str: "iti", A: -1, B: 1, F: nil},
	{Str: "al", A: -1, B: 1, F: nil},
	{Str: "ism", A: -1, B: 1, F: nil},
	{Str: "ion", A: -1, B: 2, F: nil},
	{Str: "er", A: -1, B: 1, F: nil},
	{Str: "ous", A: -1, B: 1, F: nil},
	{Str: "ant", A: -1, B: 1, F: nil},
	{Str: "ent", A: -1, B: 1, F: nil},
	{Str: "ment", A: 15, B: 1, F: nil},
	{Str: "ement", A: 16, B: 1, F: nil},
	{Str: "ou", A: -1, B: 1, F: nil},
}

var G_v = []byte{17, 65, 16, 1}

var G_v_WXY = []byte{1, 17, 65, 208, 1}

type Context struct {
	b_Y_found bool
	i_p2      int
	i_p1      int
}

func r_shortv(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	// (, line 19
	if !env.OutGroupingB(G_v_WXY, 89, 121) {
		return false
	}
	if !env.InGroupingB(G_v, 97, 121) {
		return false
	}
	if !env.OutGroupingB(G_v, 97, 121) {
		return false
	}
	return true
}

func r_R1(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	if !(context.i_p1 <= env.Cursor) {
		return false
	}
	return true
}

func r_R2(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	if !(context.i_p2 <= env.Cursor) {
		return false
	}
	return true
}

func r_Step_1a(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var among_var int32
	// (, line 24
	// [, line 25
	env.Ket = env.Cursor
	// substring, line 25
	among_var = env.FindAmongB(A_0, context)
	if among_var == 0 {
		return false
	}
	// ], line 25
	env.Bra = env.Cursor
	if among_var == 0 {
		return false
	} else if among_var == 1 {
		// (, line 26
		// <-, line 26
		if !env.SliceFrom("ss") {
			return false
		}
	} else if among_var == 2 {
		// (, line 27
		// <-, line 27
		if !env.SliceFrom("i") {
			return false
		}
	} else if among_var == 3 {
		// (, line 29
		// delete, line 29
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

func r_Step_1b(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var among_var int32
	// (, line 33
	// [, line 34
	env.Ket = env.Cursor
	// substring, line 34
	among_var = env.FindAmongB(A_2, context)
	if among_var == 0 {
		return false
	}
	// ], line 34
	env.Bra = env.Cursor
	if among_var == 0 {
		return false
	} else if among_var == 1 {
		// (, line 35
		// call R1, line 35
		if !r_R1(env, context) {
			return false
		}
		// <-, line 35
		if !env.SliceFrom("ee") {
			return false
		}
	} else if among_var == 2 {
		// (, line 37
		// test, line 38
		var v_1 = env.Limit - env.Cursor
		// gopast, line 38
	golab0:
		for {
		lab1:
			for {
				if !env.InGroupingB(G_v, 97, 121) {
					break lab1
				}
				break golab0
			}
			if env.Cursor <= env.LimitBackward {
				return false
			}
			env.PrevChar()
		}
		env.Cursor = env.Limit - v_1
		// delete, line 38
		if !env.SliceDel() {
			return false
		}
		// test, line 39
		var v_3 = env.Limit - env.Cursor
		// substring, line 39
		among_var = env.FindAmongB(A_1, context)
		if among_var == 0 {
			return false
		}
		env.Cursor = env.Limit - v_3
		if among_var == 0 {
			return false
		} else if among_var == 1 {
			// (, line 41
			{
				// <+, line 41
				var c = env.Cursor
				bra, ket := env.Cursor, env.Cursor
				env.Insert(bra, ket, "e")
				env.Cursor = c
			}
		} else if among_var == 2 {
			// (, line 44
			// [, line 44
			env.Ket = env.Cursor
			// next, line 44
			if env.Cursor <= env.LimitBackward {
				return false
			}
			env.PrevChar()
			// ], line 44
			env.Bra = env.Cursor
			// delete, line 44
			if !env.SliceDel() {
				return false
			}
		} else if among_var == 3 {
			// (, line 45
			// atmark, line 45
			if env.Cursor != context.i_p1 {
				return false
			}
			// test, line 45
			var v_4 = env.Limit - env.Cursor
			// call shortv, line 45
			if !r_shortv(env, context) {
				return false
			}
			env.Cursor = env.Limit - v_4
			{
				// <+, line 45
				var c = env.Cursor
				bra, ket := env.Cursor, env.Cursor
				env.Insert(bra, ket, "e")
				env.Cursor = c
			}
		}
	}
	return true
}

func r_Step_1c(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	// (, line 51
	// [, line 52
	env.Ket = env.Cursor
	// or, line 52
lab0:
	for {
		var v_1 = env.Limit - env.Cursor
	lab1:
		for {
			// literal, line 52
			if !env.EqSB("y") {
				break lab1
			}
			break lab0
		}
		env.Cursor = env.Limit - v_1
		// literal, line 52
		if !env.EqSB("Y") {
			return false
		}
		break lab0
	}
	// ], line 52
	env.Bra = env.Cursor
	// gopast, line 53
golab2:
	for {
	lab3:
		for {
			if !env.InGroupingB(G_v, 97, 121) {
				break lab3
			}
			break golab2
		}
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PrevChar()
	}
	// <-, line 54
	if !env.SliceFrom("i") {
		return false
	}
	return true
}

func r_Step_2(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var among_var int32
	// (, line 57
	// [, line 58
	env.Ket = env.Cursor
	// substring, line 58
	among_var = env.FindAmongB(A_3, context)
	if among_var == 0 {
		return false
	}
	// ], line 58
	env.Bra = env.Cursor
	// call R1, line 58
	if !r_R1(env, context) {
		return false
	}
	if among_var == 0 {
		return false
	} else if among_var == 1 {
		// (, line 59
		// <-, line 59
		if !env.SliceFrom("tion") {
			return false
		}
	} else if among_var == 2 {
		// (, line 60
		// <-, line 60
		if !env.SliceFrom("ence") {
			return false
		}
	} else if among_var == 3 {
		// (, line 61
		// <-, line 61
		if !env.SliceFrom("ance") {
			return false
		}
	} else if among_var == 4 {
		// (, line 62
		// <-, line 62
		if !env.SliceFrom("able") {
			return false
		}
	} else if among_var == 5 {
		// (, line 63
		// <-, line 63
		if !env.SliceFrom("ent") {
			return false
		}
	} else if among_var == 6 {
		// (, line 64
		// <-, line 64
		if !env.SliceFrom("e") {
			return false
		}
	} else if among_var == 7 {
		// (, line 66
		// <-, line 66
		if !env.SliceFrom("ize") {
			return false
		}
	} else if among_var == 8 {
		// (, line 68
		// <-, line 68
		if !env.SliceFrom("ate") {
			return false
		}
	} else if among_var == 9 {
		// (, line 69
		// <-, line 69
		if !env.SliceFrom("al") {
			return false
		}
	} else if among_var == 10 {
		// (, line 71
		// <-, line 71
		if !env.SliceFrom("al") {
			return false
		}
	} else if among_var == 11 {
		// (, line 72
		// <-, line 72
		if !env.SliceFrom("ful") {
			return false
		}
	} else if among_var == 12 {
		// (, line 74
		// <-, line 74
		if !env.SliceFrom("ous") {
			return false
		}
	} else if among_var == 13 {
		// (, line 76
		// <-, line 76
		if !env.SliceFrom("ive") {
			return false
		}
	} else if among_var == 14 {
		// (, line 77
		// <-, line 77
		if !env.SliceFrom("ble") {
			return false
		}
	}
	return true
}

func r_Step_3(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var among_var int32
	// (, line 81
	// [, line 82
	env.Ket = env.Cursor
	// substring, line 82
	among_var = env.FindAmongB(A_4, context)
	if among_var == 0 {
		return false
	}
	// ], line 82
	env.Bra = env.Cursor
	// call R1, line 82
	if !r_R1(env, context) {
		return false
	}
	if among_var == 0 {
		return false
	} else if among_var == 1 {
		// (, line 83
		// <-, line 83
		if !env.SliceFrom("al") {
			return false
		}
	} else if among_var == 2 {
		// (, line 85
		// <-, line 85
		if !env.SliceFrom("ic") {
			return false
		}
	} else if among_var == 3 {
		// (, line 87
		// delete, line 87
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

func r_Step_4(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var among_var int32
	// (, line 91
	// [, line 92
	env.Ket = env.Cursor
	// substring, line 92
	among_var = env.FindAmongB(A_5, context)
	if among_var == 0 {
		return false
	}
	// ], line 92
	env.Bra = env.Cursor
	// call R2, line 92
	if !r_R2(env, context) {
		return false
	}
	if among_var == 0 {
		return false
	} else if among_var == 1 {
		// (, line 95
		// delete, line 95
		if !env.SliceDel() {
			return false
		}
	} else if among_var == 2 {
		// (, line 96
		// or, line 96
	lab0:
		for {
			var v_1 = env.Limit - env.Cursor
		lab1:
			for {
				// literal, line 96
				if !env.EqSB("s") {
					break lab1
				}
				break lab0
			}
			env.Cursor = env.Limit - v_1
			// literal, line 96
			if !env.EqSB("t") {
				return false
			}
			break lab0
		}
		// delete, line 96
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

func r_Step_5a(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	// (, line 100
	// [, line 101
	env.Ket = env.Cursor
	// literal, line 101
	if !env.EqSB("e") {
		return false
	}
	// ], line 101
	env.Bra = env.Cursor
	// or, line 102
lab0:
	for {
		var v_1 = env.Limit - env.Cursor
	lab1:
		for {
			// call R2, line 102
			if !r_R2(env, context) {
				break lab1
			}
			break lab0
		}
		env.Cursor = env.Limit - v_1
		// (, line 102
		// call R1, line 102
		if !r_R1(env, context) {
			return false
		}
		// not, line 102
		var v_2 = env.Limit - env.Cursor
	lab2:
		for {
			// call shortv, line 102
			if !r_shortv(env, context) {
				break lab2
			}
			return false
		}
		env.Cursor = env.Limit - v_2
		break lab0
	}
	// delete, line 103
	if !env.SliceDel() {
		return false
	}
	return true
}

func r_Step_5b(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	// (, line 106
	// [, line 107
	env.Ket = env.Cursor
	// literal, line 107
	if !env.EqSB("l") {
		return false
	}
	// ], line 107
	env.Bra = env.Cursor
	// call R2, line 108
	if !r_R2(env, context) {
		return false
	}
	// literal, line 108
	if !env.EqSB("l") {
		return false
	}
	// delete, line 109
	if !env.SliceDel() {
		return false
	}
	return true
}

func Stem(env *snowballRuntime.Env) bool {
	var context = &Context{
		b_Y_found: false,
		i_p2:      0,
		i_p1:      0,
	}
	_ = context
	// (, line 113
	// unset Y_found, line 115
	context.b_Y_found = false
	// do, line 116
	var v_1 = env.Cursor
lab0:
	for {
		// (, line 116
		// [, line 116
		env.Bra = env.Cursor
		// literal, line 116
		if !env.EqS("y") {
			break lab0
		}
		// ], line 116
		env.Ket = env.Cursor
		// <-, line 116
		if !env.SliceFrom("Y") {
			return false
		}
		// set Y_found, line 116
		context.b_Y_found = true
		break lab0
	}
	env.Cursor = v_1
	// do, line 117
	var v_2 = env.Cursor
lab1:
	for {
		// repeat, line 117
	replab2:
		for {
			var v_3 = env.Cursor
		lab3:
			for range [2]struct{}{} {
				// (, line 117
				// goto, line 117
			golab4:
				for {
					var v_4 = env.Cursor
				lab5:
					for {
						// (, line 117
						if !env.InGrouping(G_v, 97, 121) {
							break lab5
						}
						// [, line 117
						env.Bra = env.Cursor
						// literal, line 117
						if !env.EqS("y") {
							break lab5
						}
						// ], line 117
						env.Ket = env.Cursor
						env.Cursor = v_4
						break golab4
					}
					env.Cursor = v_4
					if env.Cursor >= env.Limit {
						break lab3
					}
					env.NextChar()
				}
				// <-, line 117
				if !env.SliceFrom("Y") {
					return false
				}
				// set Y_found, line 117
				context.b_Y_found = true
				continue replab2
			}
			env.Cursor = v_3
			break replab2
		}
		break lab1
	}
	env.Cursor = v_2
	context.i_p1 = env.Limit
	context.i_p2 = env.Limit
	// do, line 121
	var v_5 = env.Cursor
lab6:
	for {
		// (, line 121
		// gopast, line 122
	golab7:
		for {
		lab8:
			for {
				if !env.InGrouping(G_v, 97, 121) {
					break lab8
				}
				break golab7
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
		// gopast, line 122
	golab9:
		for {
		lab10:
			for {
				if !env.OutGrouping(G_v, 97, 121) {
					break lab10
				}
				break golab9
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
		// setmark p1, line 122
		context.i_p1 = env.Cursor
		// gopast, line 123
	golab11:
		for {
		lab12:
			for {
				if !env.InGrouping(G_v, 97, 121) {
					break lab12
				}
				break golab11
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
		// gopast, line 123
	golab13:
		for {
		lab14:
			for {
				if !env.OutGrouping(G_v, 97, 121) {
					break lab14
				}
				break golab13
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
		// setmark p2, line 123
		context.i_p2 = env.Cursor
		break lab6
	}
	env.Cursor = v_5
	// backwards, line 126
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	// (, line 126
	// do, line 127
	var v_10 = env.Limit - env.Cursor
lab15:
	for {
		// call Step_1a, line 127
		if !r_Step_1a(env, context) {
			break lab15
		}
		break lab15
	}
	env.Cursor = env.Limit - v_10
	// do, line 128
	var v_11 = env.Limit - env.Cursor
lab16:
	for {
		// call Step_1b, line 128
		if !r_Step_1b(env, context) {
			break lab16
		}
		break lab16
	}
	env.Cursor = env.Limit - v_11
	// do, line 129
	var v_12 = env.Limit - env.Cursor
lab17:
	for {
		// call Step_1c, line 129
		if !r_Step_1c(env, context) {
			break lab17
		}
		break lab17
	}
	env.Cursor = env.Limit - v_12
	// do, line 130
	var v_13 = env.Limit - env.Cursor
lab18:
	for {
		// call Step_2, line 130
		if !r_Step_2(env, context) {
			break lab18
		}
		break lab18
	}
	env.Cursor = env.Limit - v_13
	// do, line 131
	var v_14 = env.Limit - env.Cursor
lab19:
	for {
		// call Step_3, line 131
		if !r_Step_3(env, context) {
			break lab19
		}
		break lab19
	}
	env.Cursor = env.Limit - v_14
	// do, line 132
	var v_15 = env.Limit - env.Cursor
lab20:
	for {
		// call Step_4, line 132
		if !r_Step_4(env, context) {
			break lab20
		}
		break lab20
	}
	env.Cursor = env.Limit - v_15
	// do, line 133
	var v_16 = env.Limit - env.Cursor
lab21:
	for {
		// call Step_5a, line 133
		if !r_Step_5a(env, context) {
			break lab21
		}
		break lab21
	}
	env.Cursor = env.Limit - v_16
	// do, line 134
	var v_17 = env.Limit - env.Cursor
lab22:
	for {
		// call Step_5b, line 134
		if !r_Step_5b(env, context) {
			break lab22
		}
		break lab22
	}
	env.Cursor = env.Limit - v_17
	env.Cursor = env.LimitBackward
	// do, line 137
	var v_18 = env.Cursor
lab23:
	for {
		// (, line 137
		// Boolean test Y_found, line 137
		if !context.b_Y_found {
			break lab23
		}
		// repeat, line 137
	replab24:
		for {
			var v_19 = env.Cursor
		lab25:
			for range [2]struct{}{} {
				// (, line 137
				// goto, line 137
			golab26:
				for {
					var v_20 = env.Cursor
				lab27:
					for {
						// (, line 137
						// [, line 137
						env.Bra = env.Cursor
						// literal, line 137
						if !env.EqS("Y") {
							break lab27
						}
						// ], line 137
						env.Ket = env.Cursor
						env.Cursor = v_20
						break golab26
					}
					env.Cursor = v_20
					if env.Cursor >= env.Limit {
						break lab25
					}
					env.NextChar()
				}
				// <-, line 137
				if !env.SliceFrom("y") {
					return false
				}
				continue replab24
			}
			env.Cursor = v_19
			break replab24
		}
		break lab23
	}
	env.Cursor = v_18
	return true
}

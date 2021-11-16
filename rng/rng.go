/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: rng.i

package rng

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L ../bin -lpractrand

#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_rng_bfbcb6a9797b2cf1(swig_intgo arg1);
extern void _wrap_delete_StateWalkingObject_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern void _wrap_StateWalkingObject_handle__SWIG_0_rng_bfbcb6a9797b2cf1(uintptr_t arg1, swig_voidp arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_1_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_2_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_3_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_4_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_5_rng_bfbcb6a9797b2cf1(uintptr_t arg1, swig_voidp arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_6_rng_bfbcb6a9797b2cf1(uintptr_t arg1, swig_voidp arg2);
extern swig_intgo _wrap_FLAG_READ_ONLY_StateWalkingObject_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_FLAG_WRITE_ONLY_StateWalkingObject_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_FLAG_CLUMSY_StateWalkingObject_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_FLAG_SEEDER_StateWalkingObject_rng_bfbcb6a9797b2cf1(void);
extern uintptr_t _wrap_StateWalkingObject_get_properties_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern _Bool _wrap_StateWalkingObject_is_read_only_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern _Bool _wrap_StateWalkingObject_is_write_only_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern _Bool _wrap_StateWalkingObject_is_clumsy_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern _Bool _wrap_StateWalkingObject_is_seeder_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern void _wrap_StateWalkingObject_handle__SWIG_7_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_8_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_9_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_StateWalkingObject_handle__SWIG_10_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_0_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_1_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_2_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_3_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_4_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_5_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_6_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_7_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_8_rng_bfbcb6a9797b2cf1(uintptr_t arg1, swig_voidp arg2);
extern uintptr_t _wrap_StateWalkingObject_LeftShift__SWIG_9_rng_bfbcb6a9797b2cf1(uintptr_t arg1, swig_voidp arg2);
extern uintptr_t _wrap_vrng_to_rng_seeder_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern uintptr_t _wrap_get_autoseeder_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern swig_intgo _wrap_OUTPUT_TYPE_salsa_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_OUTPUT_BITS_salsa_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_FLAGS_salsa_rng_bfbcb6a9797b2cf1(void);
extern uintptr_t _wrap_new_salsa_rng_bfbcb6a9797b2cf1(void);
extern void _wrap_delete_salsa_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern uintptr_t _wrap_salsa_raw32_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern void _wrap_salsa_seed__SWIG_0_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_salsa_seed__SWIG_1_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2, _Bool arg3);
extern void _wrap_salsa_seed__SWIG_2_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_salsa_seed_short__SWIG_0_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2, _Bool arg3);
extern void _wrap_salsa_seed_short__SWIG_1_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_salsa_walk_state_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_salsa_seek_forward_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3);
extern void _wrap_salsa_seek_backward_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3);
extern void _wrap_salsa_set_rounds_rng_bfbcb6a9797b2cf1(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_salsa_get_rounds_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern void _wrap_salsa_self_test_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_OUTPUT_TYPE_sfc16_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_OUTPUT_BITS_sfc16_rng_bfbcb6a9797b2cf1(void);
extern swig_intgo _wrap_FLAGS_sfc16_rng_bfbcb6a9797b2cf1(void);
extern uintptr_t _wrap_sfc16_raw16_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
extern void _wrap_sfc16_seed__SWIG_0_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_sfc16_seed_fast_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_sfc16_seed__SWIG_1_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3, uintptr_t arg4);
extern void _wrap_sfc16_walk_state_rng_bfbcb6a9797b2cf1(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_new_sfc16_rng_bfbcb6a9797b2cf1(void);
extern void _wrap_delete_sfc16_rng_bfbcb6a9797b2cf1(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_rng_bfbcb6a9797b2cf1(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrStateWalkingObject uintptr

func (p SwigcptrStateWalkingObject) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrStateWalkingObject) SwigIsStateWalkingObject() {
}

func DeleteStateWalkingObject(arg1 StateWalkingObject) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_StateWalkingObject_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_0(arg2 *bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_StateWalkingObject_handle__SWIG_0_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_1(arg2 Uint8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_1_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_2(arg2 Uint16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_2_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_3(arg2 Uint32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_3_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_4(arg2 Uint64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_4_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_5(arg2 *float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_StateWalkingObject_handle__SWIG_5_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_6(arg2 *float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_StateWalkingObject_handle__SWIG_6_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func _swig_getStateWalkingObject_FLAG_READ_ONLY_StateWalkingObject() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_FLAG_READ_ONLY_StateWalkingObject_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var StateWalkingObjectFLAG_READ_ONLY int = _swig_getStateWalkingObject_FLAG_READ_ONLY_StateWalkingObject()
func _swig_getStateWalkingObject_FLAG_WRITE_ONLY_StateWalkingObject() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_FLAG_WRITE_ONLY_StateWalkingObject_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var StateWalkingObjectFLAG_WRITE_ONLY int = _swig_getStateWalkingObject_FLAG_WRITE_ONLY_StateWalkingObject()
func _swig_getStateWalkingObject_FLAG_CLUMSY_StateWalkingObject() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_FLAG_CLUMSY_StateWalkingObject_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var StateWalkingObjectFLAG_CLUMSY int = _swig_getStateWalkingObject_FLAG_CLUMSY_StateWalkingObject()
func _swig_getStateWalkingObject_FLAG_SEEDER_StateWalkingObject() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_FLAG_SEEDER_StateWalkingObject_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var StateWalkingObjectFLAG_SEEDER int = _swig_getStateWalkingObject_FLAG_SEEDER_StateWalkingObject()
func (arg1 SwigcptrStateWalkingObject) Get_properties() (_swig_ret Uint32) {
	var swig_r Uint32
	_swig_i_0 := arg1
	swig_r = (Uint32)(SwigcptrUint32(C._wrap_StateWalkingObject_get_properties_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) Is_read_only() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_StateWalkingObject_is_read_only_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) Is_write_only() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_StateWalkingObject_is_write_only_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) Is_clumsy() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_StateWalkingObject_is_clumsy_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) Is_seeder() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_StateWalkingObject_is_seeder_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_7(arg2 Sint8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_7_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_8(arg2 Sint16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_8_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_9(arg2 Sint32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_9_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrStateWalkingObject) Handle__SWIG_10(arg2 Sint64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_StateWalkingObject_handle__SWIG_10_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (p SwigcptrStateWalkingObject) Handle(a ...interface{}) {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(*bool); !ok {
			goto check_1
		}
		p.Handle__SWIG_0(a[0].(*bool))
		return
	}
check_1:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint8); !ok {
			goto check_2
		}
		p.Handle__SWIG_1(a[0].(Uint8))
		return
	}
check_2:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint16); !ok {
			goto check_3
		}
		p.Handle__SWIG_2(a[0].(Uint16))
		return
	}
check_3:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint32); !ok {
			goto check_4
		}
		p.Handle__SWIG_3(a[0].(Uint32))
		return
	}
check_4:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint64); !ok {
			goto check_5
		}
		p.Handle__SWIG_4(a[0].(Uint64))
		return
	}
check_5:
	if argc == 1 {
		if _, ok := a[0].(*float32); !ok {
			goto check_6
		}
		p.Handle__SWIG_5(a[0].(*float32))
		return
	}
check_6:
	if argc == 1 {
		if _, ok := a[0].(*float64); !ok {
			goto check_7
		}
		p.Handle__SWIG_6(a[0].(*float64))
		return
	}
check_7:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrSint8); !ok {
			goto check_8
		}
		p.Handle__SWIG_7(a[0].(Sint8))
		return
	}
check_8:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrSint16); !ok {
			goto check_9
		}
		p.Handle__SWIG_8(a[0].(Sint16))
		return
	}
check_9:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrSint32); !ok {
			goto check_10
		}
		p.Handle__SWIG_9(a[0].(Sint32))
		return
	}
check_10:
	if argc == 1 {
		p.Handle__SWIG_10(a[0].(Sint64))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_0(arg2 Uint8) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_0_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_1(arg2 Uint16) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_1_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_2(arg2 Uint32) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_2_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_3(arg2 Uint64) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_3_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_4(arg2 Sint8) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_4_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_5(arg2 Sint16) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_5_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_6(arg2 Sint32) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_6_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_7(arg2 Sint64) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_7_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_8(arg2 *float32) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_8_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrStateWalkingObject) LeftShift__SWIG_9(arg2 *float64) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_StateWalkingObject_LeftShift__SWIG_9_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))))
	return swig_r
}

func (p SwigcptrStateWalkingObject) LeftShift(a ...interface{}) StateWalkingObject {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint8); !ok {
			goto check_1
		}
		return p.LeftShift__SWIG_0(a[0].(Uint8))
	}
check_1:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint16); !ok {
			goto check_2
		}
		return p.LeftShift__SWIG_1(a[0].(Uint16))
	}
check_2:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint32); !ok {
			goto check_3
		}
		return p.LeftShift__SWIG_2(a[0].(Uint32))
	}
check_3:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint64); !ok {
			goto check_4
		}
		return p.LeftShift__SWIG_3(a[0].(Uint64))
	}
check_4:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrSint8); !ok {
			goto check_5
		}
		return p.LeftShift__SWIG_4(a[0].(Sint8))
	}
check_5:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrSint16); !ok {
			goto check_6
		}
		return p.LeftShift__SWIG_5(a[0].(Sint16))
	}
check_6:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrSint32); !ok {
			goto check_7
		}
		return p.LeftShift__SWIG_6(a[0].(Sint32))
	}
check_7:
	if argc == 1 {
		if _, ok := a[0].(SwigcptrSint64); !ok {
			goto check_8
		}
		return p.LeftShift__SWIG_7(a[0].(Sint64))
	}
check_8:
	if argc == 1 {
		if _, ok := a[0].(*float32); !ok {
			goto check_9
		}
		return p.LeftShift__SWIG_8(a[0].(*float32))
	}
check_9:
	if argc == 1 {
		return p.LeftShift__SWIG_9(a[0].(*float64))
	}
	panic("No match for overloaded function call")
}

type StateWalkingObject interface {
	Swigcptr() uintptr
	SwigIsStateWalkingObject()
	Get_properties() (_swig_ret Uint32)
	Is_read_only() (_swig_ret bool)
	Is_write_only() (_swig_ret bool)
	Is_clumsy() (_swig_ret bool)
	Is_seeder() (_swig_ret bool)
	Handle(a ...interface{})
	LeftShift(a ...interface{}) StateWalkingObject
}

func Vrng_to_rng_seeder(arg1 PractRand_RNGs_vRNG) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_vrng_to_rng_seeder_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func Get_autoseeder(arg1 uintptr) (_swig_ret StateWalkingObject) {
	var swig_r StateWalkingObject
	_swig_i_0 := arg1
	swig_r = (StateWalkingObject)(SwigcptrStateWalkingObject(C._wrap_get_autoseeder_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))))
	return swig_r
}

type SwigcptrSalsa uintptr

func (p SwigcptrSalsa) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSalsa) SwigIsSalsa() {
}

func _swig_getsalsa_OUTPUT_TYPE_salsa() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_OUTPUT_TYPE_salsa_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var SalsaOUTPUT_TYPE int = _swig_getsalsa_OUTPUT_TYPE_salsa()
func _swig_getsalsa_OUTPUT_BITS_salsa() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_OUTPUT_BITS_salsa_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var SalsaOUTPUT_BITS int = _swig_getsalsa_OUTPUT_BITS_salsa()
func _swig_getsalsa_FLAGS_salsa() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_FLAGS_salsa_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var SalsaFLAGS int = _swig_getsalsa_FLAGS_salsa()
func NewSalsa() (_swig_ret Salsa) {
	var swig_r Salsa
	swig_r = (Salsa)(SwigcptrSalsa(C._wrap_new_salsa_rng_bfbcb6a9797b2cf1()))
	return swig_r
}

func DeleteSalsa(arg1 Salsa) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_salsa_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrSalsa) Raw32() (_swig_ret Uint32) {
	var swig_r Uint32
	_swig_i_0 := arg1
	swig_r = (Uint32)(SwigcptrUint32(C._wrap_salsa_raw32_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrSalsa) Seed__SWIG_0(arg2 Uint64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_salsa_seed__SWIG_0_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrSalsa) Seed__SWIG_1(arg2 Uint32, arg3 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	C._wrap_salsa_seed__SWIG_1_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C._Bool(_swig_i_2))
}

func (arg1 SwigcptrSalsa) Seed__SWIG_2(arg2 Uint32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_salsa_seed__SWIG_2_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (p SwigcptrSalsa) Seed(a ...interface{}) {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(SwigcptrUint64); !ok {
			goto check_1
		}
		p.Seed__SWIG_0(a[0].(Uint64))
		return
	}
check_1:
	if argc == 1 {
		p.Seed__SWIG_2(a[0].(Uint32))
		return
	}
	if argc == 2 {
		p.Seed__SWIG_1(a[0].(Uint32), a[1].(bool))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrSalsa) Seed_short__SWIG_0(arg2 Uint32, arg3 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	C._wrap_salsa_seed_short__SWIG_0_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C._Bool(_swig_i_2))
}

func (arg1 SwigcptrSalsa) Seed_short__SWIG_1(arg2 Uint32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_salsa_seed_short__SWIG_1_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (p SwigcptrSalsa) Seed_short(a ...interface{}) {
	argc := len(a)
	if argc == 1 {
		p.Seed_short__SWIG_1(a[0].(Uint32))
		return
	}
	if argc == 2 {
		p.Seed_short__SWIG_0(a[0].(Uint32), a[1].(bool))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrSalsa) Walk_state(arg2 StateWalkingObject) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_salsa_walk_state_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrSalsa) Seek_forward(arg2 Uint64, arg3 Uint64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_salsa_seek_forward_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func (arg1 SwigcptrSalsa) Seek_backward(arg2 Uint64, arg3 Uint64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_salsa_seek_backward_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func (arg1 SwigcptrSalsa) Set_rounds(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_salsa_set_rounds_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrSalsa) Get_rounds() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_salsa_get_rounds_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func SalsaSelf_test() {
	C._wrap_salsa_self_test_rng_bfbcb6a9797b2cf1()
}

type Salsa interface {
	Swigcptr() uintptr
	SwigIsSalsa()
	Raw32() (_swig_ret Uint32)
	Seed(a ...interface{})
	Seed_short(a ...interface{})
	Walk_state(arg2 StateWalkingObject)
	Seek_forward(arg2 Uint64, arg3 Uint64)
	Seek_backward(arg2 Uint64, arg3 Uint64)
	Set_rounds(arg2 int)
	Get_rounds() (_swig_ret int)
}

type SwigcptrSfc16 uintptr

func (p SwigcptrSfc16) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSfc16) SwigIsSfc16() {
}

func _swig_getsfc16_OUTPUT_TYPE_sfc16() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_OUTPUT_TYPE_sfc16_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var Sfc16OUTPUT_TYPE int = _swig_getsfc16_OUTPUT_TYPE_sfc16()
func _swig_getsfc16_OUTPUT_BITS_sfc16() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_OUTPUT_BITS_sfc16_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var Sfc16OUTPUT_BITS int = _swig_getsfc16_OUTPUT_BITS_sfc16()
func _swig_getsfc16_FLAGS_sfc16() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_FLAGS_sfc16_rng_bfbcb6a9797b2cf1())
	return swig_r
}

var Sfc16FLAGS int = _swig_getsfc16_FLAGS_sfc16()
func (arg1 SwigcptrSfc16) Raw16() (_swig_ret Uint16) {
	var swig_r Uint16
	_swig_i_0 := arg1
	swig_r = (Uint16)(SwigcptrUint16(C._wrap_sfc16_raw16_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrSfc16) Seed__SWIG_0(arg2 Uint64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_sfc16_seed__SWIG_0_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrSfc16) Seed_fast(arg2 Uint64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_sfc16_seed_fast_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrSfc16) Seed__SWIG_1(arg2 Uint16, arg3 Uint16, arg4 Uint16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3.Swigcptr()
	_swig_i_3 := arg4.Swigcptr()
	C._wrap_sfc16_seed__SWIG_1_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))
}

func (p SwigcptrSfc16) Seed(a ...interface{}) {
	argc := len(a)
	if argc == 1 {
		p.Seed__SWIG_0(a[0].(Uint64))
		return
	}
	if argc == 3 {
		p.Seed__SWIG_1(a[0].(Uint16), a[1].(Uint16), a[2].(Uint16))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrSfc16) Walk_state(arg2 StateWalkingObject) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_sfc16_walk_state_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func NewSfc16() (_swig_ret Sfc16) {
	var swig_r Sfc16
	swig_r = (Sfc16)(SwigcptrSfc16(C._wrap_new_sfc16_rng_bfbcb6a9797b2cf1()))
	return swig_r
}

func DeleteSfc16(arg1 Sfc16) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_sfc16_rng_bfbcb6a9797b2cf1(C.uintptr_t(_swig_i_0))
}

type Sfc16 interface {
	Swigcptr() uintptr
	SwigIsSfc16()
	Raw16() (_swig_ret Uint16)
	Seed_fast(arg2 Uint64)
	Seed(a ...interface{})
	Walk_state(arg2 StateWalkingObject)
}


type SwigcptrSint16 uintptr
type Sint16 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrSint16) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrUint32 uintptr
type Uint32 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrUint32) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrUint64 uintptr
type Uint64 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrUint64) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrSint8 uintptr
type Sint8 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrSint8) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrUint8 uintptr
type Uint8 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrUint8) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrUint16 uintptr
type Uint16 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrUint16) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrSint32 uintptr
type Sint32 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrSint32) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrSint64 uintptr
type Sint64 interface {
	Swigcptr() uintptr;
}
func (p SwigcptrSint64) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrPractRand_RNGs_vRNG uintptr
type PractRand_RNGs_vRNG interface {
	Swigcptr() uintptr;
}
func (p SwigcptrPractRand_RNGs_vRNG) Swigcptr() uintptr {
	return uintptr(p)
}


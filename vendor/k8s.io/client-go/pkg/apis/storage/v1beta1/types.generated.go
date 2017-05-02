/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package v1beta1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg2_v1 "k8s.io/client-go/pkg/api/v1"
	pkg1_v1 "k8s.io/client-go/pkg/apis/meta/v1"
	pkg3_types "k8s.io/client-go/pkg/types"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg2_v1.ObjectMeta
		var v1 pkg1_v1.TypeMeta
		var v2 pkg3_types.UID
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *StorageClass) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [5]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[4] = len(x.Parameters) != 0
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(5)
			} else {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy11 := &x.ObjectMeta
					yy11.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Provisioner))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("provisioner"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Provisioner))
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					if x.Parameters == nil {
						r.EncodeNil()
					} else {
						yym16 := z.EncBinary()
						_ = yym16
						if false {
						} else {
							z.F.EncMapStringStringV(x.Parameters, false, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("parameters"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Parameters == nil {
						r.EncodeNil()
					} else {
						yym17 := z.EncBinary()
						_ = yym17
						if false {
						} else {
							z.F.EncMapStringStringV(x.Parameters, false, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *StorageClass) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym18 := z.DecBinary()
	_ = yym18
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct19 := r.ContainerType()
		if yyct19 == codecSelferValueTypeMap1234 {
			yyl19 := r.ReadMapStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl19, d)
			}
		} else if yyct19 == codecSelferValueTypeArray1234 {
			yyl19 := r.ReadArrayStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl19, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *StorageClass) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys20Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys20Slc
	var yyhl20 bool = l >= 0
	for yyj20 := 0; ; yyj20++ {
		if yyhl20 {
			if yyj20 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys20Slc = r.DecodeBytes(yys20Slc, true, true)
		yys20 := string(yys20Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys20 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_v1.ObjectMeta{}
			} else {
				yyv23 := &x.ObjectMeta
				yyv23.CodecDecodeSelf(d)
			}
		case "provisioner":
			if r.TryDecodeAsNil() {
				x.Provisioner = ""
			} else {
				x.Provisioner = string(r.DecodeString())
			}
		case "parameters":
			if r.TryDecodeAsNil() {
				x.Parameters = nil
			} else {
				yyv25 := &x.Parameters
				yym26 := z.DecBinary()
				_ = yym26
				if false {
				} else {
					z.F.DecMapStringStringX(yyv25, false, d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys20)
		} // end switch yys20
	} // end for yyj20
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *StorageClass) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj27 int
	var yyb27 bool
	var yyhl27 bool = l >= 0
	yyj27++
	if yyhl27 {
		yyb27 = yyj27 > l
	} else {
		yyb27 = r.CheckBreak()
	}
	if yyb27 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj27++
	if yyhl27 {
		yyb27 = yyj27 > l
	} else {
		yyb27 = r.CheckBreak()
	}
	if yyb27 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj27++
	if yyhl27 {
		yyb27 = yyj27 > l
	} else {
		yyb27 = r.CheckBreak()
	}
	if yyb27 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_v1.ObjectMeta{}
	} else {
		yyv30 := &x.ObjectMeta
		yyv30.CodecDecodeSelf(d)
	}
	yyj27++
	if yyhl27 {
		yyb27 = yyj27 > l
	} else {
		yyb27 = r.CheckBreak()
	}
	if yyb27 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Provisioner = ""
	} else {
		x.Provisioner = string(r.DecodeString())
	}
	yyj27++
	if yyhl27 {
		yyb27 = yyj27 > l
	} else {
		yyb27 = r.CheckBreak()
	}
	if yyb27 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Parameters = nil
	} else {
		yyv32 := &x.Parameters
		yym33 := z.DecBinary()
		_ = yym33
		if false {
		} else {
			z.F.DecMapStringStringX(yyv32, false, d)
		}
	}
	for {
		yyj27++
		if yyhl27 {
			yyb27 = yyj27 > l
		} else {
			yyb27 = r.CheckBreak()
		}
		if yyb27 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj27-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *StorageClassList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym34 := z.EncBinary()
		_ = yym34
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep35 := !z.EncBinary()
			yy2arr35 := z.EncBasicHandle().StructToArray
			var yyq35 [4]bool
			_, _, _ = yysep35, yyq35, yy2arr35
			const yyr35 bool = false
			yyq35[0] = x.Kind != ""
			yyq35[1] = x.APIVersion != ""
			yyq35[2] = true
			var yynn35 int
			if yyr35 || yy2arr35 {
				r.EncodeArrayStart(4)
			} else {
				yynn35 = 1
				for _, b := range yyq35 {
					if b {
						yynn35++
					}
				}
				r.EncodeMapStart(yynn35)
				yynn35 = 0
			}
			if yyr35 || yy2arr35 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq35[0] {
					yym37 := z.EncBinary()
					_ = yym37
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq35[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym38 := z.EncBinary()
					_ = yym38
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr35 || yy2arr35 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq35[1] {
					yym40 := z.EncBinary()
					_ = yym40
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq35[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym41 := z.EncBinary()
					_ = yym41
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr35 || yy2arr35 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq35[2] {
					yy43 := &x.ListMeta
					yym44 := z.EncBinary()
					_ = yym44
					if false {
					} else if z.HasExtensions() && z.EncExt(yy43) {
					} else {
						z.EncFallback(yy43)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq35[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy45 := &x.ListMeta
					yym46 := z.EncBinary()
					_ = yym46
					if false {
					} else if z.HasExtensions() && z.EncExt(yy45) {
					} else {
						z.EncFallback(yy45)
					}
				}
			}
			if yyr35 || yy2arr35 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym48 := z.EncBinary()
					_ = yym48
					if false {
					} else {
						h.encSliceStorageClass(([]StorageClass)(x.Items), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("items"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym49 := z.EncBinary()
					_ = yym49
					if false {
					} else {
						h.encSliceStorageClass(([]StorageClass)(x.Items), e)
					}
				}
			}
			if yyr35 || yy2arr35 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *StorageClassList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym50 := z.DecBinary()
	_ = yym50
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct51 := r.ContainerType()
		if yyct51 == codecSelferValueTypeMap1234 {
			yyl51 := r.ReadMapStart()
			if yyl51 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl51, d)
			}
		} else if yyct51 == codecSelferValueTypeArray1234 {
			yyl51 := r.ReadArrayStart()
			if yyl51 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl51, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *StorageClassList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys52Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys52Slc
	var yyhl52 bool = l >= 0
	for yyj52 := 0; ; yyj52++ {
		if yyhl52 {
			if yyj52 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys52Slc = r.DecodeBytes(yys52Slc, true, true)
		yys52 := string(yys52Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys52 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ListMeta = pkg1_v1.ListMeta{}
			} else {
				yyv55 := &x.ListMeta
				yym56 := z.DecBinary()
				_ = yym56
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv55) {
				} else {
					z.DecFallback(yyv55, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv57 := &x.Items
				yym58 := z.DecBinary()
				_ = yym58
				if false {
				} else {
					h.decSliceStorageClass((*[]StorageClass)(yyv57), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys52)
		} // end switch yys52
	} // end for yyj52
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *StorageClassList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj59 int
	var yyb59 bool
	var yyhl59 bool = l >= 0
	yyj59++
	if yyhl59 {
		yyb59 = yyj59 > l
	} else {
		yyb59 = r.CheckBreak()
	}
	if yyb59 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj59++
	if yyhl59 {
		yyb59 = yyj59 > l
	} else {
		yyb59 = r.CheckBreak()
	}
	if yyb59 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj59++
	if yyhl59 {
		yyb59 = yyj59 > l
	} else {
		yyb59 = r.CheckBreak()
	}
	if yyb59 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg1_v1.ListMeta{}
	} else {
		yyv62 := &x.ListMeta
		yym63 := z.DecBinary()
		_ = yym63
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv62) {
		} else {
			z.DecFallback(yyv62, false)
		}
	}
	yyj59++
	if yyhl59 {
		yyb59 = yyj59 > l
	} else {
		yyb59 = r.CheckBreak()
	}
	if yyb59 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv64 := &x.Items
		yym65 := z.DecBinary()
		_ = yym65
		if false {
		} else {
			h.decSliceStorageClass((*[]StorageClass)(yyv64), d)
		}
	}
	for {
		yyj59++
		if yyhl59 {
			yyb59 = yyj59 > l
		} else {
			yyb59 = r.CheckBreak()
		}
		if yyb59 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj59-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSliceStorageClass(v []StorageClass, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv66 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy67 := &yyv66
		yy67.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSliceStorageClass(v *[]StorageClass, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv68 := *v
	yyh68, yyl68 := z.DecSliceHelperStart()
	var yyc68 bool
	if yyl68 == 0 {
		if yyv68 == nil {
			yyv68 = []StorageClass{}
			yyc68 = true
		} else if len(yyv68) != 0 {
			yyv68 = yyv68[:0]
			yyc68 = true
		}
	} else if yyl68 > 0 {
		var yyrr68, yyrl68 int
		var yyrt68 bool
		if yyl68 > cap(yyv68) {

			yyrg68 := len(yyv68) > 0
			yyv268 := yyv68
			yyrl68, yyrt68 = z.DecInferLen(yyl68, z.DecBasicHandle().MaxInitLen, 280)
			if yyrt68 {
				if yyrl68 <= cap(yyv68) {
					yyv68 = yyv68[:yyrl68]
				} else {
					yyv68 = make([]StorageClass, yyrl68)
				}
			} else {
				yyv68 = make([]StorageClass, yyrl68)
			}
			yyc68 = true
			yyrr68 = len(yyv68)
			if yyrg68 {
				copy(yyv68, yyv268)
			}
		} else if yyl68 != len(yyv68) {
			yyv68 = yyv68[:yyl68]
			yyc68 = true
		}
		yyj68 := 0
		for ; yyj68 < yyrr68; yyj68++ {
			yyh68.ElemContainerState(yyj68)
			if r.TryDecodeAsNil() {
				yyv68[yyj68] = StorageClass{}
			} else {
				yyv69 := &yyv68[yyj68]
				yyv69.CodecDecodeSelf(d)
			}

		}
		if yyrt68 {
			for ; yyj68 < yyl68; yyj68++ {
				yyv68 = append(yyv68, StorageClass{})
				yyh68.ElemContainerState(yyj68)
				if r.TryDecodeAsNil() {
					yyv68[yyj68] = StorageClass{}
				} else {
					yyv70 := &yyv68[yyj68]
					yyv70.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj68 := 0
		for ; !r.CheckBreak(); yyj68++ {

			if yyj68 >= len(yyv68) {
				yyv68 = append(yyv68, StorageClass{}) // var yyz68 StorageClass
				yyc68 = true
			}
			yyh68.ElemContainerState(yyj68)
			if yyj68 < len(yyv68) {
				if r.TryDecodeAsNil() {
					yyv68[yyj68] = StorageClass{}
				} else {
					yyv71 := &yyv68[yyj68]
					yyv71.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj68 < len(yyv68) {
			yyv68 = yyv68[:yyj68]
			yyc68 = true
		} else if yyj68 == 0 && yyv68 == nil {
			yyv68 = []StorageClass{}
			yyc68 = true
		}
	}
	yyh68.End()
	if yyc68 {
		*v = yyv68
	}
}

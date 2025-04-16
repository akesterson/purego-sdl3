package sdl

import "unsafe"

const (
	AlphaOpaque           = 255
	AlphaOpaqueFloat      = 1.0
	AlphaTransparent      = 0
	AlphaTransparentFloat = 0.0
)

type Colorspace uint32

const (
	ColorspaceUnknown       Colorspace = 0
	ColorspaceSRGB          Colorspace = 0x120005a0
	ColorspaceSRGBLinear    Colorspace = 0x12000500
	ColorspaceHDR10         Colorspace = 0x12002600
	ColorspaceJPEG          Colorspace = 0x220004c6
	ColorspaceBT601Limited  Colorspace = 0x211018c6
	ColorspaceBT601Full     Colorspace = 0x221018c6
	ColorspaceBT709Limited  Colorspace = 0x21100421
	ColorspaceBT709Full     Colorspace = 0x22100421
	ColorspaceBT2020Limited Colorspace = 0x21102609
	ColorspaceBT2020Full    Colorspace = 0x22102609
	ColorspaceRGBDefault    Colorspace = ColorspaceSRGB
	ColorspaceYUVDefault    Colorspace = ColorspaceJPEG
)

type ChromaLocation uint32

const (
	ChromaLocationNone ChromaLocation = iota
	ChromaLocationLeft
	ChromaLocationCenter
	ChromaLocationTopLeft
)

type MatrixCoefficients uint32

const (
	MatrixCoefficientsIdentity         MatrixCoefficients = 0
	MatrixCoefficientsBT709            MatrixCoefficients = 1
	MatrixCoefficientsUnspecified      MatrixCoefficients = 2
	MatrixCoefficientsFCC              MatrixCoefficients = 4
	MatrixCoefficientsBT470BG          MatrixCoefficients = 5
	MatrixCoefficientsBT601            MatrixCoefficients = 6
	MatrixCoefficientsSMPTE240         MatrixCoefficients = 7
	MatrixCoefficientsYCGCO            MatrixCoefficients = 8
	MatrixCoefficientsBT2020NCL        MatrixCoefficients = 9
	MatrixCoefficientsBT2020CL         MatrixCoefficients = 10
	MatrixCoefficientsSMPTE2085        MatrixCoefficients = 11
	MatrixCoefficientsChromaDerivedNCL MatrixCoefficients = 12
	MatrixCoefficientsChromaDerivedCL  MatrixCoefficients = 13
	MatrixCoefficientsICTCP            MatrixCoefficients = 14
	MatrixCoefficientsCustom           MatrixCoefficients = 31
)

type TransferCharacteristics uint32

const (
	TransferCharacteristicsUnknown      TransferCharacteristics = 0
	TransferCharacteristicsBT709        TransferCharacteristics = 1
	TransferCharacteristicsUnspecified  TransferCharacteristics = 2
	TransferCharacteristicsGamma22      TransferCharacteristics = 4
	TransferCharacteristicsGamma28      TransferCharacteristics = 5
	TransferCharacteristicsBT601        TransferCharacteristics = 6
	TransferCharacteristicsSMPTE240     TransferCharacteristics = 7
	TransferCharacteristicsLinear       TransferCharacteristics = 8
	TransferCharacteristicsLog100       TransferCharacteristics = 9
	TransferCharacteristicsLog100Sqrt10 TransferCharacteristics = 10
	TransferCharacteristicsIEC61966     TransferCharacteristics = 11
	TransferCharacteristicsBT1361       TransferCharacteristics = 12
	TransferCharacteristicsSRGB         TransferCharacteristics = 13
	TransferCharacteristicsBT202010Bit  TransferCharacteristics = 14
	TransferCharacteristicsBT202012Bit  TransferCharacteristics = 15
	TransferCharacteristicsPQ           TransferCharacteristics = 16
	TransferCharacteristicsSMPTE428     TransferCharacteristics = 17
	TransferCharacteristicsHLG          TransferCharacteristics = 18
	TransferCharacteristicsCustom       TransferCharacteristics = 31
)

type ColorPrimaries uint32

const (
	ColorPrimariesUnknown     ColorPrimaries = 0
	ColorPrimariesBT709       ColorPrimaries = 1
	ColorPrimariesUnspecified ColorPrimaries = 2
	ColorPrimariesBT470M      ColorPrimaries = 4
	ColorPrimariesBT470BG     ColorPrimaries = 5
	ColorPrimariesBT601       ColorPrimaries = 6
	ColorPrimariesSMPTE240    ColorPrimaries = 7
	ColorPrimariesGenericFilm ColorPrimaries = 8
	ColorPrimariesBT2020      ColorPrimaries = 9
	ColorPrimariesXYZ         ColorPrimaries = 10
	ColorPrimariesSMPTE431    ColorPrimaries = 11
	ColorPrimariesSMPTE432    ColorPrimaries = 12
	ColorPrimariesEBU3213     ColorPrimaries = 22
	ColorPrimariesCustom      ColorPrimaries = 31
)

type ColorRange uint32

const (
	ColorRangeUnknown ColorRange = iota
	ColorRangeLimited
	ColorRangeFull
)

type ColorType uint32

const (
	ColorTypeUnknown ColorType = iota
	ColorTypeRGB
	ColorTypeYCBCR
)

type PixelType uint32

const (
	PixelTypeUnknown PixelType = iota
	PixelTypeIndex1
	PixelTypeIndex4
	PixelTypeIndex8
	PixelTypePacked8
	PixelTypePacked16
	PixelTypePacked32
	PixelTypeArrayU8
	PixelTypeArrayU16
	PixelTypeArrayU32
	PixelTypeArrayF16
	PixelTypeArrayF32
	PixelTypeIndex2
)

type BitmapOrder uint32

const (
	BitmapOrderNone BitmapOrder = iota
	BitmapOrder4321
	BitmapOrder1234
)

type PackedOrder uint32

const (
	PackedOrderNone PackedOrder = iota
	PackedOrderXRGB
	PackedOrderRGBX
	PackedOrderARGB
	PackedOrderRGBA
	PackedOrderXBGR
	PackedOrderBGRX
	PackedOrderABGR
	PackedOrderBGRA
)

type ArrayOrder uint32

const (
	ArrayOrderNone ArrayOrder = iota
	ArrayOrderRGB
	ArrayOrderRGBA
	ArrayOrderARGB
	ArrayOrderBGR
	ArrayOrderBGRA
	ArrayOrderABGR
)

type PackedLayout uint32

const (
	PackedLayoutNone PackedLayout = iota
	PackedLayout332
	PackedLayout4444
	PackedLayout1555
	PackedLayout5551
	PackedLayout565
	PackedLayout8888
	PackedLayout2101010
	PackedLayout1010102
)

type PixelFormat uint32

const (
	PixelFormatUnknown      PixelFormat = 0
	PixelFormatIndex1LSB    PixelFormat = 0x11100100
	PixelFormatIndex1MSB    PixelFormat = 0x11200100
	PixelFormatIndex2LSB    PixelFormat = 0x1C100200
	PixelFormatIndex2MSB    PixelFormat = 0x1C200200
	PixelFormatIndex4LSB    PixelFormat = 0x12100400
	PixelFormatIndex4MSB    PixelFormat = 0x12200400
	PixelFormatIndex8       PixelFormat = 0x13000801
	PixelFormatRGB332       PixelFormat = 0x14110801
	PixelFormatXRGB4444     PixelFormat = 0x15120C02
	PixelFormatXBGR4444     PixelFormat = 0x15520C02
	PixelFormatXRGB1555     PixelFormat = 0x15130F02
	PixelFormatXBGR1555     PixelFormat = 0x15530F02
	PixelFormatARGB4444     PixelFormat = 0x15321002
	PixelFormatRGBA4444     PixelFormat = 0x15421002
	PixelFormatABGR4444     PixelFormat = 0x15721002
	PixelFormatBGRA4444     PixelFormat = 0x15821002
	PixelFormatARGB1555     PixelFormat = 0x15331002
	PixelFormatRGBA5551     PixelFormat = 0x15441002
	PixelFormatABGR1555     PixelFormat = 0x15731002
	PixelFormatBGRA5551     PixelFormat = 0x15841002
	PixelFormatRGB565       PixelFormat = 0x15151002
	PixelFormatBGR565       PixelFormat = 0x15551002
	PixelFormatRGB24        PixelFormat = 0x17101803
	PixelFormatBGR24        PixelFormat = 0x17401803
	PixelFormatXRGB8888     PixelFormat = 0x16161804
	PixelFormatRGBX8888     PixelFormat = 0x16261804
	PixelFormatXBGR8888     PixelFormat = 0x16561804
	PixelFormatBGRX8888     PixelFormat = 0x16661804
	PixelFormatARGB8888     PixelFormat = 0x16362004
	PixelFormatRGBA8888     PixelFormat = 0x16462004
	PixelFormatABGR8888     PixelFormat = 0x16762004
	PixelFormatBGRA8888     PixelFormat = 0x16862004
	PixelFormatXRGB2101010  PixelFormat = 0x16172004
	PixelFormatXBGR2101010  PixelFormat = 0x16572004
	PixelFormatARGB2101010  PixelFormat = 0x16372004
	PixelFormatABGR2101010  PixelFormat = 0x16772004
	PixelFormatRGB48        PixelFormat = 0x18103006
	PixelFormatBGR48        PixelFormat = 0x18403006
	PixelFormatRGBA64       PixelFormat = 0x18204008
	PixelFormatARGB64       PixelFormat = 0x18304008
	PixelFormatBGRA64       PixelFormat = 0x18504008
	PixelFormatABGR64       PixelFormat = 0x18604008
	PixelFormatRGB48Float   PixelFormat = 0x1A103006
	PixelFormatBGR48Float   PixelFormat = 0x1A403006
	PixelFormatRGBA64Float  PixelFormat = 0x1A204008
	PixelFormatARGB64Float  PixelFormat = 0x1A304008
	PixelFormatBGRA64Float  PixelFormat = 0x1A504008
	PixelFormatABGR64Float  PixelFormat = 0x1A604008
	PixelFormatRGB96Float   PixelFormat = 0x1B10600C
	PixelFormatBGR96Float   PixelFormat = 0x1B40600C
	PixelFormatRGBA128Float PixelFormat = 0x1B208010
	PixelFormatARGB128Float PixelFormat = 0x1B308010
	PixelFormatBGRA128Float PixelFormat = 0x1B508010
	PixelFormatABGR128Float PixelFormat = 0x1B608010
	PixelFormatYV12         PixelFormat = 0x32315659
	PixelFormatIYUV         PixelFormat = 0x56555949
	PixelFormatYUY2         PixelFormat = 0x32595559
	PixelFormatUYVY         PixelFormat = 0x59565955
	PixelFormatYVYU         PixelFormat = 0x55595659
	PixelFormatNV12         PixelFormat = 0x3231564E
	PixelFormatNV21         PixelFormat = 0x3132564E
	PixelFormatP010         PixelFormat = 0x30313050
	PixelFormatExternalOES  PixelFormat = 0x2053454F
	PixelFormatRGBA32       PixelFormat = PixelFormatABGR8888
	PixelFormatARGB32       PixelFormat = PixelFormatBGRA8888
	PixelFormatBGRA32       PixelFormat = PixelFormatARGB8888
	PixelFormatABGR32       PixelFormat = PixelFormatRGBA8888
	PixelFormatRGBX32       PixelFormat = PixelFormatXBGR8888
	PixelFormatXRGB32       PixelFormat = PixelFormatBGRX8888
	PixelFormatBGRX32       PixelFormat = PixelFormatXRGB8888
	PixelFormatXBGR32       PixelFormat = PixelFormatRGBX8888
)

type Color struct {
	R, G, B, A uint8
}

type FColor struct {
	R, G, B, A float32
}

type Palette struct {
	ncolors int32
	colors  *Color
	// Version is for internal use only, do not touch.
	Version uint32
	// Refcount is for internal use only, do not touch.
	Refcount int32
}

type PixelFormatDetails struct {
	Format                         PixelFormat
	BitsPerPixel                   uint8
	BytesPerPixel                  uint8
	Padding                        [2]uint8
	Rmask, Gmask, Bmask, Amask     uint32
	Rbits, Gbits, Bbits, Abits     uint8
	Rshift, Gshift, Bshift, Ashift uint8
}

func (p *Palette) Colors() []Color {
	if p == nil || p.colors == nil {
		return nil
	}
	return unsafe.Slice(p.colors, p.ncolors)
}

func DefinePixelFourCC(a, b, c, d byte) PixelFormat {
	return PixelFormat(FourCC(a, b, c, d))
}

func DefinePixelFormat[T BitmapOrder | PackedOrder | ArrayOrder](pixelType PixelType, order T, layout PackedLayout, bitsPerPixel, bytesPerPixel byte) PixelFormat {
	return 1<<28 | PixelFormat(pixelType)<<24 | PixelFormat(order)<<20 | PixelFormat(layout)<<16 | PixelFormat(bitsPerPixel)<<8 | PixelFormat(bytesPerPixel)
}

func CreatePalette(ncolors int32) *Palette {
	return sdlCreatePalette(ncolors)
}

func DestroyPalette(palette *Palette) {
	sdlDestroyPalette(palette)
}

// func GetMasksForPixelFormat(format PixelFormat, bpp *int32, Rmask *uint32, Gmask *uint32, Bmask *uint32, Amask *uint32) bool {
//	return sdlGetMasksForPixelFormat(format, bpp, Rmask, Gmask, Bmask, Amask)
// }

func GetPixelFormatDetails(format PixelFormat) *PixelFormatDetails {
	return sdlGetPixelFormatDetails(format)
}

// func GetPixelFormatForMasks(bpp int32, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) PixelFormat {
//	return sdlGetPixelFormatForMasks(bpp, Rmask, Gmask, Bmask, Amask)
// }

// func GetPixelFormatName(format PixelFormat) string {
//	return sdlGetPixelFormatName(format)
// }

// func GetRGB(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8)  {
//	sdlGetRGB(pixel, format, palette, r, g, b)
// }

// func GetRGBA(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8, a *uint8)  {
//	sdlGetRGBA(pixel, format, palette, r, g, b, a)
// }

func MapRGB(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8) uint32 {
	return sdlMapRGB(format, palette, r, g, b)
}

// func MapRGBA(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8, a uint8) uint32 {
//	return sdlMapRGBA(format, palette, r, g, b, a)
// }

func SetPaletteColors(palette *Palette, firstcolor int32, Colors ...Color) bool {
	ncolors := len(Colors)
	var ptrColors *Color
	if ncolors > 0 {
		ptrColors = &Colors[0]
	}
	return sdlSetPaletteColors(palette, ptrColors, firstcolor, int32(ncolors))
}

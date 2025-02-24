//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package drawing ;import (_b "github.com/unidoc/unioffice";_gd "github.com/unidoc/unioffice/color";_d "github.com/unidoc/unioffice/measurement";_gdb "github.com/unidoc/unioffice/schema/soo/dml";);func (_dde ShapeProperties )LineProperties ()LineProperties {if _dde ._fbf .Ln ==nil {_dde ._fbf .Ln =_gdb .NewCT_LineProperties ();};return LineProperties {_dde ._fbf .Ln };};

// Run is a run within a paragraph.
type Run struct{_ce *_gdb .EG_TextRun };

// SetBulletFont controls the font for the bullet character.
func (_cg ParagraphProperties )SetBulletFont (f string ){if f ==""{_cg ._ae .BuFont =nil ;}else {_cg ._ae .BuFont =_gdb .NewCT_TextFont ();_cg ._ae .BuFont .TypefaceAttr =f ;};};

// GetPosition gets the position of the shape in EMU.
func (_dba ShapeProperties )GetPosition ()(int64 ,int64 ){_dba .ensureXfrm ();if _dba ._fbf .Xfrm .Off ==nil {_dba ._fbf .Xfrm .Off =_gdb .NewCT_Point2D ();};return *_dba ._fbf .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_dba ._fbf .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;};

// AddBreak adds a new line break to a paragraph.
func (_f Paragraph )AddBreak (){_cc :=_gdb .NewEG_TextRun ();_cc .Br =_gdb .NewCT_TextLineBreak ();_f ._cf .EG_TextRun =append (_f ._cf .EG_TextRun ,_cc );};

// RunProperties controls the run properties.
type RunProperties struct{_cfd *_gdb .CT_TextCharacterProperties ;};

// SetAlign controls the paragraph alignment
func (_fb ParagraphProperties )SetAlign (a _gdb .ST_TextAlignType ){_fb ._ae .AlgnAttr =a };

// SetHeight sets the height of the shape.
func (_ag ShapeProperties )SetHeight (h _d .Distance ){_ag .ensureXfrm ();if _ag ._fbf .Xfrm .Ext ==nil {_ag ._fbf .Xfrm .Ext =_gdb .NewCT_PositiveSize2D ();};_ag ._fbf .Xfrm .Ext .CyAttr =int64 (h /_d .EMU );};

// SetNumbered controls if bullets are numbered or not.
func (_ff ParagraphProperties )SetNumbered (scheme _gdb .ST_TextAutonumberScheme ){if scheme ==_gdb .ST_TextAutonumberSchemeUnset {_ff ._ae .BuAutoNum =nil ;}else {_ff ._ae .BuAutoNum =_gdb .NewCT_TextAutonumberBullet ();_ff ._ae .BuAutoNum .TypeAttr =scheme ;};};

// SetPosition sets the position of the shape.
func (_cec ShapeProperties )SetPosition (x ,y _d .Distance ){_cec .ensureXfrm ();if _cec ._fbf .Xfrm .Off ==nil {_cec ._fbf .Xfrm .Off =_gdb .NewCT_Point2D ();};_cec ._fbf .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_b .Int64 (int64 (x /_d .EMU ));_cec ._fbf .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_b .Int64 (int64 (y /_d .EMU ));};

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_gdb .CT_TextParagraph )Paragraph {return Paragraph {x }};

// LineJoin is the type of line join
type LineJoin byte ;

// SetSize sets the font size of the run text
func (_bf RunProperties )SetSize (sz _d .Distance ){_bf ._cfd .SzAttr =_b .Int32 (int32 (sz /_d .HundredthPoint ));};func (_ad ShapeProperties )SetSolidFill (c _gd .Color ){_ad .clearFill ();_ad ._fbf .SolidFill =_gdb .NewCT_SolidColorFillProperties ();_ad ._fbf .SolidFill .SrgbClr =_gdb .NewCT_SRgbColor ();_ad ._fbf .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// X returns the inner wrapped XML type.
func (_ab Paragraph )X ()*_gdb .CT_TextParagraph {return _ab ._cf };func (_de ShapeProperties )SetNoFill (){_de .clearFill ();_de ._fbf .NoFill =_gdb .NewCT_NoFillProperties ();};

// X returns the inner wrapped XML type.
func (_fg ShapeProperties )X ()*_gdb .CT_ShapeProperties {return _fg ._fbf };

// Properties returns the paragraph properties.
func (_dg Paragraph )Properties ()ParagraphProperties {if _dg ._cf .PPr ==nil {_dg ._cf .PPr =_gdb .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_dg ._cf .PPr );};

// SetLevel sets the level of indentation of a paragraph.
func (_bef ParagraphProperties )SetLevel (idx int32 ){_bef ._ae .LvlAttr =_b .Int32 (idx )};

// Paragraph is a paragraph within a document.
type Paragraph struct{_cf *_gdb .CT_TextParagraph };

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_gdb .EG_TextRun )Run {return Run {x }};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_gfe LineProperties )SetWidth (w _d .Distance ){_gfe ._gf .WAttr =_b .Int32 (int32 (w /_d .EMU ))};

// X returns the inner wrapped XML type.
func (_ced Run )X ()*_gdb .EG_TextRun {return _ced ._ce };type ShapeProperties struct{_fbf *_gdb .CT_ShapeProperties };func (_a LineProperties )SetNoFill (){_a .clearFill ();_a ._gf .NoFill =_gdb .NewCT_NoFillProperties ()};

// X returns the inner wrapped XML type.
func (_dd ParagraphProperties )X ()*_gdb .CT_TextParagraphProperties {return _dd ._ae };type LineProperties struct{_gf *_gdb .CT_LineProperties };

// SetBold controls the bolding of a run.
func (_af RunProperties )SetBold (b bool ){_af ._cfd .BAttr =_b .Bool (b )};

// SetSize sets the width and height of the shape.
func (_aa ShapeProperties )SetSize (w ,h _d .Distance ){_aa .SetWidth (w );_aa .SetHeight (h )};

// AddRun adds a new run to a paragraph.
func (_ba Paragraph )AddRun ()Run {_e :=MakeRun (_gdb .NewEG_TextRun ());_ba ._cf .EG_TextRun =append (_ba ._cf .EG_TextRun ,_e .X ());return _e ;};func MakeShapeProperties (x *_gdb .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};

// SetSolidFill controls the text color of a run.
func (_fc RunProperties )SetSolidFill (c _gd .Color ){_fc ._cfd .NoFill =nil ;_fc ._cfd .BlipFill =nil ;_fc ._cfd .GradFill =nil ;_fc ._cfd .GrpFill =nil ;_fc ._cfd .PattFill =nil ;_fc ._cfd .SolidFill =_gdb .NewCT_SolidColorFillProperties ();_fc ._cfd .SolidFill .SrgbClr =_gdb .NewCT_SRgbColor ();_fc ._cfd .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};func (_db LineProperties )clearFill (){_db ._gf .NoFill =nil ;_db ._gf .GradFill =nil ;_db ._gf .SolidFill =nil ;_db ._gf .PattFill =nil ;};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_ae *_gdb .CT_TextParagraphProperties ;};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_fdd ShapeProperties )SetFlipHorizontal (b bool ){_fdd .ensureXfrm ();if !b {_fdd ._fbf .Xfrm .FlipHAttr =nil ;}else {_fdd ._fbf .Xfrm .FlipHAttr =_b .Bool (true );};};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_gdb .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// SetFlipVertical controls if the shape is flipped vertically.
func (_ec ShapeProperties )SetFlipVertical (b bool ){_ec .ensureXfrm ();if !b {_ec ._fbf .Xfrm .FlipVAttr =nil ;}else {_ec ._fbf .Xfrm .FlipVAttr =_b .Bool (true );};};

// SetBulletChar sets the bullet character for the paragraph.
func (_fd ParagraphProperties )SetBulletChar (c string ){if c ==""{_fd ._ae .BuChar =nil ;}else {_fd ._ae .BuChar =_gdb .NewCT_TextCharBullet ();_fd ._ae .BuChar .CharAttr =c ;};};

// X returns the inner wrapped XML type.
func (_c LineProperties )X ()*_gdb .CT_LineProperties {return _c ._gf };

// SetGeometry sets the shape type of the shape
func (_abb ShapeProperties )SetGeometry (g _gdb .ST_ShapeType ){if _abb ._fbf .PrstGeom ==nil {_abb ._fbf .PrstGeom =_gdb .NewCT_PresetGeometry2D ();};_abb ._fbf .PrstGeom .PrstAttr =g ;};func (_be LineProperties )SetSolidFill (c _gd .Color ){_be .clearFill ();_be ._gf .SolidFill =_gdb .NewCT_SolidColorFillProperties ();_be ._gf .SolidFill .SrgbClr =_gdb .NewCT_SRgbColor ();_be ._gf .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetWidth sets the width of the shape.
func (_fe ShapeProperties )SetWidth (w _d .Distance ){_fe .ensureXfrm ();if _fe ._fbf .Xfrm .Ext ==nil {_fe ._fbf .Xfrm .Ext =_gdb .NewCT_PositiveSize2D ();};_fe ._fbf .Xfrm .Ext .CxAttr =int64 (w /_d .EMU );};func (_cb ShapeProperties )clearFill (){_cb ._fbf .NoFill =nil ;_cb ._fbf .BlipFill =nil ;_cb ._fbf .GradFill =nil ;_cb ._fbf .GrpFill =nil ;_cb ._fbf .SolidFill =nil ;_cb ._fbf .PattFill =nil ;};

// SetText sets the run's text contents.
func (_ea Run )SetText (s string ){_ea ._ce .Br =nil ;_ea ._ce .Fld =nil ;if _ea ._ce .R ==nil {_ea ._ce .R =_gdb .NewCT_RegularTextRun ();};_ea ._ce .R .T =s ;};func (_fgg ShapeProperties )ensureXfrm (){if _fgg ._fbf .Xfrm ==nil {_fgg ._fbf .Xfrm =_gdb .NewCT_Transform2D ();};};

// Properties returns the run's properties.
func (_abc Run )Properties ()RunProperties {if _abc ._ce .R ==nil {_abc ._ce .R =_gdb .NewCT_RegularTextRun ();};if _abc ._ce .R .RPr ==nil {_abc ._ce .R .RPr =_gdb .NewCT_TextCharacterProperties ();};return RunProperties {_abc ._ce .R .RPr };};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_gdb .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};

// SetJoin sets the line join style.
func (_ac LineProperties )SetJoin (e LineJoin ){_ac ._gf .Round =nil ;_ac ._gf .Miter =nil ;_ac ._gf .Bevel =nil ;switch e {case LineJoinRound :_ac ._gf .Round =_gdb .NewCT_LineJoinRound ();case LineJoinBevel :_ac ._gf .Bevel =_gdb .NewCT_LineJoinBevel ();case LineJoinMiter :_ac ._gf .Miter =_gdb .NewCT_LineJoinMiterProperties ();};};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// SetFont controls the font of a run.
func (_gb RunProperties )SetFont (s string ){_gb ._cfd .Latin =_gdb .NewCT_TextFont ();_gb ._cfd .Latin .TypefaceAttr =s ;};
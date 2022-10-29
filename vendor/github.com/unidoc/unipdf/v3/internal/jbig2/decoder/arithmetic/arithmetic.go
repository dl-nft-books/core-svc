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

package arithmetic ;import (_ag "fmt";_c "github.com/unidoc/unipdf/v3/common";_aff "github.com/unidoc/unipdf/v3/internal/bitwise";_cf "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_af "io";_d "strings";);func (_bb *DecoderStats )getMps ()byte {return _bb ._ggd [_bb ._ea ]};
func (_adaa *Decoder )lpsExchange (_bgd *DecoderStats ,_be int32 ,_gf uint32 )int {_gbd :=_bgd .getMps ();if _adaa ._gb < _gf {_bgd .setEntry (int (_g [_be ][1]));_adaa ._gb =_gf ;return int (_gbd );};if _g [_be ][3]==1{_bgd .toggleMps ();};_bgd .setEntry (int (_g [_be ][2]));
_adaa ._gb =_gf ;return int (1-_gbd );};func (_dgg *DecoderStats )toggleMps (){_dgg ._ggd [_dgg ._ea ]^=1};func (_gba *DecoderStats )setEntry (_beb int ){_ge :=byte (_beb &0x7f);_gba ._agf [_gba ._ea ]=_ge };func NewStats (contextSize int32 ,index int32 )*DecoderStats {return &DecoderStats {_ea :index ,_fe :contextSize ,_agf :make ([]byte ,contextSize ),_ggd :make ([]byte ,contextSize )};
};func New (r _aff .StreamReader )(*Decoder ,error ){_cfc :=&Decoder {_cd :r ,ContextSize :[]uint32 {16,13,10,10},ReferedToContextSize :[]uint32 {13,10}};if _ca :=_cfc .init ();_ca !=nil {return nil ,_ca ;};return _cfc ,nil ;};func (_gg *Decoder )init ()error {_gg ._f =_gg ._cd .StreamPosition ();
_dgf ,_cff :=_gg ._cd .ReadByte ();if _cff !=nil {_c .Log .Debug ("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076",_cff );return _cff ;};_gg ._cfg =_dgf ;_gg ._da =uint64 (_dgf )<<16;
if _cff =_gg .readByte ();_cff !=nil {return _cff ;};_gg ._da <<=7;_gg ._dg -=7;_gg ._gb =0x8000;_gg ._ad ++;return nil ;};func (_fb *DecoderStats )Overwrite (dNew *DecoderStats ){for _afc :=0;_afc < len (_fb ._agf );_afc ++{_fb ._agf [_afc ]=dNew ._agf [_afc ];
_fb ._ggd [_afc ]=dNew ._ggd [_afc ];};};func (_agd *Decoder )DecodeInt (stats *DecoderStats )(int32 ,error ){var (_e ,_ada int32 ;_ba ,_cfcd ,_dd int ;_caf error ;);if stats ==nil {stats =NewStats (512,1);};_agd ._b =1;_cfcd ,_caf =_agd .decodeIntBit (stats );
if _caf !=nil {return 0,_caf ;};_ba ,_caf =_agd .decodeIntBit (stats );if _caf !=nil {return 0,_caf ;};if _ba ==1{_ba ,_caf =_agd .decodeIntBit (stats );if _caf !=nil {return 0,_caf ;};if _ba ==1{_ba ,_caf =_agd .decodeIntBit (stats );if _caf !=nil {return 0,_caf ;
};if _ba ==1{_ba ,_caf =_agd .decodeIntBit (stats );if _caf !=nil {return 0,_caf ;};if _ba ==1{_ba ,_caf =_agd .decodeIntBit (stats );if _caf !=nil {return 0,_caf ;};if _ba ==1{_dd =32;_ada =4436;}else {_dd =12;_ada =340;};}else {_dd =8;_ada =84;};}else {_dd =6;
_ada =20;};}else {_dd =4;_ada =4;};}else {_dd =2;_ada =0;};for _cfe :=0;_cfe < _dd ;_cfe ++{_ba ,_caf =_agd .decodeIntBit (stats );if _caf !=nil {return 0,_caf ;};_e =(_e <<1)|int32 (_ba );};_e +=_ada ;if _cfcd ==0{return _e ,nil ;}else if _cfcd ==1&&_e > 0{return -_e ,nil ;
};return 0,_cf .ErrOOB ;};func (_eef *Decoder )renormalize ()error {for {if _eef ._dg ==0{if _bg :=_eef .readByte ();_bg !=nil {return _bg ;};};_eef ._gb <<=1;_eef ._da <<=1;_eef ._dg --;if (_eef ._gb &0x8000)!=0{break ;};};_eef ._da &=0xffffffff;return nil ;
};func (_ggb *Decoder )decodeIntBit (_ec *DecoderStats )(int ,error ){_ec .SetIndex (int32 (_ggb ._b ));_dde ,_ff :=_ggb .DecodeBit (_ec );if _ff !=nil {_c .Log .Debug ("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076",_ff );
return _dde ,_ff ;};if _ggb ._b < 256{_ggb ._b =((_ggb ._b <<uint64 (1))|int64 (_dde ))&0x1ff;}else {_ggb ._b =(((_ggb ._b <<uint64 (1)|int64 (_dde ))&511)|256)&0x1ff;};return _dde ,nil ;};func (_df *DecoderStats )String ()string {_dfa :=&_d .Builder {};
_dfa .WriteString (_ag .Sprintf ("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a",len (_df ._agf )));for _bc ,_ggad :=range _df ._agf {if _ggad !=0{_dfa .WriteString (_ag .Sprintf ("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a",_bc ,_ggad ));
};};return _dfa .String ();};func (_bag *Decoder )readByte ()error {if _bag ._cd .StreamPosition ()> _bag ._f {if _ ,_cdd :=_bag ._cd .Seek (-1,_af .SeekCurrent );_cdd !=nil {return _cdd ;};};_cg ,_ee :=_bag ._cd .ReadByte ();if _ee !=nil {return _ee ;
};_bag ._cfg =_cg ;if _bag ._cfg ==0xFF{_ef ,_dae :=_bag ._cd .ReadByte ();if _dae !=nil {return _dae ;};if _ef > 0x8F{_bag ._da +=0xFF00;_bag ._dg =8;if _ ,_ga :=_bag ._cd .Seek (-2,_af .SeekCurrent );_ga !=nil {return _ga ;};}else {_bag ._da +=uint64 (_ef )<<9;
_bag ._dg =7;};}else {_cg ,_ee =_bag ._cd .ReadByte ();if _ee !=nil {return _ee ;};_bag ._cfg =_cg ;_bag ._da +=uint64 (_bag ._cfg )<<8;_bag ._dg =8;};_bag ._da &=0xFFFFFFFFFF;return nil ;};func (_eb *Decoder )DecodeIAID (codeLen uint64 ,stats *DecoderStats )(int64 ,error ){_eb ._b =1;
var _fac uint64 ;for _fac =0;_fac < codeLen ;_fac ++{stats .SetIndex (int32 (_eb ._b ));_baa ,_cfa :=_eb .DecodeBit (stats );if _cfa !=nil {return 0,_cfa ;};_eb ._b =(_eb ._b <<1)|int64 (_baa );};_gd :=_eb ._b -(1<<codeLen );return _gd ,nil ;};var (_g =[][4]uint32 {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
);func (_fcf *DecoderStats )Reset (){for _ce :=0;_ce < len (_fcf ._agf );_ce ++{_fcf ._agf [_ce ]=0;_fcf ._ggd [_ce ]=0;};};func (_fbb *DecoderStats )SetIndex (index int32 ){_fbb ._ea =index };type Decoder struct{ContextSize []uint32 ;ReferedToContextSize []uint32 ;
_cd _aff .StreamReader ;_cfg uint8 ;_da uint64 ;_gb uint32 ;_b int64 ;_dg int32 ;_ad int32 ;_f int64 ;};func (_baf *DecoderStats )Copy ()*DecoderStats {_ded :=&DecoderStats {_fe :_baf ._fe ,_agf :make ([]byte ,_baf ._fe )};for _bae :=0;_bae < len (_baf ._agf );
_bae ++{_ded ._agf [_bae ]=_baf ._agf [_bae ];};return _ded ;};type DecoderStats struct{_ea int32 ;_fe int32 ;_agf []byte ;_ggd []byte ;};func (_gbe *Decoder )mpsExchange (_de *DecoderStats ,_facc int32 )int {_cge :=_de ._ggd [_de ._ea ];if _gbe ._gb < _g [_facc ][0]{if _g [_facc ][3]==1{_de .toggleMps ();
};_de .setEntry (int (_g [_facc ][2]));return int (1-_cge );};_de .setEntry (int (_g [_facc ][1]));return int (_cge );};func (_fa *Decoder )DecodeBit (stats *DecoderStats )(int ,error ){var (_aa int ;_fc =_g [stats .cx ()][0];_cb =int32 (stats .cx ());
);defer func (){_fa ._ad ++}();_fa ._gb -=_fc ;if (_fa ._da >>16)< uint64 (_fc ){_aa =_fa .lpsExchange (stats ,_cb ,_fc );if _afe :=_fa .renormalize ();_afe !=nil {return 0,_afe ;};}else {_fa ._da -=uint64 (_fc )<<16;if (_fa ._gb &0x8000)==0{_aa =_fa .mpsExchange (stats ,_cb );
if _db :=_fa .renormalize ();_db !=nil {return 0,_db ;};}else {_aa =int (stats .getMps ());};};return _aa ,nil ;};func (_fd *DecoderStats )cx ()byte {return _fd ._agf [_fd ._ea ]};
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: notif.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bnotif.proto\"\x18\n\x07MeowReq\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\"\x17\n\x07MeowRes\x12\x0c\n\x04\x66\x61\x63t\x18\x01 \x01(\t\"9\n\x06\x43urReq\x12\x10\n\x08\x66rom_cur\x18\x01 \x01(\t\x12\x0e\n\x06to_cur\x18\x02 \x01(\t\x12\r\n\x05\x63ount\x18\x03 \x01(\x05\"u\n\x06\x43urRes\x12\x0c\n\x04\x62\x61se\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x02 \x01(\t\x12!\n\x05rates\x18\x03 \x03(\x0b\x32\x12.CurRes.RatesEntry\x1a,\n\nRatesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x02:\x02\x38\x01\x32Q\n\x05Notif\x12\"\n\nStreamMeow\x12\x08.MeowReq\x1a\x08.MeowRes0\x01\x12$\n\x0eStreamCurrency\x12\x07.CurReq\x1a\x07.CurRes0\x01\x42\x0eZ\x0cgo_server/pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'notif_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\014go_server/pb'
  _globals['_CURRES_RATESENTRY']._options = None
  _globals['_CURRES_RATESENTRY']._serialized_options = b'8\001'
  _globals['_MEOWREQ']._serialized_start=15
  _globals['_MEOWREQ']._serialized_end=39
  _globals['_MEOWRES']._serialized_start=41
  _globals['_MEOWRES']._serialized_end=64
  _globals['_CURREQ']._serialized_start=66
  _globals['_CURREQ']._serialized_end=123
  _globals['_CURRES']._serialized_start=125
  _globals['_CURRES']._serialized_end=242
  _globals['_CURRES_RATESENTRY']._serialized_start=198
  _globals['_CURRES_RATESENTRY']._serialized_end=242
  _globals['_NOTIF']._serialized_start=244
  _globals['_NOTIF']._serialized_end=325
# @@protoc_insertion_point(module_scope)

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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bnotif.proto\"\x18\n\x07MeowReq\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\"\x17\n\x07MeowRes\x12\x0c\n\x04\x66\x61\x63t\x18\x01 \x01(\t2+\n\x05Notif\x12\"\n\nStreamMeow\x12\x08.MeowReq\x1a\x08.MeowRes0\x01\x42\x0eZ\x0cgo_server/pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'notif_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\014go_server/pb'
  _globals['_MEOWREQ']._serialized_start=15
  _globals['_MEOWREQ']._serialized_end=39
  _globals['_MEOWRES']._serialized_start=41
  _globals['_MEOWRES']._serialized_end=64
  _globals['_NOTIF']._serialized_start=66
  _globals['_NOTIF']._serialized_end=109
# @@protoc_insertion_point(module_scope)
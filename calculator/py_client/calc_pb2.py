# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: calc.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ncalc.proto\"!\n\x07\x43\x61lcReq\x12\n\n\x02n1\x18\x01 \x01(\x02\x12\n\n\x02n2\x18\x02 \x01(\x02\"\x13\n\x06\x41\x64\x64Res\x12\t\n\x01r\x18\x01 \x01(\x02\"\x13\n\x06SubRes\x12\t\n\x01r\x18\x01 \x01(\x02\"\x13\n\x06\x44ivRes\x12\t\n\x01r\x18\x01 \x01(\x02\"\x13\n\x06MulRes\x12\t\n\x01r\x18\x01 \x01(\x02\x32|\n\nCalculator\x12\x1a\n\x03\x41\x64\x64\x12\x08.CalcReq\x1a\x07.AddRes\"\x00\x12\x1a\n\x03Sub\x12\x08.CalcReq\x1a\x07.SubRes\"\x00\x12\x1a\n\x03\x44iv\x12\x08.CalcReq\x1a\x07.DivRes\"\x00\x12\x1a\n\x03Mul\x12\x08.CalcReq\x1a\x07.MulRes\"\x00\x42\x0cZ\ncalc/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'calc_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\ncalc/proto'
  _globals['_CALCREQ']._serialized_start=14
  _globals['_CALCREQ']._serialized_end=47
  _globals['_ADDRES']._serialized_start=49
  _globals['_ADDRES']._serialized_end=68
  _globals['_SUBRES']._serialized_start=70
  _globals['_SUBRES']._serialized_end=89
  _globals['_DIVRES']._serialized_start=91
  _globals['_DIVRES']._serialized_end=110
  _globals['_MULRES']._serialized_start=112
  _globals['_MULRES']._serialized_end=131
  _globals['_CALCULATOR']._serialized_start=133
  _globals['_CALCULATOR']._serialized_end=257
# @@protoc_insertion_point(module_scope)

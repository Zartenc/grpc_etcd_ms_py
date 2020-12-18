# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: zarten.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='zarten.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0czarten.proto\"#\n\rZartenRequest\x12\x12\n\nzhihu_name\x18\x01 \x01(\t\"0\n\x0eZartenResponse\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x10\n\x08homepage\x18\x02 \x01(\t24\n\x06Zarten\x12*\n\x07GetInfo\x12\x0e.ZartenRequest\x1a\x0f.ZartenResponseb\x06proto3'
)




_ZARTENREQUEST = _descriptor.Descriptor(
  name='ZartenRequest',
  full_name='ZartenRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='zhihu_name', full_name='ZartenRequest.zhihu_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=16,
  serialized_end=51,
)


_ZARTENRESPONSE = _descriptor.Descriptor(
  name='ZartenResponse',
  full_name='ZartenResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='ZartenResponse.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='homepage', full_name='ZartenResponse.homepage', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=53,
  serialized_end=101,
)

DESCRIPTOR.message_types_by_name['ZartenRequest'] = _ZARTENREQUEST
DESCRIPTOR.message_types_by_name['ZartenResponse'] = _ZARTENRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ZartenRequest = _reflection.GeneratedProtocolMessageType('ZartenRequest', (_message.Message,), {
  'DESCRIPTOR' : _ZARTENREQUEST,
  '__module__' : 'zarten_pb2'
  # @@protoc_insertion_point(class_scope:ZartenRequest)
  })
_sym_db.RegisterMessage(ZartenRequest)

ZartenResponse = _reflection.GeneratedProtocolMessageType('ZartenResponse', (_message.Message,), {
  'DESCRIPTOR' : _ZARTENRESPONSE,
  '__module__' : 'zarten_pb2'
  # @@protoc_insertion_point(class_scope:ZartenResponse)
  })
_sym_db.RegisterMessage(ZartenResponse)



_ZARTEN = _descriptor.ServiceDescriptor(
  name='Zarten',
  full_name='Zarten',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=103,
  serialized_end=155,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetInfo',
    full_name='Zarten.GetInfo',
    index=0,
    containing_service=None,
    input_type=_ZARTENREQUEST,
    output_type=_ZARTENRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ZARTEN)

DESCRIPTOR.services_by_name['Zarten'] = _ZARTEN

# @@protoc_insertion_point(module_scope)
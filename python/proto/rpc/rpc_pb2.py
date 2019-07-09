# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/rpc/rpc.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/rpc/rpc.proto',
  package='iotexrpc',
  syntax='proto3',
  serialized_options=_b('Z3github.com/iotexproject/iotex-proto/golang/iotexrpc'),
  serialized_pb=_b('\n\x13proto/rpc/rpc.proto\x12\x08iotexrpc\x1a\x1fgoogle/protobuf/timestamp.proto\"\'\n\tBlockSync\x12\r\n\x05start\x18\x02 \x01(\x04\x12\x0b\n\x03\x65nd\x18\x03 \x01(\x04\"\x9b\x01\n\x0c\x42roadcastMsg\x12\x10\n\x08\x63hain_id\x18\x01 \x01(\r\x12\'\n\x08msg_type\x18\x02 \x01(\x0e\x32\x15.iotexrpc.MessageType\x12\x10\n\x08msg_body\x18\x03 \x01(\x0c\x12\x0f\n\x07peer_id\x18\x04 \x01(\t\x12-\n\ttimestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xa7\x01\n\nUnicastMsg\x12\x10\n\x08\x63hain_id\x18\x01 \x01(\r\x12\x0c\n\x04\x61\x64\x64r\x18\x02 \x01(\t\x12\'\n\x08msg_type\x18\x03 \x01(\x0e\x32\x15.iotexrpc.MessageType\x12\x10\n\x08msg_body\x18\x04 \x01(\x0c\x12\x0f\n\x07peer_id\x18\x05 \x01(\t\x12-\n\ttimestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp*^\n\x0bMessageType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\n\n\x06\x41\x43TION\x10\x01\x12\t\n\x05\x42LOCK\x10\x02\x12\r\n\tCONSENSUS\x10\x03\x12\x11\n\rBLOCK_REQUEST\x10\x04\x12\t\n\x04TEST\x10\x91NB5Z3github.com/iotexproject/iotex-proto/golang/iotexrpcb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_MESSAGETYPE = _descriptor.EnumDescriptor(
  name='MessageType',
  full_name='iotexrpc.MessageType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ACTION', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BLOCK', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONSENSUS', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BLOCK_REQUEST', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TEST', index=5, number=10001,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=435,
  serialized_end=529,
)
_sym_db.RegisterEnumDescriptor(_MESSAGETYPE)

MessageType = enum_type_wrapper.EnumTypeWrapper(_MESSAGETYPE)
UNKNOWN = 0
ACTION = 1
BLOCK = 2
CONSENSUS = 3
BLOCK_REQUEST = 4
TEST = 10001



_BLOCKSYNC = _descriptor.Descriptor(
  name='BlockSync',
  full_name='iotexrpc.BlockSync',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='start', full_name='iotexrpc.BlockSync.start', index=0,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end', full_name='iotexrpc.BlockSync.end', index=1,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=66,
  serialized_end=105,
)


_BROADCASTMSG = _descriptor.Descriptor(
  name='BroadcastMsg',
  full_name='iotexrpc.BroadcastMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='chain_id', full_name='iotexrpc.BroadcastMsg.chain_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='msg_type', full_name='iotexrpc.BroadcastMsg.msg_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='msg_body', full_name='iotexrpc.BroadcastMsg.msg_body', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='peer_id', full_name='iotexrpc.BroadcastMsg.peer_id', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='iotexrpc.BroadcastMsg.timestamp', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=108,
  serialized_end=263,
)


_UNICASTMSG = _descriptor.Descriptor(
  name='UnicastMsg',
  full_name='iotexrpc.UnicastMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='chain_id', full_name='iotexrpc.UnicastMsg.chain_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addr', full_name='iotexrpc.UnicastMsg.addr', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='msg_type', full_name='iotexrpc.UnicastMsg.msg_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='msg_body', full_name='iotexrpc.UnicastMsg.msg_body', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='peer_id', full_name='iotexrpc.UnicastMsg.peer_id', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='iotexrpc.UnicastMsg.timestamp', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=266,
  serialized_end=433,
)

_BROADCASTMSG.fields_by_name['msg_type'].enum_type = _MESSAGETYPE
_BROADCASTMSG.fields_by_name['timestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_UNICASTMSG.fields_by_name['msg_type'].enum_type = _MESSAGETYPE
_UNICASTMSG.fields_by_name['timestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['BlockSync'] = _BLOCKSYNC
DESCRIPTOR.message_types_by_name['BroadcastMsg'] = _BROADCASTMSG
DESCRIPTOR.message_types_by_name['UnicastMsg'] = _UNICASTMSG
DESCRIPTOR.enum_types_by_name['MessageType'] = _MESSAGETYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

BlockSync = _reflection.GeneratedProtocolMessageType('BlockSync', (_message.Message,), dict(
  DESCRIPTOR = _BLOCKSYNC,
  __module__ = 'proto.rpc.rpc_pb2'
  # @@protoc_insertion_point(class_scope:iotexrpc.BlockSync)
  ))
_sym_db.RegisterMessage(BlockSync)

BroadcastMsg = _reflection.GeneratedProtocolMessageType('BroadcastMsg', (_message.Message,), dict(
  DESCRIPTOR = _BROADCASTMSG,
  __module__ = 'proto.rpc.rpc_pb2'
  # @@protoc_insertion_point(class_scope:iotexrpc.BroadcastMsg)
  ))
_sym_db.RegisterMessage(BroadcastMsg)

UnicastMsg = _reflection.GeneratedProtocolMessageType('UnicastMsg', (_message.Message,), dict(
  DESCRIPTOR = _UNICASTMSG,
  __module__ = 'proto.rpc.rpc_pb2'
  # @@protoc_insertion_point(class_scope:iotexrpc.UnicastMsg)
  ))
_sym_db.RegisterMessage(UnicastMsg)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

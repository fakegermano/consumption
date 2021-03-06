# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: meterusage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


DESCRIPTOR = _descriptor.FileDescriptor(
    name="meterusage.proto",
    package="proto",
    syntax="proto3",
    serialized_options=b"Z(github.com/fakegermano/consumption/proto",
    create_key=_descriptor._internal_create_key,
    serialized_pb=b'\n\x10meterusage.proto\x12\x05proto"2\n\x11MeterUsageRequest\x12\r\n\x05limit\x18\x01 \x01(\x05\x12\x0e\n\x06offset\x18\x02 \x01(\x05"7\n\x12MeterUsageResponse\x12!\n\x06usages\x18\x01 \x03(\x0b\x32\x11.proto.MeterUsage")\n\nMeterUsage\x12\x0c\n\x04time\x18\x01 \x01(\x03\x12\r\n\x05usage\x18\x02 \x01(\x01\x32Q\n\x11MeterUsageService\x12<\n\x03Get\x12\x18.proto.MeterUsageRequest\x1a\x19.proto.MeterUsageResponse"\x00\x42*Z(github.com/fakegermano/consumption/protob\x06proto3',
)


_METERUSAGEREQUEST = _descriptor.Descriptor(
    name="MeterUsageRequest",
    full_name="proto.MeterUsageRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="limit",
            full_name="proto.MeterUsageRequest.limit",
            index=0,
            number=1,
            type=5,
            cpp_type=1,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="offset",
            full_name="proto.MeterUsageRequest.offset",
            index=1,
            number=2,
            type=5,
            cpp_type=1,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=27,
    serialized_end=77,
)


_METERUSAGERESPONSE = _descriptor.Descriptor(
    name="MeterUsageResponse",
    full_name="proto.MeterUsageResponse",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="usages",
            full_name="proto.MeterUsageResponse.usages",
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=79,
    serialized_end=134,
)


_METERUSAGE = _descriptor.Descriptor(
    name="MeterUsage",
    full_name="proto.MeterUsage",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="time",
            full_name="proto.MeterUsage.time",
            index=0,
            number=1,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="usage",
            full_name="proto.MeterUsage.usage",
            index=1,
            number=2,
            type=1,
            cpp_type=5,
            label=1,
            has_default_value=False,
            default_value=float(0),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=136,
    serialized_end=177,
)

_METERUSAGERESPONSE.fields_by_name["usages"].message_type = _METERUSAGE
DESCRIPTOR.message_types_by_name["MeterUsageRequest"] = _METERUSAGEREQUEST
DESCRIPTOR.message_types_by_name["MeterUsageResponse"] = _METERUSAGERESPONSE
DESCRIPTOR.message_types_by_name["MeterUsage"] = _METERUSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MeterUsageRequest = _reflection.GeneratedProtocolMessageType(
    "MeterUsageRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _METERUSAGEREQUEST,
        "__module__": "meterusage_pb2"
        # @@protoc_insertion_point(class_scope:proto.MeterUsageRequest)
    },
)
_sym_db.RegisterMessage(MeterUsageRequest)

MeterUsageResponse = _reflection.GeneratedProtocolMessageType(
    "MeterUsageResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _METERUSAGERESPONSE,
        "__module__": "meterusage_pb2"
        # @@protoc_insertion_point(class_scope:proto.MeterUsageResponse)
    },
)
_sym_db.RegisterMessage(MeterUsageResponse)

MeterUsage = _reflection.GeneratedProtocolMessageType(
    "MeterUsage",
    (_message.Message,),
    {
        "DESCRIPTOR": _METERUSAGE,
        "__module__": "meterusage_pb2"
        # @@protoc_insertion_point(class_scope:proto.MeterUsage)
    },
)
_sym_db.RegisterMessage(MeterUsage)


DESCRIPTOR._options = None

_METERUSAGESERVICE = _descriptor.ServiceDescriptor(
    name="MeterUsageService",
    full_name="proto.MeterUsageService",
    file=DESCRIPTOR,
    index=0,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
    serialized_start=179,
    serialized_end=260,
    methods=[
        _descriptor.MethodDescriptor(
            name="Get",
            full_name="proto.MeterUsageService.Get",
            index=0,
            containing_service=None,
            input_type=_METERUSAGEREQUEST,
            output_type=_METERUSAGERESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
    ],
)
_sym_db.RegisterServiceDescriptor(_METERUSAGESERVICE)

DESCRIPTOR.services_by_name["MeterUsageService"] = _METERUSAGESERVICE

# @@protoc_insertion_point(module_scope)

// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package contractmanagerv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Failure            protoreflect.MessageDescriptor
	fd_Failure_channel_id protoreflect.FieldDescriptor
	fd_Failure_address    protoreflect.FieldDescriptor
	fd_Failure_id         protoreflect.FieldDescriptor
	fd_Failure_ack_id     protoreflect.FieldDescriptor
	fd_Failure_ack_type   protoreflect.FieldDescriptor
)

func init() {
	file_nolus_contractmanager_v1_failure_proto_init()
	md_Failure = File_nolus_contractmanager_v1_failure_proto.Messages().ByName("Failure")
	fd_Failure_channel_id = md_Failure.Fields().ByName("channel_id")
	fd_Failure_address = md_Failure.Fields().ByName("address")
	fd_Failure_id = md_Failure.Fields().ByName("id")
	fd_Failure_ack_id = md_Failure.Fields().ByName("ack_id")
	fd_Failure_ack_type = md_Failure.Fields().ByName("ack_type")
}

var _ protoreflect.Message = (*fastReflection_Failure)(nil)

type fastReflection_Failure Failure

func (x *Failure) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Failure)(x)
}

func (x *Failure) slowProtoReflect() protoreflect.Message {
	mi := &file_nolus_contractmanager_v1_failure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Failure_messageType fastReflection_Failure_messageType
var _ protoreflect.MessageType = fastReflection_Failure_messageType{}

type fastReflection_Failure_messageType struct{}

func (x fastReflection_Failure_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Failure)(nil)
}
func (x fastReflection_Failure_messageType) New() protoreflect.Message {
	return new(fastReflection_Failure)
}
func (x fastReflection_Failure_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Failure
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Failure) Descriptor() protoreflect.MessageDescriptor {
	return md_Failure
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Failure) Type() protoreflect.MessageType {
	return _fastReflection_Failure_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Failure) New() protoreflect.Message {
	return new(fastReflection_Failure)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Failure) Interface() protoreflect.ProtoMessage {
	return (*Failure)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Failure) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ChannelId != "" {
		value := protoreflect.ValueOfString(x.ChannelId)
		if !f(fd_Failure_channel_id, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_Failure_address, value) {
			return
		}
	}
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Failure_id, value) {
			return
		}
	}
	if x.AckId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AckId)
		if !f(fd_Failure_ack_id, value) {
			return
		}
	}
	if x.AckType != "" {
		value := protoreflect.ValueOfString(x.AckType)
		if !f(fd_Failure_ack_type, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Failure) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "nolus.contractmanager.v1.Failure.channel_id":
		return x.ChannelId != ""
	case "nolus.contractmanager.v1.Failure.address":
		return x.Address != ""
	case "nolus.contractmanager.v1.Failure.id":
		return x.Id != uint64(0)
	case "nolus.contractmanager.v1.Failure.ack_id":
		return x.AckId != uint64(0)
	case "nolus.contractmanager.v1.Failure.ack_type":
		return x.AckType != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nolus.contractmanager.v1.Failure"))
		}
		panic(fmt.Errorf("message nolus.contractmanager.v1.Failure does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Failure) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "nolus.contractmanager.v1.Failure.channel_id":
		x.ChannelId = ""
	case "nolus.contractmanager.v1.Failure.address":
		x.Address = ""
	case "nolus.contractmanager.v1.Failure.id":
		x.Id = uint64(0)
	case "nolus.contractmanager.v1.Failure.ack_id":
		x.AckId = uint64(0)
	case "nolus.contractmanager.v1.Failure.ack_type":
		x.AckType = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nolus.contractmanager.v1.Failure"))
		}
		panic(fmt.Errorf("message nolus.contractmanager.v1.Failure does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Failure) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "nolus.contractmanager.v1.Failure.channel_id":
		value := x.ChannelId
		return protoreflect.ValueOfString(value)
	case "nolus.contractmanager.v1.Failure.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "nolus.contractmanager.v1.Failure.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "nolus.contractmanager.v1.Failure.ack_id":
		value := x.AckId
		return protoreflect.ValueOfUint64(value)
	case "nolus.contractmanager.v1.Failure.ack_type":
		value := x.AckType
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nolus.contractmanager.v1.Failure"))
		}
		panic(fmt.Errorf("message nolus.contractmanager.v1.Failure does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Failure) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "nolus.contractmanager.v1.Failure.channel_id":
		x.ChannelId = value.Interface().(string)
	case "nolus.contractmanager.v1.Failure.address":
		x.Address = value.Interface().(string)
	case "nolus.contractmanager.v1.Failure.id":
		x.Id = value.Uint()
	case "nolus.contractmanager.v1.Failure.ack_id":
		x.AckId = value.Uint()
	case "nolus.contractmanager.v1.Failure.ack_type":
		x.AckType = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nolus.contractmanager.v1.Failure"))
		}
		panic(fmt.Errorf("message nolus.contractmanager.v1.Failure does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Failure) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "nolus.contractmanager.v1.Failure.channel_id":
		panic(fmt.Errorf("field channel_id of message nolus.contractmanager.v1.Failure is not mutable"))
	case "nolus.contractmanager.v1.Failure.address":
		panic(fmt.Errorf("field address of message nolus.contractmanager.v1.Failure is not mutable"))
	case "nolus.contractmanager.v1.Failure.id":
		panic(fmt.Errorf("field id of message nolus.contractmanager.v1.Failure is not mutable"))
	case "nolus.contractmanager.v1.Failure.ack_id":
		panic(fmt.Errorf("field ack_id of message nolus.contractmanager.v1.Failure is not mutable"))
	case "nolus.contractmanager.v1.Failure.ack_type":
		panic(fmt.Errorf("field ack_type of message nolus.contractmanager.v1.Failure is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nolus.contractmanager.v1.Failure"))
		}
		panic(fmt.Errorf("message nolus.contractmanager.v1.Failure does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Failure) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "nolus.contractmanager.v1.Failure.channel_id":
		return protoreflect.ValueOfString("")
	case "nolus.contractmanager.v1.Failure.address":
		return protoreflect.ValueOfString("")
	case "nolus.contractmanager.v1.Failure.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "nolus.contractmanager.v1.Failure.ack_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "nolus.contractmanager.v1.Failure.ack_type":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nolus.contractmanager.v1.Failure"))
		}
		panic(fmt.Errorf("message nolus.contractmanager.v1.Failure does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Failure) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in nolus.contractmanager.v1.Failure", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Failure) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Failure) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Failure) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Failure) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Failure)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ChannelId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		if x.AckId != 0 {
			n += 1 + runtime.Sov(uint64(x.AckId))
		}
		l = len(x.AckType)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Failure)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.AckType) > 0 {
			i -= len(x.AckType)
			copy(dAtA[i:], x.AckType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AckType)))
			i--
			dAtA[i] = 0x2a
		}
		if x.AckId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AckId))
			i--
			dAtA[i] = 0x20
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ChannelId) > 0 {
			i -= len(x.ChannelId)
			copy(dAtA[i:], x.ChannelId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ChannelId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Failure)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Failure: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Failure: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ChannelId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AckId", wireType)
				}
				x.AckId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AckId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AckType", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AckType = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: nolus/contractmanager/v1/failure.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Deprecated. Used only for migration purposes.
type Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ChannelId
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Address of the failed contract
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// id of the failure under specific address
	Id uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// ACK id to restore
	AckId uint64 `protobuf:"varint,4,opt,name=ack_id,json=ackId,proto3" json:"ack_id,omitempty"`
	// Acknowledgement type
	AckType string `protobuf:"bytes,5,opt,name=ack_type,json=ackType,proto3" json:"ack_type,omitempty"`
}

func (x *Failure) Reset() {
	*x = Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nolus_contractmanager_v1_failure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_nolus_contractmanager_v1_failure_proto_rawDescGZIP(), []int{0}
}

func (x *Failure) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Failure) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Failure) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Failure) GetAckId() uint64 {
	if x != nil {
		return x.AckId
	}
	return 0
}

func (x *Failure) GetAckType() string {
	if x != nil {
		return x.AckType
	}
	return ""
}

var File_nolus_contractmanager_v1_failure_proto protoreflect.FileDescriptor

var file_nolus_contractmanager_v1_failure_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x6f, 0x6c, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6e, 0x6f, 0x6c, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x22, 0x84, 0x01, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x42, 0xeb, 0x01, 0x0a, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x6f, 0x6c, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x6c,
	0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x43, 0x58, 0xaa, 0x02, 0x18,
	0x4e, 0x6f, 0x6c, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x4e, 0x6f, 0x6c, 0x75, 0x73,
	0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x4e, 0x6f, 0x6c, 0x75, 0x73, 0x5c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x4e, 0x6f, 0x6c,
	0x75, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nolus_contractmanager_v1_failure_proto_rawDescOnce sync.Once
	file_nolus_contractmanager_v1_failure_proto_rawDescData = file_nolus_contractmanager_v1_failure_proto_rawDesc
)

func file_nolus_contractmanager_v1_failure_proto_rawDescGZIP() []byte {
	file_nolus_contractmanager_v1_failure_proto_rawDescOnce.Do(func() {
		file_nolus_contractmanager_v1_failure_proto_rawDescData = protoimpl.X.CompressGZIP(file_nolus_contractmanager_v1_failure_proto_rawDescData)
	})
	return file_nolus_contractmanager_v1_failure_proto_rawDescData
}

var file_nolus_contractmanager_v1_failure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nolus_contractmanager_v1_failure_proto_goTypes = []interface{}{
	(*Failure)(nil), // 0: nolus.contractmanager.v1.Failure
}
var file_nolus_contractmanager_v1_failure_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nolus_contractmanager_v1_failure_proto_init() }
func file_nolus_contractmanager_v1_failure_proto_init() {
	if File_nolus_contractmanager_v1_failure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nolus_contractmanager_v1_failure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Failure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nolus_contractmanager_v1_failure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nolus_contractmanager_v1_failure_proto_goTypes,
		DependencyIndexes: file_nolus_contractmanager_v1_failure_proto_depIdxs,
		MessageInfos:      file_nolus_contractmanager_v1_failure_proto_msgTypes,
	}.Build()
	File_nolus_contractmanager_v1_failure_proto = out.File
	file_nolus_contractmanager_v1_failure_proto_rawDesc = nil
	file_nolus_contractmanager_v1_failure_proto_goTypes = nil
	file_nolus_contractmanager_v1_failure_proto_depIdxs = nil
}

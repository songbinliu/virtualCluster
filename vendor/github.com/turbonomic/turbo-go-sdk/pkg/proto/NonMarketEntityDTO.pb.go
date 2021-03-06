// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NonMarketEntityDTO.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// enumerate non market entities
type NonMarketEntityDTO_NonMarketEntityType int32

const (
	NonMarketEntityDTO_CLOUD_SERVICE NonMarketEntityDTO_NonMarketEntityType = 0
	NonMarketEntityDTO_WORKFLOW      NonMarketEntityDTO_NonMarketEntityType = 1
	NonMarketEntityDTO_ACCOUNT       NonMarketEntityDTO_NonMarketEntityType = 2
)

var NonMarketEntityDTO_NonMarketEntityType_name = map[int32]string{
	0: "CLOUD_SERVICE",
	1: "WORKFLOW",
	2: "ACCOUNT",
}
var NonMarketEntityDTO_NonMarketEntityType_value = map[string]int32{
	"CLOUD_SERVICE": 0,
	"WORKFLOW":      1,
	"ACCOUNT":       2,
}

func (x NonMarketEntityDTO_NonMarketEntityType) Enum() *NonMarketEntityDTO_NonMarketEntityType {
	p := new(NonMarketEntityDTO_NonMarketEntityType)
	*p = x
	return p
}
func (x NonMarketEntityDTO_NonMarketEntityType) String() string {
	return proto.EnumName(NonMarketEntityDTO_NonMarketEntityType_name, int32(x))
}
func (x *NonMarketEntityDTO_NonMarketEntityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_NonMarketEntityType_value, data, "NonMarketEntityDTO_NonMarketEntityType")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_NonMarketEntityType(value)
	return nil
}
func (NonMarketEntityDTO_NonMarketEntityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0}
}

// The NonMarketEntityDTO message represents an Entity discovered in the target that your probe is
// monitoring and this entity does not participate in the market.
//
// Each entity must have a unique ID to identify it in the Operations Manager repository.
// Many targets provide unique IDs for their entities, or you can generate your own.
// To guarantee that it's unique, you can give the ID a prefix that identifies your
// probe and the given target.
//
// Specify entity type by setting a 'NonMarketEntityType' value to the 'entityType' field.
//
// The 'displayName' value appears in the product GUI and in reports to identify the entity.
//
type NonMarketEntityDTO struct {
	EntityType  *NonMarketEntityDTO_NonMarketEntityType `protobuf:"varint,1,req,name=entityType,enum=common_dto.NonMarketEntityDTO_NonMarketEntityType" json:"entityType,omitempty"`
	Id          *string                                 `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	DisplayName *string                                 `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
	Description *string                                 `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// Collection of entity type's specific data.
	//
	// Types that are valid to be assigned to EntityData:
	//	*NonMarketEntityDTO_CloudServiceData_
	//	*NonMarketEntityDTO_WorkflowData_
	EntityData       isNonMarketEntityDTO_EntityData `protobuf_oneof:"entity_data"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *NonMarketEntityDTO) Reset()                    { *m = NonMarketEntityDTO{} }
func (m *NonMarketEntityDTO) String() string            { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO) ProtoMessage()               {}
func (*NonMarketEntityDTO) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type isNonMarketEntityDTO_EntityData interface {
	isNonMarketEntityDTO_EntityData()
}

type NonMarketEntityDTO_CloudServiceData_ struct {
	CloudServiceData *NonMarketEntityDTO_CloudServiceData `protobuf:"bytes,500,opt,name=cloud_service_data,json=cloudServiceData,oneof"`
}
type NonMarketEntityDTO_WorkflowData_ struct {
	WorkflowData *NonMarketEntityDTO_WorkflowData `protobuf:"bytes,501,opt,name=workflow_data,json=workflowData,oneof"`
}

func (*NonMarketEntityDTO_CloudServiceData_) isNonMarketEntityDTO_EntityData() {}
func (*NonMarketEntityDTO_WorkflowData_) isNonMarketEntityDTO_EntityData()     {}

func (m *NonMarketEntityDTO) GetEntityData() isNonMarketEntityDTO_EntityData {
	if m != nil {
		return m.EntityData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetEntityType() NonMarketEntityDTO_NonMarketEntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return NonMarketEntityDTO_CLOUD_SERVICE
}

func (m *NonMarketEntityDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *NonMarketEntityDTO) GetCloudServiceData() *NonMarketEntityDTO_CloudServiceData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_CloudServiceData_); ok {
		return x.CloudServiceData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetWorkflowData() *NonMarketEntityDTO_WorkflowData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_WorkflowData_); ok {
		return x.WorkflowData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NonMarketEntityDTO) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NonMarketEntityDTO_OneofMarshaler, _NonMarketEntityDTO_OneofUnmarshaler, _NonMarketEntityDTO_OneofSizer, []interface{}{
		(*NonMarketEntityDTO_CloudServiceData_)(nil),
		(*NonMarketEntityDTO_WorkflowData_)(nil),
	}
}

func _NonMarketEntityDTO_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NonMarketEntityDTO)
	// entity_data
	switch x := m.EntityData.(type) {
	case *NonMarketEntityDTO_CloudServiceData_:
		b.EncodeVarint(500<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CloudServiceData); err != nil {
			return err
		}
	case *NonMarketEntityDTO_WorkflowData_:
		b.EncodeVarint(501<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WorkflowData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NonMarketEntityDTO.EntityData has unexpected type %T", x)
	}
	return nil
}

func _NonMarketEntityDTO_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NonMarketEntityDTO)
	switch tag {
	case 500: // entity_data.cloud_service_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NonMarketEntityDTO_CloudServiceData)
		err := b.DecodeMessage(msg)
		m.EntityData = &NonMarketEntityDTO_CloudServiceData_{msg}
		return true, err
	case 501: // entity_data.workflow_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NonMarketEntityDTO_WorkflowData)
		err := b.DecodeMessage(msg)
		m.EntityData = &NonMarketEntityDTO_WorkflowData_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NonMarketEntityDTO_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NonMarketEntityDTO)
	// entity_data
	switch x := m.EntityData.(type) {
	case *NonMarketEntityDTO_CloudServiceData_:
		s := proto.Size(x.CloudServiceData)
		n += proto.SizeVarint(500<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NonMarketEntityDTO_WorkflowData_:
		s := proto.Size(x.WorkflowData)
		n += proto.SizeVarint(501<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NonMarketEntityDTO_CloudServiceData struct {
	CloudProvider    *string                                          `protobuf:"bytes,1,req,name=cloud_provider,json=cloudProvider" json:"cloud_provider,omitempty"`
	Spend            *float32                                         `protobuf:"fixed32,2,req,name=spend" json:"spend,omitempty"`
	BillingData      *NonMarketEntityDTO_CloudServiceData_BillingData `protobuf:"bytes,3,opt,name=billing_data,json=billingData" json:"billing_data,omitempty"`
	XXX_unrecognized []byte                                           `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData) Reset()         { *m = NonMarketEntityDTO_CloudServiceData{} }
func (m *NonMarketEntityDTO_CloudServiceData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_CloudServiceData) ProtoMessage()    {}
func (*NonMarketEntityDTO_CloudServiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData) GetCloudProvider() string {
	if m != nil && m.CloudProvider != nil {
		return *m.CloudProvider
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetSpend() float32 {
	if m != nil && m.Spend != nil {
		return *m.Spend
	}
	return 0
}

func (m *NonMarketEntityDTO_CloudServiceData) GetBillingData() *NonMarketEntityDTO_CloudServiceData_BillingData {
	if m != nil {
		return m.BillingData
	}
	return nil
}

type NonMarketEntityDTO_CloudServiceData_BillingData struct {
	VirtualMachines   []*EntityDTO `protobuf:"bytes,1,rep,name=virtual_machines,json=virtualMachines" json:"virtual_machines,omitempty"`
	ReservedInstances []*EntityDTO `protobuf:"bytes,2,rep,name=reserved_instances,json=reservedInstances" json:"reserved_instances,omitempty"`
	XXX_unrecognized  []byte       `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) Reset() {
	*m = NonMarketEntityDTO_CloudServiceData_BillingData{}
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) String() string {
	return proto.CompactTextString(m)
}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) ProtoMessage() {}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetVirtualMachines() []*EntityDTO {
	if m != nil {
		return m.VirtualMachines
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetReservedInstances() []*EntityDTO {
	if m != nil {
		return m.ReservedInstances
	}
	return nil
}

type NonMarketEntityDTO_WorkflowData struct {
	Param            []*NonMarketEntityDTO_Parameter `protobuf:"bytes,1,rep,name=param" json:"param,omitempty"`
	Property         []*NonMarketEntityDTO_Property  `protobuf:"bytes,2,rep,name=property" json:"property,omitempty"`
	EntityType       *EntityDTO_EntityType           `protobuf:"varint,3,opt,name=entityType,enum=common_dto.EntityDTO_EntityType" json:"entityType,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *NonMarketEntityDTO_WorkflowData) Reset()         { *m = NonMarketEntityDTO_WorkflowData{} }
func (m *NonMarketEntityDTO_WorkflowData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_WorkflowData) ProtoMessage()    {}
func (*NonMarketEntityDTO_WorkflowData) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 1}
}

func (m *NonMarketEntityDTO_WorkflowData) GetParam() []*NonMarketEntityDTO_Parameter {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *NonMarketEntityDTO_WorkflowData) GetProperty() []*NonMarketEntityDTO_Property {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *NonMarketEntityDTO_WorkflowData) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

type NonMarketEntityDTO_Parameter struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Description      *string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Type             *string `protobuf:"bytes,3,req,name=type,def=String" json:"type,omitempty"`
	Mandatory        *bool   `protobuf:"varint,4,opt,name=mandatory,def=1" json:"mandatory,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NonMarketEntityDTO_Parameter) Reset()                    { *m = NonMarketEntityDTO_Parameter{} }
func (m *NonMarketEntityDTO_Parameter) String() string            { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_Parameter) ProtoMessage()               {}
func (*NonMarketEntityDTO_Parameter) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0, 2} }

const Default_NonMarketEntityDTO_Parameter_Type string = "String"
const Default_NonMarketEntityDTO_Parameter_Mandatory bool = true

func (m *NonMarketEntityDTO_Parameter) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NonMarketEntityDTO_Parameter) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *NonMarketEntityDTO_Parameter) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_NonMarketEntityDTO_Parameter_Type
}

func (m *NonMarketEntityDTO_Parameter) GetMandatory() bool {
	if m != nil && m.Mandatory != nil {
		return *m.Mandatory
	}
	return Default_NonMarketEntityDTO_Parameter_Mandatory
}

type NonMarketEntityDTO_Property struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NonMarketEntityDTO_Property) Reset()                    { *m = NonMarketEntityDTO_Property{} }
func (m *NonMarketEntityDTO_Property) String() string            { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_Property) ProtoMessage()               {}
func (*NonMarketEntityDTO_Property) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0, 3} }

func (m *NonMarketEntityDTO_Property) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NonMarketEntityDTO_Property) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*NonMarketEntityDTO)(nil), "common_dto.NonMarketEntityDTO")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData_BillingData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData.BillingData")
	proto.RegisterType((*NonMarketEntityDTO_WorkflowData)(nil), "common_dto.NonMarketEntityDTO.WorkflowData")
	proto.RegisterType((*NonMarketEntityDTO_Parameter)(nil), "common_dto.NonMarketEntityDTO.Parameter")
	proto.RegisterType((*NonMarketEntityDTO_Property)(nil), "common_dto.NonMarketEntityDTO.Property")
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_NonMarketEntityType", NonMarketEntityDTO_NonMarketEntityType_name, NonMarketEntityDTO_NonMarketEntityType_value)
}

func init() { proto.RegisterFile("NonMarketEntityDTO.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0xb2, 0x0e, 0x9a, 0x9b, 0xb6, 0xeb, 0xcc, 0x90, 0xa2, 0xbc, 0x50, 0x4d, 0x42, 0x54,
	0x42, 0xca, 0xa4, 0x8a, 0xa7, 0x21, 0xa1, 0xd1, 0x74, 0x88, 0x89, 0x6d, 0x9d, 0xdc, 0x8e, 0xbd,
	0x11, 0xbc, 0xc4, 0x1b, 0xd6, 0x12, 0x3b, 0x72, 0xdc, 0x4e, 0x7d, 0xe4, 0x23, 0xf8, 0x02, 0x3e,
	0x86, 0x2f, 0xe0, 0x4f, 0xe0, 0x1d, 0x25, 0xce, 0xda, 0xd0, 0x55, 0x54, 0xbc, 0xf9, 0x9e, 0xdc,
	0x73, 0xee, 0xf5, 0x49, 0x4e, 0xc0, 0x39, 0x13, 0xfc, 0x94, 0xc8, 0x5b, 0xaa, 0x8e, 0xb8, 0x62,
	0x6a, 0x36, 0x18, 0x0f, 0xbd, 0x54, 0x0a, 0x25, 0x10, 0x84, 0x22, 0x49, 0x04, 0x0f, 0x22, 0x25,
	0xdc, 0x6d, 0xbf, 0x38, 0xcf, 0x1f, 0xee, 0x7d, 0xb7, 0x00, 0x3d, 0x64, 0x22, 0x0c, 0x40, 0x8b,
	0x62, 0x3c, 0x4b, 0xa9, 0x63, 0x74, 0xcc, 0x6e, 0xab, 0xd7, 0xf3, 0x16, 0x42, 0xde, 0x8a, 0x69,
	0x4b, 0x50, 0xce, 0xc4, 0x15, 0x15, 0xd4, 0x02, 0x93, 0x45, 0x8e, 0xd9, 0x31, 0xbb, 0x16, 0x36,
	0x59, 0x84, 0x3a, 0x60, 0x47, 0x2c, 0x4b, 0x63, 0x32, 0x3b, 0x23, 0x09, 0x75, 0x36, 0x3b, 0x46,
	0xd7, 0xc2, 0x55, 0xa8, 0xe8, 0xa0, 0x59, 0x28, 0x59, 0xaa, 0x98, 0xe0, 0x4e, 0xad, 0xec, 0x58,
	0x40, 0xe8, 0x33, 0xa0, 0x30, 0x16, 0x93, 0x28, 0xc8, 0xa8, 0x9c, 0xb2, 0x90, 0x06, 0x11, 0x51,
	0xc4, 0xf9, 0x95, 0x6b, 0xd9, 0xbd, 0xfd, 0x35, 0x0b, 0xfb, 0x39, 0x73, 0xa4, 0x89, 0x03, 0xa2,
	0xc8, 0xfb, 0x0d, 0xdc, 0x0e, 0x97, 0x30, 0x34, 0x82, 0xe6, 0x9d, 0x90, 0xb7, 0xd7, 0xb1, 0xb8,
	0xd3, 0xe2, 0xbf, 0xb5, 0xf8, 0xcb, 0x35, 0xe2, 0x97, 0x25, 0xa9, 0x14, 0x6e, 0xdc, 0x55, 0x6a,
	0xf7, 0x87, 0x09, 0xed, 0xe5, 0xe9, 0xe8, 0x39, 0xb4, 0xf4, 0x5d, 0x52, 0x29, 0xa6, 0x2c, 0xa2,
	0xb2, 0xf0, 0xdd, 0xc2, 0xcd, 0x02, 0x3d, 0x2f, 0x41, 0xb4, 0x0b, 0x5b, 0x59, 0x4a, 0xb9, 0x76,
	0xd2, 0xc4, 0xba, 0x40, 0x9f, 0xa0, 0x71, 0xc5, 0xe2, 0x98, 0xf1, 0x1b, 0xbd, 0xa5, 0x5e, 0xf2,
	0xf5, 0x7f, 0x3a, 0xe0, 0xf5, 0xb5, 0x46, 0x7e, 0xc6, 0xf6, 0xd5, 0xa2, 0x70, 0xbf, 0x19, 0x60,
	0x57, 0x1e, 0xa2, 0x43, 0x68, 0x4f, 0x99, 0x54, 0x13, 0x12, 0x07, 0x09, 0x09, 0xbf, 0x30, 0x4e,
	0x33, 0xc7, 0xe8, 0x6c, 0x76, 0xed, 0xde, 0xd3, 0xea, 0xcc, 0xf9, 0x28, 0xbc, 0x5d, 0xb6, 0x9f,
	0x96, 0xdd, 0x68, 0x00, 0x48, 0xd2, 0xfc, 0xb5, 0xd1, 0x28, 0x60, 0x3c, 0x53, 0x84, 0x87, 0x34,
	0x73, 0xcc, 0x7f, 0x69, 0xec, 0xdc, 0x13, 0x8e, 0xef, 0xfb, 0xdd, 0x9f, 0x06, 0x34, 0xaa, 0x56,
	0xa3, 0x37, 0xb0, 0x95, 0x12, 0x49, 0x92, 0x72, 0x9b, 0xee, 0x1a, 0x07, 0xce, 0xf3, 0x5e, 0xaa,
	0xa8, 0xc4, 0x9a, 0x86, 0x7c, 0xa8, 0xa7, 0x52, 0xa4, 0x54, 0xaa, 0x59, 0xb9, 0xcc, 0x8b, 0x75,
	0x12, 0x65, 0x3b, 0x9e, 0x13, 0xd1, 0xe1, 0x5f, 0xf1, 0xc9, 0xdf, 0x45, 0xab, 0xd7, 0x59, 0x79,
	0x27, 0x6f, 0x75, 0x58, 0xdc, 0xaf, 0x06, 0x58, 0xf3, 0xdd, 0x10, 0x82, 0x1a, 0xcf, 0x33, 0xa2,
	0x3f, 0x88, 0xe2, 0xbc, 0x1c, 0x0e, 0xf3, 0x61, 0x38, 0x5c, 0xa8, 0x29, 0x3d, 0xdf, 0xec, 0x5a,
	0x07, 0x8f, 0x46, 0x4a, 0x32, 0x7e, 0x83, 0x0b, 0x0c, 0xed, 0x81, 0x95, 0x10, 0x1e, 0x11, 0x25,
	0xe4, 0xac, 0x08, 0x56, 0xfd, 0xa0, 0xa6, 0xe4, 0x84, 0xe2, 0x05, 0xec, 0xbe, 0x82, 0xfa, 0xfd,
	0xdd, 0x56, 0x6e, 0xb0, 0x0b, 0x5b, 0x53, 0x12, 0x4f, 0x68, 0x99, 0x69, 0x5d, 0xec, 0xf9, 0xf0,
	0x64, 0xc5, 0x9f, 0x00, 0xed, 0x40, 0xd3, 0x3f, 0x19, 0x5e, 0x0c, 0x82, 0xd1, 0x11, 0xfe, 0x78,
	0xec, 0x1f, 0xb5, 0x37, 0x50, 0x03, 0xea, 0x97, 0x43, 0xfc, 0xe1, 0xdd, 0xc9, 0xf0, 0xb2, 0x6d,
	0x20, 0x1b, 0x1e, 0xbf, 0xf5, 0xfd, 0xe1, 0xc5, 0xd9, 0xb8, 0x6d, 0xf6, 0x9b, 0x60, 0x6b, 0x33,
	0x8a, 0xaf, 0xb9, 0xbf, 0x0f, 0xcf, 0x42, 0x91, 0x78, 0xd3, 0x44, 0x4d, 0xe4, 0x95, 0xf0, 0xd2,
	0x98, 0xa8, 0x6b, 0x21, 0x93, 0xd2, 0x51, 0x2f, 0x52, 0xa2, 0xdf, 0x98, 0x0f, 0x1d, 0x8c, 0x87,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x94, 0x7a, 0xa7, 0x0e, 0x05, 0x00, 0x00,
}

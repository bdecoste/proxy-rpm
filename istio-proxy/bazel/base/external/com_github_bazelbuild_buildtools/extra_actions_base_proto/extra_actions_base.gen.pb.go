// Code generated by protoc-gen-go.
// source: extra_actions_base_proto/extra_actions_base.proto
// DO NOT EDIT!

/*
Package blaze is a generated protocol buffer package.

It is generated from these files:
	extra_actions_base_proto/extra_actions_base.proto

It has these top-level messages:
	ExtraActionSummary
	DetailedExtraActionInfo
	ExtraActionInfo
	EnvironmentVariable
	SpawnInfo
	CppCompileInfo
	CppLinkInfo
	JavaCompileInfo
	PythonInfo
*/
package blaze

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A list of extra actions and metadata for the print_action command.
type ExtraActionSummary struct {
	Action           []*DetailedExtraActionInfo `protobuf:"bytes,1,rep,name=action" json:"action,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ExtraActionSummary) Reset()                    { *m = ExtraActionSummary{} }
func (m *ExtraActionSummary) String() string            { return proto.CompactTextString(m) }
func (*ExtraActionSummary) ProtoMessage()               {}
func (*ExtraActionSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExtraActionSummary) GetAction() []*DetailedExtraActionInfo {
	if m != nil {
		return m.Action
	}
	return nil
}

// An individual action printed by the print_action command.
type DetailedExtraActionInfo struct {
	// If the given action was included in the output due to a request for a
	// specific file, then this field contains the name of that file so that the
	// caller can correctly associate the extra action with that file.
	//
	// The data in this message is currently not sufficient to run the action on a
	// production machine, because not all necessary input files are identified,
	// especially for C++.
	//
	// There is no easy way to fix this; we could require that all header files
	// are declared and then add all of them here (which would be a huge superset
	// of the files that are actually required), or we could run the include
	// scanner and add those files here.
	TriggeringFile *string `protobuf:"bytes,1,opt,name=triggering_file,json=triggeringFile" json:"triggering_file,omitempty"`
	// The actual action.
	Action           *ExtraActionInfo `protobuf:"bytes,2,req,name=action" json:"action,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DetailedExtraActionInfo) Reset()                    { *m = DetailedExtraActionInfo{} }
func (m *DetailedExtraActionInfo) String() string            { return proto.CompactTextString(m) }
func (*DetailedExtraActionInfo) ProtoMessage()               {}
func (*DetailedExtraActionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DetailedExtraActionInfo) GetTriggeringFile() string {
	if m != nil && m.TriggeringFile != nil {
		return *m.TriggeringFile
	}
	return ""
}

func (m *DetailedExtraActionInfo) GetAction() *ExtraActionInfo {
	if m != nil {
		return m.Action
	}
	return nil
}

// Provides information to an extra_action on the original action it is
// shadowing.
type ExtraActionInfo struct {
	// The label of the ActionOwner of the shadowed action.
	Owner *string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	// Only set if the owner is an Aspect.
	// Corresponds to AspectValue.AspectKey.getAspectClass.getName()
	// This field is deprecated as there might now be
	// multiple aspects applied to the same target.
	// This is the aspect name of the last aspect
	// in 'aspects' (8) field.
	AspectName *string `protobuf:"bytes,6,opt,name=aspect_name,json=aspectName" json:"aspect_name,omitempty"`
	// Only set if the owner is an Aspect.
	// Corresponds to AspectValue.AspectKey.getParameters()
	// This field is deprecated as there might now be
	// multiple aspects applied to the same target.
	// These are the aspect parameters of the last aspect
	// in 'aspects' (8) field.
	AspectParameters map[string]*ExtraActionInfo_StringList `protobuf:"bytes,7,rep,name=aspect_parameters,json=aspectParameters" json:"aspect_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// If the owner is an aspect, all aspects applied to the target
	Aspects []*ExtraActionInfo_AspectDescriptor `protobuf:"bytes,8,rep,name=aspects" json:"aspects,omitempty"`
	// An id uniquely describing the shadowed action at the ActionOwner level.
	Id *string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// The mnemonic of the shadowed action. Used to distinguish actions with the
	// same ActionType.
	Mnemonic                     *string `protobuf:"bytes,5,opt,name=mnemonic" json:"mnemonic,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *ExtraActionInfo) Reset()                    { *m = ExtraActionInfo{} }
func (m *ExtraActionInfo) String() string            { return proto.CompactTextString(m) }
func (*ExtraActionInfo) ProtoMessage()               {}
func (*ExtraActionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

var extRange_ExtraActionInfo = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*ExtraActionInfo) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ExtraActionInfo
}

func (m *ExtraActionInfo) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *ExtraActionInfo) GetAspectName() string {
	if m != nil && m.AspectName != nil {
		return *m.AspectName
	}
	return ""
}

func (m *ExtraActionInfo) GetAspectParameters() map[string]*ExtraActionInfo_StringList {
	if m != nil {
		return m.AspectParameters
	}
	return nil
}

func (m *ExtraActionInfo) GetAspects() []*ExtraActionInfo_AspectDescriptor {
	if m != nil {
		return m.Aspects
	}
	return nil
}

func (m *ExtraActionInfo) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ExtraActionInfo) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

type ExtraActionInfo_StringList struct {
	Value            []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ExtraActionInfo_StringList) Reset()                    { *m = ExtraActionInfo_StringList{} }
func (m *ExtraActionInfo_StringList) String() string            { return proto.CompactTextString(m) }
func (*ExtraActionInfo_StringList) ProtoMessage()               {}
func (*ExtraActionInfo_StringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *ExtraActionInfo_StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type ExtraActionInfo_AspectDescriptor struct {
	// Corresponds to AspectDescriptor.getName()
	AspectName *string `protobuf:"bytes,1,opt,name=aspect_name,json=aspectName" json:"aspect_name,omitempty"`
	// Corresponds to AspectDescriptor.getParameters()
	AspectParameters map[string]*ExtraActionInfo_AspectDescriptor_StringList `protobuf:"bytes,2,rep,name=aspect_parameters,json=aspectParameters" json:"aspect_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte                                                  `json:"-"`
}

func (m *ExtraActionInfo_AspectDescriptor) Reset()         { *m = ExtraActionInfo_AspectDescriptor{} }
func (m *ExtraActionInfo_AspectDescriptor) String() string { return proto.CompactTextString(m) }
func (*ExtraActionInfo_AspectDescriptor) ProtoMessage()    {}
func (*ExtraActionInfo_AspectDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 2}
}

func (m *ExtraActionInfo_AspectDescriptor) GetAspectName() string {
	if m != nil && m.AspectName != nil {
		return *m.AspectName
	}
	return ""
}

func (m *ExtraActionInfo_AspectDescriptor) GetAspectParameters() map[string]*ExtraActionInfo_AspectDescriptor_StringList {
	if m != nil {
		return m.AspectParameters
	}
	return nil
}

type ExtraActionInfo_AspectDescriptor_StringList struct {
	Value            []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ExtraActionInfo_AspectDescriptor_StringList) Reset() {
	*m = ExtraActionInfo_AspectDescriptor_StringList{}
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) String() string {
	return proto.CompactTextString(m)
}
func (*ExtraActionInfo_AspectDescriptor_StringList) ProtoMessage() {}
func (*ExtraActionInfo_AspectDescriptor_StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 2, 1}
}

func (m *ExtraActionInfo_AspectDescriptor_StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type EnvironmentVariable struct {
	// It is possible that this name is not a valid variable identifier.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The value is unescaped and unquoted.
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EnvironmentVariable) Reset()                    { *m = EnvironmentVariable{} }
func (m *EnvironmentVariable) String() string            { return proto.CompactTextString(m) }
func (*EnvironmentVariable) ProtoMessage()               {}
func (*EnvironmentVariable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EnvironmentVariable) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *EnvironmentVariable) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

// Provides access to data that is specific to spawn actions.
// Usually provided by actions using the "Spawn" & "Genrule" Mnemonics.
type SpawnInfo struct {
	Argument []string `protobuf:"bytes,1,rep,name=argument" json:"argument,omitempty"`
	// An *unordered* list of environment variables and their values.
	Variable         []*EnvironmentVariable `protobuf:"bytes,2,rep,name=variable" json:"variable,omitempty"`
	InputFile        []string               `protobuf:"bytes,4,rep,name=input_file,json=inputFile" json:"input_file,omitempty"`
	OutputFile       []string               `protobuf:"bytes,5,rep,name=output_file,json=outputFile" json:"output_file,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SpawnInfo) Reset()                    { *m = SpawnInfo{} }
func (m *SpawnInfo) String() string            { return proto.CompactTextString(m) }
func (*SpawnInfo) ProtoMessage()               {}
func (*SpawnInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SpawnInfo) GetArgument() []string {
	if m != nil {
		return m.Argument
	}
	return nil
}

func (m *SpawnInfo) GetVariable() []*EnvironmentVariable {
	if m != nil {
		return m.Variable
	}
	return nil
}

func (m *SpawnInfo) GetInputFile() []string {
	if m != nil {
		return m.InputFile
	}
	return nil
}

func (m *SpawnInfo) GetOutputFile() []string {
	if m != nil {
		return m.OutputFile
	}
	return nil
}

var E_SpawnInfo_SpawnInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*SpawnInfo)(nil),
	Field:         1003,
	Name:          "blaze.SpawnInfo.spawn_info",
	Tag:           "bytes,1003,opt,name=spawn_info,json=spawnInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

// Provides access to data that is specific to C++ compile actions.
// Usually provided by actions using the "CppCompile" Mnemonic.
type CppCompileInfo struct {
	Tool           *string  `protobuf:"bytes,1,opt,name=tool" json:"tool,omitempty"`
	CompilerOption []string `protobuf:"bytes,2,rep,name=compiler_option,json=compilerOption" json:"compiler_option,omitempty"`
	SourceFile     *string  `protobuf:"bytes,3,opt,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	OutputFile     *string  `protobuf:"bytes,4,opt,name=output_file,json=outputFile" json:"output_file,omitempty"`
	// Due to header discovery, this won't include headers unless the build is
	// actually performed. If set, this field will include the value of
	// "source_file" in addition to the headers.
	SourcesAndHeaders []string `protobuf:"bytes,5,rep,name=sources_and_headers,json=sourcesAndHeaders" json:"sources_and_headers,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *CppCompileInfo) Reset()                    { *m = CppCompileInfo{} }
func (m *CppCompileInfo) String() string            { return proto.CompactTextString(m) }
func (*CppCompileInfo) ProtoMessage()               {}
func (*CppCompileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CppCompileInfo) GetTool() string {
	if m != nil && m.Tool != nil {
		return *m.Tool
	}
	return ""
}

func (m *CppCompileInfo) GetCompilerOption() []string {
	if m != nil {
		return m.CompilerOption
	}
	return nil
}

func (m *CppCompileInfo) GetSourceFile() string {
	if m != nil && m.SourceFile != nil {
		return *m.SourceFile
	}
	return ""
}

func (m *CppCompileInfo) GetOutputFile() string {
	if m != nil && m.OutputFile != nil {
		return *m.OutputFile
	}
	return ""
}

func (m *CppCompileInfo) GetSourcesAndHeaders() []string {
	if m != nil {
		return m.SourcesAndHeaders
	}
	return nil
}

var E_CppCompileInfo_CppCompileInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*CppCompileInfo)(nil),
	Field:         1001,
	Name:          "blaze.CppCompileInfo.cpp_compile_info",
	Tag:           "bytes,1001,opt,name=cpp_compile_info,json=cppCompileInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

// Provides access to data that is specific to C++ link  actions.
// Usually provided by actions using the "CppLink" Mnemonic.
type CppLinkInfo struct {
	InputFile               []string `protobuf:"bytes,1,rep,name=input_file,json=inputFile" json:"input_file,omitempty"`
	OutputFile              *string  `protobuf:"bytes,2,opt,name=output_file,json=outputFile" json:"output_file,omitempty"`
	InterfaceOutputFile     *string  `protobuf:"bytes,3,opt,name=interface_output_file,json=interfaceOutputFile" json:"interface_output_file,omitempty"`
	LinkTargetType          *string  `protobuf:"bytes,4,opt,name=link_target_type,json=linkTargetType" json:"link_target_type,omitempty"`
	LinkStaticness          *string  `protobuf:"bytes,5,opt,name=link_staticness,json=linkStaticness" json:"link_staticness,omitempty"`
	LinkStamp               []string `protobuf:"bytes,6,rep,name=link_stamp,json=linkStamp" json:"link_stamp,omitempty"`
	BuildInfoHeaderArtifact []string `protobuf:"bytes,7,rep,name=build_info_header_artifact,json=buildInfoHeaderArtifact" json:"build_info_header_artifact,omitempty"`
	// The list of command line options used for running the linking tool.
	LinkOpt          []string `protobuf:"bytes,8,rep,name=link_opt,json=linkOpt" json:"link_opt,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CppLinkInfo) Reset()                    { *m = CppLinkInfo{} }
func (m *CppLinkInfo) String() string            { return proto.CompactTextString(m) }
func (*CppLinkInfo) ProtoMessage()               {}
func (*CppLinkInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CppLinkInfo) GetInputFile() []string {
	if m != nil {
		return m.InputFile
	}
	return nil
}

func (m *CppLinkInfo) GetOutputFile() string {
	if m != nil && m.OutputFile != nil {
		return *m.OutputFile
	}
	return ""
}

func (m *CppLinkInfo) GetInterfaceOutputFile() string {
	if m != nil && m.InterfaceOutputFile != nil {
		return *m.InterfaceOutputFile
	}
	return ""
}

func (m *CppLinkInfo) GetLinkTargetType() string {
	if m != nil && m.LinkTargetType != nil {
		return *m.LinkTargetType
	}
	return ""
}

func (m *CppLinkInfo) GetLinkStaticness() string {
	if m != nil && m.LinkStaticness != nil {
		return *m.LinkStaticness
	}
	return ""
}

func (m *CppLinkInfo) GetLinkStamp() []string {
	if m != nil {
		return m.LinkStamp
	}
	return nil
}

func (m *CppLinkInfo) GetBuildInfoHeaderArtifact() []string {
	if m != nil {
		return m.BuildInfoHeaderArtifact
	}
	return nil
}

func (m *CppLinkInfo) GetLinkOpt() []string {
	if m != nil {
		return m.LinkOpt
	}
	return nil
}

var E_CppLinkInfo_CppLinkInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*CppLinkInfo)(nil),
	Field:         1002,
	Name:          "blaze.CppLinkInfo.cpp_link_info",
	Tag:           "bytes,1002,opt,name=cpp_link_info,json=cppLinkInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

// Provides access to data that is specific to java compile actions.
// Usually provided by actions using the "Javac" Mnemonic.
type JavaCompileInfo struct {
	Outputjar        *string  `protobuf:"bytes,1,opt,name=outputjar" json:"outputjar,omitempty"`
	Classpath        []string `protobuf:"bytes,2,rep,name=classpath" json:"classpath,omitempty"`
	Sourcepath       []string `protobuf:"bytes,3,rep,name=sourcepath" json:"sourcepath,omitempty"`
	SourceFile       []string `protobuf:"bytes,4,rep,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	JavacOpt         []string `protobuf:"bytes,5,rep,name=javac_opt,json=javacOpt" json:"javac_opt,omitempty"`
	Processor        []string `protobuf:"bytes,6,rep,name=processor" json:"processor,omitempty"`
	Processorpath    []string `protobuf:"bytes,7,rep,name=processorpath" json:"processorpath,omitempty"`
	Bootclasspath    []string `protobuf:"bytes,8,rep,name=bootclasspath" json:"bootclasspath,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *JavaCompileInfo) Reset()                    { *m = JavaCompileInfo{} }
func (m *JavaCompileInfo) String() string            { return proto.CompactTextString(m) }
func (*JavaCompileInfo) ProtoMessage()               {}
func (*JavaCompileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *JavaCompileInfo) GetOutputjar() string {
	if m != nil && m.Outputjar != nil {
		return *m.Outputjar
	}
	return ""
}

func (m *JavaCompileInfo) GetClasspath() []string {
	if m != nil {
		return m.Classpath
	}
	return nil
}

func (m *JavaCompileInfo) GetSourcepath() []string {
	if m != nil {
		return m.Sourcepath
	}
	return nil
}

func (m *JavaCompileInfo) GetSourceFile() []string {
	if m != nil {
		return m.SourceFile
	}
	return nil
}

func (m *JavaCompileInfo) GetJavacOpt() []string {
	if m != nil {
		return m.JavacOpt
	}
	return nil
}

func (m *JavaCompileInfo) GetProcessor() []string {
	if m != nil {
		return m.Processor
	}
	return nil
}

func (m *JavaCompileInfo) GetProcessorpath() []string {
	if m != nil {
		return m.Processorpath
	}
	return nil
}

func (m *JavaCompileInfo) GetBootclasspath() []string {
	if m != nil {
		return m.Bootclasspath
	}
	return nil
}

var E_JavaCompileInfo_JavaCompileInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*JavaCompileInfo)(nil),
	Field:         1000,
	Name:          "blaze.JavaCompileInfo.java_compile_info",
	Tag:           "bytes,1000,opt,name=java_compile_info,json=javaCompileInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

// Provides access to data that is specific to python rules.
// Usually provided by actions using the "Python" Mnemonic.
type PythonInfo struct {
	SourceFile       []string `protobuf:"bytes,1,rep,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	DepFile          []string `protobuf:"bytes,2,rep,name=dep_file,json=depFile" json:"dep_file,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PythonInfo) Reset()                    { *m = PythonInfo{} }
func (m *PythonInfo) String() string            { return proto.CompactTextString(m) }
func (*PythonInfo) ProtoMessage()               {}
func (*PythonInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PythonInfo) GetSourceFile() []string {
	if m != nil {
		return m.SourceFile
	}
	return nil
}

func (m *PythonInfo) GetDepFile() []string {
	if m != nil {
		return m.DepFile
	}
	return nil
}

var E_PythonInfo_PythonInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*PythonInfo)(nil),
	Field:         1005,
	Name:          "blaze.PythonInfo.python_info",
	Tag:           "bytes,1005,opt,name=python_info,json=pythonInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

func init() {
	proto.RegisterType((*ExtraActionSummary)(nil), "blaze.ExtraActionSummary")
	proto.RegisterType((*DetailedExtraActionInfo)(nil), "blaze.DetailedExtraActionInfo")
	proto.RegisterType((*ExtraActionInfo)(nil), "blaze.ExtraActionInfo")
	proto.RegisterType((*ExtraActionInfo_StringList)(nil), "blaze.ExtraActionInfo.StringList")
	proto.RegisterType((*ExtraActionInfo_AspectDescriptor)(nil), "blaze.ExtraActionInfo.AspectDescriptor")
	proto.RegisterType((*ExtraActionInfo_AspectDescriptor_StringList)(nil), "blaze.ExtraActionInfo.AspectDescriptor.StringList")
	proto.RegisterType((*EnvironmentVariable)(nil), "blaze.EnvironmentVariable")
	proto.RegisterType((*SpawnInfo)(nil), "blaze.SpawnInfo")
	proto.RegisterType((*CppCompileInfo)(nil), "blaze.CppCompileInfo")
	proto.RegisterType((*CppLinkInfo)(nil), "blaze.CppLinkInfo")
	proto.RegisterType((*JavaCompileInfo)(nil), "blaze.JavaCompileInfo")
	proto.RegisterType((*PythonInfo)(nil), "blaze.PythonInfo")
	proto.RegisterExtension(E_SpawnInfo_SpawnInfo)
	proto.RegisterExtension(E_CppCompileInfo_CppCompileInfo)
	proto.RegisterExtension(E_CppLinkInfo_CppLinkInfo)
	proto.RegisterExtension(E_JavaCompileInfo_JavaCompileInfo)
	proto.RegisterExtension(E_PythonInfo_PythonInfo)
}

func init() { proto.RegisterFile("extra_actions_base_proto/extra_actions_base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1030 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0xdd, 0xa6, 0x8d, 0x4f, 0xb4, 0x69, 0x3a, 0xa5, 0xac, 0x37, 0xc0, 0x6e, 0x09, 0x88,
	0xad, 0x00, 0x79, 0x45, 0x2e, 0x16, 0x04, 0x42, 0xa8, 0xdb, 0x2d, 0x2a, 0x50, 0xd1, 0xca, 0x5d,
	0x21, 0x24, 0x84, 0xac, 0x89, 0x3d, 0x4d, 0xa7, 0xb5, 0x3d, 0xa3, 0xf1, 0x24, 0x25, 0x5c, 0xf5,
	0x25, 0xb8, 0x40, 0xe2, 0x51, 0x78, 0xa0, 0x5d, 0x10, 0x37, 0x88, 0x07, 0x40, 0xf3, 0x63, 0x3b,
	0x7f, 0xed, 0xf6, 0x6e, 0xe6, 0xfb, 0x3e, 0x9f, 0x39, 0xe7, 0x3b, 0x67, 0x26, 0x81, 0x4f, 0xc8,
	0x2f, 0x52, 0xe0, 0x08, 0xc7, 0x92, 0xb2, 0xbc, 0x88, 0x06, 0xb8, 0x20, 0x11, 0x17, 0x4c, 0xb2,
	0x27, 0x8b, 0x44, 0xa0, 0x09, 0xd4, 0x18, 0xa4, 0xf8, 0x57, 0xd2, 0x3b, 0x02, 0x74, 0xa0, 0x24,
	0x7b, 0x5a, 0x71, 0x3a, 0xca, 0x32, 0x2c, 0x26, 0xe8, 0x29, 0xac, 0x99, 0x4f, 0x7c, 0x67, 0x67,
	0x65, 0xb7, 0xd5, 0x7f, 0x18, 0x68, 0x75, 0xf0, 0x9c, 0x48, 0x4c, 0x53, 0x92, 0x4c, 0x7d, 0xf2,
	0x4d, 0x7e, 0xc6, 0x42, 0xab, 0xee, 0x09, 0xb8, 0x7f, 0x83, 0x04, 0x3d, 0x86, 0x0d, 0x29, 0xe8,
	0x70, 0x48, 0x04, 0xcd, 0x87, 0xd1, 0x19, 0x4d, 0x89, 0xef, 0xec, 0x38, 0xbb, 0x5e, 0xd8, 0xae,
	0xe1, 0xaf, 0x69, 0x4a, 0x50, 0x50, 0x9d, 0xed, 0xee, 0xb8, 0xbb, 0xad, 0xfe, 0x9b, 0xf6, 0xec,
	0x9b, 0xce, 0xfc, 0xaf, 0x01, 0x1b, 0xf3, 0x87, 0xbd, 0x01, 0x0d, 0x76, 0x95, 0x13, 0x61, 0x8f,
	0x30, 0x1b, 0xf4, 0x1e, 0xb4, 0x70, 0xc1, 0x49, 0x2c, 0xa3, 0x1c, 0x67, 0xc4, 0x5f, 0x53, 0xdc,
	0x33, 0xd7, 0x77, 0x42, 0x30, 0xf0, 0xf7, 0x38, 0x23, 0xe8, 0x67, 0xd8, 0xb4, 0x22, 0x8e, 0x05,
	0xce, 0x88, 0x24, 0xa2, 0xf0, 0xd7, 0xb5, 0x0b, 0x1f, 0x2f, 0xcf, 0x24, 0xd8, 0xd3, 0xfa, 0x93,
	0x4a, 0x7e, 0x90, 0x4b, 0x31, 0xd1, 0x81, 0x3b, 0x78, 0x8e, 0x42, 0x7b, 0xb0, 0x6e, 0xb0, 0xc2,
	0x6f, 0xea, 0xa0, 0x8f, 0x6f, 0x0d, 0xfa, 0x9c, 0x14, 0xb1, 0xa0, 0x5c, 0x32, 0x11, 0x96, 0xdf,
	0xa1, 0x36, 0xb8, 0x34, 0xf1, 0x5d, 0x5d, 0x99, 0x4b, 0x13, 0xd4, 0x85, 0x66, 0x96, 0x93, 0x8c,
	0xe5, 0x34, 0xf6, 0x1b, 0x1a, 0xad, 0xf6, 0xdd, 0x33, 0xd8, 0x5e, 0x9a, 0x1d, 0xea, 0xc0, 0xca,
	0x25, 0x99, 0x58, 0x7f, 0xd4, 0x12, 0x7d, 0x0a, 0x8d, 0x31, 0x4e, 0x47, 0x44, 0x47, 0x6e, 0xf5,
	0xdf, 0xbd, 0x21, 0xaf, 0x53, 0xa9, 0x3a, 0x75, 0x44, 0x0b, 0x19, 0x1a, 0xfd, 0xe7, 0xee, 0x67,
	0x4e, 0xf7, 0x03, 0x80, 0x9a, 0x50, 0xf6, 0x9b, 0x50, 0x6a, 0x7a, 0xbc, 0x4a, 0xe7, 0x3b, 0xdd,
	0x3f, 0x5d, 0xe8, 0xcc, 0x57, 0x86, 0x1e, 0xcd, 0xf6, 0xc5, 0xe4, 0x34, 0xdd, 0x93, 0x8b, 0x65,
	0x3d, 0x71, 0xb5, 0x7d, 0x5f, 0xde, 0xd1, 0xbe, 0xe5, 0x4d, 0x5a, 0x6c, 0x50, 0xf7, 0xea, 0xee,
	0x8e, 0x1d, 0xce, 0x3a, 0xd6, 0xbf, 0x6b, 0x2a, 0xcb, 0x2d, 0xec, 0xbd, 0xde, 0xc2, 0x0f, 0xbd,
	0xe6, 0xcb, 0xf5, 0xce, 0xf5, 0xf5, 0xf5, 0xb5, 0xdb, 0xfb, 0x0a, 0xb6, 0x0e, 0xf2, 0x31, 0x15,
	0x2c, 0xcf, 0x48, 0x2e, 0x7f, 0xc0, 0x82, 0xe2, 0x41, 0x4a, 0x10, 0x82, 0x55, 0x6b, 0xa2, 0xbb,
	0xeb, 0x85, 0x7a, 0x5d, 0xc7, 0x72, 0x35, 0x68, 0x36, 0xbd, 0x57, 0x0e, 0x78, 0xa7, 0x1c, 0x5f,
	0x99, 0x1b, 0xd3, 0x85, 0x26, 0x16, 0xc3, 0x91, 0x8a, 0x65, 0x8f, 0xac, 0xf6, 0xe8, 0x29, 0x34,
	0xc7, 0x36, 0xbe, 0x75, 0xbd, 0x5b, 0x96, 0xba, 0x98, 0x41, 0x58, 0x69, 0xd1, 0x3b, 0x00, 0x34,
	0xe7, 0x23, 0x69, 0x6e, 0xfb, 0xaa, 0x8e, 0xea, 0x69, 0x44, 0x5f, 0xf4, 0x47, 0xd0, 0x62, 0x23,
	0x59, 0xf1, 0x0d, 0xcd, 0x83, 0x81, 0x94, 0xa0, 0x7f, 0x08, 0x50, 0xa8, 0x04, 0x23, 0xaa, 0x32,
	0xbc, 0xe1, 0x1d, 0xf0, 0xff, 0x5e, 0xd7, 0xee, 0x77, 0x2c, 0x5d, 0x95, 0x14, 0x7a, 0x45, 0xb9,
	0xec, 0xfd, 0xee, 0x42, 0x7b, 0x9f, 0xf3, 0x7d, 0x96, 0x71, 0x9a, 0x12, 0x5d, 0x30, 0x82, 0x55,
	0xc9, 0x58, 0x6a, 0xfb, 0xa9, 0xd7, 0xea, 0x8d, 0x8a, 0x8d, 0x44, 0x44, 0x8c, 0xdb, 0x37, 0x48,
	0x65, 0xd5, 0x2e, 0xe1, 0x63, 0x8d, 0xaa, 0xd4, 0x0b, 0x36, 0x12, 0x31, 0x31, 0xa9, 0xaf, 0x98,
	0x89, 0x35, 0xd0, 0xb2, 0xda, 0x56, 0x8d, 0xa0, 0xae, 0x0d, 0x05, 0xb0, 0x65, 0xe4, 0x45, 0x84,
	0xf3, 0x24, 0x3a, 0x27, 0x38, 0x51, 0x43, 0x6d, 0x4c, 0xd8, 0xb4, 0xd4, 0x5e, 0x9e, 0x1c, 0x1a,
	0xa2, 0xff, 0x23, 0x74, 0x62, 0xce, 0x23, 0x9b, 0xc7, 0xed, 0x8e, 0xbc, 0x32, 0x8e, 0x6c, 0x5b,
	0x7a, 0xb6, 0xf0, 0xb0, 0x1d, 0xcf, 0xec, 0x7b, 0x7f, 0xac, 0x40, 0x6b, 0x9f, 0xf3, 0x23, 0x9a,
	0x5f, 0x6a, 0x63, 0x66, 0xbb, 0xe6, 0xbc, 0xa6, 0x6b, 0xee, 0x42, 0x65, 0x7d, 0xd8, 0xa6, 0xb9,
	0x24, 0xe2, 0x0c, 0xc7, 0x24, 0x9a, 0x96, 0x1a, 0x97, 0xb6, 0x2a, 0xf2, 0xb8, 0xfe, 0x66, 0x17,
	0x3a, 0x29, 0xcd, 0x2f, 0x23, 0x89, 0xc5, 0x90, 0xc8, 0x48, 0x4e, 0x78, 0xe9, 0x59, 0x5b, 0xe1,
	0x2f, 0x34, 0xfc, 0x62, 0xc2, 0x89, 0x6a, 0x91, 0x56, 0x16, 0x12, 0x4b, 0x1a, 0xe7, 0xa4, 0x28,
	0xec, 0x9b, 0xa7, 0x85, 0xa7, 0x15, 0xaa, 0xca, 0x28, 0x85, 0x19, 0xf7, 0xd7, 0x4c, 0x19, 0x56,
	0x93, 0x71, 0xf4, 0x05, 0x74, 0x07, 0x23, 0x9a, 0x26, 0xda, 0x49, 0x6b, 0x7f, 0x84, 0x85, 0xa4,
	0x67, 0x38, 0x96, 0xfa, 0xbd, 0xf7, 0xc2, 0xfb, 0x5a, 0xa1, 0x4c, 0x31, 0x5d, 0xd8, 0xb3, 0x34,
	0x7a, 0x00, 0x4d, 0x1d, 0x9b, 0x71, 0xa9, 0x5f, 0x71, 0x2f, 0x5c, 0x57, 0xfb, 0x63, 0x2e, 0xfb,
	0xc7, 0x70, 0x4f, 0xf5, 0x49, 0xd3, 0xb7, 0x36, 0xe9, 0x2f, 0xd3, 0x24, 0x54, 0x37, 0xa9, 0xec,
	0x40, 0xd8, 0x8a, 0xeb, 0x4d, 0xef, 0x5f, 0x17, 0x36, 0xbe, 0xc5, 0x63, 0x3c, 0x3d, 0xbb, 0x6f,
	0x83, 0x67, 0x8c, 0xbd, 0xc0, 0xe5, 0x4f, 0x5c, 0x0d, 0x28, 0x36, 0x4e, 0x71, 0x51, 0x70, 0x2c,
	0xcf, 0xed, 0xfc, 0xd6, 0x00, 0x7a, 0x08, 0x76, 0x4e, 0x35, 0xbd, 0x62, 0x2e, 0x5d, 0x8d, 0xcc,
	0x8f, 0xf6, 0xea, 0xb4, 0x40, 0xf7, 0xea, 0x2d, 0xf0, 0x2e, 0xf0, 0x18, 0xc7, 0xba, 0x7a, 0x33,
	0xaf, 0x4d, 0x0d, 0x1c, 0x73, 0xa9, 0xce, 0xe6, 0x82, 0xc5, 0xa4, 0x28, 0x98, 0x28, 0x4d, 0xaf,
	0x00, 0xf4, 0x3e, 0xdc, 0xab, 0x36, 0xfa, 0x78, 0xe3, 0xf3, 0x2c, 0xa8, 0x54, 0x03, 0xc6, 0x64,
	0x5d, 0x83, 0xb1, 0x78, 0x16, 0xec, 0xff, 0x04, 0x9b, 0xea, 0xd4, 0xbb, 0xdd, 0x88, 0x97, 0xc6,
	0xec, 0x92, 0x9e, 0xf3, 0x33, 0xdc, 0xb8, 0x98, 0x05, 0x7a, 0xbf, 0x39, 0x00, 0x27, 0x13, 0x79,
	0x6e, 0xff, 0x4e, 0xcc, 0x79, 0xe2, 0x2c, 0x78, 0xf2, 0x00, 0x9a, 0x09, 0xe1, 0xe5, 0x8d, 0xd0,
	0x03, 0x91, 0x10, 0xae, 0x1f, 0xb1, 0xef, 0xa0, 0xc5, 0x75, 0xa4, 0xdb, 0x33, 0xfc, 0xc7, 0x64,
	0xb8, 0x69, 0xe9, 0xfa, 0xf0, 0x10, 0x78, 0xb5, 0x7e, 0xf6, 0x04, 0x3e, 0x8a, 0x59, 0x16, 0x0c,
	0x19, 0x1b, 0xa6, 0x24, 0x48, 0xc8, 0x58, 0x3d, 0x5b, 0x45, 0xa0, 0xe7, 0x34, 0x48, 0xe9, 0x20,
	0xb0, 0x7f, 0xf4, 0x02, 0xfd, 0xb7, 0xef, 0xc4, 0xf9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x62, 0x64,
	0x2d, 0xeb, 0x19, 0x0a, 0x00, 0x00,
}

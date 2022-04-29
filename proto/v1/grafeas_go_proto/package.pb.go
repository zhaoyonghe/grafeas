// Copyright 2019 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: proto/v1/package.proto

package grafeas_go_proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Instruction set architectures supported by various package managers.
type Architecture int32

const (
	// Unknown architecture.
	Architecture_ARCHITECTURE_UNSPECIFIED Architecture = 0
	// X86 architecture.
	Architecture_X86 Architecture = 1
	// X64 architecture.
	Architecture_X64 Architecture = 2
)

// Enum value maps for Architecture.
var (
	Architecture_name = map[int32]string{
		0: "ARCHITECTURE_UNSPECIFIED",
		1: "X86",
		2: "X64",
	}
	Architecture_value = map[string]int32{
		"ARCHITECTURE_UNSPECIFIED": 0,
		"X86":                      1,
		"X64":                      2,
	}
)

func (x Architecture) Enum() *Architecture {
	p := new(Architecture)
	*p = x
	return p
}

func (x Architecture) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Architecture) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1_package_proto_enumTypes[0].Descriptor()
}

func (Architecture) Type() protoreflect.EnumType {
	return &file_proto_v1_package_proto_enumTypes[0]
}

func (x Architecture) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Architecture.Descriptor instead.
func (Architecture) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1_package_proto_rawDescGZIP(), []int{0}
}

// Whether this is an ordinary package version or a sentinel MIN/MAX version.
type Version_VersionKind int32

const (
	// Unknown.
	Version_VERSION_KIND_UNSPECIFIED Version_VersionKind = 0
	// A standard package version.
	Version_NORMAL Version_VersionKind = 1
	// A special version representing negative infinity.
	Version_MINIMUM Version_VersionKind = 2
	// A special version representing positive infinity.
	Version_MAXIMUM Version_VersionKind = 3
)

// Enum value maps for Version_VersionKind.
var (
	Version_VersionKind_name = map[int32]string{
		0: "VERSION_KIND_UNSPECIFIED",
		1: "NORMAL",
		2: "MINIMUM",
		3: "MAXIMUM",
	}
	Version_VersionKind_value = map[string]int32{
		"VERSION_KIND_UNSPECIFIED": 0,
		"NORMAL":                   1,
		"MINIMUM":                  2,
		"MAXIMUM":                  3,
	}
)

func (x Version_VersionKind) Enum() *Version_VersionKind {
	p := new(Version_VersionKind)
	*p = x
	return p
}

func (x Version_VersionKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version_VersionKind) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1_package_proto_enumTypes[1].Descriptor()
}

func (Version_VersionKind) Type() protoreflect.EnumType {
	return &file_proto_v1_package_proto_enumTypes[1]
}

func (x Version_VersionKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version_VersionKind.Descriptor instead.
func (Version_VersionKind) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1_package_proto_rawDescGZIP(), []int{4, 0}
}

// This represents a particular channel of distribution for a given package.
// E.g., Debian's jessie-backports dpkg mirror.
type Distribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cpe_uri in [CPE format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The CPU architecture for which packages in this distribution channel were
	// built.
	Architecture Architecture `protobuf:"varint,2,opt,name=architecture,proto3,enum=grafeas.v1.Architecture" json:"architecture,omitempty"`
	// The latest available version of this package in this distribution channel.
	LatestVersion *Version `protobuf:"bytes,3,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	// A freeform string denoting the maintainer of this package.
	Maintainer string `protobuf:"bytes,4,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	// The distribution channel-specific homepage for this package.
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// The distribution channel-specific description of this package.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Distribution) Reset() {
	*x = Distribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution) ProtoMessage() {}

func (x *Distribution) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution.ProtoReflect.Descriptor instead.
func (*Distribution) Descriptor() ([]byte, []int) {
	return file_proto_v1_package_proto_rawDescGZIP(), []int{0}
}

func (x *Distribution) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *Distribution) GetArchitecture() Architecture {
	if x != nil {
		return x.Architecture
	}
	return Architecture_ARCHITECTURE_UNSPECIFIED
}

func (x *Distribution) GetLatestVersion() *Version {
	if x != nil {
		return x.LatestVersion
	}
	return nil
}

func (x *Distribution) GetMaintainer() string {
	if x != nil {
		return x.Maintainer
	}
	return ""
}

func (x *Distribution) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Distribution) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// An occurrence of a particular package installation found within a system's
// filesystem. E.g., glibc was found in `/var/lib/dpkg/status`.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated.
	// The CPE URI in [CPE format](https://cpe.mitre.org/specification/)
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// Deprecated.
	// The version installed at this location.
	Version *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// The path from which we gathered that this package/version is installed.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_proto_v1_package_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *Location) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Location) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// PackageNote represents a particular package version.
type PackageNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the package.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated.
	// The various channels by which a package is distributed.
	Distribution []*Distribution `protobuf:"bytes,10,rep,name=distribution,proto3" json:"distribution,omitempty"`
	// The type of package; whether native or non native (e.g., ruby gems,
	// node.js packages, etc.).
	PackageType string `protobuf:"bytes,11,opt,name=package_type,json=packageType,proto3" json:"package_type,omitempty"`
	// The cpe_uri in [CPE format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	// The cpe_uri will be blank for language packages.
	CpeUri string `protobuf:"bytes,12,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The CPU architecture for which packages in this distribution channel were
	// built. Architecture will be blank for language packages.
	Architecture Architecture `protobuf:"varint,13,opt,name=architecture,proto3,enum=grafeas.v1.Architecture" json:"architecture,omitempty"`
	// The version of the package.
	Version *Version `protobuf:"bytes,14,opt,name=version,proto3" json:"version,omitempty"`
	// A freeform text denoting the maintainer of this package.
	Maintainer string `protobuf:"bytes,15,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	// The homepage for this package.
	Url string `protobuf:"bytes,16,opt,name=url,proto3" json:"url,omitempty"`
	// The description of this package.
	Description string `protobuf:"bytes,17,opt,name=description,proto3" json:"description,omitempty"`
	// Licenses that have been declared by the authors of the package.
	License *License `protobuf:"bytes,18,opt,name=license,proto3" json:"license,omitempty"`
	// Hash value, typically a file digest, that allows unique
	// identification a specific package.
	Digest []*Digest `protobuf:"bytes,19,rep,name=digest,proto3" json:"digest,omitempty"`
}

func (x *PackageNote) Reset() {
	*x = PackageNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageNote) ProtoMessage() {}

func (x *PackageNote) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageNote.ProtoReflect.Descriptor instead.
func (*PackageNote) Descriptor() ([]byte, []int) {
	return file_proto_v1_package_proto_rawDescGZIP(), []int{2}
}

func (x *PackageNote) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PackageNote) GetDistribution() []*Distribution {
	if x != nil {
		return x.Distribution
	}
	return nil
}

func (x *PackageNote) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

func (x *PackageNote) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *PackageNote) GetArchitecture() Architecture {
	if x != nil {
		return x.Architecture
	}
	return Architecture_ARCHITECTURE_UNSPECIFIED
}

func (x *PackageNote) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *PackageNote) GetMaintainer() string {
	if x != nil {
		return x.Maintainer
	}
	return ""
}

func (x *PackageNote) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PackageNote) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PackageNote) GetLicense() *License {
	if x != nil {
		return x.License
	}
	return nil
}

func (x *PackageNote) GetDigest() []*Digest {
	if x != nil {
		return x.Digest
	}
	return nil
}

// Details on how a particular software package was installed on a system.
type PackageOccurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the installed package.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// All of the places within the filesystem versions of this package
	// have been found.
	Location []*Location `protobuf:"bytes,2,rep,name=location,proto3" json:"location,omitempty"`
	// The type of package; whether native or non native (e.g., ruby gems,
	// node.js packages, etc.).
	PackageType string `protobuf:"bytes,3,opt,name=package_type,json=packageType,proto3" json:"package_type,omitempty"`
	// The cpe_uri in [CPE format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	// The cpe_uri will be blank for language packages.
	CpeUri string `protobuf:"bytes,4,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The CPU architecture for which packages in this distribution channel were
	// built. Architecture will be blank for language packages.
	Architecture Architecture `protobuf:"varint,5,opt,name=architecture,proto3,enum=grafeas.v1.Architecture" json:"architecture,omitempty"`
	// Licenses that have been declared by the authors of the package.
	License *License `protobuf:"bytes,6,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *PackageOccurrence) Reset() {
	*x = PackageOccurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageOccurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageOccurrence) ProtoMessage() {}

func (x *PackageOccurrence) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageOccurrence.ProtoReflect.Descriptor instead.
func (*PackageOccurrence) Descriptor() ([]byte, []int) {
	return file_proto_v1_package_proto_rawDescGZIP(), []int{3}
}

func (x *PackageOccurrence) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PackageOccurrence) GetLocation() []*Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *PackageOccurrence) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

func (x *PackageOccurrence) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *PackageOccurrence) GetArchitecture() Architecture {
	if x != nil {
		return x.Architecture
	}
	return Architecture_ARCHITECTURE_UNSPECIFIED
}

func (x *PackageOccurrence) GetLicense() *License {
	if x != nil {
		return x.License
	}
	return nil
}

// Version contains structured information about the version of a package.
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to correct mistakes in the version numbering scheme.
	Epoch int32 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Required only when version kind is NORMAL. The main part of the version
	// name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The iteration of the package build from the above version.
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// Whether this version is specifying part of an inclusive range. Grafeas
	// does not have the capability to specify version ranges; instead we have
	// fields that specify start version and end versions. At times this is
	// insufficient - we also need to specify whether the version is included in
	// the range or is excluded from the range. This boolean is expected to be set
	// to true when the version is included in a range.
	Inclusive bool `protobuf:"varint,6,opt,name=inclusive,proto3" json:"inclusive,omitempty"`
	// Required. Distinguishes between sentinel MIN/MAX versions and normal
	// versions.
	Kind Version_VersionKind `protobuf:"varint,4,opt,name=kind,proto3,enum=grafeas.v1.Version_VersionKind" json:"kind,omitempty"`
	// Human readable version string. This string is of the form
	// <epoch>:<name>-<revision> and is only set when kind is NORMAL.
	FullName string `protobuf:"bytes,5,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_proto_v1_package_proto_rawDescGZIP(), []int{4}
}

func (x *Version) GetEpoch() int32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Version) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *Version) GetInclusive() bool {
	if x != nil {
		return x.Inclusive
	}
	return false
}

func (x *Version) GetKind() Version_VersionKind {
	if x != nil {
		return x.Kind
	}
	return Version_VERSION_KIND_UNSPECIFIED
}

func (x *Version) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

var File_proto_v1_package_proto protoreflect.FileDescriptor

var file_proto_v1_package_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a,
	0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x07, 0x63, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x06, 0x63, 0x70, 0x65, 0x55, 0x72, 0x69, 0x12, 0x3c, 0x0a, 0x0c, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0c, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x70, 0x65, 0x55, 0x72, 0x69, 0x12, 0x2d,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0xbf, 0x03, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x70, 0x65, 0x55, 0x72, 0x69, 0x12, 0x3c, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x07,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x22, 0x96, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4f,
	0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x07,
	0x63, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x06, 0x63, 0x70, 0x65, 0x55, 0x72, 0x69, 0x12, 0x41, 0x0a, 0x0c, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2d, 0x0a,
	0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x92, 0x02, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x33, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x51,
	0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1c, 0x0a,
	0x18, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x4e, 0x49, 0x4d,
	0x55, 0x4d, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x55, 0x4d, 0x10,
	0x03, 0x2a, 0x3e, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x52, 0x43, 0x48, 0x49, 0x54, 0x45, 0x43, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x58, 0x38, 0x36, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x36, 0x34, 0x10,
	0x02, 0x42, 0x4d, 0x0a, 0x0d, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_package_proto_rawDescOnce sync.Once
	file_proto_v1_package_proto_rawDescData = file_proto_v1_package_proto_rawDesc
)

func file_proto_v1_package_proto_rawDescGZIP() []byte {
	file_proto_v1_package_proto_rawDescOnce.Do(func() {
		file_proto_v1_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_package_proto_rawDescData)
	})
	return file_proto_v1_package_proto_rawDescData
}

var file_proto_v1_package_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_v1_package_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_v1_package_proto_goTypes = []interface{}{
	(Architecture)(0),         // 0: grafeas.v1.Architecture
	(Version_VersionKind)(0),  // 1: grafeas.v1.Version.VersionKind
	(*Distribution)(nil),      // 2: grafeas.v1.Distribution
	(*Location)(nil),          // 3: grafeas.v1.Location
	(*PackageNote)(nil),       // 4: grafeas.v1.PackageNote
	(*PackageOccurrence)(nil), // 5: grafeas.v1.PackageOccurrence
	(*Version)(nil),           // 6: grafeas.v1.Version
	(*License)(nil),           // 7: grafeas.v1.License
	(*Digest)(nil),            // 8: grafeas.v1.Digest
}
var file_proto_v1_package_proto_depIdxs = []int32{
	0,  // 0: grafeas.v1.Distribution.architecture:type_name -> grafeas.v1.Architecture
	6,  // 1: grafeas.v1.Distribution.latest_version:type_name -> grafeas.v1.Version
	6,  // 2: grafeas.v1.Location.version:type_name -> grafeas.v1.Version
	2,  // 3: grafeas.v1.PackageNote.distribution:type_name -> grafeas.v1.Distribution
	0,  // 4: grafeas.v1.PackageNote.architecture:type_name -> grafeas.v1.Architecture
	6,  // 5: grafeas.v1.PackageNote.version:type_name -> grafeas.v1.Version
	7,  // 6: grafeas.v1.PackageNote.license:type_name -> grafeas.v1.License
	8,  // 7: grafeas.v1.PackageNote.digest:type_name -> grafeas.v1.Digest
	3,  // 8: grafeas.v1.PackageOccurrence.location:type_name -> grafeas.v1.Location
	0,  // 9: grafeas.v1.PackageOccurrence.architecture:type_name -> grafeas.v1.Architecture
	7,  // 10: grafeas.v1.PackageOccurrence.license:type_name -> grafeas.v1.License
	1,  // 11: grafeas.v1.Version.kind:type_name -> grafeas.v1.Version.VersionKind
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proto_v1_package_proto_init() }
func file_proto_v1_package_proto_init() {
	if File_proto_v1_package_proto != nil {
		return
	}
	file_proto_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution); i {
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
		file_proto_v1_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_proto_v1_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageNote); i {
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
		file_proto_v1_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageOccurrence); i {
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
		file_proto_v1_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_proto_v1_package_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1_package_proto_goTypes,
		DependencyIndexes: file_proto_v1_package_proto_depIdxs,
		EnumInfos:         file_proto_v1_package_proto_enumTypes,
		MessageInfos:      file_proto_v1_package_proto_msgTypes,
	}.Build()
	File_proto_v1_package_proto = out.File
	file_proto_v1_package_proto_rawDesc = nil
	file_proto_v1_package_proto_goTypes = nil
	file_proto_v1_package_proto_depIdxs = nil
}

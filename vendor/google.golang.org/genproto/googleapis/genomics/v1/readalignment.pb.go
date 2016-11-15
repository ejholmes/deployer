// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/genomics/v1/readalignment.proto
// DO NOT EDIT!

package google_genomics_v1 // import "google.golang.org/genproto/googleapis/genomics/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf3 "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A linear alignment can be represented by one CIGAR string. Describes the
// mapped position and local alignment of the read to the reference.
type LinearAlignment struct {
	// The position of this alignment.
	Position *Position `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	// The mapping quality of this alignment. Represents how likely
	// the read maps to this position as opposed to other locations.
	//
	// Specifically, this is -10 log10 Pr(mapping position is wrong), rounded to
	// the nearest integer.
	MappingQuality int32 `protobuf:"varint,2,opt,name=mapping_quality,json=mappingQuality" json:"mapping_quality,omitempty"`
	// Represents the local alignment of this sequence (alignment matches, indels,
	// etc) against the reference.
	Cigar []*CigarUnit `protobuf:"bytes,3,rep,name=cigar" json:"cigar,omitempty"`
}

func (m *LinearAlignment) Reset()                    { *m = LinearAlignment{} }
func (m *LinearAlignment) String() string            { return proto.CompactTextString(m) }
func (*LinearAlignment) ProtoMessage()               {}
func (*LinearAlignment) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *LinearAlignment) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *LinearAlignment) GetCigar() []*CigarUnit {
	if m != nil {
		return m.Cigar
	}
	return nil
}

// A read alignment describes a linear alignment of a string of DNA to a
// [reference sequence][google.genomics.v1.Reference], in addition to metadata
// about the fragment (the molecule of DNA sequenced) and the read (the bases
// which were read by the sequencer). A read is equivalent to a line in a SAM
// file. A read belongs to exactly one read group and exactly one
// [read group set][google.genomics.v1.ReadGroupSet].
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
//
// ### Reverse-stranded reads
//
// Mapped reads (reads having a non-null `alignment`) can be aligned to either
// the forward or the reverse strand of their associated reference. Strandedness
// of a mapped read is encoded by `alignment.position.reverseStrand`.
//
// If we consider the reference to be a forward-stranded coordinate space of
// `[0, reference.length)` with `0` as the left-most position and
// `reference.length` as the right-most position, reads are always aligned left
// to right. That is, `alignment.position.position` always refers to the
// left-most reference coordinate and `alignment.cigar` describes the alignment
// of this read to the reference from left to right. All per-base fields such as
// `alignedSequence` and `alignedQuality` share this same left-to-right
// orientation; this is true of reads which are aligned to either strand. For
// reverse-stranded reads, this means that `alignedSequence` is the reverse
// complement of the bases that were originally reported by the sequencing
// machine.
//
// ### Generating a reference-aligned sequence string
//
// When interacting with mapped reads, it's often useful to produce a string
// representing the local alignment of the read to reference. The following
// pseudocode demonstrates one way of doing this:
//
//     out = ""
//     offset = 0
//     for c in read.alignment.cigar {
//       switch c.operation {
//       case "ALIGNMENT_MATCH", "SEQUENCE_MATCH", "SEQUENCE_MISMATCH":
//         out += read.alignedSequence[offset:offset+c.operationLength]
//         offset += c.operationLength
//         break
//       case "CLIP_SOFT", "INSERT":
//         offset += c.operationLength
//         break
//       case "PAD":
//         out += repeat("*", c.operationLength)
//         break
//       case "DELETE":
//         out += repeat("-", c.operationLength)
//         break
//       case "SKIP":
//         out += repeat(" ", c.operationLength)
//         break
//       case "CLIP_HARD":
//         break
//       }
//     }
//     return out
//
// ### Converting to SAM's CIGAR string
//
// The following pseudocode generates a SAM CIGAR string from the
// `cigar` field. Note that this is a lossy conversion
// (`cigar.referenceSequence` is lost).
//
//     cigarMap = {
//       "ALIGNMENT_MATCH": "M",
//       "INSERT": "I",
//       "DELETE": "D",
//       "SKIP": "N",
//       "CLIP_SOFT": "S",
//       "CLIP_HARD": "H",
//       "PAD": "P",
//       "SEQUENCE_MATCH": "=",
//       "SEQUENCE_MISMATCH": "X",
//     }
//     cigarStr = ""
//     for c in read.alignment.cigar {
//       cigarStr += c.operationLength + cigarMap[c.operation]
//     }
//     return cigarStr
type Read struct {
	// The server-generated read ID, unique across all reads. This is different
	// from the `fragmentName`.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The ID of the read group this read belongs to. A read belongs to exactly
	// one read group. This is a server-generated ID which is distinct from SAM's
	// RG tag (for that value, see
	// [ReadGroup.name][google.genomics.v1.ReadGroup.name]).
	ReadGroupId string `protobuf:"bytes,2,opt,name=read_group_id,json=readGroupId" json:"read_group_id,omitempty"`
	// The ID of the read group set this read belongs to. A read belongs to
	// exactly one read group set.
	ReadGroupSetId string `protobuf:"bytes,3,opt,name=read_group_set_id,json=readGroupSetId" json:"read_group_set_id,omitempty"`
	// The fragment name. Equivalent to QNAME (query template name) in SAM.
	FragmentName string `protobuf:"bytes,4,opt,name=fragment_name,json=fragmentName" json:"fragment_name,omitempty"`
	// The orientation and the distance between reads from the fragment are
	// consistent with the sequencing protocol (SAM flag 0x2).
	ProperPlacement bool `protobuf:"varint,5,opt,name=proper_placement,json=properPlacement" json:"proper_placement,omitempty"`
	// The fragment is a PCR or optical duplicate (SAM flag 0x400).
	DuplicateFragment bool `protobuf:"varint,6,opt,name=duplicate_fragment,json=duplicateFragment" json:"duplicate_fragment,omitempty"`
	// The observed length of the fragment, equivalent to TLEN in SAM.
	FragmentLength int32 `protobuf:"varint,7,opt,name=fragment_length,json=fragmentLength" json:"fragment_length,omitempty"`
	// The read number in sequencing. 0-based and less than numberReads. This
	// field replaces SAM flag 0x40 and 0x80.
	ReadNumber int32 `protobuf:"varint,8,opt,name=read_number,json=readNumber" json:"read_number,omitempty"`
	// The number of reads in the fragment (extension to SAM flag 0x1).
	NumberReads int32 `protobuf:"varint,9,opt,name=number_reads,json=numberReads" json:"number_reads,omitempty"`
	// Whether this read did not pass filters, such as platform or vendor quality
	// controls (SAM flag 0x200).
	FailedVendorQualityChecks bool `protobuf:"varint,10,opt,name=failed_vendor_quality_checks,json=failedVendorQualityChecks" json:"failed_vendor_quality_checks,omitempty"`
	// The linear alignment for this alignment record. This field is null for
	// unmapped reads.
	Alignment *LinearAlignment `protobuf:"bytes,11,opt,name=alignment" json:"alignment,omitempty"`
	// Whether this alignment is secondary. Equivalent to SAM flag 0x100.
	// A secondary alignment represents an alternative to the primary alignment
	// for this read. Aligners may return secondary alignments if a read can map
	// ambiguously to multiple coordinates in the genome. By convention, each read
	// has one and only one alignment where both `secondaryAlignment`
	// and `supplementaryAlignment` are false.
	SecondaryAlignment bool `protobuf:"varint,12,opt,name=secondary_alignment,json=secondaryAlignment" json:"secondary_alignment,omitempty"`
	// Whether this alignment is supplementary. Equivalent to SAM flag 0x800.
	// Supplementary alignments are used in the representation of a chimeric
	// alignment. In a chimeric alignment, a read is split into multiple
	// linear alignments that map to different reference contigs. The first
	// linear alignment in the read will be designated as the representative
	// alignment; the remaining linear alignments will be designated as
	// supplementary alignments. These alignments may have different mapping
	// quality scores. In each linear alignment in a chimeric alignment, the read
	// will be hard clipped. The `alignedSequence` and
	// `alignedQuality` fields in the alignment record will only
	// represent the bases for its respective linear alignment.
	SupplementaryAlignment bool `protobuf:"varint,13,opt,name=supplementary_alignment,json=supplementaryAlignment" json:"supplementary_alignment,omitempty"`
	// The bases of the read sequence contained in this alignment record,
	// **without CIGAR operations applied** (equivalent to SEQ in SAM).
	// `alignedSequence` and `alignedQuality` may be
	// shorter than the full read sequence and quality. This will occur if the
	// alignment is part of a chimeric alignment, or if the read was trimmed. When
	// this occurs, the CIGAR for this read will begin/end with a hard clip
	// operator that will indicate the length of the excised sequence.
	AlignedSequence string `protobuf:"bytes,14,opt,name=aligned_sequence,json=alignedSequence" json:"aligned_sequence,omitempty"`
	// The quality of the read sequence contained in this alignment record
	// (equivalent to QUAL in SAM).
	// `alignedSequence` and `alignedQuality` may be shorter than the full read
	// sequence and quality. This will occur if the alignment is part of a
	// chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR
	// for this read will begin/end with a hard clip operator that will indicate
	// the length of the excised sequence.
	AlignedQuality []int32 `protobuf:"varint,15,rep,packed,name=aligned_quality,json=alignedQuality" json:"aligned_quality,omitempty"`
	// The mapping of the primary alignment of the
	// `(readNumber+1)%numberReads` read in the fragment. It replaces
	// mate position and mate strand in SAM.
	NextMatePosition *Position `protobuf:"bytes,16,opt,name=next_mate_position,json=nextMatePosition" json:"next_mate_position,omitempty"`
	// A map of additional read alignment information. This must be of the form
	// map<string, string[]> (string key mapping to a list of string values).
	Info map[string]*google_protobuf3.ListValue `protobuf:"bytes,17,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Read) Reset()                    { *m = Read{} }
func (m *Read) String() string            { return proto.CompactTextString(m) }
func (*Read) ProtoMessage()               {}
func (*Read) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *Read) GetAlignment() *LinearAlignment {
	if m != nil {
		return m.Alignment
	}
	return nil
}

func (m *Read) GetNextMatePosition() *Position {
	if m != nil {
		return m.NextMatePosition
	}
	return nil
}

func (m *Read) GetInfo() map[string]*google_protobuf3.ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*LinearAlignment)(nil), "google.genomics.v1.LinearAlignment")
	proto.RegisterType((*Read)(nil), "google.genomics.v1.Read")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/genomics/v1/readalignment.proto", fileDescriptor6)
}

var fileDescriptor6 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x4e, 0x1b, 0x47,
	0x14, 0xc7, 0xb5, 0x36, 0xa6, 0x78, 0x0c, 0xb6, 0x99, 0x4a, 0x74, 0x6b, 0x51, 0xd5, 0x35, 0x17,
	0x35, 0x17, 0xdd, 0x2d, 0xa0, 0xb6, 0xa8, 0x52, 0x94, 0x00, 0x22, 0x91, 0x23, 0x82, 0x9c, 0x45,
	0xe1, 0x76, 0x35, 0xde, 0x3d, 0x5e, 0x46, 0xec, 0xce, 0x0c, 0x33, 0xb3, 0x56, 0xfc, 0x48, 0x79,
	0xb7, 0x3c, 0x40, 0x2e, 0xa3, 0x99, 0xfd, 0x70, 0x48, 0x7c, 0x01, 0x57, 0xa0, 0xff, 0xf9, 0x9d,
	0xff, 0x9c, 0x3d, 0x1f, 0x46, 0x97, 0x09, 0xe7, 0x49, 0x0a, 0x5e, 0xc2, 0x53, 0xc2, 0x12, 0x8f,
	0xcb, 0xc4, 0x4f, 0x80, 0x09, 0xc9, 0x35, 0xf7, 0x8b, 0x10, 0x11, 0x54, 0x19, 0x8d, 0x67, 0x34,
	0x52, 0xfe, 0xe2, 0xc8, 0x97, 0x40, 0x62, 0x92, 0xd2, 0x84, 0x65, 0xc0, 0xb4, 0x67, 0x51, 0x8c,
	0x2b, 0x9b, 0x92, 0xf3, 0x16, 0x47, 0x83, 0xc9, 0xd3, 0xac, 0x89, 0xa0, 0xbe, 0x02, 0xb9, 0xa0,
	0x11, 0x44, 0x9c, 0xcd, 0x69, 0xe2, 0x13, 0xc6, 0xb8, 0x26, 0x9a, 0x72, 0xa6, 0x0a, 0xfb, 0xc1,
	0x8b, 0xe7, 0x57, 0x19, 0xd1, 0x84, 0xc8, 0x32, 0xfd, 0xd5, 0xf3, 0xd3, 0x05, 0x57, 0xd4, 0x54,
	0x50, 0x3a, 0xfc, 0x93, 0x50, 0x7d, 0x97, 0xcf, 0xbc, 0x88, 0x67, 0x7e, 0xe1, 0xe2, 0xdb, 0xc0,
	0x2c, 0x9f, 0xfb, 0x42, 0x2f, 0x05, 0x28, 0x5f, 0x69, 0x99, 0x47, 0xba, 0xfc, 0x53, 0xa4, 0x8d,
	0x3e, 0x39, 0xa8, 0x77, 0x45, 0x19, 0x10, 0x79, 0x56, 0x35, 0x0c, 0x9f, 0xa2, 0xad, 0xca, 0xdc,
	0x75, 0x86, 0xce, 0xb8, 0x73, 0xbc, 0xef, 0xfd, 0xd8, 0x3d, 0x6f, 0x5a, 0x32, 0x41, 0x4d, 0xe3,
	0x3f, 0x51, 0x2f, 0x23, 0x42, 0x50, 0x96, 0x84, 0x0f, 0x39, 0x49, 0xa9, 0x5e, 0xba, 0x8d, 0xa1,
	0x33, 0x6e, 0x05, 0xdd, 0x52, 0x7e, 0x5f, 0xa8, 0xf8, 0x04, 0xb5, 0xec, 0xe7, 0xbb, 0xcd, 0x61,
	0x73, 0xdc, 0x39, 0xfe, 0x6d, 0x9d, 0xff, 0x85, 0x01, 0x3e, 0x30, 0xaa, 0x83, 0x82, 0x1d, 0x7d,
	0xde, 0x44, 0x1b, 0x01, 0x90, 0x18, 0x77, 0x51, 0x83, 0xc6, 0xb6, 0xb4, 0x76, 0xd0, 0xa0, 0x31,
	0x1e, 0xa1, 0x1d, 0x33, 0xf2, 0x30, 0x91, 0x3c, 0x17, 0x21, 0x8d, 0xed, 0xa3, 0xed, 0xa0, 0x63,
	0xc4, 0x37, 0x46, 0x9b, 0xc4, 0xf8, 0x10, 0xed, 0x7e, 0xc3, 0x28, 0xd0, 0x86, 0x6b, 0x5a, 0xae,
	0x5b, 0x73, 0x37, 0xa0, 0x27, 0x31, 0x3e, 0x40, 0x3b, 0x73, 0x49, 0x12, 0xd3, 0x8b, 0x90, 0x91,
	0x0c, 0xdc, 0x0d, 0x8b, 0x6d, 0x57, 0xe2, 0x35, 0xc9, 0x00, 0x1f, 0xa2, 0xbe, 0x90, 0x5c, 0x80,
	0x0c, 0x45, 0x4a, 0x22, 0x30, 0xba, 0xdb, 0x1a, 0x3a, 0xe3, 0xad, 0xa0, 0x57, 0xe8, 0xd3, 0x4a,
	0xc6, 0x7f, 0x21, 0x1c, 0xe7, 0x22, 0xa5, 0x11, 0xd1, 0x10, 0x56, 0x26, 0xee, 0xa6, 0x85, 0x77,
	0xeb, 0xc8, 0xeb, 0x32, 0x60, 0x9a, 0x58, 0x3f, 0x9f, 0x02, 0x4b, 0xf4, 0x9d, 0xfb, 0x53, 0xd1,
	0xc4, 0x4a, 0xbe, 0xb2, 0x2a, 0xfe, 0x1d, 0xd9, 0x2f, 0x0c, 0x59, 0x9e, 0xcd, 0x40, 0xba, 0x5b,
	0x16, 0x42, 0x46, 0xba, 0xb6, 0x0a, 0xfe, 0x03, 0x6d, 0x17, 0xb1, 0xd0, 0x88, 0xca, 0x6d, 0x5b,
	0xa2, 0x53, 0x68, 0xa6, 0x93, 0x0a, 0xbf, 0x44, 0xfb, 0x73, 0x42, 0x53, 0x88, 0xc3, 0x05, 0xb0,
	0x98, 0xcb, 0x6a, 0x6e, 0x61, 0x74, 0x07, 0xd1, 0xbd, 0x72, 0x91, 0xad, 0xf2, 0xd7, 0x82, 0xb9,
	0xb5, 0x48, 0x39, 0xc3, 0x0b, 0x0b, 0xe0, 0x33, 0xd4, 0xae, 0x4f, 0xcd, 0xed, 0xd8, 0x6d, 0x39,
	0x58, 0x37, 0xcd, 0xef, 0x96, 0x2c, 0x58, 0x65, 0x61, 0x1f, 0xfd, 0xac, 0xcc, 0x65, 0xc5, 0x44,
	0x2e, 0xc3, 0x95, 0xd9, 0xb6, 0x7d, 0x1a, 0xd7, 0xa1, 0xd5, 0x82, 0xfe, 0x87, 0x7e, 0x51, 0xb9,
	0x10, 0xa9, 0x6d, 0xef, 0xe3, 0xa4, 0x1d, 0x9b, 0xb4, 0xf7, 0x28, 0xbc, 0x4a, 0x3c, 0x44, 0x7d,
	0x8b, 0x42, 0x1c, 0x2a, 0x78, 0xc8, 0x81, 0x45, 0xe0, 0x76, 0xed, 0x70, 0x7b, 0xa5, 0x7e, 0x53,
	0xca, 0x66, 0x0a, 0x15, 0x5a, 0xad, 0x72, 0x6f, 0xd8, 0x34, 0x53, 0x28, 0xe5, 0x6a, 0x95, 0xdf,
	0x22, 0xcc, 0xe0, 0xa3, 0x0e, 0x33, 0x33, 0xdd, 0xfa, 0x6e, 0xfa, 0x4f, 0xb8, 0x9b, 0xbe, 0xc9,
	0x7b, 0x47, 0x34, 0x54, 0x0a, 0xfe, 0x17, 0x6d, 0x50, 0x36, 0xe7, 0xee, 0xae, 0xbd, 0x8a, 0xd1,
	0xba, 0x6c, 0x33, 0x36, 0x6f, 0xc2, 0xe6, 0xfc, 0x92, 0x69, 0xb9, 0x0c, 0x2c, 0x3f, 0xb8, 0x41,
	0xed, 0x5a, 0xc2, 0x7d, 0xd4, 0xbc, 0x87, 0x65, 0x79, 0x1e, 0xe6, 0x5f, 0xfc, 0x37, 0x6a, 0x2d,
	0x48, 0x9a, 0x83, 0xbd, 0x8b, 0xce, 0xf1, 0xa0, 0xf2, 0xad, 0x7e, 0x20, 0xbc, 0x2b, 0xaa, 0xf4,
	0xad, 0x21, 0x82, 0x02, 0xfc, 0xbf, 0x71, 0xea, 0x9c, 0x1f, 0xa1, 0xbd, 0x88, 0x67, 0x6b, 0x6a,
	0x38, 0xc7, 0xa6, 0x88, 0xba, 0xab, 0x53, 0xe3, 0x32, 0x75, 0xbe, 0x38, 0xce, 0x6c, 0xd3, 0x3a,
	0x9e, 0x7c, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xa0, 0xb1, 0xbc, 0xb4, 0x05, 0x00, 0x00,
}

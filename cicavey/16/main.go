package main

import (
	"aoc2021/aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type FakeBitStream struct {
	bits string
	idx  int
}

func NewFakeBitStream(bits string) *FakeBitStream {
	return &FakeBitStream{bits, 0}
}

func NewFakeBitStreamFromHex(hex string) *FakeBitStream {
	hexMap := make(map[rune]string)
	hexMap['0'] = "0000"
	hexMap['1'] = "0001"
	hexMap['2'] = "0010"
	hexMap['3'] = "0011"
	hexMap['4'] = "0100"
	hexMap['5'] = "0101"
	hexMap['6'] = "0110"
	hexMap['7'] = "0111"
	hexMap['8'] = "1000"
	hexMap['9'] = "1001"
	hexMap['A'] = "1010"
	hexMap['B'] = "1011"
	hexMap['C'] = "1100"
	hexMap['D'] = "1101"
	hexMap['E'] = "1110"
	hexMap['F'] = "1111"
	var sb strings.Builder
	for _, c := range hex {
		sb.WriteString(hexMap[c])
	}
	return &FakeBitStream{sb.String(), 0}
}

// ReadInt inteprets the next bitSize bits as a big-endian integer
func (bs *FakeBitStream) ReadInt(bitSize int) int64 {
	value, _ := strconv.ParseInt(bs.bits[bs.idx:bs.idx+bitSize], 2, 64)
	bs.idx += bitSize
	return value
}

// Fixed size bitstream
func (bs *FakeBitStream) ReadBitStream(bitSize int) *FakeBitStream {
	value := NewFakeBitStream(bs.bits[bs.idx : bs.idx+bitSize])
	bs.idx += bitSize
	return value
}

// DOES NOT ALTER PARENT IDX
func (bs *FakeBitStream) SubStream() *FakeBitStream {
	return NewFakeBitStream(bs.bits[bs.idx:])
}

func (bs *FakeBitStream) ReadBool() bool {
	value := bs.bits[bs.idx] == '1'
	bs.idx += 1
	return value
}

func (bs *FakeBitStream) ReadChunk(bitSize int) string {
	chunk := bs.bits[bs.idx : bs.idx+bitSize]
	bs.idx += bitSize
	return chunk
}

func (bs *FakeBitStream) EOS() bool {
	return bs.idx >= len(bs.bits)
}

func (bs *FakeBitStream) Remaining() int {
	return len(bs.bits) - bs.idx
}

const PKT_TYPE_LITERAL = 4

type Packet struct {
	version    int64
	typeID     int64
	literal    int64
	subPackets Packets
}

type Packets []Packet

func (p Packet) String() string {
	return fmt.Sprintf("Type: %d, v: %d, sp: %v", p.typeID, p.literal, p.subPackets)
}

func NewLiteralPacket(version int64, literal int64) Packet {
	return Packet{
		version: version,
		typeID:  PKT_TYPE_LITERAL,
		literal: literal,
	}
}

func NewOperatorPacket(version int64, typeID int64, subpkts Packets) Packet {
	return Packet{
		version:    version,
		typeID:     typeID,
		subPackets: subpkts,
	}
}

func ParsePackets(bs *FakeBitStream, limit int) Packets {

	var packets Packets

	pktCount := 0
	// 11 is the smallest possible packet size
	// 3 + 3 + 5            = 11
	// 3 + 3 + 1 + 11 + ... = 18+
	// 3 + 3 + 1 + 15 + ... = 21+
	for bs.Remaining() >= 11 {

		version := bs.ReadInt(3)
		typeID := bs.ReadInt(3)

		if typeID == PKT_TYPE_LITERAL {
			// read 5 bits at a time

			var acc strings.Builder
			for {
				chunk := bs.ReadChunk(5)
				acc.WriteString(chunk[1:])
				if chunk[0] == '0' {
					break
				}
			}
			val, _ := strconv.ParseInt(acc.String(), 2, 64)

			packets = append(packets, NewLiteralPacket(version, val))
		} else {
			// operator packet
			lengthTypeID := bs.ReadBool()
			var subpkts Packets
			if lengthTypeID {
				// 11 bits
				subPacketCount := bs.ReadInt(11)
				subbs := bs.SubStream()
				subpkts = ParsePackets(subbs, int(subPacketCount))
				bs.idx += subbs.idx
			} else {
				// 15 bits
				subPacketBits := bs.ReadInt(15)
				subbs := bs.ReadBitStream(int(subPacketBits))
				subpkts = ParsePackets(subbs, -1)
			}

			packets = append(packets, NewOperatorPacket(version, typeID, subpkts))
		}

		pktCount++

		if pktCount == limit {
			break
		}
	}

	return packets
}

func SumVersions(pkts Packets) int64 {
	var sum int64
	for _, pkt := range pkts {
		sum += pkt.version
		if len(pkt.subPackets) > 0 {
			sum += SumVersions(pkt.subPackets)
		}
	}
	return sum
}

func EvalPacket(pkt Packet) int64 {

	switch pkt.typeID {
	case 0: //sum
		var sum int64
		for _, spkt := range pkt.subPackets {
			sum += EvalPacket(spkt)
		}
		return sum
	case 1: //prod
		var prod int64 = 1
		for _, spkt := range pkt.subPackets {
			prod *= EvalPacket(spkt)
		}
		return prod
	case 2: //min
		var min int64 = math.MaxInt64
		for _, spkt := range pkt.subPackets {
			v := EvalPacket(spkt)
			if v < min {
				min = v
			}
		}
		return min
	case 3: //max
		var max int64 = math.MinInt64
		for _, spkt := range pkt.subPackets {
			v := EvalPacket(spkt)
			if v > max {
				max = v
			}
		}
		return max
	case 4: //literal
		return pkt.literal // we return here because a top-level literal never makes sense
	case 5: //gte
		a := EvalPacket(pkt.subPackets[0])
		b := EvalPacket(pkt.subPackets[1])
		if a > b {
			return 1
		} else {
			return 0
		}
	case 6: //lte
		a := EvalPacket(pkt.subPackets[0])
		b := EvalPacket(pkt.subPackets[1])
		if a < b {
			return 1
		} else {
			return 0
		}
	case 7: //eq
		a := EvalPacket(pkt.subPackets[0])
		b := EvalPacket(pkt.subPackets[1])
		if a == b {
			return 1
		} else {
			return 0
		}
	}

	return 0
}

func main() {

	lines := aoc.ReadLines("input.txt")

	hex := lines[0]

	bs := NewFakeBitStreamFromHex(hex)

	pkts := ParsePackets(bs, -1)
	fmt.Println(SumVersions(pkts))

	// First packet has to be an operator...
	fmt.Println(EvalPacket(pkts[0]))
}

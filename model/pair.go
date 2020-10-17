package model

import (
	ptypes "github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	pb "github.com/strawhatboy/map_reduce/proto"
)

type PairCount struct {
	First  string
	Second string
}

func (p PairCount) toPbAny() *any.Any {
	m := &pb.StringMessage{Str: p.Second}
	a, _ := ptypes.MarshalAny(m)
	return a
}

func (p PairCount) ToPbPair() *pb.MapPair {
	return &pb.MapPair{First: p.First, Second: p.toPbAny()}
}

func (p PairCount) FromPbPair(pp *pb.MapPair) {
	strMsg := pb.StringMessage{}
	ptypes.UnmarshalAny(pp.Second, &strMsg)
	p.First = pp.First
	p.Second = strMsg.Str
}

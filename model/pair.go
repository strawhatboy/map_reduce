package model

import (
	ptypes "github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	pb "github.com/strawhatboy/map_reduce/proto"
)

type PairCount struct {
	First  string
	Second int64
}

func (p PairCount) toPbAny() *any.Any {
	m := &pb.NumberMessage{Num: p.Second}
	a, _ := ptypes.MarshalAny(m)
	return a
}

func (p PairCount) ToPbPair() *pb.MapPair {
	return &pb.MapPair{First: p.First, Second: p.toPbAny()}
}

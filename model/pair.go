package model

import (
	ptypes "github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	pb "github.com/strawhatboy/map_reduce/proto"
	log "github.com/sirupsen/logrus"
)

type PairCount struct {
	First  string
	Second string
}

func (p PairCount) toPbAny() *any.Any {
	m := &pb.StringMessage{Str: p.Second}
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		log.Error("failed convert PairCount: ", p, " to PbAny, err: ", err)
	}
	return a
}

func (p PairCount) ToPbPair() *pb.MapPair {
	return &pb.MapPair{First: p.First, Second: p.Second}
}

func (p PairCount) FromPbPair(pp *pb.MapPair) {
	// strMsg := pb.StringMessage{}
	// ptypes.UnmarshalAny(pp.Second, &strMsg)
	p.First = pp.First
	// p.Second = strMsg.Str
	p.Second = pp.Second
}

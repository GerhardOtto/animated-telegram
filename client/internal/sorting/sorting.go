package sorting

import pb "github.com/gerhardotto/animated-telegram/client/backendservice"

// CompareFunc defines the comparator signature used by all sort algorithms.
type CompareFunc func(a, b *pb.DataItem) bool

// ByIntValAsc orders DataItems by their integer value ascending.
var ByIntValAsc CompareFunc = func(a, b *pb.DataItem) bool {
	return a.GetIntVal() < b.GetIntVal()
}

// ByStringValAsc orders DataItems by their string value ascending.
var ByStringValAsc CompareFunc = func(a, b *pb.DataItem) bool {
	return a.GetStringval() < b.GetStringval()
}

// ByIntValDesc orders DataItems by their integer value descending.
var ByIntValDesc CompareFunc = func(a, b *pb.DataItem) bool {
	return a.GetIntVal() > b.GetIntVal()
}

// ByStringValDesc orders DataItems by their string value descending.
var ByStringValDesc CompareFunc = func(a, b *pb.DataItem) bool {
	return a.GetStringval() > b.GetStringval()
}

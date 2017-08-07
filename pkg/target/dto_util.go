package target

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

const (
	defaultInfiniteCapacity = 1E10
	defaultTransactionCapacity = 100
)

// no capacity for bought commodity
func CreateResourceCommodityBought(res *Resource, ctype proto.CommodityDTO_CommodityType) (*proto.CommodityDTO, error) {
	return builder.NewCommodityDTOBuilder(ctype).
		Used(res.Used).
		Create()
}

func CreateResourceCommodity(res *Resource, ctype proto.CommodityDTO_CommodityType) (*proto.CommodityDTO, error) {
	return builder.NewCommodityDTOBuilder(ctype).
		Capacity(res.Capacity).
		Used(res.Used).
		Create()
}

func CreateKeyCommodity(key string, ctype proto.CommodityDTO_CommodityType) (*proto.CommodityDTO, error) {
	return builder.
		NewCommodityDTOBuilder(ctype).
		Key(key).
		Capacity(defaultInfiniteCapacity).
		Create()
}


func CreateTransactionCommodity(key string, ctype proto.CommodityDTO_CommodityType) (*proto.CommodityDTO, error) {
	return builder.
	NewCommodityDTOBuilder(ctype).
		Key(key).
		Capacity(defaultTransactionCapacity).
		Create()
}
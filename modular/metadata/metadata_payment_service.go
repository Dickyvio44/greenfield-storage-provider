package metadata

import (
	"context"

	"cosmossdk.io/math"
	payment_types "github.com/bnb-chain/greenfield/x/payment/types"

	"github.com/bnb-chain/greenfield-storage-provider/modular/metadata/types"
	"github.com/bnb-chain/greenfield-storage-provider/pkg/log"
	model "github.com/bnb-chain/greenfield-storage-provider/store/bsdb"
)

// GfSpGetPaymentByBucketName get bucket payment info by a bucket name
func (r *MetadataModular) GfSpGetPaymentByBucketName(ctx context.Context, req *types.GfSpGetPaymentByBucketNameRequest) (resp *types.GfSpGetPaymentByBucketNameResponse, err error) {
	var (
		streamRecord *model.StreamRecord
		res          *payment_types.StreamRecord
	)

	ctx = log.Context(ctx, req)

	streamRecord, err = r.baseApp.GfBsDB().GetPaymentByBucketName(req.BucketName, req.IncludePrivate)
	if err != nil {
		log.CtxErrorw(ctx, "failed to get payment by bucket name", "error", err)
		return
	}

	if streamRecord != nil {
		res = &payment_types.StreamRecord{
			Account:           streamRecord.Account.String(),
			CrudTimestamp:     streamRecord.CrudTimestamp,
			NetflowRate:       math.NewIntFromBigInt(streamRecord.NetflowRate.Raw()),
			StaticBalance:     math.NewIntFromBigInt(streamRecord.StaticBalance.Raw()),
			BufferBalance:     math.NewIntFromBigInt(streamRecord.BufferBalance.Raw()),
			LockBalance:       math.NewIntFromBigInt(streamRecord.LockBalance.Raw()),
			Status:            payment_types.StreamAccountStatus(payment_types.StreamAccountStatus_value[streamRecord.Status]),
			SettleTimestamp:   streamRecord.SettleTimestamp,
			OutFlowCount:      streamRecord.OutFlowCount,
			FrozenNetflowRate: math.NewIntFromBigInt(streamRecord.FrozenNetflowRate.Raw()),
		}
	}

	resp = &types.GfSpGetPaymentByBucketNameResponse{StreamRecord: res}
	log.CtxInfow(ctx, "succeed to get payment by bucket name")
	return resp, nil
}

// GfSpGetPaymentByBucketID get bucket payment info by a bucket id
func (r *MetadataModular) GfSpGetPaymentByBucketID(ctx context.Context, req *types.GfSpGetPaymentByBucketIDRequest) (resp *types.GfSpGetPaymentByBucketIDResponse, err error) {
	var (
		streamRecord *model.StreamRecord
		res          *payment_types.StreamRecord
	)

	ctx = log.Context(ctx, req)

	streamRecord, err = r.baseApp.GfBsDB().GetPaymentByBucketID(req.BucketId, req.IncludePrivate)
	if err != nil {
		log.CtxErrorw(ctx, "failed to get payment by bucket id", "error", err)
		return
	}

	if streamRecord != nil {
		res = &payment_types.StreamRecord{
			Account:           streamRecord.Account.String(),
			CrudTimestamp:     streamRecord.CrudTimestamp,
			NetflowRate:       math.NewIntFromBigInt(streamRecord.NetflowRate.Raw()),
			StaticBalance:     math.NewIntFromBigInt(streamRecord.StaticBalance.Raw()),
			BufferBalance:     math.NewIntFromBigInt(streamRecord.BufferBalance.Raw()),
			LockBalance:       math.NewIntFromBigInt(streamRecord.LockBalance.Raw()),
			Status:            payment_types.StreamAccountStatus(payment_types.StreamAccountStatus_value[streamRecord.Status]),
			SettleTimestamp:   streamRecord.SettleTimestamp,
			OutFlowCount:      streamRecord.OutFlowCount,
			FrozenNetflowRate: math.NewIntFromBigInt(streamRecord.FrozenNetflowRate.Raw()),
		}
	}

	resp = &types.GfSpGetPaymentByBucketIDResponse{StreamRecord: res}
	log.CtxInfow(ctx, "succeed to get payment by bucket id")
	return resp, nil
}

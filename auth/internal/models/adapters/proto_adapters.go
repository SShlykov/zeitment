package adapters

import (
	"github.com/SShlykov/zeitment/auth/internal/models/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

func ProtoToNullDt(dt *timestamppb.Timestamp) types.Null[time.Time] {
	if dt != nil {
		return types.Null[time.Time]{Valid: true, Value: dt.AsTime()}
	}
	return types.Null[time.Time]{Valid: false}
}

func ProtoToNullInt(i *wrapperspb.Int32Value) types.Null[int] {
	if i != nil {
		return types.Null[int]{Valid: true, Value: int(i.Value)}
	}
	return types.Null[int]{Valid: false}
}

func ProtoToNullInt64(i *wrapperspb.Int64Value) types.Null[int64] {
	if i != nil {
		return types.Null[int64]{Valid: true, Value: i.Value}
	}
	return types.Null[int64]{Valid: false}
}

func ProtoToNullString(str *wrapperspb.StringValue) types.Null[string] {
	if str != nil {
		return types.Null[string]{Valid: true, Value: str.Value}
	}
	return types.Null[string]{Valid: false}
}

func NullInt64ToProto(i types.Null[int64]) *wrapperspb.Int64Value {
	if i.Valid {
		return wrapperspb.Int64(i.Value)
	}
	return wrapperspb.Int64(0)
}

func NullIntToProto(i types.Null[int]) *wrapperspb.Int32Value {
	if i.Valid {
		return wrapperspb.Int32(int32(i.Value))
	}
	return wrapperspb.Int32(0)
}

func NullStringToProto(str types.Null[string]) *wrapperspb.StringValue {
	if str.Valid {
		return wrapperspb.String(str.Value)
	}
	return nil
}

func NullDtToProto(dt types.Null[time.Time]) *timestamppb.Timestamp {
	if dt.Valid {
		return timestamppb.New(dt.Value)
	}
	return nil
}

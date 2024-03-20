package adapters

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

func ProtoToNullDt(dt *timestamppb.Timestamp) sql.Null[time.Time] {
	if dt != nil {
		return sql.Null[time.Time]{Valid: true, V: dt.AsTime()}
	}
	return sql.Null[time.Time]{Valid: false}
}

func ProtoToNullInt(i *wrapperspb.Int32Value) sql.Null[int] {
	if i != nil {
		return sql.Null[int]{Valid: true, V: int(i.Value)}
	}
	return sql.Null[int]{Valid: false}
}

func ProtoToNullInt64(i *wrapperspb.Int64Value) sql.Null[int64] {
	if i != nil {
		return sql.Null[int64]{Valid: true, V: i.Value}
	}
	return sql.Null[int64]{Valid: false}
}

func ProtoToNullString(str *wrapperspb.StringValue) sql.Null[string] {
	if str != nil {
		return sql.Null[string]{Valid: true, V: str.Value}
	}
	return sql.Null[string]{Valid: false}
}

func NullInt64ToProto(i sql.Null[int64]) *wrapperspb.Int64Value {
	if i.Valid {
		return wrapperspb.Int64(i.V)
	}
	return wrapperspb.Int64(0)
}

func NullIntToProto(i sql.Null[int]) *wrapperspb.Int32Value {
	if i.Valid {
		return wrapperspb.Int32(int32(i.V))
	}
	return wrapperspb.Int32(0)
}

func NullStringToProto(str sql.Null[string]) *wrapperspb.StringValue {
	if str.Valid {
		return wrapperspb.String(str.V)
	}
	return nil
}

func NullDtToProto(dt sql.Null[time.Time]) *timestamppb.Timestamp {
	if dt.Valid {
		return timestamppb.New(dt.V)
	}
	return nil
}

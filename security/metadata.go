package security

import (
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	SignAppidKey = "wm-appid"
	SignSignKey  = "wm-sign"
	SignTimeKey  = "wm-time"
)

type SignInfo struct {
	appid string
	sign  string
	time  int64
}

func (s *SignInfo) Verify(sign string) bool {
	return s.sign == sign
}

func (s *SignInfo) Before(t int64) bool {
	return s.time < t
}

func (s *SignInfo) Appid() string {
	return s.appid
}

func (s *SignInfo) Sign() string {
	return s.sign
}

func (s *SignInfo) Time() int64 {
	return s.time
}

func ParseSignFromIncommingMetadata(incomming metadata.MD) (*SignInfo, error) {
	if len(incomming.Get(SignAppidKey)) == 0 ||
		len(incomming.Get(SignTimeKey)) == 0 ||
		len(incomming.Get(SignSignKey)) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing auth fileds")
	}

	signTime, err := strconv.ParseInt(incomming.Get(SignTimeKey)[0], 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid sign time")
	}

	return &SignInfo{
		appid: incomming.Get(SignAppidKey)[0],
		sign:  incomming.Get(SignSignKey)[0],
		time:  signTime,
	}, nil
}

package security

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/spacegrower/tools/utils"
)

const (
	TOKEN_KEY = "sg-token"
)

func GenSign(appid, secret string, signTime int64) string {
	signString := fmt.Sprintf("appid=%s&secret=%s&time=%d", appid, secret, signTime)
	signString += fmt.Sprintf("%s&key=%s", signString, utils.MD5(signString))
	return utils.MD5(signString)
}

type TokenClaims struct {
	Appid        string            `json:"aid"`
	ServiceTitle string            `json:"st"`
	User         string            `json:"u"` // 对应平台的用户唯一标识
	Fields       map[string]string `json:"f"`
	ExpireTime   int64             `json:"exp"` // 过期时间 时间戳
	NotBefore    int64             `json:"nbf"` // 生效时间 时间戳
}

func (t *TokenClaims) UserInt64() int64 {
	if t.User == "" {
		return 0
	}
	id, _ := strconv.ParseInt(t.User, 10, 64)
	return id
}

func (t *TokenClaims) Field(key string) string {
	if t.Fields == nil {
		return ""
	}

	return t.Fields[key]
}

func GenerateJWT(info TokenClaims, signBytes []byte) (string, error) {
	claims := jwt.MapClaims{}

	t := reflect.TypeOf(info)
	v := reflect.ValueOf(info)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		claims[tag] = v.Field(i).Interface()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}
	return token.SignedString(privateKey)
}

func ParseJWT(tokenString string, key []byte) (*TokenClaims, error) {
	result := &TokenClaims{}
	_, err := jwt.Parse(tokenString, func(i2 *jwt.Token) (i interface{}, e error) {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("invalied public key, %s", err.Error()))
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	parts := strings.Split(tokenString, ".")
	claimBytes, _ := jwt.DecodeSegment(parts[1])

	if err = json.Unmarshal(claimBytes, &result); err != nil {
		return result, fmt.Errorf("invalid jwt, %w", err)
	}
	return result, nil
}

func NewPerRPCCredentialOfSign(appid, secret string) *perRPCCredentialOfSign {
	return &perRPCCredentialOfSign{
		appid:  appid,
		secret: secret,
	}
}

// perRPCCredential implements "grpccredentials.PerRPCCredentials" interface.
type perRPCCredentialOfSign struct {
	appid  string
	secret string
}

func (rc *perRPCCredentialOfSign) RequireTransportSecurity() bool { return false }

func (rc *perRPCCredentialOfSign) GetRequestMetadata(ctx context.Context, s ...string) (map[string]string, error) {
	now := time.Now().Unix()
	return map[string]string{signAppidKey: rc.appid, signSignKey: GenSign(rc.appid, rc.secret, now), signTimeKey: strconv.FormatInt(now, 10)}, nil
}

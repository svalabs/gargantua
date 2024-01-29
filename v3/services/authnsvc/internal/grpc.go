package authnservice

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/hobbyfarm/gargantua/v3/pkg/errors"
	authnProto "github.com/hobbyfarm/gargantua/v3/protos/authn"
	userProto "github.com/hobbyfarm/gargantua/v3/protos/user"
	"google.golang.org/grpc/codes"
)

type GrpcAuthnServer struct {
	authnProto.UnimplementedAuthNServer
	userClient userProto.UserSvcClient
}

func NewGrpcAuthNServer(userClient userProto.UserSvcClient) *GrpcAuthnServer {
	return &GrpcAuthnServer{userClient: userClient}
}

func (a *GrpcAuthnServer) AuthN(c context.Context, ar *authnProto.AuthNRequest) (*userProto.User, error) {
	token := ar.GetToken()
	if len(token) == 0 {
		glog.Errorf("no bearer token passed, authentication failed")
		return &userProto.User{}, errors.GrpcError(
			codes.InvalidArgument,
			"missing the following properties from type 'AuthNRequest': token",
			ar,
		)
	}

	var finalToken string

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) == 1 {
		finalToken = strings.TrimSpace(splitToken[0])
	} else {
		finalToken = strings.TrimSpace(splitToken[1])
	}

	user, err := a.validateToken(c, finalToken)
	if err != nil {
		glog.Infof("could not validate token: %s", err)
		return &userProto.User{}, errors.GrpcError(
			codes.Unauthenticated,
			"could not validate token: %s",
			ar,
			err,
		)
	}
	return user, nil
}

func (a *GrpcAuthnServer) validateToken(ctx context.Context, token string) (*userProto.User, error) {
	user, err := a.validate(ctx, token)

	if err != nil {
		glog.Errorf("error validating user %v", err)
		return &userProto.User{}, fmt.Errorf("authentication failed")
	}

	return user, nil
}

func (a *GrpcAuthnServer) validate(ctx context.Context, tokenString string) (*userProto.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		var user *userProto.User
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			var err error
			user, err = a.userClient.GetUserByEmail(ctx, &userProto.GetUserByEmailRequest{Email: fmt.Sprint(claims["email"])})
			if err != nil {
				glog.Errorf("could not find user that matched token %s", fmt.Sprint(claims["email"]))
				return &userProto.User{}, fmt.Errorf("could not find user that matched token %s", fmt.Sprint(claims["email"]))
			}
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(user.Password), nil
	})

	if err != nil {
		glog.Errorf("error while validating user: %v", err)
		return &userProto.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := a.userClient.GetUserByEmail(ctx, &userProto.GetUserByEmailRequest{Email: fmt.Sprint(claims["email"])})
		if err != nil {
			return &userProto.User{}, err
		} else {
			return user, nil
		}
	}
	glog.Errorf("error while validating user")
	return &userProto.User{}, fmt.Errorf("error while validating user")
}

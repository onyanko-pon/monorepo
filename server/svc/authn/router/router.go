package router

import (
	svchandler "github.com/onyanko-pon/monorepo/server/svc/authn/svc_handler"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

func SvcRouter() error {
	authn, err := svchandler.InitAuthn()
	if err != nil {
		return err
	}

	if err := svcrouter.AddHandler(svcrouter.UserTokenVerify, authn.VerifyToken); err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.UserTokenEncode, authn.EncodeToken); err != nil {
		return err
	}

	twauth1, err := svchandler.InitTwitterAuth1()
	if err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.TwitterOAuth1FetchAccessToken, twauth1.GetAccessToken); err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.TwitterOAuth1VerifyAccessToken, twauth1.VerifyAccessToken); err != nil {
		return err
	}

	return nil
}

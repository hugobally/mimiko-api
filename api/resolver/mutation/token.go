package mutation

import (
	"context"
	"errors"
	"github.com/hugobally/mimiko/backend/auth"
	"github.com/hugobally/mimiko/backend/prisma"
	"time"
)

func (r *Resolver) GetToken(ctx context.Context, appType prisma.AppType) (*prisma.LinkedApp, error) {
	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	apps, err := r.Prisma.LinkedApps(&prisma.LinkedAppsParams{
		Where: &prisma.LinkedAppWhereInput{
			User: &prisma.UserWhereInput{ID: &user.ID},
			Type: &appType,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}
	if len(apps) == 0 {
		return nil, errors.New("no linked app found for the user")
	}

	app := &apps[0]
	exp, _ := time.Parse(time.RFC3339, *app.TokenExpiry)

	if exp.After(time.Now().Add(5 * time.Minute)) {
		return app, nil
	}

	// Sample session token refresh

	newToken, err := r.Spotify.CreateClientCredentialsToken()
	expStr := time.Now().Add(time.Duration(newToken.ExpiresIn) * time.Second).Format(time.RFC3339)
	app, err = r.Prisma.UpdateLinkedApp(prisma.LinkedAppUpdateParams{
		Data: prisma.LinkedAppUpdateInput{
			AccessToken:  &newToken.AccessToken,
			TokenExpiry:  &expStr,
			RefreshToken: nil,
		},
		Where: prisma.LinkedAppWhereUniqueInput{
			ID: &app.ID,
		},
	}).Exec(ctx)

	// Uncomment this when re-implementing proper spotify login ; for now we're always in sample session mode
	//
	//if app.RefreshToken == nil {
	//	newToken, err := r.Spotify.CreateClientCredentialsToken()
	//	expStr := time.Now().Add(time.Duration(newToken.ExpiresIn) * time.Second).Format(time.RFC3339)
	//	app, err = r.Prisma.UpdateLinkedApp(prisma.LinkedAppUpdateParams{
	//		Data: prisma.LinkedAppUpdateInput{
	//			AccessToken:  &newToken.AccessToken,
	//			TokenExpiry:  &expStr,
	//			RefreshToken: nil,
	//		},
	//		Where: prisma.LinkedAppWhereUniqueInput{
	//			ID: &app.ID,
	//		},
	//	}).Exec(ctx)
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return app, nil
	//}
	//
	//newToken, err := r.Spotify.RefreshToken(*app.RefreshToken)

	//expStr := time.Now().Add(time.Duration(newToken.ExpiresIn) * time.Second).Format(time.RFC3339)
	//var refreshToken *string
	//if newToken.RefreshToken != "" {
	//	refreshToken = &newToken.RefreshToken
	//} else {
	//	refreshToken = app.RefreshToken
	//}
	//
	//app, err = r.Prisma.UpdateLinkedApp(prisma.LinkedAppUpdateParams{
	//	Data: prisma.LinkedAppUpdateInput{
	//		AccessToken:  &newToken.AccessToken,
	//		TokenExpiry:  &expStr,
	//		RefreshToken: refreshToken,
	//	},
	//	Where: prisma.LinkedAppWhereUniqueInput{
	//		ID: &app.ID,
	//	},
	//}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return app, nil
}

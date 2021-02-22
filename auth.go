package main

import (
	"github.com/gogearbox/gearbox"
	"github.com/google/uuid"
)

func authMiddleware(ctx gearbox.Context) {
	xAuthID, _ := uuid.Parse(ctx.Get("X-Auth-ID"))
	xAuthToken := ctx.Get("X-Auth-Token")

	ok, err := DB.Model(&Owner{}).Where("id=? AND token=?", xAuthID, xAuthToken).Exists()

	if ok != true || err != nil {
		ctx.Status(gearbox.StatusUnauthorized).SendJSON(&Response{"unauthorized"})
	} else {
		ctx.SetLocal("ownerId", xAuthID)
		ctx.Next()
	}
}

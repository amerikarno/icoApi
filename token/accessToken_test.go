package token

import (
	"testing"

	"github.com/amerikarno/icoApi/models"
)

func TestACExpire(t *testing.T) {

	ac := NewAccessClaims(&models.JwtUserModel{ID: "id", Email: "email", Permission: "permission", UserID: "userid", LoginStatus: "loginstatus"})
	expire := ac.IsExpired()
	if expire {
		t.Errorf("err: expect not expire but expired")
	}
	mock.forward(accessExp)
	expire = ac.IsExpired()
	if !expire {
		t.Errorf("err: expect expire but not")
	}

}

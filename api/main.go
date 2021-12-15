package api

import (
	"github.com/bufsnake/ldap-server/pkg/datas"
	"github.com/gin-gonic/gin"
)

type API struct {
	data *datas.Data
	sign string
}

func NewAPI(data *datas.Data, sign string) API {
	return API{data: data, sign: sign}
}

func (a *API) Verify(c *gin.Context) {
	sign := c.Query("sign")
	if sign != a.sign {
		c.Status(401)
		return
	}
	search := c.Query("search")
	if search == "" {
		c.Status(403)
		return
	}
	flag := a.data.VerifyData(search)
	if flag {
		c.Status(200)
		return
	}
	c.Status(404)
}

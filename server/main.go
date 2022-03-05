package server

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type Server struct {
	r *gin.Engine
}

func New() *Server {
	s := &Server{
		r: gin.Default(),
	}
	s.r.SetTrustedProxies(nil)
	s.r.GET("/:src", func(c *gin.Context) {
		src := c.Param("src")
		q, err := qrcode.New(src, qrcode.Medium)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
		q.DisableBorder = true
		buf := new(bytes.Buffer)
		q.Write(256, buf)
		c.DataFromReader(http.StatusOK, int64(buf.Len()), "image/png", buf, map[string]string{})
	})
	s.r.GET("/b64/:src", func(c *gin.Context) {
		src, err := base64.URLEncoding.DecodeString(c.Param("src"))
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprint(err))
			return
		}
		q, err := qrcode.New(string(src), qrcode.Medium)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
		q.DisableBorder = true
		buf := new(bytes.Buffer)
		q.Write(256, buf)
		c.DataFromReader(http.StatusOK, int64(buf.Len()), "image/png", buf, map[string]string{})
	})
	return s
}

func (s *Server) Run(host string) {
	s.r.Run(host)
}

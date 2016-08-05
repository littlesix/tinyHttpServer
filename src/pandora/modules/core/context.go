package core

import (
	"fmt"
	"mime"
	"net/http"
	"strings"
	"time"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Page           IPage
}

func (ctx *Context) WriteString(content string) {
	ctx.ResponseWriter.Write([]byte(content))
}

func (ctx *Context) Error(code int, message string) {
	http.Error(ctx.ResponseWriter, message, code)
}

func (ctx *Context) Redirect(urlStr string, args ...interface{}) {
	http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf(urlStr, args...), http.StatusSeeOther)
}

func (ctx *Context) ContentType(ext string) {
	ctype := mime.TypeByExtension(ext)

	if ctype != "" {
		ctx.ResponseWriter.Header().Set("Content-Type", ctype)
	}
}

func (ctx *Context) FormValue(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) SetHeader(hdr string, val string, unique bool) {
	if unique {
		ctx.ResponseWriter.Header().Set(hdr, val)
	} else {
		ctx.ResponseWriter.Header().Add(hdr, val)
	}
}

func (ctx *Context) SetCookie(name string, value string, age int64, path string, domain string, httpOnly bool) {
	var expires time.Time
	if age != 0 {
		expires = time.Unix(time.Now().Unix()+age, 0)
	}
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		Expires:  expires,
		HttpOnly: httpOnly,
	}
	http.SetCookie(ctx.ResponseWriter, cookie)
}

func webTime(t time.Time) string {
	ftime := t.Format(time.RFC1123)
	if strings.HasSuffix(ftime, "UTC") {
		ftime = ftime[0:len(ftime)-3] + "GMT"
	}
	return ftime
}

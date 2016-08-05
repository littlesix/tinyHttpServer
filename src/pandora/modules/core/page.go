package core

import (
	"net/http"
	"pandora/modules/logger"
	"pandora/vars"
	"strings"
	"text/template"
)

var tag = "page"

type WebPage interface {
}

type PageFunc func(p *Page) bool
type Page struct {
	Ctx                *Context
	Parent             interface{}
	TemplatePath       string
	isResponseByWriter bool
	// execute while page onload if return false break onload!;
	DoOnload        PageFunc
	//返回
	BeforeDoExecute PageFunc
}
type IPage interface {
	Prepare(ct *Context)
	Init(rule string)
	OnLoad()
	Execute()
}

func (p *Page) Init(rule string) {
	logger.D(tag, "init ", rule)
	p.TemplatePath = vars.TemplatePath + rule
}
func (p *Page) OnLoad() {
	if p.DoOnload != nil {
		p.DoOnload(p)
	}
}

func (p *Page) Prepare(ct *Context) {
	p.Ctx = ct
	p.Parent = p.Ctx.Page
}

func (p *Page) Get(key string) string {
	return p.Ctx.FormValue(key)
}
func (p *Page) SetCookie(key, value string) {
	var c = http.Cookie{Name: key, Value: value}
	http.SetCookie(p.Ctx.ResponseWriter, &c)
}
func (p *Page) GetCookie(key string) string {
	c, _ := p.Ctx.Request.Cookie(key)
	if c != nil {
		return c.Value
	} else {
		return ""
	}
}
func (p *Page) ResponseToClient(str string) {
	p.Ctx.WriteString(str)
}
func (p *Page) Redirect(url string) {
	var str string = `<html><meta http-equiv="refresh" content="0; url=` + url + `" /></html>`
	if strings.Contains(url, ".js") || strings.Contains(url, "callback=") {
		str = "window.lcation.href='" + url + "'"
	}

	//p.TemplatePath = "/blank.html"
	p.Ctx.WriteString(str)
	return
}
func (p *Page) Execute() {
	if p.BeforeDoExecute != nil && !p.BeforeDoExecute(p) {
		return
	}

	if p.isResponseByWriter {
		return
	}
	t, err := template.ParseFiles(p.TemplatePath)
	if err != nil {
		logger.E(tag, err)
	}

	er := t.Execute(p.Ctx.ResponseWriter, p.Parent)
	if er != nil {
		logger.E(tag, er)
	}
}

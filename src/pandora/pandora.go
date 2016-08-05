package pandora

import (
	"bufio"

	"fmt"
	"net/http"
	"os"
	"pandora/consts"
	"pandora/modules/conf"
	"pandora/modules/core"
	"pandora/modules/logger"
	"pandora/vars"

	"bytes"
)

type SiteServer struct {
	ListenAddr      string
	EnableGZip      bool
	TemplatePath    string
	WebRoot         string
	HomeUrl         string
	HomeIndexFiles  []string
}

var (
	siteServer     *SiteServer
	RunMode        string = "dev"
	basePathLength        = 0
)

func GetServer() *SiteServer {
	return siteServer
}

func init() {
	siteServer = &SiteServer{
		ListenAddr:   ":8080",
		EnableGZip:   true,
		TemplatePath: "./webroot",
		WebRoot:      "./www",
		HomeUrl:      "./",
		HomeIndexFiles: []string{"index.html","default.html","index.htm"},
	}

	if _, err := os.Stat("./logs"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./logs", os.ModeDir)
		}
	}

}

func loadConfig(cfgFilePath string) {
	baseCfg, err := conf.LoadConfig(cfgFilePath)
	vars.Conf = baseCfg
	if err != nil {
		logger.E("pandora", " load config err:", err)
	}
	var s *SiteServer = GetServer()

	s.ListenAddr = baseCfg.GetString("listen")
	s.WebRoot = baseCfg.GetString("web_root")
	s.TemplatePath = baseCfg.GetString("template_path")
	checkAvilable(s)
}

func checkAvilable(s *SiteServer) {
	if len(s.ListenAddr) < 3 {
		logger.E("pandora", "端口配置出错,请检查配置文件！")
	}
	if len(s.WebRoot) < 1 {
		logger.E("pandora", "网站根目录配置出错,请检查配置文件！")
	}
	if len(s.TemplatePath) < 1 {
		logger.E("pandora", "模板目录配置出错,请检查配置文件！")
	}
	logger.D(s.WebRoot)
}
func Init(cfgFilePath string) {
	loadConfig(cfgFilePath)
	//go initCtrl()
}

func Start() {
	defer func() {
		if err := recover(); err != nil {
			logger.E("pandora", "pandora error ", err)
		}
	}()
	GetServer().startHttp()
}

func (s *SiteServer) startHttp() {
	logger.W("pandora", "starting server, listening at", s.ListenAddr)

	if err := http.ListenAndServe(s.ListenAddr, newServerMux()); err != nil {
		logger.E("pandora:", " starting server error:", err.Error())
	}

}

var apiMap = make(map[string]http.Handler)

func newServerMux() *http.ServeMux {
	sm := http.NewServeMux()
	sm.HandleFunc("/", httpPageHandler)
	for k, v := range apiMap {
		sm.Handle(k, v)
	}
	return sm
}

func RouteApi(rule string, fitStuct http.Handler) {
	apiMap[rule] = fitStuct
}
func Route(rule string, fitStuct core.IPage, templatePath string) *core.Router {
	fitStuct.Init(templatePath)
	return core.Route(rule, fitStuct)
}

func httpPageHandler(w http.ResponseWriter, r *http.Request) {
	bPath := []byte(r.URL.Path)
	bRoot := []byte(GetServer().WebRoot)
	var routeMatched *core.RouteMatched

	if routeMatched = core.MatchRoot(bPath); routeMatched == nil {
		var fstr string = string(append(bRoot, bPath...))

		if(bytes.HasSuffix(bPath,consts.P_SLASH)){
			for _,indexF:=range GetServer().HomeIndexFiles{
				_,err:=os.Stat(fstr+indexF)
				if(err==nil){
					http.ServeFile(w, r, fstr+indexF)
					return;
				}
			}
			fstr="404.html"
		}

			http.ServeFile(w, r, fstr)

		return
	}

	var p core.IPage = routeMatched.Page
	ctx := buildContext(w, r, p)
	p.Prepare(ctx)
	p.OnLoad()
	p.Execute()

}

func buildContext(rw http.ResponseWriter, req *http.Request, p core.IPage) *core.Context {
	/*
		if req.Method == "POST" {
			req.ParseForm()
		}
	*/
	req.ParseForm()
	return &core.Context{ResponseWriter: rw, Request: req, Page: p}
}

func initCtrl() {
	/**
	  按ctrl+z ,回车终止程序
	*/
	var reader = bufio.NewReader(os.Stdin)
	for {

		keyCode, _ := reader.ReadByte()
		//fmt.Println("press key:", keyCode)
		//90 shift + z +enter
		//26 ctrl +z +enter
		if keyCode == 120 {
			fmt.Println("exiting")
			os.Exit(0)
		}
	}
}

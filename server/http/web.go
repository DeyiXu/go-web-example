package http

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	homeController "github.com/DeyiXu/go-web-example/controller/home"
	authController "github.com/DeyiXu/go-web-example/controller/auth"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/logger"
	"github.com/spf13/viper"
)

func loadTemplates(templatesDir string) multitemplate.Render {
	r := multitemplate.New()
	// 加载布局
	layouts, err := filepath.Glob(filepath.Join(templatesDir, "layouts/*.tmpl"))
	if err != nil {
		panic(err)
	}
	// 加载错误页面
	errors, err := filepath.Glob(filepath.Join(templatesDir, "errors/*.tmpl"))
	if err != nil {
		panic(err)
	}
	for _, errPage := range errors {
		tmplName := fmt.Sprintf("error_%s", filepath.Base(errPage))
		logger.Debugf("load error tmpl:%s", tmplName)
		r.AddFromFiles(tmplName, errPage)
	}
	// 页面文件夹
	pages, err := ioutil.ReadDir(filepath.Join(templatesDir, "pages"))
	if err != nil {
		panic(err)
	}
	for _, page := range pages {
		if !page.IsDir() {
			continue
		}
		for _, layout := range layouts {
			pageItems, err := filepath.Glob(filepath.Join(templatesDir, fmt.Sprintf("pages/%s/*.tmpl", page.Name())))
			if err != nil {
				panic(err)
			}
			files := []string{
				layout,
			}
			files = append(files, pageItems...)
			tmplName := fmt.Sprintf("%s_pages_%s", filepath.Base(layout), page.Name())
			logger.Debugf("load page tmpl:%s", tmplName)
			r.AddFromFiles(tmplName, files...)
		}
	}
	// 加载单页面
	singles, err := filepath.Glob(filepath.Join(templatesDir, "singles/*.tmpl"))
	if err != nil {
		panic(err)
	}
	for _, singlePage := range singles {
		tmplName := fmt.Sprintf("singles_%s", filepath.Base(singlePage))
		logger.Debugf("load single tmpl:%s", tmplName)
		r.AddFromFiles(tmplName, singlePage)
	}
	return r
}

func setWeb(engine *gin.Engine) {
	engine.HTMLRender = loadTemplates(viper.GetString("web.templates_dir"))
	// session
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))
	// gizp
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	// file server
	engine.Static("/assets", viper.GetString("web.assets_dir"))
}

func setWebRouter(engine *gin.RouterGroup) {
	engine.GET("/",ngin.WebControllerFunc(homeController.Index,"index"))
	engine.GET("/login", ngin.WebControllerFunc(authController.Login, "login"))
	engine.GET("/register", ngin.WebControllerFunc(authController.Register, "register"))
}

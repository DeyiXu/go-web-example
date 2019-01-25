package http

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/DeyiXu/go-web-example/controller/web"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
		r.AddFromFilesFuncs(tmplName, loadFuncMap(), errPage)
	}

	// 加载局部页面
	partials, err := filepath.Glob(filepath.Join(templatesDir, "partials/*.tmpl"))
	if err != nil {
		panic(err)
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
			files = append(files, partials...)
			files = append(files, pageItems...)
			tmplName := fmt.Sprintf("%s_pages_%s", filepath.Base(layout), page.Name())
			logger.Debugf("load page tmpl:%s", tmplName)
			r.AddFromFilesFuncs(tmplName, loadFuncMap(), files...)
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
		r.AddFromFilesFuncs(tmplName, loadFuncMap(), singlePage)
	}
	return r
}

func loadFuncMap() template.FuncMap {
	return template.FuncMap{
		"getMenuData":       web.Auth.GetMenuData,
		"getNavigationData": web.Auth.GetNavigationData,
	}
}

func setWeb(engine *gin.Engine) {
	// 404 page
	engine.NoRoute(web.Error.Error404)

	engine.HTMLRender = loadTemplates(viper.GetString("web.templates_dir"))
	// session
	gob.Register(gin.H{})
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))
	// gizp
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	// file server
	engine.Static("/assets", viper.GetString("web.assets_dir"))
}

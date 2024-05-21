package templateServiceModule

import (
	"bytes"
	"fmt"
	"github.com/Xi-Yuer/cms/dto"
	templateResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/template"
	"github.com/Xi-Yuer/cms/utils"
	"sync"
	"text/template"
)

var TemplateService = &templateService{}

type templateService struct{}

func (t *templateService) CreateTemplate(params *dto.CreateTemplateRequestParams) (*dto.CreateTemplateResponse, error) {
	var wg sync.WaitGroup
	errChan := make(chan error, 3)
	results := make(map[string]string)
	var mu sync.Mutex

	templates := map[string]string{
		"controller":         "template/server/controller/controller.tmpl",
		"service":            "template/server/service/service.tmpl",
		"repository":         "template/server/repository/repository.tmpl",
		"route":              "template/server/route/route.tmpl",
		"dto":                "template/server/dto/dto.tmpl",
		"webReactSearchForm": "template/web/react/hook/searchForm.tmpl",
		"webReactTableHook":  "template/web/react/hook/crudHook.tmpl",
		"webReactTable":      "template/web/react/pages/crud.tmpl",
	}

	wg.Add(len(templates))

	for key, path := range templates {
		go func(key, path string) {
			defer wg.Done()
			result, err := parseTemplate(path, params)
			if err != nil {
				errChan <- err
				return
			}

			mu.Lock()
			results[key] = result
			mu.Unlock()
		}(key, path)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return &dto.CreateTemplateResponse{
		Server: templateResponsiesModules.Server{
			ControllerFile: templateResponsiesModules.Code{
				Code: results["controller"],
				Lang: "go",
			},
			ServiceFile: templateResponsiesModules.Code{
				Code: results["service"],
				Lang: "go",
			},
			RepositoryFile: templateResponsiesModules.Code{
				Code: results["repository"],
				Lang: "go",
			},
			RouteFile: templateResponsiesModules.Code{
				Code: results["route"],
				Lang: "go",
			},
			DTOFile: templateResponsiesModules.Code{
				Code: results["dto"],
				Lang: "go",
			},
		},
		Web: templateResponsiesModules.Web{
			React: templateResponsiesModules.React{
				SearchForm: templateResponsiesModules.Code{
					Code: results["webReactSearchForm"],
					Lang: "typescript",
				},
				TableHook: templateResponsiesModules.Code{
					Code: results["webReactTableHook"],
					Lang: "typescript",
				},
				Table: templateResponsiesModules.Code{
					Code: results["webReactTable"],
					Lang: "typescript",
				},
			},
		},
	}, nil
}

func (t *templateService) DownloadTemplateZip(params *dto.DownloadTemplateRequestParams) ([]byte, error) {
	// 定义文件路径和内容的映射
	Controller := fmt.Sprintf("%vTemplate/server/controller/%v.go", params.TableName, params.TableName)
	Service := fmt.Sprintf("%vTemplate/server/service/%v.go", params.TableName, params.TableName)
	Repository := fmt.Sprintf("%vTemplate/server/repository/%v.go", params.TableName, params.TableName)
	Route := fmt.Sprintf("%vTemplate/server/route/%v.go", params.TableName, params.TableName)
	DTO := fmt.Sprintf("%vTemplate/server/dto/%v.go", params.TableName, params.TableName)
	SearchForm := fmt.Sprintf("%vTemplate/web/react/hook/searchForm.tsx", params.TableName)
	TableHook := fmt.Sprintf("%vTemplate/web/react/hook/%vtableHook.tsx", params.TableName, params.TableName)
	Table := fmt.Sprintf("%vTemplate/web/react/pages/%vTablePage.tsx", params.TableName, params.TableName)
	files := map[string]string{
		Controller: params.Controller,
		Service:    params.Service,
		Repository: params.Repository,
		Route:      params.Route,
		DTO:        params.DTO,
		SearchForm: params.SearchForm,
		Table:      params.Table,
		TableHook:  params.TableHook,
	}

	return utils.CreateFilesAndZip(files)
}

func parseTemplate(path string, params any) (string, error) {
	var parseTemplate bytes.Buffer
	parseControllerExecute, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	if err := parseControllerExecute.Execute(&parseTemplate, params); err != nil {
		return "", err
	}
	return parseTemplate.String(), nil
}

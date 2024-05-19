package templateServiceModule

import (
	"bytes"
	"github.com/Xi-Yuer/cms/dto"
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
		"controller": "template/controller/controller.tmpl",
		"service":    "template/service/service.tmpl",
		"repository": "template/repository/repository.tmpl",
		"dto":        "template/dto/dto.tmpl",
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
		Controller: results["controller"],
		Service:    results["service"],
		Repository: results["repository"],
		DTO:        results["dto"],
	}, nil
}

func parseTemplate(path string, params any) (string, error) {
	var parseTemplate bytes.Buffer
	parseControllerExecute, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	if err := parseControllerExecute.Execute(&parseTemplate, params); err != nil {
		return "", err
	}
	return parseTemplate.String(), nil
}

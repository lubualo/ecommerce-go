package routers

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/handlers"
)

func Router(request events.APIGatewayV2HTTPRequest, urlPrefix string) (*events.APIGatewayProxyResponse) {
	path := strings.Replace(request.RawPath, urlPrefix, "", -1)
	method := request.RequestContext.HTTP.Method
	fmt.Println("Processing " + path + " > " + method)
	firstSegment := getFirstPathSegment(path)

	entityHandler, _ := handlers.CreateHandler(firstSegment)
	switch method {
	case Get:
		return entityHandler.Get(request)
	case Post:
		return entityHandler.Post(request)
	case Put:
		return entityHandler.Put(request)
	case Delete:
		return entityHandler.Delete(request)
	}



	// firstSegment := getFirstPathSegment(path)
	// entityRouter, err := CreateRouter(firstSegment)
	// if err != nil {
	// 	return 400, "unable to create router: " + err.Error()
	// }
	// switch method {
	// case Get:
	// 	return entityRouter.Get(request)
	// case Post:
	// 	return entityRouter.Post(request)
	// case Put:
	// 	return entityRouter.Put(request)
	// case Delete:
	// 	return entityRouter.Delete(request)
	// default:
	// 	return 405, "method not allowed"
	// }
}

func getFirstPathSegment(path string) string {
	// Remove leading/trailing slashes
	trimmed := strings.Trim(path, "/")
	segments := strings.Split(trimmed, "/")
	if len(segments) > 0 && segments[0] != "" {
		return segments[0]
	}
	return ""
}

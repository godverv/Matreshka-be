package docs

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"

	swaggerui "github.com/Red-Sock/go-swagger-ui"
)

//go:embed all:grpc
var swaggers embed.FS

const (
	BasePath    = "/docs/"
	swaggerPath = BasePath + "swaggers/"
)

func Swagger() (p string, handler http.HandlerFunc) {
	mux := http.NewServeMux()

	mux.Handle(BasePath, swaggerui.NewHandler(
		swaggerui.WithBasePath(BasePath),
		swaggerui.WithHTMLTitle("Swagger"),
		swaggerui.WithSpecURLs("Api/grpc/matreshkaBeApi",
			[]swaggerui.SpecURL{
				{
					Name: "Api/grpc/matreshkaBeApi",
					URL:  path.Join(swaggerPath, "api/grpc/matreshka-be_api.swagger.json"),
				},
			}),
		swaggerui.WithShowExtensions(true),
	))

	stripped, err := fs.Sub(swaggers, "grpc")
	if err != nil {
		log.Fatal(err)
	}

	ffs := http.StripPrefix(swaggerPath, http.FileServer(http.FS(stripped)))
	mux.Handle(swaggerPath, ffs)

	return BasePath, mux.ServeHTTP
}

package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"lai.com/GraphQL_Server/graph"
	"lai.com/GraphQL_Server/graph/generated"
	"lai.com/GraphQL_Server/service"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultPort = "8080"

func init() {
	//初始化数据表
	service.InitTables()
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	gin.SetMode("debug")
	r := gin.New()
	r.Any("/", FromStd(playground.Handler("GraphQL playground", "/query")))
	r.Any("/query", FromStd(srv.ServeHTTP))

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 30,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

}

// 上下文密钥类型
type CtxKeyT struct{}

//保存在request.context中的gin ctx的密钥
var GinCtxKey CtxKeyT

// 将标准处理程序转换为gin.Handler，并嵌入gin上下文
func FromStd(handler http.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r2 := ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), GinCtxKey, ctx))
		handler(ctx.Writer, r2)
	}
}

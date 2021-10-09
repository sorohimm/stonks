package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"

	"net/http"
	"stonks/internal/interfaces"

	"github.com/go-playground/validator/v10"
)

type NewsControllers struct {
	Log         *zap.SugaredLogger
	NewsService interfaces.INewsService
}

func CurrentUTCDate() string {
	t := time.Now()
	year := t.UTC().Year()
	month := t.UTC().Month()
	var m = int(month)
	day := t.UTC().Day()
	date := fmt.Sprintf("%s-%s-%s", strconv.Itoa(year), strconv.Itoa(m), strconv.Itoa(day))

	return date
}

func (c *NewsControllers) GetNews(ctx *gin.Context) {

	validate := validator.New()
	var err error

	parameter := ctx.Request.URL.Query()
	if parameter.Has("q") {
		err = validate.Var(parameter["q"], "required")
		if err != nil {
			c.Log.Warn("invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "The request must contain the 'c' field "})
			return
		}
	} else {
		c.Log.Fatal("invalid request")
	}

	if parameter.Has("sort_by") {
		switch parameter["sort_by"][0] {
		case "relevancy", "date", "rank":
		default:
			c.Log.Warn("invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'sort' parameter"})
			return
		}
	}

	if parameter.Has("page") {
		err = validate.Var(parameter["page"], "omitempty,numeric,min=1,max=100")
		if err != nil {
			c.Log.Warn("invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'page' parameter"})
			return
		}
	}


	if parameter.Has("page_size") {
		err = validate.Var(parameter["page_size"], "omitempty,numeric,min=1,max=100")
		if err != nil {
			c.Log.Warn("invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'page' parameter"})
			return
		}
	}

	if parameter.Has("from") {
		err = validate.Var(parameter["from"], "datetime")
		if err == nil {
			c.Log.Warn("invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'from' parameter"})
			return
		}
	}


	if parameter.Has("to") {
		err = validate.Var(parameter["to"], "datetime")
		if err == nil {
			c.Log.Warn("invalid request")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'from' parameter"})
			return
		}
	}

	resp, err :=c.NewsService.GetNews(parameter)
	if err != nil {
		c.Log.Warn("unknown error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid value of the 'to' parameter"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

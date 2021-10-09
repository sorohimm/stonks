package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"regexp"
	"stonks/internal/config"
	"strconv"
	"time"

	"net/http"
	"stonks/internal/interfaces"

	"github.com/go-playground/validator/v10"
)

type NewsControllers struct {
	Log         *zap.SugaredLogger
	NewsService interfaces.INewsService
	Config *config.Config
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

	company := ctx.Request.URL.Query().Get("c")

	err := validate.Var(company, "required")
	if err != nil {
		c.Log.Panic("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "The request must contain the 'c' field "})
		return
	}

	sort := ctx.DefaultQuery("sort", "relevancy")

	switch sort {
	case "relevancy", "date", "rank":
	default:
		c.Log.Panic("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'sort' parameter"})
		return
	}

	page := ctx.DefaultQuery("page", "50")
	err = validate.Var(page, "omitempty,numeric,min=1,max=100")
	if err != nil {
		c.Log.Panic("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'page' parameter"})
		return
	}

	pageSize := ctx.DefaultQuery("page_size", "15")
	err = validate.Var(pageSize, "omitempty,numeric,min=1,max=100")
	if err != nil {
		c.Log.Panic("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'page' parameter"})
		return
	}

	to := ctx.DefaultQuery("to", "2021-10-9")
	re := regexp.MustCompile("((19|20)\\d\\d)/(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])")
	if re == nil {
		c.Log.Panic("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'to' parameter"})
		return
	}

	from := ctx.DefaultQuery("from", "2021-10-8")
	re = regexp.MustCompile("((19|20)\\d\\d)/(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])")
	if re == nil {
		c.Log.Panic("invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid value of the 'from' parameter"})
		return
	}

	body := "https://api.newscatcherapi.com/v2/search?"
	query := fmt.Sprintf("%sq=%s&sort_by=%s&page=%s&page_size=%s&to=%s&from=%s", body, company, sort, page, pageSize, to, from)

	resp, err :=c.NewsService.GetNews(query)
	if err != nil {
		c.Log.Panic("invalid request")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid value of the 'to' parameter"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

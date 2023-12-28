package ginger

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type IEngine interface {
	GetGin() *gin.Engine
	SetGin(*gin.Engine)
	GetCron() *cron.Cron
	SetCron(*cron.Cron)
	Run(...string)
}

type Engine struct {
	Gin  *gin.Engine
	Cron *cron.Cron
}

func (e *Engine) GetGin() *gin.Engine {
	return e.Gin
}

func (e *Engine) SetGin(gin *gin.Engine) {
	e.Gin = gin
}

func (e *Engine) GetCron() *cron.Cron {
	return e.Cron
}

func (e *Engine) SetCron(cron *cron.Cron) {
	e.Cron = cron
}

func (e *Engine) Run(addr ...string) {
	e.Cron.Start()
	e.Gin.Run(addr...)
}

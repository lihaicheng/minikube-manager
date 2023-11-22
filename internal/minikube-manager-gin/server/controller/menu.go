package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lihaicheng/minikube-manager/internal/minikube-manager-gin/store/dao"
	"go.uber.org/zap"
	"net/http"
)

type Menu struct {
	HomeInfo Info   `json:"homeInfo"`
	LogoInfo Info   `json:"logoInfo"`
	MenuInfo []Info `json:"menuInfo"`
}

type Info struct {
	Title  string `json:"title,omitempty"`
	Image  string `json:"image,omitempty"`
	Href   string `json:"href"`
	Target string `json:"target,omitempty"`
	Child  []Info `json:"child,omitempty"`
}

func DefaultInitMenuHandler(c *gin.Context) {
	// JSON文本作为字符串
	filter := make(map[string]interface{}, 0)
	filter["flag"] = "test"
	menu, _ := dao.Menu().Get(filter)
	zap.L().Info("DefaultInitMenuHandler: menu value is " + menu.Value)
	// 使用c.Data将JSON文本发送到前端
	c.Data(http.StatusOK, "application/json", []byte(menu.Value))
}

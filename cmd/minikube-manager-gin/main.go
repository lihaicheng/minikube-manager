package main

import (
	"flag"
	"fmt"
	svc "github.com/kardianos/service"
	"github.com/lihaicheng/minikube-manager/internal/minikube-manager-gin/pkg/config"
	"github.com/lihaicheng/minikube-manager/internal/minikube-manager-gin/pkg/logger"
	"github.com/lihaicheng/minikube-manager/internal/minikube-manager-gin/server"
	"github.com/lihaicheng/minikube-manager/internal/minikube-manager-gin/store/mysql"
	"go.uber.org/zap"
	"os"
)

const (
	linuxDefaultConfigFile = "configs/minikube-manager.conf"
)

var (
	versionNumber string
	buildTime     string
	configFile    string
	version       bool
)

type serviceWrapper struct{}

func (w *serviceWrapper) Start(s svc.Service) error {
	// Start should not block. Do the actual work async.
	go w.run()
	return nil
}

func (w *serviceWrapper) run() {

	if err := config.InitConfig(configFile); err != nil {
		os.Exit(1)
	}

	if err := logger.InitLogger(config.Config); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// Database initialization
	zap.L().Info("init database.")
	err := mysql.InitDB(config.Config)
	if err != nil {
		zap.L().Error("init database failed", zap.Error(err))
	}

	_, err = server.New()
	if err != nil {
		zap.L().Error("init server failed", zap.Error(err))
	}

}
func (w *serviceWrapper) Stop(s svc.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func init() {
	flag.StringVar(&configFile, "c", linuxDefaultConfigFile, "config file path.")

	flag.BoolVar(&version, "v", false, "print version number.")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "\nusage: %s [install|uninstall|start|stop|restart]\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
}

func main() {
	options := make(svc.KeyValue)
	options["Type"] = "forking"
	svcConfig := &svc.Config{
		Name:        "github.com/lihaicheng/minikube-manager",
		DisplayName: "github.com/lihaicheng/minikube-manager",
		Description: "github.com/lihaicheng/minikube-manager",
		Arguments:   []string{"-c", linuxDefaultConfigFile},
		Option:      options,
	}
	wrapper := &serviceWrapper{}
	s, err := svc.New(wrapper, svcConfig)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}

	if len(os.Args) > 1 {
		for _, v := range svc.ControlAction {
			if os.Args[1] == v {
				err = svc.Control(s, os.Args[1])
				if err != nil {
					fmt.Println(err)
				}
				return
			}
		}
	}
	err = s.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}

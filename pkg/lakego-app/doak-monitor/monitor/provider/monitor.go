package provider

import (
    "github.com/deatil/lakego-doak/lakego/router"
    "github.com/deatil/lakego-doak/lakego/provider"

    "github.com/deatil/lakego-doak-admin/admin/support/route"

    monitorRouter "github.com/deatil/lakego-doak-monitor/monitor/route"
)

/**
 * 服务提供者
 *
 * @create 2022-7-3
 * @author deatil
 */
type Monitor struct {
    provider.ServiceProvider
}

// 引导
func (this *Monitor) Boot() {
    // 路由
    this.loadRoute()
}

/**
 * 导入路由
 */
func (this *Monitor) loadRoute() {
    // 后台路由
    route.AddRoute(func(engine *router.RouterGroup) {
        monitorRouter.Route(engine)
    })
}


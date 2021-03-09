/**
 * @file main store
 */

import Vue from 'vue'
import Vuex from 'vuex'
import cookie from 'cookie'

import http from '@open/api'
import { unifyObjectStyle, json2Query } from '@open/common/util'

import depot from '@open/store/modules/depot'
import mc from '@open/store/modules/mc'
import cluster from '@open/store/modules/cluster'
import resource from '@open/store/modules/resource'
import app from '@open/store/modules/app'
import variable from '@open/store/modules/variable'
import configuration from '@open/store/modules/configuration'
import templateset from '@open/store/modules/templateset'
import network from '@open/store/modules/network'
import mesosTemplate from '@open/store/modules/mesos-template'
import k8sTemplate from '@open/store/modules/k8s-template'
import helm from '@open/store/modules/helm'

import menuConfig from './menu-config'

Vue.use(Vuex)

// cookie 中 zh-cn / en
let lang = cookie.parse(document.cookie).blueking_language || 'zh-cn'
if (['zh-CN', 'zh-cn', 'cn', 'zhCN', 'zhcn'].indexOf(lang) > -1) {
    lang = 'zh-CN'
} else {
    lang = 'en-US'
}

const { menuList, k8sMenuList } = menuConfig(lang)

const store = new Vuex.Store({
    // 模块
    modules: {
        depot,
        mc,
        cluster,
        resource,
        app,
        variable,
        configuration,
        templateset,
        network,
        mesosTemplate,
        k8sTemplate,
        helm
    },
    // 公共 store
    state: {
        curProject: null,
        mainContentLoading: false,
        // 系统当前登录用户
        user: {},
        // 左侧导航
        sideMenu: {
            // 在线的 project
            onlineProjectList: [],
            // 左侧导航 menu 集合
            menuList: menuList,
            k8sMenuList: k8sMenuList
        },

        // 当前语言环境
        lang: lang,
        isEn: lang === 'en-US',

        // 是否允许路由跳转
        allowRouterChange: true,

        crdInstanceList: []
    },
    // 公共 getters
    getters: {
        mainContentLoading: state => state.mainContentLoading,
        user: state => state.user,
        lang: state => state.lang
    },
    // 公共 mutations
    mutations: {
        /**
         * 设置内容区的 loading 是否显示
         *
         * @param {Object} state store state
         * @param {boolean} loading 是否显示 loading
         */
        setMainContentLoading (state, loading) {
            state.mainContentLoading = loading
        },

        /**
         * 更新当前用户 user
         *
         * @param {Object} state store state
         * @param {Object} user user 对象
         */
        updateUser (state, user) {
            state.user = Object.assign({}, user)
        },

        /**
         * 更改当前项目信息
         *
         * @param {Object} state store state
         * @param {String} projectId
         */
        updateCurProject (state, projectCode) {
            const project = state.sideMenu.onlineProjectList.find(project => project.project_code === projectCode)
            if (project) {
                state.curProject = Object.assign({}, project)
                state.sideMenu.k8sMenuList = k8sMenuList
                state.sideMenu.menuList = menuList
            }
        },

        /**
         * 修改 state.allowRouterChange
         *
         * @param {Object} state store state
         * @param {boolean} val 值
         */
        updateAllowRouterChange (state, val) {
            state.allowRouterChange = val
        },

        /**
         * 更新 store 中的 onlineProjectList
         *
         * @param {Object} state store state
         * @param {list} list 项目列表
         */
        forceUpdateOnlineProjectList (state, list) {
            state.sideMenu.onlineProjectList.splice(0, state.sideMenu.onlineProjectList.length, ...list)
        },

        /**
         * 更新 store 中的 menuList
         *
         * @param {Object} state store state
         * @param {list} list menu 列表
         */
        forceUpdateMenuList (state, list) {
            if (state.curProject && state.curProject.kind === 1) {
                state.sideMenu.k8sMenuList.splice(0, state.sideMenu.k8sMenuList.length, ...list)
            } else {
                state.sideMenu.menuList.splice(0, state.sideMenu.menuList.length, ...list)
            }
        },

        /**
         * 更新 store 中的 menuList
         *
         * @param {Object} state store state
         * @param {list} list menu 列表
         */
        forceUpdateDevOpsMenuList (state, list) {
            state.sideMenu.devOpsMenuList.splice(0, state.sideMenu.devOpsMenuList.length, ...list)
        },

        /**
         * 更新crdInstanceList
         * @param {Object} state store state
         * @param {Object} data data
         */
        updateCrdInstanceList (state, data) {
            state.crdInstanceList = data
        }
    },
    actions: {
        /**
         * 获取用户信息
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        userInfo (context, params, config = {}) {
            // return http.get(`/app/index?invoke=userInfo`, {}, config)
            return http.get(`${DEVOPS_BCS_API_URL}/api/user/`, params, config).then(response => {
                const userData = response.data || {}
                context.commit('updateUser', userData)
                return userData
            })
        },

        /**
         * 获取项目列表
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        getProjectList (context, params, config = {}) {
            return http.get(DEVOPS_BCS_API_URL + '/api/projects/', params, config).then(response => {
                let list = []
                const online = []
                const offline = []
                const adminList = []
                if (response.data && response.data.length) {
                    list = response.data
                }
                list.forEach(item => {
                    if (item.is_offlined) {
                        offline.push(item)
                    } else {
                        if (item.permissions && item.permissions['modify:project:btn'] === 0) {
                            adminList.push(item)
                        }
                        // 通过审批
                        if (item.approval_status === 2) {
                            online.push(item)
                        }
                    }
                })
                context.commit('forceUpdateOnlineProjectList', online)
                return list
            })
        },

        /**
         * 根据项目 id 查询项目的权限
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        getProjectPerm (context, { projectCode }, config = {}) {
            return http.get(`${DEVOPS_BCS_API_URL}/api/projects/${projectCode}/`)
        },

        /**
         * 根据 pathName 来判断 menuList 中的哪一个 menu 应该被选中
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         *
         * @return {Promise} promise 对象
         */
        updateMenuListSelected (context, { pathName, idx, projectType }) {
            return new Promise((resolve, reject) => {
                const list = []
                const tmp = []
                let invokeStr = ''
                switch (idx) {
                    case 'devops':
                        tmp.splice(0, 0, ...context.state.sideMenu.devOpsMenuList)
                        invokeStr = 'forceUpdateDevOpsMenuList'
                        break
                    case 'bcs':
                        if ((context.state.curProject && context.state.curProject.kind === 1) || projectType === 'k8s') {
                            tmp.splice(0, 0, ...context.state.sideMenu.k8sMenuList)
                        } else {
                            tmp.splice(0, 0, ...context.state.sideMenu.menuList)
                        }
                        invokeStr = 'forceUpdateMenuList'
                        break
                    default:
                }

                // 清掉 menuList 里的选中
                tmp.forEach(m => {
                    m.isSelected = false
                    m.isOpen = false
                    if (m.children) {
                        m.isChildSelected = false
                        m.children.forEach(childItem => {
                            childItem.isSelected = false
                        })
                    }
                })
                list.splice(0, 0, ...tmp)

                let continueLoop = true

                const len = list.length
                for (let i = len - 1; i >= 0; i--) {
                    if (!continueLoop) {
                        break
                    }
                    const menu = list[i]
                    if ((menu.pathName || []).indexOf(pathName) > -1) {
                        // clearMenuListSelected(list)
                        menu.isSelected = true
                        continueLoop = false
                        break
                    }
                    if (menu.children) {
                        const childrenLen = menu.children.length
                        for (let j = childrenLen - 1; j >= 0; j--) {
                            if ((menu.children[j].pathName || []).indexOf(pathName) > -1) {
                                // clearMenuListSelected(list)
                                menu.isOpen = true
                                menu.isChildSelected = true
                                menu.children[j].isSelected = true
                                continueLoop = false
                                break
                            }
                        }
                    }
                }

                context.commit(invokeStr, list)
            })
        },

        /**
         * 获取资源权限
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} 参数 对象
         */
        getResourcePermissions (context, params, config = {}) {
            return http.post(`${DEVOPS_BCS_API_URL}/api/perm/verify/`, params, config)
        },

        /**
         * 获取多个资源权限
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} 参数 对象
         */
        getMultiResourcePermissions (context, params, config = {}) {
            return http.post(`${DEVOPS_BCS_API_URL}/api/perm/multi/verify/`, params, config)
        },

        /**
         * 获取关联 CC 的列表
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        getCCList (context, params = {}, config = {}) {
            return http.get(`${DEVOPS_BCS_API_URL}/api/cc/?${json2Query(params)}`, params, config)
        },

        /**
         * 停用/启用屏蔽
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        editProject (context, params, config = {}) {
            const projectId = params.project_id
            return http.put(`${DEVOPS_BCS_API_URL}/api/projects/${projectId}/`, params, config)
        },

        /**
         * 获取项目信息
         *
         * @param {Object} context store 上下文对象
         * @param {Object} params 请求参数
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        getProject (context, params, config = {}) {
            const projectId = params.projectId
            return http.get(`${DEVOPS_BCS_API_URL}/api/projects/${projectId}/`, params, config)
        },

        /**
         * 项目启用日志采集功能
         *
         * @param {Object} context store 上下文对象
         * @param {string} projectId 项目 id
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        enableLogPlans (context, projectId, config = {}) {
            return http.post(
                `${DEVOPS_BCS_API_URL}/api/datalog/projects/${projectId}/log_plans/`,
                {},
                config
            )
        },

        /**
         * 获取项目日志采集信息
         *
         * @param {Object} context store 上下文对象
         * @param {string} projectId 项目 id
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        getLogPlans (context, projectId, config = {}) {
            return http.get(
                `${DEVOPS_BCS_API_URL}/api/datalog/projects/${projectId}/log_plans/`,
                {},
                config
            )
        },

        /**
         * 查询crd列表 (新)
         *
         * @param {Object} context store 上下文对象
         * @param {Object} projectId, clusterId, crdKind
         * @param {Object} config 请求的配置
         *
         * @return {Promise} promise 对象
         */
        getBcsCrdsList (context, { projectId, params = {} }, config = {}) {
            context.commit('updateCrdInstanceList', [])
            const url = `${DEVOPS_BCS_API_URL}/api/bcs_crd/projects/${projectId}/crds/?${json2Query(params)}`
            return http.get(url, {}, config).then(res => {
                context.commit('updateCrdInstanceList', res.data)
                return res
            })
        }
    }
})

/**
 * hack vuex dispatch, add third parameter `config` to the dispatch method
 *
 * 需要对单独的请求做配置的话，无论是 get 还是 post，store.dispatch 都需要三个参数，例如：
 * store.dispatch('example/btn1', {btn: 'btn1'}, {fromCache: true})
 * 其中第二个参数指的是请求本身的参数，第三个参数指的是请求的配置，如果请求本身没有参数，那么
 * 第二个参数也必须占位，store.dispatch('example/btn1', {}, {fromCache: true})
 * 在 store 中需要如下写法：
 * btn1 ({commit, state, dispatch}, params, config) {
 *     return http.get(`/app/index?invoke=btn1`, params, config)
 * }
 *
 * @param {Object|string} _type vuex type
 * @param {Object} _payload vuex payload
 * @param {Object} config config 参数，主要指 http 的参数，详见 src/api/index initConfig
 *
 * @return {Promise} 执行请求的 promise
 */
store.dispatch = function (_type, _payload, config = {}) {
    const { type, payload } = unifyObjectStyle(_type, _payload)
    const action = { type, payload, config }
    const entry = store._actions[type]

    if (!entry) {
        if (NODE_ENV !== 'production') {
            console.error(`[vuex] unknown action type: ${type}`)
        }
        return
    }

    store._actionSubscribers.forEach(sub => {
        return sub(action, store.state)
    })

    return entry.length > 1
        ? Promise.all(entry.map(handler => handler(payload, config)))
        : entry[0](payload, config)
}

export default store

<template>
    <div class="biz-content">
        <div class="biz-top-bar">
            <div class="biz-tpl-title">
                {{$t('Helm Chart仓库')}}
            </div>
            <bk-guide>
                <a href="javascript:void(0);" class="bk-text-button" @click.stop.prevent="showGuide">{{$t('如何推送Helm Chart到项目仓库？')}}</a>
            </bk-guide>
        </div>

        <guide ref="clusterGuide"></guide>

        <bk-dialog
            :is-show.sync="helmDialog.isShow"
            :width="500"
            :has-footer="false"
            :title="$t('项目Chart仓库配置信息')"
            @cancel="hideHelmDialog">
            <div slot="content">
                <div class="helm-repos-detail" v-if="reposData">
                    <div class="repos-item" v-for="repo in reposData.privateRepos" :key="repo.url">
                        <div class="wrapper mb10">
                            <h2 class="label">{{$t('项目Chart仓库地址')}}：</h2>
                            <p class="url">
                                {{repo.url}}
                                <bk-tooltip placement="top" :content="$t('复制')">
                                    <span :data-clipboard-text="repo.url" class="copy-btn bk-text-button bk-icon icon-clipboard ml5" style="color: #999;"></span>
                                </bk-tooltip>
                            </p>
                        </div>
                        <div class="auth" v-for="auth in repo.auths" :key="auth.credentials_decoded.username">
                            <div>username：{{auth.credentials_decoded.username}}</div>
                            <div>password：
                                <template v-if="repo.isPasswordShow">{{auth.credentials_decoded.password}}</template>
                                <template v-else>************</template>
                                <bk-tooltip placement="top" :content="repo.isPasswordShow ? $t('隐藏') : $t('查看')">
                                    <span :class="['bk-text-button bk-icon',{ 'icon-eye': !repo.isPasswordShow, 'icon-eye-slash': repo.isPasswordShow }]" @click="togglePassword(repo)" style="color: #999;"></span>
                                </bk-tooltip>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </bk-dialog>

        <div class="biz-content-wrapper biz-tpl-wrapper" style="padding: 0; margin: 0;" v-bkloading="{ isLoading: showLoading, opacity: 0.1 }">
            <template v-if="!showLoading">
                <div class="biz-panel-header" style="padding: 20px;">
                    <div class="left">
                        <bk-button type="primary" @click="syncHelmTpl" :loading="isTplSynLoading">{{$t('同步仓库')}}</bk-button>
                        <span class="biz-tip f13 ml5">{{$t('同步仓库中的Helm Chart')}}</span>
                        <a class="bk-text-button f13 ml10" href="javascript:void(0);" @click="getHelmDeops">{{$t('查看项目Chart仓库配置信息')}}</a>
                    </div>
                    <div class="right">
                        <div class="biz-search-input" style="width: 300px;">
                            <input type="text" class="bk-form-input" :placeholder="$t('输入关键字，按Enter搜索')" v-model="searchKeyword" @keyup.enter="search">
                            <a href="javascript:void(0)" class="biz-search-btn" v-if="!searchKeyword">
                                <i class="bk-icon icon-search" style="color: #c3cdd7;"></i>
                            </a>
                            <a href="javascript:void(0)" class="biz-search-btn" v-else @click.stop.prevent="clearSearch">
                                <i class="bk-icon icon-close-circle-shape"></i>
                            </a>
                        </div>
                    </div>
                </div>

                <app-exception
                    v-if="exceptionCode && !showLoading"
                    :type="exceptionCode.code"
                    :text="exceptionCode.msg">
                </app-exception>

                <template>
                    <svg style="display: none;">
                        <title>{{$t('模板集默认图标')}}</title>
                        <symbol id="biz-set-icon" viewBox="0 0 60 60">
                            <g id="图层_6">
                                <g id="图层_32_1_">
                                    <path class="st0" d="M12,8v4H8c-1.1,0-2,0.9-2,2v42c0,1.1,0.9,2,2,2h42c1.1,0,2-0.9,2-2v-4h4c1.1,0,2-0.9,2-2V8c0-1.1-0.9-2-2-2
                                        H14C12.9,6,12,6.9,12,8z M48,48v4v2H10V16h2h4h32V48z M54,48h-2V14c0-1.1-0.9-2-2-2H16v-2h38V48z" />
                                </g>
                                <path class="st1" d="M45.7,33.7h-1.8l-3.4-8.3l1.3-1.3c0.5-0.5,0.5-1.3,0-1.8l0,0c-0.5-0.5-1.3-0.5-1.8,0l-1.3,1.3l-8.4-3.5v-1.8
                                    c0-0.7-0.6-1.3-1.3-1.3l0,0c-0.7,0-1.3,0.6-1.3,1.3V20l-8.4,3.5l-1.2-1.2c-0.5-0.5-1.3-0.5-1.8,0l0,0c-0.5,0.5-0.5,1.3,0,1.8
                                    l1.2,1.2L14,33.7h-1.8c-0.7,0-1.3,0.6-1.3,1.3l0,0c0,0.7,0.6,1.3,1.3,1.3H14l3.5,8.4L16.2,46c-0.5,0.5-0.5,1.3,0,1.8l0,0
                                    c0.5,0.5,1.3,0.5,1.8,0l1.3-1.3l8.3,3.4v1.8c0,0.7,0.6,1.3,1.3,1.3l0,0c0.7,0,1.3-0.6,1.3-1.3v-1.9l8.3-3.4l1.3,1.3
                                    c0.5,0.5,1.3,0.5,1.8,0l0,0c0.5-0.5,0.5-1.3,0-1.8l-1.3-1.3l3.4-8.3h1.9c0.7,0,1.3-0.6,1.3-1.3l0,0C47,34.3,46.4,33.7,45.7,33.7z
                                     M30.3,23.4l6,2.5l-4.6,4.6c-0.4-0.2-0.9-0.4-1.3-0.6v-6.5H30.3z M27.7,23.4V30c-0.5,0.1-0.9,0.3-1.4,0.6l-4.7-4.7L27.7,23.4z
                                     M19.9,27.7l4.7,4.7c-0.2,0.4-0.4,0.9-0.5,1.3h-6.6L19.9,27.7z M17.4,36.3H24c0.1,0.5,0.3,0.9,0.6,1.3l-4.7,4.7L17.4,36.3z
                                     M27.7,46.5l-6-2.5l4.7-4.7c0.4,0.2,0.8,0.4,1.3,0.5V46.5z M29,37.5c-1.4,0-2.6-1.2-2.6-2.6c0-1.4,1.2-2.6,2.6-2.6s2.6,1.2,2.6,2.6
                                    C31.6,36.4,30.4,37.5,29,37.5z M30.3,46.5v-6.6c0.5-0.1,0.9-0.3,1.3-0.5l4.6,4.6L30.3,46.5z M38,42.2l-4.6-4.6
                                    c0.2-0.4,0.4-0.8,0.6-1.3h6.5L38,42.2z M34,33.7c-0.1-0.5-0.3-0.9-0.5-1.3l4.6-4.6l2.5,6H34V33.7z" />
                                <g class="st2">
                                    <path class="st3" d="M41,49H17c-1.1,0-2-0.9-2-2V23c0-1.1,0.9-2,2-2h24c1.1,0,2,0.9,2,2v24C43,48.1,42.1,49,41,49z" />
                                </g>
                                <g>
                                    <path class="st0" d="M42.2,25c-1.9,0-2.9,0.5-2.9,1.5v17.1c0,1,1,1.5,2.9,1.5v1.8H31.4V45c2,0,3-0.5,3-1.5v-8H23.6v8
                                        c0,1,1,1.5,3,1.5v1.8H15.8V45c1.9,0,2.8-0.5,2.8-1.5V26.4c0-1-0.9-1.5-2.8-1.5V23h10.8v2c-2,0-3,0.5-3,1.5v6.8h10.8v-6.8
                                        c0-1-1-1.5-3-1.5v-1.9h10.8V25z" />
                                </g>
                            </g>
                        </symbol>
                    </svg>

                    <div class="bk-tab2" style="border-left: none; border-right: none;">
                        <div class="bk-tab2-head is-fill">
                            <div class="bk-tab2-nav" style="width: 100%;">
                                <div :title="$t('项目仓库')" :class="['tab2-nav-item', { 'active': tabActiveName === 'privateRepo' }]" @click="tabActiveName = 'privateRepo'" style="width: 200px;">
                                    {{$t('项目仓库')}}
                                </div>
                                <div title="公共仓库" :class="['tab2-nav-item', { 'active': tabActiveName === 'publicRepo' }]" @click="tabActiveName = 'publicRepo'" style="width: 200px;">
                                    {{$t('公共仓库')}}
                                </div>
                            </div>
                        </div>
                        <div class="bk-tab2-content">
                            <div class="biz-namespace mt20">
                                <table class="bk-table biz-templateset-table mb20">
                                    <thead>
                                        <tr>
                                            <th style="width: 120px; padding-left: 0;" class="center">{{$t('图标')}}</th>
                                            <th style="width: 230px; padding-left: 20px;">{{$t('Helm Chart名称')}}</th>
                                            <th style="width: 120px; padding-left: 0;">{{$t('版本')}}</th>
                                            <th style="padding-left: 0;">{{$t('描述')}}</th>
                                            <th style="width: 170px; padding-left: 0;">{{$t('最近更新')}}</th>
                                            <th style="width: 200px; padding-left: 0;">{{$t('操作')}}</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <template v-if="tplList.length">
                                            <tr
                                                v-for="template in tplList"
                                                :key="template.id">
                                                <td colspan="7">
                                                    <table class="biz-inner-table">
                                                        <tr>
                                                            <td class="logo">
                                                                <div class="logo-wrapper" v-if="template.icon && isImage(template.icon)">
                                                                    <img :src="template.icon">
                                                                </div>
                                                                <svg class="biz-set-icon" v-else>
                                                                    <use xlink:href="#biz-set-icon"></use>
                                                                </svg>
                                                            </td>
                                                            <td class="data">
                                                                <bk-tooltip placement="top" :delay="500">
                                                                    <p class="tpl-name">
                                                                        <router-link class="bk-text-button bk-primary bk-button-small" :to="{ name: 'helmTplDetail', params: { tplId: template.id } }">{{template.name}}</router-link>
                                                                    </p>
                                                                    <template slot="content">
                                                                        <p>{{template.name}}</p>
                                                                    </template>
                                                                </bk-tooltip>
                                                            </td>
                                                            <td class="version">
                                                                <template v-if="template.defaultChartVersion">
                                                                    <bk-tooltip placement="top" :delay="500">
                                                                        <p class="tpl-version">
                                                                            {{template.defaultChartVersion.version}}
                                                                        </p>
                                                                        <template slot="content">
                                                                            <p>{{template.defaultChartVersion.version}}</p>
                                                                        </template>
                                                                    </bk-tooltip>
                                                                </template>
                                                                <template v-else>--</template>
                                                            </td>
                                                            <td class="description">
                                                                <p class="text">{{template.description}}</p>
                                                            </td>
                                                            <td class="update">
                                                                {{template.changed_at}}
                                                            </td>
                                                            <td class="action">
                                                                <router-link class="bk-button bk-primary mr5" :to="{ name: 'helmTplInstance', params: { tplId: template.id } }">{{$t('部署')}}</router-link>
                                                                <bk-dropdown-menu class="dropdown-menu" :align="'right'" ref="dropdown" v-if="tabActiveName === 'privateRepo'">
                                                                    <button class="bk-button bk-default btn" slot="dropdown-trigger" style="width: 82px; position: relative;">
                                                                        <span>{{$t('更多')}}</span>
                                                                        <i class="bk-icon icon-angle-down dropdown-menu-angle-down ml0" style="font-size: 10px;"></i>
                                                                    </button>
                                                                    <ul class="bk-dropdown-list" slot="dropdown-content">
                                                                        <li>
                                                                            <a href="javascript:void(0)" @click="handleRemoveChart(template)">{{$t('删除Chart')}}</a>
                                                                        </li>
                                                                        <li>
                                                                            <a href="javascript:void(0)" @click="showChooseDialog(template)">{{$t('删除版本')}}</a>
                                                                        </li>
                                                                    </ul>
                                                                </bk-dropdown-menu>
                                                            </td>
                                                        </tr>
                                                    </table>
                                                </td>
                                            </tr>
                                        </template>
                                        <template v-if="!tplList.length && !showLoading">
                                            <tr>
                                                <td colspan="6">
                                                    <div class="biz-empty-message" style="padding: 80px;">
                                                        <template v-if="isSearchMode">
                                                            {{$t('无数据')}}
                                                        </template>
                                                        <template v-else>
                                                            <span style="vertical-align: middle;">{{$t('无数据，请尝试')}}</span> <a href="javascript:void(0);" class="bk-text-button" @click="syncHelmTpl">{{$t('同步仓库')}}</a>
                                                        </template>
                                                    </div>
                                                </td>
                                            </tr>
                                        </template>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </template>
            </template>
        </div>

        <bk-dialog
            :is-show.sync="delInstanceDialogConf.isShow"
            :width="delInstanceDialogConf.width"
            :title="delInstanceDialogConf.title"
            :close-icon="delInstanceDialogConf.closeIcon"
            :quick-close="false"
            :ext-cls="'biz-config-templateset-del-instance-dialog'"
            @cancel="delInstanceDialogConf.isShow = false">
            <div slot="content" v-bkloading="{ isLoading: isDeleting }">
                <div class="content-inner">
                    <div class="bk-form bk-form-vertical" style="margin-bottom: 20px;">
                        <div class="bk-form-item">
                            <label class="bk-label fl">
                                {{$t('选择要删除的版本')}}:
                            </label>
                            <div class="bk-form-content">
                                <div class="bk-dropdown-box" style="width: 300px;">
                                    <bk-selector
                                        :placeholder="$t('请选择')"
                                        :setting-key="'id'"
                                        :display-key="'version'"
                                        :selected.sync="delInstanceDialogConf.versionIndex"
                                        :list="delInstanceDialogConf.versions"
                                        @item-selected="getReleaseByVersion">
                                    </bk-selector>
                                </div>
                            </div>
                        </div>
                    </div>
                    <template v-if="!isReleaseLoading">
                        <!-- <template v-if="delInstanceDialogConf.versionIndex && !delInstanceDialogConf.releases.length">
                            <transition name="fadeDown">
                                <div class="content-trigger-wrapper" style="text-align: left;">
                                    {{$t('没有Release，可进行版本删除操作')}}
                                </div>
                            </transition>
                        </template> -->
                        <template v-if="delInstanceDialogConf && delInstanceDialogConf.releases.length">
                            <p style="font-weight: bold; color: #737987; font-size: 14px; text-align: left;">
                                当前版本含有的Release: <span class="biz-tip f12" style="font-weight: normal;">(格式：命名空间:Release名称)</span>
                            </p>
                            <ul class="key-list mt10 mb10">
                                <li v-for="release of delInstanceDialogConf.releases" :key="release.name">
                                    <span class="key">{{release.namespace}}</span>
                                    <span class="value">{{release.name}}</span>
                                </li>
                            </ul>
                            <div>
                                {{$t('您需要先删除所有Release，再进行版本删除操作')}}
                            </div>
                        </template>
                    </template>
                    <template v-else>
                        <div v-bkloading="{ isLoading: isReleaseLoading }" style="min-height: 50px;"></div>
                    </template>
                </div>
            </div>
            <template slot="footer">
                <div class="bk-dialog-outer">
                    <template v-if="isVersionDeleting">
                        <button type="button" class="bk-dialog-btn bk-dialog-btn-confirm bk-btn-primary is-disabled">
                            {{$t('删除中...')}}
                        </button>
                        <button type="button" class="bk-dialog-btn bk-dialog-btn-cancel is-disabled">
                            {{$t('取消')}}
                        </button>
                    </template>
                    <template v-else>
                        <template v-if="delInstanceDialogConf.versionIndex && !delInstanceDialogConf.releases.length">
                            <button type="button" class="bk-dialog-btn bk-dialog-btn-confirm bk-btn-primary"
                                @click="confirmDelVersion">
                                {{$t('提交')}}
                            </button>
                        </template>
                        <template v-else>
                            <button type="button" class="bk-dialog-btn bk-dialog-btn-confirm bk-btn-primary is-disabled" disabled>
                                {{$t('提交')}}
                            </button>
                        </template>
                        
                        <button type="button" class="bk-dialog-btn bk-dialog-btn-cancel" @click="cancelDelVersion">
                            {{$t('取消')}}
                        </button>
                    </template>
                </div>
            </template>
        </bk-dialog>

        <bk-dialog
            :is-show.sync="delTemplateDialogConf.isShow"
            :width="delTemplateDialogConf.width"
            :close-icon="false"
            :ext-cls="'biz-config-templateset-copy-dialog'"
            :has-header="false"
            :quick-close="false">
            <div slot="content" style="padding: 0 20px;">
                <div style="color: #333; font-size: 20px">
                    Chart【{{delTemplateDialogConf.title}}】包含以下Releases：
                    <span class="biz-tip">(格式：命名空间:Release名称)</span>
                </div>
                <ul class="key-list mt10 mb10">
                    <li v-for="release of delTemplateDialogConf.releases" :key="release.name">
                        <span class="key">{{release.namespace}}</span>
                        <span class="value">{{release.name}}</span>
                    </li>
                </ul>
                <div>
                    {{$t('您需要先删除所有Release，再进行Chart删除操作')}}
                </div>
            </div>
            <template slot="footer">
                <div class="bk-dialog-outer">
                    <button type="button" style="width: 96px;" class="bk-dialog-btn bk-dialog-btn-confirm bk-btn-primary"
                        @click="delTemplateCancel">
                        {{$t('关闭')}}
                    </button>
                </div>
            </template>
        </bk-dialog>

    </div>
</template>

<script>
    import Guide from './guide'
    import Clipboard from 'clipboard'
    import { catchErrorHandler } from '@open/common/util'

    export default {
        components: {
            Guide
        },
        data () {
            return {
                tplList: [],
                publicTplList: [],
                privateTplList: [],
                tplListCache: [],
                showLoading: false,
                exceptionCode: null,
                searchKeyword: '',
                isSearchMode: false,
                curProjectId: '',
                isTplSynLoading: false,
                isRepoDataLoading: false,
                tabActiveName: 'privateRepo',
                delTemplateDialogConf: {
                    isShow: false,
                    width: 650,
                    title: '',
                    closeIcon: false,
                    template: {},
                    releases: []
                },
                isReleaseLoading: false,
                delInstanceDialogConf: {
                    isShow: false,
                    width: 550,
                    title: '',
                    closeIcon: false,
                    template: {},
                    versions: [],
                    releases: [],
                    versionIndex: 0
                },
                reposData: {
                    publicRepos: [],
                    privateRepos: []
                },
                helmDialog: {
                    isShow: false
                }
            }
        },
        computed: {
            curProject () {
                return this.$store.state.curProject
            },
            projectId () {
                this.curProjectId = this.$route.params.projectId
                return this.curProjectId
            },
            projectCode () {
                return this.$route.params.projectCode
            }
        },
        watch: {
            searchKeyword (newVal, oldVal) {
                // 如果删除，为空时触发搜索
                if (oldVal && !newVal) {
                    this.search()
                }
            },
            curProjectId () {
                // 如果不是k8s类型的项目，无法访问些页面，重定向回集群首页
                if (this.curProject && (this.curProject.kind !== PROJECT_K8S && this.curProject.kind !== PROJECT_TKE)) {
                    this.$router.push({
                        name: 'clusterMain',
                        params: {
                            projectId: this.projectId,
                            projectCode: this.projectCode
                        }
                    })
                }
            },
            tabActiveName (val) {
                this.setTplList()
            }
        },
        mounted () {
            this.getTplList()
        },
        methods: {
            /**
             * 显示/隐藏模板仓库密码
             * @param  {object} repo 模板仓库
             */
            togglePassword (repo) {
                repo.isPasswordShow = !repo.isPasswordShow
            },

            /**
             * 显示引导层(如何推送Helm Chart到项目仓库？)
             */
            showGuide () {
                this.$refs.clusterGuide.show()
            },

            /**
             * 获取集群对应的helm仓库信息
             * @param  {object} cluster 集群
             */
            async getHelmDeops (cluster) {
                this.reposData = {
                    publicRepos: [],
                    privateRepos: []
                }
                this.isRepoDataLoading = true

                try {
                    const res = await this.$store.dispatch('helm/getHelmDeops', {
                        projectId: this.projectId,
                        clusterId: cluster.cluster_id
                    })

                    const repos = res.data.results
                    repos.forEach(item => {
                        item.isPasswordShow = false
                        // 区分私有和公有
                        if (item.name === 'public-repo') {
                            this.reposData.publicRepos.push(item)
                        } else {
                            this.reposData.privateRepos.push(item)
                        }
                    })
                    this.helmDialog.isShow = true

                    setTimeout(() => {
                        this.clipboardInstance = new Clipboard('.copy-btn')
                        this.clipboardInstance.on('success', e => {
                            this.$bkMessage({
                                theme: 'success',
                                message: this.$t('复制成功')
                            })
                        })
                    }, 2000)
                } catch (e) {
                    catchErrorHandler(e, this)
                } finally {
                    this.isRepoDataLoading = false
                }
            },

            /**
             * 隐藏helm仓库信息
             */
            hideHelmDialog () {
                this.helmDialog.isShow = false
            },

            /**
             * 同步仓库
             */
            async syncHelmTpl () {
                if (this.isTplSynLoading) {
                    return false
                }

                this.isTplSynLoading = true
                try {
                    await this.$store.dispatch('helm/syncHelmTpl', { projectId: this.projectId })

                    this.$bkMessage({
                        theme: 'success',
                        message: this.$t('同步成功')
                    })
                    this.getTplList()
                } catch (e) {
                    catchErrorHandler(e, this)
                } finally {
                    this.isTplSynLoading = false
                }
            },

            /**
             * 简单判断是否为图片
             * @param  {string} img 图片url
             * @return {Boolean} true/false
             */
            isImage (img) {
                if (!img) {
                    return false
                }
                if (img.startsWith('http://') || img.startsWith('https://') || img.startsWith('data:image/')) {
                    return true
                }
                return false
            },

            /**
             * 获取模板列表
             */
            async getTplList () {
                const projectId = this.projectId
                this.showLoading = true

                try {
                    const res = await this.$store.dispatch('helm/getTplList', projectId)

                    const publicRepo = []
                    const privateRepo = []

                    // 进行分类，包括项目仓库和私有仓库
                    const tplList = res.data.filter(item => {
                        return item.defaultChartVersion
                    })

                    tplList.forEach(item => {
                        if (item.repository.name === 'public-repo') {
                            publicRepo.push(item)
                        } else {
                            privateRepo.push(item)
                        }
                    })
                    this.publicTplList = publicRepo
                    this.privateTplList = privateRepo
                    this.setTplList()
                } catch (e) {
                    catchErrorHandler(e, this)
                } finally {
                    this.showLoading = false
                }
            },

            /**
             * 根据当前显示公有/私有模板
             */
            setTplList () {
                this.clearSearch()
                if (this.tabActiveName === 'publicRepo') {
                    this.tplList = JSON.parse(JSON.stringify(this.publicTplList))
                } else {
                    this.tplList = JSON.parse(JSON.stringify(this.privateTplList))
                }
                this.tplListCache = JSON.parse(JSON.stringify(this.tplList))
            },

            /**
             * 搜索
             */
            search () {
                const keyword = this.searchKeyword
                if (keyword) {
                    const results = this.tplListCache.filter(item => {
                        if (item.name.indexOf(keyword) > -1) {
                            return true
                        } else {
                            return false
                        }
                    })
                    this.tplList.splice(0, this.tplList.length, ...results)
                    this.isSearchMode = true
                } else {
                    this.tplList.splice(0, this.tplList.length, ...this.tplListCache)
                    this.isSearchMode = false
                }
            },

            /**
             * 清除搜索
             */
            clearSearch () {
                this.searchKeyword = ''
                this.search()
            },

            async handleRemoveChart (template) {
                // if (!template.permissions.delete) {
                //     const params = {
                //         project_id: this.projectId,
                //         policy_code: 'delete',
                //         resource_code: template.id,
                //         resource_name: template.name,
                //         resource_type: 'templates'
                //     }
                //     await this.$store.dispatch('getResourcePermissions', params)
                // }
                try {
                    // 先检测当前Chart是否有release
                    const res = await this.$store.dispatch('helm/getExistReleases', {
                        projectId: this.projectId,
                        templateId: template.id
                    })

                    // 如果没有release，可删除
                    if (!res.data.length) {
                        const me = this
                        me.$bkInfo({
                            title: ``,
                            clsName: 'biz-remove-dialog',
                            content: me.$createElement('p', {
                                class: 'biz-confirm-desc'
                            }, `${this.$t('确定要Chart')}【${template.name}】?`),
                            async confirmFn () {
                                me.deleteTemplate(template)
                            }
                        })
                    } else {
                        this.delTemplateDialogConf.isShow = true
                        this.delTemplateDialogConf.title = template.name
                        this.delTemplateDialogConf.template = Object.assign({}, template)
                        this.delTemplateDialogConf.releases = res.data
                    }
                } catch (e) {
                    catchErrorHandler(e, this)
                }
            },

            /**
             * 确认删除Chart
             * @param {Object} template 当前模板集对象
             */
            async deleteTemplate (template) {
                try {
                    await this.$store.dispatch('helm/removeTemplate', {
                        projectId: this.projectId,
                        templateId: template.id
                    })

                    this.$bkMessage({
                        theme: 'success',
                        message: this.$t('删除成功')
                    })
                    this.getTplList()
                } catch (e) {
                    catchErrorHandler(e, this)
                }
            },

            /**
             * 确认删除版本
             * @param {Object} template 当前模板集对象
             */
            async deleteTemplateVersion () {
                this.isVersionDeleting = true
                try {
                    await this.$store.dispatch('helm/removeTemplate', {
                        projectId: this.projectId,
                        templateId: this.delInstanceDialogConf.template.id,
                        versionId: this.delInstanceDialogConf.versionIndex
                    })

                    this.$bkMessage({
                        theme: 'success',
                        message: this.$t('删除成功')
                    })

                    this.getTplList()
                    this.cancelDelVersion()
                } catch (e) {
                    catchErrorHandler(e, this)
                } finally {
                    this.isVersionDeleting = false
                }
            },

            showTemplateVersion () {
                const template = Object.assign({}, this.delTemplateDialogConf.template)
                this.delTemplateDialogConf.isShow = false
                this.delTemplateDialogConf.title = ''
                this.delTemplateDialogConf.template = Object.assign({}, {})
                this.delTemplateDialogConf.releases = []
                this.showChooseDialog(template)
            },

            delTemplateCancel () {
                this.delTemplateDialogConf.isShow = false
                this.delTemplateDialogConf.title = ''
                this.delTemplateDialogConf.template = Object.assign({}, {})
                this.delTemplateDialogConf.releases = []
            },

            async showChooseDialog (template) {
                // 之前没选择过，那么展开第一个
                this.delInstanceDialogConf.title = `${this.$t('删除')}【${template.name}】${this.$t('Chart的版本')}`
                this.delInstanceDialogConf.isShow = true
                this.delInstanceDialogConf.template = Object.assign({}, template)
                this.delInstanceDialogConf.versions = []
                this.delInstanceDialogConf.releases = []
                this.delInstanceDialogConf.versionIndex = 0

                this.getTplVersions()
            },

            async getTplVersions () {
                const projectId = this.projectId
                const tplId = this.delInstanceDialogConf.template.id

                try {
                    const res = await this.$store.dispatch('helm/getTplVersions', { projectId, tplId })
                    this.delInstanceDialogConf.versions = res.data.results
                    this.delInstanceDialogConf.releases = []
                } catch (e) {
                    catchErrorHandler(e, this)
                }
            },

            async getReleaseByVersion (versionId) {
                this.isReleaseLoading = true
                try {
                    const res = await this.$store.dispatch('helm/getExistReleases', {
                        projectId: this.projectId,
                        templateId: this.delInstanceDialogConf.template.id,
                        versionId: versionId
                    })

                    this.delInstanceDialogConf.releases = res.data
                } catch (e) {
                    catchErrorHandler(e, this)
                } finally {
                    this.isReleaseLoading = false
                }
            },

            cancelDelVersion () {
                // 之前没选择过，那么展开第一个
                this.delInstanceDialogConf.title = ''
                this.delInstanceDialogConf.isShow = false
                this.delInstanceDialogConf.template = {}
                this.delInstanceDialogConf.versions = []
                this.delInstanceDialogConf.releases = []
                this.delInstanceDialogConf.versionIndex = 0
            },

            /**
             * 删除命名空间弹层确认
             */
            async confirmDelVersion () {
                const me = this
                if (!this.delInstanceDialogConf.versionIndex) {
                    this.$bkMessage({
                        theme: 'error',
                        message: this.$t('请选择Chart版本')
                    })
                    return
                }

                this.$bkInfo({
                    title: this.$t('确定删除版本？'),
                    async confirmFn () {
                        me.deleteTemplateVersion()
                    }
                })
            }
        }
    }
</script>

<style scoped>
    @import './tpl-list.css';
</style>

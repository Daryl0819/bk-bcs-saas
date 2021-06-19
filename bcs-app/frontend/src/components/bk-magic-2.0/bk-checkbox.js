export default {
    name: 'bk-checkbox',
    props: {
        value: {
            type: [String, Number, Boolean],
            default: undefined
        }
    },
    watch: {
        value () {
            this.$forceUpdate()
        }
    },
    render (h) {
        const slots = Object.keys(this.$slots)
            .reduce((arr, key) => arr.concat(this.$slots[key]), [])
            .map(vnode => {
                vnode.context = this._self
                return vnode
            })
        return h('bcs-checkbox', {
            on: this.$listeners, // 透传事件
            props: this.$props, // 透传props
            scopedSlots: this.$scopedSlots, // 透传scopedSlots
            attrs: this.$attrs // 透传属性，非props
        }, slots)
    }
}

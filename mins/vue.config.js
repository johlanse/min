module.exports = {
    // 配置网页的路径
    publicPath: '/static/dist',
    // 配置生成的html名字
    indexPath: 'default.html',
    chainWebpack: config => {
        config
            .plugin('html')
            .tap(args => {
                // 配置前端title
                args[0].title= 'MINS'
                return args
            })
    }
}

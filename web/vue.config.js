module.exports = {
    pages: {
        index: {
            // entry for the page
            entry: "src/main.js",

            // the source template
            template: "src/views/templates/index.html",

            // output as dist/index.html
            filename: "index.html",

            // when using title option,
            // template title tag needs to be <title><%= htmlWebpackPlugin.options.title %></title>
            title: "BKMS",

            // chunks to include on this page, by default includes
            // extracted common chunks and vendor chunks.
            chunks: ["chunk-vendors", "chunk-common", "index"]
        },

    },
    configureWebpack: {
        resolve:{
            alias:{
                crypto: false
            },
        },
    },
    devServer:{
        host:'localhost',
        port:'3041',
        proxy:{
            '/api':{
                target: 'http://localhost:3040/api',
                changeOrgin: true,
                pathRewrite:{
                    '^/api':'',
                }
   
            }
        }
    }
};

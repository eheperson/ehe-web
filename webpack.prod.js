const path = require('path');
const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
    mode: 'development',
    devServer: {
        static: {
            directory: path.join(__dirname, 'static/dist'),
        },
        port: 3000,
        proxy: {
            '/api': 'http://localhost:8080'
        }
    }
});

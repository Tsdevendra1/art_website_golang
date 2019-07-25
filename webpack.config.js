const HtmlWebPackPlugin = require("html-webpack-plugin");
var path = require('path');
module.exports = {
    entry: './src/scripts/index.tsx',
    module: {
        rules: [
            {
                test: /\.(ts|tsx)?$/,
                include: path.resolve(__dirname, 'src'),
                use: [
                    {
                        loader: 'babel-loader'
                    }
                ]
            },
            {
                test: /\.html$/,
                use: [
                    {
                        loader: "html-loader"
                    }
                ]
            }
        ]
    },
    watch: true,
    resolve: {
        extensions: ['.tsx', '.ts', '.js']
    },
    output: {
        filename: 'bundle.js',
        path: path.resolve(__dirname, 'dist')
    },
    plugins: [
        new HtmlWebPackPlugin({
            template: "./src/templates/index.html",
        })
    ]
};
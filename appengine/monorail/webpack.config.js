/* Copyright 2019 The Chromium Authors. All Rights Reserved.
 *
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

const path = require('path');
const webpack = require('webpack');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;

const config = {
  entry: {
    'mr-app': './static_src/elements/mr-app/mr-app.js',
    'mr-profile-page': './static_src/elements/chdir/mr-profile-page/mr-profile-page.js',
    'ezt-element-package': './static_src/elements/ezt/ezt-element-package.js',
    'ezt-footer-scripts-package': './static_src/elements/ezt/ezt-footer-scripts-package.js',
  },
  devtool: 'eval-source-map',
  plugins: [],
  resolve: {
    extensions: ['.js'],
    modules: ['node_modules', 'static_src'],
  },
  output: {
    filename: '[name].min.js',
    path: path.resolve(__dirname, 'static/dist'),
    // __webpack_public_path__ is used to dynamically set the public path to
    // the App Engine version URL.
    publicPath: '/static/dist/',
  },
  externals: {
    moment: 'moment',
  },
};

module.exports = (env, argv) => {
  if (argv.mode === 'production') {
    // Settings for deploying JS to production.
    config.devtool = 'source-map';

    config.plugins = config.plugins.concat([
      new webpack.DefinePlugin(
        {'process.env.NODE_ENV': '"production"'}
      ),
    ]);
  }

  if (argv.analyze) {
    config.plugins.push(new BundleAnalyzerPlugin());
  }
  return config;
};

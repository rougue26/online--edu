const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  // 禁用ESLint检查
  lintOnSave: false,
  // 配置开发服务器代理解决跨域问题
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8082',
        changeOrigin: true,
        timeout: 10000, // 增加超时设置到10秒
        onProxyRes: (proxyRes) => {
          proxyRes.headers['Cross-Origin-Resource-Policy'] = 'cross-origin';
          proxyRes.headers['Cross-Origin-Opener-Policy'] = 'same-origin-allow-popups';
          proxyRes.headers['Access-Control-Allow-Origin'] = '*';
        }
      },
      // 为视频资源添加代理
      '/videos': {
        target: 'https://sample-videos.com',
        changeOrigin: true,
        pathRewrite: {
          '^/videos': ''
        },
        timeout: 15000,
        headers: {
          'Origin': 'https://sample-videos.com',
          'Referer': 'https://sample-videos.com/',
          'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
        },
        onProxyRes: (proxyRes) => {
          proxyRes.headers['Cross-Origin-Resource-Policy'] = 'cross-origin';
          proxyRes.headers['Cross-Origin-Opener-Policy'] = 'same-origin-allow-popups';
          proxyRes.headers['Access-Control-Allow-Origin'] = '*';
        },
        onError: (err, req, res) => {
          console.error('Video proxy error:', err);
          res.writeHead(502, { 'Content-Type': 'application/json' });
          res.end(JSON.stringify({
            error: 'Video proxy error',
            message: 'Failed to fetch video'
          }));
        }
      },
      // 为图片资源添加单独的代理
    '/photos': {
      target: 'https://picsum.photos',
      changeOrigin: true,
      pathRewrite: {
        '^/photos': ''
      },
      timeout: 15000, // 增加超时设置到15秒
      headers: {
        'Origin': 'https://picsum.photos',
        'Referer': 'https://picsum.photos/',
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
      },
      // 处理 ORB 问题
      onProxyRes: (proxyRes) => {
        proxyRes.headers['Cross-Origin-Resource-Policy'] = 'cross-origin';
        proxyRes.headers['Cross-Origin-Opener-Policy'] = 'unsafe-none';
        proxyRes.headers['Access-Control-Allow-Origin'] = '*';
        proxyRes.headers['Access-Control-Allow-Methods'] = 'GET, POST, PUT, DELETE, OPTIONS';
        proxyRes.headers['Access-Control-Allow-Headers'] = 'Content-Type, Authorization';
        // 添加更多安全头部
        proxyRes.headers['Content-Security-Policy'] = 'default-src \'self\'; img-src \'self\' https://picsum.photos data:;';
        proxyRes.headers['X-Content-Type-Options'] = 'nosniff';
      },
      // 处理代理错误
      onError: (err, req, res) => {
        console.error('Proxy error:', err);
        res.writeHead(502, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({
          error: 'Proxy error',
          message: 'Failed to fetch image from picsum.photos'
        }));
      }
    },
    // 为fastly.picsum.photos添加代理
    '/fastly-photos': {
      target: 'https://fastly.picsum.photos',
      changeOrigin: true,
      pathRewrite: {
        '^/fastly-photos': ''
      },
      timeout: 15000,
      headers: {
        'Origin': 'https://fastly.picsum.photos',
        'Referer': 'https://fastly.picsum.photos/',
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
      },
      onProxyRes: (proxyRes) => {
        proxyRes.headers['Cross-Origin-Resource-Policy'] = 'cross-origin';
        proxyRes.headers['Cross-Origin-Opener-Policy'] = 'unsafe-none';
        proxyRes.headers['Access-Control-Allow-Origin'] = '*';
        proxyRes.headers['Access-Control-Allow-Methods'] = 'GET, POST, PUT, DELETE, OPTIONS';
        proxyRes.headers['Access-Control-Allow-Headers'] = 'Content-Type, Authorization';
        proxyRes.headers['Content-Security-Policy'] = 'default-src \'self\'; img-src \'self\' https://fastly.picsum.photos data:;';
        proxyRes.headers['X-Content-Type-Options'] = 'nosniff';
      },
      onError: (err, req, res) => {
        console.error('Fastly proxy error:', err);
        res.writeHead(502, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({
          error: 'Fastly proxy error',
          message: 'Failed to fetch image from fastly.picsum.photos'
        }));
      }
    }
    },
    // 启用webpack-dev-server的错误覆盖层
    client: {
      overlay: {
        errors: true,
        warnings: false
      }
    },
    // 增加WebSocket支持以提高开发体验
    setupMiddlewares: (middlewares, devServer) => {
      if (!devServer) {
        throw new Error('webpack-dev-server is not defined');
      }

      return middlewares;
    }
  },
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '南知教育系统'
      return args
    })
  }
})

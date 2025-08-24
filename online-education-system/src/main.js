import { createApp } from 'vue'
import App from './App.vue'
// 引入Element Plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// 引入路由
import router from './router'

// 解决ResizeObserver循环警告
// 方案: 使用requestAnimationFrame和错误抑制相结合
(function() {
  const originalResizeObserver = window.ResizeObserver;
  if (!originalResizeObserver) return;

  // 跟踪已经记录的错误，避免重复报告
  const reportedErrors = new Set();

  window.ResizeObserver = class ResizeObserver extends originalResizeObserver {
    constructor(callback) {
      // 使用requestAnimationFrame包装回调，避免同步更新导致的循环
      const wrappedCallback = (...args) => {
        requestAnimationFrame(() => {
          try {
            callback(...args);
          } catch (error) {
            const errorMessage = error.message || '';
            // 只忽略特定的ResizeObserver循环错误
            if (errorMessage.includes('ResizeObserver loop completed with undelivered notifications')) {
              // 避免重复报告相同的错误
              if (!reportedErrors.has(errorMessage)) {
                console.warn('忽略ResizeObserver循环警告:', error);
                reportedErrors.add(errorMessage);
              }
            } else {
              // 对于其他错误，仍然抛出
              throw error;
            }
          }
        });
      };

      super(wrappedCallback);
    }
  };
})();

const app = createApp(App)
// 使用Element Plus
app.use(ElementPlus)
// 使用路由
app.use(router)

// 引入API插件
import api from './api'
// 使用API插件
app.use(api)

app.mount('#app')

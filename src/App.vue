<template>
  <div id="app">
    <!-- 导航栏 -->
    <el-header>
      <div class="logo">
        <img src="./assets/logo.svg" alt="教育系统logo">
        <span>在线教育系统</span>
      </div>
      <el-menu :default-active="activeIndex" mode="horizontal" @select="handleMenuSelect">
        <el-menu-item index="home">首页</el-menu-item>
        <el-sub-menu index="categories">
          <template #title>课程分类</template>
          <template v-if="loading">
            <el-menu-item index="loading"><el-skeleton /></el-menu-item>
          </template>
          <template v-else-if="categories.length > 0">
            <el-menu-item v-for="category in categories" :key="category.id" :index="`category-${category.id}`">{{ category.name }}</el-menu-item>
          </template>
          <template v-else>
            <el-menu-item index="no-categories">暂无分类数据</el-menu-item>
          </template>
        </el-sub-menu>
        <el-menu-item index="popular-courses">热门课程</el-menu-item>
        <el-menu-item index="community">社区讨论</el-menu-item>
        <el-menu-item index="videos">视频中心</el-menu-item>
        <el-menu-item index="about">关于我们</el-menu-item>
      </el-menu>
      <div class="user-actions">
        <el-button v-if="!isLogin" type="primary" size="small" @click="goToLogin">登录</el-button>
        <el-button v-if="!isLogin" size="small" @click="goToRegister">注册</el-button>
        <el-dropdown v-if="isLogin" placement="bottom-end">
          <span class="user-avatar">
            <img :src="userInfo.avatar" alt="用户头像">
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="goToProfile">个人中心</el-dropdown-item>
              <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>

    <!-- 主要内容区域 -->
    <el-main>
      <template v-if="$route.meta.loading">
        <div class="loading-container">
          <el-skeleton active />
        </div>
      </template>
      <router-view/>
    </el-main>

    <!-- 页脚 -->
    <el-footer>
      <div class="footer-content">
        <div class="footer-section">
          <h3>关于我们</h3>
          <p>提供高质量的在线教育课程，帮助你提升技能，实现职业发展。</p>
        </div>
        <div class="footer-section">
          <h3>联系我们</h3>
          <p>邮箱：contact@example.com</p>
          <p>电话：123-4567-8910</p>
        </div>
        <div class="footer-section">
          <h3>快速链接</h3>
          <ul>
            <li><a href="#" @click.prevent="goToHome">首页</a></li>
            <li><a href="#" @click.prevent="goToCourseList('all')">课程分类</a></li>
            <li><a href="#" @click.prevent="goToCourseList('popular')">热门课程</a></li>
          </ul>
        </div>
      </div>
      <div class="copyright">
        南知教育系统
      </div>
    </el-footer>
  </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { ElMessage } from 'element-plus';
import { courseCategoryAPI } from './api/index';

export default {
  name: 'App',
  setup() {
    const router = useRouter();
    const route = useRoute();
    const activeIndex = ref('home');
    const isLogin = ref(!!localStorage.getItem('token'));
    const userInfo = ref({
      avatar: 'https://picsum.photos/100/100?random=1'
    });
    const categories = ref([]);
    const loading = ref(false);

    // 获取课程分类
    const fetchCategories = async () => {
      try {
        loading.value = true;
        console.log('开始获取课程分类...');
        const response = await courseCategoryAPI.getAllCategories();
        categories.value = response.data || [];
        console.log('获取课程分类成功:', categories.value);
        if (categories.value.length === 0) {
          ElMessage.info('当前没有可用的课程分类');
        }
      } catch (error) {
        console.error('获取课程分类失败:', error);
        const errorMsg = error.response ? error.response.data.message || '服务器错误' : error.message;
        console.error('错误详情:', errorMsg);
        ElMessage.error(`获取课程分类失败: ${errorMsg}`);
        // 提供默认分类数据，避免菜单为空
        categories.value = [
          { id: 1, name: '前端开发' },
          { id: 2, name: '后端开发' },
          { id: 3, name: '移动开发' },
          { id: 4, name: '数据科学' }
        ];
      } finally {
        loading.value = false;
      }
    };

    // 菜单选择处理
    const handleMenuSelect = (key) => {
      switch (key) {
        case 'home':
          goToHome();
          break;
        case 'popular-courses':
          goToCourseList('popular');
          break;
        case 'community':
          goToCommunity();
          break;
        case 'videos':
          goToVideos();
          break;
        default:
          if (key.startsWith('category-')) {
            const categoryId = key.replace('category-', '');
            goToCourseList(categoryId);
          }
      }
    };

    // 导航方法
    const goToHome = () => {
      router.push({ name: 'home' });
    };

    const goToCourseList = (category) => {
      router.push({ name: 'courses', query: { category } });
    };

    const goToLogin = () => {
      router.push({ name: 'login' });
    };

    const goToRegister = () => {
      router.push({ name: 'register' });
    };

    const goToProfile = () => {
      router.push({ name: 'profile' });
    };

    const goToVideos = () => {
      router.push({ name: 'videos' });
    };

    const goToCommunity = () => {
      router.push({ name: 'community' });
    };

    const logout = () => {
      localStorage.removeItem('token');
      isLogin.value = false;
      ElMessage.success('退出登录成功');
      router.push({ name: 'home' });
    };

    // 监听路由变化
    watch(
      () => route,
      (to) => {
        // 根据路由更新菜单激活状态
    if (to) {
      if (to.name === 'home') {
        activeIndex.value = 'home';
      } else if (to.name === 'courseList') {
        if (to.query.category && to.query.category !== 'all') {
          activeIndex.value = `category-${to.query.category}`;
        } else {
              activeIndex.value = 'popular-courses';
            }
          } else if (to.name === 'community') {
            activeIndex.value = 'community';
          } else if (to.name === 'videos') {
            activeIndex.value = 'videos';
          }
        }
      },
      { immediate: true }
    );

    onMounted(() => {
      fetchCategories();
    });

    return {
      activeIndex,
      isLogin,
      userInfo,
      categories,
      handleMenuSelect,
      goToHome,
      goToCourseList,
      goToLogin,
      goToRegister,
      goToProfile,
      goToVideos,
      goToCommunity,
      logout
    };
  }
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.el-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #fff;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 0 50px;
}

.logo {
  display: flex;
  align-items: center;
}

.logo img {
  width: 40px;
  height: 40px;
  margin-right: 10px;
}

.logo span {
  font-size: 20px;
  font-weight: bold;
  color: #1890ff;
}

.user-actions {
  display: flex;
  gap: 10px;
  align-items: center;
}

.user-avatar img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  cursor: pointer;
}

.el-main {
  flex: 1;
  padding: 20px 50px;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 400px;
}

.el-carousel {
  margin-bottom: 30px;
}

.el-carousel img {
  width: 100%;
  height: 400px;
  object-fit: cover;
}

.category-section,
.popular-courses {
  margin-bottom: 40px;
}

.category-section h2,
.popular-courses h2 {
  font-size: 24px;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 2px solid #1890ff;
  display: inline-block;
}

.image {
  width: 100%;
  height: 160px;
  object-fit: cover;
}

.category-info {
  padding: 10px;
}

.category-title {
  font-size: 16px;
  font-weight: bold;
}

.course-count {
  color: #999;
  font-size: 14px;
  margin-top: 5px;
}

.course-image {
  width: 100%;
  height: 160px;
  object-fit: cover;
}

.course-info {
  padding: 10px;
}

.course-title {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
}

.course-teacher {
  color: #666;
  font-size: 14px;
  margin-bottom: 5px;
}

.course-price {
  color: #f56c6c;
  font-size: 16px;
  font-weight: bold;
}

.el-footer {
  background-color: #f5f5f5;
  padding: 30px 50px;
  color: #666;
}

.footer-content {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.footer-section {
  flex: 1;
  margin-right: 30px;
}

.footer-section h3 {
  font-size: 18px;
  margin-bottom: 15px;
  color: #333;
}

.footer-section ul {
  list-style: none;
  padding: 0;
}

.footer-section li {
  margin-bottom: 8px;
}

.footer-section a {
  color: #666;
  text-decoration: none;
}

.footer-section a:hover {
  color: #1890ff;
}

.copyright {
  text-align: center;
  padding-top: 20px;
  border-top: 1px solid #e5e5e5;
}
</style>

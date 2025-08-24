<template>
  <div class="admin-container">
    <el-container style="height: 100vh; border: 1px solid #eee;">
      <!-- 侧边栏 -->
      <el-aside width="200px" style="background-color: #f5f7fa;">
        <div class="logo-container">
          <h2>后台管理系统</h2>
        </div>
        <el-menu
          default-active="1"
          class="el-menu-vertical-demo"
          router
          background-color="#f5f7fa"
          text-color="#333"
          active-text-color="#1890ff"
        >
          <el-menu-item index="1" route="/admin/dashboard">
            <el-icon><Dashboard /></el-icon>
            <span>仪表盘</span>
          </el-menu-item>
          <el-sub-menu index="2">
            <template #title>
              <el-icon><Cpu /></el-icon>
              <span>课程管理</span>
            </template>
            <el-menu-item index="2-1" route="/admin/dashboard/courses">课程列表</el-menu-item>
            <el-menu-item index="2-2" route="/admin/dashboard/courses/create">创建课程</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title>
              <el-icon><User /></el-icon>
              <span>用户管理</span>
            </template>
            <el-menu-item index="3-1" route="/admin/dashboard/users">用户列表</el-menu-item>
            <el-menu-item index="3-2" route="/admin/dashboard/roles">角色权限</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="4">
            <template #title>
              <el-icon><VideoCamera /></el-icon>
              <span>视频管理</span>
            </template>
            <el-menu-item index="4-1" route="/admin/dashboard/videos">视频列表</el-menu-item>
            <el-menu-item index="4-2" route="/admin/dashboard/videos/create">上传视频</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="5">
            <template #title>
              <el-icon><Message /></el-icon>
              <span>社区管理</span>
            </template>
            <el-menu-item index="5-1" route="/admin/dashboard/posts">帖子管理</el-menu-item>
            <el-menu-item index="5-2" route="/admin/dashboard/comments">评论管理</el-menu-item>
          </el-sub-menu>
          <el-menu-item index="6" route="/admin/dashboard/settings">
            <el-icon><Setting /></el-icon>
            <span>系统设置</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <!-- 主内容区 -->
      <el-container>
        <el-header style="text-align: right; font-size: 12px; border-bottom: 1px solid #eee;">
          <div class="user-info">
            <span>管理员</span>
            <el-dropdown>
              <span class="el-dropdown-link">
                <el-avatar :size="32"><User /></el-avatar>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>个人中心</el-dropdown-item>
                  <el-dropdown-item>退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-header>

        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { Dashboard, Cpu, User, VideoCamera, Message, Setting, Logout } from '@element-plus/icons-vue';

// 检查是否有权限访问
const router = useRouter();

onMounted(() => {
  const isAdmin = localStorage.getItem('isAdmin');
  if (!isAdmin) {
    router.push('/login');
  }
});
</script>

<style scoped>
.admin-container {
  height: 100%;
  width: 100%;
}

.logo-container {
  text-align: center;
  padding: 20px 0;
  border-bottom: 1px solid #eee;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 0 20px;
  height: 100%;
}
</style>
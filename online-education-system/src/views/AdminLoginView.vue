<template>
  <div class="login-container admin-login">
    <div class="login-form-wrapper">
      <div class="logo-container">
        <img src="@/assets/logo.png" alt="教育平台logo" class="logo">
        <h2 class="logo-text">后台管理系统</h2>
      </div>      
      <el-form ref="loginFormRef" :model="loginForm" :rules="rules" class="login-form">
        <el-form-item prop="username">
          <el-input
            v-model="loginForm.username"
            placeholder="请输入管理员用户名"
            prefix-icon="User"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入管理员密码"
            prefix-icon="Lock"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item>
          <div class="form-footer">
            <el-checkbox v-model="loginForm.remember">记住密码</el-checkbox>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="login-button"
            @click="handleLogin"
            :loading="loading"
          >
            管理员登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { userAPI } from '@/api/index';

// 路由和导航
const router = useRouter();

// 状态管理
const loginForm = reactive({
  username: '',
  password: '',
  remember: false
});

const rules = {
  username: [
    { required: true, message: '请输入管理员用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入管理员密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ]
};

const loading = ref(false);
const loginFormRef = ref(null);

// 生命周期钩子
onMounted(() => {
  // 检查是否有记住的密码
  const savedAdminUsername = localStorage.getItem('rememberedAdminUsername');
  if (savedAdminUsername) {
    loginForm.username = savedAdminUsername;
    loginForm.remember = true;
  }
});

// 处理登录
const handleLogin = () => {
  loginFormRef.value.validate(valid => {
    if (valid) {
      loading.value = true;

      // 处理登录
          userAPI.login({ username: loginForm.username, password: loginForm.password })
            .then(response => {
              console.log('Admin login response:', response);
              if (!response || !response.token) {
                throw new Error('登录失败: 无效的响应');
              }

              // 验证是否为管理员
              console.log('User role:', response.user?.role);
              if (response.user?.role !== 'admin') {
                throw new Error('登录失败: 非管理员账户');
              }

          const token = response.token;
              console.log('Token:', token);
              localStorage.setItem('token', token);
              localStorage.setItem('userId', response.user?.ID || '');
              localStorage.setItem('userAvatar', response.user?.Avatar || '');
              localStorage.setItem('isAdmin', 'true');
              console.log('Local storage after login:', {
                token: localStorage.getItem('token'),
                userId: localStorage.getItem('userId'),
                isAdmin: localStorage.getItem('isAdmin')
              });

              // 记住密码
              if (loginForm.remember) {
                localStorage.setItem('rememberedAdminUsername', loginForm.username);
              } else {
                localStorage.removeItem('rememberedAdminUsername');
              }

              loading.value = false;
              ElMessage.success('管理员登录成功');
              console.log('Redirecting to admin dashboard...');
              router.push('/admin/dashboard');
        })
        .catch(error => {
          loading.value = false;
          ElMessage.error(error.message || '登录失败: 用户名或密码错误');
        });
    }
  });
};
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f7fa;
}

.admin-login {
  background-color: #eef1f5;
}

.login-form-wrapper {
  width: 400px;
  padding: 30px;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.logo-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 30px;
}

.logo {
  width: 80px;
  height: 80px;
  margin-bottom: 15px;
}

.logo-text {
  font-size: 20px;
  font-weight: bold;
  color: #303133;
}

.login-form {
  width: 100%;
}

.form-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.forget-password {
  color: #409eff;
  text-decoration: none;
}

.login-button {
  width: 100%;
}

.register-link {
  margin-top: 20px;
  text-align: center;
  color: #606266;
}

.register-link a {
  color: #409eff;
  text-decoration: none;
}
</style>
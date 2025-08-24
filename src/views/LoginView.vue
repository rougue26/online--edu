<template>
  <div class="login-container">
    <div class="login-form-wrapper">
      <div class="logo-container">
        <img src="@/assets/logo.png" alt="教育平台logo" class="logo">
        <h2 class="logo-text">在线教育平台</h2>
      </div>
      <el-form ref="loginFormRef" :model="loginForm" :rules="rules" class="login-form">
        <el-form-item prop="username">
          <el-input
            v-model="loginForm.username"
            placeholder="请输入用户名或邮箱"
            prefix-icon="User"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item>
          <div class="form-footer">
            <el-checkbox v-model="loginForm.remember">记住密码</el-checkbox>
            <a href="#" class="forget-password">忘记密码?</a>
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
            登录
          </el-button>
        </el-form-item>
      </el-form>
      <div class="register-link">
        还没有账号? <a href="#" @click.prevent="toRegister">立即注册</a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { userAPI } from '@/api/index';

// 路由和导航
const route = useRoute();
const router = useRouter();

// 状态管理
const loginForm = reactive({
  username: '',
  password: '',
  remember: false
});

const rules = {
  username: [
    { required: true, message: '请输入用户名或邮箱', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ]
};

const loading = ref(false);
const loginFormRef = ref(null);

// 生命周期钩子
onMounted(() => {
  // 检查是否有记住的密码
  const savedUsername = localStorage.getItem('rememberedUsername');
  if (savedUsername) {
    loginForm.username = savedUsername;
    loginForm.remember = true;
  }
});

// 处理登录
const handleLogin = () => {
  loginFormRef.value.validate(valid => {
    if (valid) {
      loading.value = true;

      // 实际应用中，调用登录API
      userAPI.login({ username: loginForm.username, password: loginForm.password })
        .then(response => {
          console.log('Login response:', response);
          if (!response) {
            throw new Error('Empty response from server');
          }
          if (!response.token) {
            throw new Error('Token not found in response');
          }
          const token = response.token;
          localStorage.setItem('token', token);
          localStorage.setItem('userId', response.user?.ID || '');
          localStorage.setItem('userAvatar', response.user?.Avatar || '');
          // 根据用户角色设置isAdmin标志
          const isAdmin = response.user?.Role === 'admin';
          localStorage.setItem('isAdmin', String(isAdmin));

          // 记住密码
          if (loginForm.remember) {
            localStorage.setItem('rememberedUsername', loginForm.username);
          } else {
            localStorage.removeItem('rememberedUsername');
          }

          loading.value = false;
          ElMessage.success('登录成功');

          // 跳转到首页或之前的页面
          const redirect = route.query.redirect || '/';
          router.push(redirect);
          
          // 登录成功后自动刷新页面一次
          setTimeout(() => {
            router.go(0);
          }, 100);
        })
        .catch(error => {
          loading.value = false;
          ElMessage.error('登录失败: ' + (error.message || '用户名或密码错误'));
        });
    }
  });
};

// 跳转到注册页面
const toRegister = () => {
  router.push({ name: 'register' });
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 20px;
}

.login-form-wrapper {
  width: 100%;
  max-width: 400px;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 40px 30px;
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
  font-size: 24px;
  font-weight: bold;
  color: #333;
}

.login-form {
  width: 100%;
}

.form-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
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
}

.register-link a {
  color: #409eff;
  text-decoration: none;
}
</style>
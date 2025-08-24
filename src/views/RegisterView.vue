<template>
  <div class="register-container">
    <div class="register-form-wrapper">
      <div class="logo-container">
        <img src="@/assets/logo.png" alt="教育平台logo" class="logo">
        <h2 class="logo-text">在线教育平台</h2>
      </div>
      <el-form ref="registerFormRef" :model="registerForm" :rules="rules" class="register-form">
        <el-form-item prop="username">
          <el-input
            v-model="registerForm.username"
            placeholder="请输入用户名"
            prefix-icon="User"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item prop="email">
          <el-input
            v-model="registerForm.email"
            placeholder="请输入邮箱"
            prefix-icon="Message"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="registerForm.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item prop="confirmPassword">
          <el-input
            v-model="registerForm.confirmPassword"
            type="password"
            placeholder="请确认密码"
            prefix-icon="Lock"
            :validate-event="false"
          />
        </el-form-item>
        <el-form-item prop="agreement">
          <el-checkbox v-model="registerForm.agreement">
            我已阅读并同意 <a href="#" class="agreement-link">用户协议</a> 和 <a href="#" class="agreement-link">隐私政策</a>
          </el-checkbox>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="register-button"
            @click="handleRegister"
            :loading="loading"
          >
            立即注册
          </el-button>
        </el-form-item>
      </el-form>
      <div class="login-link">
        已有账号? <a href="#" @click.prevent="toLogin">立即登录</a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { userAPI } from '@/api/index';

// 路由和导航
const router = useRouter();

// 状态管理
const registerForm = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
  agreement: false
});

const loading = ref(false);
const registerFormRef = ref(null);

// 自定义验证器
const validateConfirmPassword = (rule, value, callback) => {
  if (value !== registerForm.password) {
    callback(new Error('两次输入的密码不一致'));
  } else {
    callback();
  }
};

const validateAgreement = (rule, value, callback) => {
  if (!value) {
    callback(new Error('请阅读并同意用户协议和隐私政策'));
  } else {
    callback();
  }
};

// 表单规则
const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '用户名长度在3到20个字符之间', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' },
    { pattern: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,}$/, message: '密码必须包含字母和数字', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' }
  ],
  agreement: [
    { validator: validateAgreement, trigger: 'change' }
  ]
};

// 处理注册
const handleRegister = () => {
  registerFormRef.value.validate(valid => {
    if (valid) {
      loading.value = true;

      // 准备注册数据
      const registerData = {
        username: registerForm.username,
        email: registerForm.email,
        password: registerForm.password
      };

      // 实际应用中，调用注册API
      userAPI.register(registerData)
        .then(response => {
          loading.value = false;
          console.log('注册成功响应:', response);
          ElMessage.success('注册成功，请登录');
          router.push({ name: 'login' });
        })
        .catch(error => {
          loading.value = false;
          console.error('注册失败错误:', error);
          console.error('错误响应数据:', error.response?.data);
          ElMessage.error('注册失败: ' + (error.message || '未知错误'));
        });
    }
  });
};

// 跳转到登录页面
const toLogin = () => {
  router.push({ name: 'login' });
};
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 20px;
}

.register-form-wrapper {
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

.register-form {
  width: 100%;
}

.agreement-link {
  color: #409eff;
  text-decoration: none;
}

.register-button {
  width: 100%;
}

.login-link {
  margin-top: 20px;
  text-align: center;
}

.login-link a {
  color: #409eff;
  text-decoration: none;
}
</style>
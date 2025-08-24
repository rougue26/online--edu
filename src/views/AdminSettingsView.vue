<template>
  <div class="admin-settings-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>系统设置</span>
        </div>
      </template>
      <el-tabs v-model="activeTab">
        <el-tab-pane label="基本设置" name="basic">
          <el-form ref="basicForm" :model="basicForm" label-width="120px">
            <el-form-item label="网站名称" prop="siteName">
              <el-input v-model="basicForm.siteName" placeholder="请输入网站名称"></el-input>
            </el-form-item>
            <el-form-item label="网站描述" prop="siteDescription">
              <el-input v-model="basicForm.siteDescription" type="textarea" placeholder="请输入网站描述"></el-input>
            </el-form-item>
            <el-form-item label="网站Logo">
              <el-upload
                class="avatar-uploader"
                action="#"
                :show-file-list="false"
                :before-upload="beforeUploadLogo"
                :on-success="onSuccessLogo"
              >
                <img v-if="basicForm.siteLogo" :src="basicForm.siteLogo" class="avatar">
                <div v-else class="upload-placeholder">
                  <el-icon><Plus /></el-icon>
                </div>
              </el-upload>
            </el-form-item>
            <el-form-item label="联系方式" prop="contactInfo">
              <el-input v-model="basicForm.contactInfo" placeholder="请输入联系方式"></el-input>
            </el-form-item>
            <el-form-item label="版权信息" prop="copyrightInfo">
              <el-input v-model="basicForm.copyrightInfo" placeholder="请输入版权信息"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleBasicSubmit">保存设置</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="安全设置" name="security">
          <el-form ref="securityForm" :model="securityForm" label-width="120px">
            <el-form-item label="密码有效期(天)" prop="passwordExpiryDays">
              <el-input v-model.number="securityForm.passwordExpiryDays" placeholder="请输入密码有效期"></el-input>
            </el-form-item>
            <el-form-item label="登录失败次数限制" prop="loginFailedLimit">
              <el-input v-model.number="securityForm.loginFailedLimit" placeholder="请输入登录失败次数限制"></el-input>
            </el-form-item>
            <el-form-item label="登录超时时间(分钟)" prop="loginTimeoutMinutes">
              <el-input v-model.number="securityForm.loginTimeoutMinutes" placeholder="请输入登录超时时间"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSecuritySubmit">保存设置</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="邮件设置" name="email">
          <el-form ref="emailForm" :model="emailForm" label-width="120px">
            <el-form-item label="SMTP服务器" prop="smtpServer">
              <el-input v-model="emailForm.smtpServer" placeholder="请输入SMTP服务器"></el-input>
            </el-form-item>
            <el-form-item label="SMTP端口" prop="smtpPort">
              <el-input v-model.number="emailForm.smtpPort" placeholder="请输入SMTP端口"></el-input>
            </el-form-item>
            <el-form-item label="发件人邮箱" prop="senderEmail">
              <el-input v-model="emailForm.senderEmail" placeholder="请输入发件人邮箱"></el-input>
            </el-form-item>
            <el-form-item label="发件人名称" prop="senderName">
              <el-input v-model="emailForm.senderName" placeholder="请输入发件人名称"></el-input>
            </el-form-item>
            <el-form-item label="SMTP用户名" prop="smtpUsername">
              <el-input v-model="emailForm.smtpUsername" placeholder="请输入SMTP用户名"></el-input>
            </el-form-item>
            <el-form-item label="SMTP密码" prop="smtpPassword">
              <el-input v-model="emailForm.smtpPassword" type="password" placeholder="请输入SMTP密码"></el-input>
            </el-form-item>
            <el-form-item label="启用SSL/TLS" prop="enableSsl">
              <el-switch v-model="emailForm.enableSsl"></el-switch>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleEmailSubmit">保存设置</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
// 假设存在settingsAPI
// import { settingsAPI } from '../api/index';
import { Plus } from '@element-plus/icons-vue';

export default {
  name: 'AdminSettingsView',
  components: {
    Plus
  },
  setup() {
    const activeTab = ref('basic');
    const basicForm = ref({
      siteName: '在线教育系统',
      siteDescription: '提供高质量的在线教育课程',
      siteLogo: '',
      contactInfo: 'contact@example.com',
      copyrightInfo: '© 2023 在线教育系统 版权所有'
    });
    const securityForm = ref({
      passwordExpiryDays: 90,
      loginFailedLimit: 5,
      loginTimeoutMinutes: 30
    });
    const emailForm = ref({
      smtpServer: 'smtp.example.com',
      smtpPort: 465,
      senderEmail: 'noreply@example.com',
      senderName: '在线教育系统',
      smtpUsername: 'noreply@example.com',
      smtpPassword: '',
      enableSsl: true
    });
    const basicFormRef = ref(null);
    const securityFormRef = ref(null);
    const emailFormRef = ref(null);

    // 获取系统设置
    const fetchSettings = async () => {
      try {
        // 模拟数据，实际项目中应替换为API调用
        // const res = await settingsAPI.getSettings();
        // basicForm.value = res.basic || basicForm.value;
        // securityForm.value = res.security || securityForm.value;
        // emailForm.value = res.email || emailForm.value;
      } catch (error) {
        console.error('获取系统设置失败:', error);
      }
    };

    // 上传Logo前检查
    const beforeUploadLogo = (file) => {
      const isImage = file.type.startsWith('image/');
      const isLt5M = file.size / 1024 / 1024 < 5;

      if (!isImage) {
        ElMessage.error('只能上传图片格式');
        return false;
      }
      if (!isLt5M) {
        ElMessage.error('图片大小不能超过 5MB');
        return false;
      }
      return true;
    };

    // Logo上传成功
    const onSuccessLogo = (response, file, fileList) => {
      // 模拟Logo URL
      basicForm.value.siteLogo = 'https://example.com/logo/' + Date.now() + '.' + file.name.split('.').pop();
      ElMessage.success('Logo上传成功');
    };

    // 提交基本设置
    const handleBasicSubmit = async () => {
      try {
        await basicFormRef.value.validate();
        // 调用API保存基本设置
        // await settingsAPI.saveBasicSettings(basicForm.value);
        ElMessage.success('基本设置保存成功');
      } catch (error) {
        console.error('保存基本设置失败:', error);
        ElMessage.error('保存基本设置失败，请重试');
      }
    };

    // 提交安全设置
    const handleSecuritySubmit = async () => {
      try {
        await securityFormRef.value.validate();
        // 调用API保存安全设置
        // await settingsAPI.saveSecuritySettings(securityForm.value);
        ElMessage.success('安全设置保存成功');
      } catch (error) {
        console.error('保存安全设置失败:', error);
        ElMessage.error('保存安全设置失败，请重试');
      }
    };

    // 提交邮件设置
    const handleEmailSubmit = async () => {
      try {
        await emailFormRef.value.validate();
        // 调用API保存邮件设置
        // await settingsAPI.saveEmailSettings(emailForm.value);
        ElMessage.success('邮件设置保存成功');
      } catch (error) {
        console.error('保存邮件设置失败:', error);
        ElMessage.error('保存邮件设置失败，请重试');
      }
    };

    onMounted(() => {
      fetchSettings();
    });

    return {
      activeTab,
      basicForm,
      securityForm,
      emailForm,
      basicFormRef,
      securityFormRef,
      emailFormRef,
      beforeUploadLogo,
      onSuccessLogo,
      handleBasicSubmit,
      handleSecuritySubmit,
      handleEmailSubmit
    };
  }
};
</script>

<style scoped>
.admin-settings-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.avatar {
  width: 120px;
  height: 40px;
  display: block;
  object-fit: contain;
}

.upload-placeholder {
  width: 120px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px dashed #d9d9d9;
  background-color: #f5f5f5;
}
</style>
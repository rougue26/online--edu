<template>
  <div class="admin-video-upload-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>上传视频</span>
        </div>
      </template>
      <el-form ref="videoForm" :model="videoForm" label-width="120px">
        <el-form-item label="视频标题" prop="title">
          <el-input v-model="videoForm.title" placeholder="请输入视频标题"></el-input>
        </el-form-item>
        <el-form-item label="视频描述" prop="description">
          <el-input v-model="videoForm.description" type="textarea" placeholder="请输入视频描述"></el-input>
        </el-form-item>
        <el-form-item label="视频文件">
          <el-upload
            class="upload-demo"
            drag
            action="#"
            :before-upload="beforeUpload"
            :on-progress="onProgress"
            :on-success="onSuccess"
            :on-error="onError"
            :file-list="fileList"
            multiple="false"
            :auto-upload="false"
          >
            <el-icon class="el-icon--upload"><Upload /></el-icon>
            <div class="el-upload__text">
              拖拽文件到此处，或<em>点击上传</em>
            </div>
            <template #tip>
              <div class="el-upload__tip">
                支持 MP4、MOV 格式，文件大小不超过 200MB
              </div>
            </template>
          </el-upload>
          <div v-if="uploadProgress > 0 && uploadProgress < 100" class="progress-container">
            <el-progress :percentage="uploadProgress"></el-progress>
          </div>
        </el-form-item>
        <el-form-item label="封面图片">
          <el-upload
            class="avatar-uploader"
            action="#"
            :show-file-list="false"
            :before-upload="beforeUploadCover"
            :on-success="onSuccessCover"
          >
            <img v-if="videoForm.coverImageUrl" :src="videoForm.coverImageUrl" class="avatar">
            <div v-else class="upload-placeholder">
              <el-icon><Plus /></el-icon>
            </div>
          </el-upload>
          <div class="el-upload__tip">点击上传封面图片 (建议尺寸 16:9)</div>
        </el-form-item>
        <el-form-item label="视频分类">
          <el-select v-model="videoForm.categoryId" placeholder="请选择分类">
            <el-option v-for="category in categories" :key="category.id" :label="category.name" :value="category.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="视频时长" prop="duration">
          <el-input v-model="videoForm.duration" placeholder="请输入视频时长(秒)"></el-input>
        </el-form-item>
        <el-form-item label="是否公开">
          <el-switch v-model="videoForm.isPublic"></el-switch>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">提交</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { videoAPI } from '../api/index';
import { Upload, Plus } from '@element-plus/icons-vue';

export default {
  name: 'AdminVideoUploadView',
  components: {
    Upload,
    Plus
  },
  setup() {
    const videoForm = ref({
      title: '',
      description: '',
      videoUrl: '',
      coverImageUrl: '',
      categoryId: '',
      duration: 0,
      isPublic: true
    });
    const fileList = ref([]);
    const uploadProgress = ref(0);
    const categories = ref([]);
    const videoFormRef = ref(null);

    // 获取视频分类
    const fetchCategories = async () => {
      try {
        const res = await videoAPI.getVideoCategories();
        categories.value = res || [];
      } catch (error) {
        console.error('获取视频分类失败:', error);
      }
    };

    // 上传前检查
    const beforeUpload = (file) => {
      const isMP4 = file.type === 'video/mp4';
      const isMOV = file.type === 'video/quicktime';
      const isLt200M = file.size / 1024 / 1024 < 200;

      if (!isMP4 && !isMOV) {
        ElMessage.error('只能上传 MP4 或 MOV 格式的视频');
        return false;
      }
      if (!isLt200M) {
        ElMessage.error('视频大小不能超过 200MB');
        return false;
      }
      return true;
    };

    // 上传进度
    const onProgress = (event, file, fileList) => {
      uploadProgress.value = Math.floor(event.percent);
    };

    // 上传成功
    const onSuccess = (response, file, fileList) => {
      // 实际项目中，这里会处理上传成功后的逻辑
      // 例如获取视频URL等
      uploadProgress.value = 0;
      ElMessage.success('视频上传成功');
      // 模拟视频URL
      videoForm.value.videoUrl = 'https://example.com/videos/' + Date.now() + '.' + file.name.split('.').pop();
    };

    // 上传失败
    const onError = (err, file, fileList) => {
      uploadProgress.value = 0;
      ElMessage.error('视频上传失败，请重试');
    };

    // 封面上传前检查
    const beforeUploadCover = (file) => {
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

    // 封面上传成功
    const onSuccessCover = (response, file, fileList) => {
      // 模拟封面URL
      videoForm.value.coverImageUrl = 'https://example.com/covers/' + Date.now() + '.' + file.name.split('.').pop();
      ElMessage.success('封面上传成功');
    };

    // 提交表单
    const handleSubmit = async () => {
      try {
        await videoFormRef.value.validate();
        // 调用API提交视频信息
        const res = await videoAPI.createVideo(videoForm.value);
        ElMessage.success('视频创建成功');
        // 重置表单
        handleReset();
      } catch (error) {
        console.error('提交视频失败:', error);
        ElMessage.error('提交视频失败，请重试');
      }
    };

    // 重置表单
    const handleReset = () => {
      videoFormRef.value.resetFields();
      fileList.value = [];
      uploadProgress.value = 0;
      videoForm.value = {
        title: '',
        description: '',
        videoUrl: '',
        coverImageUrl: '',
        categoryId: '',
        duration: 0,
        isPublic: true
      };
    };

    onMounted(() => {
      fetchCategories();
    });

    return {
      videoForm,
      fileList,
      uploadProgress,
      categories,
      videoFormRef,
      beforeUpload,
      onProgress,
      onSuccess,
      onError,
      beforeUploadCover,
      onSuccessCover,
      handleSubmit,
      handleReset
    };
  }
};
</script>

<style scoped>
.admin-video-upload-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.avatar {
  width: 178px;
  height: 100px;
  display: block;
  object-fit: cover;
}

.upload-placeholder {
  width: 178px;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px dashed #d9d9d9;
  background-color: #f5f5f5;
}

.progress-container {
  margin-top: 10px;
  width: 100%;
}
</style>
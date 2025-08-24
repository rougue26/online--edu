<template>
  <div class="admin-course-create-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>创建课程</span>
        </div>
      </template>

      <el-form
        ref="courseFormRef"
        :model="courseForm"
        :rules="courseRules"
        label-width="100px"
        class="course-form"
      >
        <el-form-item label="课程名称" prop="title">
          <el-input
            v-model="courseForm.title"
            placeholder="请输入课程名称"
            :maxlength="100"
          />
        </el-form-item>

        <el-form-item label="课程分类" prop="categoryId">
          <el-select
            v-model="courseForm.categoryId"
            placeholder="请选择课程分类"
          >
            <el-option
              v-for="category in categories"
              :key="category.id"
              :label="category.name"
              :value="category.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="讲师" prop="teacher">
          <el-input
            v-model="courseForm.teacher"
            placeholder="请输入讲师名称"
            :maxlength="50"
          />
        </el-form-item>

        <el-form-item label="价格" prop="price">
          <el-input
            v-model="courseForm.price"
            placeholder="请输入课程价格"
            type="number"
            :min="0"
          />
        </el-form-item>

        <el-form-item label="课程封面" prop="image">
          <el-upload
            class="avatar-uploader"
            action="/api/upload"
            :show-file-list="false"
            :on-success="handleImageUploadSuccess"
            :before-upload="beforeImageUpload"
          >
            <img
              v-if="courseForm.image"
              :src="courseForm.image"
              class="avatar"
            >
            <div v-else class="uploader-trigger">
              <el-icon><Plus /></el-icon>
              <div class="text">上传封面</div>
            </div>
          </el-upload>
        </el-form-item>

        <el-form-item label="课程简介" prop="description">
          <el-input
            v-model="courseForm.description"
            placeholder="请输入课程简介"
            type="textarea"
            :rows="4"
          />
        </el-form-item>

        <el-form-item label="课程状态" prop="status">
          <el-radio-group v-model="courseForm.status">
            <el-radio label="draft">草稿</el-radio>
            <el-radio label="published">已发布</el-radio>
            <el-radio label="archived">已归档</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm">提交</el-button>
          <el-button @click="resetForm">重置</el-button>
          <el-button @click="navigateBack">返回</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { courseAPI, courseCategoryAPI } from '@/api/index';
import { Plus } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

// 路由
const router = useRouter();

// 状态管理
const courseFormRef = ref(null);
const categories = ref([]);
const loading = ref(false);

// 表单数据
const courseForm = reactive({
  title: '',
  categoryId: '',
  teacher: '',
  price: 0,
  image: '',
  description: '',
  status: 'draft'
});

// 表单规则
const courseRules = {
  title: [
    { required: true, message: '请输入课程名称', trigger: 'blur' },
    { max: 100, message: '课程名称不能超过100个字符', trigger: 'blur' }
  ],
  categoryId: [
    { required: true, message: '请选择课程分类', trigger: 'change' }
  ],
  teacher: [
    { required: true, message: '请输入讲师名称', trigger: 'blur' },
    { max: 50, message: '讲师名称不能超过50个字符', trigger: 'blur' }
  ],
  price: [
    { required: true, message: '请输入课程价格', trigger: 'blur' },
    { type: 'number', message: '课程价格必须是数字', trigger: 'blur' },
    { min: 0, message: '课程价格不能小于0', trigger: 'blur' }
  ]
};

// 获取课程分类
const fetchCategories = async () => {
  try {
    const response = await courseCategoryAPI.getAllCategories();
    categories.value = response || [];
  } catch (error) {
    console.error('获取课程分类失败:', error);
  }
};

// 提交表单
const submitForm = async () => {
  try {
    await courseFormRef.value.validate();
    loading.value = true;

    const response = await courseAPI.createCourse(courseForm);
    ElMessage.success('课程创建成功');
    router.push('/admin/courses');
  } catch (error) {
    if (error.name !== 'ValidationError') {
      console.error('创建课程失败:', error);
      ElMessage.error('创建课程失败');
    }
  } finally {
    loading.value = false;
  }
};

// 重置表单
const resetForm = () => {
  courseFormRef.value.resetFields();
};

// 返回课程列表
const navigateBack = () => {
  router.push('/admin/courses');
};

// 图片上传前检查
const beforeImageUpload = (file) => {
  const isJPG = file.type === 'image/jpeg' || file.type === 'image/png';
  const isLt2M = file.size / 1024 / 1024 < 2;

  if (!isJPG) {
    ElMessage.error('上传图片只能是 JPG/PNG 格式!');
  }
  if (!isLt2M) {
    ElMessage.error('上传图片大小不能超过 2MB!');
  }

  return isJPG && isLt2M;
};

// 图片上传成功处理
const handleImageUploadSuccess = (response) => {
  if (response.success) {
    courseForm.image = response.data.url;
    ElMessage.success('图片上传成功');
  } else {
    ElMessage.error('图片上传失败');
  }
};

// 初始加载
onMounted(() => {
  fetchCategories();
});
</script>

<style scoped>
.admin-course-create-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.course-form {
  margin-top: 20px;
}

.avatar-uploader {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.avatar {
  width: 200px;
  height: 150px;
  object-fit: cover;
  border-radius: 4px;
}

.uploader-trigger {
  width: 200px;
  height: 150px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: 1px dashed #d9d9d9;
  border-radius: 4px;
  cursor: pointer;
}

.text {
  margin-top: 10px;
  color: #909399;
}
</style>
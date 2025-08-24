<template>
  <div class="admin-course-list-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>课程管理</span>
          <el-button
            type="primary"
            size="small"
            @click="navigateToCreateCourse"
            icon="Plus"
          >
            创建课程
          </el-button>
        </div>
      </template>

      <div class="search-bar">
        <el-input
          v-model="searchQuery"
          placeholder="搜索课程名称"
          prefix-icon="Search"
          class="search-input"
        />
        <el-select
          v-model="selectedCategory"
          placeholder="选择分类"
          class="category-select"
        >
          <el-option label="全部" value="" />
          <el-option
            v-for="category in categories"
            :key="category.id"
            :label="category.name"
            :value="category.id"
          />
        </el-select>
        <el-button type="primary" @click="fetchCourses">搜索</el-button>
      </div>

      <el-table
        v-loading="loading"
        :data="courses"
        border
        style="width: 100%"
      >
        <el-table-column prop="id" label="课程ID" width="80" />
        <el-table-column prop="title" label="课程名称" width="200" />
        <el-table-column
          prop="categoryId"
          label="课程分类"
          width="120"
          :formatter="formatCategory"
        />
        <el-table-column prop="teacher" label="讲师" width="120" />
        <el-table-column prop="price" label="价格" width="100" />
        <el-table-column prop="studentCount" label="学习人数" width="100" />
        <el-table-column prop="rating" label="评分" width="80" />
        <el-table-column prop="status" label="状态" width="100" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button
              type="primary"
              size="small"
              @click="navigateToEditCourse(scope.row.id)"
            >
              编辑
            </el-button>
            <el-button
              type="danger"
              size="small"
              @click="deleteCourse(scope.row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { courseAPI, courseCategoryAPI } from '@/api/index';
import { ElMessageBox, ElMessage } from 'element-plus';

// 路由
const router = useRouter();

// 状态管理
const courses = ref([]);
const categories = ref([]);
const searchQuery = ref('');
const selectedCategory = ref('');
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const loading = ref(false);

// 获取课程列表
const fetchCourses = async () => {
  try {
    loading.value = true;
    const params = {
      page: currentPage.value,
      pageSize: pageSize.value
    };

    if (searchQuery.value) {
      params.title = searchQuery.value;
    }

    if (selectedCategory.value) {
      params.categoryId = selectedCategory.value;
    }

    const response = await courseAPI.getCourseList(params);
    // 数据验证和处理
    console.log('课程列表API响应:', response);
    if (response && typeof response === 'object') {
      // 确保courses.value始终是数组
      courses.value = Array.isArray(response.data) ? response.data : [];
      // 确保total.value是数字
      total.value = typeof response.total === 'number' ? response.total : 0;
    } else {
      console.error('API响应格式不符合预期:', response);
      courses.value = [];
      total.value = 0;
    }
  } catch (error) {
    console.error('获取课程列表失败:', error);
    courses.value = [];
    total.value = 0;
  } finally {
    loading.value = false;
  }
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

// 格式化分类名称
const formatCategory = (row) => {
  const category = categories.value.find(c => c.id === row.categoryId);
  return category ? category.name : '未分类';
};

// 处理分页变化
const handleSizeChange = (size) => {
  pageSize.value = size;
  fetchCourses();
};

const handleCurrentChange = (current) => {
  currentPage.value = current;
  fetchCourses();
};

// 导航到创建课程页面
const navigateToCreateCourse = () => {
  router.push('/admin/dashboard/courses/create');
};

// 导航到编辑课程页面
const navigateToEditCourse = (id) => {
  router.push(`/admin/dashboard/courses/edit/${id}`);
};

// 添加编辑课程路由的支持
router.addRoute('adminDashboard', {
  path: 'courses/edit/:id',
  name: 'adminEditCourse',
  component: () => import('../views/AdminCourseCreateView.vue'),
  meta: { requiresAuth: true, requiresAdmin: true }
});

// 删除课程
const deleteCourse = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这门课程吗？', '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    });

    await courseAPI.deleteCourse(id);
    ElMessage.success('课程删除成功');
    fetchCourses();
  } catch (error) {
    if (error.name !== 'Cancel') {
      console.error('删除课程失败:', error);
      ElMessage.error('删除课程失败');
    }
  }
};

// 初始加载
onMounted(() => {
  fetchCategories();
  fetchCourses();
});
</script>

<style scoped>
.admin-course-list-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-bar {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.search-input {
  flex: 1;
  min-width: 200px;
}

.category-select {
  min-width: 150px;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
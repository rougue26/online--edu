<template>
  <div class="course-list-container">
    <div class="search-and-filter">
      <el-input
        v-model="searchQuery"
        placeholder="搜索课程..."
        prefix-icon="Search"
        class="search-input"
      /> 
      <el-select
        v-model="selectedCategory"
        placeholder="选择分类"
        class="category-select"
        @change="handleCategoryChange"
      >
        <el-option value="">全部分类</el-option>
        <el-option v-for="category in processedCategories" :key="category.id" :value="category.id">
          {{ category.displayName }}
        </el-option>
      </el-select>
      <el-select
        v-model="sortOption"
        placeholder="排序方式"
        class="sort-select"
      >
        <el-option label="价格从低到高" value="price-asc" /> 
        <el-option label="价格从高到低" value="price-desc" /> 
        <el-option label="最新上线" value="newest" /> 
        <el-option label="人气最高" value="popular" /> 
      </el-select>
    </div>

    <div class="course-list">
      <el-row :gutter="20">
        <el-col :span="6" v-for="course in filteredCourses" :key="course.id">
          <el-card :body-style="{ padding: '0px' }" @click="navigateToCourseDetail(course.id)">
            <img :src="course.image" class="course-image" @error="handleImageError($event, course)">
            <div class="course-info">
              <div class="course-title">{{ course.title }}</div>
              <div class="course-teacher">{{ course.teacher }}</div>
              <div class="course-rating">
                <el-rate
                  :model-value="course.rating"
                  disabled
                  :max="5"
                  size="small"
                /> 
                <span class="rating-count">({{ course.ratingCount }})</span>
              </div>
              <div class="course-price">{{ course.price }}</div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="pagination-container">
      <el-pagination
        v-model:current-page="currentPage"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        v-loading="loading"
      /> 
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watchEffect } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { ElSelect, ElOption, ElInput, ElPagination } from 'element-plus';
import { courseAPI, courseCategoryAPI } from '@/api/index';

// 路由和导航
const router = useRouter();
const route = useRoute();

// 状态管理
const searchQuery = ref('');
const selectedCategory = ref('');
const sortOption = ref('');
const currentPage = ref(1);
const pageSize = ref(8);
const categories = ref([]);
const allCourses = ref([]);
const loading = ref(false);
const total = ref(0);

// 计算属性，处理分类数据
const processedCategories = computed(() => {
  // 确保categories.value是数组
  if (!Array.isArray(categories.value)) {
    return [];
  }
  return categories.value.map((category, index) => ({
    ...category,
    displayName: category.name || `分类${index + 1}`
  }));
});

// 计算属性 - 过滤和排序后的课程
const filteredCourses = computed(() => {
  // 实际过滤和排序已在API请求中处理
  return allCourses.value;
});

// 处理图片加载错误
const handleImageError = (event, course) => {
  event.target.src = '/photos/300/200?random=' + course.id;
};

// 获取课程分类
const fetchCategories = async () => {
  try {
    const response = await courseCategoryAPI.getAllCategories();
    categories.value = response.data || [];
  } catch (error) {
    console.error('获取课程分类失败:', error);
  }
};

// 获取课程列表
const fetchCourses = async () => {
  loading.value = true;
  try {
    const params = {
      page: currentPage.value,
      pageSize: pageSize.value,
      categoryId: selectedCategory.value || undefined,
      sortBy: sortOption.value.split('-')[0] || undefined,
      sortOrder: sortOption.value.split('-')[1] || 'asc'
    };

    if (searchQuery.value) {
      params.search = searchQuery.value;
    }

    const response = await courseAPI.getCourseList(params);
    allCourses.value = response.data || [];
    total.value = response.total || 0;
  } catch (error) {
    console.error('获取课程列表失败:', error);
  } finally {
    loading.value = false;
  }
};

// 从路由参数获取分类ID
const categoryId = route.query.category || '';
if (categoryId) {
  selectedCategory.value = categoryId;
}

// 监听路由变化
watchEffect(() => {
  const newCategoryId = route.query.category || '';
  if (newCategoryId !== selectedCategory.value) {
    selectedCategory.value = newCategoryId;
  }
});

// 监听分类、排序和搜索变化
watchEffect(() => {
  // 重置当前页
  currentPage.value = 1;
  fetchCourses();
});

// 生命周期钩子
onMounted(() => {
  fetchCategories();
  fetchCourses();
});

// 处理分页
const handleSizeChange = (size) => {
  pageSize.value = size;
  currentPage.value = 1;
};

const handleCurrentChange = (current) => {
  currentPage.value = current;
};

// 导航到课程详情
const navigateToCourseDetail = (courseId) => {
  if (courseId) {
    router.push({ name: 'courseDetail', params: { id: courseId } });
  } else {
    console.error('无法导航到课程详情：缺少课程ID');
  }
};

// 处理分类变化
const handleCategoryChange = (categoryId) => {
  router.push({ name: 'courses', query: { category: categoryId || '' } });
};
</script>

<style scoped>
.course-list-container {
  padding: 20px 50px;
}

/* 分类选择器样式 */
.category-select {
  width: 180px;
}

/* 排序选择器样式 */
.sort-select {
  width: 180px;
}

.search-and-filter {
  display: flex;
  gap: 20px;
  margin-bottom: 30px;
  flex-wrap: wrap;
}

.search-input {
  flex: 1;
  min-width: 300px;
}

.category-select, .sort-select {
  width: 180px;
}

.course-image {
  width: 100%;
  height: 160px;
  object-fit: cover;
  cursor: pointer;
}

.course-info {
  padding: 10px;
}

.course-title {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
  cursor: pointer;
}

.course-teacher {
  color: #666;
  font-size: 14px;
  margin-bottom: 5px;
}

.course-rating {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
}

.rating-count {
  margin-left: 5px;
  color: #999;
  font-size: 12px;
}

.course-price {
  color: #f56c6c;
  font-size: 16px;
  font-weight: bold;
}

.pagination-container {
  margin-top: 40px;
  display: flex;
  justify-content: center;
}
</style>
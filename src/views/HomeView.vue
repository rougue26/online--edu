<template>
  <div class="home-container">
    <!-- 轮播图 -->
    <el-carousel height="400px" :interval="5000">
      <el-carousel-item>
        <img src="/photos/1200/400?random=1" alt="轮播图1" class="carousel-image">
      </el-carousel-item>
      <el-carousel-item>
        <img src="/photos/1200/400?random=2" alt="轮播图2" class="carousel-image">
      </el-carousel-item>
      <el-carousel-item>
        <img src="/photos/1200/400?random=3" alt="轮播图3" class="carousel-image">
      </el-carousel-item>
    </el-carousel>

    <!-- 课程分类 -->
    <div class="category-section">
      <h2>热门分类</h2>
      <template v-if="loadingCategories">
        <el-row :gutter="20">
          <el-col :span="6" v-for="i in 4" :key="i">
            <el-card :body-style="{ padding: '0px' }">
              <el-skeleton :loading="true">
                <template #image>
                  <div class="category-image-skeleton"></div>
                </template>
                <template #text>
                  <div class="category-info-skeleton">
                    <div class="category-title-skeleton"></div>
                    <div class="course-count-skeleton"></div>
                  </div>
                </template>
              </el-skeleton>
            </el-card>
          </el-col>
        </el-row>
      </template>
      <template v-else-if="categories.length > 0">
        <el-row :gutter="20">
          <el-col :span="6" v-for="category in categories" :key="category.id">
            <el-card :body-style="{ padding: '0px' }" @click="navigateToCourses(category.id)">
              <img :src="category.image" class="category-image">
              <div class="category-info">
                <div class="category-title">{{ category.name }}</div>
                <div class="course-count">{{ category.courseCount }} 门课程</div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </template>
      <template v-else>
        <div class="no-data">暂无分类数据</div>
      </template>
    </div>

    <!-- 热门课程 -->
    <div class="popular-courses">
      <h2>热门推荐</h2>
      <template v-if="loadingCourses">
        <el-row :gutter="20">
          <el-col :span="6" v-for="i in 4" :key="i">
            <el-card :body-style="{ padding: '0px' }">
              <el-skeleton :loading="true">
                <template #image>
                  <div class="course-image-skeleton"></div>
                </template>
                <template #text>
                  <div class="course-info-skeleton">
                    <div class="course-title-skeleton"></div>
                    <div class="course-teacher-skeleton"></div>
                    <div class="course-price-skeleton"></div>
                  </div>
                </template>
              </el-skeleton>
            </el-card>
          </el-col>
        </el-row>
      </template>
      <template v-else-if="popularCourses.length > 0">
        <el-row :gutter="20">
          <el-col :span="6" v-for="course in popularCourses" :key="course.id">
            <el-card :body-style="{ padding: '0px' }" @click="navigateToCourseDetail(course.id)">
              <img :src="course.image" class="course-image">
              <div class="course-info">
                <div class="course-title">{{ course.title }}</div>
                <div class="course-teacher">{{ course.teacher }}</div>
                <div class="course-price">{{ course.price }}</div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </template>
      <template v-else>
        <div class="no-data">暂无推荐课程</div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { courseCategoryAPI, courseAPI } from '@/api/index';
import { ElMessage } from 'element-plus';

const router = useRouter();
const categories = ref([]);
const popularCourses = ref([]);
const loadingCategories = ref(true);
const loadingCourses = ref(true);

// 获取课程分类
const fetchCategories = async () => {
  try {
    loadingCategories.value = true;
    const response = await courseCategoryAPI.getAllCategories();
    categories.value = Array.isArray(response.data) ? response.data : [];
    // 为每个分类添加图片和课程数量（实际项目中这些数据应该来自API）
    categories.value = categories.value.map(category => ({
      ...category,
      image: `/photos/400/200?random=${category.id + 3}`,},{
      courseCount: Math.floor(Math.random() * 100) + 50
    }));
  } catch (error) {
    console.error('获取课程分类失败:', error);
    ElMessage.error('获取课程分类失败');
    // 提供默认数据以防API失败
    categories.value = [
      { id: 1, name: '编程开发', image: '/photos/400/200?random=4', courseCount: 120 },
      { id: 2, name: '设计创意', image: '/photos/400/200?random=5', courseCount: 86 },
      { id: 3, name: '职业技能', image: '/photos/400/200?random=6', courseCount: 95 },
      { id: 4, name: '语言学习', image: '/photos/400/200?random=7', courseCount: 78 }
    ];
  } finally {
    loadingCategories.value = false;
  }
};

// 获取热门课程
const fetchPopularCourses = async () => {
  try {
    loadingCourses.value = true;
    // 获取热门课程（假设API支持sort=popular参数）
    const response = await courseAPI.getCourseList({ sort: 'popular', limit: 4 });
    // 确保popularCourses.value始终是一个数组
    popularCourses.value = Array.isArray(response.data) ? response.data : [];
    // 为课程添加图片（实际项目中这些数据应该来自API）
    popularCourses.value = popularCourses.value.map((course, index) => ({
      ...course,
      image: `/photos/400/200?random=${index + 8}`,},{
      price: course.price ? `¥${course.price}` : '免费'
    }));
  } catch (error) {
    console.error('获取热门课程失败:', error);
    ElMessage.error('获取热门课程失败');
    // 提供默认数据以防API失败
    popularCourses.value = [
      { id: 1, title: 'Vue.js从入门到精通', teacher: '张老师', price: '¥299', image: '/photos/400/200?random=8' },
      { id: 2, title: 'Python数据分析', teacher: '李老师', price: '¥399', image: '/photos/400/200?random=9' },
      { id: 3, title: 'UI设计实战', teacher: '王老师', price: '¥499', image: '/photos/400/200?random=10' },
      { id: 4, title: 'Web前端开发', teacher: '赵老师', price: '¥599', image: '/photos/400/200?random=11' }
    ];
  } finally {
    loadingCourses.value = false;
  }
};

// 导航到课程列表
const navigateToCourses = (categoryId) => {
  router.push({ name: 'courses', query: { category: categoryId } });
};

// 导航到课程详情
const navigateToCourseDetail = (courseId) => {
  router.push({ name: 'courseDetail', params: { id: courseId } });
};

onMounted(() => {
  fetchCategories();
  fetchPopularCourses();
});
</script>

<style scoped>
.home-container {
  padding: 20px 50px;
}

.carousel-image {
  width: 100%;
  height: 400px;
  object-fit: cover;
}

.category-section,
.popular-courses {
  margin-top: 40px;
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

.category-image {
  width: 100%;
  height: 160px;
  object-fit: cover;
  cursor: pointer;
}

.category-info {
  padding: 10px;
}

.category-title {
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
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

.course-price {
  color: #f56c6c;
  font-size: 16px;
  font-weight: bold;
}

/* 骨架屏样式 */
.category-image-skeleton,
.course-image-skeleton {
  width: 100%;
  height: 160px;
  background-color: #f5f5f5;
}

.category-info-skeleton,
.course-info-skeleton {
  padding: 10px;
}

.category-title-skeleton,
.course-title-skeleton {
  width: 80%;
  height: 20px;
  background-color: #f5f5f5;
  margin-bottom: 10px;
  border-radius: 4px;
}

.course-count-skeleton,
.course-teacher-skeleton {
  width: 60%;
  height: 16px;
  background-color: #f5f5f5;
  margin-bottom: 10px;
  border-radius: 4px;
}

.course-price-skeleton {
  width: 40%;
  height: 18px;
  background-color: #f5f5f5;
  border-radius: 4px;
}

/* 无数据状态样式 */
.no-data {
  text-align: center;
  padding: 30px 0;
  color: #999;
  font-size: 16px;
  background-color: #f9f9f9;
  border-radius: 8px;
}
</style>
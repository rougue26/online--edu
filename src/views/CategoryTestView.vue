<template>
  <div class="category-test-container">
    <h1>分类测试页面</h1>
    
    <!-- 直接显示原始数据 -->
    <div class="raw-data">
      <h3>原始分类数据:</h3>
      <pre>{{ JSON.stringify(categories, null, 2) }}</pre>
    </div>

    <!-- 显示处理后的分类数据 -->
    <div class="processed-data">
      <h3>处理后的分类数据:</h3>
      <pre>{{ JSON.stringify(processedCategories, null, 2) }}</pre>
    </div>
    
    <!-- 使用Vue指令显示 -->
    <div class="vue-display">
      <h3>Vue指令显示:</h3>
      <ul>
        <li v-for="category in processedCategories" :key="category.id">
          {{ category.id }} - {{ category.displayName }}
        </li>
      </ul>
    </div>
    
    <!-- 使用原生select -->
    <div class="native-select-container">
      <h3>原生select:</h3>
      <select class="test-select">
        <option value="">选择分类</option>
        <option v-for="category in processedCategories" :key="category.id" :value="category.id">
          {{ category.displayName }} (ID: {{ category.id }})
        </option>
      </select>
    </div>
    
    <!-- 使用Element Plus select -->
    <div class="el-select-container">
      <h3>Element Plus select:</h3>
      <el-select v-model="selectedCategory" placeholder="选择分类" style="width: 200px;">
        <el-option
          v-for="category in processedCategories"
          :key="category.id"
          :value="category.id"
          :label="category.displayName"
        ></el-option>
      </el-select>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue';
import { ElSelect, ElOption } from 'element-plus';
import { courseCategoryAPI } from '@/api';

export default {
  name: 'CategoryTestView',
  components: {
    ElSelect,
    ElOption
  },
  setup() {
    const categories = ref([]);
    const selectedCategory = ref('');

    // 计算属性，处理分类数据
    const processedCategories = computed(() => {
      return categories.value.map((category, index) => ({
        ...category,
        displayName: category.name || `分类${index + 1}`
      }));
    });

    const fetchCategories = async () => {
      try {
        console.log('开始获取分类数据');
        const response = await courseCategoryAPI.getAllCategories();
        console.log('分类数据响应:', response);
        console.log('分类数据内容:', response.data);
        categories.value = response.data;
        console.log('分类数据已设置:', categories.value);
        console.log('处理后的分类数据:', processedCategories.value);
      } catch (error) {
        console.error('获取分类数据失败:', error);
      }
    };

    onMounted(() => {
      fetchCategories();
    });

    return {
      categories,
      selectedCategory
    };
  }
};
</script>

<style scoped>
.category-test-container {
  padding: 20px;
}

.raw-data {
  margin-bottom: 20px;
  padding: 10px;
  background-color: #f5f5f5;
  border-radius: 4px;
}

.processed-data {
  margin-bottom: 20px;
  padding: 10px;
  background-color: #e6f7ff;
  border-radius: 4px;
}

.vue-display {
  margin-bottom: 20px;
  padding: 10px;
  border: 1px solid #eee;
  border-radius: 4px;
}

.native-select-container {
  margin-bottom: 20px;
  padding: 10px;
  border: 1px solid #eee;
  border-radius: 4px;
}

.el-select-container {
  margin-bottom: 20px;
  padding: 10px;
  border: 1px solid #eee;
  border-radius: 4px;
}

.test-select {
  width: 200px;
  padding: 8px 12px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  font-size: 14px;
}
</style>
<template>
  <div class="admin-user-list-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>用户列表</span>
          <el-button type="primary" size="small" @click="handleAddUser">新增用户</el-button>
        </div>
      </template>
      <el-table :data="users" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80"></el-table-column>
        <el-table-column prop="username" label="用户名" width="180"></el-table-column>
        <el-table-column prop="email" label="邮箱"></el-table-column>
        <el-table-column prop="role" label="角色" width="100"></el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="180"></el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleEditUser(scope.row)">编辑</el-button>
            <el-button type="danger" size="small" @click="handleDeleteUser(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          :page-sizes="[10, 20, 50, 100]"
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
import { userAPI } from '@/api/index';
import { ElMessageBox, ElMessage } from 'element-plus';

// 状态管理
const users = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const loading = ref(false);

// 路由
const router = useRouter();

// 获取用户列表
const fetchUsers = async () => {
  try {
    loading.value = true;
    console.log('开始获取用户列表，页码:', currentPage.value, '每页数量:', pageSize.value);
    const res = await userAPI.getUserList({
      page: currentPage.value,
      pageSize: pageSize.value
    });
    // 数据验证和处理
    console.log('用户列表API响应:', res);
    if (res && typeof res === 'object') {
      // 确保users.value始终是数组
      users.value = Array.isArray(res.list) ? res.list : [];
      // 确保total.value是数字
      total.value = typeof res.total === 'number' ? res.total : 0;
    } else {
      console.error('API响应格式不符合预期:', res);
      users.value = [];
      total.value = 0;
      ElMessage.error('获取用户列表失败: 响应格式不正确');
    }
  } catch (error) {
    console.error('获取用户列表失败:', error);
    // 详细错误信息处理
    let errorMessage = '未知错误';
    if (error.message) {
      errorMessage = error.message;
    } else if (error.response && error.response.data && error.response.data.message) {
      errorMessage = error.response.data.message;
    }
    ElMessage.error('获取用户列表失败: ' + errorMessage);
    users.value = [];
    total.value = 0;
  } finally {
    loading.value = false;
  }
};

// 处理分页大小变化
const handleSizeChange = (size) => {
  pageSize.value = size;
  fetchUsers();
};

// 处理当前页变化
const handleCurrentChange = (current) => {
  currentPage.value = current;
  fetchUsers();
};

// 添加用户管理相关路由
router.addRoute('adminDashboard', {
  path: 'users/create',
  name: 'adminCreateUser',
  component: () => import('../views/AdminUserCreateView.vue'),
  meta: { requiresAuth: true, requiresAdmin: true }
});

router.addRoute('adminDashboard', {
  path: 'users/:id/edit',
  name: 'adminEditUser',
  component: () => import('../views/AdminUserEditView.vue'),
  meta: { requiresAuth: true, requiresAdmin: true }
});

// 处理新增用户
const handleAddUser = () => {
  router.push('/admin/dashboard/users/create');
};

// 处理编辑用户
const handleEditUser = (user) => {
  router.push(`/admin/dashboard/users/${user.id}/edit`);
};

// 处理删除用户
const handleDeleteUser = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除该用户吗？', '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    });

    await userAPI.deleteUser(id);
    ElMessage.success('用户删除成功');
    fetchUsers(); // 重新获取用户列表
  } catch (error) {
    if (error.name !== 'Cancel') {
      console.error('删除用户失败:', error);
      ElMessage.error('删除用户失败: ' + (error.message || '未知错误'));
    }
  }
};

// 初始加载
onMounted(() => {
  fetchUsers();
});
</script>

<style scoped>
.admin-user-list-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
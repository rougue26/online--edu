<template>
  <div class="admin-role-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>角色权限管理</span>
          <el-button type="primary" size="small" @click="handleAddRole">新增角色</el-button>
        </div>
      </template>
      <el-table :data="roles" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80"></el-table-column>
        <el-table-column prop="name" label="角色名称" width="180"></el-table-column>
        <el-table-column prop="description" label="角色描述"></el-table-column>
        <el-table-column label="权限" width="200">
          <template #default="scope">
            <el-tag v-for="permission in scope.row.permissions" :key="permission" size="small" style="margin-right: 5px;">{{ permission }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleEditRole(scope.row)">编辑</el-button>
            <el-button type="danger" size="small" @click="handleDeleteRole(scope.row.id)">删除</el-button>
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

<script>
import { ref, onMounted } from 'vue';
// 假设存在roleAPI
// import { roleAPI } from '../api/index';

export default {
  name: 'AdminRoleView',
  setup() {
    const roles = ref([]);
    const currentPage = ref(1);
    const pageSize = ref(10);
    const total = ref(0);
    const loading = ref(false);

    // 获取角色列表
    const fetchRoles = async () => {
      try {
        loading.value = true;
        // 模拟数据，实际项目中应替换为API调用
        // const res = await roleAPI.getRoleList({
        //   page: currentPage.value,
        //   pageSize: pageSize.value
        // });
        // 模拟数据
        roles.value = [
          {
            id: 1,
            name: '超级管理员',
            description: '拥有系统所有权限',
            permissions: ['全部权限']
          },
          {
            id: 2,
            name: '课程管理员',
            description: '负责课程管理',
            permissions: ['课程管理', '视频管理']
          },
          {
            id: 3,
            name: '用户管理员',
            description: '负责用户管理',
            permissions: ['用户管理']
          }
        ];
        total.value = roles.value.length;
      } catch (error) {
        console.error('获取角色列表失败:', error);
      } finally {
        loading.value = false;
      }
    };

    // 处理分页大小变化
    const handleSizeChange = (size) => {
      pageSize.value = size;
      fetchRoles();
    };

    // 处理当前页变化
    const handleCurrentChange = (current) => {
      currentPage.value = current;
      fetchRoles();
    };

    // 处理新增角色
    const handleAddRole = () => {
      alert('新增角色功能待实现');
    };

    // 处理编辑角色
    const handleEditRole = (role) => {
      alert(`编辑角色 ${role.name} 功能待实现`);
    };

    // 处理删除角色
    const handleDeleteRole = (id) => {
      alert(`删除角色 ID: ${id} 功能待实现`);
    };

    onMounted(() => {
      fetchRoles();
    });

    return {
      roles,
      currentPage,
      pageSize,
      total,
      loading,
      handleSizeChange,
      handleCurrentChange,
      handleAddRole,
      handleEditRole,
      handleDeleteRole
    };
  }
};
</script>

<style scoped>
.admin-role-container {
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
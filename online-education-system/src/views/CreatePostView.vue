<template>
  <div class="create-post-container">
    <el-card class="create-post-card">
      <template #header>
        <div class="card-header">
          <span>发布新帖子</span>
        </div>
      </template>
      <el-form ref="postForm" :model="postForm" :rules="rules" class="post-form">
        <el-form-item prop="title">
          <el-input v-model="postForm.title" placeholder="请输入帖子标题"></el-input>
        </el-form-item>
        <el-form-item prop="category">
          <el-select v-model="postForm.category" placeholder="请选择帖子分类">
            <el-option v-for="category in categories" :key="category.value" :label="category.label" :value="category.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="content">
          <div class="editor-container">
            <div class="editor-toolbar">
              <el-button type="text" @click="format('bold')"><b>B</b></el-button>
              <el-button type="text" @click="format('italic')"><i>I</i></el-button>
              <el-button type="text" @click="format('underline')"><u>U</u></el-button>
              <el-button type="text" @click="format('strikeThrough')"><s>S</s></el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" @click="format('insertOrderedList')">1. 列表</el-button>
              <el-button type="text" @click="format('insertUnorderedList')">• 列表</el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" @click="format('justifyLeft')"><i class="el-icon-align-left"></i></el-button>
              <el-button type="text" @click="format('justifyCenter')"><i class="el-icon-align-center"></i></el-button>
              <el-button type="text" @click="format('justifyRight')"><i class="el-icon-align-right"></i></el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" @click="insertImage()"><i class="el-icon-picture"></i></el-button>
            </div>
            <div
              ref="editor"
              class="editor-content"
              contenteditable="true"
              @input="onContentChange"
              placeholder="请输入帖子内容..."
            ></div>
          </div>
        </el-form-item>
        <el-form-item>
          <div class="form-actions">
            <el-button @click="resetForm">重置</el-button>
            <el-button type="primary" @click="submitForm">发布帖子</el-button>
          </div>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { postAPI } from '@/api/index'
export default {
  


  name: 'CreatePostView',
  data() {
    return {
      postForm: {
        title: '',
        category: '',
        content: ''
      },
      categories: [
        { label: '学习方法', value: '学习方法' },
        { label: '技术讨论', value: '技术讨论' },
        { label: '求职经验', value: '求职经验' },
        { label: '资源分享', value: '资源分享' },
        { label: '其他', value: '其他' }
      ],
      rules: {
        title: [
          { required: true, message: '请输入帖子标题', trigger: 'blur' },
          { min: 5, max: 50, message: '标题长度在 5 到 50 个字符之间', trigger: 'blur' }
        ],
        category: [
          { required: true, message: '请选择帖子分类', trigger: 'change' }
        ],
        content: [
          { required: true, message: '请输入帖子内容', trigger: 'blur' },
          { min: 10, message: '内容至少 10 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  mounted() {
    // 初始化编辑器内容
    this.$refs.editor.innerHTML = ''
  },
  methods: {
    onContentChange() {
      this.postForm.content = this.$refs.editor.innerHTML
    },
    format(command) {
      document.execCommand(command, false, null)
      this.$refs.editor.focus()
      this.onContentChange()
    },
    insertImage() {
      const imageUrl = prompt('请输入图片URL:')
      if (imageUrl) {
        document.execCommand('insertImage', false, imageUrl)
        this.onContentChange()
      }
    },
    submitForm() {
      this.$refs.postForm.validate(valid => {
        if (valid) {
          // 提交到API
          const postData = {
            title: this.postForm.title,
            content: this.postForm.content,
            category: this.postForm.category
          }
          
          postAPI.createPost(postData)
            .then(response => {
              this.$message.success('帖子发布成功')
              setTimeout(() => {
                this.$router.push({ name: 'community' })
              }, 1500)
            })
            .catch(error => {
              this.$message.error('帖子发布失败: ' + (error.response?.data?.message || error.message))
            })
        } else {
          return false
        }
      })
    },
    resetForm() {
      this.$refs.postForm.resetFields()
      this.$refs.editor.innerHTML = ''
      this.postForm.content = ''
    }
  }
}
</script>

<style scoped>
.create-post-container {
  padding: 20px;
}

.card-header {
  font-size: 18px;
  font-weight: bold;
}

.post-form {
  margin-top: 20px;
}

.editor-container {
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  overflow: hidden;
}

.editor-toolbar {
  background-color: #f5f7fa;
  padding: 10px;
  border-bottom: 1px solid #dcdfe6;
  display: flex;
  align-items: center;
}

.editor-content {
  min-height: 300px;
  padding: 15px;
  outline: none;
}

.editor-content:focus {
  border-color: #409eff;
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>
<template>
  <div class="video-upload-container">
    <el-card class="video-upload-card">
      <template #header>
        <div class="card-header">
          <span>上传视频</span>
        </div>
      </template>
      <el-form ref="videoForm" :model="videoForm" :rules="rules" class="video-form">
        <el-form-item prop="title">
          <el-input v-model="videoForm.title" placeholder="请输入视频标题"></el-input>
        </el-form-item>
        <el-form-item prop="description">
          <el-input
            v-model="videoForm.description"
            type="textarea"
            placeholder="请输入视频描述"
            rows="4"
          ></el-input>
        </el-form-item>
        <el-form-item prop="category">
          <el-select v-model="videoForm.category" placeholder="请选择视频分类">
            <el-option v-for="category in categories" :key="category.value" :label="category.label" :value="category.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="videoFile">
          <div class="upload-container">
            <el-upload
              ref="upload"
              :auto-upload="false"
              :on-change="handleFileChange"
              :file-list="fileList"
              accept=".mp4,.avi,.mov,.wmv"
              class="upload-demo"
            >
              <el-button slot="trigger" size="small" type="primary">选取视频文件</el-button>
              <div slot="tip" class="el-upload__tip">支持上传 MP4, AVI, MOV, WMV 格式文件，最大 1GB</div>
            </el-upload>
            <div v-if="videoForm.videoFile" class="file-info">
              <div class="file-name">{{ videoForm.videoFile.name }}</div>
              <div class="file-size">{{ formatFileSize(videoForm.videoFile.size) }}</div>
            </div>
          </div>
        </el-form-item>
        <el-form-item prop="coverImage">
          <div class="upload-container">
            <el-upload
              ref="coverUpload"
              :auto-upload="false"
              :on-change="handleCoverChange"
              :file-list="coverFileList"
              accept=".jpg,.jpeg,.png,.gif"
              class="upload-demo"
            >
              <el-button slot="trigger" size="small" type="primary">选取封面图片</el-button>
              <div slot="tip" class="el-upload__tip">支持上传 JPG, JPEG, PNG, GIF 格式文件，建议尺寸 16:9</div>
            </el-upload>
            <div v-if="videoForm.coverImage" class="file-info">
              <div class="file-name">{{ videoForm.coverImage.name }}</div>
              <div class="file-size">{{ formatFileSize(videoForm.coverImage.size) }}</div>
              <img :src="coverPreviewUrl" alt="封面预览" class="cover-preview" v-if="coverPreviewUrl">
            </div>
          </div>
        </el-form-item>
        <el-form-item prop="isPublic">
          <el-radio-group v-model="videoForm.isPublic">
            <el-radio :label="true">公开</el-radio>
            <el-radio :label="false">私有</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <div class="form-actions">
            <el-button @click="resetForm">重置</el-button>
            <el-button type="primary" @click="submitForm">开始上传</el-button>
          </div>
        </el-form-item>
      </el-form>

      <!-- 上传进度条 -->
      <div v-if="uploadProgress > 0 && uploadProgress < 100" class="upload-progress">
        <el-progress :percentage="uploadProgress" status="success"></el-progress>
        <div class="progress-text">{{ progressText }}</div>
      </div>

      <!-- 上传完成提示 -->
      <div v-if="uploadSuccess" class="upload-success">
        <el-alert
          title="上传成功"
          type="success"
          show-icon
        >
          <div slot="desc">
            <p>视频已成功上传！</p>
            <el-button type="primary" size="small" @click="goToVideoDetail">查看视频</el-button>
          </div>
        </el-alert>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'VideoUploadView',
  data() {
    return {
      videoForm: {
        title: '',
        description: '',
        category: '',
        videoFile: null,
        coverImage: null,
        isPublic: true
      },
      categories: [
        { label: '编程开发', value: 'programming' },
        { label: '设计创意', value: 'design' },
        { label: '职业技能', value: 'skills' },
        { label: '学术研究', value: 'academic' },
        { label: '其他', value: 'other' }
      ],
      rules: {
        title: [
          { required: true, message: '请输入视频标题', trigger: 'blur' },
          { min: 2, max: 100, message: '标题长度在 2 到 100 个字符之间', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入视频描述', trigger: 'blur' },
          { min: 10, max: 500, message: '描述长度在 10 到 500 个字符之间', trigger: 'blur' }
        ],
        category: [
          { required: true, message: '请选择视频分类', trigger: 'change' }
        ],
        videoFile: [
          { required: true, message: '请选择视频文件', trigger: 'change' },
          { validator: this.validateVideoFile, trigger: 'change' }
        ],
        coverImage: [
          { required: true, message: '请选择封面图片', trigger: 'change' },
          { validator: this.validateCoverImage, trigger: 'change' }
        ]
      },
      fileList: [],
      coverFileList: [],
      coverPreviewUrl: '',
      uploadProgress: 0,
      progressText: '',
      uploadSuccess: false,
      uploadedVideoId: null
    }
  },
  methods: {
    handleFileChange(file, fileList) {
      this.fileList = fileList
      this.videoForm.videoFile = file.raw
    },
    handleCoverChange(file, fileList) {
      this.coverFileList = fileList
      this.videoForm.coverImage = file.raw
      this.createCoverPreview(file.raw)
    },
    createCoverPreview(file) {
      const reader = new FileReader()
      reader.onload = (e) => {
        this.coverPreviewUrl = e.target.result
      }
      reader.readAsDataURL(file)
    },
    formatFileSize(size) {
      if (size < 1024) {
        return `${size} B`
      } else if (size < 1048576) {
        return `${(size / 1024).toFixed(2)} KB`
      } else if (size < 1073741824) {
        return `${(size / 1048576).toFixed(2)} MB`
      } else {
        return `${(size / 1073741824).toFixed(2)} GB`
      }
    },
    validateVideoFile(rule, value, callback) {
      if (!value) {
        return callback(new Error('请选择视频文件'))
      }

      const validTypes = ['.mp4', '.avi', '.mov', '.wmv']
      const fileExtension = value.name.substring(value.name.lastIndexOf('.')).toLowerCase()
      if (!validTypes.includes(fileExtension)) {
        return callback(new Error('只支持 MP4, AVI, MOV, WMV 格式的视频文件'))
      }

      if (value.size > 1073741824) {
        return callback(new Error('视频文件大小不能超过 1GB'))
      }

      callback()
    },
    validateCoverImage(rule, value, callback) {
      if (!value) {
        return callback(new Error('请选择封面图片'))
      }

      const validTypes = ['.jpg', '.jpeg', '.png', '.gif']
      const fileExtension = value.name.substring(value.name.lastIndexOf('.')).toLowerCase()
      if (!validTypes.includes(fileExtension)) {
        return callback(new Error('只支持 JPG, JPEG, PNG, GIF 格式的图片文件'))
      }

      if (value.size > 10485760) {
        return callback(new Error('图片文件大小不能超过 10MB'))
      }

      callback()
    },
    submitForm() {
      this.$refs.videoForm.validate(valid => {
        if (valid) {
          this.uploadProgress = 0
          this.uploadSuccess = false
          this.progressText = '准备上传...'

          const interval = setInterval(() => {
            this.uploadProgress += Math.floor(Math.random() * 10)
            if (this.uploadProgress >= 100) {
              this.uploadProgress = 100
              clearInterval(interval)
              this.progressText = '上传完成'
              this.uploadedVideoId = Math.floor(Math.random() * 1000)
              this.uploadSuccess = true
            } else if (this.uploadProgress >= 50) {
              this.progressText = '视频处理中...'
            } else {
              this.progressText = '上传中...'
            }
          }, 500)
        } else {
          return false
        }
      })
    },
    resetForm() {
      this.$refs.videoForm.resetFields()
      this.fileList = []
      this.coverFileList = []
      this.coverPreviewUrl = ''
      this.videoForm.videoFile = null
      this.videoForm.coverImage = null
      this.uploadProgress = 0
      this.uploadSuccess = false
      this.uploadedVideoId = null
    },
    goToVideoDetail() {
      this.$router.push({ name: 'videoDetail', params: { id: this.uploadedVideoId } })
    }
  }
}
</script>

<style scoped>
.video-upload-container {
  padding: 20px;
}

.card-header {
  font-size: 18px;
  font-weight: bold;
}

.video-form {
  margin-top: 20px;
}

.upload-container {
  margin-top: 10px;
}

.file-info {
  margin-top: 10px;
  padding: 10px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.file-name {
  font-weight: bold;
}

.file-size {
  color: #666;
  font-size: 14px;
}

.cover-preview {
  max-width: 300px;
  margin-top: 10px;
  border-radius: 4px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}

.upload-progress {
  margin-top: 20px;
}

.progress-text {
  text-align: center;
  margin-top: 10px;
  color: #666;
}

.upload-success {
  margin-top: 20px;
}
</style>
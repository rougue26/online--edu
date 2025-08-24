<template>
  <div class="post-detail-container">
    <template v-if="loading">
      <div class="loading-container">
        <el-skeleton active :count="3" />
      </div>
    </template>
    <template v-else>
      <el-card class="post-detail-card">
        <template #header>
          <div class="post-header">
            <img :src="post.authorAvatar" alt="作者头像" class="author-avatar">
            <div class="author-info">
              <h3 class="author-name">{{ post.authorName }}</h3>
              <p class="post-time">{{ formatDate(post.createdAt) }} · {{ post.category }}</p>
            </div>
          </div>
        </template>
        <div class="post-body">
          <h1 class="post-title">{{ post.title }}</h1>
          <div class="post-content" v-html="post.content"></div>
        </div>
        <div class="post-footer">
          <div class="post-actions">
            <el-button type="primary" icon="el-icon-thumbs-up" :disabled="isLiked" @click="likePost">{{ likeCount }} {{ isLiked ? '已点赞' : '点赞' }}</el-button>
            <el-button icon="el-icon-share">分享</el-button>
          </div>
        </div>
      </el-card>

      <div class="comments-section">
        <h2>评论 ({{ comments.length }})</h2>
        <el-form ref="commentForm" :model="commentForm" class="comment-form">
          <el-form-item prop="content">
            <el-input
              v-model="commentForm.content"
              type="textarea"
              placeholder="写下你的评论..."
              :rows="4"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitComment">发表评论</el-button>
          </el-form-item>
        </el-form>

        <div class="comment-list">
          <el-card v-for="comment in comments" :key="comment.id" class="comment-card">
            <div class="comment-header">
              <img :src="comment.userAvatar" alt="用户头像" class="user-avatar">
              <div class="user-info">
                <h4 class="user-name">{{ comment.userName }}</h4>
                <p class="comment-time">{{ formatDate(comment.createdAt) }}</p>
              </div>
            </div>
            <div class="comment-content">{{ comment.content }}</div>
            <div class="comment-actions">
              <span class="like-action" @click="likeComment(comment.id)">
                <i class="el-icon-thumbs-up"></i> {{ comment.likeCount }} {{ comment.isLiked ? '已点赞' : '点赞' }}
              </span>
            </div>
          </el-card>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import { postAPI } from '@/api/index';
import { ElMessage } from 'element-plus';
export default {
  name: 'PostDetailView',
  props: {
    id: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      loading: true,
      post: {
        id: 1,
        title: '如何有效地学习编程？',
        content: '<p>学习编程是一个需要耐心和实践的过程。以下是一些有效的学习方法：</p><ol><li><strong>设定明确的目标</strong>：确定你想学习的编程语言和领域，设定具体的学习目标。</li><li><strong>从基础开始</strong>：扎实掌握编程基础概念，如变量、数据类型、控制流等。</li><li><strong>大量实践</strong>：通过实际项目练习，将理论知识应用到实践中。</li><li><strong>阅读优秀代码</strong>：学习他人的代码风格和解决问题的方法。</li><li><strong>加入社区</strong>：参与编程社区讨论，向他人请教和分享经验。</li></ol><p>记住，编程学习没有捷径，坚持和实践是成功的关键。</p>',
        authorName: '张教授',
        authorAvatar: 'https://picsum.photos/100/100?random=1',
        createdAt: '2023-11-15T10:30:00',
        category: '学习方法',
        viewCount: 320
      },
      likeCount: 45,
      isLiked: false,
      comments: [
        {
          id: 1,
          userName: '小李',
          userAvatar: 'https://picsum.photos/100/100?random=4',
          content: '非常有用的建议，谢谢分享！',
          createdAt: '2023-11-15T11:20:00',
          likeCount: 8,
          isLiked: false
        },
        {
          id: 2,
          userName: '小王',
          userAvatar: 'https://picsum.photos/100/100?random=5',
          content: '我觉得项目实践真的很重要，通过做项目可以学到很多书本上学不到的东西。',
          createdAt: '2023-11-15T13:45:00',
          likeCount: 12,
          isLiked: false
        }
      ],
      commentForm: {
        content: ''
      }
    }
  },
  created() {
    this.fetchPostDetail();
  },
  methods: {
    async fetchPostDetail() {
      try {
        this.loading = true;
        const response = await postAPI.getPostDetail(this.id);
        if (response && response.code === 200 && response.data) {
          const postData = response.data;
          this.post = {
            ...postData,
            authorAvatar: postData.avatar || `https://picsum.photos/100/100?random=${postData.id}`,
            authorName: postData.user_name || postData.username || '匿名用户',
            createdAt: postData.created_at || postData.createdAt || new Date().toISOString(),
            category: postData.category || '未分类',
            viewCount: postData.view_count || postData.viewCount || 0
          };
          this.likeCount = postData.like_count || postData.likeCount || 0;
          // 获取评论
          this.fetchComments();
        } else {
          ElMessage.error('获取帖子详情失败');
          console.error('获取帖子详情失败:', response);
        }
      } catch (error) {
        ElMessage.error('获取帖子详情失败: ' + (error.message || '未知错误'));
        console.error('获取帖子详情失败:', error);
      } finally {
        this.loading = false;
      }
    },
    fetchComments() {
      // 模拟获取评论
      // 实际应用中应该调用API获取评论
      setTimeout(() => {
        this.comments = [
          {
            id: 1,
            userName: '小李',
            userAvatar: 'https://picsum.photos/100/100?random=4',
            content: '非常有用的建议，谢谢分享！',
            createdAt: '2023-11-15T11:20:00',
            likeCount: 8,
            isLiked: false
          },
          {
            id: 2,
            userName: '小王',
            userAvatar: 'https://picsum.photos/100/100?random=5',
            content: '我觉得项目实践真的很重要，通过做项目可以学到很多书本上学不到的东西。',
            createdAt: '2023-11-15T13:45:00',
            likeCount: 12,
            isLiked: false
          }
        ];
      }, 500);
    },
    formatDate(dateString) {
      const date = new Date(dateString)
      return date.toLocaleDateString()
    },
    likePost() {
      if (this.isLiked) {
        this.likeCount--
      } else {
        this.likeCount++
      }
      this.isLiked = !this.isLiked
    },
    likeComment(commentId) {
      const comment = this.comments.find(c => c.id === commentId)
      if (comment) {
        if (comment.isLiked) {
          comment.likeCount--
        } else {
          comment.likeCount++
        }
        comment.isLiked = !comment.isLiked
      }
    },
    submitComment() {
      if (!this.commentForm.content.trim()) {
        this.$message.warning('评论内容不能为空')
        return
      }

      // 模拟提交评论
      const newComment = {
        id: this.comments.length + 1,
        userName: '当前用户',
        userAvatar: 'https://picsum.photos/100/100?random=6',
        content: this.commentForm.content,
        createdAt: new Date().toISOString(),
        likeCount: 0,
        isLiked: false
      }

      this.comments.unshift(newComment)
      this.commentForm.content = ''
      this.$message.success('评论发表成功')
    }
  }
}
</script>

<style scoped>
.post-detail-container {
  padding: 20px;
}

.loading-container {
  padding: 20px;
}

.post-detail-card {
  margin-bottom: 20px;
}

.post-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.author-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  margin-right: 15px;
}

.author-info {
  flex: 1;
}

.author-name {
  font-size: 18px;
  font-weight: bold;
  margin: 0;
}

.post-time {
  font-size: 14px;
  color: #999;
  margin: 0;
}

.post-title {
  font-size: 24px;
  margin: 20px 0;
  color: #333;
}

.post-content {
  color: #333;
  line-height: 1.6;
  margin-bottom: 20px;
}

.post-content p {
  margin-bottom: 15px;
}

.post-content ol {
  margin-left: 20px;
  margin-bottom: 15px;
}

.post-footer {
  margin-top: 20px;
}

.post-actions {
  display: flex;
  gap: 10px;
}

.comments-section {
  margin-top: 30px;
}

.comment-form {
  margin-bottom: 30px;
}

.comment-list {
  display: grid;
  gap: 15px;
}

.comment-card {
  padding: 15px;
}

.comment-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.user-info {
  flex: 1;
}

.user-name {
  font-size: 16px;
  font-weight: bold;
  margin: 0;
}

.comment-time {
  font-size: 12px;
  color: #999;
  margin: 0;
}

.comment-content {
  color: #333;
  margin-bottom: 10px;
}

.comment-actions {
  font-size: 14px;
  color: #666;
}

.like-action {
  cursor: pointer;
}

.like-action:hover {
  color: #409eff;
}
</style>
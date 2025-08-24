// src/api/index.js
import axios from 'axios';

// 创建axios实例
const api = axios.create({
  baseURL: '/api', // 使用相对路径，通过vue.config.js中的代理转发
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  }
});

// 请求拦截器 - 添加认证token
api.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token');
    console.log('Request URL:', config.url);
    console.log('Token from localStorage:', token);
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
      console.log('Authorization header added:', config.headers.Authorization);
    } else {
      console.log('No token found in localStorage');
    }
    return config;
  },
  error => {
    console.error('Request interceptor error:', error);
    return Promise.reject(error);
  }
);

// 响应拦截器 - 处理错误
api.interceptors.response.use(
  response => {
    return response.data;
  },
  error => {
    // 处理401未授权错误
    if (error.response && error.response.status === 401) {
      localStorage.removeItem('token');
      localStorage.removeItem('isAdmin');
      // 根据当前路由判断跳转到哪个登录页面
      if (window.location.pathname.startsWith('/admin')) {
        window.location.href = '/admin/login';
      } else {
        window.location.href = '/login';
      }
    }
    return Promise.reject(error);
  }
);

// 用户相关API
export const userAPI = {
  login: (data) => api.post('/users/login', data),
  register: (data) => api.post('/users/register', data),
  getProfile: () => api.get('/users/profile'),
  updateProfile: (data) => api.put('/users/profile', data),
  changePassword: (data) => api.put('/users/password', data),
  getUserList: (params) => api.get('/users', { params }),
  deleteUser: (id) => api.delete(`/users/${id}`),
  createUser: (data) => api.post('/users', data),
  getUserDetail: (id) => api.get(`/users/${id}`),
  updateUser: (id, data) => api.put(`/users/${id}`, data)
};

// 课程分类相关API
export const courseCategoryAPI = {
  getAllCategories: () => api.get('/course-categories'),
  getCategoryById: (id) => api.get(`/course-categories/${id}`)
};

// 课程相关API
export const courseAPI = {
  getCourseList: (params) => api.get('/courses', { params }),
  getCourseDetail: (id) => api.get(`/courses/${id}`),
  createCourse: (data) => api.post('/courses', data),
  updateCourse: (id, data) => api.put(`/courses/${id}`, data),
  deleteCourse: (id) => api.delete(`/courses/${id}`),
  enrollCourse: (courseId) => api.post(`/user-courses/${courseId}`),
  getUserCourses: () => api.get('/user-courses'),
  getUserCourseById: (courseId) => api.get(`/user-courses/${courseId}`),
  unenrollCourse: (courseId) => api.delete(`/user-courses/${courseId}`)
};

// 社区帖子相关API
export const postAPI = {
  getPostList: (params) => api.get('/posts', { params }),
  getPostDetail: (id) => api.get(`/posts/${id}`),
  createPost: (data) => api.post('/posts', data),
  updatePost: (id, data) => api.put(`/posts/${id}`, data),
  deletePost: (id) => api.delete(`/posts/${id}`),
  getUserPosts: () => api.get('/posts/user/posts')
};

// 视频相关API
export const videoAPI = {
  getVideoList: (params) => api.get('/videos', { params }),
  getVideoDetail: (id) => api.get(`/videos/${id}`),
  getVideoCategories: () => api.get('/videos/categories')
};

// 评论相关API
export const commentAPI = {
  getCommentList: (params) => api.get('/comments', { params }),
  deleteComment: (id) => api.delete(`/comments/${id}`)
};

// 将API方法绑定到Vue 3全局属性
const install = (app) => {
  app.config.globalProperties.$api = {
    userAPI,
    courseCategoryAPI,
    courseAPI,
    postAPI,
    videoAPI
  }
}

export default install;
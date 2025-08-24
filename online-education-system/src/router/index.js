import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CourseListView from '../views/CourseListView.vue'
import CategoryTestView from '../views/CategoryTestView.vue'
import CourseDetailView from '../views/CourseDetailView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import ProfileView from '../views/ProfileView.vue'
import CommunityView from '../views/CommunityView.vue'
import PostDetailView from '../views/PostDetailView.vue'
import CreatePostView from '../views/CreatePostView.vue'
import VideoUploadView from '../views/VideoUploadView.vue'
import VideoDetailView from '../views/VideoDetailView.vue'
import VideoListView from '../views/VideoListView.vue'
import AdminDashboardView from '../views/AdminDashboardView.vue'
import AdminHomeView from '../views/AdminHomeView.vue'
import AdminCourseListView from '../views/AdminCourseListView.vue'
import AdminCourseCreateView from '../views/AdminCourseCreateView.vue'
import AdminLoginView from '../views/AdminLoginView.vue'
import AdminUserListView from '../views/AdminUserListView.vue'
import AdminRoleView from '../views/AdminRoleView.vue'
import AdminVideoListView from '../views/AdminVideoListView.vue'
import AdminVideoUploadView from '../views/AdminVideoUploadView.vue'
import AdminPostListView from '../views/AdminPostListView.vue'
import AdminCommentListView from '../views/AdminCommentListView.vue'
import AdminSettingsView from '../views/AdminSettingsView.vue'

const routes = [
    // 分类测试路由
    {
      path: '/category-test',
      name: 'categoryTest',
      component: CategoryTestView
    },
    // 后台管理路由
    {
      path: '/admin/login',
      name: 'adminLogin',
      component: AdminLoginView
    },
    {
      path: '/admin/dashboard',
      name: 'adminDashboard',
      component: AdminDashboardView,
      meta: { requiresAuth: true, requiresAdmin: true },
      children: [
        {
          path: '',
          name: 'adminHome',
          component: AdminHomeView
        },
        {
          path: 'courses',
          name: 'adminCourses',
          component: AdminCourseListView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'courses/create',
          name: 'adminCreateCourse',
          component: AdminCourseCreateView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'users',
          name: 'adminUsers',
          component: AdminUserListView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'roles',
          name: 'adminRoles',
          component: AdminRoleView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'videos',
          name: 'adminVideos',
          component: AdminVideoListView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'videos/create',
          name: 'adminCreateVideo',
          component: AdminVideoUploadView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'posts',
          name: 'adminPosts',
          component: AdminPostListView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'comments',
          name: 'adminComments',
          component: AdminCommentListView,
          meta: { requiresAuth: true, requiresAdmin: true }
        },
        {
          path: 'settings',
          name: 'adminSettings',
          component: AdminSettingsView,
          meta: { requiresAuth: true, requiresAdmin: true }
        }
      ]
    },
    // 重定向
    {
      path: '/admin',
      redirect: '/admin/login'
    },
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: { loading: true }
    },
    {
      path: '/videos',
      name: 'videos',
      component: VideoListView,
      meta: { loading: true }
    },
  {    path: '/courses',    name: 'courses',    component: CourseListView,    meta: { loading: true }  },
  {
    path: '/courses/:id',
    name: 'courseDetail',
    component: CourseDetailView,
    props: true
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
      path: '/profile',
      name: 'profile',
      component: ProfileView,
      meta: { requiresAuth: true }
    },
    {
      path: '/community',
      name: 'community',
      component: CommunityView,
      meta: { loading: true }
    },
    {
      path: '/post/:id',
      name: 'postDetail',
      component: PostDetailView,
      props: true
    },
    {
      path: '/post/create',
      name: 'createPost',
      component: CreatePostView,
      meta: { requiresAuth: true }
    },
    {
      path: '/video/upload',
      name: 'videoUpload',
      component: VideoUploadView,
      meta: { requiresAuth: true }
    },
    {
      path: '/video/:id',
      name: 'videoDetail',
      component: VideoDetailView,
      props: true
    }
  ]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  // 检查是否需要登录
  const requiresAuth = to.meta.requiresAuth || false
  const requiresAdmin = to.meta.requiresAdmin || false
  const isLoggedIn = localStorage.getItem('token')
  const isAdmin = localStorage.getItem('isAdmin') === 'true'

  // 如果是管理员路由且未登录，重定向到管理员登录页面
  if (requiresAdmin && !isLoggedIn) {
    next('/admin/login')
  } else if (requiresAuth && !isLoggedIn) {
    next('/login')
  } else if (requiresAdmin && !isAdmin) {
    // 如果需要管理员权限但不是管理员，则重定向到首页
    next('/')
  } else {
    next()
  }
})

export default router
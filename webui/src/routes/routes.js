import DashboardLayout from '../components/Dashboard/Layout/DashboardLayout.vue'
import LoginLayout from '../components/Login/Layout/LoginLayout.vue'
// GeneralViews
import NotFound from '../components/GeneralViews/NotFoundPage.vue'

// Admin pages
import Overview from 'src/components/Dashboard/Views/Overview.vue'

const routes = [
  {
    path: '/',
    component: DashboardLayout,
    redirect: '/login'
  },
  {
    path: '/login',
    component: LoginLayout
  },
  {
    path: '/admin',
    component: DashboardLayout,
    redirect: '/admin/stats',
    children: [
      {
        path: 'overview',
        name: 'overview',
        component: Overview
      }
    ]
  },
  { path: '*', component: NotFound }
]

export default routes

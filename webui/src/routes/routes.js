import DashboardLayout from '../components/Dashboard/Layout/DashboardLayout.vue'
import LoginLayout from '../components/Login/Layout/LoginLayout.vue'
// GeneralViews
import NotFound from '../components/GeneralViews/NotFoundPage.vue'

// Admin pages
import Overview from 'src/components/Dashboard/Views/Overview.vue'
import Assets from 'src/components/Dashboard/Views/Assets.vue'

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
    path: '/book',
    component: DashboardLayout,
    redirect: '/book/summary',
    children: [
      {
        path: 'summary',
        name: 'Summary',
        component: Overview
      },
      {
        path: 'Accounts',
        name: 'Accounts',
        component: Overview
      },
      {
        path: 'assets',
        name: 'Currencies & Assets',
        component: Assets
      },
      {
        path: 'transactions',
        name: 'Transactions',
        component: Overview
      }
    ]
  },
  { path: '*', component: NotFound }
]

export default routes

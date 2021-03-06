import DashboardLayout from '../components/Dashboard/Layout/DashboardLayout.vue'
import LoginLayout from '../components/Login/Layout/LoginLayout.vue'
// GeneralViews
import NotFound from '../components/GeneralViews/NotFoundPage.vue'

// Admin pages
import Overview from 'src/components/Dashboard/Views/Overview.vue'
import Assets from 'src/components/Dashboard/Views/Assets.vue'
import Accounts from 'src/components/Dashboard/Views/Accounts.vue'
import AccountBalances from 'src/components/Dashboard/Views/AccountBalances.vue'
import AccountMovements from 'src/components/Dashboard/Views/AccountMovements.vue'
import Transactions from 'src/components/Dashboard/Views/Transactions.vue'
import EditTransaction from 'src/components/Dashboard/Views/EditTransaction.vue'

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
        path: 'accounts',
        name: 'Accounts',
        component: Accounts
      },
      {
        path: 'accounts/:acc_id/movements',
        name: 'Account Movements',
        component: AccountMovements
      },
      {
        path: 'accounts/:acc_id/balances',
        name: 'Account Balances',
        component: AccountBalances
      },
      {
        path: 'assets',
        name: 'Currencies & Assets',
        component: Assets
      },
      {
        path: 'transactions',
        name: 'Transactions',
        component: Transactions
      },
      {
        path: 'transactions/:tr_id',
        name: 'Add Transaction',
        component: EditTransaction
      }
    ]
  },
  { path: '*', component: NotFound }
]

export default routes

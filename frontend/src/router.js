import Layout from './components/Layout.vue'
import Hotels from './components/Hotels.vue'
import HotelView from './components/HotelView.vue'
import SignUpForm from './components/SignUpForm.vue'
import SignInForm from './components/SignInForm.vue'
import UserView from './components/UserView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: '/',
                component: Hotels,
            },
            {
                path: 'hotels',
                alias: '/',
                component: Hotels,
            },
            {
                path: 'hotels/:id',
                name: 'hotel',
                component: HotelView
            },
            {
                path: 'signup',
                component: SignUpForm,
            },
            {
                path: 'signin',
                component: SignInForm,
            },
            {
                path: 'profile',
                component: UserView,
            }
        ],
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
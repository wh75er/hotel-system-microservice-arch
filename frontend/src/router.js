import Layout from './components/Layout.vue'
import Hotels from './components/Hotels.vue'
import HotelView from './components/HotelView.vue'
import SignUpForm from './components/SignUpForm.vue'
import SignInForm from './components/SignInForm.vue'
import UserView from './components/UserView.vue'
import HotelForm from './components/HotelForm.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: '/',
                name: 'home',
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
                props: true,
                component: HotelView
            },
            {
                path: 'signup',
                component: SignUpForm,
            },
            {
                path: 'signin',
                name: 'signIn',
                component: SignInForm,
            },
            {
                path: 'profile',
                component: UserView,
            },
            {
                path: 'hotels/create',
                component: HotelForm,
            }
        ],
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
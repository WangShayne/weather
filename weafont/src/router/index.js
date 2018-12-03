import Vue from 'vue'
import Router from 'vue-router'
import cityList from '@/components/cityList'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'cityList',
      component: cityList
    }
  ]
})

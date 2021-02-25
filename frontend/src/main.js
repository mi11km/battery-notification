import 'core-js/stable';
import 'regenerator-runtime/runtime';
import '@mdi/font/css/materialdesignicons.css';
import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import App from './App.vue';
import Wails from '@wailsapp/runtime';

Vue.use(Vuetify);

Vue.config.productionTip = false;
Vue.config.devtools = true;

Wails.Init(() => {
    new Vue({
        vuetify: new Vuetify({
            icons: {
                iconfont: 'mdi'
            },
        }),
        render: h => h(App)
    }).$mount('#app');
});

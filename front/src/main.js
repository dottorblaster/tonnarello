var Vue = require('vue');
var VueRouter = require('vue-router');
var VueResource = require('vue-resource');

var Sidebar = require('./sidebar.vue');
var App = require('./app.vue');

Vue.use(VueRouter);
Vue.use(VueResource);

var application = new Vue({
	el: 'body',
	components: {
		sidebar: Sidebar,
		app: App
	}
});

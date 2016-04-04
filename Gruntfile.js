module.exports = function(grunt) {
	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),
		watch: {
			sass: {
				files: ['public/sass/*.scss'],
				tasks: ['sass']
			},
			browserify: {
				files: ['front/src/*.js', 'front/src/*.vue'],
				tasks: ['browserify']
			}
		},
		sass: {
			dist: {
				options: {
					loadPath: ['node_modules/materialize-css/sass']
				},
				files: {
					'front/public/css/style.css': 'front/src/style.scss'
				}
			}
		},
		browserify: {
			build: {
				src: ['front/src/*.js', 'front/src/*.vue'],
				dest: 'front/public/bundle.js'
			},
			options: {
				transform: ['vueify']
			}
		}
	});

	grunt.loadNpmTasks('grunt-browserify');
	grunt.loadNpmTasks('grunt-contrib-sass');
	grunt.loadNpmTasks('grunt-contrib-watch');

	grunt.registerTask('default', ['sass', 'browserify']);
};

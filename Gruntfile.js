module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    watch: {
      sass: {
        files: ['public/sass/*.scss'],
        tasks: ['sass']
      },
      browserify: {
        files: ['public/js/*.js', 'public/js/components/*.js', '!public/js/bundle.js'],
        tasks: ['browserify']
      }
    },
    sass: {
      dist: {
        options: {
          loadPath: ['node_modules/milligram-scss/dist/']
        },
        files: {
          'public/css/style.css': 'public/sass/style.scss'
        }
      }
    },
    browserify: {
      build: {
        src: ['public/js/*.js', 'public/js/components/*.js', '!public/js/bundle.js'],
        dest: 'public/js/bundle.js'
      }
    }
  });

  grunt.loadNpmTasks('grunt-browserify');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', ['sass', 'browserify']);
};

const gulp = require('gulp');
const sass = require('gulp-sass')

gulp.task('sass', function(
    gulp.src('./src/scss/main.scss')
        .pipe(sass())
        .pipe(gulp.dest('./dist/css'))
))
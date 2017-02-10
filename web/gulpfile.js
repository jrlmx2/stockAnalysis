var gulp = require('gulp');
var watch = require('gulp-watch');
var coffee = require('gulp-coffee');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var imagemin = require('gulp-imagemin');
var sourcemaps = require('gulp-sourcemaps');
var changed = require("gulp-changed");
var cached = require("gulp-cached");
var remember = require("gulp-remember");
var newer = require("gulp-newer");
var miniHTML = require("gulp-minify-html");
var cleanCSS = require("gulp-clean-css");
var less = require("gulp-less");
var del = require('del');
var config = require("./config.json");
var pump = require("pump");

gulp.task('cleanBuild', function(){
  return del(['./build/css/**/*','./build/js/**/*']);
});

//define task scripts. Before we run scripts, we run clean build
gulp.task('scripts',function(cb){
  watchJavascript(cb);
});

gulp.task('style',function(cb){
  watchStyle(cb);
});

var calls=0;
function debug(fn,args){
  console.log("Running function number "+calls);
  calls++;
  console.log("Args "+args);
  return fn(args);
}

function watchStyle(cb){
  if(config.usingLess){
    return watch(config.paths.less, {ignoreInitial:false,verbose:true,name:"Gulp Less Watching"}, function(){
      console.log("Building to" +config.paths.buildDir + config.paths.cssDest);
      pump([
        gulp.src(config.paths.less),
        sourcemaps.init(),
        less(config.less||{}),
        cleanCSS(),
        sourcemaps.write(),
        gulp.dest(config.paths.buildDir + config.paths.cssDest)
      ],cb);
      console.log("finished");
    });
  } else {
        console.log(config.paths.less);
    return watch(config.paths.css, {ignoreInitial:false,verbose:true,name:"Gulp CSS Watching"}, function(){
            console.log("Building to" +config.paths.buildDir + config.paths.cssDest);
      pump([
        gulp.src(config.paths.css),
        sourcemaps.init(),
        cleanCSS(),
        sourcemaps.write(),
        gulp.dest(config.paths.buildDir + config.paths.cssDest)
      ],cb);
    });
          console.log("finished");
  }
}

function watchJavascript(cb){
  if(config.usingCoffee){
    return watch(config.paths.coffee, {ignoreInitial:false,verbose:true,name:"Gulp CoffeeScript Watching"},function(){
      console.log("Making coffee");
      pump([
        gulp.src(config.paths.coffee),
        coffee({bare: true}),
        gulp.dest(config.paths.coffeeDest)
      ]);
      if(config.prod){
        console.log("combining javascript");
        pump([
          gulp.src(config.paths.js),
          sourcemaps.init(),
          uglyify(config.uglify||{}),
          concat(config.paths.minifiedScript),
          sourcemaps.write(),
          dest(config.paths.buildDir + config.paths.jsDest)
        ],cb);
      }
      console.log("finished");
    });
  } else {
    return watch(config.paths.js, {ignoreInitial:false,verbose:true,name:"Gulp Javascript Watching"},function(){
      if(config.prod){
        console.log("combining javascript");
        pump([
          gulp.src(config.paths.js),
          sourcemaps.init(),
          uglyify(config.uglify||{}),
          concat(config.paths.minifiedScript),
          sourcemaps.write(),
          dest(config.paths.buildDir + config.paths.jsDest)
        ],cb);
      }
      console.log("finished");
    });
  }
}

gulp.task('updateView',function(cb){
  watch(config.paths.inject,  {ignoreInitial:false,verbose:true,name:"Gulp Include Watching",debounceDelay:2000},function(){
    pump([
      gulp.src(config.paths.page),
      gulp.src(inject(config.paths.inject)),
      gulp.dest(config.paths.buildDir)
    ],cb);
  })
});

// The default task (called when you run `gulp` from cli)
gulp.task('default', ['scripts','style','updateView']);

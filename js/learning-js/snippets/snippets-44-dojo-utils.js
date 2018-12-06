var	noop = function(){ },

    isEmpty = function(it){
        for(var p in it){
            return 0;
        }
        return 1;
    },

    toString = {}.toString,

    isFunction = function(it){
        return toString.call(it) == "[object Function]";
    },

    isString = function(it){
        return toString.call(it) == "[object String]";
    },

    isArray = function(it){
        return toString.call(it) == "[object Array]";
    },

    forEach = function(vector, callback){
        if (vector){
            for (var i = 0; i < vector.length;){
                callback(vector[i++]);
            }
        }
    },

    mix = function(dest, src){
        for(var p in src){
            dest[p] = src[p];
        }
        return dest;
    },



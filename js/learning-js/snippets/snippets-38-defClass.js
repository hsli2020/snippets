var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { 
        for (var key in parent) {
            if (__hasProp.call(parent, key))
                child[key] = parent[key];
        }
        function ctor() {
            this.constructor = child; 
        } 
        ctor.prototype = parent.prototype; 
        child.prototype = new ctor(); 
        child.__super__ = parent.prototype; 
        return child; 
    };

var Users = (function(_super) {
    __extends(Users, _super);

    function Users() {
        return Users.__super__.constructor.apply(this, arguments);
    }

    Users.prototype.model = User;
    Users.prototype.url = "/api/users";
    Users.prototype.initialize = function() {
        return console.log('Users Collection initialize');
    };
    Users.prototype.parse = function(data) {
        return data.result.users;
    };

    return Users;

})(Backbone.Collection);


function newClass(classObj, superClass) {
	if (!classObj) classObj = {};

	if (typeof classObj.__construct !== "function") 
        classObj.__construct = function() {};

	var f = classObj.__construct;

	f.extend = function(classObj) { return newClass(classObj, this); }
	
	if(superClass) {
		for(var i in superClass.prototype)
            f.prototype[i] = superClass.prototype[i];
		classObj.__super = superClass.prototype;
	}
	
	for(var j in classObj) {
		if(superClass && typeof classObj[j] == "function") {
			f.prototype[j] = (function(func, superClass) {
				return function() {
					var tmpSuper = this.__super;
					this.__super = superClass.prototype;
					var result = func.apply(this, arguments);
					this.__super = tmpSuper;
					return result;
				};
			})(classObj[j], superClass);
		} else {
			f.prototype[j] = classObj[j];
		}
	}
	return f;
}

var Class = newClass({});

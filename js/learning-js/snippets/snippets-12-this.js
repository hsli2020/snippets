
var obj = {

    specialFunc:  function ()   { console.log('specialFunc'); },
    anotherFunc:  function ()   { console.log('anotherFunc'); },
    getAsyncData: function (cb) { cb(); },

    render: function () {
        var that = this;  // !!!
        this.getAsyncData(function () {
            that.specialFunc();
            that.anotherFunc();
        });
    }
};

//obj.render();

var bad = {

    specialFunc:  function ()   { console.log('specialFunc'); },
    anotherFunc:  function ()   { console.log('anotherFunc'); },
    getAsyncData: function (cb) { cb(); },

    render: function () {
        this.getAsyncData(function () {
            this.specialFunc();
            this.anotherFunc();
        });
    }
};

//bad.render();


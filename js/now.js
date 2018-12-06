function now() {
    var d = new Date();
    var offset = d.getTimezoneOffset()*60*1000;
    d.setTime(d.getTime() - offset);

    var s = d.toISOString();
    return s.substr(0, 10) + ' ' + s.substr(11, 8);
}
//console.log(now());

function startOfWeek(date) {
    date.setDate(date.getDate()-date.getDay());
    return date;
}

//var dt = new Date(); 
//console.log(startOfWeek(dt).toISOString());

Date.prototype.getFirstDayOfWeek = function() {
    return (new Date(this.setDate(this.getDate() - this.getDay())));
}
Date.prototype.getLastDayOfWeek = function() {
    return (new Date(this.setDate(this.getDate() - this.getDay() + 6)));
}
Date.prototype.toLocal = function() {
    return (new Date(this.getTime() - this.getTimezoneOffset()*60*1000));
}
Date.prototype.print = function() {
    console.log(this.format());
}
Date.prototype.format = function() {
    var s = this.toISOString();
    return (s.substr(0, 10) + ' ' + s.substr(11, 8));
}
Date.prototype.dateString = function() {
    return this.toISOString().substr(0, 10);
}
Date.prototype.timeString = function() {
    return this.toISOString().substr(11, 8);
}

var today = new Date();
//console.log(today.getFirstDayOfWeek());
//console.log(today.getLastDayOfWeek());
//today.getFirstDayOfWeek().print();
//today.getLastDayOfWeek().print();
today.toLocal().print();
//console.log(today.dateString());
//console.log(today.timeString());

var Calendar = function() {
    this.namesOfWeek  = [ 'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat' ];
    this.namesOfMonth = [ 'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec' ];
    this.daysOfMonths = [ 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 ];
};

Calendar.prototype.isLeap = function(y) {
    return (y % 100 == 0) && (y % 400 == 0) ? true : (y % 4 == 0 ? true : false);
};

Calendar.prototype.getDates = function(date) {
    var Y = date.getFullYear();
    var M = date.getMonth();

    if (this.isLeap(Y)) {
        this.daysOfMonths[1] = 29;
    }

    var totalDays = 7*6;

    var firstDay = new Date(Y, M, 1);
    var daysPrevMonth = firstDay.getDay() || 7;
    var daysThisMonth = this.daysOfMonths[M];
    var daysNextMonth = totalDays - daysThisMonth - daysPrevMonth;

    // Prev month
    var prevY, prevM;

    if (M == 0) {
        prevY = Y - 1;
        prevM = 11;
    } else {
        prevY = Y;
        prevM = M - 1;
    }
    var prevMax = this.daysOfMonths[prevM];

    // Next month
    var nextY, nextM;

    if (M == 11) {
        nextY = Y + 1;
        nextM = 0;
    } else {
        nextY = Y;
        prevM = M + 1;
    }

    var dates = [];

    // Add dates in prev month
    for (var i=daysPrevMonth; i>0; i--){
        dates.push({
            year:  prevY,
            month: prevM,
            day:   (prevMax - i + 1),
            id:    'prev' + (prevMax - i + 1),
            attr:  'prevmonth'
        });
    }

    // Add dates in this month
    for(var i=1; i<=daysThisMonth; i++){
        dates.push({
            year:  Y,
            month: M,
            day:   i,
            id:    'day' + i,
            attr:  'thismonth'
        });
    }

    // Add dates in next month
    for (var i=1; i<=daysNextMonth; i++){
        dates.push({
            year:  nextY,
            month: nextM,
            day:   i,
            id:    'next' + i,
            attr:  'nextmonth'
        });
    }

    return dates;
};

Calendar.prototype.getHtml = function(dateStr) {
    var y, m, d;
    [ y, m, d ] = dateStr.split('-');
    var date = new Date(y, m-1, 1);
    var days = this.getDates(date);

    var str = `<table class="w3-table w3-bordered" id="calendar">`;

    var monthYear = this.namesOfMonth[date.getMonth()] + ", " + date.getFullYear();
    str += `<caption><h3 style="margin-top:0;">${monthYear}</h3></caption>`;

    str += "<tr>";
    for(var i=0; i<7; i++) {
        str += `<th class="daysheader">${this.namesOfWeek[i]}</th>`;
    }
    str += "</tr>";

    for (var i=0; i<days.length; i++) {
        if ((i+1)%7 == 1) {
            str += "<tr>";
        }

        var d = days[i];

        str += `<td class="${d.attr}" id="${d.id}">${d.day}</td>`;
        if ((i+1)%7 == 0) {
            str += "</tr>";
        }
    }
    str += "</table>";

    return str;
};

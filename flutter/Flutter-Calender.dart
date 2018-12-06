#/example/lib/main.dart

import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:http/http.dart';
import 'dart:convert';
import 'package:calendar/calendar.dart';
import 'package:calendar/controllers.dart';
import 'package:calendar/utils.dart';

Widget appBar = new AppBar(title: new Text('Calendar Example'), backgroundColor: Colors.blue[500]);
ThemeData theme = new ThemeData(primarySwatch: Colors.blue);

/// INFO(mperrotte):
/// - Define the model class that will represent your data.
/// - Extend the DataModel for functionality and type checking
class Event extends DataModel {
  Event(
    Map<String, dynamic> data,
  )
      : id = data['id'],
        title = data['title'],
        url = data['url'],
        start = DateTime.parse(data['date_start']),
        end = DateTime.parse(data['date_end']),
        details = data['details'],
        super(
          year: DateTime.parse(data['date_start']).year,
          month: DateTime.parse(data['date_start']).month,
          date: DateTime.parse(data['date_start']).day,
        );

  final int id;
  final String title;
  final String url;
  final DateTime start;
  final DateTime end;
  final dynamic details;
}

EventsView renderEventsHandler(List<DataModel> events) {
  return new EventsView(
    itemBuilder: (BuildContext context, int index) {
      if (index >= events.length) {
        return null;
      }
      return new Dismissible(
        // key needs to be more unique
        key: new ValueKey<int>((events[index] as Event).id),
        direction: DismissDirection.horizontal,
        onDismissed: (DismissDirection direction) {
          print(direction);
          removeEventAction(index);
          print('done removing');
        },
        onResize: () {
          print('resize');
        },
        background: new Container(
          decoration: new BoxDecoration(
            color: theme.primaryColor,
          ),
          child: new ListTile(
            leading: new Icon(
              Icons.input,
              color: Colors.white,
              size: 36.0,
            ),
          ),
        ),
        child: new Container(
          height: 75.0,
          decoration: new BoxDecoration(
            color: theme.canvasColor,
            border: new Border(
              bottom: new BorderSide(color: theme.dividerColor),
            ),
          ),
          child: new ListTile(
            title: new Text((events[index] as Event).title),
            subtitle: new Text((events[index] as Event).url),
            isThreeLine: true,
          ),
        ),
      );
    },
  );
}

dynamic fetchDataHandler(String uri) async {
  Response response = await get(uri);
  String json = response.body;
  if (json == null) {
    // NOTE(mperrotte): no body value from response; bail out;
    return null;
  }
  JsonDecoder decoder = new JsonDecoder();
  try {
    dynamic payload = decoder.convert(json);
    return payload;
  } on ArgumentError {
    // NOTE(mperrotte): unable to decode json payload
    return null;
  }
}

List<Event> parseDataHandler(List<Map<String, dynamic>> payload) {
  List<Event> events = new List<Event>();
  for (Map<String, dynamic> item in payload) {
    events.add(new Event(item));
  }
  return events;
}

String getUriHandler() {
  return 'https://gist.githubusercontent.com/mikemimik/2da9df7ca94f0bdfd493bf70e53ec3b6/raw/f4e0fe43219b02091f4eec79ebd10e37185e213d/data-events.json';
}

CalendarController calendarController = new CalendarController(
  renderEvents: renderEventsHandler,
);
DataController dataController = new DataController(
  fetchData: fetchDataHandler,
  parseData: parseDataHandler,
  getUri: getUriHandler,
);

void main() {
  debugPaintSizeEnabled = false;

  runApp(
    new MaterialApp(
      title: 'Calendar Example',
      theme: theme,
      home: new Scaffold(
        appBar: appBar,
        body: new Calendar(
          calendarController: calendarController,
          dataController: dataController,
        ),
      ),
    ),
  );
}

#/lib/calendar.dart

/// Calendar Exports
/// The calendar library implements a [Calendar] component with a [CalendarController].

export 'src/calendar/calendar.dart';
export 'src/calendar/calendar_controller.dart';

export 'src/events/events_view.dart';
export 'src/data/data_model.dart';

#/lib/controllers.dart

/// Controller Exports
export 'src/calendar/calendar_controller.dart';
export 'src/data/data_controller.dart';

#/lib/src/actions.dart

import 'package:flutter_flux/flutter_flux.dart';
import 'package:calendar/controllers.dart';
import 'package:calendar/src/types.dart';

final Action<RenderableView> switchViewAction = new Action<RenderableView>();
final Action<Map<String, int>> selectDateAction = new Action<Map<String, int>>();
final Action<Map<String, dynamic>> fetchDataAction = new Action<Map<String, dynamic>>();
final Action<int> removeEventAction = new Action<int>();

final Action<CalendarController> setCalendarControllerAction = new Action<CalendarController>();
final Action<DataController> setDataControllerAction = new Action<DataController>();

#/lib/src/calendar/calendar.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_flux/flutter_flux.dart';

import 'package:calendar/src/calendar/calendar_view.dart';
import 'package:calendar/src/calendar/calendar_events_view.dart';
import 'package:calendar/utils.dart';
import 'package:calendar/controllers.dart';

// TODO(mperrotte): fetch data
class Calendar extends StoreWatcher {
  Calendar({
    Key key,
    DateTime initializeDate,
    @required this.calendarController,
    @required this.dataController,
  })
      : year = (initializeDate == null) ? new DateTime.now().year : initializeDate.year,
        month = (initializeDate == null) ? new DateTime.now().month : initializeDate.month,
        date = (initializeDate == null) ? new DateTime.now().day : initializeDate.day,
        super(key: key);

  final int year;
  final int month;
  final int date;
  final CalendarController calendarController;
  final DataController dataController;

  @override
  void initStores(ListenToStore listenToStore) {
    listenToStore(calendarStoreToken);
    listenToStore(controllerStoreToken);

    // INFO(mperrotte): Fetch data after we've initialized the listeners
    String uri = dataController.getUri();
    dynamic payload = dataController.fetchData(uri);
    fetchDataAction({
      'payload': payload,
      'parseData': dataController.parseData,
    });
  }

  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    final CalendarStore calendarStore = stores[calendarStoreToken];
    Widget component;

    // INFO(mperrotte): put controllers into store
    setCalendarControllerAction(calendarController);
    setDataControllerAction(dataController);

    switch (calendarStore.currentView) {
      case RenderableView.calendar:
        component = new CalendarView(
          year: year,
          month: month,
          date: date,
        );
        break;
      case RenderableView.events:
        component = new CalendarEventsView(
          year: calendarStore.selectedYear,
          month: calendarStore.selectedMonth,
          date: calendarStore.selectedDate,
          child: calendarController.renderEvents(
            calendarStore.eventsByDate[calendarStore.selectedDate],
          ),
        );
        break;
      case RenderableView.event:
        break;
    }
    return new Material(
      child: component,
    );
  }
}

#/lib/src/calendar/calendar_controller.dart

import 'package:flutter/foundation.dart';
import 'package:calendar/src/types.dart';

abstract class BaseCalendarController {
  BaseCalendarController({
    @required this.renderEvents,
  });

  final RenderEventsCallback renderEvents;
}

class CalendarController extends BaseCalendarController {
  CalendarController({
    @required renderEvents,
  })
      : super(
          renderEvents: renderEvents,
        );
}

#/lib/src/calendar/calendar_events_view.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:calendar/src/events/events_header_view.dart';
import 'package:calendar/src/events/events_footer_view.dart';

class CalendarEventsView extends StatelessWidget {
  CalendarEventsView({
    @required this.year,
    @required this.month,
    @required this.date,
    @required this.child,
  });

  final int year;
  final int month;
  final int date;
  final Widget child;

  @override
  Widget build(BuildContext context) {
    ThemeData theme = Theme.of(context);

    return new Container(
      constraints: new BoxConstraints(),
      child: new Column(
        // crossAxisAlignment: CrossAxisAlignment.stretch,
        // mainAxisAlignment: MainAxisAlignment.spaceBetween,

        children: <Widget>[
          new EventsHeader(year, month, date),
          child,
          new EventsFooter(theme),
        ],
      ),
    );
  }
}

#/lib/src/calendar/calendar_event_icon.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

class CalendarViewEventIcon extends StatelessWidget {
  CalendarViewEventIcon({
    @required this.backgroundColor,
  });

  final Color backgroundColor;

  @override
  Widget build(BuildContext context) {
    return new Container(
      height: 5.0,
      width: 5.0,
      margin: const EdgeInsets.all(0.5),
      decoration: new BoxDecoration(
        color: backgroundColor,
      ),
    );
  }
}

#/lib/src/calendar/calendar_event_icon_row.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'package:calendar/src/calendar/calendar_event_icon.dart';

class CalendarViewEventIconRow extends StatelessWidget {
  CalendarViewEventIconRow({
    @required this.eventIcons,
  });

  final List<CalendarViewEventIcon> eventIcons;

  @override
  Widget build(BuildContext context) {
    return new Row(
      children: eventIcons,
    );
  }
}

#/lib/src/calendar/calendar_view.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'package:calendar/src/month/month.dart';
import 'package:calendar/controllers.dart';
import 'package:calendar/utils.dart';

class CalendarView extends StatelessWidget {
  CalendarView({
    @required this.year,
    @required this.month,
    @required this.date,
  });

  final int year;
  final int month;
  final int date;

  @override
  Widget build(BuildContext context) {
    return new Container(
      child: new Column(
        children: <Widget>[
          new CalendarViewHeader(monthName: MonthNames[month - 1]['long']),
          new Month(
            year: year,
            month: month,
            date: date,
          ),
        ],
      ),
    );
  }
}

class CalendarViewHeader extends StatelessWidget {
  CalendarViewHeader({
    @required this.monthName,
  });

  final String monthName;
  @override
  Widget build(BuildContext context) {
    return new Container(
      height: 40.0,
      margin: const EdgeInsets.only(top: 5.0, bottom: 10.0),
      child: new Align(
        alignment: FractionalOffset.center,
        child: new Text(monthName),
      ),
    );
  }
}

#/lib/src/data/data_controller.dart

import 'package:flutter/foundation.dart';
import 'package:calendar/src/types.dart';

class BaseDataController {
  BaseDataController({
    @required this.fetchData,
    @required this.parseData,
    @required this.getUri,
  });

  final FetchDataCallback fetchData;
  final ParseDataCallback parseData;
  final GetDataUriCallback getUri;
}

class DataController extends BaseDataController {
  DataController({
    @required fetchData,
    @required parseData,
    @required getUri,
  })
      : super(
          fetchData: fetchData,
          parseData: parseData,
          getUri: getUri,
        );
}

#/lib/src/data/data_model.dart

import 'package:flutter/foundation.dart';

abstract class DataModel {
  DataModel({
    @required this.year,
    @required this.month,
    @required this.date,
  });

  final int year;
  final int month;
  final int date;
}

#/lib/src/day/day.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_flux/flutter_flux.dart';

import 'package:calendar/src/calendar/calendar_event_icon_row.dart';
import 'package:calendar/src/calendar/calendar_event_icon.dart';
import 'package:calendar/src/data/data_model.dart';
import 'package:calendar/utils.dart';

class Day extends StoreWatcher {
  Day({
    @required this.year,
    @required this.month,
    @required this.date,
    List<DataModel> events,
    this.today: false,
  })
      : textStyle = today ? new TextStyle(color: Colors.orange[500]) : new TextStyle();

  final int year;
  final int month;
  final int date;
  final bool today;
  final TextStyle textStyle;

  // Functions
  List<CalendarViewEventIcon> _generateEventIcons(int count) {
    List<CalendarViewEventIcon> eventIcons = new List<CalendarViewEventIcon>();
    for (int i = 0; i < count; i++) {
      eventIcons.add(
        new CalendarViewEventIcon(
          backgroundColor: Colors.red[500],
        ),
      );
    }
    return eventIcons;
  }

  List<Widget> _generateEventIconRows(CalendarStore store) {
    List<Widget> children = new List<Widget>();
    if (store.eventsByDate[date] != null && store.eventsByDate[date].length != 0) {
      num rows = store.eventsByDate[date].length / 4;
      num rowCount = rows.round();
      num remainder = store.eventsByDate[date].length.remainder(4);
      for (int i = 0; i < rowCount; i++) {
        if (i == rowCount - 1) {
          int iconCount = (remainder == 0) ? 4 : remainder;
          children.add(
            new CalendarViewEventIconRow(
              eventIcons: _generateEventIcons(iconCount),
            ),
          );
        } else {
          children.add(
            new CalendarViewEventIconRow(
              eventIcons: _generateEventIcons(4),
            ),
          );
        }
      }
    }
    return children;
  }

  @override
  void initStores(ListenToStore listenToStore) {
    listenToStore(calendarStoreToken);
  }

  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    final CalendarStore calendarStore = stores[calendarStoreToken];

    return new Flexible(
      child: new InkWell(
        onTap: () {
          Map<String, int> currentDay = new Map.fromIterables(
            ['year', 'month', 'date'],
            [year, month, date],
          );
          selectDateAction(currentDay);
          switchViewAction(RenderableView.events);
        },
        child: new Container(
          height: 60.0,
          decoration: new BoxDecoration(
              border: new Border.all(
            color: Colors.black,
            width: 0.5,
          )),
          padding: const EdgeInsets.all(4.0),
          child: new Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              new Align(
                alignment: FractionalOffset.centerRight,
                child: new Container(
                  child: new Text(
                    date.toString(),
                    style: textStyle,
                  ),
                ),
              ),
              new Column(
                mainAxisAlignment: MainAxisAlignment.end,
                children: _generateEventIconRows(calendarStore),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

#/lib/src/day/day_header.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_flux/flutter_flux.dart';
import 'day.dart';

class DayHeader extends Day {
  DayHeader({
    @required this.day,
  });

  final String day;

  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    return new Flexible(
      child: new Container(
        decoration: new BoxDecoration(
          border: new Border.all(
            color: Colors.black,
            width: 0.5,
          ),
        ),
        padding: const EdgeInsets.all(4.0),
        child: new Align(
          alignment: FractionalOffset.center,
          child: new Text(day),
        ),
      ),
    );
  }
}

#/lib/src/events/events_footer_view.dart

import 'package:flutter/material.dart';
import 'package:calendar/src/types.dart';
import 'package:calendar/src/actions.dart';

class EventsFooter extends StatelessWidget {
  EventsFooter(
    this.theme,
  );

  final ThemeData theme;

  @override
  Widget build(BuildContext context) {
    return new Container(
      decoration: new BoxDecoration(
        color: theme.accentColor,
      ),
      child: new SizedBox(
        height: 48.0,
        child: new Row(
          children: <Widget>[
            new IconButton(
              iconSize: 36.0,
              icon: new Icon(
                Icons.chevron_left,
                size: 36.0,
              ),
              onPressed: () {
                switchViewAction(RenderableView.calendar);
              },
            )
          ],
        ),
      ),
    );
  }
}

#/lib/src/events/events_header_view.dart

import 'package:flutter/material.dart';
import 'package:calendar/src/types.dart';

class EventsHeader extends StatelessWidget {
  EventsHeader(
    this.year,
    this.month,
    this.date,
  );

  final int year;
  final int month;
  final int date;

  @override
  Widget build(BuildContext context) {
    return new Center(
      child: new Container(
        constraints: new BoxConstraints(),
        margin: const EdgeInsets.only(
          top: 12.0,
          bottom: 12.0,
        ),
        child: new SizedBox(
          height: 15.0,
          child: new Text('Day: ' + MonthNames[month - 1]['long'] + ' $date, $year'),
        ),
      ),
    );
  }
}

#/lib/src/events/events_view.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_flux/flutter_flux.dart';
import 'package:calendar/src/stores.dart';

class EventsView extends StoreWatcher {
  EventsView({
    @required this.itemBuilder,
  });

  final IndexedWidgetBuilder itemBuilder;

  @override
  void initStores(ListenToStore listenToStore) {
    listenToStore(calendarStoreToken);
  }

  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    final CalendarStore calendarStore = stores[calendarStoreToken];

    return new Container(
      height: 504.0,
      child: new ListView.builder(
        itemCount: calendarStore.eventsByDate[calendarStore.selectedDate].length,
        itemBuilder: itemBuilder,
      ),
    );
  }
}

#/lib/src/month/month.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_flux/flutter_flux.dart';

import 'package:calendar/src/week/week_header.dart';
import 'package:calendar/src/week/week.dart';
import 'package:calendar/src/day/day.dart';
import 'package:calendar/utils.dart';

class Month extends StoreWatcher {
  Month({
    @required this.year,
    @required this.month,
    @required this.date,
  });

  final int year;
  final int month;
  final int date;

  // Functions
  List<Day> _padMonthBeginning(DateTime firstDay) {
    List<Day> days = new List<Day>();
    int firstWeekday = firstDay.weekday;
    DateTime lastDayPrevMonth = new DateTime(year, month, 0);
    DateTime prevMonth = new DateTime(year, month - 1);
    if (firstWeekday < 7) {
      // NOTE(mperrotte): ignore if sunday (no padding needed)
      for (int i = 0; i < firstWeekday; i++) {
        days.add(new Day(
          year: (month == 1) ? year - 1 : year,
          month: prevMonth.month,
          date: lastDayPrevMonth.day - i,
          today: false,
        ));
      }
    }
    return days;
  }

  List<Day> _padMonthEnding(DateTime lastDay) {
    List<Day> days = new List<Day>();
    int lastWeekday = lastDay.weekday;
    DateTime firstDayNextMonth = new DateTime(year, month + 1, 1);
    DateTime nextMonth = new DateTime(year, month + 1);
    int remainingDays = ((6 - lastWeekday) == -1) ? 6 : (6 - lastWeekday);
    for (int i = 0; i < remainingDays; i++) {
      days.add(new Day(
        year: (month == 12) ? year + 1 : year,
        month: nextMonth.month,
        date: firstDayNextMonth.day + i,
        today: false,
      ));
    }
    return days;
  }

  List<Day> _generateMonthDays(DateTime firstDay, DateTime lastDay) {
    List<Day> days = <Day>[];
    for (int i = firstDay.day; i <= lastDay.day; i++) {
      if (i == date) {
        days.add(new Day(year: year, month: month, date: i, today: true));
      } else {
        days.add(new Day(year: year, month: month, date: i, today: false));
      }
    }
    // INFO(mperrotte): add padding to the beginnging of the month
    days.insertAll(0, _padMonthBeginning(firstDay));
    // INFO(mperrotte): add padding to the ending of the month
    days.addAll(_padMonthEnding(lastDay));
    return days;
  }

  List<Week> _generateMonthWeeks() {
    DateTime firstDay = new DateTime(year, month, 1);
    DateTime lastDay = new DateTime(year, month + 1, 0);
    List<Day> monthDays = _generateMonthDays(firstDay, lastDay);
    List<Week> weekList = new List<Week>();
    for (int weeknum = 0; weeknum < (monthDays.length / 7); weeknum++) {
      List<Day> weekDays = new List<Day>();
      // TODO(mperrotte): Look into using `List.sublist` here instead
      for (int weekday = (weeknum * 7); weekday < (weeknum * 7) + 7; weekday++) {
        weekDays.add(monthDays[weekday]);
      }
      weekList.add(new Week(days: weekDays));
    }
    return weekList;
  }

  @override
  void initStores(ListenToStore listenToStore) {
    listenToStore(calendarStoreToken);
  }

  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    List<Week> monthWeeks = _generateMonthWeeks();

    return new Expanded(
      child: new Container(
        margin: const EdgeInsets.all(8.0),
        child: new Column(
          children: [
            [generateWeekHeader('short')],
            monthWeeks,
          ].expand((dynamic week) {
            return week;
          }).toList(),
        ),
      ),
    );
  }
}

#/lib/src/stores.dart

import 'package:flutter_flux/flutter_flux.dart';

import 'package:calendar/src/data/data_model.dart';
import 'package:calendar/src/actions.dart';
import 'package:calendar/src/types.dart';
import 'package:calendar/controllers.dart';

Map<int, List<DataModel>> reduceEvents(List<DataModel> events) {
  Map<int, List<DataModel>> eventsByDate = new Map<int, List<DataModel>>();
  return events.fold(eventsByDate, (prev, event) {
    if (prev.containsKey(event.date)) {
      prev[event.date].add(event);
    } else {
      prev[event.date] = new List<DataModel>();
      prev[event.date].add(event);
    }
    return prev;
  });
}

class CalendarStore extends Store {
  CalendarStore() {
    // NOTE(mperrotte): define action/reducers
    triggerOnAction(switchViewAction, (RenderableView view) {
      _currentView = view;
    });
    triggerOnAction(selectDateAction, (Map<String, int> currentDay) {
      _selectedDate = currentDay['date'];
      _selectedMonth = currentDay['month'];
      _selectedYear = currentDay['year'];
    });
    triggerOnAction(removeEventAction, (int index) {
      _events.removeAt(index);
      _eventsByDate = reduceEvents(_events);
    });

    _fetchDataReducer(Map<String, dynamic> funcs) async {
      dynamic payload = await funcs['payload'];
      List<DataModel> data = funcs['parseData'](payload);
      _events = data;
      _eventsByDate = reduceEvents(_events);
      trigger();
    }

    fetchDataAction.listen(_fetchDataReducer);
  }

  // NOTE(mperrotte): define store properties
  RenderableView _currentView = RenderableView.calendar;
  int _selectedYear;
  int _selectedMonth;
  int _selectedDate;
  List<DataModel> _events = new List<DataModel>();
  Map<int, List<DataModel>> _eventsByDate = new Map<int, List<DataModel>>();

  // NOTE(mperrotte): define selectors
  RenderableView get currentView => _currentView;
  int get selectedYear => _selectedYear;
  int get selectedMonth => _selectedMonth;
  int get selectedDate => _selectedDate;
  List<DataModel> get events => _events;
  Map<int, List<DataModel>> get eventsByDate => _eventsByDate;
  Map<String, int> get currentDay => {
        'year': _selectedYear,
        'month': _selectedMonth,
        'date': _selectedDate,
      };
}

class ControllerStore extends Store {
  ControllerStore() {
    // setCalendarControllerAction.listen((CalendarController controller) {
    //   _calendarController = controller;
    //   trigger();
    // });
    triggerOnAction(setCalendarControllerAction, (CalendarController controller) {
      _calendarController = controller;
    });
    triggerOnAction(setDataControllerAction, (DataController controller) {
      _dataController = controller;
    });
  }

  CalendarController _calendarController;
  DataController _dataController;

  // NOTE(mperrotte): created single selector with type interface
  getController(ControllerType controller) {
    switch (controller) {
      case ControllerType.calendar:
        return _calendarController;
      case ControllerType.data:
        return _dataController;
    }
  }
}

// NOTE(mperrotte): define store tokens
final StoreToken calendarStoreToken = new StoreToken(new CalendarStore());
final StoreToken controllerStoreToken = new StoreToken(new ControllerStore());

#/lib/src/types.dart

// import 'package:flutter/widgets.dart';
import 'package:calendar/src/events/events_view.dart';
import 'package:calendar/src/data/data_model.dart';

const Map<int, Map<String, String>> MonthNames = const {
  0: const {'short': 'Jan', 'long': 'January'},
  1: const {'short': 'Feb', 'long': 'February'},
  2: const {'short': 'Mar', 'long': 'March'},
  3: const {'short': 'Apr', 'long': 'April'},
  4: const {'short': 'May', 'long': 'May'},
  5: const {'short': 'June', 'long': 'June'},
  6: const {'short': 'July', 'long': 'July'},
  7: const {'short': 'Aug', 'long': 'August'},
  8: const {'short': 'Sep', 'long': 'September'},
  9: const {'short': 'Oct', 'long': 'October'},
  10: const {'short': 'Nov', 'long': 'November'},
  11: const {'short': 'Dec', 'long': 'December'}
};

const Map<int, Map<String, String>> DayNames = const {
  0: const {'short': 'Mon', 'long': 'Monday'},
  1: const {'short': 'Tue', 'long': 'Tuesday'},
  2: const {'short': 'Wed', 'long': 'Wednesday'},
  3: const {'short': 'Thur', 'long': 'Thursday'},
  4: const {'short': 'Fri', 'long': 'Friday'},
  5: const {'short': 'Sat', 'long': 'Saturday'},
  6: const {'short': 'Sun', 'long': 'Sunday'}
};

enum RenderableView {
  calendar,
  events,
  event,
}

enum ControllerType {
  calendar,
  data,
}

typedef String DataCallback();
typedef EventsView RenderEventsCallback(List<DataModel> events);
typedef dynamic FetchDataCallback(String uri);
typedef String GetDataUriCallback();
typedef List<DataModel> ParseDataCallback(dynamic payload);
typedef List<DataModel> FilterEventsCallback(List<DataModel> events, Map<String, int> currentDay);

#/lib/src/week/week.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:calendar/src/day/day.dart';

class Week extends StatelessWidget {
  Week({
    @required this.days,
  });

  final List<Day> days;

  @override
  Widget build(BuildContext context) {
    return new Row(
      children: days,
    );
  }
}

#/lib/src/week/week_header.dart

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'package:calendar/src/day/day_header.dart';
import 'package:calendar/utils.dart';
import 'week.dart';

class WeekHeader extends Week {
  WeekHeader({
    @required this.days,
  });

  final List<DayHeader> days;

  @override
  Widget build(BuildContext context) {
    return new Row(
      children: days,
    );
  }
}

WeekHeader generateWeekHeader(String nameFormat) {
  assert(nameFormat != null);
  List<DayHeader> days = new List<DayHeader>();
  // INFO(mperrotte): add sunday first
  days.add(new DayHeader(
    day: DayNames[0][nameFormat],
  ));
  for (int i = 0; i < 6; i++) {
    days.add(new DayHeader(
      day: DayNames[i][nameFormat],
    ));
  }
  return new WeekHeader(
    days: days,
  );
}

#/lib/utils.dart

/// Utility Exports
export 'src/actions.dart';
export 'src/stores.dart';
export 'src/types.dart';

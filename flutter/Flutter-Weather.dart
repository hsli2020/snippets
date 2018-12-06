#/lib/main.dart

import 'package:flutter/material.dart';
import 'package:sunshine/ui/HomePage.dart';

void main() {
  runApp(new MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      title: 'Flutter Demo',
      theme: theme,
      home: new HomePage(),
    );
  }
}

class Pages {
  static const HOME = "/";
  static const DETAIL = "/detail";
}

var theme = new ThemeData (
  primarySwatch: Colors.blue,
  fontFamily: 'Dosis'
);

#/lib/ui/HomePage.dart

import 'package:flutter/material.dart';
import 'package:sunshine/ui/forecast/Forecast.dart';
import 'package:sunshine/ui/weather/Weather.dart';

class HomePage extends StatelessWidget{
  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      body: new Container(
        child: new Column(
          children: <Widget>[
            new AspectRatio(child: new Weather(), aspectRatio: 750.0/805.0),
            new Expanded(child: new Forecast()),
          ],
        )
      ),
    );
  }
}

#/lib/res/Res.dart

import 'dart:ui';

/* App specific color suite */
class $Colors {
  static const Color empress = const Color(0xFF756C6F);
  static const Color ghostWhite = const Color(0xFFF6F6F7);
  static const Color quartz = const Color(0xFFE1E0E8);
  static const Color blueHaze = const Color(0xFFB3B2C2);
  static const Color lavender = const Color(0xFFCBCAD6);
  static const Color blueParis = const Color(0xDD595877);
}

/* Image paths */
class $Asset {
  static const String dotFull = "assets/img/dotBlueishFull.png";
  static const String dotEmpty = "assets/img/dotEmpty.png";
  static const String backgroundParis = "assets/img/parisback.png";
  static const String pressure = "assets/img/pres.png";
  static const String humidity = "assets/img/water.png";
  static const String wind = "assets/img/wind.png";
}

#/lib/model/Condition.dart

class Condition {
  int id;
  String description;

  Condition(this.id, this.description);

  String getAssetString() {
    if (id >= 200 && id <= 299)
      return "assets/img/d7s.png";
    else if (id >= 300 && id <= 399)
      return "assets/img/d6s.png";
    else if (id >= 500 && id <= 599)
      return "assets/img/d5s.png";
    else if (id >= 600 && id <= 699)
      return "assets/img/d8s.png";
    else if (id >= 700 && id <= 799)
      return "assets/img/d9s.png";
    else if (id >= 300 && id <= 399)
      return "assets/img/d6s.png";
    else if (id == 800)
      return "assets/img/d1s.png";
    else if (id == 801)
      return "assets/img/d2s.png";
    else if (id == 802)
      return "assets/img/d3s.png";
    else if (id == 803 || id == 804)
      return "assets/img/d4s.png";

    print("Unknown condition ${id}");
    return "assets/img/n1s.png";
  }
}

#/lib/model/ForecastData.dart

import 'dart:convert';

import 'package:sunshine/model/Condition.dart';

class ForecastData {

  List<ForecastWeather> forecastList;

  ForecastData(this.forecastList);

  static ForecastData deserialize(String json) {
    JsonDecoder decoder = new JsonDecoder();
    var map = decoder.convert(json);

    var list = map["list"];
    List<ForecastWeather> forecast = [];

    for (var weatherMap in list) {
      forecast.add(ForecastWeather._deserialize(weatherMap));
    }

    return new ForecastData(forecast);
  }
}

class ForecastWeather {
  String temperature;
  Condition condition;
  DateTime dateTime;

  double pressure;
  double humidity;
  double wind;
  //Wind, rain, etc.

  ForecastWeather(this.temperature, this.condition, this.dateTime, this.pressure, this.humidity, this.wind);

  static ForecastWeather _deserialize(Map<String, dynamic> map) {
    String description = map["weather"][0]["description"];
    int conditionId = map["weather"][0]["id"];
    Condition condition = new Condition(conditionId, description);

    double temperature = map["main"]["temp"].toDouble();
    double humidity = map["main"]["humidity"].toDouble();
    double pressure = map["main"]["pressure"].toDouble();
    double wind = map["wind"]["speed"].toDouble();
    int epochTimeMs = map["dt"]*1000;
    DateTime dateTime = new DateTime.fromMillisecondsSinceEpoch(epochTimeMs);

    return new ForecastWeather(temperature.toString(), condition, dateTime, pressure, humidity, wind);
  }
}

#/lib/model/WeatherData.dart

import 'dart:convert';

import 'package:sunshine/model/Condition.dart';

class WeatherData {
  String temperature;
  Condition condition;

  WeatherData(this.temperature, this.condition);

  static WeatherData deserialize(String json) {
    JsonDecoder decoder = new JsonDecoder();
    var map = decoder.convert(json);

    String description = map["weather"][0]["description"];
    int id = map["weather"][0]["id"];
    Condition condition = new Condition(id, description);

    double temperature = map["main"]["temp"].toDouble();

    return new WeatherData(temperature.toString(), condition);
  }
}

#/lib/network/ApiClient.dart

import 'package:http/http.dart' as http;

import 'package:sunshine/model/WeatherData.dart';
import 'package:sunshine/model/ForecastData.dart';

import 'dart:async';

class ApiClient {
  static ApiClient _instance;

  static ApiClient getInstance() {
    if (_instance == null) {
      _instance = new ApiClient();
    }
    return _instance;
  }


  Future<WeatherData> getWeather() async {
    http.Response response = await http.get(
      Uri.encodeFull(Endpoints.WEATHER),
      headers: {
        "Accept": "application/json"
      }
    );

    return WeatherData.deserialize(response.body);
  }

  Future<ForecastData> getForecast() async {
    http.Response response = await http.get(
      Uri.encodeFull(Endpoints.FORECAST),
      headers: {
        "Accept": "application/json"
      }
    );

    return ForecastData.deserialize(response.body);
  }

}

class Endpoints {
  static const _ENDPOINT = "http://api.openweathermap.org/data/2.5";
  static const WEATHER = _ENDPOINT + "/weather?lat=43.509645&lon=16.445783&APPID=af29567e139fe06b6c2d050515cdff0c&units=metric";
  static const FORECAST = _ENDPOINT + "/forecast?lat=43.509645&lon=16.445783&APPID=af29567e139fe06b6c2d050515cdff0c&units=metric";
}

#/lib/store/ForecastStore.dart

import 'dart:async';

import 'package:flutter_flux/flutter_flux.dart';
import 'package:sunshine/model/ForecastData.dart';
import 'package:sunshine/network/ApiClient.dart';

class ForecastStore extends Store {

  /// Forecast by day: list of days with each containing list of
  /// [ForecastWeather] through day
  List<List<ForecastWeather>> forecastByDay;

  ForecastStore() {

    triggerOnAction(updateForecast, (dynamic) {
      _updateForecast();
    });
  }

  _updateForecast() {
    Future<ForecastData> fForecastData = ApiClient.getInstance().getForecast();
    fForecastData.then((content) {
      ForecastData forecastData = content;
      this.forecastByDay = groupForecastListByDay(forecastData);
      trigger();
    }).catchError((e) {
      print(e);
    });
  }

  static List<List<ForecastWeather>> groupForecastListByDay(
      ForecastData forecastData) {
    if (forecastData == null) return null;

    List<List<ForecastWeather>> forecastListByDay = [];
    final forecastList = forecastData.forecastList;

    int currentDay = forecastList[0].dateTime.day;
    List<ForecastWeather> intermediateList = [];

    for (var forecast in forecastList) {
      if (currentDay == forecast.dateTime.day) {
        intermediateList.add(forecast);
      } else {
        forecastListByDay.add(intermediateList);
        currentDay = forecast.dateTime.day;
        intermediateList = [];
        intermediateList.add(forecast);
      }
    }

    forecastListByDay.add(intermediateList);
    return forecastListByDay;
  }

}

final Action updateForecast = new Action();
final StoreToken forecastStoreToken = new StoreToken(new ForecastStore());

#/lib/store/StatelessStoreWidget.dart

import 'package:flutter/src/widgets/framework.dart';
import 'package:flutter_flux/flutter_flux.dart';

/// Workaround for flutter_flux implementation when state class is dealing
/// with store.
///
/// The widget then should be like regulare StateWidget, delegating build
/// method to State class. StoreWatcher works like this, but has issue where
/// that requires you to have concrete implementation of build and initStores
/// methods, even though they are not used and we only need createState method.

abstract class StatelessStoreWidget extends StoreWatcher {

  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    // This shouldn't be called if setState is implemented
    throw new Exception("Implement setState method in your widget");
  }

  @override
  void initStores(ListenToStore listenToStore) {
    // This shouldn't be called if setState is implemented
    throw new Exception("Implement setState method in your widget");
  }
}

#/lib/store/WeatherStore.dart

import 'dart:async';

import 'package:flutter_flux/flutter_flux.dart';
import 'package:sunshine/model/Condition.dart';
import 'package:sunshine/model/WeatherData.dart';
import 'package:sunshine/network/ApiClient.dart';

class WeatherStore extends Store {

  WeatherData weatherData;

  WeatherStore() {
    // TODO make loading widget from here
    this.weatherData = new WeatherData("", new Condition(0, "Loading"));

    triggerOnAction(actionUpdateWeather, (dynamic) {
      _updateWeather();
    });
  }

  void _updateWeather() {
    var apiClient = ApiClient.getInstance();
    Future<WeatherData> fWeatherData = apiClient.getWeather();
    fWeatherData
        .then((content) {
          this.weatherData = content;
          trigger();
    }).catchError((e) {
        this.weatherData = null;
        // todo trigger error
    });
  }
}

// Token and actions
final Action actionUpdateWeather = new Action();
final StoreToken weatherStoreToken = new StoreToken(new WeatherStore());

#/lib/ui/forecast/Forecast.dart

import 'package:flutter/material.dart';
import 'package:flutter_flux/flutter_flux.dart';

import 'package:sunshine/res/Res.dart';
import 'package:sunshine/store/ForecastStore.dart';

import 'package:sunshine/ui/forecast/ForecastPager.dart';
import 'package:flutter_flux/src/store_watcher.dart';

class Forecast extends StoreWatcher {

  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    ForecastStore store = stores[forecastStoreToken];
    if (store.forecastByDay == null) return new Container();

    return new Stack(
      children: <Widget>[
        new Image(
          image: new AssetImage("assets/img/bottom_parisback.png"),
          fit: BoxFit.fitWidth,
        ),
        new Container(
          child: new ForecastPager(store.forecastByDay),
          decoration: new BoxDecoration(
              color: Colors.white,
              shape: BoxShape.rectangle,
              borderRadius: new BorderRadius.only(
                topLeft: new Radius.circular(15.0),
                topRight: new Radius.circular(15.0),
              ),
              boxShadow: <BoxShadow>[
                new BoxShadow(
                    color: Colors.black12,
                    blurRadius: 10.0,
                    offset: new Offset(0.0, -10.0))
              ]),
        ),
      ],
    );
  }

  @override
  void initStores(ListenToStore listenToStore) {
    listenToStore(forecastStoreToken);
    updateForecast.call(); // Initial load
  }
}

#/lib/ui/forecast/ForecastList.dart

import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

import 'package:sunshine/model/ForecastData.dart';
import 'package:sunshine/res/Res.dart';
import 'package:sunshine/ui/ForecastDetailPage.dart';

import 'package:sunshine/ui/widgets/TextWithExponent.dart';

final timeFormat = new DateFormat('HH');

class ForecastList extends StatelessWidget {
  final List<ForecastWeather> _forecast;

  ForecastList(this._forecast);

  @override
  Widget build(BuildContext context) {
    return new ListView.builder(
      itemBuilder: (BuildContext context, int index) =>
          new _ForecastListItem(_forecast[index]),
      itemCount: _forecast == null ? 0 : _forecast.length,
    );
  }
}

class _ForecastListItem extends StatelessWidget {
  final ForecastWeather weather;

  _ForecastListItem(this.weather);

  void clicked(BuildContext context) {
    Navigator.of(context).push(ForecastDetailPage.getRoute(weather));
  }

  @override
  Widget build(BuildContext context) {
    final time = timeFormat.format(weather.dateTime);

    return new Material(
        child: new InkWell(
            onTap: () => clicked(context),
            child: new Container(
                height: 65.0,
                padding:
                    new EdgeInsets.symmetric(horizontal: 12.0, vertical: 8.0),
                child: new Stack(
                  children: <Widget>[
                    new Align(
                      child: new TextWithExponent(time, "h"),
                      alignment: FractionalOffset.centerLeft,
                    ),
                    new Positioned(
                      child: new Row(
                        children: <Widget>[
                          new Container(
                            child: new Image.asset(
                              weather.condition.getAssetString(),
                              height: 46.0,
                              width: 46.0,
                              fit: BoxFit.scaleDown,
                              color: $Colors.blueParis,
                            ),
                            margin: new EdgeInsets.only(right: 8.0),
                          ),
                          new Container(
                            width: 80.0,
                            alignment: FractionalOffset.centerRight,
                            child: new Text(
                              weather.temperature + "°C",
                              style: new TextStyle(fontSize: 20.0),
                            ),
                          ),
                        ],
                      ),
                      right: 0.0,
                    )
                  ],
                ))));
  }
}

#/lib/ui/forecast/ForecastPager.dart

import 'package:flutter/material.dart';

import 'package:sunshine/model/ForecastData.dart';
import 'package:sunshine/res/Res.dart';

import 'package:intl/intl.dart';
import 'package:sunshine/ui/forecast/ForecastList.dart';
import 'package:sunshine/ui/widgets/DotPageIndicator.dart';
import 'package:sunshine/ui/widgets/TextWithExponent.dart';

final weekdayFormat = new DateFormat('EEE');

class ForecastPager extends StatefulWidget {
  var _forecastByDay;
  ForecastPager(this._forecastByDay);

  @override
  _ForecastPagerState createState() {
    return new _ForecastPagerState(_forecastByDay);
  }
}

class _ForecastPagerState extends State<ForecastPager> {
  var currentPage = 0;
  List<List<ForecastWeather>> _forecastByDay;

  _ForecastPagerState(this._forecastByDay);

  @override
  void didUpdateWidget(ForecastPager oldWidget) {
    setState(() {
      this._forecastByDay = widget._forecastByDay;
    });
  }

  @override
  Widget build(BuildContext context) {
    DateTime currentDateTime;
    final pageCount = _forecastByDay != null ? _forecastByDay.length : 0;

    if (_forecastByDay != null) {
      if (_forecastByDay[currentPage].length > 0) {
        currentDateTime = _forecastByDay[currentPage][0].dateTime;
      }
    }

    return new Column(
      children: <Widget>[
        new _ForecastWeekTabs(currentDateTime, currentPage, pageCount),
        new Expanded(
            child: new PageView.builder(
              itemBuilder: (BuildContext context, int index) =>
              new ForecastList(_forecastByDay[index]),
              itemCount: pageCount,
              scrollDirection: Axis.horizontal,
              onPageChanged: (index) => this.setState(() {this.currentPage = index;}),
            )),
      ],
    );
  }
}

class _ForecastWeekTabs extends StatelessWidget {
  final DateTime dateTime;
  final int currentPage;
  final int pageCount;

  _ForecastWeekTabs(this.dateTime, this.currentPage, this.pageCount);


  @override
  Widget build(BuildContext context) {
    final textStyle = new TextStyle(fontSize: 24.0);
    final int dayOfMonth = dateTime != null ? dateTime.day : 0;
    String dayMonthSuffix = "";

    final weekDay = weekdayFormat.format(dateTime).toString();
    if (dayOfMonth == 1) {
      dayMonthSuffix += "st";
    } else if (dayOfMonth == 2) {
      dayMonthSuffix += "nd";
    } else {
      dayMonthSuffix += "th";
    }

    return new Container(
      child: new Container(
        child: new Stack(
          children: <Widget>[
            new Container(
              child: new Text(weekDay, style: textStyle),
              padding: new EdgeInsets.only(left: 36.0),
            ),
            new Align(
              child: new DotPageIndicator(this.currentPage, this.pageCount),
              alignment: FractionalOffset.center,
            ),
            new Positioned(
                child: new TextWithExponent(
                  dayOfMonth.toString(),
                  dayMonthSuffix,
                  textSize: 24.0,
                  exponentTextSize: 18.0,
                ),
                right: 36.0),
          ],
        ),
        padding: new EdgeInsets.symmetric(vertical: 8.0),
      ),
      decoration: new BoxDecoration(
          border: new Border(bottom: new BorderSide(color: Colors.black12))),
    );
  }
}

#/lib/ui/ForecastDetailPage.dart

import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:sunshine/model/ForecastData.dart';
import 'package:sunshine/ui/forecast_detail/ForecastDetail.dart';

final monthFormat = new DateFormat('MMMM');

class ForecastDetailPage extends StatelessWidget {
  ForecastWeather weather;

  ForecastDetailPage(this.weather);

  static MaterialPageRoute getRoute(ForecastWeather forecastWeather) {
    return new MaterialPageRoute(builder: (BuildContext context) {
      return new ForecastDetailPage(forecastWeather);
    });
  }

  @override
  Widget build(BuildContext context) {
    var title = weather.dateTime.hour.toString() + "h";
    var month = monthFormat.format(weather.dateTime);
    title += " • " + weather.dateTime.day.toString() + " " + month;

    return new Scaffold(
      appBar: new AppBar(title: new Text(title)),
      body: new Container(
          decoration: new BoxDecoration(
              gradient: new LinearGradient(
                  colors: [
                    const Color(0x99338600),
                    const Color(0x9900CCFF),
                    const Color(0xAA0077FF),
                  ],
                  begin: const FractionalOffset(0.0, 0.0),
                  end: const FractionalOffset(0.7, 1.0),
                  stops: [0.0, 0.7, 1.0],
                  tileMode: TileMode.clamp)),
          child: new Center(
            child: new ForecastDetail(weather),
          )),
    );
  }
}

#/lib/ui/forecast_detail/ForecastDetail.dart

import 'package:flutter/material.dart';
import 'package:sunshine/model/ForecastData.dart';
import 'package:sunshine/res/Res.dart';

class ForecastDetail extends StatelessWidget {
  final ForecastWeather weather;

  ForecastDetail(this.weather);

  @override
  Widget build(BuildContext context) {
    return new Container(
      child: new Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          new _MainWeatherInfo(weather),
          new Container(margin: const EdgeInsets.all(20.0), height: 1.0, width: 220.0, color: Colors.black45,),
          new _WeatherInfo(weather.wind.toString(), $Asset.wind),
          new _WeatherInfo(weather.pressure.toString(), $Asset.pressure),
          new _WeatherInfo(weather.humidity.toString(), $Asset.humidity),
        ],
      ),
    );
  }
}

class _MainWeatherInfo extends StatelessWidget {
  final ForecastWeather weather;

  _MainWeatherInfo(this.weather);

  @override
  Widget build(BuildContext context) {
    return new Container(
        child: new Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
          new Image.asset(
            this.weather.condition.getAssetString(),
            width: 50.0,
            color: $Colors.blueParis,
          ),
          new Padding(
            padding: new EdgeInsets.only(left: 16.0),
            child: new Text(
              this.weather.condition.description,
              style: new TextStyle(fontSize: 32.0),
            ),
          )
        ]));
  }
}

class _WeatherInfo extends StatelessWidget {
  final String info;
  final String imageAsset;

  _WeatherInfo(this.info, this.imageAsset);

  @override
  Widget build(BuildContext context) {
    return new Container(child: new Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        new Image.asset(imageAsset, width: 30.0, color: $Colors.blueParis,),
        new Padding(
          padding: const EdgeInsets.only(left: 16.0, bottom: 12.0, top: 12.0),
          child: new Text(info, style: new TextStyle(fontSize: 20.0),),
        )
      ],
    ));
  }
}

#/lib/ui/weather/Weather.dart

import 'package:flutter/material.dart';
import 'package:sunshine/model/Condition.dart';
import 'package:sunshine/model/WeatherData.dart';

import 'package:sunshine/res/Res.dart';
import 'package:sunshine/store/WeatherStore.dart';
import 'package:sunshine/ui/widgets/TextWithExponent.dart';
import 'package:flutter_flux/flutter_flux.dart';

class Weather extends StoreWatcher {
  @override
  Widget build(BuildContext context, Map<StoreToken, Store> stores) {
    WeatherStore store = stores[weatherStoreToken];
    WeatherData weatherData = store.weatherData;

    return new Container(
        decoration: new BoxDecoration(
            image: new DecorationImage(
          image: new AssetImage($Asset.backgroundParis),
          fit: BoxFit.cover,
        )),
        child: new Row(
          children: <Widget>[
            new Flexible(
              child: new WeatherInfo(weatherData),
            ),
          ],
        ));
  }

  @override
  void initStores(ListenToStore listenToStore) {
    listenToStore(weatherStoreToken);
    actionUpdateWeather.call(); // Initial load
  }
}

class WeatherInfo extends StatelessWidget {
  WeatherInfo(WeatherData this._weather);

  final WeatherData _weather;

  @override
  Widget build(BuildContext context) {
    final roundedTemperature = this._weather.temperature.split(".")[0] + "°";
    final condition = '${this._weather.condition.description[0]
        .toUpperCase()}${this._weather
        .condition.description.substring(1)}';

    return new Container(
      child: new Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          new Text(
            "Paris",
            style: new TextStyle(
                fontSize: 21.0,
                fontWeight: FontWeight.w700,
                color: $Colors.blueParis),
          ),
          new Text(
            condition,
            style: new TextStyle(
              fontSize: 18.0,
              color: $Colors.blueParis,
            ),
          ),
          new Text(roundedTemperature,
              style: new TextStyle(
                  fontSize: 72.0,
                  color: $Colors.blueParis,
                  fontFamily: "Roboto")),
        ],
      ),
      padding: new EdgeInsets.only(left: 64.0),
    );
  }
}

#/lib/ui/widgets/DotPageIndicator.dart

import 'package:flutter/material.dart';

import 'package:sunshine/res/Res.dart';

class DotPageIndicator extends StatelessWidget {
  final int currentPage;
  final int pagesCount;

  DotPageIndicator(this.currentPage, this.pagesCount);

  @override
  Widget build(BuildContext context) {
    var dots = <Widget>[];
    final dotEmpty = new Flexible(
        child: new Image(
          image: new AssetImage($Asset.dotEmpty),
          width: 15.0,
          height: 15.0,
        ));

    final dotFull = new Flexible(
        child: new Image(
          image: new AssetImage($Asset.dotFull),
          width: 15.0,
          height: 15.0,
        ));

    for (var i=0; i<pagesCount; i++) {
      if (i == currentPage)
        dots.add(dotFull);
      else
        dots.add(dotEmpty);
    }

    return new Container(
      padding: new EdgeInsets.only(top: 8.0),
      alignment: FractionalOffset.center,
      child: new Row(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: dots,
      ),
    );
  }
}

#/lib/ui/widgets/GradientAppBar.dart

import 'package:flutter/material.dart';

class GradientAppBar extends StatelessWidget {
  final String title;
  final double barHeight = 66.0;

  GradientAppBar(this.title);

  @override
  Widget build(BuildContext context) {
    final double statusBarHeight = MediaQuery
        .of(context)
        .padding
        .top;

    return new Container(
      padding: new EdgeInsets.only(top: statusBarHeight),
      height: barHeight + statusBarHeight,
      decoration: new BoxDecoration(
        gradient: new LinearGradient(colors:[const Color(0x553366FF), const Color(0x5500CCFF),],
        begin: const FractionalOffset(0.0, 0.0),
        end: const FractionalOffset(1.0, 0.0),
        stops: [0.0, 1.0],
        tileMode: TileMode.clamp)
      ),
      child: new Center(
        child: new Text(
            title,
            style: const TextStyle(
                color: Colors.black54,
                fontWeight: FontWeight.w400,
                fontSize: 34.0
            )
        ),
      ),
    );
  }
}

#/lib/ui/widgets/TextWithExponent.dart

import 'package:flutter/material.dart';

class TextWithExponent extends StatelessWidget {
  final String text;
  final String exponentText;
  final double textSize;
  final double exponentTextSize;

  TextWithExponent(this.text, this.exponentText, {this.textSize = 25.0, this.exponentTextSize = 18.0});

  @override
  Widget build(BuildContext context) {
    return new Row(children: <Widget>[
      new Text(text, style: new TextStyle(fontSize: textSize),),
      new Container(
        child: new Text(exponentText, style: new TextStyle(fontSize: exponentTextSize)),
        margin: new EdgeInsets.only(bottom: (textSize - exponentTextSize)),
      ),
    ],);
  }
}

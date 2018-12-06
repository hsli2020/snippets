// https://github.com/chenyuantao/flutter_calculator

// main.dart

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'number.dart';
import 'operator.dart';
import 'result.dart';

void main() => runApp(new MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return CupertinoApp(
      title: 'Flutter Calculator',
      debugShowCheckedModeBanner: false,
      home: CalculatorPage(),
    );
  }
}

class CalculatorPage extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => CalculatorState();
}

class CalculatorState extends State<StatefulWidget> {
  List<Result> results = [];
  String currentDisplay = '0';

  onResultButtonPressed(display) {
    if (results.length > 0) {
      var result = results[results.length - 1];
      if (display == '=') {
        result.result = result.oper.calculate(
            double.parse(result.firstNum), double.parse(result.secondNum));
      } else if (display == '<<') {
        results.removeLast();
      }
    }
    pickCurrentDisplay();
  }

  onOperatorButtonPressed(Operator oper) {
    if (results.length > 0) {
      var result = results[results.length - 1];
      if (result.result != null) {
        var newRes = Result();
        newRes.firstNum = currentDisplay;
        newRes.oper = oper;
        results.add(newRes);
      } else if (result.firstNum != null) {
        result.oper = oper;
      }
    }
    pickCurrentDisplay();
  }

  onNumberButtonPressed(Number number) {
    var result = results.length > 0 ? results[results.length - 1] : Result();
    if (result.firstNum == null || result.oper == null) {
      result.firstNum = number.apply(currentDisplay);
    } else if (result.result == null) {
      if (result.secondNum == null) {
        currentDisplay = '0';
      }
      result.secondNum = number.apply(currentDisplay);
    } else {
      var newRes = Result();
      currentDisplay = '0';
      newRes.firstNum = number.apply(currentDisplay);
      results.add(newRes);
    }
    if (results.length == 0) {
      results.add(result);
    }
    pickCurrentDisplay();
  }

  pickCurrentDisplay() {
    this.setState(() {
      var display = '0';
      results.removeWhere((item) =>
          item.firstNum == null && item.oper == null && item.secondNum == null);
      if (results.length > 0) {
        var result = results[results.length - 1];
        if (result.result != null) {
          display = format(result.result);
        } else if (result.secondNum != null && result.oper != null) {
          display = result.secondNum;
        } else if (result.firstNum != null) {
          display = result.firstNum;
        }
      }
      currentDisplay = display;
    });
  }

  String format(num number) {
    if (number == number.toInt()) {
      return number.toInt().toString();
    }
    return number.toString();
  }

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      child: Container(
          color: Colors.grey[100],
          child: Column(
            children: <Widget>[
              Expanded(
                key: Key('Current_Display'),
                flex: 2,
                child: FractionallySizedBox(
                  widthFactor: 1.0,
                  heightFactor: 1.0,
                  child: Container(
                    color: Colors.lightBlue[300],
                    alignment: Alignment.bottomRight,
                    padding: const EdgeInsets.all(16.0),
                    child: ResultDisplay(result: currentDisplay),
                  ),
                ),
              ),
              Expanded(
                  key: Key('History_Display'),
                  child: FractionallySizedBox(
                      widthFactor: 1.0,
                      heightFactor: 1.0,
                      child: Container(
                        color: Colors.black54,
                        child: ListView(
                            scrollDirection: Axis.horizontal,
                            reverse: true,
                            // mainAxisAlignment: MainAxisAlignment.end,
                            children: results.reversed.map((result) {
                              return HistoryBlock(result: result);
                            }).toList()),
                      )),
                  flex: 1),
              Expanded(
                  key: Key('Number_Button_Line_1'),
                  child: NumberButtonLine(
                    array: [
                      NormalNumber('1'),
                      NormalNumber('2'),
                      NormalNumber('3')
                    ],
                    onPress: onNumberButtonPressed,
                  ),
                  flex: 1),
              Expanded(
                  key: Key('Number_Button_Line_2'),
                  child: NumberButtonLine(
                    array: [
                      NormalNumber('4'),
                      NormalNumber('5'),
                      NormalNumber('6')
                    ],
                    onPress: onNumberButtonPressed,
                  ),
                  flex: 1),
              Expanded(
                  key: Key('Number_Button_Line_3'),
                  child: NumberButtonLine(
                    array: [
                      NormalNumber('7'),
                      NormalNumber('8'),
                      NormalNumber('9')
                    ],
                    onPress: onNumberButtonPressed,
                  ),
                  flex: 1),
              Expanded(
                  key: Key('Number_Button_Line_4'),
                  child: NumberButtonLine(
                    array: [SymbolNumber(), NormalNumber('0'), DecimalNumber()],
                    onPress: onNumberButtonPressed,
                  ),
                  flex: 1),
              Expanded(
                  key: Key('Operator_Group'),
                  child: OperatorGroup(onOperatorButtonPressed),
                  flex: 1),
              Expanded(
                  key: Key('Result_Button_Area'),
                  child: Row(
                    children: <Widget>[
                      ResultButton(
                        display: '<<',
                        color: Colors.red,
                        onPress: onResultButtonPressed,
                      ),
                      ResultButton(
                          display: '=',
                          color: Colors.green,
                          onPress: onResultButtonPressed),
                    ],
                  ),
                  flex: 1)
            ],
          )),
    );
  }
}

// number.dart

import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';

typedef void PressOperationCallback(Number number);

abstract class Number {
  String display;
  String apply(String original);
}

class NormalNumber extends Number {
  NormalNumber(String display) {
    this.display = display;
  }
  apply(original) {
    if (original == '0') {
      return display;
    } else {
      return original + display;
    }
  }
}

class SymbolNumber extends Number {
  @override
  String get display => '+/-';
  @override
  String apply(String original) {
    int index = original.indexOf('-');
    if (index == -1 && original != '0') {
      return '-' + original;
    } else {
      return original.replaceFirst(new RegExp(r'-'), '');
    }
  }
}

class DecimalNumber extends Number {
  @override
  String get display => ('.');
  @override
  String apply(String original) {
    int index = original.indexOf('.');
    if (index == -1) {
      return original + '.';
    } else if (index == original.length) {
      return original.replaceFirst(new RegExp(r'.'), '');
    } else {
      return original;
    }
  }
}

class NumberButtonLine extends StatelessWidget {
  NumberButtonLine({@required this.array, this.onPress})
      : assert(array != null);
  final List<Number> array;
  final PressOperationCallback onPress;

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Row(children: <Widget>[
        NumberButton(
            number: array[0],
            pad: EdgeInsets.only(bottom: 4.0),
            onPress: onPress),
        NumberButton(
            number: array[1],
            pad: EdgeInsets.only(left: 4.0, right: 4.0, bottom: 4.0),
            onPress: onPress),
        NumberButton(
            number: array[2],
            pad: EdgeInsets.only(bottom: 4.0),
            onPress: onPress)
      ]),
    );
  }
}

class NumberButton extends StatefulWidget {
  const NumberButton({@required this.number, @required this.pad, this.onPress})
      : assert(number != null),
        assert(pad != null);
  final Number number;
  final EdgeInsetsGeometry pad;
  final PressOperationCallback onPress;
  @override
  State<StatefulWidget> createState() => new NumberButtonState();
}

class NumberButtonState extends State<NumberButton> {
  bool pressed = false;
  @override
  Widget build(BuildContext context) {
    return Expanded(
        flex: 1,
        child: Padding(
          padding: widget.pad,
          child: GestureDetector(
            onTap: () {
              if (widget.onPress != null) {
                widget.onPress(widget.number);
                setState(() {
                  pressed = true;
                });
                Future.delayed(
                    const Duration(milliseconds: 200),
                    () => setState(() {
                          pressed = false;
                        }));
              }
            },
            child: Container(
              alignment: Alignment.center,
              color: pressed ? Colors.grey[200] : Colors.white,
              child: Text(
                '${widget.number.display}',
                style: TextStyle(fontSize: 30.0, color: Colors.grey),
              ),
            ),
          ),
        ));
  }
}

// operator.dart

import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';

typedef void PressOperationCallback(Operator oper);

abstract class Operator {
  String display;
  Color color;
  num calculate(num first, num second);
}

class AddOperator extends Operator {
  @override
  String get display => '+';
  @override
  Color get color => Colors.pink[300];
  @override
  calculate(first, second) {
    return first + second;
  }
}

class SubOperator extends Operator {
  @override
  String get display => '-';
  @override
  Color get color => Colors.orange[300];
  @override
  calculate(first, second) {
    return first - second;
  }
}

class MultiOperator extends Operator {
  @override
  String get display => 'x';
  @override
  Color get color => Colors.lightBlue[300];
  @override
  calculate(first, second) {
    return first * second;
  }
}

class DivisionOperator extends Operator {
  @override
  String get display => '÷';
  @override
  Color get color => Colors.purple[300];
  @override
  calculate(first, second) {
    return first / second;
  }
}

class OperatorGroup extends StatelessWidget {
  OperatorGroup(this.onOperatorButtonPressed);
  final PressOperationCallback onOperatorButtonPressed;
  @override
  Widget build(BuildContext context) {
    return Row(
      children: <Widget>[
        OperatorButton(
          oper: AddOperator(),
          onPress: onOperatorButtonPressed,
        ),
        OperatorButton(
          oper: SubOperator(),
          onPress: onOperatorButtonPressed,
        ),
        OperatorButton(
          oper: MultiOperator(),
          onPress: onOperatorButtonPressed,
        ),
        OperatorButton(
          oper: DivisionOperator(),
          onPress: onOperatorButtonPressed,
        ),
      ],
    );
  }
}

class OperatorButton extends StatefulWidget {
  OperatorButton({@required this.oper, this.onPress})
      : assert(Operator != null);
  final Operator oper;
  final PressOperationCallback onPress;

  @override
  State<StatefulWidget> createState() => OperatorButtonState();
}

class OperatorButtonState extends State<OperatorButton> {
  bool pressed = false;

  @override
  Widget build(BuildContext context) {
    return Expanded(
        flex: 1,
        child: Padding(
            padding: EdgeInsets.all(16.0),
            child: GestureDetector(
              onTap: () {
                if (widget.onPress != null) {
                  widget.onPress(widget.oper);
                  setState(() {
                    pressed = true;
                  });
                  Future.delayed(
                      const Duration(milliseconds: 200),
                      () => setState(() {
                            pressed = false;
                          }));
                }
              },
              child: Container(
                alignment: Alignment.center,
                decoration: BoxDecoration(
                    color: pressed
                        ? Color.alphaBlend(Colors.white30, widget.oper.color)
                        : widget.oper.color,
                    borderRadius: BorderRadius.all(Radius.circular(100.0))),
                child: Text(
                  '${widget.oper.display}',
                  style: TextStyle(fontSize: 30.0, color: Colors.white),
                ),
              ),
            )));
  }
}

// result.dart

import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'operator.dart';

typedef void PressOperationCallback(display);

class Result {
  Result();
  String firstNum;
  String secondNum;
  Operator oper;
  num result;
}

class ResultButton extends StatefulWidget {
  ResultButton({@required this.display, @required this.color, this.onPress});
  final String display;
  final Color color;
  final PressOperationCallback onPress;

  @override
  State<StatefulWidget> createState() => ResultButtonState();
}

class ResultButtonState extends State<ResultButton> {
  bool pressed = false;

  @override
  Widget build(BuildContext context) {
    return Expanded(
        flex: 1,
        child: Padding(
            padding: EdgeInsets.only(
                left: 10.0, right: 10.0, top: 10.0, bottom: 24.0),
            child: GestureDetector(
              onTap: () {
                if (widget.onPress != null) {
                  widget.onPress(widget.display);
                  setState(() {
                    pressed = true;
                  });
                  Future.delayed(
                      const Duration(milliseconds: 200),
                      () => setState(() {
                            pressed = false;
                          }));
                }
              },
              child: Container(
                alignment: Alignment.center,
                decoration: BoxDecoration(
                    color: pressed ? Colors.grey[200] : null,
                    border: Border.all(color: widget.color, width: 2.0),
                    borderRadius: BorderRadius.all(Radius.circular(16.0))),
                child: Text(
                  '${widget.display}',
                  style: TextStyle(
                      fontSize: 36.0,
                      color: widget.color,
                      fontWeight: FontWeight.w300),
                ),
              ),
            )));
  }
}

class ResultDisplay extends StatelessWidget {
  ResultDisplay({this.result});
  final String result;
  @override
  Widget build(BuildContext context) {
    return Text(
      '$result',
      softWrap: false,
      overflow: TextOverflow.fade,
      textScaleFactor: 7.5 / result.length > 1.0 ? 1.0 : 7.5 / result.length,
      style: TextStyle(
          fontSize: 80.0, fontWeight: FontWeight.w500, color: Colors.black),
    );
  }
}

class HistoryBlock extends StatelessWidget {
  HistoryBlock({this.result});
  final Result result;
  @override
  Widget build(BuildContext context) {
    var text = '';
    if (result.secondNum != null) {
      text = '${result.firstNum} ${result.oper.display} ${result.secondNum}';
    } else if (result.oper != null) {
      text = '${result.firstNum} ${result.oper.display} ?';
    } else if (result.firstNum != null) {
      text = '${result.firstNum}';
    }
    return Padding(
      padding: EdgeInsets.only(top: 16.0, bottom: 16.0, right: 16.0),
      child: Container(
        padding: EdgeInsets.all(16.0),
        alignment: Alignment.center,
        decoration: BoxDecoration(
            color: result.oper != null ? result.oper.color : Colors.white54,
            borderRadius: BorderRadius.all(Radius.circular(16.0))),
        child:
            Text(text, style: TextStyle(fontSize: 30.0, color: Colors.black54)),
      ),
    );
  }
}

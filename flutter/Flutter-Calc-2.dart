// Here's the code that works with Flutter2.0
import 'package:flutter/material.dart';

void main() {
    runApp(new MaterialApp(
        title: 'Tip Calculator',
        home: new TipCalculator()
    ));
}

class TipCalculator extends StatefulWidget {
    @override
    TipCalculatorState createState() {
        return new TipCalculatorState();
    }
}

class TipCalculatorState extends State<tipcalculator> {
    double billAmount = 0.0;
    double tipPercentage = 0.0;

    @override
    Widget build(BuildContext context) {
        TextField billAmountField = new TextField(
            keyboardType: TextInputType.number,
            decoration: new InputDecoration(labelText: "Bill amount(\$)"),
            onChanged: (String value) {
                try {
                    billAmount = double.parse(value.toString());
                } catch (exception) {
                    billAmount = 0.0;
                }
            }
        );

        // Create another input field
        TextField tipPercentageField = new TextField(
            keyboardType: TextInputType.number,
            decoration: new InputDecoration(labelText: "Tip %", hintText: "15" ),
            onChanged: (String value) {
                try {
                    tipPercentage = double.parse(value.toString());
                } catch (exception) {
                    tipPercentage = 0.0;
                }
            }
        );

        // Create button
        RaisedButton calculateButton = new RaisedButton(
            child: new Text("Calculate"),
            onPressed: () {
                // Calculate tip and total
                double calculatedTip = billAmount * tipPercentage / 100.0;
                double total = billAmount + calculatedTip;

                // Generate dialog
                AlertDialog dialog = new AlertDialog(
                    content: new Text("Tip: \$$calculatedTip \n"
                        "Total: \$$total")
                    );

                // Show dialog
                showDialog(context: context, child: dialog); // More code goes here
            }
        );

        Container container = new Container(
            padding: const EdgeInsets.all(16.0),
            child: new Column(
                children: [ billAmountField, tipPercentageField, calculateButton ]
            )
        );

        AppBar appBar = new AppBar(title: new Text("Tip Calculator"));

        Scaffold scaffold = new Scaffold(appBar: appBar, body: container);

        return scaffold;
    }
}

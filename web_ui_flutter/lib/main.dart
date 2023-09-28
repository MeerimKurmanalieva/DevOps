import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp();
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'API Client',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final TextEditingController _nameController = TextEditingController();
  String _serverResponse = '';

  void sendRequest() {
    final String name = _nameController.text;
    setState(() {
      _serverResponse = 'Hello $name!';
    });
  }

  void createEC2Instance() {
    // Implement the logic to create an EC2 instance here
    // You can use a package like http or Dio to make API requests to your backend
  }

  void deleteEC2Instance() {
    // Implement the logic to delete an EC2 instance here
    // You can use a package like http or Dio to make API requests to your backend
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('API Client'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              'Name:',
              style: TextStyle(fontSize: 16.0),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextField(
                controller: _nameController,
                decoration: InputDecoration(
                  hintText: 'Enter your name',
                ),
              ),
            ),
            ElevatedButton(
              onPressed: sendRequest,
              child: Text('Submit'),
            ),
            SizedBox(height: 20.0),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                ElevatedButton(
                  onPressed: createEC2Instance,
                  child: Text('Create EC2 Instance'),
                ),
                SizedBox(width: 20.0),
                ElevatedButton(
                  onPressed: deleteEC2Instance,
                  child: Text('Delete EC2 Instance'),
                ),
              ],
            ),
            SizedBox(height: 20.0),
            Text(
              'Server Response:',
              style: TextStyle(fontSize: 16.0),
            ),
            Text(
              _serverResponse,
              style: TextStyle(fontSize: 20.0),
            ),
          ],
        ),
      ),
    );
  }
}

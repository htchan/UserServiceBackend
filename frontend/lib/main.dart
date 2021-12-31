import 'package:flutter/material.dart';
import 'package:flutter_web_plugins/flutter_web_plugins.dart';

import './userChosenPage.dart';
import './loginPage.dart';

class Strategy extends HashUrlStrategy {
  final PlatformLocation _platformLocation;

  Strategy([
    PlatformLocation _platformLocation = const BrowserPlatformLocation(),
  ]) : _platformLocation = _platformLocation,
      super(_platformLocation);

  @override
  String prepareExternalUrl(String internalUrl) {
    return internalUrl.isEmpty
      ? '${_platformLocation.pathname}${_platformLocation.search}'
      : '$internalUrl';
  }

  @override
  String getPath() {
    String path = _platformLocation.pathname + _platformLocation.search;
    if (!_platformLocation.hash.startsWith('#/')) {
      path += _platformLocation.hash;
    }
    return path;
  }
}

void main() {
  // setPathUrlStrategy();
  setUrlStrategy(Strategy());
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  String url = 'http://localhost:8080/api/users';

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'UserService',
      theme: ThemeData(
        textTheme: Theme.of(context).textTheme.apply(
          fontSizeFactor: 1.25,
        ),
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      initialRoute: '/',
      onGenerateRoute: (settings) {
        var uri = Uri.parse(settings.name??"");
        String service = uri.queryParameters["service"]??"";
        if (uri.path == LoginPage.route) {
          return MaterialPageRoute(
            builder: (context) => LoginPage(url, service),
            settings: settings,
          );
        } else if (uri.path == UserChosenPage.route) {
          return MaterialPageRoute(
            builder: (context) => UserChosenPage(url, service),
            settings: settings
          );
        } else {
          return MaterialPageRoute(
            builder: (context) => UserChosenPage(url, service),
            settings: settings
          );
        }
      },
    );
  }
}

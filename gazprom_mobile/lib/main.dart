import 'package:flutter/material.dart';
import 'package:webview_flutter/webview_flutter.dart';

import 'ui/core/themes/colors.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Gazprom mobile',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: const WebViewScreen(title: 'GazpromBank'),
    );
  }
}

class WebViewScreen extends StatefulWidget {
  final String title;
  const WebViewScreen({super.key, required this.title});

  @override
  State<WebViewScreen> createState() => _WebViewScreenState();
}

class _WebViewScreenState extends State<WebViewScreen> {
  late final WebViewController _controller;

  @override
  void initState() {
    super.initState();
    _controller = WebViewController()
      ..setJavaScriptMode(JavaScriptMode.unrestricted)
      ..loadRequest(Uri.parse('https://web24.team'));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // appBar: AppBar(
      //     title: Text(widget.title, style: TextStyle(color: AppColors.white)),
      //     backgroundColor: AppColors.lightBlueAccent,
      //   ),
      body: WebViewWidget(controller: _controller),
    );
  }
}

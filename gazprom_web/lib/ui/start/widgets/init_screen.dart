import 'package:flutter/material.dart';

import 'start_screen.dart';

class InitScreen extends StatefulWidget {
  const InitScreen({super.key});

  @override
  State<InitScreen> createState() => _InitScreenState();
}

class _InitScreenState extends State<InitScreen> {
  late Widget currentScreen;

  @override
  void initState() {
    super.initState();
    currentScreen = StartScreen(
      onTap: (newScreen) {
        setState(() {
          currentScreen = newScreen;
        });
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return currentScreen;
  }
}
import 'package:flutter/material.dart';

abstract final class AppColors {
  static const Color white = Colors.white;
  static const Color lightBlueAccent = Colors.lightBlueAccent;
  static const Color transparent = Colors.transparent;
}

class Background extends StatelessWidget {
  final Widget? child;
  const Background({super.key, this.child});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
            gradient: LinearGradient(
              begin: Alignment.topCenter,
              end: Alignment.bottomCenter,
              colors: [
                Color(0xFF57FBFF),
                Color(0xFF1A1EEF),
              ],
              stops: [0.0, 1.0],
            ),
      ),
      child: child ?? child,
    );
  }
}
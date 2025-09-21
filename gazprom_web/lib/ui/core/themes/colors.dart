import 'package:flutter/material.dart';

abstract final class AppColors {
  static const Color white = Colors.white;
  static const Color white26 = Color.fromARGB(26, 255, 255, 255);
  static const Color black = Colors.black;
  static const Color black20 = Color.fromARGB(20, 0, 0, 0);
  static const Color blue1 = Color(0xFF1A1CEF);
  static const Color blue2 = Color(0xFF100DB5);
  static const Color transparent = Colors.transparent;

  static const Color unselectedItem = Color.fromARGB(255, 153, 153, 153);

  static const LinearGradient purpleGradient = LinearGradient(
    begin: Alignment.topCenter,
    end: Alignment.bottomCenter,
    colors: [Color(0xFFFF82BE), Color(0xFFDD42DB)],
    stops: [0.0, 1.0],
  );

  static const LinearGradient blackGradient = LinearGradient(
    begin: Alignment.bottomLeft,
    end: Alignment.topRight,
    colors: [Color(0xFF161313), Color(0xFF4C4C4C)],
    stops: [0.0, 1.0],
  );
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
          colors: [Color(0xFF57FBFF), Color(0xFF1A1EEF)],
          stops: [0.0, 1.0],
        ),
      ),
      child: child ?? child,
    );
  }
}

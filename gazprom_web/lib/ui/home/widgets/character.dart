import 'dart:math';
import 'dart:ui';
import 'package:flutter/material.dart';

class AnimatedSmiley extends StatefulWidget {
  final double size;
  const AnimatedSmiley({super.key, this.size = 200});

  @override
  State<AnimatedSmiley> createState() => _AnimatedSmileyState();
}

class _AnimatedSmileyState extends State<AnimatedSmiley>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(seconds: 3),
    )..repeat(reverse: true); // бесконечная анимация "туда-сюда"
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: _controller,
      builder: (_, __) {
        return CustomPaint(
          size: Size.square(widget.size),
          painter: _SmileyPainter(animationValue: _controller.value),
        );
      },
    );
  }
}

class _SmileyPainter extends CustomPainter {
  final double animationValue;
  _SmileyPainter({required this.animationValue});

  @override
  void paint(Canvas canvas, Size size) {
    final center = Offset(size.width / 2, size.height / 2);
    final radius = size.shortestSide / 2;

    // 1) Лицо
    final facePaint = Paint()..color = const Color(0xFFFFE066);
    canvas.drawCircle(center, radius, facePaint);

    // 2) Обводка
    final stroke = Paint()
      ..color = Colors.orange.shade700
      ..style = PaintingStyle.stroke
      ..strokeWidth = radius * 0.05;
    canvas.drawCircle(center, radius - stroke.strokeWidth / 2, stroke);

    // 3) Глаза
    final eyePaint = Paint()..color = Colors.black;
    final eyeOffsetX = radius * 0.45;
    final eyeOffsetY = radius * 0.25;
    final eyeRadius = radius * 0.15;

    final leftEye = Offset(center.dx - eyeOffsetX, center.dy - eyeOffsetY);
    final rightEye = Offset(center.dx + eyeOffsetX, center.dy - eyeOffsetY);

    canvas.drawCircle(leftEye, eyeRadius, eyePaint);
    canvas.drawCircle(rightEye, eyeRadius, eyePaint);

    // 4) Зрачки (двигаются по синусоиде)
    final pupilPaint = Paint()..color = Colors.white;
    final pupilOffset = Offset(
      sin(animationValue * 2 * pi) * eyeRadius * 0.4,
      cos(animationValue * 2 * pi) * eyeRadius * 0.2,
    );
    final pupilRadius = eyeRadius * 0.5;

    canvas.drawCircle(leftEye + pupilOffset, pupilRadius, pupilPaint);
    canvas.drawCircle(rightEye + pupilOffset, pupilRadius, pupilPaint);

    // 5) Улыбка (анимируем кривую)
    final mouthPaint = Paint()
      ..color = Colors.black
      ..style = PaintingStyle.stroke
      ..strokeWidth = radius * 0.08
      ..strokeCap = StrokeCap.round;

    final mouthWidth = radius * 1.1;
    final mouthTop = center.dy + radius * 0.2;
    final mouthLeft = Offset(center.dx - mouthWidth / 2, mouthTop);
    final mouthRight = Offset(center.dx + mouthWidth / 2, mouthTop);

    // animationValue (0..1) управляет "глубиной улыбки"
    final smileDepth = lerpDouble(0.2, 0.7, animationValue)!;
    final controlPoint = Offset(center.dx, center.dy + radius * smileDepth);

    final path = Path()
      ..moveTo(mouthLeft.dx, mouthLeft.dy)
      ..quadraticBezierTo(
          controlPoint.dx, controlPoint.dy, mouthRight.dx, mouthRight.dy);

    canvas.drawPath(path, mouthPaint);
  }

  @override
  bool shouldRepaint(covariant _SmileyPainter oldDelegate) =>
      oldDelegate.animationValue != animationValue;
}

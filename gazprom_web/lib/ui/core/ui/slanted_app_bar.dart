import 'package:flutter/material.dart';

import '../themes/colors.dart';

class SlantedAppBar extends StatelessWidget implements PreferredSizeWidget {
  final String title;

  const SlantedAppBar({super.key, required this.title});

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        // фон AppBar
        Container(
          color: AppColors.transparent,
          height: preferredSize.height,
        ),

        // сама линия
        Positioned.fill(
          child: CustomPaint(
            painter: _SlantedLinePainter(),
          ),
        ),

        // контент AppBar (текст, кнопки и т.д.)
        Positioned(
          left: 16,
          bottom: 12,
          child: Text(
            title,
            style: const TextStyle(color: Colors.white, fontSize: 20),
          ),
        ),
      ],
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(80);
}

class _SlantedLinePainter extends CustomPainter {
  @override
  void paint(Canvas canvas, Size size) {
    final paint = Paint()
      ..color = Colors.white
      ..strokeWidth = 3
      ..style = PaintingStyle.stroke;

    final path = Path();

    // начинаем справа снизу
    path.moveTo(size.width, size.height);

    // линия горизонтально до 4/5 ширины
    final cutPoint = size.width * 0.2; // 20% от левого края
    path.lineTo(cutPoint, size.height);

    // от cutPoint вниз под 45° к левому краю
    path.lineTo(0, size.height + cutPoint);

    canvas.drawPath(path, paint);
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) => false;
}

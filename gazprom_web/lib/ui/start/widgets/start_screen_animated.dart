import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:google_fonts/google_fonts.dart';
import 'dart:math';

import '../../core/themes/colors.dart';
import '../../core/themes/theme.dart';
import '../../core/ui/svg_picture_widget.dart';
import 'welcome_screen.dart';

class StartScreenAnimated extends StatefulWidget {
  const StartScreenAnimated({super.key});

  @override
  State<StartScreenAnimated> createState() => _StartScreenAnimatedState();
}

class _StartScreenAnimatedState extends State<StartScreenAnimated>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  late Animation<double> _sizeAnimation;
  late Animation<Offset> _positionAnimation;
  late Animation<double> _rotationAnimation;

  @override
  void initState() {
    super.initState();

    _controller = AnimationController(
      duration: const Duration(seconds: 2),
      vsync: this,
    );

    // Размер (увеличение в 2 раза)
    _sizeAnimation = Tween<double>(
      begin: 100,
      end: 2000,
    ).animate(CurvedAnimation(parent: _controller, curve: Curves.easeInOut));

    // Позиция (с верхнего левого в центр)
    _positionAnimation = Tween<Offset>(
      begin: const Offset(0, 0),
      end: const Offset(0.5, 0.2), // относительное смещение
    ).animate(CurvedAnimation(parent: _controller, curve: Curves.easeInOut));

    // Поворот (на -45 градусов)
    _rotationAnimation = Tween<double>(
      begin: 0,
      end: -pi / 4,
    ).animate(CurvedAnimation(parent: _controller, curve: Curves.easeInOut));

    _controller.forward();

    _controller.addStatusListener((status) {
      if (status == AnimationStatus.completed) {
        context.go('/welcome');
      }
    });
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          // Анимация логотипа
          AnimatedBuilder(
            animation: _controller,
            builder: (_, child) {
              final screenSize = MediaQuery.of(context).size;
              return Positioned(
                left:
                    _positionAnimation.value.dx *
                    (screenSize.width - _sizeAnimation.value),
                top:
                    _positionAnimation.value.dy *
                    (screenSize.height - _sizeAnimation.value),
                child: Transform.rotate(
                  angle: _rotationAnimation.value,
                  child: SvgPictureWidget(
                    path: 'assets/other/logo_1.svg',
                    width: _sizeAnimation.value,
                    height: _sizeAnimation.value * (52 / 76),
                    color: AppColors.black,
                  ),
                ),
              );
            },
          ),

          // Второй логотип рядом
          Positioned(
            top: 50,
            left: 100,
            child: SvgPictureWidget(
              path: 'assets/other/logo_2.svg',
              width: 269.4,
              height: 50.17,
              color: AppColors.black,
            ),
          ),

          // Тексты снизу
          Align(
            alignment: Alignment.bottomCenter,
            child: Padding(
              padding: const EdgeInsets.only(bottom: 50),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text(
                    'GOTOCASH',
                    style: GoogleFonts.righteous(
                      fontSize: 24,
                      color: AppColors.black,
                      fontWeight: AppFontWeights.bolt,
                    ),
                  ),
                  Text(
                    'by',
                    style: GoogleFonts.righteous(
                      fontSize: 18,
                      color: AppColors.black,
                      fontWeight: AppFontWeights.bolt300,
                    ),
                  ),
                  Text(
                    'SWIFTLY',
                    style: GoogleFonts.righteous(
                      fontSize: 20,
                      color: AppColors.black,
                      fontWeight: AppFontWeights.bolt300,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

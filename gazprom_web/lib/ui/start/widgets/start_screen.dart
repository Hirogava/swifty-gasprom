import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/core/ui/button_widget.dart';
import 'package:gazprom_web/ui/core/ui/svg_picture_widget.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../core/themes/colors.dart';
import '../../core/themes/theme.dart';

class StartScreen extends StatelessWidget {
  final ValueChanged<Widget>? onTap;
  const StartScreen({super.key, this.onTap});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.white,
      body: SafeArea(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.symmetric(
                horizontal: 16,
                vertical: 100,
              ),
              child: Align(
                alignment: Alignment.topLeft,
                child: SvgPictureWidget(
                  path: 'assets/other/card_logo.svg',
                  width: 216,
                  height: 31,
                  color: AppColors.black,
                ),
              ),
            ),

            Expanded(
              child: Center(
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
                    const SizedBox(height: 20),
                    ButtonWidget(
                      buttonText: 'Начать',
                      textColor: AppColors.black,
                      backgroundColor: AppColors.transparent,
                      borderColor: AppColors.black,
                      width: 138,
                      height: 45,
                      onTap: onTap,
                    ),
                  ],
                ),
              ),
            ),

            Padding(
              padding: const EdgeInsets.only(bottom: 20),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  SvgPictureWidget(
                    path: 'assets/other/mind.svg',
                    width: 52,
                    height: 59,
                    color: AppColors.black,
                  ),
                  SvgPictureWidget(
                    path: 'assets/other/play.svg',
                    width: 59,
                    height: 37,
                    color: AppColors.black,
                  ),
                  SvgPictureWidget(
                    path: 'assets/other/knowlegs.svg',
                    width: 62,
                    height: 52,
                    color: AppColors.black,
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
